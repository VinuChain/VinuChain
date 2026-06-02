package opera

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

func TestUpdateRules(t *testing.T) {
	require := require.New(t)

	exp := FakeNetRules()

	exp.Dag.MaxParents = 5
	exp.Economy.MinGasPrice = big.NewInt(7)
	exp.Blocks.MaxBlockGas = 1000
	got, err := UpdateRules(exp, []byte(`{"Dag":{"MaxParents":5},"Economy":{"MinGasPrice":7},"Blocks":{"MaxBlockGas":1000}}`))
	require.NoError(err)
	require.Equal(exp.String(), got.String(), "mutate fields")

	got, err = UpdateRules(exp, []byte(`{"Name":"xxx","NetworkID":1}`))
	require.NoError(err)
	require.Equal(exp.String(), got.String(), "readonly fields")

	got, err = UpdateRules(exp, []byte(`{}`))
	require.NoError(err)
	require.Equal(exp.String(), got.String(), "empty diff")

	_, err = UpdateRules(exp, []byte(`}{`))
	require.Error(err)
}

func TestUpdateRulesGovernanceBounds(t *testing.T) {
	require := require.New(t)
	base := FakeNetRules()

	// MaxParents below minimum
	_, err := UpdateRules(base, []byte(`{"Dag":{"MaxParents":0}}`))
	require.Error(err, "MaxParents=0 should be rejected")
	require.Contains(err.Error(), "MaxParents")

	_, err = UpdateRules(base, []byte(`{"Dag":{"MaxParents":2}}`))
	require.Error(err, "MaxParents=2 should be rejected")
	require.Contains(err.Error(), "MaxParents")

	// MaxEventGas=0 halts chain
	_, err = UpdateRules(base, []byte(`{"Economy":{"Gas":{"MaxEventGas":0}}}`))
	require.Error(err, "MaxEventGas=0 should be rejected")
	require.Contains(err.Error(), "MaxEventGas")

	// EventGas > MaxEventGas
	_, err = UpdateRules(base, []byte(`{"Economy":{"Gas":{"EventGas":999999999}}}`))
	require.Error(err, "EventGas > MaxEventGas should be rejected")
	require.Contains(err.Error(), "EventGas exceeds")

	// MaxEpochGas=0
	_, err = UpdateRules(base, []byte(`{"Epochs":{"MaxEpochGas":0}}`))
	require.Error(err, "MaxEpochGas=0 should be rejected")
	require.Contains(err.Error(), "MaxEpochGas")

	// ShortGasPower.AllocPerSec=0
	_, err = UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"AllocPerSec":0}}}`))
	require.Error(err, "ShortGasPower.AllocPerSec=0 should be rejected")
	require.Contains(err.Error(), "ShortGasPower.AllocPerSec")

	// LongGasPower.AllocPerSec=0
	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"AllocPerSec":0}}}`))
	require.Error(err, "LongGasPower.AllocPerSec=0 should be rejected")
	require.Contains(err.Error(), "LongGasPower.AllocPerSec")

	// MaxBlockGas=0
	_, err = UpdateRules(base, []byte(`{"Blocks":{"MaxBlockGas":0}}`))
	require.Error(err, "MaxBlockGas=0 should be rejected")
	require.Contains(err.Error(), "MaxBlockGas")

	// MisbehaviourProofGas > MaxEventGas/2
	bigGas := base.Economy.Gas.MaxEventGas/2 + 1
	diff := []byte(`{"Economy":{"Gas":{"MisbehaviourProofGas":` + big.NewInt(int64(bigGas)).String() + `}}}`)
	_, err = UpdateRules(base, diff)
	require.Error(err, "MisbehaviourProofGas > MaxEventGas/2 should be rejected")
	require.Contains(err.Error(), "MisbehaviourProofGas")

	// ShortGasPower.AllocPerSec exceeds upper bound
	_, err = UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"AllocPerSec":2000000000000}}}`))
	require.Error(err, "ShortGasPower.AllocPerSec > 1e12 should be rejected")
	require.Contains(err.Error(), "ShortGasPower.AllocPerSec")

	// LongGasPower.AllocPerSec exceeds upper bound
	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"AllocPerSec":2000000000000}}}`))
	require.Error(err, "LongGasPower.AllocPerSec > 1e12 should be rejected")
	require.Contains(err.Error(), "LongGasPower.AllocPerSec")

	// MaxAllocPeriod exceeds upper bound (> 1 week in nanoseconds)
	_, err = UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"MaxAllocPeriod":700000000000000}}}`))
	require.Error(err, "ShortGasPower.MaxAllocPeriod > 1 week should be rejected")
	require.Contains(err.Error(), "ShortGasPower.MaxAllocPeriod")

	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"MaxAllocPeriod":700000000000000}}}`))
	require.Error(err, "LongGasPower.MaxAllocPeriod > 1 week should be rejected")
	require.Contains(err.Error(), "LongGasPower.MaxAllocPeriod")

	// Valid changes should still work
	_, err = UpdateRules(base, []byte(`{"Dag":{"MaxParents":10}}`))
	require.NoError(err, "valid MaxParents=10 should succeed")

	// ExtraDataGas too large: 128 * 100000 = 12800000 exceeds MaxEventGas (~10028000)
	_, err = UpdateRules(base, []byte(`{"Economy":{"Gas":{"ExtraDataGas":100000}}}`))
	require.Error(err, "ExtraDataGas that makes maxEmptyEventGas exceed MaxEventGas should be rejected")

	// ParentGas too large: 7 * 1600000 = 11200000 exceeds MaxEventGas
	_, err = UpdateRules(base, []byte(`{"Economy":{"Gas":{"ParentGas":1600000}}}`))
	require.Error(err, "ParentGas that makes maxEmptyEventGas exceed MaxEventGas should be rejected")
}

