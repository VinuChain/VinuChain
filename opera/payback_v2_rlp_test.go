package opera

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// TestPaybackV2_RLPRoundtrip pins the bitfield encoding so a future
// refactor cannot silently change the on-the-wire layout of the
// PaybackV2 flag. The flag rides bit 1<<12 in the Upgrades bitmap.
func TestPaybackV2_RLPRoundtrip(t *testing.T) {
	u := Upgrades{PaybackV2: true}
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &u), "encode must succeed")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the same bytes")
	require.True(t, decoded.PaybackV2, "PaybackV2 must round-trip as true")
	require.False(t, decoded.PaybackV2Patch, "PaybackV2Patch must not spuriously appear after decode")
	require.False(t, decoded.Berlin, "no other flag must spuriously appear after decode")
	require.False(t, decoded.London, "no other flag must spuriously appear after decode")
	require.False(t, decoded.Podgorica, "no other flag must spuriously appear after decode")
	require.False(t, decoded.SfcV2, "no other flag must spuriously appear after decode")
	require.False(t, decoded.Elemont, "no other flag must spuriously appear after decode")
	require.False(t, decoded.SfcV2Patch5, "no other flag must spuriously appear after decode")
}

// TestPaybackV2_BitfieldDoesNotClashWithOtherFlags confirms the new bit
// (1<<12) is disjoint from every previously assigned bit. A copy-paste
// regression that re-used an earlier bit number would let either flag
// silently set the other on decode; this test catches that at build time.
func TestPaybackV2_BitfieldDoesNotClashWithOtherFlags(t *testing.T) {
	flags := map[string]uint64{
		"Berlin":                  berlinBit,
		"London":                  londonBit,
		"Llr":                     llrBit,
		"Podgorica":               podgoricaBit,
		"SfcV2":                   sfcV2Bit,
		"Elemont":                 elemontBit,
		"SfcV2Patch":              sfcV2PatchBit,
		"SfcV2Patch2":             sfcV2Patch2Bit,
		"SfcV2Patch3":             sfcV2Patch3Bit,
		"SfcV2Patch4":             sfcV2Patch4Bit,
		"ElemontPubkeyValidation": elemontPubkeyValidationBit,
		"SfcV2Patch5":             sfcV2Patch5Bit,
		"PaybackV2":               paybackV2Bit,
		"PaybackV2Patch":          paybackV2PatchBit,
	}
	seen := map[uint64]string{}
	for name, bit := range flags {
		if other, dup := seen[bit]; dup {
			t.Fatalf("bit %#x reused by %s and %s — RLP bitmap collision", bit, name, other)
		}
		seen[bit] = name
	}
	require.Equal(t, uint64(1<<12), uint64(paybackV2Bit), "paybackV2Bit must be 1<<12 (next free bit after sfcV2Patch5Bit)")
	require.Equal(t, uint64(1<<13), uint64(paybackV2PatchBit), "paybackV2PatchBit must be 1<<13 (next free bit after paybackV2Bit)")
}

// TestPaybackV2_MainnetAndLegacyConstructorsStayFalse defends against an
// accidental flip on networks that have NOT yet completed the PaybackV2
// rollout. Testnet is intentionally activated (see
// TestPaybackV2_TestnetActivatedWithNonSentinelAddress), so it's no
// longer in scope here.
func TestPaybackV2_MainnetAndLegacyConstructorsStayFalse(t *testing.T) {
	cases := []struct {
		name string
		fn   func() Rules
	}{
		{"MainNetRules", MainNetRules},
		{"TestNetRules", TestNetRules},
		{"VinuChainMainNetRules", VinuChainMainNetRules},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			require.False(t, c.fn().Upgrades.PaybackV2,
				"%s.Upgrades.PaybackV2 must be false until the PaybackV2 rollout completes on this network", c.name)
			require.False(t, c.fn().Upgrades.PaybackV2Patch,
				"%s.Upgrades.PaybackV2Patch must be false unless that network needs a corrected-contract rebind", c.name)
		})
	}
	// Fakenet may have PaybackV2 enabled for in-process activation tests.
	// No assertion on FakeNetRules / LegacyFakeNetRules.
}

// TestPaybackV2_TestnetActivatedWithCorrectedPatch pins the source state after
// the 2026-05-16 corrected PaybackV2 deployment. The original PaybackV2 flag
// remains true, and PaybackV2Patch provides the new one-shot edge that rebinds
// already-active testnet nodes to the corrected contract.
func TestPaybackV2_TestnetActivatedWithCorrectedPatch(t *testing.T) {
	rules := VinuChainTestNetRules()
	require.True(t, rules.Upgrades.PaybackV2,
		"VinuChainTestNetRules.Upgrades.PaybackV2 remains true")
	require.True(t, rules.Upgrades.PaybackV2Patch,
		"VinuChainTestNetRules.Upgrades.PaybackV2Patch must stage the corrected-contract rebind on already-active testnet")

	addr, err := PaybackV2ContractAddress(VinuChainTestNetworkID)
	require.NoError(t, err)
	require.False(t, PaybackV2AddressIsSentinel(addr),
		"paybackV2TestnetAddress must be the corrected non-sentinel QuotaContractV2 deployment")
}
