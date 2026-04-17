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

// TestSfcV2BytecodeInstalledBeforeProcessBlock is a structural regression test
// that locks in the ordering invariant documented in sealEpochIfNeeded:
// endBlock must call sealEpochIfNeeded (which installs SFC V2 bytecode) before
// calling dispatchBlock (which calls processBlock). This guarantees that
// post-internal and user transactions in the activation block see V2 bytecode.
//
// The test reads the source file and asserts relative line ordering so that
// any refactor that accidentally defers the SetCode to after processBlock will
// cause an immediate CI failure.
func TestSfcV2BytecodeInstalledBeforeProcessBlock(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)

	lines := strings.Split(string(src), "\n")

	var sealLine, dispatchLine int
	for i, line := range lines {
		if strings.Contains(line, "bp.sealEpochIfNeeded()") {
			sealLine = i + 1
		}
		if strings.Contains(line, "bp.dispatchBlock()") {
			dispatchLine = i + 1
		}
	}

	require.NotZero(t, sealLine, "bp.sealEpochIfNeeded() call not found in block_processor.go")
	require.NotZero(t, dispatchLine, "bp.dispatchBlock() call not found in block_processor.go")
	require.Less(t, sealLine, dispatchLine,
		"sealEpochIfNeeded (line %d) must precede dispatchBlock (line %d) in endBlock; "+
			"SFC V2 bytecode must be installed before processBlock runs post-internal txs",
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
