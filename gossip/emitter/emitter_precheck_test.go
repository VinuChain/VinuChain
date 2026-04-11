package emitter

import (
	"math"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/emitter/ancestor"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/piecefunc"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter"
)

// TestIsAllowedToEmit_IdleGate_OptimisticPreCheck documents the contract for
// the pre-check in createEvent. The pre-check intentionally passes eTxs=true
// (optimistic: "assume I will have txs") so that an idle validator with
// pending txs in the pool is not blocked by the idle gate. The real gate is
// the post-check with the actual tx count. Reverts audit finding E-04
// (cycle 53), which flipped the pre-check to eTxs=false and caused idle
// validators to stall pending txs until the Max-interval force-emit fired
// (60-second tx inclusion cliff on testnet).
//
// With eTxs=true, isAllowedToEmit must return true for an idle validator
// whose previous emission was recent (well under intervals.Max).
// With eTxs=false, isAllowedToEmit must return false under the same
// conditions — this is the regression the pre-check was bypassing.
func TestIsAllowedToEmit_IdleGate_OptimisticPreCheck(t *testing.T) {
	em, external, _, _, _ := newTestEmitter(t)

	// isAllowedToEmit queries the world for the latest block index to
	// compute passedBlocks. Return 0 so passedBlocks == 0 and no
	// block-slack-based force-emit fires.
	external.EXPECT().GetLatestBlockIndex().Return(idx.Block(0)).AnyTimes()

	// Set emitter state so we land squarely inside the idle-gate check:
	//   - previous emission 200ms ago: passedTime ~ 200ms, which is
	//     >= intervals.Min (110ms) and << intervals.Max (~55s after
	//     RandomizeEmitTime). That skips both the Min-interval gate and
	//     the Max-interval force-emit.
	//   - originatedTxs empty (default after NewEmitter) so idle() is true.
	//   - stakeRatio for our validator set to 0 so the "top validator"
	//     branch collapses passedTimeIdle onto passedTime.
	now := time.Now()
	em.prevEmittedAtTime = now.Add(-200 * time.Millisecond)
	em.prevIdleTime = now.Add(-200 * time.Millisecond)
	em.prevEmittedAtBlock = 0
	em.stakeRatio = map[idx.ValidatorID]uint64{em.config.Validator.ID: 0}

	// Build a minimal mutable event owned by our validator with plenty of
	// gas power so the low-power branches are skipped.
	e := &inter.MutableEventPayload{}
	e.SetCreator(em.config.Validator.ID)
	e.SetGasPowerLeft(inter.GasPowerLeft{
		Gas: [inter.GasPowerConfigs]uint64{math.MaxUint64, math.MaxUint64},
	})

	// metric = 1.0 (full DecimalUnit) so adjustedPassedTime tracks
	// passedTime and no efficiency-metric branch trips.
	metric := ancestor.Metric(piecefunc.DecimalUnit)

	require.True(t, em.idle(),
		"precondition: emitter must be idle (empty originatedTxs buffer)")

	// Optimistic pre-check: eTxs=true must pass the idle gate even though
	// idle()==true and passedTime << intervals.Max. This is the behaviour
	// createEvent relies on so that addTxs can attach pending pool txs.
	require.True(t,
		em.isAllowedToEmit(e, true, metric, nil),
		"eTxs=true: idle validator with recent prior emission must be allowed "+
			"— the pre-check is intentionally optimistic so addTxs can run")

	// Pessimistic pre-check: eTxs=false must be blocked by the idle gate
	// (passedTime < intervals.Max && em.idle() && !eTxs). This is the
	// regression path: when createEvent passed eTxs=false here, every
	// pre-check returned nil, pending txs never reached addTxs, and
	// inclusion stalled until the Max-interval force-emit fired.
	require.False(t,
		em.isAllowedToEmit(e, false, metric, nil),
		"eTxs=false: idle gate must block emission — if this returns true "+
			"the idle gate is broken, if createEvent's pre-check passes "+
			"eTxs=false here the testnet tx-inclusion stall regresses")
}

// TestCreateEvent_PreCheckUsesOptimisticETxs is a source-level regression
// guard for audit finding E-04 (cycle 53). The pre-check in createEvent must
// invoke isAllowedToEmit with eTxs=true. Combined with
// TestIsAllowedToEmit_IdleGate_OptimisticPreCheck (which pins the contract
// of isAllowedToEmit under idle+eTxs=true vs idle+eTxs=false), this guards
// against re-regression of the 60-second testnet tx-inclusion stall that
// shipped when the pre-check was flipped to eTxs=false.
//
// Rationale for a source-level test: exercising the full createEvent path
// requires a live quorum indexer, a non-nil DagIndex, and a selfParent —
// scaffolding that is fragile and noisy. The pre-check is a one-line
// invariant and is worth pinning directly.
func TestCreateEvent_PreCheckUsesOptimisticETxs(t *testing.T) {
	src, err := os.ReadFile("emitter.go")
	require.NoError(t, err, "must be able to read emitter.go from package dir")

	// Match the first isAllowedToEmit(...) call in createEvent — the
	// pre-check on the line following the "Pre-check if event should be
	// emitted" comment block. We assert it passes eTxs=true.
	preCheck := regexp.MustCompile(
		`Pre-check if event should be emitted[^\n]*\n(?:[^\n]*\n){1,10}?\s*if !em\.isAllowedToEmit\(mutEvent,\s*(true|false),`)
	match := preCheck.FindSubmatch(src)
	require.NotNil(t, match, "could not locate the createEvent pre-check in emitter.go")
	require.Equal(t, "true", string(match[1]),
		"createEvent pre-check must pass eTxs=true (optimistic): "+
			"eTxs=false causes idle validators to stall pending txs until "+
			"the Max-interval force-emit fires (E-04 regression, cycle 53). "+
			"See TestIsAllowedToEmit_IdleGate_OptimisticPreCheck for the "+
			"contract this guard protects.")
}
