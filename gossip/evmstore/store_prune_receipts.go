package evmstore

import (
	"fmt"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

const receiptPruneCompactionThreshold = 100_000

// PruneReceiptsUpTo deletes receipts, tx positions, and internal tx data
// for all blocks up to and including n.
// Returns the total count of deleted entries.
func (s *Store) PruneReceiptsUpTo(n idx.Block) (int, error) {
	count := 0

	if err := s.pruneReceipts(n, &count); err != nil {
		return count, fmt.Errorf("pruning receipts: %w", err)
	}

	if err := s.pruneTxPositions(n, &count); err != nil {
		return count, fmt.Errorf("pruning tx positions: %w", err)
	}

	if count >= receiptPruneCompactionThreshold {
		if err := s.compactPrunedTables(); err != nil {
			log.Error("Failed to compact pruned tables", "err", err)
		}
	}

	return count, nil
}

func (s *Store) pruneReceipts(n idx.Block, count *int) error {
	batch := s.table.Receipts.NewBatch()
	it := s.table.Receipts.NewIterator(nil, nil)
	defer it.Release()

	start := time.Now()
	logged := start

	for it.Next() {
		key := it.Key()
		blockNum := idx.BytesToBlock(key)
		if blockNum > n {
			break
		}

		if err := batch.Delete(key); err != nil {
			return err
		}
		s.cache.Receipts.Remove(blockNum)
		*count++

		if time.Since(logged) > 8*time.Second {
			log.Info("Pruning receipts", "deleted", *count, "elapsed", common.PrettyDuration(time.Since(start)))
			logged = time.Now()
		}

		if batch.ValueSize() >= kvdb.IdealBatchSize {
			if err := batch.Write(); err != nil {
				return err
			}
			batch.Reset()
		}
	}
	if err := it.Error(); err != nil {
		return err
	}

	if batch.ValueSize() > 0 {
		return batch.Write()
	}
	return nil
}

func (s *Store) pruneTxPositions(n idx.Block, count *int) error {
	batch := s.table.TxPositions.NewBatch()
	txsBatch := s.table.Txs.NewBatch()

	it := s.table.TxPositions.NewIterator(nil, nil)
	defer it.Release()

	flushBatches := func() error {
		if batch.ValueSize() > 0 {
			if err := batch.Write(); err != nil {
				return err
			}
			batch.Reset()
		}
		if txsBatch.ValueSize() > 0 {
			if err := txsBatch.Write(); err != nil {
				return err
			}
			txsBatch.Reset()
		}
		return nil
	}

	start := time.Now()
	logged := start

	for it.Next() {
		var pos TxPosition
		if err := rlp.DecodeBytes(it.Value(), &pos); err != nil {
			return fmt.Errorf("decoding TxPosition: %w", err)
		}
		if pos.Block > n {
			continue
		}

		key := it.Key()
		txid := common.BytesToHash(key)

		if err := batch.Delete(key); err != nil {
			return err
		}
		if err := txsBatch.Delete(key); err != nil {
			return err
		}
		s.cache.TxPositions.Remove(txid.String())
		*count++

		if time.Since(logged) > 8*time.Second {
			log.Info("Pruning tx positions", "deleted", *count, "elapsed", common.PrettyDuration(time.Since(start)))
			logged = time.Now()
		}

		if batch.ValueSize()+txsBatch.ValueSize() >= kvdb.IdealBatchSize {
			if err := flushBatches(); err != nil {
				return err
			}
		}
	}
	if err := it.Error(); err != nil {
		return err
	}

	return flushBatches()
}

func (s *Store) compactPrunedTables() error {
	start := time.Now()
	for _, db := range []interface{ Compact([]byte, []byte) error }{s.table.Receipts, s.table.TxPositions, s.table.Txs} {
		for b := 0x00; b <= 0xf0; b += 0x10 {
			lo := []byte{byte(b)}
			hi := []byte{byte(b + 0x10)}
			if b == 0xf0 {
				hi = nil
			}
			if err := db.Compact(lo, hi); err != nil {
				return fmt.Errorf("compaction failed: %w", err)
			}
		}
	}
	log.Info("Compaction finished", "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}
