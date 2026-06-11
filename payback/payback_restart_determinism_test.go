package payback

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// ============================================================================
// M1 PaybackCache restart-determinism — RESOLVED
//
// Audit findings A1 (HIGH) / T2 (HIGH):
//   "PaybackCache is volatile yet feeds consensus state (restart-determinism
//    hypothesis). A validator that restarts mid-epoch loses its accumulated
//    PaybackUsedMap and would compute a *larger* available payback (quotaUsed=0)
//    for subsequent blocks than a peer that never restarted."
//
// VERDICT: A1 CONFIRMED as a real consensus bug, then FIXED (Option A — replay
// the current epoch on startup, no schema change).
//
// The confirmed forward-sealing causal chain was:
//
//   1. On startup, svc.paybackCache is constructed EMPTY (service.go,
//      NewPaybackCache with no warm-up). ReexecuteBlocks used a LOCAL
//      reexecCache and did NOT install accumulated quotaUsed back into
//      svc.paybackCache.
//   2. GetConsensusCallbacks (via engine.Bootstrap) passed the still-empty
//      svc.paybackCache into every future block processor.
//   3. The first NEW block sealed after restart read quotaUsed=0 →
//      GetAvailablePaybackByAddress returned the FULL epoch quota →
//      larger FeeRefund → larger AddBalance (state_transition.go) → different
//      block.Root than non-restarted peers → consensus split.
//
// THE FIX (gossip/service.go Start() → gossip/c_block_callbacks.go
// WarmUpPaybackCache): after RecoverEVM and before the engine starts sealing,
// the current epoch's already-sealed (tx, receipt) pairs are replayed through
// the LIVE svc.paybackCache via PrepareForBlock/AddTransaction/FinishBlock —
// exactly the lifecycle the live block processor uses, with the per-block epoch
// derived from GetHistoryEpochState(FindBlockEpoch(b)).Epoch (the same
// derivation ReexecuteBlocks uses). The warmed PaybackUsedMap therefore equals a
// never-restarted node's accumulation at head, so the first new block seals an
// identical FeeRefund and block.Root.
//
// WHAT WAS NEVER BROKEN (and is still correct):
//   - Already-sealed receipts are never recomputed post-restart (pinned in
//     gossip/payback_restart_recovery_test.go).
//   - At an epoch boundary all caches reset to zero, so the bug was always
//     bounded to one epoch (TestPaybackUsedMapResetsAtEpochBoundary).
//
// The tests below now assert CONVERGENCE: a restarted-then-warmed cache produces
// the same quotaUsed (and therefore the same available payback / FeeRefund /
// AddBalance / block.Root) as a never-restarted cache.
// ============================================================================