// TestValidateRulesBounds_MinGasPriceZero verifies that MinGasPrice=0 is rejected by
// validateRulesBounds. A zero base fee removes all spam protection — any number of
// zero-cost transactions can flood the mempool and halt the chain.
func TestValidateRulesBounds_MinGasPriceZero(t *testing.T) {
	require := require.New(t)
	base := FakeNetRules()

	_, err := UpdateRules(base, []byte(`{"Economy":{"MinGasPrice":0}}`))
	require.Error(err, "MinGasPrice=0 should be rejected")
	require.Contains(err.Error(), "MinGasPrice")

	// Positive value must be accepted
	_, err = UpdateRules(base, []byte(`{"Economy":{"MinGasPrice":1}}`))
	require.NoError(err, "MinGasPrice=1 must be accepted")
}

// TestValidateRulesBounds_MaxAllocPeriodMinimum verifies that MaxAllocPeriod values below
// 1 second are rejected for both Short and Long gas power rules. Values below 1 second
// cause maxTotalGasPower() to return 0 after integer division (AllocPerSec×MaxAllocPeriod/1e9),
// which triggers a div-by-zero panic in constructiveGasPrice on all accepting nodes.
func TestValidateRulesBounds_MaxAllocPeriodMinimum(t *testing.T) {
	require := require.New(t)
	base := FakeNetRules()

	// 1 ns — passes non-zero check but product floors to 0 after /1e9
	_, err := UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"MaxAllocPeriod":1}}}`))
	require.Error(err, "ShortGasPower.MaxAllocPeriod=1ns should be rejected (< 1 second)")
	require.Contains(err.Error(), "ShortGasPower.MaxAllocPeriod")

	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"MaxAllocPeriod":1}}}`))
	require.Error(err, "LongGasPower.MaxAllocPeriod=1ns should be rejected (< 1 second)")
	require.Contains(err.Error(), "LongGasPower.MaxAllocPeriod")

	// 999,999,999 ns — one nanosecond below 1 second, still floors to 0
	_, err = UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"MaxAllocPeriod":999999999}}}`))
	require.Error(err, "ShortGasPower.MaxAllocPeriod=999999999ns should be rejected (< 1 second)")

	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"MaxAllocPeriod":999999999}}}`))
	require.Error(err, "LongGasPower.MaxAllocPeriod=999999999ns should be rejected (< 1 second)")

	// Exactly 1 second (1,000,000,000 ns) — must be accepted
	_, err = UpdateRules(base, []byte(`{"Economy":{"ShortGasPower":{"MaxAllocPeriod":1000000000}}}`))
	require.NoError(err, "ShortGasPower.MaxAllocPeriod=1s must be accepted")

	_, err = UpdateRules(base, []byte(`{"Economy":{"LongGasPower":{"MaxAllocPeriod":1000000000}}}`))
	require.NoError(err, "LongGasPower.MaxAllocPeriod=1s must be accepted")
}

