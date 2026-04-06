package topicsdb

import (
	"math"
	"sync"
	"sync/atomic"
)

type (
	posCounter struct {
		pos   int
		count int
	}

	blockCounter struct {
		wait  chan struct{}
		count int
	}

	synchronizator struct {
		mu        sync.Mutex
		threads   sync.WaitGroup
		positions map[int]*posCounter
		goNext    int32
		minBlock  uint64
		blocks    map[uint64]*blockCounter
		firstErr  error
	}
)

func newSynchronizator() *synchronizator {
	s := &synchronizator{
		positions: make(map[int]*posCounter),
		minBlock:  0,
		blocks:    make(map[uint64]*blockCounter),
	}
	atomic.StoreInt32(&s.goNext, 1)

	return s
}

func (s *synchronizator) Halt() {
	atomic.StoreInt32(&s.goNext, 0)

	s.mu.Lock()
	defer s.mu.Unlock()

	for n := range s.blocks {
		if n != s.minBlock {
			close(s.blocks[n].wait)
		}
	}
}

func (s *synchronizator) GoNext(n uint64) (prev uint64, gonext bool) {
	if atomic.LoadInt32(&s.goNext) == 0 {
		return
	}

	if n > s.minBlock {
		s.mu.Lock()
		prev = s.minBlock
		s.enqueueBlock(n)
		s.dequeueBlock()
		wait := s.blocks[n].wait
		s.mu.Unlock()
		// wait for other threads
		<-wait
	}

	gonext = atomic.LoadInt32(&s.goNext) != 0
	return
}

func (s *synchronizator) StartThread(pos int, num int) {
	s.threads.Add(1)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.enqueueBlock(s.minBlock)

	if _, ok := s.positions[pos]; ok {
		s.positions[pos].count++
	} else {
		s.positions[pos] = &posCounter{pos, 1}
	}
}

func (s *synchronizator) FinishThread(pos int, num int) {
	s.threads.Done()

	s.mu.Lock()
	defer s.mu.Unlock()

	s.positions[pos].count--
	s.dequeueBlock()
}

func (s *synchronizator) enqueueBlock(n uint64) {
	if _, ok := s.blocks[n]; ok {
		s.blocks[n].count++
	} else {
		s.blocks[n] = &blockCounter{
			wait:  make(chan struct{}),
			count: 1,
		}
	}
}

func (s *synchronizator) dequeueBlock() {
	s.blocks[s.minBlock].count--
	if s.blocks[s.minBlock].count < 1 {

		for _, pos := range s.positions {
			if pos.count < 1 {
				atomic.StoreInt32(&s.goNext, 0)
				break
			}
		}

		delete(s.blocks, s.minBlock)
		if len(s.blocks) < 1 {
			return
		}
		// find new minBlock
		s.minBlock = math.MaxUint64
		for b := range s.blocks {
			if s.minBlock > b {
				s.minBlock = b
			}
		}
		close(s.blocks[s.minBlock].wait)
	}
}

func (s *synchronizator) WaitForThreads() {
	s.threads.Wait()
}

func (s *synchronizator) PositionsCount() int {
	return len(s.positions)
}

// SetError records the first iterator error encountered by any scan goroutine.
func (s *synchronizator) SetError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.firstErr == nil {
		s.firstErr = err
	}
}

// Err returns the first iterator error recorded by SetError, if any.
func (s *synchronizator) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.firstErr
}