// makeRefundTx builds a signed tx and a successful receipt carrying feeRefund
// wei, with the sender cached so tx.From() returns the signer address — exactly
// the shape WarmUpPaybackCache feeds into AddTransaction when replaying an
// already-sealed block. Returns the tx, receipt, and the sender address.
func makeRefundTx(t *testing.T, feeRefund *big.Int) (*types.Transaction, *types.Receipt, common.Address) {
	t.Helper()
	gasPrice := big.NewInt(1e9)
	gasUsed := uint64(21000)

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	signer := types.HomesteadSigner{}
	tx, err := types.SignTx(
		types.NewTransaction(0, common.HexToAddress("0xBBBB"), big.NewInt(0), gasUsed, gasPrice, nil),
		signer, key,
	)
	if err != nil {
		t.Fatalf("SignTx: %v", err)
	}
	// Cache the sender so tx.From() works (mirrors the live tx.AsMessage path
	// and the explicit types.Sender call WarmUpPaybackCache performs).
	if _, err := types.Sender(signer, tx); err != nil {
		t.Fatalf("Sender: %v", err)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	receipt := &types.Receipt{
		Status:    types.ReceiptStatusSuccessful,
		GasUsed:   gasUsed,
		FeeRefund: new(big.Int).Set(feeRefund),
	}
	return tx, receipt, addr
}

// stakeRulesForWarmUp returns opera.Rules whose Economy.QuotaCacheAddress matches
// testQuotaContractAddress so that stake() txs sent to it are recorded into
// StakesMap by AddTransaction during the replay (mirrors the live rules a block
// processor uses).
func stakeRulesForWarmUp() opera.Rules {
	var rules opera.Rules
	rules.Economy.QuotaCacheAddress = testQuotaContractAddress
	return rules
}

// makeStakeTx builds a signed stake() tx (selector only) sent to the quota
// contract, with the sender cached, plus a successful receipt. AddTransaction
// records it into StakesMap[epoch] keyed by the sender with Amount = tx.Value().
// This is the exact shape WarmUpPaybackCache feeds into AddTransaction when
// replaying an already-sealed stake block. The signing key is returned so the
// caller can sign a second stake() tx from the SAME address (e.g. staking again
// in a later epoch).
func makeStakeTx(t *testing.T, stakeAmount *big.Int) (*types.Transaction, *types.Receipt, common.Address, *ecdsa.PrivateKey) {
	t.Helper()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	tx, receipt := signStakeTx(t, key, 0, stakeAmount)
	return tx, receipt, crypto.PubkeyToAddress(key.PublicKey), key
}

// signStakeTx signs a stake() tx with the given key and nonce, caching the
// sender, and returns it with a successful receipt.
func signStakeTx(t *testing.T, key *ecdsa.PrivateKey, nonce uint64, stakeAmount *big.Int) (*types.Transaction, *types.Receipt) {
	t.Helper()
	gasPrice := big.NewInt(1e9)
	gasUsed := uint64(50000)
	signer := types.HomesteadSigner{}
	tx, err := types.SignTx(
		types.NewTransaction(nonce, testQuotaContractAddress, new(big.Int).Set(stakeAmount), gasUsed, gasPrice, stakeSelector),
		signer, key,
	)
	if err != nil {
		t.Fatalf("SignTx: %v", err)
	}
	if _, err := types.Sender(signer, tx); err != nil {
		t.Fatalf("Sender: %v", err)
	}
	return tx, &types.Receipt{
		Status:  types.ReceiptStatusSuccessful,
		GasUsed: gasUsed,
	}
}

// sumStakesForEpoch returns the per-address sum of recorded stake amounts for the
// given epoch as decimal strings (pointer-identity-independent for require.Equal).
func sumStakesForEpoch(pc *PaybackCache, epoch idx.Epoch) map[common.Address]string {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	out := make(map[common.Address]string)
	es, ok := pc.StakesMap[epoch]
	if !ok {
		return out
	}
	for addr, stakes := range es.StakesByAddress {
		sum := big.NewInt(0)
		for _, st := range stakes {
			sum.Add(sum, st.Amount)
		}
		out[addr] = sum.String()
	}
	return out
}

// TestWarmUpReconstructsStakesMapForPreviousEpoch is the cache-layer property
// test for StakesMap[E-1] reconstruction (audit rework). It pins that replaying
// the already-sealed blocks of BOTH epoch E-1 and epoch E reconstructs StakesMap
// for the PREVIOUS epoch E-1, not only the current epoch E.
//
// An address that staked in BOTH E-1 and E takes different branches in
// calculateFullDurationLocked depending on whether StakesMap[E-1] is populated
// (getSumStakeByAddressSplitLocked returns sumPrev>0 → prevEpochState.Duration()
// branch; sumPrev==0 → time-since-last-stake branch). A warm-up that replayed
// only epoch E would leave StakesMap[E-1] empty → wrong branch → wrong
// fullDuration → wrong FeeRefund → consensus split.
//
// This is a cache-layer property check with an explicit CONTRAST leg: a third
// cache replays only epoch E's blocks (the shape of a hypothetical broken
// warm-up) and must end with StakesMap[E-1] empty — demonstrating in-test that
// the E-1..E replay window is what the property hinges on, not merely that two
// identical replays agree. The production start-block selection itself is
// pinned by the gossip-level
// TestWarmUpReconstructsStakesMapForPreviousEpochEndToEnd, which drives
// WarmUpPaybackCache directly (the audit finding P1 claiming an epoch-E-only
// start was REFUTED).
func TestWarmUpReconstructsStakesMapForPreviousEpoch(t *testing.T) {
	const prevEpoch = idx.Epoch(4)
	const curEpoch = idx.Epoch(5)
	rules := stakeRulesForWarmUp()
	blockTime := time.Now()

	// One staking address stakes in E-1 and again in E (the dangerous case),
	// plus a second address that stakes only in E.
	stakeE1Tx, stakeE1Rcpt, stakerBoth, bothKey := makeStakeTx(t, big.NewInt(1000))
	stakeEOnlyTx, stakeEOnlyRcpt, stakerEOnly, _ := makeStakeTx(t, big.NewInt(500))

	// Re-stake by the SAME address (stakerBoth) in epoch E, signed with the same
	// key so tx.From() resolves to stakerBoth and AddTransaction records it into
	// StakesMap[E] under the same address.
	stakeE2ndTx, stakeE2ndRcpt := signStakeTx(t, bothKey, 1, big.NewInt(2000))

	// Build the already-sealed block sequence:
	//   epoch E-1: [stakeE1Tx]  (stakerBoth stakes)
	//   epoch E:   [stakeEOnlyTx] (stakerEOnly stakes), [stakeE2ndTx] (stakerBoth re-stakes)

	type sealedBlock struct {
		epoch idx.Epoch
		txs   []*types.Transaction
		rcpts []*types.Receipt
	}
	blocks := []sealedBlock{
		{epoch: prevEpoch, txs: []*types.Transaction{stakeE1Tx}, rcpts: []*types.Receipt{stakeE1Rcpt}},
		{epoch: curEpoch, txs: []*types.Transaction{stakeEOnlyTx, stakeE2ndTx}, rcpts: []*types.Receipt{stakeEOnlyRcpt, stakeE2ndRcpt}},
	}

	replayAll := func(pc *PaybackCache) {
		for _, blk := range blocks {
			pc.PrepareForBlock(blk.epoch, rules, blockTime)
			for i := range blk.txs {
				if err := pc.AddTransaction(blk.txs[i], blk.rcpts[i]); err != nil {
					t.Fatalf("AddTransaction: %v", err)
				}
			}
			pc.FinishBlock()
		}
	}

	// Never-restarted node: records stakes as blocks are sealed across E-1 and E.
	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	replayAll(warm)

	// Restarted node: empty cache, then warm-up replays the SAME blocks
	// (including epoch E-1's stake block — this is the P1 fix).
	restarted, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	replayAll(restarted)

	// CONTRAST LEG: an epoch-E-only replay (the hypothetical broken warm-up
	// shape) must leave StakesMap[E-1] EMPTY — proving the property below is
	// decided by the replay window, not satisfied by any replay whatsoever.
	eOnly, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	for _, blk := range blocks {
		if blk.epoch != curEpoch {
			continue
		}
		eOnly.PrepareForBlock(blk.epoch, rules, blockTime)
		for i := range blk.txs {
			if err := eOnly.AddTransaction(blk.txs[i], blk.rcpts[i]); err != nil {
				t.Fatalf("AddTransaction (E-only leg): %v", err)
			}
		}
		eOnly.FinishBlock()
	}
	if got := sumStakesForEpoch(eOnly, prevEpoch); len(got) != 0 {
		t.Fatalf("epoch-E-only replay unexpectedly populated StakesMap[E-1] (%v); the convergence property below would be vacuous", got)
	}
	if got := sumStakesForEpoch(eOnly, curEpoch); len(got) == 0 {
		t.Fatal("epoch-E-only replay must still populate StakesMap[E]; contrast leg is broken otherwise")
	}

	// The never-restarted node MUST hold StakesMap for BOTH E-1 and E
	// (cleanupOldEpochsLocked preserves epoch >= E-2). Assert non-vacuously.
	warmPrev := sumStakesForEpoch(warm, prevEpoch)
	warmCur := sumStakesForEpoch(warm, curEpoch)
	if len(warmPrev) == 0 {
		t.Fatal("never-restarted cache must hold StakesMap[E-1]; test would be vacuous otherwise")
	}
	if _, ok := warmPrev[stakerBoth]; !ok {
		t.Fatalf("never-restarted cache StakesMap[E-1] must contain stakerBoth %s", stakerBoth.Hex())
	}

	// The restarted-then-warmed cache must reconstruct BOTH epochs identically.
	gotPrev := sumStakesForEpoch(restarted, prevEpoch)
	gotCur := sumStakesForEpoch(restarted, curEpoch)

	if len(gotPrev) == 0 {
		t.Fatal("CONVERGENCE FAILURE: warmed cache StakesMap[E-1] is EMPTY — the warm-up replayed only epoch E. " +
			"An address that staked in both E-1 and E would take a different duration branch → different FeeRefund → split.")
	}
	if !equalStakeSums(warmPrev, gotPrev) {
		t.Fatalf("StakesMap[E-1] mismatch: never-restarted=%v warmed=%v", warmPrev, gotPrev)
	}
	if !equalStakeSums(warmCur, gotCur) {
		t.Fatalf("StakesMap[E] mismatch: never-restarted=%v warmed=%v", warmCur, gotCur)
	}

	// Sanity: stakerEOnly appears only in E, stakerBoth appears in both.
	if _, ok := gotCur[stakerEOnly]; !ok {
		t.Fatalf("warmed StakesMap[E] must contain stakerEOnly %s", stakerEOnly.Hex())
	}
	if _, ok := gotPrev[stakerBoth]; !ok {
		t.Fatalf("warmed StakesMap[E-1] must contain stakerBoth %s", stakerBoth.Hex())
	}
	if _, ok := gotCur[stakerBoth]; !ok {
		t.Fatalf("warmed StakesMap[E] must contain stakerBoth %s (re-staked in E)", stakerBoth.Hex())
	}
}

// equalStakeSums compares two address→decimal-string stake-sum maps.
func equalStakeSums(a, b map[common.Address]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// TestForwardSealingConvergesAfterMidEpochRestart is the success-criterion test
// for the A1 fix. It models the exact production scenario:
//
//   - A never-restarted "warm" node accumulates quotaUsed across earlier
//     intra-epoch blocks (via AddTransaction, the live mechanism).
//   - A "restarted" node starts with an EMPTY cache (as svc.paybackCache is
//     constructed) and then runs the warm-up: it replays the SAME already-sealed
//     (tx, receipt) pairs of the current epoch through AddTransaction — exactly
//     what WarmUpPaybackCache does.
//
// After warm-up, both caches must report the SAME quotaUsed for the next
// (to-be-sealed) block. Identical quotaUsed → identical available payback →
// identical FeeRefund → identical AddBalance → identical block.Root. No split.
func TestForwardSealingConvergesAfterMidEpochRestart(t *testing.T) {
	const epoch = idx.Epoch(5)
	rules := opera.Rules{}
	blockTime := time.Now()

	// Two already-sealed blocks earlier in the epoch, each consuming quota for a
	// different sender. PaybackUsedMap is keyed per address, so per-address
	// convergence is the exact consensus property the fix must guarantee.
	tx1, rcpt1, addr1 := makeRefundTx(t, big.NewInt(300))
	tx2, rcpt2, addr2 := makeRefundTx(t, big.NewInt(200))
	txs := []*types.Transaction{tx1, tx2}
	rcpts := []*types.Receipt{rcpt1, rcpt2}

	// replay mirrors WarmUpPaybackCache: one PrepareForBlock/AddTransaction/
	// FinishBlock cycle per already-sealed block of the current epoch.
	replay := func(pc *PaybackCache) {
		for i := range txs {
			pc.PrepareForBlock(epoch, rules, blockTime)
			if err := pc.AddTransaction(txs[i], rcpts[i]); err != nil {
				t.Fatalf("AddTransaction: %v", err)
			}
			pc.FinishBlock()
		}
	}

	// ---- Warm cache: node that never restarted. ----
	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	replay(warm)

	// ---- Restarted cache: starts empty, then runs the warm-up replay. ----
	restarted, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	// Sanity: a fresh cache really is empty before warm-up (the pre-fix state
	// that produced the divergence).
	restarted.PrepareForBlock(epoch, rules, blockTime)
	restarted.mu.RLock()
	preWarm := restarted.getQuotaUsedLocked(addr1)
	restarted.mu.RUnlock()
	restarted.FinishBlock()
	if preWarm.Sign() != 0 {
		t.Fatalf("fresh cache must be empty before warm-up, got %s", preWarm)
	}
	// Warm-up: replay the SAME already-sealed blocks into the live cache.
	replay(restarted)

	// ---- Next block to be sealed: both caches must agree per address. ----
	warm.PrepareForBlock(epoch, rules, blockTime)
	restarted.PrepareForBlock(epoch, rules, blockTime)
	defer warm.FinishBlock()
	defer restarted.FinishBlock()

	for _, a := range []common.Address{addr1, addr2} {
		warm.mu.RLock()
		warmUsed := warm.getQuotaUsedLocked(a)
		warm.mu.RUnlock()
		restarted.mu.RLock()
		restartedUsed := restarted.getQuotaUsedLocked(a)
		restarted.mu.RUnlock()
		if warmUsed.Cmp(restartedUsed) != 0 {
			t.Fatalf("CONVERGENCE FAILURE for %s: warm quotaUsed=%s, restarted (warmed) quotaUsed=%s — "+
				"the A1 fix is broken; a mid-epoch restart would seal a different block.Root",
				a.Hex(), warmUsed, restartedUsed)
		}
	}

	// Assert the accumulation is the expected non-zero value, so a regression
	// that silently warms to zero (which would also "converge" but to the wrong
	// value) is caught.
	warm.mu.RLock()
	got1 := warm.getQuotaUsedLocked(addr1)
	got2 := warm.getQuotaUsedLocked(addr2)
	warm.mu.RUnlock()
	if got1.Cmp(big.NewInt(300)) != 0 {
		t.Fatalf("warm cache quotaUsed for addr1 want 300, got %s", got1)
	}
	if got2.Cmp(big.NewInt(200)) != 0 {
		t.Fatalf("warm cache quotaUsed for addr2 want 200, got %s", got2)
	}
}

// TestPaybackCacheIsVolatileWithinEpoch demonstrates the cache-layer volatility
// as an isolated property: accumulated quotaUsed in a warm cache differs from a
// freshly-constructed one at mid-epoch. This underpins the confirmed bug.
func TestPaybackCacheIsVolatileWithinEpoch(t *testing.T) {
	const epoch = idx.Epoch(5)
	rules := opera.Rules{}
	blockTime := time.Now()
	addr := common.HexToAddress("0xA11CE")

	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr] = big.NewInt(300)
	warm.mu.Unlock()
	warm.FinishBlock()

	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr].Add(warm.PaybackUsedMap[addr], big.NewInt(200))
	warm.mu.Unlock()
	warm.mu.RLock()
	warmQuotaUsed := warm.getQuotaUsedLocked(addr)
	warm.mu.RUnlock()
	warm.FinishBlock()

	if warmQuotaUsed.Cmp(big.NewInt(500)) != 0 {
		t.Fatalf("warm cache quotaUsed: want 500, got %s", warmQuotaUsed)
	}

	fresh, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	fresh.PrepareForBlock(epoch, rules, blockTime)
	fresh.mu.RLock()
	freshQuotaUsed := fresh.getQuotaUsedLocked(addr)
	fresh.mu.RUnlock()
	fresh.FinishBlock()

	if freshQuotaUsed.Sign() != 0 {
		t.Fatalf("fresh cache quotaUsed: want 0, got %s", freshQuotaUsed)
	}

	if warmQuotaUsed.Cmp(freshQuotaUsed) == 0 {
		t.Fatal("expected warm and fresh caches to diverge mid-epoch")
	}
}