func TestMainNetRulesRLP(t *testing.T) {
	rules := MainNetRules()
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
}

func TestRulesBerlinRLP(t *testing.T) {
	rules := MainNetRules()
	rules.Upgrades.Berlin = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.Berlin)
}

func TestRulesLondonRLP(t *testing.T) {
	rules := MainNetRules()
	rules.Upgrades.London = true
	rules.Upgrades.Berlin = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.Berlin)
	require.True(decodedRules.Upgrades.London)
}

func TestRulesShanghaiRLP(t *testing.T) {
	rules := MainNetRules()
	rules.Upgrades.Berlin = true
	rules.Upgrades.London = true
	rules.Upgrades.Shanghai = true
	rules.Upgrades.Cancun = true
	rules.Upgrades.Prague = true
	rules.Upgrades.VinuBLS12381 = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.Berlin)
	require.True(decodedRules.Upgrades.London)
	require.True(decodedRules.Upgrades.Shanghai)
	require.True(decodedRules.Upgrades.Cancun)
	require.True(decodedRules.Upgrades.Prague)
	require.True(decodedRules.Upgrades.VinuBLS12381)
}

func TestRulesEthereumForkDefaultsRLP(t *testing.T) {
	cases := []struct {
		mk             func() Rules
		expectShanghai bool
		expectCancun   bool
		expectPrague   bool
		expectVinuBLS  bool
		network        string
	}{
		{MainNetRules, false, false, false, false, "MainNetRules"},
		{TestNetRules, false, false, false, false, "TestNetRules"},
		{VinuChainMainNetRules, true, true, true, false, "VinuChainMainNetRules"},
		{VinuChainTestNetRules, true, true, true, true, "VinuChainTestNetRules"},
		{FakeNetRules, true, true, true, true, "FakeNetRules"},
		{LegacyFakeNetRules, true, true, true, true, "LegacyFakeNetRules"},
	}
	for _, c := range cases {
		rules := c.mk()
		require := require.New(t)
		require.Equal(c.expectShanghai, rules.Upgrades.Shanghai,
			"%s Shanghai default mismatch", c.network)
		require.Equal(c.expectCancun, rules.Upgrades.Cancun,
			"%s Cancun default mismatch", c.network)
		require.Equal(c.expectPrague, rules.Upgrades.Prague,
			"%s Prague default mismatch", c.network)
		require.Equal(c.expectVinuBLS, rules.Upgrades.VinuBLS12381,
			"%s VinuBLS12381 default mismatch", c.network)

		b, err := rlp.EncodeToBytes(rules)
		require.NoError(err)

		decodedRules := Rules{}
		require.NoError(rlp.DecodeBytes(b, &decodedRules))

		require.Equal(rules.String(), decodedRules.String())
		require.Equal(c.expectShanghai, decodedRules.Upgrades.Shanghai,
			"%s Upgrades.Shanghai must round-trip through RLP", c.network)
		require.Equal(c.expectCancun, decodedRules.Upgrades.Cancun,
			"%s Upgrades.Cancun must round-trip through RLP", c.network)
		require.Equal(c.expectPrague, decodedRules.Upgrades.Prague,
			"%s Upgrades.Prague must round-trip through RLP", c.network)
		require.Equal(c.expectVinuBLS, decodedRules.Upgrades.VinuBLS12381,
			"%s Upgrades.VinuBLS12381 must round-trip through RLP", c.network)
	}
}

func TestEvmChainConfigShanghaiActivationHeight(t *testing.T) {
	before := Upgrades{Berlin: true, London: true}
	after := before
	after.Shanghai = true
	after.Cancun = true
	after.Prague = true
	rules := VinuChainTestNetRules()

	cfg := rules.EvmChainConfig([]UpgradeHeight{
		{Upgrades: before, Height: 0},
		{Upgrades: after, Height: 123},
	})

	require.Equal(t, big.NewInt(0), cfg.BerlinBlock)
	require.Equal(t, big.NewInt(0), cfg.LondonBlock)
	require.Equal(t, big.NewInt(123), cfg.ShanghaiBlock)
	require.Equal(t, big.NewInt(123), cfg.CancunBlock)
	require.Equal(t, big.NewInt(123), cfg.PragueBlock)
}

