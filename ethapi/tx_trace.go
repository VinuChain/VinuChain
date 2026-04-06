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
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/txtrace"
	"github.com/Fantom-foundation/go-opera/utils/signers/gsignercache"
)

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
	ctx context.Context, blockCtx vm.BlockContext, msg types.Message,
	state *state.StateDB, block *evmcore.EvmBlock, tx *types.Transaction, index uint64,
	status uint64, chainConfig *params.ChainConfig) (*[]txtrace.ActionTrace, error) {

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

	txContext := evmcore.NewEVMTxContext(msg)
	vmenv := vm.NewEVM(blockCtx, txContext, state, chainConfig, cfg)

	go func() {
		<-ctx.Done()
		vmenv.Cancel()
	}()

	gp := new(evmcore.GasPool).AddGas(msg.Gas())
	state.Prepare(tx.Hash(), int(index))
	result, err := evmcore.ApplyMessage(vmenv, msg, gp, big.NewInt(0))

	if result != nil {
		txTracer.SetGasUsed(result.UsedGas)
	}
	txTracer.ProcessTx()
	traceActions := txTracer.GetTraceActions()
	state.Finalise(true)
	if err != nil {
		errTrace := txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, err)
		at := []txtrace.ActionTrace{*errTrace}
		if status == 1 {
			log.Error("State mismatch replaying tx: failed but receipt shows success", "txHash", tx.Hash().String(), "err", err)
			return nil, fmt.Errorf("state mismatch replaying tx %s: failed but receipt shows success", tx.Hash().String())
		}
		return &at, nil
	}
	if vmenv.Cancelled() {
		log.Info("EVM was canceled due to timeout when replaying transaction", "txHash", tx.Hash().String())
		return nil, fmt.Errorf("timeout when replaying tx")
	}
	if result != nil && result.Err != nil {
		if len(*traceActions) == 0 {
			log.Error("error in result when replaying transaction", "txHash", tx.Hash().String(), "err", result.Err.Error())
			errTrace := txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, result.Err)
			return &[]txtrace.ActionTrace{*errTrace}, nil
		}
		if status == 1 {
			log.Error("State mismatch replaying tx: reverted but receipt shows success", "txHash", tx.Hash().String(), "err", result.Err)
			return nil, fmt.Errorf("state mismatch replaying tx %s: reverted but receipt shows success", tx.Hash().String())
		}
		return traceActions, nil
	}
	if status == 0 {
		log.Error("State mismatch replaying tx: succeeded but receipt shows failure", "txHash", tx.Hash().String())
		return nil, fmt.Errorf("state mismatch replaying tx %s: succeeded but receipt shows failure", tx.Hash().String())
	}
	return traceActions, nil
}

