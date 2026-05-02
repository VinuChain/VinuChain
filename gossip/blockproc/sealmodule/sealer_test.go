package sealmodule

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/lachesis"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/drivertype"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
)

// newTestSealer constructs a sealer with a configurable set of existing
// validators and a next-epoch profile map. Callers override individual
// BlockState / EpochState fields after construction.
func newTestSealer(t *testing.T, rules opera.Rules, existingIDs []idx.ValidatorID, next iblockproc.ValidatorProfiles) *OperaEpochsSealer {
	t.Helper()
	validators := pos.EqualWeightValidators(existingIDs, 100)
	states := make([]iblockproc.ValidatorBlockState, len(existingIDs))
	for i := range states {
		states[i] = iblockproc.ValidatorBlockState{
			Originated:     new(big.Int),
			DirtyGasRefund: 1_000,
			Uptime:         inter.Timestamp(5_000),
			LastBlock:      idx.Block(42),
			LastOnlineTime: inter.Timestamp(100),
		}
	}
	bs := iblockproc.BlockState{
		ValidatorStates:       states,
		NextValidatorProfiles: next,
		EpochGas:              0,
	}
	es := iblockproc.EpochState{
		Epoch:          5,
		EpochStart:     inter.Timestamp(1_000),
		PrevEpochStart: inter.Timestamp(500),
		Validators:     validators,
		Rules:          rules,
	}
	return &OperaEpochsSealer{
		block: iblockproc.BlockCtx{Idx: 100, Time: inter.Timestamp(10_000)},
		es:    es,
		bs:    bs,
	}
}

func defaultRules() opera.Rules {
	return opera.Rules{
		Epochs: opera.EpochsRules{
			MaxEpochGas:      1_000_000,
			MaxEpochDuration: inter.Timestamp(60_000),
		},
	}
}

func profile(weight int64, pubkey []byte) drivertype.Validator {
	return drivertype.Validator{
		Weight: big.NewInt(weight),
		PubKey: validatorpk.PubKey{Type: validatorpk.Types.Secp256k1, Raw: pubkey},
	}
}

// --- EpochSealing triggers ------------------------------------------------

func TestEpochSealing_GasThresholdReached(t *testing.T) {
	rules := defaultRules()
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{})
	s.bs.EpochGas = rules.Epochs.MaxEpochGas
	require.True(t, s.EpochSealing(), "gas >= MaxEpochGas must seal")
}

func TestEpochSealing_DurationExceeded(t *testing.T) {
	rules := defaultRules()
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{})
	s.es.EpochStart = inter.Timestamp(1_000)
	s.block.Time = s.es.EpochStart + rules.Epochs.MaxEpochDuration
	require.True(t, s.EpochSealing(), "duration >= MaxEpochDuration must seal")
}

func TestEpochSealing_AdvanceEpochsPending(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{})
	s.bs.AdvanceEpochs = 3
	require.True(t, s.EpochSealing(), "AdvanceEpochs > 0 must seal")
}

func TestEpochSealing_CheatersPresent(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1, 2}, iblockproc.ValidatorProfiles{})
	s.bs.EpochCheaters = lachesis.Cheaters{idx.ValidatorID(2)}
	require.True(t, s.EpochSealing(), "cheaters present must seal")
}

func TestEpochSealing_NoConditionsMet(t *testing.T) {
	rules := defaultRules()
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{})
	s.bs.EpochGas = rules.Epochs.MaxEpochGas - 1
	s.block.Time = s.es.EpochStart + 1 // far below MaxEpochDuration
	require.False(t, s.EpochSealing(), "below thresholds must not seal")
}

// --- SealEpoch mechanics --------------------------------------------------

func TestSealEpoch_BumpsEpochAndResetsGasAndCheaters(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0xAA, 0xBB}),
	})
	s.bs.EpochGas = 900_000
	s.bs.EpochCheaters = lachesis.Cheaters{idx.ValidatorID(1)}
	s.bs.CheatersWritten = 2

	oldEpoch := s.es.Epoch
	newBS, newES := s.SealEpoch()

	require.Equal(t, oldEpoch+1, newES.Epoch)
	require.Equal(t, uint64(0), newBS.EpochGas)
	require.Empty(t, newBS.EpochCheaters)
	require.Equal(t, uint32(0), newBS.CheatersWritten)
	require.Equal(t, s.block.Time, newES.EpochStart)
}