func TestEvmChainConfigCancunCanActivateAfterShanghai(t *testing.T) {
	before := Upgrades{Berlin: true, London: true, Shanghai: true}
	after := before
	after.Cancun = true
	rules := VinuChainTestNetRules()

	cfg := rules.EvmChainConfig([]UpgradeHeight{
		{Upgrades: before, Height: 0},
		{Upgrades: after, Height: 456},
	})

	require.Equal(t, big.NewInt(0), cfg.ShanghaiBlock)
	require.Equal(t, big.NewInt(456), cfg.CancunBlock)
	require.Nil(t, cfg.PragueBlock)
	require.Nil(t, cfg.VinuBLSBlock)
}

func TestEvmChainConfigPragueCanActivateAfterCancun(t *testing.T) {
	before := Upgrades{Berlin: true, London: true, Shanghai: true, Cancun: true}
	after := before
	after.Prague = true
	rules := VinuChainTestNetRules()

	cfg := rules.EvmChainConfig([]UpgradeHeight{
		{Upgrades: before, Height: 0},
		{Upgrades: after, Height: 789},
	})

	require.Equal(t, big.NewInt(0), cfg.ShanghaiBlock)
	require.Equal(t, big.NewInt(0), cfg.CancunBlock)
	require.Equal(t, big.NewInt(789), cfg.PragueBlock)
	require.Nil(t, cfg.VinuBLSBlock)
}

func TestEvmChainConfigVinuBLSCanActivateAfterPrague(t *testing.T) {
	before := Upgrades{Berlin: true, London: true, Shanghai: true, Cancun: true, Prague: true}
	after := before
	after.VinuBLS12381 = true
	rules := VinuChainTestNetRules()

	cfg := rules.EvmChainConfig([]UpgradeHeight{
		{Upgrades: before, Height: 0},
		{Upgrades: after, Height: 987},
	})

	require.Equal(t, big.NewInt(0), cfg.ShanghaiBlock)
	require.Equal(t, big.NewInt(0), cfg.CancunBlock)
	require.Equal(t, big.NewInt(0), cfg.PragueBlock)
	require.Equal(t, big.NewInt(987), cfg.VinuBLSBlock)
}

func TestEvmChainConfigVinuChainMainNetForksActive(t *testing.T) {
	// VinuChainMainNetRules() stages Shanghai/Cancun/Prague active for the
	// full-parity mainnet upgrade (2026-05-29). With a single UpgradeHeight at
	// block 0 carrying those flags, EvmChainConfig must resolve all three fork
	// blocks to 0 (active from genesis-of-this-rules-window) rather than nil.
	rules := VinuChainMainNetRules()
	cfg := rules.EvmChainConfig([]UpgradeHeight{{Upgrades: rules.Upgrades, Height: 0}})

	require.Equal(t, big.NewInt(0), cfg.ShanghaiBlock,
		"mainnet rules now stage Shanghai active")
	require.Equal(t, big.NewInt(0), cfg.CancunBlock,
		"mainnet rules now stage Cancun active")
	require.Equal(t, big.NewInt(0), cfg.PragueBlock,
		"mainnet rules now stage Prague/EIP-7702 active")
	require.Nil(t, cfg.VinuBLSBlock,
		"mainnet must not stage VinuBLS12381 until a separate activation")
}

func TestEvmChainConfigEthereumForksDisabledNil(t *testing.T) {
	// Guard the disabled-fork -> nil mapping with an explicit forks-off rules
	// value (MainNetRules is the legacy Fantom constructor that keeps the
	// Ethereum forks off), independent of VinuChain mainnet's staged state.
	rules := MainNetRules()
	cfg := rules.EvmChainConfig([]UpgradeHeight{{Upgrades: rules.Upgrades, Height: 0}})

	require.Nil(t, cfg.ShanghaiBlock,
		"forks-off rules must leave Shanghai inactive (nil block)")
	require.Nil(t, cfg.CancunBlock,
		"forks-off rules must leave Cancun inactive (nil block)")
	require.Nil(t, cfg.PragueBlock,
		"forks-off rules must leave Prague inactive (nil block)")
	require.Nil(t, cfg.VinuBLSBlock,
		"forks-off rules must leave VinuBLS12381 inactive (nil block)")
}