// traceBlock replays all transactions in the block and returns their traces.
func (s *PublicTxTraceAPI) traceBlock(ctx context.Context, block *evmcore.EvmBlock, txHash *common.Hash, traceIndex *[]hexutil.Uint) (*[]txtrace.ActionTrace, error) {
	var (
		blockNumber   int64
		parentBlockNr rpc.BlockNumber
	)

	if block != nil && block.NumberU64() > 0 {
		blockNumber = block.Number.Int64()
		parentBlockNr = rpc.BlockNumber(blockNumber - 1)
	} else {
		return nil, fmt.Errorf("invalid block for tracing")
	}

	if s.b.CurrentBlock().Number.Int64() < blockNumber {
		return nil, fmt.Errorf("invalid block %v for tracing, current block is %v", blockNumber, s.b.CurrentBlock())
	}

	callTrace := txtrace.CallTrace{
		Actions: make([]txtrace.ActionTrace, 0),
	}

	signer := gsignercache.Wrap(types.MakeSigner(s.b.ChainConfig(), block.Number))
	blockCtx := s.b.GetBlockContext(block.Header())

	allTxOK := true
	for _, tx := range block.Transactions {
		traces, err := s.b.TxTraceByHash(ctx, tx.Hash())
		if err == nil {
			if txHash == nil || *txHash == tx.Hash() {
				callTrace.AddTraces(traces, traceIndex)
				if txHash != nil {
					break
				}
			}
		} else {
			allTxOK = false
			break
		}
	}

	if !allTxOK {
		stateDB, _, err := s.b.StateAndHeaderByNumberOrHash(ctx, rpc.BlockNumberOrHash{BlockNumber: &parentBlockNr})
		if err != nil {
			return nil, fmt.Errorf("cannot get state for block %v, error: %v", block.NumberU64(), err.Error())
		}
		receipts, err := s.b.GetReceiptsByNumber(ctx, rpc.BlockNumber(blockNumber))
		if err != nil {
			log.Debug("Cannot get receipts for block", "block", blockNumber, "err", err.Error())
			return nil, fmt.Errorf("cannot get receipts for block %v, error: %v", block.NumberU64(), err.Error())
		}

		callTrace = txtrace.CallTrace{
			Actions: make([]txtrace.ActionTrace, 0),
		}

		for i, tx := range block.Transactions {
			if txHash == nil || *txHash == tx.Hash() {
				log.Info("Replaying transaction", "txHash", tx.Hash().String())
				index := uint64(i)
				msg, err := tx.AsMessage(signer, block.BaseFee)
				if err != nil {
					callTrace.AddTrace(txtrace.GetErrorTrace(block.Hash, *block.Number, nil, tx.To(), tx.Hash(), index, errors.New("not able to decode tx")))
					continue
				}
				from := msg.From()
				if tx.To() != nil && *tx.To() == sfc.ContractAddress {
					errTrace := txtrace.GetErrorTrace(block.Hash, *block.Number, &from, tx.To(), tx.Hash(), index, errors.New("sfc tx"))
					at := []txtrace.ActionTrace{*errTrace}
					callTrace.AddTrace(errTrace)
					if jsonTraceBytes, err := json.Marshal(&at); err == nil {
						if saveErr := s.b.TxTraceSave(ctx, tx.Hash(), jsonTraceBytes); saveErr != nil {
							log.Debug("Cannot save sfc tx trace", "txHash", tx.Hash().String(), "err", saveErr)
						}
					}
				} else {
					txTraces, err := s.traceTx(ctx, blockCtx, msg, stateDB, block, tx, index, receipts[i].Status, s.b.ChainConfig())
					if err != nil {
						log.Debug("Cannot get transaction trace", "txHash", tx.Hash().String(), "err", err.Error())
						callTrace.AddTrace(txtrace.GetErrorTraceFromMsg(&msg, block.Hash, *block.Number, tx.Hash(), index, err))
					} else {
						callTrace.AddTraces(txTraces, traceIndex)
						if jsonTraceBytes, marshalErr := json.Marshal(txTraces); marshalErr != nil {
							log.Error("Cannot marshal tx traces for storage", "txHash", tx.Hash().String(), "err", marshalErr)
						} else {
							s.b.TxTraceSave(ctx, tx.Hash(), jsonTraceBytes)
						}
					}
				}
				if txHash != nil {
					break
				}
			} else if txHash != nil {
				log.Info("Replaying transaction without trace", "txHash", tx.Hash().String())
				msg, err := tx.AsMessage(signer, block.BaseFee)
				if err != nil {
					return nil, fmt.Errorf("cannot decode tx %s: %w", tx.Hash().String(), err)
				}
				stateDB.Prepare(tx.Hash(), i)
				vmConfig := opera.DefaultVMConfig
				vmConfig.NoBaseFee = true
				vmConfig.Debug = false
				vmConfig.Tracer = nil
				vmenv := vm.NewEVM(blockCtx, evmcore.NewEVMTxContext(msg), stateDB, s.b.ChainConfig(), vmConfig)
				res, applyErr := evmcore.ApplyMessage(vmenv, msg, new(evmcore.GasPool).AddGas(msg.Gas()), big.NewInt(0))
				failed := applyErr != nil
				if applyErr != nil {
					log.Error("Cannot replay transaction", "txHash", tx.Hash().String(), "err", applyErr.Error())
				}
				if res != nil && res.Err != nil {
					failed = true
					log.Debug("Error replaying transaction", "txHash", tx.Hash().String(), "err", res.Err.Error())
				}
				stateDB.Finalise(true)
				if (failed && receipts[i].Status == 1) || (!failed && receipts[i].Status == 0) {
					log.Error("State mismatch replaying tx without trace", "txHash", tx.Hash().String())
					return nil, fmt.Errorf("state mismatch replaying tx %s", tx.Hash().String())
				}
			}
		}
	}

	if len(callTrace.Actions) == 0 {
		if traceIndex != nil || txHash != nil {
			return nil, nil
		}
		emptyTrace := txtrace.CallTrace{
			Actions: make([]txtrace.ActionTrace, 0),
		}
		blockTrace := txtrace.NewActionTrace(block.Hash, *block.Number, common.Hash{}, 0, "empty")
		txAction := txtrace.NewAddressAction(&common.Address{}, 0, []byte{}, nil, hexutil.Big{}, nil)
		blockTrace.Action = txAction
		blockTrace.Error = "Empty block"
		emptyTrace.AddTrace(blockTrace)
		return &emptyTrace.Actions, nil
	}

	return &callTrace.Actions, nil
}

