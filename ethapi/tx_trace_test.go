package ethapi

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

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
