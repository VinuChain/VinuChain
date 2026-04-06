package txtrace

import (
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Fantom-foundation/go-opera/logger"
)

// Store is a transaction traces persistent storage working over physical key-value database.
type Store struct {
	mainDB kvdb.Store
	logger.Instance
}

// NewStore creates store over key-value db.
func NewStore(mainDB kvdb.Store) *Store {
	s := &Store{
		mainDB:   mainDB,
		Instance: logger.New("TxTrace Store"),
	}
	return s
}

// Close closes underlying database.
func (s *Store) Close() {
	_ = s.mainDB.Close()
}

// SetTxTrace stores []byte representation of transaction traces.
func (s *Store) SetTxTrace(txID common.Hash, txTraces []byte) error {
	return s.mainDB.Put(txID.Bytes(), txTraces)
}

// GetTx returns stored transaction traces, or nil if not found.
// Returns an error only for genuine I/O failures; a missing key is not an error.
func (s *Store) GetTx(txID common.Hash) ([]byte, error) {
	buf, err := s.mainDB.Get(txID.Bytes())
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// RemoveTxTrace removes key and []byte representation of transaction traces.
func (s *Store) RemoveTxTrace(txID common.Hash) error {
	return s.mainDB.Delete(txID.Bytes())
}

// ForEachTxtrace iterates over all transaction traces in db.
func (s *Store) ForEachTxtrace(onEvent func(key common.Hash, traces []byte) bool) {
	it := s.mainDB.NewIterator(nil, nil)
	defer it.Release()
	for it.Next() {
		if !onEvent(common.BytesToHash(it.Key()), it.Value()) {
			return
		}
	}
}
