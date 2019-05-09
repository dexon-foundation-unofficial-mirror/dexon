// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package lds implements the Light Dexon Subprotocol.
package lds

import (
	"fmt"
	"sync"
	"time"

	"github.com/dexon-foundation/dexon/accounts"
	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/common/hexutil"
	"github.com/dexon-foundation/dexon/consensus"
	"github.com/dexon-foundation/dexon/consensus/dexcon"
	"github.com/dexon-foundation/dexon/core"
	"github.com/dexon-foundation/dexon/core/bloombits"
	"github.com/dexon-foundation/dexon/core/rawdb"
	"github.com/dexon-foundation/dexon/core/types"
	"github.com/dexon-foundation/dexon/dex"
	"github.com/dexon-foundation/dexon/dex/downloader"
	"github.com/dexon-foundation/dexon/eth/filters"
	"github.com/dexon-foundation/dexon/eth/gasprice"
	"github.com/dexon-foundation/dexon/event"
	"github.com/dexon-foundation/dexon/internal/ethapi"
	"github.com/dexon-foundation/dexon/light"
	"github.com/dexon-foundation/dexon/log"
	"github.com/dexon-foundation/dexon/node"
	"github.com/dexon-foundation/dexon/p2p"
	"github.com/dexon-foundation/dexon/p2p/discv5"
	"github.com/dexon-foundation/dexon/params"
	"github.com/dexon-foundation/dexon/rpc"
)

type LightDexon struct {
	ldsCommons

	odr         *LdsOdr
	relay       *LdsTxRelay
	chainConfig *params.ChainConfig
	// Channel for shutting down the service
	shutdownChan chan bool

	// Handlers
	peers      *peerSet
	txPool     *light.TxPool
	blockchain *light.LightChain
	serverPool *serverPool
	reqDist    *requestDistributor
	retriever  *retrieveManager

	bloomRequests chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer  *core.ChainIndexer

	ApiBackend *LdsApiBackend

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	networkId     uint64
	netRPCService *ethapi.PublicNetAPI

	wg sync.WaitGroup
}

// TODO: implement this
type dummyGov struct{}

func (d *dummyGov) GetRoundHeight(round uint64) uint64 {
	return 0
}

func New(ctx *node.ServiceContext, config *dex.Config) (*LightDexon, error) {
	chainDb, err := dex.CreateDB(ctx, config, "lightchaindata")
	if err != nil {
		return nil, err
	}
	chainConfig, genesisHash, genesisErr := core.SetupGenesisBlock(chainDb, config.Genesis)
	if _, isCompat := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !isCompat {
		return nil, genesisErr
	}
	log.Info("Initialised chain configuration", "config", chainConfig)

	peers := newPeerSet()
	quitSync := make(chan struct{})

	ldex := &LightDexon{
		ldsCommons: ldsCommons{
			chainDb: chainDb,
			config:  config,
			iConfig: light.DefaultClientIndexerConfig,
		},
		chainConfig:    chainConfig,
		eventMux:       ctx.EventMux,
		peers:          peers,
		reqDist:        newRequestDistributor(peers, quitSync),
		accountManager: ctx.AccountManager,
		engine:         dexcon.New(),
		shutdownChan:   make(chan bool),
		networkId:      config.NetworkId,
		bloomRequests:  make(chan chan *bloombits.Retrieval),
		bloomIndexer:   dex.NewBloomIndexer(chainDb, params.BloomBitsBlocksClient, params.HelperTrieConfirmations),
	}

	ldex.relay = NewLdsTxRelay(peers, ldex.reqDist)
	ldex.serverPool = newServerPool(chainDb, quitSync, &ldex.wg)
	ldex.retriever = newRetrieveManager(peers, ldex.reqDist, ldex.serverPool)

	ldex.odr = NewLdsOdr(chainDb, light.DefaultClientIndexerConfig, ldex.retriever)
	ldex.chtIndexer = light.NewChtIndexer(chainDb, ldex.odr, params.CHTFrequencyClient, params.HelperTrieConfirmations)
	ldex.bloomTrieIndexer = light.NewBloomTrieIndexer(chainDb, ldex.odr, params.BloomBitsBlocksClient, params.BloomTrieFrequency)
	ldex.odr.SetIndexers(ldex.chtIndexer, ldex.bloomTrieIndexer, ldex.bloomIndexer)

	// Note: NewLightChain adds the trusted checkpoint so it needs an ODR with
	// indexers already set but not started yet
	if ldex.blockchain, err = light.NewLightChain(ldex.odr, ldex.chainConfig, ldex.engine); err != nil {
		return nil, err
	}
	// Note: AddChildIndexer starts the update process for the child
	ldex.bloomIndexer.AddChildIndexer(ldex.bloomTrieIndexer)
	ldex.chtIndexer.Start(ldex.blockchain)
	ldex.bloomIndexer.Start(ldex.blockchain)

	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		ldex.blockchain.SetHead(compat.RewindTo)
		rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}

	ldex.txPool = light.NewTxPool(ldex.chainConfig, ldex.blockchain, ldex.relay)
	if ldex.protocolManager, err = NewProtocolManager(ldex.chainConfig, light.DefaultClientIndexerConfig, true, config.NetworkId, ldex.eventMux, ldex.engine, ldex.peers, ldex.blockchain, nil, chainDb, ldex.odr, ldex.relay, ldex.serverPool, quitSync, &ldex.wg, &dummyGov{}); err != nil {
		return nil, err
	}
	ldex.ApiBackend = &LdsApiBackend{ldex, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.DefaultGasPrice
	}
	ldex.ApiBackend.gpo = gasprice.NewOracle(ldex.ApiBackend, gpoParams)
	return ldex, nil
}

