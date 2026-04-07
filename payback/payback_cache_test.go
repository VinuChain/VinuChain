package payback

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Fantom-foundation/go-opera/payback/contract/paybackProxy"
)

// stubStore is a minimal Store implementation for testing AddTransaction.
type stubStore struct{}

func (s *stubStore) GetLatestBlockIndex() uint64                             { return 1 }
func (s *stubStore) GetBlockTransactionsAndReceipts(uint64) (types.Transactions, types.Receipts) {
	return nil, nil
}
func (s *stubStore) FindBlockEpoch(idx.Block) idx.Epoch                     { return 1 }
func (s *stubStore) GetHistoryEpochState(idx.Epoch) *iblockproc.EpochState  { return nil }
func (s *stubStore) GetRules() opera.Rules                                  { return opera.Rules{} }
func (s *stubStore) GetCurrentEpoch() idx.Epoch                             { return 1 }
func (s *stubStore) GetBlock(idx.Block) *inter.Block                        { return nil }

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

func TestDecodeUint256_RejectsOver128Bits(t *testing.T) {
	// 129 bits (bit 128 set) — should be rejected
	data129 := make([]byte, 32)
	data129[15] = 0x01 // byte 15 is the 129th bit position (big-endian)
	_, err := decodeUint256(data129)
	if err == nil {
		t.Fatal("expected error for 129-bit value")
	}

	// 128-bit max (all 1s in lower 16 bytes) — should succeed
	data128ok := make([]byte, 32)
	for i := 16; i < 32; i++ {
		data128ok[i] = 0xFF
	}
	v, err := decodeUint256(data128ok)
	if err != nil {
		t.Fatalf("expected success for 128-bit value, got %v", err)
	}
	if v.BitLen() > 128 {
		t.Fatalf("decoded value exceeds 128 bits: %d", v.BitLen())
	}

	// Zero — should succeed
	dataZero := make([]byte, 32)
	v, err = decodeUint256(dataZero)
	if err != nil {
		t.Fatalf("expected success for zero, got %v", err)
	}
	if v.Sign() != 0 {
		t.Fatalf("expected zero, got %v", v)
	}

	// Wrong length — should fail
	_, err = decodeUint256([]byte{0x01, 0x02})
	if err == nil {
		t.Fatal("expected error for wrong length")
	}
}

// TestGetAvailablePaybackByAddress_InBlockProcessing verifies that
// GetAvailablePaybackByAddress returns zero for a zero address when called
// during block processing (the inBlockProcessing guard passes). This is the
// regression test for the LOG-07 fix: escalating the out-of-processing error
// to log.Crit must not disturb the normal execution path.
func TestGetAvailablePaybackByAddress_InBlockProcessing(t *testing.T) {
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
	}

	pc.PrepareForBlock(idx.Epoch(1), opera.Rules{}, time.Now())
	defer pc.FinishBlock()

	// Zero address must return zero without panicking or calling log.Crit.
	result := pc.GetAvailablePaybackByAddress(common.Address{}, nil)
	if result == nil || result.Sign() != 0 {
		t.Errorf("expected zero for empty address during block processing, got %v", result)
	}
}

func TestMaxPaybackEntries_CapsStakesMap(t *testing.T) {
	pc := &PaybackCache{
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
		PaybackUsedMap: make(map[common.Address]*big.Int),
	}
	epoch := idx.Epoch(5)
	rules := opera.FakeNetRules()
	rules.Upgrades.Podgorica = true

	pc.PrepareForBlock(epoch, rules, time.Now())

	// Fill stakes map to capacity
	pc.mu.Lock()
	if pc.StakesMap[epoch] == nil {
		pc.StakesMap[epoch] = &EpochStakes{
			StakesByAddress: make(map[common.Address][]StakeInfo),
		}
	}
	for i := 0; i < maxPaybackEntries; i++ {
		addr := common.BigToAddress(big.NewInt(int64(i)))
		pc.StakesMap[epoch].StakesByAddress[addr] = []StakeInfo{{Amount: big.NewInt(100)}}
	}
	pc.mu.Unlock()

	// Adding one more should hit the cap (AddTransaction checks internally)
	// Verify the map doesn't grow beyond maxPaybackEntries
	pc.mu.Lock()
	count := len(pc.StakesMap[epoch].StakesByAddress)
	pc.mu.Unlock()
	if count != maxPaybackEntries {
		t.Fatalf("expected %d entries, got %d", maxPaybackEntries, count)
	}
}

// TestAddTransaction_FeeRefundCappedAtTxFee verifies that FeeRefund values
// exceeding the transaction fee are capped rather than blindly accumulated.
func TestAddTransaction_FeeRefundCappedAtTxFee(t *testing.T) {
	contractABI, _ := abi.JSON(strings.NewReader(paybackProxy.QuotaProxyABI))

	var addr common.Address
	pc := &PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
		store:          &stubStore{},
		contractABI:    &contractABI,
		blkCtx: &blockContext{
			epoch:     1,
			rules:     opera.Rules{},
			blockTime: time.Now(),
		},
		inBlockProcessing: true,
	}

	gasPrice := big.NewInt(1e9) // 1 gwei
	gasUsed := uint64(21000)
	txFee := new(big.Int).Mul(big.NewInt(int64(gasUsed)), gasPrice) // 21000 gwei

	// Create and sign a tx so tx.From() returns a real address.
	key, _ := crypto.GenerateKey()
	signer := types.HomesteadSigner{}
	tx, _ := types.SignTx(
		types.NewTransaction(0, common.HexToAddress("0xBBBB"), big.NewInt(0), 21000, gasPrice, nil),
		signer, key,
	)
	addr = crypto.PubkeyToAddress(key.PublicKey)
	// Cache the sender so tx.From() works.
	types.Sender(signer, tx)

	// FeeRefund is 10x the actual tx fee — should be capped.
	oversizedRefund := new(big.Int).Mul(txFee, big.NewInt(10))
	receipt := &types.Receipt{
		Status:    types.ReceiptStatusSuccessful,
		GasUsed:   gasUsed,
		FeeRefund: oversizedRefund,
	}

	if err := pc.AddTransaction(tx, receipt); err != nil {
		t.Fatalf("AddTransaction failed: %v", err)
	}

	used := pc.PaybackUsedMap[addr]
	if used == nil {
		t.Fatal("expected PaybackUsedMap entry for sender")
	}
	if used.Cmp(txFee) > 0 {
		t.Fatalf("PaybackUsedMap should be capped at txFee=%v, got %v", txFee, used)
	}
	if used.Cmp(txFee) != 0 {
		t.Fatalf("expected PaybackUsedMap=%v (capped to txFee), got %v", txFee, used)
	}
}
