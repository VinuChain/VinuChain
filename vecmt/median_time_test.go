package vecmt

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/vecfc"
	"github.com/stretchr/testify/assert"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/dag/tdag"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/kvdb/memorydb"

	"github.com/Fantom-foundation/go-opera/inter"
)

func TestMedianTimeOnIndex(t *testing.T) {
	nodes := tdag.GenNodes(5)
	weights := []pos.Weight{5, 4, 3, 2, 1}
	validators := pos.ArrayToValidators(nodes, weights)

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), nil)

	assertar := assert.New(t)
	{ // seq=0
		e := hash.ZeroEvent
		// validator indexes are sorted by weight amount
		before := NewHighestBefore(idx.Validator(validators.Len()))

		before.VSeq.Set(0, vecfc.BranchSeq{Seq: 0})
		before.VTime.Set(0, 100)

		before.VSeq.Set(1, vecfc.BranchSeq{Seq: 0})
		before.VTime.Set(1, 100)

		before.VSeq.Set(2, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(2, 10)

		before.VSeq.Set(3, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(3, 10)

		before.VSeq.Set(4, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(4, 10)

		vi.SetHighestBefore(e, before)
		assertar.Equal(inter.Timestamp(1), vi.MedianTime(e, 1))
	}

	{ // fork seen = true
		e := hash.ZeroEvent
		// validator indexes are sorted by weight amount
		before := NewHighestBefore(idx.Validator(validators.Len()))

		before.SetForkDetected(0)
		before.VTime.Set(0, 100)

		before.SetForkDetected(1)
		before.VTime.Set(1, 100)

		before.VSeq.Set(2, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(2, 10)

		before.VSeq.Set(3, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(3, 10)

		before.VSeq.Set(4, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(4, 10)

		vi.SetHighestBefore(e, before)
		assertar.Equal(inter.Timestamp(10), vi.MedianTime(e, 1))
	}

	{ // normal
		e := hash.ZeroEvent
		// validator indexes are sorted by weight amount
		before := NewHighestBefore(idx.Validator(validators.Len()))

		before.VSeq.Set(0, vecfc.BranchSeq{Seq: 1})
		before.VTime.Set(0, 11)

		before.VSeq.Set(1, vecfc.BranchSeq{Seq: 2})
		before.VTime.Set(1, 12)

		before.VSeq.Set(2, vecfc.BranchSeq{Seq: 2})
		before.VTime.Set(2, 13)

		before.VSeq.Set(3, vecfc.BranchSeq{Seq: 3})
		before.VTime.Set(3, 14)

		before.VSeq.Set(4, vecfc.BranchSeq{Seq: 4})
		before.VTime.Set(4, 15)

		vi.SetHighestBefore(e, before)
		assertar.Equal(inter.Timestamp(12), vi.MedianTime(e, 1))
	}

}

// TestMedianTimeAllCheaters verifies that when every validator is detected as a
// fork cheater (honestTotalWeight == 0), MedianTime returns the caller-supplied
// defaultTime rather than panicking or calling crit. This is the catastrophic
// all-cheaters guard introduced in the elemont audit.
func TestMedianTimeAllCheaters(t *testing.T) {
	nodes := tdag.GenNodes(3)
	weights := []pos.Weight{3, 2, 1}
	validators := pos.ArrayToValidators(nodes, weights)

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), nil)

	e := hash.ZeroEvent
	before := NewHighestBefore(idx.Validator(validators.Len()))

	// Mark all validators as fork cheaters — their weights must be zeroed out.
	for i := idx.Validator(0); i < idx.Validator(validators.Len()); i++ {
		before.SetForkDetected(i)
		before.VTime.Set(i, inter.Timestamp(1000+uint64(i)))
	}
	vi.SetHighestBefore(e, before)

	const defaultTime = inter.Timestamp(42)
	result := vi.MedianTime(e, defaultTime)
	if result != defaultTime {
		t.Errorf("all-cheaters case: got %d, want defaultTime=%d", result, defaultTime)
	}
}

// TestMedianTimeElemont_StableSort verifies that enabling the Elemont upgrade
// causes MedianTime to use a stable sort with a weight tie-breaker. With equal
// timestamps across all validators, both modes must return the same timestamp
// (since the median value is independent of sort order when all times are equal),
// but the test confirms no panic or crit is triggered and the result is correct.
func TestMedianTimeElemont_StableSort(t *testing.T) {
	nodes := tdag.GenNodes(4)
	// Assign unequal weights so weight tie-breaking could matter.
	weights := []pos.Weight{5, 3, 2, 1}
	validators := pos.ArrayToValidators(nodes, weights)

	for _, elemont := range []bool{false, true} {
		vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
		vi.Reset(validators, memorydb.New(), nil)
		vi.SetElemont(elemont)

		e := hash.ZeroEvent
		before := NewHighestBefore(idx.Validator(validators.Len()))
		before.elemont = elemont

		// All validators report the same creation time. The stable-sort tie-breaker
		// by weight must not change the final timestamp output.
		for i := idx.Validator(0); i < idx.Validator(validators.Len()); i++ {
			before.VSeq.Set(i, vecfc.BranchSeq{Seq: 1})
			before.VTime.Set(i, inter.Timestamp(100))
		}
		vi.SetHighestBefore(e, before)

		result := vi.MedianTime(e, inter.Timestamp(1))
		if result != inter.Timestamp(100) {
			t.Errorf("elemont=%v equal-timestamps case: got %d, want 100", elemont, result)
		}
	}
}

func TestMedianTimeOnDAG(t *testing.T) {
	dagAscii := `
 ║
 nodeA001
 ║
 nodeA012
 ║            ║
 ║            nodeB001
 ║            ║            ║
 ║            ╠═══════════ nodeC001
 ║║           ║            ║            ║
 ║╚══════════─╫─══════════─╫─══════════ nodeD001
║║            ║            ║            ║
╚ nodeA002════╬════════════╬════════════╣
 ║║           ║            ║            ║
 ║╚══════════─╫─══════════─╫─══════════ nodeD002
 ║            ║            ║            ║
 nodeA003════─╫─══════════─╫─═══════════╣
 ║            ║            ║
 ╠════════════nodeB002     ║
 ║            ║            ║
 ╠════════════╫═══════════ nodeC002
`

	weights := []pos.Weight{3, 4, 2, 1}
	genesisTime := inter.Timestamp(1)
	creationTimes := map[string]inter.Timestamp{
		"nodeA001": inter.Timestamp(111),
		"nodeB001": inter.Timestamp(112),
		"nodeC001": inter.Timestamp(13),
		"nodeD001": inter.Timestamp(14),
		"nodeA002": inter.Timestamp(120),
		"nodeD002": inter.Timestamp(20),
		"nodeA012": inter.Timestamp(120),
		"nodeA003": inter.Timestamp(20),
		"nodeB002": inter.Timestamp(20),
		"nodeC002": inter.Timestamp(35),
	}
	medianTimes := map[string]inter.Timestamp{
		"nodeA001": genesisTime,
		"nodeB001": genesisTime,
		"nodeC001": inter.Timestamp(13),
		"nodeD001": genesisTime,
		"nodeA002": inter.Timestamp(112),
		"nodeD002": genesisTime,
		"nodeA012": genesisTime,
		"nodeA003": inter.Timestamp(20),
		"nodeB002": inter.Timestamp(20),
		"nodeC002": inter.Timestamp(35),
	}
	t.Run("testMedianTimeOnDAG", func(t *testing.T) {
		testMedianTime(t, dagAscii, weights, creationTimes, medianTimes, genesisTime)
	})
}

func testMedianTime(t *testing.T, dagAscii string, weights []pos.Weight, creationTimes map[string]inter.Timestamp, medianTimes map[string]inter.Timestamp, genesis inter.Timestamp) {
	assertar := assert.New(t)

	var ordered dag.Events
	nodes, _, named := tdag.ASCIIschemeForEach(dagAscii, tdag.ForEachEvent{
		Process: func(e dag.Event, name string) {
			ordered = append(ordered, &eventWithCreationTime{e, creationTimes[name]})
		},
	})

	validators := pos.ArrayToValidators(nodes, weights)

	events := make(map[hash.Event]dag.Event)
	getEvent := func(id hash.Event) dag.Event {
		return events[id]
	}

	vi := NewIndex(func(err error) { panic(err) }, LiteConfig())
	vi.Reset(validators, memorydb.New(), getEvent)

	// push
	for _, e := range ordered {
		events[e.ID()] = e
		assertar.NoError(vi.Add(e))
		vi.Flush()
	}

	// check
	for name, e := range named {
		expected, ok := medianTimes[name]
		if !ok {
			continue
		}
		assertar.Equal(expected, vi.MedianTime(e.ID(), genesis), name)
	}
}
