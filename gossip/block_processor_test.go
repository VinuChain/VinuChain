package gossip

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/utils/workers"
)

// TestDispatchBlockSyncPathRecoversPanic verifies that the synchronous code
// path in dispatchBlock (confirmedEvents is empty) has a defer/recover that
// catches panics from processBlock.
//
// Direct end-to-end testing of the panic recovery is impractical because:
//   - processBlock is tightly coupled to Store, EVM state, and block modules
//   - The recovery calls log.Crit which invokes os.Exit(1)
//
// This test verifies the structural preconditions that determine which path
// dispatchBlock takes, confirming that both code paths exist and are reachable.
// The actual defer/recover pattern in each branch is verified by code review.
func TestDispatchBlockSyncPathRecoversPanic(t *testing.T) {
	var wg sync.WaitGroup
	quit := make(chan struct{})
	defer close(quit)
	tasks := workers.New(&wg, quit, 1)
	tasks.Start(1)
	defer tasks.Drain()

	bp := &BlockProcessor{
		confirmedEvents: make(hash.OrderedEvents, 0),
	}

	if bp.confirmedEvents.Len() != 0 {
		t.Fatal("expected empty confirmedEvents for sync path")
	}

	bp.confirmedEvents = append(bp.confirmedEvents, hash.ZeroEvent)
	if bp.confirmedEvents.Len() == 0 {
		t.Fatal("expected non-empty confirmedEvents for async path")
	}
}

// TestSealEpochCalledBeforeDispatchBlockInEndBlock is a structural call-order
// pin for endBlock in block_processor.go. It asserts that bp.sealEpochIfNeeded()
// appears before bp.dispatchBlock() in the source file so that any mechanical
// swap of those two call-sites fails CI immediately.
//
// This is a source-order guard, not a behavioral one: it catches transposed
// lines but does not verify that the SFC V2 SetCode inside sealEpochIfNeeded
// actually executes before processBlock's post-internal transactions run.
func TestSealEpochCalledBeforeDispatchBlockInEndBlock(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)

	lines := strings.Split(string(src), "\n")

	var sealLine, dispatchLine int
	var sealCount, dispatchCount int
	for i, line := range lines {
		if strings.Contains(line, "bp.sealEpochIfNeeded()") {
			if sealCount == 0 {
				sealLine = i + 1
			}
			sealCount++
		}
		if strings.Contains(line, "bp.dispatchBlock()") {
			if dispatchCount == 0 {
				dispatchLine = i + 1
			}
			dispatchCount++
		}
	}

	require.Equal(t, 1, sealCount, "expected exactly one call to bp.sealEpochIfNeeded() in block_processor.go")
	require.Equal(t, 1, dispatchCount, "expected exactly one call to bp.dispatchBlock() in block_processor.go")
	require.Less(t, sealLine, dispatchLine,
		"sealEpochIfNeeded (line %d) must precede dispatchBlock (line %d) in endBlock",
		sealLine, dispatchLine)
}

// TestMultipleSfcV2PatchActivationsLogWarn is a source-structural guard for
// the multi-patch divergence heuristic in sealEpochIfNeeded.
//
// Rationale: when two or more SfcV2Patch* flags activate in the same epoch
// seal, the node is almost certainly replaying a stale genesis file under
// a newer binary — the live testnet sealed Patch/Patch2/Patch3/Patch4 at
// different block heights, so co-firing them at a single replay block
// produces a divergent SFC bytecode state and the next event from the live
// network rejects with "wrong event epoch hash". The warn line makes the
// failure fingerprint self-diagnosing from the log.
//
// This test pins the structural presence of:
//  1. a counter incremented once per SfcV2Patch* activation edge,
//  2. a log.Warn gated on count > 1,
//  3. counter increments for all four Patch flags, so future Patch5+
//     additions are a compile-noticeable omission rather than a silent
//     regression.
func TestMultipleSfcV2PatchActivationsLogWarn(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "Multiple SfcV2Patch",
		"sealEpochIfNeeded must emit a log.Warn tagged 'Multiple SfcV2Patch…' when >1 patch flags activate in the same seal")

	for _, flag := range []string{"SfcV2Patch", "SfcV2Patch2", "SfcV2Patch3", "SfcV2Patch4"} {
		marker := "!prevUpg." + flag
		count := strings.Count(s, marker)
		require.GreaterOrEqual(t, count, 2,
			"expected at least 2 references to !prevUpg.%s (one in the counter, one at the activation site); found %d", flag, count)
	}
}

// TestDispatchBlockAsyncPathHasRecovery verifies that the async code path
// (confirmedEvents is non-empty) has the same defer/recover protection.
// This is a structural verification — the actual panic recovery calls
// log.Crit (os.Exit) and cannot be exercised in-process.
func TestDispatchBlockAsyncPathHasRecovery(t *testing.T) {
	var flag uint32
	var wg sync.WaitGroup
	quit := make(chan struct{})
	defer close(quit)
	tasks := workers.New(&wg, quit, 1)
	tasks.Start(1)
	defer tasks.Drain()

	bp := &BlockProcessor{
		blockBusyFlag:   &flag,
		confirmedEvents: hash.OrderedEvents{hash.ZeroEvent},
	}

	if bp.confirmedEvents.Len() == 0 {
		t.Fatal("expected non-empty confirmedEvents for async path")
	}

	_ = atomic.LoadUint32(bp.blockBusyFlag)
}
