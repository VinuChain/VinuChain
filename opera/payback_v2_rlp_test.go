package opera

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
	require.False(t, decoded.SfcV2Patch6, "no other flag must spuriously appear after decode")
	require.False(t, decoded.Shanghai, "no other flag must spuriously appear after decode")
	require.False(t, decoded.Cancun, "no other flag must spuriously appear after decode")
	require.False(t, decoded.Prague, "no other flag must spuriously appear after decode")
	require.False(t, decoded.VinuBLS12381, "no other flag must spuriously appear after decode")
	require.False(t, decoded.VinuLatestEVM, "no other flag must spuriously appear after decode")
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
		"SfcV2Patch6":             sfcV2Patch6Bit,
		"Shanghai":                shanghaiBit,
		"Cancun":                  cancunBit,
		"Prague":                  pragueBit,
		"VinuBLS12381":            vinuBLS12381Bit,
		"VinuLatestEVM":           vinuLatestEVMBit,
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
	require.Equal(t, uint64(1<<14), uint64(sfcV2Patch6Bit), "sfcV2Patch6Bit must be 1<<14 (next free bit after paybackV2PatchBit)")
	require.Equal(t, uint64(1<<15), uint64(shanghaiBit), "shanghaiBit must be 1<<15 (next free bit after sfcV2Patch6Bit)")
	require.Equal(t, uint64(1<<16), uint64(cancunBit), "cancunBit must be 1<<16 (next free bit after shanghaiBit)")
	require.Equal(t, uint64(1<<17), uint64(pragueBit), "pragueBit must be 1<<17 (next free bit after cancunBit)")
	require.Equal(t, uint64(1<<18), uint64(vinuBLS12381Bit), "vinuBLS12381Bit must be 1<<18 (next free bit after pragueBit)")
	require.Equal(t, uint64(1<<19), uint64(vinuLatestEVMBit), "vinuLatestEVMBit must be 1<<19 (next free bit after vinuBLS12381Bit)")
}

func TestEthereumForkBitsKnownRLP(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &Upgrades{Shanghai: true, Cancun: true}),
		"encode must succeed")
	require.Equal(t, "c483018000", hex.EncodeToString(buf.Bytes()),
		"Shanghai/Cancun bitmap wire shape must stay stable")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the fixture bytes")
	require.True(t, decoded.Shanghai, "Shanghai must decode from bit 1<<15")
	require.True(t, decoded.Cancun, "Cancun must decode from bit 1<<16")
	require.False(t, decoded.Prague, "Prague must not decode from Shanghai/Cancun bits")
	require.False(t, decoded.VinuBLS12381, "VinuBLS12381 must not decode from Shanghai/Cancun bits")
	require.False(t, decoded.VinuLatestEVM, "VinuLatestEVM must not decode from Shanghai/Cancun bits")
	require.False(t, decoded.PaybackV2, "PaybackV2 must not decode from Shanghai/Cancun bits")
}

func TestPragueBitKnownRLP(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &Upgrades{Prague: true}),
		"encode must succeed")
	require.Equal(t, "c483020000", hex.EncodeToString(buf.Bytes()),
		"Prague bitmap wire shape must stay stable")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the fixture bytes")
	require.True(t, decoded.Prague, "Prague must decode from bit 1<<17")
	require.False(t, decoded.Cancun, "Cancun must not decode from the Prague bit")
	require.False(t, decoded.VinuBLS12381, "VinuBLS12381 must not decode from the Prague bit")
	require.False(t, decoded.VinuLatestEVM, "VinuLatestEVM must not decode from the Prague bit")
}

func TestVinuBLSBitKnownRLP(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &Upgrades{VinuBLS12381: true}),
		"encode must succeed")
	require.Equal(t, "c483040000", hex.EncodeToString(buf.Bytes()),
		"VinuBLS12381 bitmap wire shape must stay stable")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the fixture bytes")
	require.True(t, decoded.VinuBLS12381, "VinuBLS12381 must decode from bit 1<<18")
	require.False(t, decoded.Prague, "Prague must not decode from the VinuBLS12381 bit")
	require.False(t, decoded.VinuLatestEVM, "VinuLatestEVM must not decode from the VinuBLS12381 bit")
}

