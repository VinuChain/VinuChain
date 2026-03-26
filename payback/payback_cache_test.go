package payback

import (
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
)

// TestEpochCleanupRunsDuringPrepareForBlock verifies that PaybackUsedMap is
// reset when the epoch advances. This is the regression test for RA-PB-001:
// prior to the fix, cleanupOldEpochsLocked was dead code because the
// inBlockProcessing guard prevented it from ever executing.
func TestEpochCleanupRunsDuringPrepareForBlock(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	addr := common.HexToAddress("0x1234")

	// Simulate epoch 5: populate PaybackUsedMap with quota usage
	pc.PaybackUsedMap[addr] = big.NewInt(1000)
	pc.StakesMap[idx.Epoch(5)] = &EpochStakes{
		StakesByAddress: make(map[common.Address][]StakeInfo),
	}

	rules := opera.Rules{}
	blockTime := time.Now()

	// PrepareForBlock in epoch 5 — first call, should clean up
	pc.PrepareForBlock(idx.Epoch(5), rules, blockTime)
	pc.FinishBlock()

	// PaybackUsedMap should be reset since lastCleanedEpoch was 0 < 5
	if len(pc.PaybackUsedMap) != 0 {
		t.Fatalf("expected PaybackUsedMap to be empty after epoch transition, got %d entries", len(pc.PaybackUsedMap))
	}

	// Simulate usage in epoch 5
	pc.PaybackUsedMap[addr] = big.NewInt(500)

	// PrepareForBlock again in epoch 5 — should NOT clean (already cleaned)
	pc.PrepareForBlock(idx.Epoch(5), rules, blockTime)
	pc.FinishBlock()

	if pc.PaybackUsedMap[addr] == nil || pc.PaybackUsedMap[addr].Cmp(big.NewInt(500)) != 0 {
		t.Fatalf("expected PaybackUsedMap to retain epoch 5 data on same-epoch PrepareForBlock, got %v", pc.PaybackUsedMap[addr])
	}

	// Advance to epoch 6 — should clean up epoch 5's used quota
	pc.PaybackUsedMap[addr] = big.NewInt(750)
	pc.PrepareForBlock(idx.Epoch(6), rules, blockTime)
	pc.FinishBlock()

	if len(pc.PaybackUsedMap) != 0 {
		t.Fatalf("expected PaybackUsedMap to be empty after epoch 6 transition, got %d entries", len(pc.PaybackUsedMap))
	}
}

// TestStakesMapPrunedOnEpochAdvance verifies that old epochs are removed
// from StakesMap when the epoch advances past the cutoff (currentEpoch - 2).
func TestStakesMapPrunedOnEpochAdvance(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	// Populate epochs 3, 4, 5
	for _, e := range []idx.Epoch{3, 4, 5} {
		pc.StakesMap[e] = &EpochStakes{
			StakesByAddress: make(map[common.Address][]StakeInfo),
		}
	}

	rules := opera.Rules{}
	blockTime := time.Now()

	// Advance to epoch 6 — cutoff is 6-2=4, so epoch 3 should be pruned
	pc.PrepareForBlock(idx.Epoch(6), rules, blockTime)
	pc.FinishBlock()

	if _, ok := pc.StakesMap[idx.Epoch(3)]; ok {
		t.Fatal("expected epoch 3 to be pruned from StakesMap (cutoff=4)")
	}
	if _, ok := pc.StakesMap[idx.Epoch(4)]; !ok {
		t.Fatal("expected epoch 4 to be retained in StakesMap")
	}
	if _, ok := pc.StakesMap[idx.Epoch(5)]; !ok {
		t.Fatal("expected epoch 5 to be retained in StakesMap")
	}
}

