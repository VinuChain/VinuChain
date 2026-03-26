package concurrent

import (
	"sync"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type ValidatorEventsSet struct {
	sync.RWMutex
	// Val is the underlying map. Callers MUST hold Lock()/RLock() before accessing.
	// Use Get() for safe single-key reads, Set() for safe single-key writes.
	Val map[idx.ValidatorID]hash.Event
}

func WrapValidatorEventsSet(v map[idx.ValidatorID]hash.Event) *ValidatorEventsSet {
	return &ValidatorEventsSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}

func (s *ValidatorEventsSet) Get(id idx.ValidatorID) (hash.Event, bool) {
	s.RLock()
	defer s.RUnlock()
	v, ok := s.Val[id]
	return v, ok
}

func (s *ValidatorEventsSet) Set(id idx.ValidatorID, val hash.Event) {
	s.Lock()
	defer s.Unlock()
	s.Val[id] = val
}
