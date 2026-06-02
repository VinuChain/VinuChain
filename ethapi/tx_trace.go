// Copyright 2015 The go-ethereum Authors
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

package ethapi

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/txtrace"
	"github.com/Fantom-foundation/go-opera/utils/signers/gsignercache"
)

var zeroPayback = big.NewInt(0)

// PublicTxTraceAPI provides an API to access transaction tracing.
type PublicTxTraceAPI struct {
	b Backend
}

// NewPublicTxTraceAPI creates a new transaction trace API.
func NewPublicTxTraceAPI(b Backend) *PublicTxTraceAPI {
	return &PublicTxTraceAPI{b}
}

// traceTx executes a single transaction with the trace logger enabled.
func (s *PublicTxTraceAPI) traceTx(
	ctx context.Context, msg types.Message,
	state *state.StateDB, block *evmcore.EvmBlock, tx *types.Transaction, index uint64,
	receipt *types.Receipt) (*[]txtrace.ActionTrace, error) {

	cfg := opera.DefaultVMConfig
	cfg.Debug = true
	txTracer := txtrace.NewTraceStructLogger(nil)
	cfg.Tracer = txTracer
	cfg.NoBaseFee = true

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	txTracer.SetTx(tx.Hash())
	txTracer.SetFrom(msg.From())
	txTracer.SetTo(msg.To())
	txTracer.SetValue(*msg.Value())
	txTracer.SetBlockHash(block.Hash)
	txTracer.SetBlockNumber(block.Number)
	txTracer.SetTxIndex(uint(index))
	txTracer.SetGasUsed(tx.Gas())

	vmenv, _, err := s.b.GetEVM(ctx, msg, state, block.Header(), &cfg)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		vmenv.Cancel()
	}()

	gp := new(evmcore.GasPool).AddGas(msg.Gas())
	state.Prepare(tx.Hash(), int(index))
	result, err := evmcore.ApplyMessage(vmenv, msg, gp, replayPayback(receipt))

	if result != nil {
		txTracer.SetGasUsed(result.UsedGas)
	}
	txTracer.ProcessTx()
	traceActions := txTracer.GetTraceActions()
	state.Finalise(true)
	if err != nil {
		errTrace := txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, err)
		at := []txtrace.ActionTrace{*errTrace}
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Error("State mismatch replaying tx: failed but receipt shows success", "txHash", tx.Hash().String(), "err", err)
			return nil, fmt.Errorf("state mismatch replaying tx %s: failed but receipt shows success", tx.Hash().String())
		}
		return &at, nil
	}
	if vmenv.Cancelled() {
		log.Warn("EVM was canceled due to timeout when replaying transaction", "txHash", tx.Hash().String())
		return nil, fmt.Errorf("timeout when replaying tx")
	}
	if result != nil && result.Err != nil {
		if len(*traceActions) == 0 {
			log.Error("error in result when replaying transaction", "txHash", tx.Hash().String(), "err", result.Err.Error())
			errTrace := txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, result.Err)
			return &[]txtrace.ActionTrace{*errTrace}, nil
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			log.Error("State mismatch replaying tx: reverted but receipt shows success", "txHash", tx.Hash().String(), "err", result.Err)
			return nil, fmt.Errorf("state mismatch replaying tx %s: reverted but receipt shows success", tx.Hash().String())
		}
		return traceActions, nil
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Error("State mismatch replaying tx: succeeded but receipt shows failure", "txHash", tx.Hash().String())
		return nil, fmt.Errorf("state mismatch replaying tx %s: succeeded but receipt shows failure", tx.Hash().String())
	}
	return traceActions, nil
}

func replayPayback(receipt *types.Receipt) *big.Int {
	if receipt == nil || receipt.FeeRefund == nil {
		return zeroPayback
	}
	return new(big.Int).Set(receipt.FeeRefund)
}

