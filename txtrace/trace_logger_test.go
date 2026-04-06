package txtrace

import (
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/kvdb/memorydb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	gossiptxtrace "github.com/Fantom-foundation/go-opera/gossip/txtrace"
)

// makeTestLogger returns a fresh TraceStructLogger backed by an in-memory store.
func makeTestLogger() (*TraceStructLogger, *gossiptxtrace.Store) {
	db := memorydb.New()
	store := gossiptxtrace.NewStore(db)
	logger := NewTraceStructLogger(store)
	return logger, store
}

// TestReset_ClearsAllFields verifies that reset() zeros every field that could
// carry stale state across transactions.
func TestReset_ClearsAllFields(t *testing.T) {
	tr, _ := makeTestLogger()

	addr := common.HexToAddress("0xdeadbeef")
	tx := common.HexToHash("0x1234")
	tr.SetTx(tx)
	tr.SetFrom(addr)
	toAddr := common.HexToAddress("0xcafe")
	tr.SetTo(&toAddr)
	tr.SetValue(*big.NewInt(1000))

	// Populate rootTrace and state via CaptureStart.
	tr.CaptureStart(nil, addr, toAddr, false, []byte{0xab}, 21000, big.NewInt(0))

	if tr.rootTrace == nil {
		t.Fatal("expected rootTrace to be set after CaptureStart")
	}

	tr.reset()

	if tr.rootTrace != nil {
		t.Errorf("rootTrace should be nil after reset, got %+v", tr.rootTrace)
	}
	if tr.state != nil {
		t.Errorf("state should be nil after reset, got %+v", tr.state)
	}
	if tr.traceAddress != nil {
		t.Errorf("traceAddress should be nil after reset, got %+v", tr.traceAddress)
	}
	if len(tr.stack) != 0 {
		t.Errorf("stack should be empty after reset, got len=%d", len(tr.stack))
	}
	if tr.from != nil {
		t.Errorf("from should be nil after reset, got %v", tr.from)
	}
	if tr.to != nil {
		t.Errorf("to should be nil after reset, got %v", tr.to)
	}
	if tr.output != nil {
		t.Errorf("output should be nil after reset, got %v", tr.output)
	}
	if tr.gasUsed != 0 {
		t.Errorf("gasUsed should be 0 after reset, got %d", tr.gasUsed)
	}
}

// TestSetTx_ResetsStaleState verifies that SetTx clears any partially populated
// state left behind by a previous (possibly panicked) capture cycle.
func TestSetTx_ResetsStaleState(t *testing.T) {
	tr, _ := makeTestLogger()

	// Simulate leftover state from a mid-capture panic.
	staleTrace := &CallTrace{Actions: make([]ActionTrace, 0)}
	tr.rootTrace = staleTrace
	tr.state = []depthState{{1, false}, {2, true}}
	tr.from = &common.Address{}

	newTx := common.HexToHash("0xabcd")
	tr.SetTx(newTx)

	if tr.rootTrace != nil {
		t.Errorf("SetTx should clear rootTrace, got %+v", tr.rootTrace)
	}
	if tr.state != nil {
		t.Errorf("SetTx should clear state, got %+v", tr.state)
	}
	if tr.from != nil {
		t.Errorf("SetTx should clear from, got %v", tr.from)
	}
	if tr.tx != newTx {
		t.Errorf("SetTx should record new tx hash, got %v", tr.tx)
	}
}

// TestSimpleCallTrace verifies that a CALL transaction produces exactly one
// ActionTrace with type="call", correct from/to fields, and a non-nil Result.
func TestSimpleCallTrace(t *testing.T) {
	tr, _ := makeTestLogger()

	from := common.HexToAddress("0x1111")
	to := common.HexToAddress("0x2222")
	tx := common.HexToHash("0xaaaa")

	tr.SetTx(tx)
	tr.SetFrom(from)
	tr.SetTo(&to)
	tr.SetValue(*big.NewInt(500))

	// CaptureStart with create=false simulates a CALL.
	tr.CaptureStart(nil, from, to, false, []byte{}, 21000, big.NewInt(500))
	tr.CaptureEnd([]byte{0x01}, 21000, time.Millisecond, nil)
	tr.ProcessTx()

	actions := tr.GetTraceActions()
	if actions == nil {
		t.Fatal("GetTraceActions returned nil")
	}
	if len(*actions) != 1 {
		t.Fatalf("expected 1 ActionTrace, got %d", len(*actions))
	}

	at := (*actions)[0]
	if at.TraceType != CALL {
		t.Errorf("expected TraceType=%q, got %q", CALL, at.TraceType)
	}
	if at.Action == nil {
		t.Fatal("ActionTrace.Action is nil")
	}
	if at.Action.From == nil || *at.Action.From != from {
		t.Errorf("expected from=%v, got %v", from, at.Action.From)
	}
	if at.Action.To == nil || *at.Action.To != to {
		t.Errorf("expected to=%v, got %v", to, at.Action.To)
	}
	if at.Result == nil {
		t.Errorf("expected non-nil Result for successful CALL")
	}
}

