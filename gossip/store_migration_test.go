package gossip

import (
	"errors"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/stretchr/testify/require"
)

// errIterator is a kvdb.Iteratee whose first Next() call returns false and
// sets a non-nil Error(), simulating a LevelDB read failure.
type errIterator struct {
	err error
}

func (e *errIterator) NewIterator(prefix []byte, start []byte) kvdb.Iterator {
	return &errIter{err: e.err}
}

type errIter struct {
	err error
}

func (i *errIter) Next() bool    { return false }
func (i *errIter) Error() error  { return i.err }
func (i *errIter) Key() []byte   { return nil }
func (i *errIter) Value() []byte { return nil }
func (i *errIter) Release()      {}

// Verify errIter satisfies the ethdb.Iterator interface.
var _ ethdb.Iterator = (*errIter)(nil)

// TestIsEmptyDBPropagatesIteratorError verifies that isEmptyDB returns an error
// when the underlying iterator reports a DB read failure rather than silently
// treating the failure as an empty table.
func TestIsEmptyDBPropagatesIteratorError(t *testing.T) {
	dbErr := errors.New("simulated LevelDB read failure")
	db := &errIterator{err: dbErr}

	empty, err := isEmptyDB(db)
	require.Error(t, err, "isEmptyDB must propagate iterator errors")
	require.True(t, empty, "empty should be true when error occurs (no valid keys)")
	require.ErrorIs(t, err, dbErr)
}

// TestIsEmptyDBTrueForEmpty verifies that isEmptyDB returns true and no error
// for a genuinely empty table.
func TestIsEmptyDBTrueForEmpty(t *testing.T) {
	db := &errIterator{err: nil}

	empty, err := isEmptyDB(db)
	require.NoError(t, err)
	require.True(t, empty)
}