func (s *PublicTxTraceAPI) replayTxWithoutTrace(ctx context.Context, msg types.Message, state *state.StateDB, block *evmcore.EvmBlock, tx *types.Transaction, index int, receipt *types.Receipt) error {
	state.Prepare(tx.Hash(), index)
	vmConfig := opera.DefaultVMConfig
	vmConfig.NoBaseFee = true
	vmConfig.Debug = false
	vmConfig.Tracer = nil
	vmenv, _, err := s.b.GetEVM(ctx, msg, state, block.Header(), &vmConfig)
	if err != nil {
		return err
	}
	res, applyErr := evmcore.ApplyMessage(vmenv, msg, new(evmcore.GasPool).AddGas(msg.Gas()), replayPayback(receipt))
	failed := applyErr != nil
	if applyErr != nil {
		log.Error("Cannot replay transaction", "txHash", tx.Hash().String(), "err", applyErr.Error())
	}
	if res != nil && res.Err != nil {
		failed = true
		log.Debug("Error replaying transaction", "txHash", tx.Hash().String(), "err", res.Err.Error())
	}
	state.Finalise(true)
	if (failed && receipt.Status == types.ReceiptStatusSuccessful) || (!failed && receipt.Status == types.ReceiptStatusFailed) {
		log.Error("State mismatch replaying tx without trace", "txHash", tx.Hash().String())
		return fmt.Errorf("state mismatch replaying tx %s", tx.Hash().String())
	}
	return nil
}

// traceBlock replays all transactions in the block and returns their traces.
func (s *PublicTxTraceAPI) traceBlock(ctx context.Context, block *evmcore.EvmBlock, txHash *common.Hash, traceIndex *[]hexutil.Uint) (*[]txtrace.ActionTrace, error) {
	var (
		blockNumber   int64
		parentBlockNr rpc.BlockNumber
	)

	if block == nil {
		return nil, fmt.Errorf("invalid block for tracing")
	}
	blockNumber = block.Number.Int64()

	callTrace := txtrace.CallTrace{
		Actions: make([]txtrace.ActionTrace, 0),
	}
	if block.NumberU64() == 0 {
		if len(block.Transactions) == 0 {
			return &callTrace.Actions, nil
		}
		return nil, fmt.Errorf("cannot trace genesis block with transactions")
	}
	parentBlockNr = rpc.BlockNumber(blockNumber - 1)

	if s.b.CurrentBlock().Number.Int64() < blockNumber {
		return nil, fmt.Errorf("invalid block %v for tracing, current block is %v", blockNumber, s.b.CurrentBlock())
	}

	signer := gsignercache.Wrap(types.MakeSigner(s.b.ChainConfig(), block.Number))
	stateDB, _, err := s.b.StateAndHeaderByNumberOrHash(ctx, rpc.BlockNumberOrHash{BlockNumber: &parentBlockNr})
	if err != nil {
		return nil, fmt.Errorf("cannot get state for block %v, error: %v", block.NumberU64(), err.Error())
	}
	receipts, err := s.b.GetReceiptsByNumber(ctx, rpc.BlockNumber(blockNumber))
	if err != nil {
		log.Debug("Cannot get receipts for block", "block", blockNumber, "err", err.Error())
		return nil, fmt.Errorf("cannot get receipts for block %v, error: %v", block.NumberU64(), err.Error())
	}
	if len(receipts) != len(block.Transactions) {
		return nil, fmt.Errorf("receipt count mismatch for block %d: %d txs but %d receipts", blockNumber, len(block.Transactions), len(receipts))
	}

	for i, tx := range block.Transactions {
		if txHash == nil || *txHash == tx.Hash() {
			log.Debug("Replaying transaction", "txHash", tx.Hash().String())
			index := uint64(i)
			msg, err := evmcore.TxAsMessage(tx, signer, block.BaseFee)
			if err != nil {
				callTrace.AddTrace(txtrace.GetErrorTrace(block.Hash, *block.Number, nil, tx.To(), tx.Hash(), index, errors.New("not able to decode tx")))
				continue
			}
			from := msg.From()
			if tx.To() != nil && *tx.To() == sfc.ContractAddress {
				if err := s.replayTxWithoutTrace(ctx, msg, stateDB, block, tx, i, receipts[i]); err != nil {
					return nil, err
				}
				callTrace.AddTrace(txtrace.GetErrorTrace(block.Hash, *block.Number, &from, tx.To(), tx.Hash(), index, errors.New("sfc tx")))
			} else {
				txTraces, err := s.traceTx(ctx, msg, stateDB, block, tx, index, receipts[i])
				if err != nil {
					log.Debug("Cannot get transaction trace", "txHash", tx.Hash().String(), "err", err.Error())
					callTrace.AddTrace(txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, err))
				} else {
					callTrace.AddTraces(txTraces, traceIndex)
				}
			}
			if txHash != nil {
				break
			}
		} else {
			log.Debug("Replaying transaction without trace", "txHash", tx.Hash().String())
			msg, err := evmcore.TxAsMessage(tx, signer, block.BaseFee)
			if err != nil {
				return nil, fmt.Errorf("cannot decode tx %s: %w", tx.Hash().String(), err)
			}
			if err := s.replayTxWithoutTrace(ctx, msg, stateDB, block, tx, i, receipts[i]); err != nil {
				return nil, err
			}
		}
	}

	if len(callTrace.Actions) == 0 {
		return &callTrace.Actions, nil
	}

	return &callTrace.Actions, nil
}