// Block implements trace_block: returns all transaction traces in the given block.
func (s *PublicTxTraceAPI) Block(ctx context.Context, numberOrHash rpc.BlockNumberOrHash) (*[]txtrace.ActionTrace, error) {
	blockNr, _ := numberOrHash.Number()
	defer func(start time.Time) {
		log.Info("Executing trace_block call finished", "blockNr", blockNr.Int64(), "runtime", time.Since(start))
	}(time.Now())
	block, err := s.b.BlockByNumber(ctx, blockNr)
	if err != nil {
		log.Debug("Cannot get block from db", "blockNr", blockNr)
		return nil, err
	}
	return s.traceBlock(ctx, block, nil, nil)
}

// Transaction implements trace_transaction: returns traces for a single transaction.
func (s *PublicTxTraceAPI) Transaction(ctx context.Context, hash common.Hash) (*[]txtrace.ActionTrace, error) {
	defer func(start time.Time) {
		log.Info("Executing trace_transaction call finished", "txHash", hash.String(), "runtime", time.Since(start))
	}(time.Now())
	return s.traceTxHash(ctx, hash, nil)
}

// Get implements trace_get: returns the trace at the specified index position.
func (s *PublicTxTraceAPI) Get(ctx context.Context, hash common.Hash, traceIndex []hexutil.Uint) (*[]txtrace.ActionTrace, error) {
	defer func(start time.Time) {
		log.Info("Executing trace_get call finished", "txHash", hash.String(), "index", traceIndex, "runtime", time.Since(start))
	}(time.Now())
	return s.traceTxHash(ctx, hash, &traceIndex)
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
	callTrace := txtrace.CallTrace{
		Actions: make([]txtrace.ActionTrace, 0),
	}
	traces, err := s.b.TxTraceByHash(ctx, hash)
	if err == nil && len(*traces) > 0 {
		callTrace.AddTraces(traces, traceIndex)
	}
	if len(callTrace.Actions) != 0 {
		return &callTrace.Actions, nil
	}
	return s.traceBlock(ctx, block, &hash, traceIndex)
}

// FilterArgs represents the arguments for specifying trace targets.
type FilterArgs struct {
	FromAddress *[]common.Address      `json:"fromAddress"`
	ToAddress   *[]common.Address      `json:"toAddress"`
	FromBlock   *rpc.BlockNumberOrHash `json:"fromBlock"`
	ToBlock     *rpc.BlockNumberOrHash `json:"toBlock"`
	After       uint                   `json:"after"`
	Count       uint                   `json:"count"`
}