func TestRulesSfcV2Patch2RLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.SfcV2Patch2)
}

func TestVinuChainTestNetRulesQuotaCacheAddress(t *testing.T) {
	require.Equal(t, "0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D", VinuChainTestNetRules().Economy.QuotaCacheAddress.Hex())
}

func TestVinuChainMainNetRulesQuotaCacheAddress(t *testing.T) {
	// Mainnet QuotaCacheAddress must point at the live TransparentUpgradeableProxy
	// 0x1c4269fb...cd0acda6 (the address sealed in live mainnet chaindata and
	// confirmed via vc_getRules), NOT the implementation 0x9D6Aa03a... that
	// DefaultEconomyRules() carries. The constructor value only governs
	// fresh-install replay from genesis (QuotaCacheAddress is governance-protected
	// on the live chain via marshal.go), so pinning the proxy here is the
	// fresh-install-safety correction. See deployment-log.md "Mainnet rules.go
	// stale-impl fix".
	require := require.New(t)
	require.Equal("0x1c4269fBBD4a8254F69383eeF6aF720bCD0aCda6", VinuChainMainNetRules().Economy.QuotaCacheAddress.Hex(),
		"mainnet QuotaCacheAddress must be the proxy, not the implementation")
	require.NotEqual("0x9D6Aa03a8D4AcF7b43c562f349Ee45b3214c3bbF", VinuChainMainNetRules().Economy.QuotaCacheAddress.Hex(),
		"mainnet QuotaCacheAddress must not be the stale implementation address")
}

// TestVinuChainMainNetRulesUpgradeFlags pins the staged mainnet full-parity
// upgrade set (decided 2026-05-29). It is the single guardrail that documents
// exactly which flags the next mainnet binary will carry. Update it
// deliberately when the mainnet upgrade scope changes — never to make a build
// pass by accident.
func TestVinuChainMainNetRulesUpgradeFlags(t *testing.T) {
	require := require.New(t)
	up := VinuChainMainNetRules().Upgrades

	// Active for the full-parity upgrade.
	require.True(up.Berlin, "Berlin")
	require.True(up.London, "London")
	require.True(up.Shanghai, "Shanghai")
	require.True(up.Cancun, "Cancun")
	require.True(up.Prague, "Prague")
	require.True(up.Llr, "Llr")
	require.True(up.Podgorica, "Podgorica")
	require.True(up.SfcV2, "SfcV2")
	require.True(up.Elemont, "Elemont")
	require.True(up.ElemontPubkeyValidation, "ElemontPubkeyValidation")

	// Intentionally NOT set on mainnet: SfcV2Patch* are testnet-only re-flash
	// flags (mainnet's first SfcV2 activation installs GetLatestContractBin,
	// which already contains every later bytecode fix). PaybackV2/Patch are a
	// separate, later mainnet release blocked on deploying QuotaContractV2 and
	// baking paybackV2MainnetAddress (EnforcePaybackV2StartupCheck panics if the
	// flag is set while the address is the zero sentinel).
	require.False(up.SfcV2Patch, "SfcV2Patch must stay false on mainnet")
	require.False(up.SfcV2Patch2, "SfcV2Patch2 must stay false on mainnet")
	require.False(up.SfcV2Patch3, "SfcV2Patch3 must stay false on mainnet")
	require.False(up.SfcV2Patch4, "SfcV2Patch4 must stay false on mainnet")
	require.False(up.SfcV2Patch5, "SfcV2Patch5 must stay false on mainnet")
	require.False(up.SfcV2Patch6, "SfcV2Patch6 must stay false on mainnet")
	require.False(up.VinuBLS12381, "VinuBLS12381 must stay false on mainnet until its separate testnet-first rollout")
	require.False(up.PaybackV2, "PaybackV2 must stay false on mainnet until its separate release")
	require.False(up.PaybackV2Patch, "PaybackV2Patch must stay false on mainnet until its separate release")
}

