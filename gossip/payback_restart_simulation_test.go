package gossip

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/payback"
	"github.com/Fantom-foundation/go-opera/utils/signers/internaltx"
)

// TestWarmUpPaybackCacheConvergesWithNeverRestartedNode is the end-to-end
// success-criterion test for the A1 fix (PaybackCache restart determinism),
// covering the PaybackUsedMap divergence mechanism:
//
//   - PaybackUsedMap (A1): the warmed cache must reconstruct the never-restarted
//     accumulation of FeeRefund-per-address for the current epoch E.
//
// The fakenet genesis deploys no quota contract at QuotaCacheAddress, so the live
// FeeRefund is always 0 and PaybackUsedMap would never be populated by sealing
// alone (this is exactly why the earlier version of this test was VACUOUS —
// comparing two empty maps). To make it non-vacuous we inject crafted non-zero
// FeeRefund into the RAW stored receipts of the sealed current-epoch blocks, then
// derive the never-restarted reference by replaying those same (tx, receipt)
// pairs, and assert the warm-up reproduces them exactly.
func TestWarmUpPaybackCacheConvergesWithNeverRestartedNode(t *testing.T) {
	const validators = idx.Validator(3)
	env := newTestEnv(2, validators)
	defer env.Close()

	const sameEpoch = time.Second
	startEpoch := env.store.GetEpoch()

	// Seal several blocks of transfers within one epoch so the store holds real
	// txs and raw receipts for the current epoch.
	for round := 0; round < 4; round++ {
		_, err := env.ApplyTxs(sameEpoch,
			env.Transfer(1, 2, big.NewInt(1000)),
			env.Transfer(2, 3, big.NewInt(1000)),
			env.Transfer(3, 1, big.NewInt(1000)),
		)
		require.NoError(t, err)
	}
	require.Equal(t, startEpoch, env.store.GetEpoch(),
		"test must stay within one epoch so PaybackUsedMap is not reset mid-run")

	head := env.store.GetLatestBlockIndex()
	firstBlock := env.paybackWarmUpFirstBlock(env.store.GetEpoch(), head)
	require.LessOrEqual(t, firstBlock, head)

	// Use the same signer the env signed with so sender recovery matches the
	// warm-up's (both LatestSignerForChainID over NetworkID==ChainID).
	signer := env.EthAPI.signer

	// Inject a crafted non-zero FeeRefund into the first NON-INTERNAL tx of every
	// tx-bearing block in the warm-up window, and build the reference
	// never-restarted PaybackUsedMap by accumulating those refunds per sender.
	// Internal txs (sealing/SFC) carry no recoverable EIP-155 signature and are
	// dropped by AddTransaction on both the live and warm-up paths, so we skip
	// them here too.
	expectedUsed := make(map[common.Address]*big.Int)
	injected := 0
	feeRefund := big.NewInt(7777)
	for b := firstBlock; b <= head; b++ {
		block := env.store.GetBlock(b)
		if block == nil {
			continue
		}
		txs := env.store.GetBlockTxs(b, block)
		if len(txs) == 0 {
			continue
		}
		stored, _ := env.store.evm.GetRawReceipts(b)
		if len(stored) != len(txs) {
			continue
		}
		userIdx := -1
		for i, tx := range txs {
			if internaltx.IsInternal(tx) {
				continue
			}
			if _, err := types.Sender(signer, tx); err != nil {
				continue
			}
			userIdx = i
			break
		}
		if userIdx < 0 {
			continue
		}
		// Add a non-zero FeeRefund to the chosen user tx's stored receipt.
		stored[userIdx].FeeRefund = new(big.Int).Set(feeRefund)
		env.store.evm.SetRawReceipts(b, stored)

		from, err := types.Sender(signer, txs[userIdx])
		require.NoError(t, err)
		if expectedUsed[from] == nil {
			expectedUsed[from] = big.NewInt(0)
		}
		expectedUsed[from].Add(expectedUsed[from], feeRefund)
		injected++
	}
	require.Greater(t, injected, 0,
		"test setup must inject FeeRefund into at least one sealed block or it is vacuous")

	expectedSnapshot := make(map[common.Address]string, len(expectedUsed))
	for a, v := range expectedUsed {
		expectedSnapshot[a] = v.String()
	}
	require.NotEmpty(t, expectedSnapshot,
		"the never-restarted reference PaybackUsedMap must be non-empty (guards against a vacuous test)")

	// ---- Simulate a mid-epoch restart: rebuild svc.paybackCache empty. ----
	freshStore := NewPaybackStore(env.store)
	fresh, err := payback.NewPaybackCache(freshStore, env.store.GetRules().Economy.QuotaCacheMaxAddresses)
	require.NoError(t, err)
	env.paybackCache = fresh

	require.Empty(t, snapshotUsedMap(env.paybackCache),
		"a freshly constructed payback cache must be empty before warm-up")

	// ---- Run the A1 fix. ----
	env.WarmUpPaybackCache()

	// ---- Assert non-vacuous convergence. ----
	warmedSnapshot := snapshotUsedMap(env.paybackCache)
	require.NotEmpty(t, warmedSnapshot,
		"WarmUpPaybackCache must reconstruct a NON-EMPTY PaybackUsedMap from the injected FeeRefunds; "+
			"an empty result means the warm-up read the wrong receipts or skipped the blocks")
	require.Equal(t, expectedSnapshot, warmedSnapshot,
		"WarmUpPaybackCache must reconstruct the never-restarted PaybackUsedMap exactly; "+
			"any difference would make the restarted node seal a different block.Root")
}

