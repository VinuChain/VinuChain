package gaspowercheck

import (
	"math"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
)

// mockEvent is a minimal implementation of inter.EventI for testing calcGasPower.
type mockEvent struct {
	epoch        idx.Epoch
	creator      idx.ValidatorID
	selfParent   *hash.Event
	medianTime   inter.Timestamp
	gasPowerLeft inter.GasPowerLeft
	gasPowerUsed uint64
}

// dag.Event interface
func (e *mockEvent) Epoch() idx.Epoch           { return e.epoch }
func (e *mockEvent) Creator() idx.ValidatorID   { return e.creator }
func (e *mockEvent) SelfParent() *hash.Event    { return e.selfParent }
func (e *mockEvent) MedianTime() inter.Timestamp { return e.medianTime }
func (e *mockEvent) GasPowerLeft() inter.GasPowerLeft { return e.gasPowerLeft }
func (e *mockEvent) GasPowerUsed() uint64        { return e.gasPowerUsed }
func (e *mockEvent) Seq() idx.Event              { return 1 }
func (e *mockEvent) Frame() idx.Frame            { return 0 }
func (e *mockEvent) Lamport() idx.Lamport        { return 0 }
func (e *mockEvent) Parents() hash.Events        { return nil }
func (e *mockEvent) IsSelfParent(hash.Event) bool { return false }
func (e *mockEvent) ID() hash.Event              { return hash.Event{} }
func (e *mockEvent) String() string              { return "mockEvent" }
func (e *mockEvent) Size() int                   { return 0 }

// inter.EventI additional interface
func (e *mockEvent) Version() uint8                    { return 1 }
func (e *mockEvent) NetForkID() uint16                 { return 0 }
func (e *mockEvent) CreationTime() inter.Timestamp     { return e.medianTime }
func (e *mockEvent) PrevEpochHash() *hash.Hash         { return nil }
func (e *mockEvent) Extra() []byte                     { return nil }
func (e *mockEvent) HashToSign() hash.Hash             { return hash.Hash{} }
func (e *mockEvent) Locator() inter.EventLocator       { return inter.EventLocator{} }
func (e *mockEvent) AnyTxs() bool                      { return false }
func (e *mockEvent) AnyBlockVotes() bool               { return false }
func (e *mockEvent) AnyEpochVote() bool                { return false }
func (e *mockEvent) AnyMisbehaviourProofs() bool       { return false }
func (e *mockEvent) PayloadHash() hash.Hash            { return hash.Hash{} }

// TestCalcGasPower_GasRefundSaturation verifies that when a validator's GasRefund
// is math.MaxUint64 (saturated by DirtyGasRefund) and prevGasPowerLeft > 0, the
// addition saturates at math.MaxUint64 rather than wrapping around.
//
// The observable symptom before the fix: the wrapped prevGasPowerLeft of 0 propagates
// into CalcValidatorGasPower, which (with zero elapsed time) yields a final gas power
// of 0. After the fix, prevGasPowerLeft saturates at MaxUint64 and CalcValidatorGasPower
// caps the result at maxGasPower (3_600_000_000 = AllocPerSec 1_000_000 × MaxAllocPeriod 3600s).
//
// Using prevTime == medianTime (zero elapsed time) ensures that no freshly allocated
// gas obscures the difference between the two code paths.
func TestCalcGasPower_GasRefundSaturation(t *testing.T) {
	const validatorID idx.ValidatorID = 1
	const epochNum idx.Epoch = 2
	const prevGasPowerLeftVal uint64 = 1 // non-zero: this triggers wrap-around before fix

	// Single validator with non-zero stake so CalcValidatorGasPowerPerSec is meaningful.
	builder := pos.NewBuilder()
	builder.Set(validatorID, pos.Weight(1_000_000))
	validators := builder.Build()

	// Timestamp shared by both prevEvent and the new event so that elapsed time = 0,
	// meaning zero gas is allocated. This isolates the refund-addition path.
	const sharedTimestamp inter.Timestamp = 1_000_000

	// prevEvent carries prevGasPowerLeft from the previous epoch.
	prevEvent := iblockproc.EventInfo{
		GasPowerLeft: inter.GasPowerLeft{
			Gas: [inter.GasPowerConfigs]uint64{prevGasPowerLeftVal, prevGasPowerLeftVal},
		},
		Time: sharedTimestamp,
	}
	// Set ID to non-zero so the branch `if validatorState.PrevEpochEvent.ID != hash.ZeroEvent`
	// is entered, loading prevGasPowerLeft from the previous epoch event.
	prevEvent.ID[0] = 0x01

	validatorState := ValidatorState{
		PrevEpochEvent: prevEvent,
		GasRefund:      math.MaxUint64, // saturated sentinel — triggers wrap without fix
	}

	// StartupAllocPeriod=0 keeps startup gas = 0 so the startup floor cannot rescue
	// a wrapped prevGasPowerLeft. maxGasPower = AllocPerSec × MaxAllocPeriod = 3_600_000_000.
	cfg := Config{
		Idx:                0,
		AllocPerSec:        1_000_000,
		MaxAllocPeriod:     inter.Timestamp(time.Hour),
		MinEnsuredAlloc:    1_000_000,
		StartupAllocPeriod: 0,
		MinStartupGas:      0,
		Podgorica:          false,
	}

	ctx := &ValidationContext{
		Epoch:           epochNum,
		Configs:         [inter.GasPowerConfigs]Config{cfg, cfg},
		EpochStart:      sharedTimestamp,
		Validators:      validators,
		ValidatorStates: []ValidatorState{validatorState},
	}

	// New event: no selfParent (epoch-first event) so the GasRefund branch is taken.
	// medianTime == prevEvent.Time → zero elapsed time → zero allocated gas.
	event := &mockEvent{
		epoch:      epochNum,
		creator:    validatorID,
		selfParent: nil,
		medianTime: sharedTimestamp,
	}

	// calcGasPower is package-private; call it directly from the test.
	got := calcGasPower(event, nil, ctx, cfg)

	// With saturation fix:
	//   prevGasPowerLeft = 1 + MaxUint64 → saturates to MaxUint64
	//   CalcValidatorGasPower: 0 allocated + MaxUint64 → capped at maxGasPower = 3_600_000_000
	//   result = 3_600_000_000 >= wantMin (1_000_000)
	//
	// Without saturation fix:
	//   prevGasPowerLeft = 1 + MaxUint64 → wraps to 0
	//   CalcValidatorGasPower: 0 allocated + 0 = 0 (below maxGasPower, no cap)
	//   result = 0 < 1_000_000 → test fails, exposing the bug
	const wantMin uint64 = 1_000_000
	if got < wantMin {
		t.Errorf("calcGasPower = %d, want >= %d: uint64 overflow wrap-around in GasRefund addition (saturation fix missing)",
			got, wantMin)
	}
}

