package concurrent

import (
	"sync"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type ValidatorEpochsSet struct {
	sync.RWMutex
	// Val is the underlying map. Callers MUST hold Lock()/RLock() before accessing.
	// Use Get() for safe single-key reads, Set() for safe single-key writes.
	Val map[idx.ValidatorID]idx.Epoch
}

func WrapValidatorEpochsSet(v map[idx.ValidatorID]idx.Epoch) *ValidatorEpochsSet {
	return &ValidatorEpochsSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}

func (s *ValidatorEpochsSet) Get(id idx.ValidatorID) (idx.Epoch, bool) {
	s.RLock()
	defer s.RUnlock()
	v, ok := s.Val[id]
	return v, ok
}

func (s *ValidatorEpochsSet) Set(id idx.ValidatorID, val idx.Epoch) {
	s.Lock()
	defer s.Unlock()
	s.Val[id] = val
}