func TestSealEpoch_ExistingValidatorInheritsStateAndResetsDirtyGas(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1, 2}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
		2: profile(100, []byte{0x02}),
	})
	// Seed ValidatorStates for old validators 1 and 2.
	s.bs.ValidatorStates[0].DirtyGasRefund = 777
	s.bs.ValidatorStates[0].Uptime = inter.Timestamp(50)
	s.bs.ValidatorStates[0].LastEvent = iblockproc.EventInfo{Time: inter.Timestamp(60)}
	s.bs.ValidatorStates[1].DirtyGasRefund = 888

	newBS, newES := s.SealEpoch()

	require.Len(t, newBS.ValidatorStates, 2)
	for _, v := range newBS.ValidatorStates {
		require.Equal(t, uint64(0), v.DirtyGasRefund, "DirtyGasRefund must reset at epoch seal")
		require.Equal(t, inter.Timestamp(0), v.Uptime, "Uptime must reset at epoch seal")
	}
	// Old DirtyGasRefund becomes new epoch's GasRefund.
	gasRefunds := []uint64{newES.ValidatorStates[0].GasRefund, newES.ValidatorStates[1].GasRefund}
	require.Contains(t, gasRefunds, uint64(777))
	require.Contains(t, gasRefunds, uint64(888))
}

func TestSealEpoch_NewValidatorGetsLastBlockAndTime(t *testing.T) {
	// Start with validator 1 existing; validator 2 appears only in NextValidatorProfiles.
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
		2: profile(100, []byte{0x02}), // new
	})

	newBS, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len())
	// Find which entry corresponds to validator 2 in the new validator set.
	newIdx := newES.Validators.GetIdx(idx.ValidatorID(2))
	require.Equal(t, s.block.Idx, newBS.ValidatorStates[newIdx].LastBlock,
		"new validator LastBlock must equal current block idx")
	require.Equal(t, s.block.Time, newBS.ValidatorStates[newIdx].LastOnlineTime,
		"new validator LastOnlineTime must equal current block time")
	// Originated must be initialized to zero big.Int, not nil.
	require.NotNil(t, newBS.ValidatorStates[newIdx].Originated)
	require.Equal(t, 0, newBS.ValidatorStates[newIdx].Originated.Sign())
}

func TestSealEpoch_ElemontEmptyPubkeyValidatorSkipped(t *testing.T) {
	rules := defaultRules()
	rules.Upgrades.Elemont = true
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),                            // valid pubkey
		2: profile(100, nil),                                     // empty pubkey — must be skipped
		3: profile(200, []byte{0x03, 0x04}),                      // valid pubkey
	})

	_, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len(), "empty-pubkey validator must be skipped under Elemont")
	require.True(t, newES.Validators.Exists(idx.ValidatorID(1)))
	require.False(t, newES.Validators.Exists(idx.ValidatorID(2)))
	require.True(t, newES.Validators.Exists(idx.ValidatorID(3)))
}

func TestSealEpoch_PreElemontEmptyPubkeyValidatorAdmitted(t *testing.T) {
	rules := defaultRules() // Elemont = false
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
		2: profile(100, nil), // empty pubkey — still admitted pre-Elemont
	})

	_, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len(),
		"pre-Elemont behaviour: empty-pubkey validator must still be included")
	require.True(t, newES.Validators.Exists(idx.ValidatorID(2)))
}

// malformedProfile builds a profile with the val-16 testnet shape: type=0x04,
// Raw len=64. The pubkey is non-empty so the existing Elemont empty-pubkey
// guard does NOT skip it; only the new ElemontPubkeyValidation guard catches it.
func malformedProfile(weight int64) drivertype.Validator {
	return drivertype.Validator{
		Weight: big.NewInt(weight),
		PubKey: validatorpk.PubKey{Type: 0x04, Raw: make([]byte, 64)},
	}
}

func TestSealEpoch_ElemontPubkeyValidationMalformedSkipped(t *testing.T) {
	rules := defaultRules()
	rules.Upgrades.Elemont = true
	rules.Upgrades.ElemontPubkeyValidation = true
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, make([]byte, 65)),  // canonical-shape pubkey: type already 0xc0 in profile()
		2: malformedProfile(100),           // val-16 shape: type=0x04, len(Raw)=64 — must be skipped
		3: profile(200, make([]byte, 65)),  // canonical-shape pubkey
	})

	_, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len(),
		"malformed-pubkey validator must be skipped under ElemontPubkeyValidation")
	require.True(t, newES.Validators.Exists(idx.ValidatorID(1)))
	require.False(t, newES.Validators.Exists(idx.ValidatorID(2)))
	require.True(t, newES.Validators.Exists(idx.ValidatorID(3)))
}

