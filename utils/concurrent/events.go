package concurrent

import (
	"sync"

	"github.com/Fantom-foundation/lachesis-base/hash"
)

type EventsSet struct {
	sync.RWMutex
	// Val is the underlying map. Callers MUST hold Lock()/RLock() before accessing.
	// Use Contains() for safe single-key reads, Add()/Remove() for safe single-key writes.
	Val hash.EventsSet
}

func WrapEventsSet(v hash.EventsSet) *EventsSet {
	return &EventsSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}

func (s *EventsSet) Contains(id hash.Event) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.Val[id]
	return ok
}

func (s *EventsSet) Add(id hash.Event) {
	s.Lock()
	defer s.Unlock()
	s.Val[id] = struct{}{}
}

func (s *EventsSet) Remove(id hash.Event) {
	s.Lock()
	defer s.Unlock()
	delete(s.Val, id)
}
