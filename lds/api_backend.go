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

package lds

import (
	"context"
	"math/big"

	"github.com/dexon-foundation/dexon/accounts"
	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/common/math"
	"github.com/dexon-foundation/dexon/core"
	"github.com/dexon-foundation/dexon/core/bloombits"
	"github.com/dexon-foundation/dexon/core/rawdb"
	"github.com/dexon-foundation/dexon/core/state"
	"github.com/dexon-foundation/dexon/core/types"
	"github.com/dexon-foundation/dexon/core/vm"
	"github.com/dexon-foundation/dexon/eth/gasprice"
	"github.com/dexon-foundation/dexon/ethdb"
	"github.com/dexon-foundation/dexon/event"
	"github.com/dexon-foundation/dexon/internal/ethapi"
	"github.com/dexon-foundation/dexon/light"
	"github.com/dexon-foundation/dexon/params"
	"github.com/dexon-foundation/dexon/rpc"
)

type LdsApiBackend struct {
	dex *LightDexon
	gpo *gasprice.Oracle
}

func (b *LdsApiBackend) ChainConfig() *params.ChainConfig {
	return b.dex.chainConfig
}

func (b *LdsApiBackend) CurrentBlock() *types.Block {
	return types.NewBlockWithHeader(b.dex.BlockChain().CurrentHeader())
}

func (b *LdsApiBackend) SetHead(number uint64) {
	b.dex.protocolManager.downloader.Cancel()
	b.dex.blockchain.SetHead(number)
}

func (b *LdsApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	if blockNr == rpc.LatestBlockNumber || blockNr == rpc.PendingBlockNumber {
		return b.dex.blockchain.CurrentHeader(), nil
	}
	return b.dex.blockchain.GetHeaderByNumberOdr(ctx, uint64(blockNr))
}

func (b *LdsApiBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.dex.blockchain.GetHeaderByHash(hash), nil
}

func (b *LdsApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, err
	}
	return b.GetBlock(ctx, header.Hash())
}

func (b *LdsApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	return light.NewState(ctx, header, b.dex.odr), header, nil
}

func (b *LdsApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.dex.blockchain.GetBlockByHash(ctx, blockHash)
}

func (b *LdsApiBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	if number := rawdb.ReadHeaderNumber(b.dex.chainDb, hash); number != nil {
		return light.GetBlockReceipts(ctx, b.dex.odr, hash, *number)
	}
	return nil, nil
}

func (b *LdsApiBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	if number := rawdb.ReadHeaderNumber(b.dex.chainDb, hash); number != nil {
		return light.GetBlockLogs(ctx, b.dex.odr, hash, *number)
	}
	return nil, nil
}

func (b *LdsApiBackend) GetTd(hash common.Hash) *big.Int {
	return b.dex.blockchain.GetTdByHash(hash)
}

func (b *LdsApiBackend) GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header) (*vm.EVM, func() error, error) {
	state.SetBalance(msg.From(), math.MaxBig256)
	context := core.NewEVMContext(msg, header, b.dex.blockchain, nil)
	return vm.NewEVM(context, state, b.dex.chainConfig, vm.Config{}), state.Error, nil
}

func (b *LdsApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.dex.txPool.Add(ctx, signedTx)
}

func (b *LdsApiBackend) SendTxs(ctx context.Context, signedTxs []*types.Transaction) []error {
	b.dex.txPool.AddBatch(ctx, signedTxs)
	return nil
}

func (b *LdsApiBackend) RemoveTx(txHash common.Hash) {
	b.dex.txPool.RemoveTx(txHash)
}

func (b *LdsApiBackend) GetPoolTransactions() (types.Transactions, error) {
	return b.dex.txPool.GetTransactions()
}

func (b *LdsApiBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	return b.dex.txPool.GetTransaction(txHash)
}

func (b *LdsApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.dex.txPool.GetNonce(ctx, addr)
}

func (b *LdsApiBackend) Stats() (pending int, queued int) {
	return b.dex.txPool.Stats(), 0
}

func (b *LdsApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.dex.txPool.Content()
}

func (b *LdsApiBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.dex.txPool.SubscribeNewTxsEvent(ch)
}

func (b *LdsApiBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.dex.blockchain.SubscribeChainEvent(ch)
}

func (b *LdsApiBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.dex.blockchain.SubscribeChainHeadEvent(ch)
}

func (b *LdsApiBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.dex.blockchain.SubscribeChainSideEvent(ch)
}

func (b *LdsApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.dex.blockchain.SubscribeLogsEvent(ch)
}

func (b *LdsApiBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.dex.blockchain.SubscribeRemovedLogsEvent(ch)
}

func (b *LdsApiBackend) Downloader() ethapi.Downloader {
	return b.dex.Downloader()
}

func (b *LdsApiBackend) ProtocolVersion() int {
	return b.dex.LdsVersion() + 10000
}

func (b *LdsApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *LdsApiBackend) ChainDb() ethdb.Database {
	return b.dex.chainDb
}

func (b *LdsApiBackend) EventMux() *event.TypeMux {
	return b.dex.eventMux
}

func (b *LdsApiBackend) AccountManager() *accounts.Manager {
	return b.dex.accountManager
}

func (b *LdsApiBackend) RPCGasCap() *big.Int {
	return b.dex.config.RPCGasCap
}

func (b *LdsApiBackend) BloomStatus() (uint64, uint64) {
	if b.dex.bloomIndexer == nil {
		return 0, 0
	}
	sections, _, _ := b.dex.bloomIndexer.Sections()
	return params.BloomBitsBlocksClient, sections
}

func (b *LdsApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.dex.bloomRequests)
	}
}
