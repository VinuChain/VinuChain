package gossip

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera"
)

// TestCheckUnprotectedTxsPolicy pins the mainnet replay-protection invariant:
// AllowUnprotectedTxs is refused only on mainnet (NetworkID 207) and permitted on
// every other network. NewService calls this helper, so a regression here is the
// canary for the guard being weakened or dropped.
func TestCheckUnprotectedTxsPolicy(t *testing.T) {
	cases := []struct {
		name      string
		allow     bool
		networkID uint64
		wantErr   bool
	}{
		{"mainnet + flag on is refused", true, opera.VinuChainMainNetworkID, true},
		{"mainnet + flag off is allowed", false, opera.VinuChainMainNetworkID, false},
		{"testnet + flag on is allowed", true, opera.VinuChainTestNetworkID, false},
		{"staging + flag on is allowed", true, opera.VinuChainStagingNetworkID, false},
		{"fakenet + flag on is allowed", true, opera.VinuChainNewNetworkID, false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := checkUnprotectedTxsPolicy(tc.allow, tc.networkID)
			if tc.wantErr && err == nil {
				t.Fatalf("expected refusal for allow=%v networkID=%d", tc.allow, tc.networkID)
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected refusal for allow=%v networkID=%d: %v", tc.allow, tc.networkID, err)
			}
		})
	}
}