func TestSealEpoch_ElemontWithoutPubkeyValidationAdmitsMalformed(t *testing.T) {
	// This is the chain-replay-safety case: the existing testnet chaindata
	// admitted val 16 (malformed pubkey, non-empty). On replay with
	// Elemont=true but ElemontPubkeyValidation=false, the malformed validator
	// must STILL be admitted bit-for-bit identical to today.
	rules := defaultRules()
	rules.Upgrades.Elemont = true
	rules.Upgrades.ElemontPubkeyValidation = false
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, make([]byte, 65)),
		2: malformedProfile(100), // val-16 shape — must STILL be admitted pre-flag
	})

	_, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len(),
		"pre-ElemontPubkeyValidation behaviour: malformed-pubkey validator must be admitted")
	require.True(t, newES.Validators.Exists(idx.ValidatorID(2)))
}

func TestSealEpoch_ElemontPubkeyValidation_StillSkipsEmpty(t *testing.T) {
	// ElemontPubkeyValidation must not regress the empty-pubkey skip.
	rules := defaultRules()
	rules.Upgrades.Elemont = true
	rules.Upgrades.ElemontPubkeyValidation = true
	s := newTestSealer(t, rules, []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, make([]byte, 65)), // canonical: type=0xc0 + 65 raw bytes
		2: profile(100, nil),               // empty — must be skipped
		3: profile(200, make([]byte, 65)), // canonical
	})

	_, newES := s.SealEpoch()

	require.Equal(t, idx.Validator(2), newES.Validators.Len())
	require.False(t, newES.Validators.Exists(idx.ValidatorID(2)),
		"empty-pubkey skip must remain in effect under ElemontPubkeyValidation")
}

func TestSealEpoch_AdvanceEpochsDecrementsByOne(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
	})
	s.bs.AdvanceEpochs = 5

	newBS, _ := s.SealEpoch()
	require.Equal(t, idx.Epoch(4), newBS.AdvanceEpochs)
}

func TestSealEpoch_AdvanceEpochsZeroStaysZero(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
	})
	s.bs.AdvanceEpochs = 0

	newBS, _ := s.SealEpoch()
	require.Equal(t, idx.Epoch(0), newBS.AdvanceEpochs, "must not wrap below zero")
}

func TestSealEpoch_DirtyRulesPromoted(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
	})
	newRules := defaultRules()
	newRules.Epochs.MaxEpochGas = 2_000_000
	s.bs.DirtyRules = &newRules

	newBS, newES := s.SealEpoch()

	require.Equal(t, uint64(2_000_000), newES.Rules.Epochs.MaxEpochGas)
	require.Nil(t, newBS.DirtyRules, "DirtyRules must be consumed on seal")
}

func TestSealEpoch_PrevEpochStartSnapshotted(t *testing.T) {
	s := newTestSealer(t, defaultRules(), []idx.ValidatorID{1}, iblockproc.ValidatorProfiles{
		1: profile(100, []byte{0x01}),
	})
	oldStart := s.es.EpochStart

	_, newES := s.SealEpoch()
	require.Equal(t, oldStart, newES.PrevEpochStart,
		"PrevEpochStart must capture the previous EpochStart before overwrite")
	require.Equal(t, s.block.Time, newES.EpochStart)
}

// --- Start / Update lifecycle --------------------------------------------

func TestStart_InstantiatesProcessor(t *testing.T) {
	m := New()
	require.NotNil(t, m)

	block := iblockproc.BlockCtx{Idx: 7, Time: inter.Timestamp(1234)}
	bs := iblockproc.BlockState{}
	es := iblockproc.EpochState{Rules: defaultRules()}

	p := m.Start(block, bs, es)
	require.NotNil(t, p)
	require.False(t, p.EpochSealing(), "empty state must not trigger sealing")
}

func TestUpdate_RefreshesBlockStateAndEpochState(t *testing.T) {
	m := New()
	block := iblockproc.BlockCtx{Idx: 1}
	p := m.Start(block, iblockproc.BlockState{}, iblockproc.EpochState{Rules: defaultRules()})

	rules := defaultRules()
	rules.Epochs.MaxEpochGas = 42
	newBS := iblockproc.BlockState{EpochGas: 42}
	newES := iblockproc.EpochState{Rules: rules}
	p.Update(newBS, newES)

	require.True(t, p.EpochSealing(), "updated state must be visible to EpochSealing")
}