func TestRulesSfcV2Patch3RLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.SfcV2Patch3)
}

func TestRulesSfcV2Patch4RLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.SfcV2Patch4)
}

func TestRulesSfcV2Patch4FalseRLP(t *testing.T) {
	rules := MainNetRules()
	require := require.New(t)

	// Sanity — this is the network that specifically does NOT enable Patch4.
	require.False(rules.Upgrades.SfcV2Patch4, "MainNetRules should not have SfcV2Patch4 enabled; update this test if mainnet activates Patch4")

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.False(decodedRules.Upgrades.SfcV2Patch4, "Upgrades.SfcV2Patch4 must round-trip as false through RLP")
}

func TestRulesElemontPubkeyValidationTrueRLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	rules.Upgrades.ElemontPubkeyValidation = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.ElemontPubkeyValidation,
		"Upgrades.ElemontPubkeyValidation must round-trip as true through RLP")
}

func TestRulesElemontPubkeyValidationDefaultsRLP(t *testing.T) {
	// The legacy Fantom MainNetRules constructor leaves ElemontPubkeyValidation
	// false; VinuChainMainNetRules() stages it true for the full-parity mainnet
	// upgrade (2026-05-29), matching testnet (which activated it in
	// v2.0.14-elemont alongside SfcV2Patch5). All defaults must round-trip
	// bit-for-bit through RLP.
	cases := []struct {
		mk      func() Rules
		expect  bool
		network string
	}{
		{MainNetRules, false, "MainNetRules"},
		{VinuChainMainNetRules, true, "VinuChainMainNetRules"},
		{VinuChainTestNetRules, true, "VinuChainTestNetRules"},
	}
	for _, c := range cases {
		rules := c.mk()
		require := require.New(t)
		require.Equal(c.expect, rules.Upgrades.ElemontPubkeyValidation,
			"%s ElemontPubkeyValidation default mismatch", c.network)

		b, err := rlp.EncodeToBytes(rules)
		require.NoError(err)

		decodedRules := Rules{}
		require.NoError(rlp.DecodeBytes(b, &decodedRules))

		require.Equal(rules.String(), decodedRules.String())
		require.Equal(c.expect, decodedRules.Upgrades.ElemontPubkeyValidation,
			"%s Upgrades.ElemontPubkeyValidation must round-trip through RLP", c.network)
	}
}

func TestRulesSfcV2Patch5TrueRLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	rules.Upgrades.SfcV2Patch5 = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.SfcV2Patch5,
		"Upgrades.SfcV2Patch5 must round-trip as true through RLP")
}

func TestRulesSfcV2Patch5DefaultsRLP(t *testing.T) {
	// Mainnet and legacy constructors leave SfcV2Patch5 false — mainnet has
	// not yet activated SfcV2 and will install the latest available bytecode
	// directly on its first activation, so re-flash flags are unnecessary.
	// Testnet activated the flag in v2.0.14-elemont alongside the Cycle-161
	// re-flash. Both defaults must round-trip bit-for-bit through RLP.
	cases := []struct {
		mk      func() Rules
		expect  bool
		network string
	}{
		{MainNetRules, false, "MainNetRules"},
		{VinuChainMainNetRules, false, "VinuChainMainNetRules"},
		{VinuChainTestNetRules, true, "VinuChainTestNetRules"},
	}
	for _, c := range cases {
		rules := c.mk()
		require := require.New(t)
		require.Equal(c.expect, rules.Upgrades.SfcV2Patch5,
			"%s SfcV2Patch5 default mismatch", c.network)

		b, err := rlp.EncodeToBytes(rules)
		require.NoError(err)

		decodedRules := Rules{}
		require.NoError(rlp.DecodeBytes(b, &decodedRules))

		require.Equal(rules.String(), decodedRules.String())
		require.Equal(c.expect, decodedRules.Upgrades.SfcV2Patch5,
			"%s Upgrades.SfcV2Patch5 must round-trip through RLP", c.network)
	}
}