// TestWarmUpReconstructsStakesMapForPreviousEpochEndToEnd pins the StakesMap[E-1]
// reconstruction property end-to-end. It seals real stake() txs to the fakenet
// QuotaCacheAddress across an epoch boundary so the live block processor records
// StakesMap[E-1] and StakesMap[E], then simulates a mid-epoch restart and asserts
// WarmUpPaybackCache reconstructs StakesMap for the PREVIOUS epoch E-1.
//
// StakesMap[E-1] is consensus-relevant: calculateFullDurationLocked branches on
// whether an address has stakes in BOTH E and E-1, so a warm-up that began at the
// first block of epoch E (leaving StakesMap[E-1] empty) would diverge. This test
// FAILS against such a true-epoch-E-only start (verified during the audit rework:
// forcing the start to the first block of FindBlockEpoch==E leaves StakesMap[E-1]
// empty and the require.NotEmpty/Equal assertions below fire).
//
// NOTE: the production warm-up's start block — GetHistoryBlockEpochState(E-1)
// .LastBlock.Idx + 1 — is the first block of epoch E-1 (see paybackWarmUpFirstBlock
// for the epoch-index convention), so it already satisfies this property. The
// audit finding P1 (which claimed that lookup yielded the first block of epoch E)
// was REFUTED; this test guards the property against future regressions.
func TestWarmUpReconstructsStakesMapForPreviousEpochEndToEnd(t *testing.T) {
	const validators = idx.Validator(3)
	env := newTestEnv(2, validators)
	defer env.Close()

	quotaAddr := env.store.GetRules().Economy.QuotaCacheAddress
	require.NotEqual(t, common.Address{}, quotaAddr,
		"fakenet rules must bind a QuotaCacheAddress for stake() detection")

	// Stake in epoch E-1 (validator 1 stakes to the quota contract). Capture the
	// epoch the stake actually sealed in.
	_, err := env.ApplyTxs(time.Second, stakeTx(env, 1, big.NewInt(111)))
	require.NoError(t, err)
	prevEpoch := env.store.GetEpoch()

	// Advance EXACTLY one epoch deterministically (EmitUntil stops at the first
	// boundary crossing), so the prior stake lands in E-1 and is preserved into E.
	env.t = env.t.Add(nextEpoch)
	require.NoError(t, env.EmitUntil(func() bool {
		return env.store.GetEpoch() > prevEpoch
	}))
	curEpoch := env.store.GetEpoch()
	require.Equal(t, prevEpoch+1, curEpoch,
		"exactly one epoch must elapse so StakesMap[E-1] is preserved (cleanup deletes only < E-2)")

	// Stake again in the current epoch E (same staker) so the address has stakes
	// in BOTH E-1 and E — the dangerous calculateFullDurationLocked case.
	_, err = env.ApplyTxs(time.Second, stakeTx(env, 1, big.NewInt(222)))
	require.NoError(t, err)
	require.Equal(t, curEpoch, env.store.GetEpoch(),
		"the second stake must stay within epoch E")

	// Never-restarted reference: the live cache must hold StakesMap for E-1 and E.
	warmPrev := env.paybackCache.SnapshotStakesByEpoch(prevEpoch)
	warmCur := env.paybackCache.SnapshotStakesByEpoch(curEpoch)
	require.NotEmpty(t, warmPrev,
		"never-restarted cache must hold StakesMap[E-1]; test would be vacuous otherwise")
	require.NotEmpty(t, warmCur, "never-restarted cache must hold StakesMap[E]")

	// ---- Simulate a mid-epoch restart. ----
	freshStore := NewPaybackStore(env.store)
	fresh, err := payback.NewPaybackCache(freshStore, env.store.GetRules().Economy.QuotaCacheMaxAddresses)
	require.NoError(t, err)
	env.paybackCache = fresh
	require.Empty(t, fresh.SnapshotStakesByEpoch(prevEpoch),
		"a freshly constructed cache must hold no stakes before warm-up")

	// ---- Run the A1/P1 fix. ----
	env.WarmUpPaybackCache()

	// ---- Assert StakesMap[E-1] (and E) reconstruction. ----
	gotPrev := env.paybackCache.SnapshotStakesByEpoch(prevEpoch)
	gotCur := env.paybackCache.SnapshotStakesByEpoch(curEpoch)
	require.NotEmpty(t, gotPrev,
		"WarmUpPaybackCache must reconstruct StakesMap[E-1]; an empty result means the warm-up "+
			"replayed only epoch E (the OLD bug) — an address staking in both E-1 and E would "+
			"take a different duration branch → different FeeRefund → consensus split")
	require.Equal(t, warmPrev, gotPrev,
		"warmed StakesMap[E-1] must match the never-restarted node's exactly")
	require.Equal(t, warmCur, gotCur,
		"warmed StakesMap[E] must match the never-restarted node's exactly")
}

