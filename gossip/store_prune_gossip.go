package gossip

import (
	"fmt"
	"time"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const pruneCompactionThreshold = 100_000

// PruneEpochData deletes all gossip data for epochs up to and including pruneUpTo.
// Returns the number of deleted entries and any error.
func (s *Store) PruneEpochData(pruneUpTo idx.Epoch) (int, error) {
	start := time.Now()
	total := 0
	var firstBlockOfEpoch idx.Block

	for epoch := idx.Epoch(1); epoch <= pruneUpTo; epoch++ {
		n, nextFirst, err := s.pruneOneEpoch(epoch, firstBlockOfEpoch)
		if err != nil {
			return total, fmt.Errorf("pruning epoch %d: %w", epoch, err)
		}
		total += n
		firstBlockOfEpoch = nextFirst
	}

	if total >= pruneCompactionThreshold {
		if err := s.runGossipCompaction(); err != nil {
			return total, err
		}
	}

	log.Info("Gossip epoch pruning complete",
		"pruneUpTo", pruneUpTo,
		"deleted", total,
		"elapsed", common.PrettyDuration(time.Since(start)),
	)
	return total, nil
}

// pruneTableByPrefix deletes all entries in tbl with the given prefix using batch writes.
// Returns the number of deleted entries.
func pruneTableByPrefix(tbl kvdb.Store, prefix []byte) (int, error) {
	count := 0
	batch := tbl.NewBatch()

	flush := func() error {
		if batch.ValueSize() == 0 {
			return nil
		}
		if err := batch.Write(); err != nil {
			return err
		}
		batch.Reset()
		return nil
	}

	it := tbl.NewIterator(prefix, nil)
	for it.Next() {
		if err := batch.Delete(it.Key()); err != nil {
			it.Release()
			_ = flush()
			return count, err
		}
		count++
		if batch.ValueSize() >= kvdb.IdealBatchSize {
			if err := flush(); err != nil {
				it.Release()
				return count, err
			}
		}
	}
	if err := it.Error(); err != nil {
		it.Release()
		return count, err
	}
	it.Release()

	if err := flush(); err != nil {
		return count, err
	}
	return count, nil
}

// pruneOneEpoch removes all stored data for a single epoch.
// firstBlock is the first block of this epoch (0 means unknown/skip block range).
// Returns the count of deleted entries and the first block of the next epoch.
func (s *Store) pruneOneEpoch(epoch idx.Epoch, firstBlock idx.Block) (int, idx.Block, error) {
	count := 0
	prefix := epoch.Bytes()

	// Collect event IDs from iterator keys for cache eviction before deleting.
	// Iterator keys hold the canonical ID used as the cache key; re-decoding
	// the event from its value and calling ID() produces a different hash.
	var eventIDs []hash.Event
	{
		it := s.table.Events.NewIterator(prefix, nil)
		for it.Next() {
			eventIDs = append(eventIDs, hash.BytesToEvent(it.Key()))
		}
		it.Release()
	}

	// Retrieve block range for this epoch before deleting BlockEpochStateHistory.
	bs, _ := s.GetHistoryBlockEpochState(epoch)

	// Delete epoch-prefixed tables.
	for _, tbl := range []kvdb.Store{
		s.table.Events,
		s.table.LlrBlockVotes,
		s.table.LlrEpochVotes,
		s.table.LlrEpochVoteIndex,
	} {
		n, err := pruneTableByPrefix(tbl, prefix)
		if err != nil {
			return count, 0, err
		}
		count += n
	}

	// Delete single-key epoch entries.
	for _, tbl := range []kvdb.Store{s.table.LlrEpochResults, s.table.BlockEpochStateHistory} {
		if err := tbl.Delete(prefix); err != nil {
			return count, 0, err
		}
		count++
	}

	var lastBlock idx.Block
	if bs != nil {
		lastBlock = bs.LastBlock.Idx
	}

	// Delete block-indexed entries for this epoch's block range.
	if bs != nil && firstBlock > 0 && lastBlock >= firstBlock {
		for blk := firstBlock; blk <= lastBlock; blk++ {
			blockKey := blk.Bytes()

			if err := s.table.LlrBlockResults.Delete(blockKey); err != nil {
				return count, 0, err
			}
			count++

			// LlrBlockVotesIndex: keyed block(4)+epoch(4)+hash(32); iterate by block prefix.
			n, err := pruneTableByPrefix(s.table.LlrBlockVotesIndex, blockKey)
			if err != nil {
				return count, 0, err
			}
			count += n
		}
	}

	// Evict deleted entries from LRU caches.
	for _, id := range eventIDs {
		s.cache.Events.Remove(id)
		s.cache.EventsHeaders.Remove(id)
		s.cache.EventIDs.Remove(id)
	}
	s.cache.BlockEpochStateHistory.Remove(epoch)
	s.cache.LlrBlockVotesIndex = NewVotesCache(s.cfg.Cache.LlrBlockVotesIndexes, s.flushLlrBlockVoteWeight)
	s.cache.LlrEpochVoteIndex = NewVotesCache(s.cfg.Cache.LlrEpochVotesIndexes, s.flushLlrEpochVoteWeight)

	nextFirstBlock := idx.Block(0)
	if bs != nil {
		nextFirstBlock = lastBlock + 1
	}
	return count, nextFirstBlock, nil
}

func (s *Store) runGossipCompaction() error {
	tables := []kvdb.Store{
		s.table.Events,
		s.table.LlrBlockVotes,
		s.table.LlrEpochVotes,
		s.table.LlrEpochVoteIndex,
		s.table.LlrEpochResults,
		s.table.BlockEpochStateHistory,
		s.table.LlrBlockResults,
		s.table.LlrBlockVotesIndex,
	}
	cstart := time.Now()
	for b := 0x00; b <= 0xf0; b += 0x10 {
		start := []byte{byte(b)}
		end := []byte{byte(b + 0x10)}
		if b == 0xf0 {
			end = nil
		}
		log.Info("Compacting gossip DB",
			"range", fmt.Sprintf("%#x-%#x", start, end),
			"elapsed", common.PrettyDuration(time.Since(cstart)),
		)
		for _, tbl := range tables {
			if err := tbl.Compact(start, end); err != nil {
				return fmt.Errorf("gossip DB compaction failed: %w", err)
			}
		}
	}
	log.Info("Gossip DB compaction done", "elapsed", common.PrettyDuration(time.Since(cstart)))
	return nil
}
