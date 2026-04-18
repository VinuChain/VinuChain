package payback

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	sfcContract "github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
)

// GetAvailablePaybackByAddressRPC computes the currently available payback
// balance for an address in an RPC-safe manner.
//
// Differences from the block-processing path (GetAvailablePaybackByAddress):
//
//   - Never reads or writes pc.blkCtx. Block processing owns blkCtx
//     exclusively; an RPC that touched it would corrupt in-flight block
//     computations under concurrent load.
//   - Never mutates pc.StakesMap. The production path auto-creates an empty
//     epoch entry on first access; RPC traffic must not slowly bloat the
//     StakesMap with empty entries keyed by every address ever queried.
//   - Obtains epoch, rules, and prevEpochState directly from the store.
//   - Runs the four SFC/payback-proxy StaticCalls against the caller-supplied
//     EVM, which the caller has primed against a specific sealed block.
//
// Returns zero (not an error) for the zero address, for networks where
// Podgorica is not active, and for addresses that stake below the minimum.
// Returns an error only when an EVM StaticCall fails or the epoch state is
// unusable.
func (pc *PaybackCache) GetAvailablePaybackByAddressRPC(
	ctx context.Context,
	address common.Address,
	evm *vm.EVM,
) (*big.Int, error) {
	if address == (common.Address{}) {
		return big.NewInt(0), nil
	}
	if pc.store == nil {
		return big.NewInt(0), errors.New("payback store is nil")
	}

	rules := pc.store.GetRules()
	if rules.NetworkID == 0 {
		return big.NewInt(0), errors.New("rules are empty")
	}
	if !rules.Upgrades.Podgorica {
		return big.NewInt(0), nil
	}
	contractAddr := rules.Economy.QuotaCacheAddress
	if contractAddr == (common.Address{}) {
		return big.NewInt(0), nil
	}

	if err := ctx.Err(); err != nil {
		return big.NewInt(0), err
	}

	addressTotalStake, err := pc.getAddressTotalStake(address, evm, contractAddr)
	if err != nil {
		log.Debug("GetAvailablePaybackByAddressRPC: getAddressTotalStake", "err", err)
		return big.NewInt(0), err
	}
	minStake, err := pc.getMinStake(address, evm, contractAddr)
	if err != nil {
		log.Debug("GetAvailablePaybackByAddressRPC: getMinStake", "err", err)
		return big.NewInt(0), err
	}
	if addressTotalStake.Cmp(minStake) < 0 {
		return big.NewInt(0), nil
	}

	if err := ctx.Err(); err != nil {
		return big.NewInt(0), err
	}

	// Read epoch-derived inputs under RLock without touching blkCtx or
	// mutating StakesMap.
	pc.mu.RLock()
	currentEpoch := pc.store.GetCurrentEpoch()
	if currentEpoch < 1 {
		pc.mu.RUnlock()
		return big.NewInt(0), nil
	}
	quotaUsed := pc.getQuotaUsedLocked(address)
	stakesCurrent, stakesPrev := pc.getSumStakeByAddressSplitLocked(address, currentEpoch, currentEpoch-1)
	currentEpochStakes := pc.readStakesForEpochNoCreateLocked(currentEpoch, address)
	pc.mu.RUnlock()

	// Epoch-state fetch happens outside any lock.
	prevEpochState := pc.store.GetHistoryEpochState(currentEpoch - 1)

	blockTime := pc.latestSealedBlockTimeRPC()

	fullDuration := computeFullDurationRPC(stakesCurrent, stakesPrev, currentEpochStakes, prevEpochState, blockTime)

	baseRewardPerSecond, err := pc.getBaseRewardPerSecond(address, evm)
	if err != nil {
		log.Debug("GetAvailablePaybackByAddressRPC: getBaseRewardPerSecond", "err", err)
		return big.NewInt(0), err
	}
	totalStakeSFC, err := pc.getTotalStake(address, sfcContract.ContractAddress, evm)
	if err != nil {
		log.Debug("GetAvailablePaybackByAddressRPC: getTotalStake(SFC)", "err", err)
		return big.NewInt(0), err
	}
	totalStakeQuota, err := pc.getTotalStake(address, contractAddr, evm)
	if err != nil {
		log.Debug("GetAvailablePaybackByAddressRPC: getTotalStake(Quota)", "err", err)
		return big.NewInt(0), err
	}
	sumTotalStake := new(big.Int).Add(totalStakeSFC, totalStakeQuota)
	if sumTotalStake.Sign() == 0 {
		return big.NewInt(0), nil
	}

	multiplier := big.NewInt(1e10)
	addressTotalStakeMultiplied := new(big.Int).Mul(addressTotalStake, multiplier)
	sliceOfBaseRewardPerSecondPercent := new(big.Int).Div(addressTotalStakeMultiplied, sumTotalStake)

	paybackSum := pc.calculatePayback(multiplier, baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent, fullDuration)
	paybackSum = new(big.Int).Sub(paybackSum, quotaUsed)
	if paybackSum.Sign() < 0 {
		return big.NewInt(0), nil
	}
	return paybackSum, nil
}

// readStakesForEpochNoCreateLocked returns a copy of stake entries for the
// address in the given epoch without inserting any new map entries. Caller
// must hold pc.mu (read or write).
func (pc *PaybackCache) readStakesForEpochNoCreateLocked(epoch idx.Epoch, address common.Address) []StakeInfo {
	epochStakes, ok := pc.StakesMap[epoch]
	if !ok {
		return nil
	}
	stakes, ok := epochStakes.StakesByAddress[address]
	if !ok {
		return nil
	}
	dst := make([]StakeInfo, len(stakes))
	copy(dst, stakes)
	return dst
}

// latestSealedBlockTimeRPC returns the block time of the latest sealed block.
// It never consults blkCtx so that block-processing state remains untouched.
func (pc *PaybackCache) latestSealedBlockTimeRPC() time.Time {
	block := pc.store.GetBlock(idx.Block(pc.store.GetLatestBlockIndex()))
	if block == nil {
		return time.Time{}
	}
	return block.Time.Time()
}

// computeFullDurationRPC is a pure function equivalent of
// calculateFullDurationLocked that takes all inputs as parameters, so the
// RPC path does not need access to blkCtx or StakesMap.
func computeFullDurationRPC(
	stakesCurrent, stakesPrev *big.Int,
	currentEpochStakes []StakeInfo,
	prevEpochState *iblockproc.EpochState,
	latestBlockTime time.Time,
) int64 {
	if stakesCurrent.Sign() > 0 && stakesPrev.Sign() == 0 {
		if len(currentEpochStakes) == 0 {
			return 0
		}
		lastStake := currentEpochStakes[len(currentEpochStakes)-1].Timestamp
		dur := latestBlockTime.Sub(lastStake)
		secs := int64(dur / 1e9)
		if secs < 0 {
			return 0
		}
		return secs
	}

	if prevEpochState == nil {
		return 0
	}
	durationPrevEpochBySec := prevEpochState.Duration() / 1e9
	const maxReasonableEpochSec = 365 * 24 * 3600
	if durationPrevEpochBySec > maxReasonableEpochSec {
		return 0
	}
	secs := int64(durationPrevEpochBySec)
	if secs < 0 {
		return 0
	}
	return secs
}
