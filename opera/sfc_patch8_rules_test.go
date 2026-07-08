package opera

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// TestSfcV2Patch8_RLPRoundtrip pins the bitfield encoding so a future refactor
// cannot silently change the on-the-wire layout of the SfcV2Patch8 flag. The
// flag rides bit 1<<21 in the Upgrades bitmap.
func TestSfcV2Patch8_RLPRoundtrip(t *testing.T) {
	u := Upgrades{SfcV2Patch8: true}
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &u), "encode must succeed")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the same bytes")
	require.True(t, decoded.SfcV2Patch8, "SfcV2Patch8 must round-trip as true")
	require.False(t, decoded.SfcV2Patch7, "SfcV2Patch7 must not spuriously appear after decode")
	require.False(t, decoded.SfcV2Patch6, "SfcV2Patch6 must not spuriously appear after decode")
	require.False(t, decoded.SfcV2Patch5, "SfcV2Patch5 must not spuriously appear after decode")
	require.False(t, decoded.PaybackV2, "PaybackV2 must not spuriously appear after decode")
	require.False(t, decoded.PaybackV2Patch, "PaybackV2Patch must not spuriously appear after decode")
	require.False(t, decoded.VinuLatestEVM, "VinuLatestEVM must not spuriously appear after decode")
	require.False(t, decoded.Berlin, "no other flag must spuriously appear after decode")
	require.False(t, decoded.SfcV2, "no other flag must spuriously appear after decode")
}

// TestSfcV2Patch8_BitfieldDoesNotClashWithOtherFlags confirms the new bit
// (1<<21) is disjoint from every previously assigned bit.
func TestSfcV2Patch8_BitfieldDoesNotClashWithOtherFlags(t *testing.T) {
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
	}
	seen := map[uint64]string{}
	for name, bit := range flags {
		if other, dup := seen[bit]; dup {
			t.Fatalf("bit %#x reused by %s and %s — RLP bitmap collision", bit, name, other)
		}
		seen[bit] = name
	}
	require.Equal(t, uint64(1<<21), uint64(sfcV2Patch8Bit), "sfcV2Patch8Bit must be 1<<21 (next free bit after sfcV2Patch7Bit)")
}

// TestSfcV2Patch8_DefaultsFalseOnRealConstructors defends against an accidental
// flip on any network other than the testnet that needs the reactivation
// reflash. Mainnet ships the permanent fix via GetLatestContractBin() at its
// first SfcV2 activation, so it must leave SfcV2Patch8 unset.
func TestSfcV2Patch8_DefaultsFalseOnRealConstructors(t *testing.T) {
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
			require.False(t, c.fn().Upgrades.SfcV2Patch8,
				"%s.Upgrades.SfcV2Patch8 must be false — only VinuChainTestNetRules stages the reactivation reflash", c.name)
		})
	}
}

// TestSfcV2Patch8_TestnetActivated pins that testnet stages the SfcV2Patch8
// reflash (Cycle-163 self-service-reactivation bytecode).
func TestSfcV2Patch8_TestnetActivated(t *testing.T) {
	require.True(t, VinuChainTestNetRules().Upgrades.SfcV2Patch8,
		"VinuChainTestNetRules.Upgrades.SfcV2Patch8 must be true to reflash Cycle-163 self-service-reactivation bytecode")
}