// TestPaybackUsedMapResetsAtEpochBoundary proves that warm and fresh caches
// converge at epoch boundaries. This bounds the A1 bug to within a single epoch:
// a restart that lands ON or AFTER an epoch boundary cannot diverge at all.
func TestPaybackUsedMapResetsAtEpochBoundary(t *testing.T) {
	rules := opera.Rules{}
	blockTime := time.Now()
	addr := common.HexToAddress("0xB0B")

	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	warm.PrepareForBlock(idx.Epoch(7), rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr] = big.NewInt(900)
	warm.mu.Unlock()
	warm.FinishBlock()

	// Cross into epoch 8 — cleanup resets PaybackUsedMap.
	warm.PrepareForBlock(idx.Epoch(8), rules, blockTime)
	warm.mu.RLock()
	warmQuotaUsed := warm.getQuotaUsedLocked(addr)
	warm.mu.RUnlock()
	warm.FinishBlock()

	fresh, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	fresh.PrepareForBlock(idx.Epoch(8), rules, blockTime)
	fresh.mu.RLock()
	freshQuotaUsed := fresh.getQuotaUsedLocked(addr)
	fresh.mu.RUnlock()
	fresh.FinishBlock()

	if warmQuotaUsed.Sign() != 0 {
		t.Fatalf("warm cache after epoch boundary: want 0, got %s", warmQuotaUsed)
	}
	if freshQuotaUsed.Sign() != 0 {
		t.Fatalf("fresh cache in new epoch: want 0, got %s", freshQuotaUsed)
	}
	if warmQuotaUsed.Cmp(freshQuotaUsed) != 0 {
		t.Fatalf("warm (%s) != fresh (%s) across epoch boundary", warmQuotaUsed, freshQuotaUsed)
	}
}