// TestStakesMapUnboundedGrowthPrevented verifies that StakesMap does not grow
// without bound across many epochs. After processing 100 epochs, only epochs
// within the retention window (currentEpoch - 2) should remain.
func TestStakesMapUnboundedGrowthPrevented(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	addr := common.HexToAddress("0xDEAD")
	rules := opera.Rules{}
	blockTime := time.Now()

	const totalEpochs = 100

	for epoch := idx.Epoch(1); epoch <= totalEpochs; epoch++ {
		pc.StakesMap[epoch] = &EpochStakes{
			StakesByAddress: map[common.Address][]StakeInfo{
				addr: {{Amount: big.NewInt(int64(epoch) * 1000), Timestamp: blockTime}},
			},
		}
		pc.PaybackUsedMap[addr] = big.NewInt(int64(epoch))

		pc.PrepareForBlock(epoch, rules, blockTime)
		pc.FinishBlock()
	}

	// Retention window is [currentEpoch-2, currentEpoch], so at most 3 epochs.
	const maxRetained = 3
	if len(pc.StakesMap) > maxRetained {
		t.Fatalf("expected at most %d epochs in StakesMap after %d epoch transitions, got %d",
			maxRetained, totalEpochs, len(pc.StakesMap))
	}

	cutoff := idx.Epoch(totalEpochs) - 2
	for epoch := range pc.StakesMap {
		if epoch < cutoff {
			t.Fatalf("epoch %d should have been pruned (cutoff=%d)", epoch, cutoff)
		}
	}

	if _, ok := pc.StakesMap[totalEpochs]; !ok {
		t.Fatal("expected current epoch to be retained")
	}
	if _, ok := pc.StakesMap[totalEpochs-1]; !ok {
		t.Fatal("expected previous epoch (N-1) to be retained")
	}
	if _, ok := pc.StakesMap[totalEpochs-2]; !ok {
		t.Fatal("expected epoch N-2 to be retained")
	}
}

// TestStakesMapPruningAtEarlyEpochs verifies that cleanup handles the edge
// case where currentEpoch <= 2 without underflow.
func TestStakesMapPruningAtEarlyEpochs(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	addr := common.HexToAddress("0xBEEF")
	rules := opera.Rules{}
	blockTime := time.Now()

	pc.StakesMap[1] = &EpochStakes{
		StakesByAddress: map[common.Address][]StakeInfo{
			addr: {{Amount: big.NewInt(500), Timestamp: blockTime}},
		},
	}
	pc.PaybackUsedMap[addr] = big.NewInt(100)

	pc.PrepareForBlock(idx.Epoch(1), rules, blockTime)
	pc.FinishBlock()

	if len(pc.PaybackUsedMap) != 0 {
		t.Fatalf("expected PaybackUsedMap to be reset at epoch 1, got %d entries", len(pc.PaybackUsedMap))
	}

	if _, ok := pc.StakesMap[1]; !ok {
		t.Fatal("expected epoch 1 stakes to be retained when currentEpoch <= 2")
	}

	pc.PrepareForBlock(idx.Epoch(2), rules, blockTime)
	pc.FinishBlock()

	if _, ok := pc.StakesMap[1]; !ok {
		t.Fatal("expected epoch 1 stakes to be retained when currentEpoch == 2")
	}
}

// TestCleanupDoesNotRunDuringBlockProcessing verifies that the
// inBlockProcessing flag prevents cleanup from running mid-block,
// even if called directly (defense in depth).
func TestCleanupIdempotentWithinSameEpoch(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	addr := common.HexToAddress("0xABCD")
	rules := opera.Rules{}
	blockTime := time.Now()

	// First block in epoch 10 — cleanup runs
	pc.PrepareForBlock(idx.Epoch(10), rules, blockTime)

	// Simulate quota usage during block processing
	pc.PaybackUsedMap[addr] = big.NewInt(999)

	pc.FinishBlock()

	// Second block still in epoch 10 — cleanup should NOT run
	pc.PrepareForBlock(idx.Epoch(10), rules, blockTime)
	pc.FinishBlock()

	// Quota from first block should persist within the same epoch
	if pc.PaybackUsedMap[addr] == nil || pc.PaybackUsedMap[addr].Cmp(big.NewInt(999)) != 0 {
		t.Fatalf("expected quota to persist within same epoch, got %v", pc.PaybackUsedMap[addr])
	}
}
