// Copyright 2018 The dexon-consensus Authors
// This file is part of the dexon-consensus library.
//
// The dexon-consensus library is free software: you can redistribute it
// and/or modify it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The dexon-consensus library is distributed in the hope that it will be
// useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the dexon-consensus library. If not, see
// <http://www.gnu.org/licenses/>.

package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dexon-foundation/dexon/dex/consensus/common"
	"github.com/dexon-foundation/dexon/dex/consensus/core"
	"github.com/dexon-foundation/dexon/dex/consensus/core/crypto"
	"github.com/dexon-foundation/dexon/dex/consensus/core/db"
	"github.com/dexon-foundation/dexon/dex/consensus/core/test"
	"github.com/dexon-foundation/dexon/dex/consensus/core/types"
	"github.com/dexon-foundation/dexon/dex/consensus/core/utils"
)

type node struct {
	ID      types.NodeID
	con     *core.Consensus
	app     *test.App
	gov     *test.Governance
	rEvt    *utils.RoundEvent
	db      db.Database
	network *test.Network
	logger  common.Logger
}

func setupNodes(
	dMoment time.Time,
	prvKeys []crypto.PrivateKey,
	seedGov *test.Governance) (nodes map[types.NodeID]*node, err error) {
	var (
		wg        sync.WaitGroup
		initRound uint64
	)
	// Setup peer server at transport layer.
	server := test.NewFakeTransportServer()
	serverChannel, err := server.Host()
	if err != nil {
		return
	}
	// setup nodes.
	nodes = make(map[types.NodeID]*node)
	wg.Add(len(prvKeys))
	for i, k := range prvKeys {
		var (
			dbInst db.Database
			f      *os.File
			rEvt   *utils.RoundEvent
		)
		dbInst, err = db.NewMemBackedDB()
		if err != nil {
			return
		}
		// Prepare essential modules: app, gov, db.
		networkModule := test.NewNetwork(k.PublicKey(), test.NetworkConfig{
			Type:          test.NetworkTypeFake,
			DirectLatency: &test.FixedLatencyModel{},
			GossipLatency: &test.FixedLatencyModel{},
			Marshaller:    test.NewDefaultMarshaller(nil)},
		)
		gov := seedGov.Clone()
		gov.SwitchToRemoteMode(networkModule)
		gov.NotifyRound(initRound, types.GenesisHeight)
		networkModule.AttachNodeSetCache(utils.NewNodeSetCache(gov))
		f, err = os.Create(fmt.Sprintf("log.%d.log", i))
		if err != nil {
			panic(err)
		}
		logger := common.NewCustomLogger(
			log.New(f, "", log.LstdFlags|log.Lmicroseconds))
		rEvt, err = utils.NewRoundEvent(context.Background(), gov, logger,
			types.Position{Height: types.GenesisHeight}, core.ConfigRoundShift)
		if err != nil {
			return
		}
		nID := types.NewNodeID(k.PublicKey())
		nodes[nID] = &node{
			ID:      nID,
			app:     test.NewApp(initRound+1, gov, rEvt),
			gov:     gov,
			db:      dbInst,
			logger:  logger,
			rEvt:    rEvt,
			network: networkModule,
		}
		go func() {
			defer wg.Done()
			if err := networkModule.Setup(serverChannel); err != nil {
				panic(err)
			}
			go networkModule.Run()
		}()
	}
	// Make sure transport layer is ready.
	if err = server.WaitForPeers(uint32(len(prvKeys))); err != nil {
		return
	}
	wg.Wait()
	for _, k := range prvKeys {
		node := nodes[types.NewNodeID(k.PublicKey())]
		// Now is the consensus module.
		node.con = core.NewConsensus(
			dMoment,
			node.app,
			node.gov,
			node.db,
			node.network,
			k,
			node.logger,
		)
	}
	return
}