func TestVinuLatestEVMBitKnownRLP(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, rlp.Encode(&buf, &Upgrades{VinuLatestEVM: true}),
		"encode must succeed")
	require.Equal(t, "c483080000", hex.EncodeToString(buf.Bytes()),
		"VinuLatestEVM bitmap wire shape must stay stable")

	var decoded Upgrades
	require.NoError(t, rlp.DecodeBytes(buf.Bytes(), &decoded),
		"decode must succeed against the fixture bytes")
	require.True(t, decoded.VinuLatestEVM, "VinuLatestEVM must decode from bit 1<<19")
	require.False(t, decoded.VinuBLS12381, "VinuBLS12381 must not decode from the VinuLatestEVM bit")
}

// TestPaybackV2_LegacyConstructorsStayFalse defends against an accidental flip
// on networks that have NOT completed the PaybackV2 rollout. Both VinuChain
// testnet (2026-05-16) and VinuChain mainnet (2026-08-21) are intentionally
// activated and are covered by TestPaybackV2_TestnetActivatedWithNonSentinelAddress
// and TestPaybackV2_MainnetActivatedWithNonSentinelAddress respectively. What is
// left here are the legacy Fantom-inherited constructors, which have no deployed
// QuotaContractV2 and must never gain one by accident.
func TestPaybackV2_LegacyConstructorsStayFalse(t *testing.T) {
	cases := []struct {
		name string
		fn   func() Rules
	}{
		{"MainNetRules", MainNetRules},
		{"TestNetRules", TestNetRules},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			require.False(t, c.fn().Upgrades.PaybackV2,
				"%s.Upgrades.PaybackV2 must be false — this legacy constructor has no deployed QuotaContractV2", c.name)
			require.False(t, c.fn().Upgrades.PaybackV2Patch,
				"%s.Upgrades.PaybackV2Patch must be false unless that network needs a corrected-contract rebind", c.name)
		})
	}
	// Fakenet may have PaybackV2 enabled for in-process activation tests.
	// No assertion on FakeNetRules / LegacyFakeNetRules.
}

// TestPaybackV2_MainnetActivatedWithNonSentinelAddress pins the 2026-08-21
// mainnet rollout: QuotaContractV2 was deployed at
// 0x5D989A2d65d049e2198D91d8ddc31C918f2544AB (deployer nonce 0, owner
// 0xf9c82B1117e8BeA97843042521B8FBC93044f347) and both the mainnet and the
// staging address slots were baked, so VinuChainMainNetRules() may enable
// PaybackV2. The pair is asserted together deliberately: the flag without a
// real address is precisely the shape EnforcePaybackV2StartupCheck() panics on,
// and staging inherits the flag from mainnet.
func TestPaybackV2_MainnetActivatedWithNonSentinelAddress(t *testing.T) {
	require.True(t, VinuChainMainNetRules().Upgrades.PaybackV2,
		"mainnet rules enable PaybackV2 for the 2026-08-29 full-parity release")
	require.False(t, VinuChainMainNetRules().Upgrades.PaybackV2Patch,
		"mainnet crosses the PaybackV2 edge once with the correct address, so it never needs the rebind patch")

	mainnetAddr, err := PaybackV2ContractAddress(VinuChainMainNetworkID)
	require.NoError(t, err)
	require.Equal(t,
		common.HexToAddress("0x5D989A2d65d049e2198D91d8ddc31C918f2544AB"), mainnetAddr,
		"mainnet V2 address must be the deployed QuotaContractV2, not a predicted or edited value")

	stagingAddr, err := PaybackV2ContractAddress(VinuChainStagingNetworkID)
	require.NoError(t, err)
	require.False(t, PaybackV2AddressIsSentinel(stagingAddr),
		"staging inherits PaybackV2 from mainnet rules, so its slot must also be non-sentinel")
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