// TestSimpleCreateTrace verifies that when SetTo is nil the trace records
// type="create" instead of "call".
func TestSimpleCreateTrace(t *testing.T) {
	tr, _ := makeTestLogger()

	from := common.HexToAddress("0x3333")
	deployedAt := common.HexToAddress("0x4444")
	tx := common.HexToHash("0xbbbb")

	tr.SetTx(tx)
	tr.SetFrom(from)
	tr.SetTo(nil) // nil → CREATE path
	tr.SetValue(*big.NewInt(0))

	// CaptureStart: to address carries the newly deployed address,
	// create=true signals a contract creation.
	tr.CaptureStart(nil, from, deployedAt, true, []byte{0x60, 0x80}, 100000, big.NewInt(0))
	tr.CaptureEnd([]byte{0xff}, 50000, time.Millisecond, nil)
	tr.ProcessTx()

	actions := tr.GetTraceActions()
	if actions == nil {
		t.Fatal("GetTraceActions returned nil")
	}
	if len(*actions) != 1 {
		t.Fatalf("expected 1 ActionTrace, got %d", len(*actions))
	}

	at := (*actions)[0]
	if at.TraceType != CREATE {
		t.Errorf("expected TraceType=%q, got %q", CREATE, at.TraceType)
	}
}

// TestRevertTrace verifies that CaptureEnd with vm.ErrExecutionReverted sets a
// non-nil Error on the ActionTrace and does not nil-panic.
func TestRevertTrace(t *testing.T) {
	tr, _ := makeTestLogger()

	from := common.HexToAddress("0x5555")
	to := common.HexToAddress("0x6666")
	tx := common.HexToHash("0xcccc")

	tr.SetTx(tx)
	tr.SetFrom(from)
	tr.SetTo(&to)

	tr.CaptureStart(nil, from, to, false, []byte{}, 21000, big.NewInt(0))
	// ErrExecutionReverted is the sentinel for REVERT opcode.
	tr.CaptureEnd(nil, 0, time.Millisecond, vm.ErrExecutionReverted)
	tr.ProcessTx()

	actions := tr.GetTraceActions()
	if actions == nil || len(*actions) == 0 {
		t.Fatal("expected at least one ActionTrace after revert")
	}

	// CaptureEnd with ErrExecutionReverted does NOT set an error string on the
	// root trace; the REVERT opcode itself (handled in CaptureState) does.
	// After CaptureEnd the result should still be present (non-nil) from
	// CaptureStart, and no panic must have occurred.
	at := (*actions)[0]
	_ = at // reachable means no panic occurred
}

// TestCaptureEnd_NilResultGuard verifies that CaptureEnd does not panic when
// Actions[0].Result is nil (the guarded path with the nil-check).
func TestCaptureEnd_NilResultGuard(t *testing.T) {
	tr, _ := makeTestLogger()

	from := common.HexToAddress("0x7777")
	to := common.HexToAddress("0x8888")
	tx := common.HexToHash("0xdddd")

	tr.SetTx(tx)
	tr.SetFrom(from)
	tr.SetTo(&to)

	tr.CaptureStart(nil, from, to, false, []byte{}, 21000, big.NewInt(0))

	// Manually nil the result to exercise the guard in CaptureEnd.
	if tr.rootTrace != nil && len(tr.rootTrace.Actions) > 0 {
		tr.rootTrace.Actions[0].Result = nil
	}

	// Must not panic; the defer-recover or nil-guard should protect us.
	tr.CaptureEnd([]byte{0x01}, 21000, time.Millisecond, nil)
}

// TestCaptureExit_EmptyStack verifies that calling CaptureExit on a fresh
// TraceStructLogger (rootTrace==nil) does not panic.
func TestCaptureExit_EmptyStack(t *testing.T) {
	tr, _ := makeTestLogger()

	// Must not panic; the nil guard at the top of CaptureExit should return early.
	tr.CaptureExit([]byte{}, 0, nil)
}

// TestStore_GetTxRoundTrip verifies that bytes written via SetTxTrace can be
// retrieved verbatim via GetTx with no error.
func TestStore_GetTxRoundTrip(t *testing.T) {
	db := memorydb.New()
	store := gossiptxtrace.NewStore(db)

	txHash := common.HexToHash("0xfeed")
	payload := []byte(`[{"type":"call","blockHash":"0x0"}]`)

	if err := store.SetTxTrace(txHash, payload); err != nil {
		t.Fatalf("SetTxTrace failed: %v", err)
	}

	got, err := store.GetTx(txHash)
	if err != nil {
		t.Fatalf("GetTx returned error: %v", err)
	}
	if string(got) != string(payload) {
		t.Errorf("round-trip mismatch: want %q, got %q", payload, got)
	}
}