func TestRulesSfcV2Patch6TrueRLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	rules.Upgrades.SfcV2Patch6 = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.SfcV2Patch6,
		"Upgrades.SfcV2Patch6 must round-trip as true through RLP")
}

func TestRulesSfcV2Patch6DefaultsRLP(t *testing.T) {
	// Mainnet constructors leave SfcV2Patch6 false. Mainnet has not yet
	// activated SfcV2 and will install the latest available bytecode directly
	// on its first activation, so re-flash flags are unnecessary. Testnet
	// activates Patch6 as the Cycle-162 orphan-delegation backfill edge.
	cases := []struct {
		mk      func() Rules
		expect  bool
		network string
	}{
		{MainNetRules, false, "MainNetRules"},
		{VinuChainMainNetRules, false, "VinuChainMainNetRules"},
		{VinuChainTestNetRules, true, "VinuChainTestNetRules"},
	}
	for _, c := range cases {
		rules := c.mk()
		require := require.New(t)
		require.Equal(c.expect, rules.Upgrades.SfcV2Patch6,
			"%s SfcV2Patch6 default mismatch", c.network)

		b, err := rlp.EncodeToBytes(rules)
		require.NoError(err)

		decodedRules := Rules{}
		require.NoError(rlp.DecodeBytes(b, &decodedRules))

		require.Equal(rules.String(), decodedRules.String())
		require.Equal(c.expect, decodedRules.Upgrades.SfcV2Patch6,
			"%s Upgrades.SfcV2Patch6 must round-trip through RLP", c.network)
	}
}

func TestRulesPaybackV2PatchTrueRLP(t *testing.T) {
	rules := VinuChainTestNetRules()
	rules.Upgrades.PaybackV2Patch = true
	require := require.New(t)

	b, err := rlp.EncodeToBytes(rules)
	require.NoError(err)

	decodedRules := Rules{}
	require.NoError(rlp.DecodeBytes(b, &decodedRules))

	require.Equal(rules.String(), decodedRules.String())
	require.True(decodedRules.Upgrades.PaybackV2Patch,
		"Upgrades.PaybackV2Patch must round-trip as true through RLP")
}

func TestRulesPaybackV2PatchDefaultsRLP(t *testing.T) {
	cases := []struct {
		mk      func() Rules
		expect  bool
		network string
	}{
		{MainNetRules, false, "MainNetRules"},
		{VinuChainMainNetRules, false, "VinuChainMainNetRules"},
		{VinuChainTestNetRules, true, "VinuChainTestNetRules"},
	}
	for _, c := range cases {
		rules := c.mk()
		require := require.New(t)
		require.Equal(c.expect, rules.Upgrades.PaybackV2Patch,
			"%s PaybackV2Patch default mismatch", c.network)

		b, err := rlp.EncodeToBytes(rules)
		require.NoError(err)

		decodedRules := Rules{}
		require.NoError(rlp.DecodeBytes(b, &decodedRules))

		require.Equal(rules.String(), decodedRules.String())
		require.Equal(c.expect, decodedRules.Upgrades.PaybackV2Patch,
			"%s Upgrades.PaybackV2Patch must round-trip through RLP", c.network)
	}
}

func TestRulesBerlinCompatibilityRLP(t *testing.T) {
	require := require.New(t)

	b1, err := rlp.EncodeToBytes(Upgrades{
		Berlin: true,
	})
	require.NoError(err)

	b2, err := rlp.EncodeToBytes(struct {
		Berlin bool
	}{true})
	require.NoError(err)

	require.Equal(b2, b1)
}

func TestGasRulesLLRCompatibilityRLP(t *testing.T) {
	require := require.New(t)

	b1, err := rlp.EncodeToBytes(GasRules{
		MaxEventGas:          1,
		EventGas:             2,
		ParentGas:            3,
		ExtraDataGas:         4,
		BlockVotesBaseGas:    0,
		BlockVoteGas:         0,
		EpochVoteGas:         0,
		MisbehaviourProofGas: 0,
	})
	require.NoError(err)

	b2, err := rlp.EncodeToBytes(struct {
		MaxEventGas  uint64
		EventGas     uint64
		ParentGas    uint64
		ExtraDataGas uint64
	}{1, 2, 3, 4})
	require.NoError(err)

	require.Equal(b2, b1)
}
