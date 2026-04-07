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
