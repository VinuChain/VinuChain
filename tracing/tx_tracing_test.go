package tracing

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestFinishTx_RemovesSpan verifies that FinishTx removes the span so that
// a subsequent FinishTx is a no-op and a subsequent StartTx can register a
// fresh span for the same hash.
func TestFinishTx_RemovesSpan(t *testing.T) {
	SetEnabled(true)
	defer SetEnabled(false)

	h := common.HexToHash("0x1234")

	StartTx(h, "test.enter")

	// Duplicate StartTx must be a no-op (span already tracked).
	StartTx(h, "test.enter2")

	FinishTx(h, "test.exit")

	// Span must be gone: a second FinishTx must not panic or double-finish.
	FinishTx(h, "test.exit2")

	// Re-registering after finish must succeed — proves the span was removed.
	StartTx(h, "test.reenter")

	txSpansMu.Lock()
	_, ok := txSpans[h]
	txSpansMu.Unlock()
	if !ok {
		t.Fatal("expected span to be re-registered after FinishTx + StartTx")
	}

	// Clean up.
	FinishTx(h, "test.cleanup")
}

// TestSetEnabled_DrainsPendingSpans verifies that SetEnabled(false) finishes
// and removes all in-flight spans so they don't leak.
func TestSetEnabled_DrainsPendingSpans(t *testing.T) {
	SetEnabled(true)

	h1 := common.HexToHash("0xaaaa")
	h2 := common.HexToHash("0xbbbb")
	StartTx(h1, "test.enter")
	StartTx(h2, "test.enter")

	txSpansMu.RLock()
	count := len(txSpans)
	txSpansMu.RUnlock()
	if count != 2 {
		t.Fatalf("expected 2 spans before disable, got %d", count)
	}

	SetEnabled(false)

	txSpansMu.RLock()
	count = len(txSpans)
	txSpansMu.RUnlock()
	if count != 0 {
		t.Fatalf("expected 0 spans after SetEnabled(false), got %d", count)
	}
}

// TestFinishTx_SkippedOperation verifies the "skipped" operation tag used by
// BlockProcessor when a transaction is excluded from execution.
func TestFinishTx_SkippedOperation(t *testing.T) {
	SetEnabled(true)
	defer SetEnabled(false)

	h := common.HexToHash("0xdead")

	StartTx(h, "EthAPIBackend.SendTx()")
	FinishTx(h, "BlockProcessor.skipped")

	txSpansMu.Lock()
	_, ok := txSpans[h]
	txSpansMu.Unlock()
	if ok {
		t.Fatal("span must be removed after FinishTx with skipped operation")
	}
}
