package concurrent

import (
	"sync"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type ValidatorBlocksSet struct {
	sync.RWMutex
	// Val is the underlying map. Callers MUST hold Lock()/RLock() before accessing.
	// Use Get() for safe single-key reads, Set() for safe single-key writes.
	Val map[idx.ValidatorID]idx.Block
}

func WrapValidatorBlocksSet(v map[idx.ValidatorID]idx.Block) *ValidatorBlocksSet {
	return &ValidatorBlocksSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}

func (s *ValidatorBlocksSet) Get(id idx.ValidatorID) (idx.Block, bool) {
	s.RLock()
	defer s.RUnlock()
	v, ok := s.Val[id]
	return v, ok
}

func (s *ValidatorBlocksSet) Set(id idx.ValidatorID, val idx.Block) {
	s.Lock()
	defer s.Unlock()
	s.Val[id] = val
}
