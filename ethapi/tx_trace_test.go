package ethapi

import (
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