// Block implements trace_block: returns all transaction traces in the given block.
func (s *PublicTxTraceAPI) Block(ctx context.Context, numberOrHash rpc.BlockNumberOrHash) (*[]txtrace.ActionTrace, error) {
	blockNr, _ := numberOrHash.Number()
	defer func(start time.Time) {
		log.Debug("Executing trace_block call finished", "blockNr", blockNr.Int64(), "runtime", time.Since(start))
	}(time.Now())

	var (
		block *evmcore.EvmBlock
		err   error
	)
	if blockHash, ok := numberOrHash.Hash(); ok {
		block, err = s.b.BlockByHash(ctx, blockHash)
	} else {
		block, err = s.b.BlockByNumber(ctx, blockNr)
	}
	if err != nil {
		log.Debug("Cannot get block from db", "blockNr", blockNr)
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block not found")
	}
	return s.traceBlock(ctx, block, nil, nil)
}

// Transaction implements trace_transaction: returns traces for a single transaction.
func (s *PublicTxTraceAPI) Transaction(ctx context.Context, hash common.Hash) (*[]txtrace.ActionTrace, error) {
	defer func(start time.Time) {
		log.Debug("Executing trace_transaction call finished", "txHash", hash.String(), "runtime", time.Since(start))
	}(time.Now())
	return s.traceTxHash(ctx, hash, nil)
}

// Get implements trace_get: returns the trace at the specified index position.
func (s *PublicTxTraceAPI) Get(ctx context.Context, hash common.Hash, traceIndex []hexutil.Uint) (*txtrace.ActionTrace, error) {
	defer func(start time.Time) {
		log.Debug("Executing trace_get call finished", "txHash", hash.String(), "index", traceIndex, "runtime", time.Since(start))
	}(time.Now())
	traces, err := s.traceTxHash(ctx, hash, &traceIndex)
	if err != nil || traces == nil || len(*traces) == 0 {
		return nil, err
	}
	return &(*traces)[0], nil
}

func (s *PublicTxTraceAPI) traceTxHash(ctx context.Context, hash common.Hash, traceIndex *[]hexutil.Uint) (*[]txtrace.ActionTrace, error) {
	tx, blockNumber, _, _ := s.b.GetTransaction(ctx, hash)
	if tx == nil {
		return nil, fmt.Errorf("transaction not found: %s", hash.Hex())
	}
	blkNr := rpc.BlockNumber(blockNumber)
	block, err := s.b.BlockByNumber(ctx, blkNr)
	if err != nil {
		log.Debug("Cannot get block from db", "blockNr", blkNr)
		return nil, err
	}
	return s.traceBlock(ctx, block, &hash, traceIndex)
}
