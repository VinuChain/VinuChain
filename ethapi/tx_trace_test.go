package ethapi

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/txtrace"
)

// blockErrBackend is a minimal Backend stub that returns a fixed error from
// BlockByNumber and panics if any other method is called. It is only safe to
// use in filterWorker tests where BlockByNumber is the first (and only) call
// before the worker returns on error.
type blockErrBackend struct {
	Backend // nil — satisfies interface; panics if any other method is called
	err     error
}

func (b *blockErrBackend) BlockByNumber(_ context.Context, _ rpc.BlockNumber) (*evmcore.EvmBlock, error) {
	return nil, b.err
}

// TestFilterWorker_PropagatesBlockError verifies that a BlockByNumber failure
// inside filterWorker is sent on errCh rather than silently discarded.
func TestFilterWorker_PropagatesBlockError(t *testing.T) {
	wantErr := errors.New("db unavailable")
	api := &PublicTxTraceAPI{b: &blockErrBackend{err: wantErr}}

	blocks := make(chan rpc.BlockNumber, 1)
	blocks <- rpc.BlockNumber(42)
	close(blocks)

	results := make(chan txtrace.ActionTrace, 1)
	errCh := make(chan error, 1)

	filterWorker(0, api, context.Background(), blocks, results, nil, nil, errCh)
	close(results)

	select {
	case got := <-errCh:
		if got != wantErr {
			t.Fatalf("expected %v, got %v", wantErr, got)
		}
	default:
		t.Fatal("filterWorker swallowed the BlockByNumber error: errCh is empty")
	}
}

// receiptMismatchBackend is a minimal Backend stub for TestTraceBlock_ReceiptCountMismatch.
// It returns a block with one transaction but zero receipts, to trigger the
// receipts[i] bounds check in traceBlock.
type receiptMismatchBackend struct {
	Backend // nil embedded — panics on any uncovered method
	block   *evmcore.EvmBlock
	stateDB *state.StateDB
}

func (b *receiptMismatchBackend) CurrentBlock() *evmcore.EvmBlock { return b.block }
func (b *receiptMismatchBackend) ChainConfig() *params.ChainConfig { return params.TestChainConfig }
func (b *receiptMismatchBackend) GetBlockContext(_ *evmcore.EvmHeader) vm.BlockContext {
	return vm.BlockContext{}
}
func (b *receiptMismatchBackend) TxTraceByHash(_ context.Context, _ common.Hash) (*[]txtrace.ActionTrace, error) {
	return nil, errors.New("not cached")
}
func (b *receiptMismatchBackend) StateAndHeaderByNumberOrHash(_ context.Context, _ rpc.BlockNumberOrHash) (*state.StateDB, *evmcore.EvmHeader, error) {
	return b.stateDB, nil, nil
}
func (b *receiptMismatchBackend) GetReceiptsByNumber(_ context.Context, _ rpc.BlockNumber) (types.Receipts, error) {
	return types.Receipts{}, nil // 0 receipts vs 1 tx in the block — mismatch
}

// TestTraceBlock_ReceiptCountMismatch verifies that traceBlock returns an error
// (not a panic) when GetReceiptsByNumber returns fewer receipts than the block
// has transactions. This guards against DB inconsistency causing an index panic.
func TestTraceBlock_ReceiptCountMismatch(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	stateDB, err := state.New(common.Hash{}, state.NewDatabase(db), nil)
	if err != nil {
		t.Fatalf("state.New: %v", err)
	}

	// Sign the tx so that tx.AsMessage succeeds and the code reaches receipts[i].
	key, _ := crypto.GenerateKey()
	signer := types.MakeSigner(params.TestChainConfig, big.NewInt(10))
	tx, _ := types.SignTx(
		types.NewTransaction(0, common.Address{1}, big.NewInt(0), 21000, big.NewInt(1), nil),
		signer, key,
	)
	block := &evmcore.EvmBlock{
		EvmHeader: evmcore.EvmHeader{
			Number: big.NewInt(10),
		},
		Transactions: types.Transactions{tx},
	}

	backend := &receiptMismatchBackend{block: block, stateDB: stateDB}
	api := &PublicTxTraceAPI{b: backend}

	_, traceErr := api.traceBlock(context.Background(), block, nil, nil)
	if traceErr == nil {
		t.Fatal("expected error for receipt count mismatch, got nil")
	}
}

// TestFilterArgs_AddressCap verifies that validateFilterArgs rejects fromAddress
// and toAddress slices that exceed maxFilterAddresses.
func TestFilterArgs_AddressCap(t *testing.T) {
	makeAddrs := func(n int) *[]common.Address {
		addrs := make([]common.Address, n)
		return &addrs
	}

	tests := []struct {
		name    string
		args    FilterArgs
		wantErr bool
	}{
		{
			name:    "nil addresses allowed",
			args:    FilterArgs{},
			wantErr: false,
		},
		{
			name:    "single fromAddress allowed",
			args:    FilterArgs{FromAddress: makeAddrs(1)},
			wantErr: false,
		},
		{
			name:    "exactly maxFilterAddresses fromAddresses allowed",
			args:    FilterArgs{FromAddress: makeAddrs(maxFilterTraceAddresses)},
			wantErr: false,
		},
		{
			name:    "one over maxFilterAddresses fromAddresses rejected",
			args:    FilterArgs{FromAddress: makeAddrs(maxFilterTraceAddresses + 1)},
			wantErr: true,
		},
		{
			name:    "exactly maxFilterAddresses toAddresses allowed",
			args:    FilterArgs{ToAddress: makeAddrs(maxFilterTraceAddresses)},
			wantErr: false,
		},
		{
			name:    "one over maxFilterAddresses toAddresses rejected",
			args:    FilterArgs{ToAddress: makeAddrs(maxFilterTraceAddresses + 1)},
			wantErr: true,
		},
		{
			name:    "both at max allowed",
			args:    FilterArgs{FromAddress: makeAddrs(maxFilterTraceAddresses), ToAddress: makeAddrs(maxFilterTraceAddresses)},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateFilterArgs(tc.args)
			if tc.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
