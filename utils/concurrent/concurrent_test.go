package concurrent

import (
	"sync"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

// ValidatorBlocksSet

func TestValidatorBlocksSet_GetSet(t *testing.T) {
	s := WrapValidatorBlocksSet(map[idx.ValidatorID]idx.Block{})
	if _, ok := s.Get(1); ok {
		t.Fatal("expected miss on empty set")
	}
	s.Set(1, 42)
	v, ok := s.Get(1)
	if !ok || v != 42 {
		t.Fatalf("got (%v, %v), want (42, true)", v, ok)
	}
}

func TestValidatorBlocksSet_Concurrent(t *testing.T) {
	s := WrapValidatorBlocksSet(map[idx.ValidatorID]idx.Block{})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		vid := idx.ValidatorID(i % 5)
		blk := idx.Block(i)
		go func() { defer wg.Done(); s.Set(vid, blk) }()
		go func() { defer wg.Done(); s.Get(vid) }()
	}
	wg.Wait()
}

// ValidatorEventsSet

func TestValidatorEventsSet_GetSet(t *testing.T) {
	s := WrapValidatorEventsSet(map[idx.ValidatorID]hash.Event{})
	if _, ok := s.Get(1); ok {
		t.Fatal("expected miss on empty set")
	}
	ev := hash.Event{0x01}
	s.Set(1, ev)
	v, ok := s.Get(1)
	if !ok || v != ev {
		t.Fatalf("got (%v, %v), want (%v, true)", v, ok, ev)
	}
}

func TestValidatorEventsSet_Concurrent(t *testing.T) {
	s := WrapValidatorEventsSet(map[idx.ValidatorID]hash.Event{})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		vid := idx.ValidatorID(i % 5)
		ev := hash.Event{byte(i)}
		go func() { defer wg.Done(); s.Set(vid, ev) }()
		go func() { defer wg.Done(); s.Get(vid) }()
	}
	wg.Wait()
}

// ValidatorEpochsSet

func TestValidatorEpochsSet_GetSet(t *testing.T) {
	s := WrapValidatorEpochsSet(map[idx.ValidatorID]idx.Epoch{})
	if _, ok := s.Get(1); ok {
		t.Fatal("expected miss on empty set")
	}
	s.Set(1, 7)
	v, ok := s.Get(1)
	if !ok || v != 7 {
		t.Fatalf("got (%v, %v), want (7, true)", v, ok)
	}
}

func TestValidatorEpochsSet_Concurrent(t *testing.T) {
	s := WrapValidatorEpochsSet(map[idx.ValidatorID]idx.Epoch{})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		vid := idx.ValidatorID(i % 5)
		epoch := idx.Epoch(i)
		go func() { defer wg.Done(); s.Set(vid, epoch) }()
		go func() { defer wg.Done(); s.Get(vid) }()
	}
	wg.Wait()
}

// EventsSet

func TestEventsSet_ContainsAddRemove(t *testing.T) {
	s := WrapEventsSet(hash.EventsSet{})
	ev := hash.Event{0xAB}
	if s.Contains(ev) {
		t.Fatal("expected miss on empty set")
	}
	s.Add(ev)
	if !s.Contains(ev) {
		t.Fatal("expected hit after Add")
	}
	s.Remove(ev)
	if s.Contains(ev) {
		t.Fatal("expected miss after Remove")
	}
}

func TestEventsSet_Concurrent(t *testing.T) {
	s := WrapEventsSet(hash.EventsSet{})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(3)
		ev := hash.Event{byte(i)}
		go func() { defer wg.Done(); s.Add(ev) }()
		go func() { defer wg.Done(); s.Contains(ev) }()
		go func() { defer wg.Done(); s.Remove(ev) }()
	}
	wg.Wait()
}
