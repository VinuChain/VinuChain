package opera

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPaybackV2ActivationActuallyMovesQuotaCacheAddress pins the pair that the
// activation depends on, which no other test covers.
//
// gossip/block_processor.go's rebindPaybackV2 reads
// PaybackV2ContractAddress(NetworkID) and overwrites
// Economy.QuotaCacheAddress with it. The existing activation tests are
// source-structural — they assert that the assignment statement is present in
// the file — so they still pass if the two addresses are equal or if the wrong
// one is baked. Both of those are silent failures: the flag flips, the log line
// prints, and the fee-refund cache keeps resolving the abandoned V1 proxy while
// nothing errors.
//
// For every network whose rule constructor enables PaybackV2, assert the baked
// address is real AND different from the constructor's QuotaCacheAddress, so
// activation is a genuine move rather than a no-op.
func TestPaybackV2ActivationActuallyMovesQuotaCacheAddress(t *testing.T) {
	cases := []struct {
		name      string
		networkID uint64
		rules     Rules
	}{
		{"mainnet", VinuChainMainNetworkID, VinuChainMainNetRules()},
		{"testnet", VinuChainTestNetworkID, VinuChainTestNetRules()},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if !c.rules.Upgrades.PaybackV2 {
				t.Skip("PaybackV2 not enabled on this network")
			}
			baked, err := PaybackV2ContractAddress(c.networkID)
			require.NoError(t, err, "a network with PaybackV2 enabled must have an address slot")
			require.False(t, PaybackV2AddressIsSentinel(baked),
				"PaybackV2 is enabled but the address slot is the zero sentinel — rebindPaybackV2 log.Crits at the activation seal")
			require.NotEqual(t, c.rules.Economy.QuotaCacheAddress, baked,
				"the baked V2 address equals the rules' existing QuotaCacheAddress, so activation would rebind to the address already in use — a silent no-op that leaves fee refunds on the old contract")
		})
	}
}

// TestPaybackV2StagingInheritsAWorkingRebind covers the network that has no rule
// constructor of its own. Staging is synthesised from mainnet rules with the
// NetworkID rewritten, so it inherits Upgrades.PaybackV2 but resolves its
// address from the staging slot — the one combination that boots fine and then
// fails at the seal.
func TestPaybackV2StagingInheritsAWorkingRebind(t *testing.T) {
	staging := VinuChainMainNetRules()
	staging.NetworkID = VinuChainStagingNetworkID
	if !staging.Upgrades.PaybackV2 {
		t.Skip("PaybackV2 not enabled on mainnet, so staging does not inherit it")
	}
	baked, err := PaybackV2ContractAddress(staging.NetworkID)
	require.NoError(t, err)
	require.False(t, PaybackV2AddressIsSentinel(baked),
		"staging inherits PaybackV2=true from mainnet, so its own slot must resolve to a real contract")
	require.NotEqual(t, staging.Economy.QuotaCacheAddress, baked,
		"staging activation must move QuotaCacheAddress, not rebind it to itself")
}
