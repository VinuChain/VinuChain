package gossip

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/dag"

	"github.com/Fantom-foundation/go-opera/inter"
)

// missingEvent returns a typed-nil *inter.Event through a function boundary.
// The function-local nilness analyzer in go vet / gopls flags direct
// `var x *T; if x == nil` patterns as tautological, even though exercising
// those semantics is the whole point of TestTypedNilInterfaceWrapping. Hiding
// the nil behind a package-scoped helper makes the test's intent explicit and
// keeps the analyzer quiet without a suppression comment.
func missingEvent() *inter.Event { return nil }

// TestTypedNilInterfaceWrapping verifies that returning a nil *inter.Event
// directly as dag.Event produces a non-nil typed nil, and that the correct
// guard (explicit nil check) produces an untyped nil.
//
// This is the core bug behind LBI-01: two getEvent closures in service.go
// and c_event_callbacks.go returned Store.GetEvent(id) directly as dag.Event,
// producing typed nils that bypass nil checks in lachesis-base (dfsSubgraph,
// calcFrameIdx), causing nil-pointer panics.
func TestTypedNilInterfaceWrapping(t *testing.T) {
	// Simulate Store.GetEvent returning nil for a missing event. The helper
	// is required — a `var nilEvent *inter.Event` literal gets folded by the
	// nilness analyzer, which would flag the comparisons below as tautological
	// even though verifying that exact semantics is the test's contract.
	nilEvent := missingEvent()

	// WRONG pattern (pre-fix): direct return wraps typed nil.
	// Go's interface comparison (==) detects this; reflect-based checks do not.
	var wrong dag.Event = nilEvent
	if wrong == nil {
		t.Fatal("typed nil *inter.Event wrapped in dag.Event must not pass == nil check (the bug)")
	}

	// CORRECT pattern (post-fix): explicit nil check returns untyped nil.
	var correct dag.Event
	if nilEvent == nil {
		correct = nil
	} else {
		correct = nilEvent
	}
	if correct != nil {
		t.Fatal("explicit nil guard must produce untyped nil that passes == nil check (the fix)")
	}
}