func lesTopic(genesisHash common.Hash, protocolVersion uint) discv5.Topic {
	var name string
	switch protocolVersion {
	case lpv2:
		name = "LES2"
	default:
		panic(nil)
	}
	return discv5.Topic(name + "@" + common.Bytes2Hex(genesisHash.Bytes()[0:8]))
}

type LightDummyAPI struct{}

// Etherbase is the address that mining rewards will be send to
func (l *LightDummyAPI) Etherbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase)
func (l *LightDummyAPI) Coinbase() (common.Address, error) {
	return common.Address{}, fmt.Errorf("not supported")
}

// Hashrate returns the POW hashrate
func (l *LightDummyAPI) Hashrate() hexutil.Uint {
	return 0
}

// Mining returns an indication if this node is currently mining.
func (l *LightDummyAPI) Mining() bool {
	return false
}

// APIs returns the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (l *LightDexon) APIs() []rpc.API {
	return append(ethapi.GetAPIs(l.ApiBackend), []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   &LightDummyAPI{},
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(l.protocolManager.downloader, l.eventMux),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(l.ApiBackend, true),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   l.netRPCService,
			Public:    true,
		},
	}...)
}

func (l *LightDexon) ResetWithGenesisBlock(gb *types.Block) {
	l.blockchain.ResetWithGenesisBlock(gb)
}

func (l *LightDexon) BlockChain() *light.LightChain { return l.blockchain }
func (l *LightDexon) TxPool() *light.TxPool         { return l.txPool }
func (l *LightDexon) Engine() consensus.Engine      { return l.engine }
func (l *LightDexon) LdsVersion() int               { return int(ClientProtocolVersions[0]) }
func (l *LightDexon) Downloader() ethapi.Downloader { return l.protocolManager.downloader }
func (l *LightDexon) EventMux() *event.TypeMux      { return l.eventMux }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (l *LightDexon) Protocols() []p2p.Protocol {
	return l.makeProtocols(ClientProtocolVersions)
}

// Start implements node.Service, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (l *LightDexon) Start(srvr *p2p.Server) error {
	log.Warn("Light client mode is an experimental feature")
	l.startBloomHandlers(params.BloomBitsBlocksClient)
	l.netRPCService = ethapi.NewPublicNetAPI(srvr, l.networkId)
	// clients are searching for the first advertised protocol in the list
	protocolVersion := AdvertiseProtocolVersions[0]
	l.serverPool.start(srvr, lesTopic(l.blockchain.Genesis().Hash(), protocolVersion))
	l.protocolManager.Start(l.config.LightPeers)
	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Ethereum protocol.
func (l *LightDexon) Stop() error {
	l.odr.Stop()
	l.bloomIndexer.Close()
	l.chtIndexer.Close()
	l.blockchain.Stop()
	l.protocolManager.Stop()
	l.txPool.Stop()
	l.engine.Close()

	l.eventMux.Stop()

	time.Sleep(time.Millisecond * 200)
	l.chainDb.Close()
	close(l.shutdownChan)

	return nil
}