// maxFilterTraceAddresses is the maximum number of addresses allowed in a
// trace_filter fromAddress or toAddress list. Large lists cause unbounded
// map allocations before any block data is read.
const maxFilterTraceAddresses = 1000

// validateFilterArgs returns an error if any address list exceeds the cap.
func validateFilterArgs(args FilterArgs) error {
	if args.FromAddress != nil && len(*args.FromAddress) > maxFilterTraceAddresses {
		return fmt.Errorf("fromAddress list too large: %d addresses, maximum is %d", len(*args.FromAddress), maxFilterTraceAddresses)
	}
	if args.ToAddress != nil && len(*args.ToAddress) > maxFilterTraceAddresses {
		return fmt.Errorf("toAddress list too large: %d addresses, maximum is %d", len(*args.ToAddress), maxFilterTraceAddresses)
	}
	return nil
}

// Filter implements trace_filter: returns traces matching the filter criteria.
func (s *PublicTxTraceAPI) Filter(ctx context.Context, args FilterArgs) (*[]txtrace.ActionTrace, error) {
	if err := validateFilterArgs(args); err != nil {
		return nil, err
	}

	defer func(start time.Time) {
		var data []interface{}
		if args.FromBlock != nil && args.FromBlock.BlockNumber != nil {
			data = append(data, "fromBlock", args.FromBlock.BlockNumber.Int64())
		}
		if args.ToBlock != nil && args.ToBlock.BlockNumber != nil {
			data = append(data, "toBlock", args.ToBlock.BlockNumber.Int64())
		}
		data = append(data, "time", time.Since(start))
		log.Info("Executing trace_filter call finished", data...)
	}(time.Now())

	var (
		fromBlock, toBlock rpc.BlockNumber
		mainErr            error
	)
	if args.FromBlock != nil && args.FromBlock.BlockNumber != nil {
		fromBlock = *args.FromBlock.BlockNumber
		if fromBlock == rpc.LatestBlockNumber || fromBlock == rpc.PendingBlockNumber {
			fromBlock = rpc.BlockNumber(s.b.CurrentBlock().NumberU64())
		}
	}
	if args.ToBlock != nil && args.ToBlock.BlockNumber != nil {
		toBlock = *args.ToBlock.BlockNumber
		if toBlock == rpc.LatestBlockNumber || toBlock == rpc.PendingBlockNumber {
			toBlock = rpc.BlockNumber(s.b.CurrentBlock().NumberU64())
		}
	} else {
		toBlock = rpc.BlockNumber(s.b.CurrentBlock().NumberU64())
	}

	const maxFilterBlockRange = 1000
	if toBlock > fromBlock && uint64(toBlock-fromBlock) >= maxFilterBlockRange {
		return nil, fmt.Errorf("block range too large: %d blocks requested, maximum is %d", uint64(toBlock-fromBlock)+1, maxFilterBlockRange)
	}

	var traceAdded, traceCount uint
	var fromAddresses, toAddresses map[common.Address]struct{}
	if args.FromAddress != nil {
		fromAddresses = make(map[common.Address]struct{})
		for _, addr := range *args.FromAddress {
			fromAddresses[addr] = struct{}{}
		}
	}
	if args.ToAddress != nil {
		toAddresses = make(map[common.Address]struct{})
		for _, addr := range *args.ToAddress {
			toAddresses[addr] = struct{}{}
		}
	}

	callTrace := txtrace.CallTrace{
		Actions: make([]txtrace.ActionTrace, 0),
	}

	if args.Count == 0 {
		workerCount := runtime.NumCPU() / 2
		if workerCount < 1 {
			workerCount = 1
		}
		blocks := make(chan rpc.BlockNumber, workerCount)
		results := make(chan txtrace.ActionTrace, workerCount*64)

		// Producer runs in its own goroutine so workers breaking early
		// doesn't deadlock the caller.
		stopProducer := make(chan struct{})
		go func() {
			defer close(blocks)
			for i := fromBlock; i <= toBlock; i++ {
				select {
				case blocks <- i:
				case <-stopProducer:
					return
				}
			}
		}()

		var wg sync.WaitGroup
		for w := 0; w < workerCount; w++ {
			wg.Add(1)
			wId := w
			go func() {
				defer wg.Done()
				defer func() {
					if r := recover(); r != nil {
						log.Error("filterWorker panicked", "worker", wId, "recover", r)
					}
				}()
				filterWorker(wId, s, ctx, blocks, results, fromAddresses, toAddresses)
			}()
		}

		var wgResult sync.WaitGroup
		wgResult.Add(1)
		go func() {
			defer wgResult.Done()
			for trace := range results {
				callTrace.AddTrace(&trace)
			}
		}()

		wg.Wait()
		close(stopProducer)
		close(results)
		wgResult.Wait()

		sort.SliceStable(callTrace.Actions, func(i, j int) bool {
			bi := callTrace.Actions[i].BlockNumber.Uint64()
			bj := callTrace.Actions[j].BlockNumber.Uint64()
			if bi != bj {
				return bi < bj
			}
			return callTrace.Actions[i].TransactionPosition < callTrace.Actions[j].TransactionPosition
		})
	} else {
	blocks:
		for i := fromBlock; i <= toBlock; i++ {
			select {
			case <-ctx.Done():
				mainErr = ctx.Err()
				break blocks
			default:
			}
			block, err := s.b.BlockByNumber(ctx, i)
			if err != nil {
				mainErr = err
				break
			}
			if block != nil && block.Transactions.Len() > 0 {
				traces, err := s.traceBlock(ctx, block, nil, nil)
				if err != nil {
					mainErr = err
					break
				}
				for _, trace := range *traces {
					if args.Count == 0 || traceAdded < args.Count {
						addTrace := true
						if args.FromAddress != nil || args.ToAddress != nil {
							if args.FromAddress != nil {
								if trace.Action.From == nil {
									addTrace = false
								} else if _, ok := fromAddresses[*trace.Action.From]; !ok {
									addTrace = false
								}
							}
							if args.ToAddress != nil {
								if trace.Action.To == nil {
									addTrace = false
								} else if _, ok := toAddresses[*trace.Action.To]; !ok {
									addTrace = false
								}
							}
						}
						if addTrace {
							if traceCount >= args.After {
								callTrace.AddTrace(&trace)
								traceAdded++
							}
							traceCount++
						}
					} else {
						break blocks
					}
				}
			}
		}
	}

	if mainErr != nil {
		return nil, mainErr
	}

	return &callTrace.Actions, nil
}

func filterWorker(id int,
	s *PublicTxTraceAPI,
	ctx context.Context,
	blocks <-chan rpc.BlockNumber,
	results chan<- txtrace.ActionTrace,
	fromAddresses map[common.Address]struct{},
	toAddresses map[common.Address]struct{}) {

	for i := range blocks {
		block, err := s.b.BlockByNumber(ctx, i)
		if err != nil {
			break
		}
		if block != nil && block.Transactions.Len() > 0 {
			traces, err := s.traceBlock(ctx, block, nil, nil)
			if err != nil {
				break
			}
			for _, trace := range *traces {
				addTrace := true
				if len(fromAddresses) > 0 {
					if trace.Action.From == nil {
						addTrace = false
					} else if _, ok := fromAddresses[*trace.Action.From]; !ok {
						addTrace = false
					}
				}
				if len(toAddresses) > 0 {
					if trace.Action.To == nil {
						addTrace = false
					} else if _, ok := toAddresses[*trace.Action.To]; !ok {
						addTrace = false
					}
				}
				if addTrace {
					select {
					case results <- trace:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}
}