// TestCalcGasPower_GasRefundSaturation_NonZeroElapsedTime verifies that when
// prevGasPowerLeft is saturated to math.MaxUint64 (from DirtyGasRefund accumulation)
// and the event has non-zero elapsed time (allocating additional gas), CalcValidatorGasPower
// does not overflow when adding the allocated gas to the saturated prevGasPowerLeft.
//
// Without the clamp inside CalcValidatorGasPower, the unsigned addition
//
//	allocatedGas (e.g. 1_000_000) + math.MaxUint64
//
// wraps to a near-zero value, causing the validator's events to be rejected by peers.
func TestCalcGasPower_GasRefundSaturation_NonZeroElapsedTime(t *testing.T) {
	const validatorID idx.ValidatorID = 1
	const epochNum idx.Epoch = 2
	const prevGasPowerLeftVal uint64 = 1

	builder := pos.NewBuilder()
	builder.Set(validatorID, pos.Weight(1_000_000))
	validators := builder.Build()

	// Non-zero elapsed time: 1 second between prevEvent and new event.
	// This causes CalcValidatorGasPower to allocate 1_000_000 gas (AllocPerSec).
	const prevTimestamp inter.Timestamp = 1_000_000
	const currTimestamp inter.Timestamp = prevTimestamp + inter.Timestamp(time.Second)

	prevEvent := iblockproc.EventInfo{
		GasPowerLeft: inter.GasPowerLeft{
			Gas: [inter.GasPowerConfigs]uint64{prevGasPowerLeftVal, prevGasPowerLeftVal},
		},
		Time: prevTimestamp,
	}
	prevEvent.ID[0] = 0x01

	validatorState := ValidatorState{
		PrevEpochEvent: prevEvent,
		GasRefund:      math.MaxUint64,
	}

	cfg := Config{
		Idx:                0,
		AllocPerSec:        1_000_000,
		MaxAllocPeriod:     inter.Timestamp(time.Hour),
		MinEnsuredAlloc:    1_000_000,
		StartupAllocPeriod: 0,
		MinStartupGas:      0,
		Podgorica:          false,
	}

	ctx := &ValidationContext{
		Epoch:           epochNum,
		Configs:         [inter.GasPowerConfigs]Config{cfg, cfg},
		EpochStart:      prevTimestamp,
		Validators:      validators,
		ValidatorStates: []ValidatorState{validatorState},
	}

	event := &mockEvent{
		epoch:      epochNum,
		creator:    validatorID,
		selfParent: nil,
		medianTime: currTimestamp,
	}

	got := calcGasPower(event, nil, ctx, cfg)

	// maxGasPower = AllocPerSec(1_000_000) × MaxAllocPeriod(3600s) = 3_600_000_000.
	// After clamp: prevGasPowerLeft = 3_600_000_000; allocated = 1_000_000;
	// sum = 4_600_000_000 > maxGasPower → caps to 3_600_000_000.
	//
	// Without the clamp: prevGasPowerLeft = math.MaxUint64; allocated(1_000_000) + MaxUint64
	// wraps to 999_999; result = 999_999, far below any reasonable threshold.
	const maxGasPower uint64 = 3_600_000_000
	if got != maxGasPower {
		t.Errorf("calcGasPower = %d, want %d: CalcValidatorGasPower overflows when prevGasPowerLeft is math.MaxUint64 and elapsed time is non-zero",
			got, maxGasPower)
	}
}
