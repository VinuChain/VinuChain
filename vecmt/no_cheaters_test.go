package vecmt

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/dag/tdag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/kvdb/memorydb"

	"github.com/Fantom-foundation/go-opera/inter"
)

// makeTestEvent builds a tdag.TestEvent with the given fields and a unique ID.
// idSuffix is placed in the last byte so events with the same epoch/lamport remain distinct.
// If seq > 1, the first element of parents is used as the self-parent by SelfParent().
func makeTestEvent(creator idx.ValidatorID, seq idx.Event, lamport idx.Lamport, idSuffix uint8, parents ...hash.Event) *tdag.TestEvent {
	e := &tdag.TestEvent{}
	e.SetCreator(creator)
	e.SetSeq(seq)
	e.SetLamport(lamport)
	e.SetParents(hash.Events{})
	for _, p := range parents {
		e.AddParent(p)
	}
	var id [24]byte
	id[23] = idSuffix
	e.SetID(id)
	return e
}

// addToIndex wraps e with a CreationTime of inter.Timestamp(seq), stores it in events,
// adds it to vi, and flushes. Panics on Add error (consistent with other vecmt tests).
func addToIndex(t *testing.T, vi *Index, events map[hash.Event]dag.Event, e *tdag.TestEvent) {
	t.Helper()
	wrapped := &eventWithCreationTime{e, inter.Timestamp(e.Seq())}
	events[e.ID()] = wrapped
	if err := vi.Add(wrapped); err != nil {
		t.Fatalf("vi.Add: %v", err)
	}
	vi.Flush()
}

// TestNoCheaters_NilSelfParent verifies that a nil selfParent causes NoCheaters to return
// the options slice unchanged (fast path, no DAG traversal).
func TestNoCheaters_NilSelfParent(t *testing.T) {
	nodes := tdag.GenNodes(2)
	validators := pos.ArrayToValidators(nodes, []pos.Weight{1, 1})

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), nil)

	options := hash.Events{hash.ZeroEvent}
	result := vi.NoCheaters(nil, options)
	if len(result) != len(options) {
		t.Errorf("nil selfParent: got %d options, want %d", len(result), len(options))
	}
}

// TestNoCheaters_NoFork_ReturnsAll verifies that when no fork exists in the DAG,
// NoCheaters returns all candidate options unfiltered.
func TestNoCheaters_NoFork_ReturnsAll(t *testing.T) {
	nodes := tdag.GenNodes(2)
	nodeA, nodeB := nodes[0], nodes[1]
	validators := pos.ArrayToValidators(nodes, []pos.Weight{1, 1})

	events := make(map[hash.Event]dag.Event)
	getEvent := func(id hash.Event) dag.Event { return events[id] }

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), getEvent)

	a0 := makeTestEvent(nodeA, 1, 1, 0)
	b0 := makeTestEvent(nodeB, 1, 1, 1)
	// b1 is the self-parent; it references a0 as a cross-parent (no fork from A)
	b1 := makeTestEvent(nodeB, 2, 2, 2, b0.ID(), a0.ID())
	for _, e := range []*tdag.TestEvent{a0, b0, b1} {
		addToIndex(t, vi, events, e)
	}

	selfParent := b1.ID()
	options := hash.Events{a0.ID(), b0.ID()}
	result := vi.NoCheaters(&selfParent, options)
	if len(result) != 2 {
		t.Errorf("no-fork: got %d results, want 2", len(result))
	}
}

// buildForkIndex constructs a 2-validator index where nodeA creates a fork visible to b1.
// The setup:
//
//	a0 (A, seq=1)   b0 (B, seq=1)   a0fork (A, seq=1, new branch — fork)
//	                b1 (B, seq=2, parents=[b0, a0, a0fork]) ← sees both A branches → A is fork-detected
//
// Returns the index and the ID of b1 (used as selfParent in NoCheaters calls).
func buildForkIndex(t *testing.T, elemont bool) (*Index, *tdag.TestEvent, *tdag.TestEvent, hash.Event) {
	t.Helper()
	nodes := tdag.GenNodes(2)
	nodeA, nodeB := nodes[0], nodes[1]
	validators := pos.ArrayToValidators(nodes, []pos.Weight{1, 1})

	events := make(map[hash.Event]dag.Event)
	getEvent := func(id hash.Event) dag.Event { return events[id] }

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), getEvent)
	vi.SetElemont(elemont)

	a0 := makeTestEvent(nodeA, 1, 1, 0)
	b0 := makeTestEvent(nodeB, 1, 1, 1)
	// a0fork: same creator as a0, seq=1, no self-parent → engine allocates a new branch (fork)
	a0fork := makeTestEvent(nodeA, 1, 1, 2)
	// b1 has both A branches as parents, so its vectors see A as a fork cheater
	b1 := makeTestEvent(nodeB, 2, 3, 3, b0.ID(), a0.ID(), a0fork.ID())
	for _, e := range []*tdag.TestEvent{a0, b0, a0fork, b1} {
		addToIndex(t, vi, events, e)
	}

	if !vi.Engine.AtLeastOneFork() {
		t.Fatal("expected fork to be detected after adding conflicting events from same creator")
	}
	return vi, a0, a0fork, b1.ID()
}

// TestNoCheaters_ForkDetected_Elemont verifies that in elemont mode, NoCheaters uses
// the merged (validator-indexed) view and filters all events from the forking validator.
func TestNoCheaters_ForkDetected_Elemont(t *testing.T) {
	vi, a0, a0fork, b1ID := buildForkIndex(t, true)

	options := hash.Events{a0.ID(), a0fork.ID()}
	result := vi.NoCheaters(&b1ID, options)
	if len(result) != 0 {
		t.Errorf("elemont fork-filter: expected 0 results (cheater filtered), got %d", len(result))
	}
}

// TestNoCheaters_ForkDetected_PreElemont verifies that in pre-elemont mode, NoCheaters uses
// the raw branch-indexed view and also filters all events from the forking validator.
func TestNoCheaters_ForkDetected_PreElemont(t *testing.T) {
	vi, a0, a0fork, b1ID := buildForkIndex(t, false)

	options := hash.Events{a0.ID(), a0fork.ID()}
	result := vi.NoCheaters(&b1ID, options)
	if len(result) != 0 {
		t.Errorf("pre-elemont fork-filter: expected 0 results (cheater filtered), got %d", len(result))
	}
}
