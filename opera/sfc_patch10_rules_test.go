package opera

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// TestSfcV2Patch10_RLPRoundtrip pins the bitfield encoding so a future refactor
// cannot silently change the on-the-wire layout of the SfcV2Patch10 flag. The
// flag rides bit 1<<23 in the Upgrades bitmap.
func TestSfcV2Patch10_RLPRoundtrip(t *testing.T) {
	u := Upgrades{SfcV2Patch10: true}
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &u), "encode must succeed")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the same bytes")
	require.True(t, decoded.SfcV2Patch10, "SfcV2Patch10 must round-trip as true")
	require.False(t, decoded.SfcV2Patch9, "SfcV2Patch9 must not spuriously appear after decode")
	require.False(t, decoded.SfcV2Patch8, "SfcV2Patch8 must not spuriously appear after decode")
	require.False(t, decoded.SfcV2Patch7, "SfcV2Patch7 must not spuriously appear after decode")
	require.False(t, decoded.PaybackV2, "PaybackV2 must not spuriously appear after decode")
	require.False(t, decoded.PaybackV2Patch, "PaybackV2Patch must not spuriously appear after decode")
	require.False(t, decoded.VinuLatestEVM, "VinuLatestEVM must not spuriously appear after decode")
	require.False(t, decoded.Berlin, "no other flag must spuriously appear after decode")
	require.False(t, decoded.SfcV2, "no other flag must spuriously appear after decode")
}

// TestSfcV2Patch10_BitfieldDoesNotClashWithOtherFlags confirms the new bit
// (1<<23) is disjoint from every previously assigned bit.
func TestSfcV2Patch10_BitfieldDoesNotClashWithOtherFlags(t *testing.T) {
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
		"SfcV2Patch6":             sfcV2Patch6Bit,
		"Shanghai":                shanghaiBit,
		"Cancun":                  cancunBit,
		"Prague":                  pragueBit,
		"VinuBLS12381":            vinuBLS12381Bit,
		"VinuLatestEVM":           vinuLatestEVMBit,
		"SfcV2Patch7":             sfcV2Patch7Bit,
		"SfcV2Patch8":             sfcV2Patch8Bit,
		"SfcV2Patch9":             sfcV2Patch9Bit,
		"SfcV2Patch10":            sfcV2Patch10Bit,
	}
	seen := map[uint64]string{}
	for name, bit := range flags {
		if other, dup := seen[bit]; dup {
			t.Fatalf("bit %#x reused by %s and %s — RLP bitmap collision", bit, name, other)
		}
		seen[bit] = name
	}
	require.Equal(t, uint64(1<<23), uint64(sfcV2Patch10Bit), "sfcV2Patch10Bit must be 1<<23 (next free bit after sfcV2Patch9Bit)")
}

// TestSfcV2Patch10_DefaultsFalseOnRealConstructors defends against an
// accidental flip on any network other than the testnet that needs the
// lockup-preservation reflash. Mainnet ships the permanent fix via
// GetLatestContractBin() at its first SfcV2 activation, so it must leave
// SfcV2Patch10 unset.
func TestSfcV2Patch10_DefaultsFalseOnRealConstructors(t *testing.T) {
	cases := []struct {
		name string
		fn   func() Rules
	}{
		{"MainNetRules", MainNetRules},
		{"TestNetRules", TestNetRules},
		{"FakeNetRules", FakeNetRules},
		{"LegacyFakeNetRules", LegacyFakeNetRules},
		{"VinuChainMainNetRules", VinuChainMainNetRules},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			require.False(t, c.fn().Upgrades.SfcV2Patch10,
				"%s.Upgrades.SfcV2Patch10 must be false — only VinuChainTestNetRules stages the lockup-preservation reflash", c.name)
		})
	}
}

// TestSfcV2Patch10_TestnetActivated pins that testnet stages the SfcV2Patch10
// reflash (Cycle-165 lockup-preservation bytecode).
func TestSfcV2Patch10_TestnetActivated(t *testing.T) {
	require.True(t, VinuChainTestNetRules().Upgrades.SfcV2Patch10,
		"VinuChainTestNetRules.Upgrades.SfcV2Patch10 must be true to reflash Cycle-165 lockup-preservation bytecode")
}
