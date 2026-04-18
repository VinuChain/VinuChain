package ethapi

import (
	"context"
	"math/big"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/txtrace"
)

// stubBackend implements the full Backend interface with no-op stubs.
// Tests embed this and override only the methods they need.
type stubBackend struct {
	extRPC bool
}

func (b *stubBackend) ExtRPCEnabled() bool                             { return b.extRPC }
func (b *stubBackend) Progress() PeerProgress                          { return PeerProgress{} }
func (b *stubBackend) SuggestGasTipCap(_ context.Context, _ uint64) *big.Int { return new(big.Int) }
func (b *stubBackend) EffectiveMinGasPrice(_ context.Context) *big.Int { return new(big.Int) }
func (b *stubBackend) ChainDb() ethdb.Database                         { return nil }
func (b *stubBackend) AccountManager() *accounts.Manager {
	return accounts.NewManager(&accounts.Config{})
}
func (b *stubBackend) RPCGasCap() uint64        { return 0 }
func (b *stubBackend) RPCTxFeeCap() float64     { return 0 }
func (b *stubBackend) UnprotectedAllowed() bool { return false }
func (b *stubBackend) CalcBlockExtApi() bool    { return false }

func (b *stubBackend) HeaderByNumber(_ context.Context, _ rpc.BlockNumber) (*evmcore.EvmHeader, error) {
	return nil, nil
}
func (b *stubBackend) HeaderByHash(_ context.Context, _ common.Hash) (*evmcore.EvmHeader, error) {
	return nil, nil
}
func (b *stubBackend) BlockByNumber(_ context.Context, _ rpc.BlockNumber) (*evmcore.EvmBlock, error) {
	return nil, nil
}
func (b *stubBackend) StateAndHeaderByNumberOrHash(_ context.Context, _ rpc.BlockNumberOrHash) (*state.StateDB, *evmcore.EvmHeader, error) {
	return nil, nil, nil
}
func (b *stubBackend) ResolveRpcBlockNumberOrHash(_ context.Context, _ rpc.BlockNumberOrHash) (idx.Block, error) {
	return 0, nil
}
func (b *stubBackend) BlockByHash(_ context.Context, _ common.Hash) (*evmcore.EvmBlock, error) {
	return nil, nil
}
func (b *stubBackend) GetReceiptsByNumber(_ context.Context, _ rpc.BlockNumber) (types.Receipts, error) {
	return nil, nil
}
func (b *stubBackend) GetTd(_ common.Hash) *big.Int { return nil }
func (b *stubBackend) GetEVM(_ context.Context, _ evmcore.Message, _ *state.StateDB, _ *evmcore.EvmHeader, _ *vm.Config) (*vm.EVM, func() error, error) {
	return nil, nil, nil
}
func (b *stubBackend) MinGasPrice() *big.Int { return new(big.Int) }
func (b *stubBackend) MaxGasLimit() uint64   { return 0 }
func (b *stubBackend) SendTx(_ context.Context, _ *types.Transaction) error { return nil }
func (b *stubBackend) GetTransaction(_ context.Context, _ common.Hash) (*types.Transaction, uint64, uint64, error) {
	return nil, 0, 0, nil
}
func (b *stubBackend) GetPoolTransactions() (types.Transactions, error)    { return nil, nil }
func (b *stubBackend) GetPoolTransaction(_ common.Hash) *types.Transaction { return nil }
func (b *stubBackend) GetPoolNonce(_ context.Context, _ common.Address) (uint64, error) {
	return 0, nil
}
func (b *stubBackend) Stats() (int, int) { return 0, 0 }
func (b *stubBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return nil, nil
}
func (b *stubBackend) TxPoolContentFrom(_ common.Address) (types.Transactions, types.Transactions) {
	return nil, nil
}
func (b *stubBackend) SubscribeNewTxsNotify(_ chan<- evmcore.NewTxsNotify) event.Subscription {
	return nil
}
func (b *stubBackend) ChainConfig() *params.ChainConfig       { return &params.ChainConfig{ChainID: new(big.Int)} }
func (b *stubBackend) CurrentBlock() *evmcore.EvmBlock         { return nil }
func (b *stubBackend) GetEventPayload(_ context.Context, _ string) (*inter.EventPayload, error) {
	return nil, nil
}
func (b *stubBackend) GetEvent(_ context.Context, _ string) (*inter.Event, error) {
	return nil, nil
}
func (b *stubBackend) GetHeads(_ context.Context, _ rpc.BlockNumber) (hash.Events, error) {
	return nil, nil
}
func (b *stubBackend) CurrentEpoch(_ context.Context) idx.Epoch { return 0 }
func (b *stubBackend) SealedEpochTiming(_ context.Context) (inter.Timestamp, inter.Timestamp) {
	return 0, 0
}
func (b *stubBackend) GetEpochBlockState(_ context.Context, _ rpc.BlockNumber) (*iblockproc.BlockState, *iblockproc.EpochState, error) {
	return nil, nil, nil
}
func (b *stubBackend) GetDowntime(_ context.Context, _ idx.ValidatorID) (idx.Block, inter.Timestamp, error) {
	return 0, 0, nil
}
func (b *stubBackend) GetUptime(_ context.Context, _ idx.ValidatorID) (*big.Int, error) {
	return nil, nil
}
func (b *stubBackend) GetOriginatedFee(_ context.Context, _ idx.ValidatorID) (*big.Int, error) {
	return nil, nil
}
func (b *stubBackend) GetPaybackBalance(_ context.Context, _ common.Address, _ *rpc.BlockNumberOrHash) (*big.Int, error) {
	return new(big.Int), nil
}
func (b *stubBackend) GetBlockContext(_ *evmcore.EvmHeader) vm.BlockContext {
	return vm.BlockContext{}
}
func (b *stubBackend) TxTraceByHash(_ context.Context, _ common.Hash) (*[]txtrace.ActionTrace, error) {
	return nil, nil
}
func (b *stubBackend) TxTraceSave(_ context.Context, _ common.Hash, _ []byte) error {
	return nil
}
