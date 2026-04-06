package topicsdb

import (
	"context"
	"errors"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/Fantom-foundation/lachesis-base/kvdb/memorydb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/logger"
)

// errIterator wraps a real iterator but injects an error after returning
// the first key/value pair.
type errIterator struct {
	inner kvdb.Iterator
	count int
	err   error
}

func (e *errIterator) Next() bool {
	if e.count > 0 {
		return false
	}
	ok := e.inner.Next()
	if ok {
		e.count++
	}
	return ok
}

func (e *errIterator) Error() error {
	if e.count > 0 {
		return e.err
	}
	return e.inner.Error()
}

func (e *errIterator) Key() []byte   { return e.inner.Key() }
func (e *errIterator) Value() []byte { return e.inner.Value() }
func (e *errIterator) Release()      { e.inner.Release() }

// errInjectStore wraps a real kvdb.Store but replaces NewIterator with one
// that injects an error after the first result.
type errInjectStore struct {
	kvdb.Store
	injectedErr error
}

func (s *errInjectStore) NewIterator(prefix []byte, start []byte) kvdb.Iterator {
	return &errIterator{
		inner: s.Store.NewIterator(prefix, start),
		err:   s.injectedErr,
	}
}

// TestSearchParallelIteratorError verifies that a LevelDB iterator error
// during topic scanning is propagated to the caller of FindInBlocks.
// Currently this test FAILS because searchParallel silently discards the error.
func TestSearchParallelIteratorError(t *testing.T) {
	logger.SetTestMode(t)

	addr := randAddress()
	topic := common.BytesToHash([]byte("topic1"))

	index, err := New(memorydb.NewProducer(""))
	require.NoError(t, err)

	// Push enough logs so the iterator has at least one entry to return
	// before injecting the error.
	for i := 0; i < 3; i++ {
		err := index.Push(&types.Log{
			BlockNumber: uint64(i + 1),
			Address:     addr,
			Topics:      []common.Hash{topic},
		})
		require.NoError(t, err)
	}

	// Inject an error store on the Topic table.
	injectedErr := errors.New("simulated LevelDB read error")
	index.table.Topic = &errInjectStore{
		Store:       index.table.Topic,
		injectedErr: injectedErr,
	}

	_, searchErr := index.FindInBlocks(context.Background(), idx.Block(1), idx.Block(10), [][]common.Hash{
		{addr.Hash()},
		{topic},
	})
	require.ErrorIs(t, searchErr, injectedErr, "iterator error must be propagated from FindInBlocks")
}
