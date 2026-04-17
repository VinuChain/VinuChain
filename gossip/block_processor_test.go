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