// stakeTx builds a signed stake() tx (selector only) from validator `from` to the
// fakenet QuotaCacheAddress, so the live block processor records it into
// StakesMap. The tx executes as a no-op against the (codeless) quota address and
// succeeds, which is sufficient for stake recording (AddTransaction records the
// stake purely from tx fields + a Successful receipt).
func stakeTx(env *testEnv, from idx.ValidatorID, amount *big.Int) *types.Transaction {
	sender := env.Address(from)
	nonce, _ := env.PendingNonceAt(context.TODO(), sender)
	env.incNonce(sender)
	key := makefakegenesis.FakeKey(from)
	gp := env.store.GetRules().Economy.MinGasPrice
	quotaAddr := env.store.GetRules().Economy.QuotaCacheAddress
	// stake() selector = first 4 bytes of keccak256("stake()").
	data := crypto.Keccak256([]byte("stake()"))[:4]
	// Use maxGasLimit: a tx carrying calldata needs more than the 21000 intrinsic
	// gas the transfer path uses, else it fails intrinsic-gas validation.
	tx := types.NewTransaction(nonce, quotaAddr, amount, maxGasLimit, gp, data)
	tx, err := types.SignTx(tx, env.EthAPI.signer, key)
	if err != nil {
		panic(err)
	}
	return tx
}

// snapshotUsedMap returns a deterministic copy of a cache's PaybackUsedMap keyed
// by address with decimal-string values, suitable for require.Equal comparison
// independent of *big.Int pointer identity.
func snapshotUsedMap(pc *payback.PaybackCache) map[common.Address]string {
	raw := pc.SnapshotUsedMap()
	out := make(map[common.Address]string, len(raw))
	for addr, v := range raw {
		out[addr] = v.String()
	}
	return out
}
