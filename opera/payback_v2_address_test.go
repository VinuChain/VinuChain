package opera

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestPaybackV2ContractAddress_UnknownNetwork rejects unknown networks so
// a misconfigured fork (custom networkID) cannot silently fall through
// the activation gate.
func TestPaybackV2ContractAddress_UnknownNetwork(t *testing.T) {
	addr, err := PaybackV2ContractAddress(0xdeadbeef)
	require.Error(t, err, "unknown networkID must return an error")
	require.Equal(t, common.Address{}, addr,
		"error path must return zero address so caller cannot accidentally use it")
}

// TestPaybackV2ContractAddress_KnownNetworks returns the sentinel zero
// until the real addresses are baked in. Re-keying when an address is
// recorded would catch an accidental commit that wipes the value.
func TestPaybackV2ContractAddress_KnownNetworks(t *testing.T) {
	cases := []struct {
		name      string
		networkID uint64
	}{
		{"testnet", VinuChainTestNetworkID},
		{"mainnet", VinuChainMainNetworkID},
		{"staging", VinuChainStagingNetworkID},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			addr, err := PaybackV2ContractAddress(c.networkID)
			require.NoError(t, err, "known networkID must not error")
			// Until the contract is deployed, the address is the zero sentinel.
			// When the V2 address is recorded in payback_v2_address.go, this
			// assertion will start failing — at which point this test should
			// be updated to assert the EXACT address that was baked in.
			require.Equal(t, common.Address{}, addr,
				"address slot for %s must remain the zero sentinel until the V2 contract is deployed", c.name)
		})
	}
}

// TestPaybackV2AddressIsSentinel returns true for the zero address and
// false for any non-zero address.
func TestPaybackV2AddressIsSentinel(t *testing.T) {
	require.True(t, PaybackV2AddressIsSentinel(common.Address{}),
		"zero address must be detected as the sentinel")
	require.False(t, PaybackV2AddressIsSentinel(common.HexToAddress("0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D")),
		"the V1 proxy address must not be flagged as a sentinel")
	require.False(t, PaybackV2AddressIsSentinel(common.HexToAddress("0x0000000000000000000000000000000000000001")),
		"any non-zero address must not be flagged as a sentinel")
}

// TestSetPaybackV2ContractAddressForTesting round-trips testnet, mainnet,
// and staging values through the helper, and confirms the restore closure
// reverts to the previous value.
func TestSetPaybackV2ContractAddressForTesting(t *testing.T) {
	target := common.HexToAddress("0x1111111111111111111111111111111111111111")
	for _, networkID := range []uint64{VinuChainTestNetworkID, VinuChainMainNetworkID, VinuChainStagingNetworkID} {
		networkID := networkID
		t.Run("network_"+commonHex(networkID), func(t *testing.T) {
			before, err := PaybackV2ContractAddress(networkID)
			require.NoError(t, err)
			restore := SetPaybackV2ContractAddressForTesting(networkID, target)
			t.Cleanup(restore)

			got, err := PaybackV2ContractAddress(networkID)
			require.NoError(t, err)
			require.Equal(t, target, got, "override must take effect")

			restore()
			restored, err := PaybackV2ContractAddress(networkID)
			require.NoError(t, err)
			require.Equal(t, before, restored, "restore closure must revert to original")
		})
	}
}

// commonHex returns the 4-char lowercase hex of the network ID, used only
// for subtest naming so failures point at the network in question.
func commonHex(v uint64) string {
	const hexDigits = "0123456789abcdef"
	if v == 0 {
		return "0"
	}
	out := make([]byte, 0, 16)
	for v > 0 {
		out = append([]byte{hexDigits[v&0xf]}, out...)
		v >>= 4
	}
	return string(out)
}

// TestEnforcePaybackV2StartupCheck_PassesWithFlagOff is the default
// scaffold state: every network constructor leaves PaybackV2=false, so the
// startup check must pass even with sentinel addresses everywhere. This is
// the state v2.0.16-elemont ships in.
func TestEnforcePaybackV2StartupCheck_PassesWithFlagOff(t *testing.T) {
	// Sanity: confirm the constructors do indeed have PaybackV2=false at scaffold time.
	require.False(t, VinuChainTestNetRules().Upgrades.PaybackV2,
		"scaffold-state testnet rules must NOT have PaybackV2 enabled until the contract is deployed and recorded")
	require.False(t, VinuChainMainNetRules().Upgrades.PaybackV2,
		"mainnet must NOT have PaybackV2 enabled until its mainnet rollout prerequisites complete")
	require.NotPanics(t, func() { EnforcePaybackV2StartupCheck() },
		"startup check must pass when every constructor leaves PaybackV2 off, regardless of sentinel address state")
}
