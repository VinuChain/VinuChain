package payback

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-opera/opera"
	sfcContract "github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/payback/contract/paybackProxy"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxType string

const (
	TxTypeStake   TxType = "stake"
	TxTypeUnstake TxType = "unstake"
	TxTypeNone    TxType = "none"
)

type TxInfo struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Type    TxType
}

type BlockInfo struct {
	BlockNumber   uint64
	Txs           []TxInfo
	BaseFeePerGas *big.Int
}

type StakeInfo struct {
	Amount    *big.Int
	Timestamp time.Time
}

type EpochStakes struct {
	StakesByAddress map[common.Address][]StakeInfo
}

type PaybackCache struct {
	mu sync.RWMutex

	// PaybackUsedMap contains last quota used
	PaybackUsedMap map[common.Address]*big.Int

	// StakesMap contains stakes and unstakes
	StakesMap map[idx.Epoch]*EpochStakes

	// ABI to get names of called methods
	ContractABI *abi.ABI

	store Store

	evm *vm.EVM

	contractAddress common.Address
}

func (pc *PaybackCache) deleteCurrentBlock() error {
	return nil
}

func (pc *PaybackCache) AddTransaction(tx *types.Transaction, receipt *types.Receipt) error {
	if receipt.Status == types.ReceiptStatusSuccessful {
		if pc.store == nil {
			return nil
		}
		epoch := pc.store.GetCurrentEpoch()
		txtype := getTxType(tx, *pc.ContractABI)
		if txtype == TxTypeStake {
			if receipt.BlockNumber == nil {
				return errors.New("receipt.BlockNumber is nil")
			}

			block := pc.store.GetBlock(idx.Block(receipt.BlockNumber.Uint64()))
			if block == nil {
				return fmt.Errorf("block %d not found in store", receipt.BlockNumber.Uint64())
			}

			pc.mu.Lock()
			if pc.StakesMap[epoch] == nil {
				pc.StakesMap[epoch] = &EpochStakes{
					StakesByAddress: make(map[common.Address][]StakeInfo),
				}
			}

			stakes := pc.StakesMap[epoch].StakesByAddress[tx.From()]
			stakes = append(stakes, StakeInfo{
				Amount:    tx.Value(),
				Timestamp: block.Time.Time(),
			})
			pc.StakesMap[epoch].StakesByAddress[tx.From()] = stakes
			pc.mu.Unlock()
		}

		if receipt.FeeRefund != nil && receipt.FeeRefund.Sign() > 0 {
			pc.mu.Lock()
			if _, ok := pc.PaybackUsedMap[tx.From()]; !ok {
				pc.PaybackUsedMap[tx.From()] = big.NewInt(0)
			}
			pc.PaybackUsedMap[tx.From()].Add(pc.PaybackUsedMap[tx.From()], receipt.FeeRefund)
			pc.mu.Unlock()
		}

	}
	return nil
}

func (pc *PaybackCache) GetQuotaUsed(address common.Address) *big.Int {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	if _, ok := pc.PaybackUsedMap[address]; ok {
		return new(big.Int).Set(pc.PaybackUsedMap[address])
	}
	return big.NewInt(0)
}

func NewPaybackCache(store Store) (*PaybackCache, error) {
	pc := PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*EpochStakes),
		store:          store,
	}

	if store.GetRules().Upgrades.Podgorica {
		pc.contractAddress = store.GetRules().Economy.QuotaCacheAddress
		log.Info("NewPaybackCache: PaybackCacheAddress is set", "address", pc.contractAddress.String())
	} else {
		log.Info("HardFork Podgorica is not activated", "status", store.GetRules().Upgrades.Podgorica)
		log.Info("NewPaybackCache:", "Rules", store.GetRules())
	}

	abiQuotaProxy, err := abi.JSON(strings.NewReader(paybackProxy.QuotaProxyABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse payback proxy ABI: %w", err)
	}
	pc.ContractABI = &abiQuotaProxy

	return &pc, nil
}

// String prints all data in cache
func (pc *PaybackCache) String() string {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("PaybackUsedMap: %v\n", pc.PaybackUsedMap))
	sb.WriteString(fmt.Sprintf("StakesMap: %v\n", pc.StakesMap))
	return sb.String()
}

func getTxType(tx *types.Transaction, abi abi.ABI) TxType {
	if tx.Data() != nil && len(tx.Data()) >= 4 {
		if method, err := abi.MethodById(tx.Data()[:4]); err == nil {
			if method.Name == "stake" {
				return TxTypeStake
			}

			if method.Name == "unstake" {
				return TxTypeUnstake
			}
		}
	}
	return TxTypeNone
}

// GetAvailablePaybackByAddress calculates the available quota for a given address.
func (pc *PaybackCache) GetAvailablePaybackByAddress(address common.Address) *big.Int {
	// Initialize payback to zero.
	payback := big.NewInt(0)

	// Return zero payback if the address is empty.
	if address == (common.Address{}) {
		return payback
	}

	// Check and update the contract address.
	if err := pc.checkAndUpdateContractAddress(); err != nil {
		log.Warn("GetAvailablePaybackByAddress:", "error", err)
		return payback
	}

	// Retrieve total stake for the address and handle errors.
	addressTotalStake, err := pc.getAddressTotalStake(address)
	if err != nil {
		log.Warn("GetAvailablePaybackByAddress:", "error", err)
	}

	// Get minimum stake required and handle errors.
	minStake, err := pc.getMinStake(address)
	if err != nil {
		log.Warn("GetAvailablePaybackByAddress:", "error", err)
	}

	// Return zero payback if the address's total stake is below the minimum stake.
	if addressTotalStake.Cmp(minStake) < 0 {
		return payback
	}

	// Retrieve the current epoch.
	currentEpoch := pc.store.GetCurrentEpoch()

	// Guard against epoch underflow: need at least epoch 1 for prevEpoch.
	if currentEpoch < 1 {
		log.Warn("GetAvailablePaybackByAddress: currentEpoch is 0, cannot compute payback")
		return payback
	}
	prevEpoch := currentEpoch - 1

	// Initialize stakes map for the current epoch if not already done.
	pc.mu.Lock()
	if pc.StakesMap[currentEpoch] == nil {
		pc.StakesMap[currentEpoch] = &EpochStakes{
			StakesByAddress: make(map[common.Address][]StakeInfo),
		}

		// Cleanup old epochs.
		pc.cleanupOldEpochsLocked()
	}
	pc.mu.Unlock()

	// Calculate the base reward per second and the total stake from SFC contract and Quota contract.
	baseRewardPerSecond, sumTotalStake, err := pc.calculateStakeDetails(address)
	if err != nil {
		log.Warn("GetAvailablePaybackByAddress:", "error", err)
		return payback
	}

	log.Debug("GetAvailablePaybackByAddress:", "baseRewardPerSecond", baseRewardPerSecond)
	log.Debug("GetAvailablePaybackByAddress:", "sumTotalStake", sumTotalStake)

	// Apply a multiplier for calculation precision.
	multiplier := big.NewInt(1e10)
	addressTotalStakeMultiplied := new(big.Int).Mul(addressTotalStake, multiplier)
	log.Debug("GetAvailablePaybackByAddress:", "addressTotalStakeMultiplied", addressTotalStakeMultiplied)

	// Calculate the slice of base reward per second based on the stake percentage.
	sliceOfBaseRewardPerSecondPercent := new(big.Int).Div(addressTotalStakeMultiplied, sumTotalStake)
	log.Debug("GetAvailablePaybackByAddress:", "sliceOfBaseRewardPerSecondPercent * multiplier(1e10)", sliceOfBaseRewardPerSecondPercent)

	// Calculate full duration for payback calculations
	fullDuration := pc.calculateFullDuration(address, currentEpoch, prevEpoch)
	log.Debug("GetAvailablePaybackByAddress:", "fullDuration", fullDuration)

	// Calculate total available payback based on base reward and duration.
	paybackSum := pc.calculatePayback(multiplier, baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent, fullDuration)

	// Subtract used payback from the total available payback.
	paybackSum = paybackSum.Sub(paybackSum, pc.GetQuotaUsed(address))
	log.Debug("GetAvailablePaybackByAddress:", "paybackSum after subtracting used payback", paybackSum)

	// Return zero if the total payback is negative after adjustments.
	if paybackSum.Cmp(big.NewInt(0)) < 0 {
		log.Warn("GetAvailablePaybackByAddress: paybackSum is negative", "paybackSum", paybackSum)
		return big.NewInt(0)
	}

	// Return the calculated payback.
	return paybackSum
}

// calculatePayback calculates the available quota based on the base reward per second, the percentage of total stake, and the full duration.
func (pc *PaybackCache) calculatePayback(multiplier, baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent *big.Int, fullDuration int64) *big.Int {
	// Calculate actual base reward per second value using the calculated percentage.
	sliceOfBaseRewardPerSecondValueMultiplied := new(big.Int).Mul(baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent)
	sliceOfBaseRewardPerSecondValue := new(big.Int).Div(sliceOfBaseRewardPerSecondValueMultiplied, multiplier)

	// Calculate total available quota based on base reward and duration.
	paybackSum := big.NewInt(0)
	paybackSum = paybackSum.Mul(sliceOfBaseRewardPerSecondValue, big.NewInt(fullDuration))

	return paybackSum
}

// checkAndUpdateContractAddress ensures the contract address is current and updates it if necessary.
func (pc *PaybackCache) checkAndUpdateContractAddress() error {
	if pc.store == nil {
		return errors.New("store is nil")
	}

	if pc.store.GetRules() == (opera.Rules{}) {
		return errors.New("rules are empty")
	}

	newContractAddress := pc.store.GetRules().Economy.QuotaCacheAddress
	if pc.contractAddress == (common.Address{}) || pc.contractAddress != newContractAddress {
		if pc.store.GetRules().Upgrades.Podgorica {
			pc.mu.Lock()
			pc.contractAddress = newContractAddress
			pc.mu.Unlock()
			log.Info("checkAndUpdateContractAddress: QuotaCacheAddress is updated", "address", pc.contractAddress.String())
		} else {
			log.Info("checkAndUpdateContractAddress: HardFork Podgorica not activated", "status", pc.store.GetRules().Upgrades.Podgorica)
			return errors.New("Podgorica upgrade  not activated")
		}
	}

	return nil
}

// calculateStakeDetails computes the total stakes and base reward for an address.
func (pc *PaybackCache) calculateStakeDetails(address common.Address) (*big.Int, *big.Int, error) {
	baseRewardPerSecond, err := pc.getBaseRewardPerSecond(address)
	if err != nil {
		log.Warn("calculateStakeDetails:", "error", err)
		return nil, nil, err
	}

	totalStakeSFC, err := pc.getTotalStake(address, sfcContract.ContractAddress)
	if err != nil {
		log.Warn("calculateStakeDetails:", "error", err)
		return nil, nil, err
	}

	totalStakeQuota, err := pc.getTotalStake(address, pc.contractAddress)
	if err != nil {
		log.Warn("calculateStakeDetails:", "error", err)
		return nil, nil, err
	}

	sumTotalStake := totalStakeSFC.Add(totalStakeSFC, totalStakeQuota)
	if sumTotalStake.Cmp(big.NewInt(0)) == 0 {
		log.Warn("calculateStakeDetails", "sumTotalStake", "is zero")
		return nil, nil, errors.New("sum of total stakes is zero")
	}

	return baseRewardPerSecond, sumTotalStake, nil
}

// calculateFullDuration calculates the effective duration of stakes for quota calculations.
func (pc *PaybackCache) calculateFullDuration(address common.Address, currentEpoch, prevEpoch idx.Epoch) int64 {
	sumStakeByAddress := pc.getSumStakeByAddress(address, prevEpoch, currentEpoch)
	var fullDuration int64

	if sumStakeByAddress.Cmp(big.NewInt(0)) != 0 {
		sumStakeByAddressCurrentEpoch := pc.getSumStakeByAddress(address, currentEpoch, currentEpoch)
		sumStakeByAddressPrevEpoch := pc.getSumStakeByAddress(address, prevEpoch, prevEpoch)

		// Calculate the duration if stakes exist in the current epoch but not in the previous.
		if sumStakeByAddressCurrentEpoch.Cmp(big.NewInt(0)) > 0 && sumStakeByAddressPrevEpoch.Cmp(big.NewInt(0)) == 0 {
			block := pc.store.GetBlock(idx.Block(pc.store.GetLatestBlockIndex()))
			if block == nil {
				log.Warn("calculateFullDuration: latest block not found in store")
				return 0
			}
			lastBlockTime := block.Time.Time()

			pc.mu.RLock()
			stakes := pc.getStakesForEpochLocked(currentEpoch, address)
			pc.mu.RUnlock()

			if len(stakes) == 0 {
				log.Warn("calculateFullDuration: no stakes found for current epoch despite non-zero sum")
				return 0
			}
			lastTimeStake := stakes[len(stakes)-1].Timestamp
			fullDuration = int64(lastBlockTime.Sub(lastTimeStake) / 1e9)
		}
	}

	// Use previous epoch duration as full duration if no current stakes.
	if fullDuration == 0 {
		epochState := pc.store.GetHistoryEpochState(prevEpoch)
		if epochState == nil {
			log.Warn("calculateFullDuration: epoch state not found for prevEpoch", "prevEpoch", prevEpoch)
			return 0
		}
		durationPrevEpochBySec := epochState.Duration() / 1e9
		fullDuration = int64(durationPrevEpochBySec)
	}

	return fullDuration
}

func (pc *PaybackCache) getAddressTotalStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	packedData, err := pc.ContractABI.Pack("getStake", address)
	if err != nil {
		return big.NewInt(0), err
	}

	result, _, vmerr = pc.evm.StaticCall(sender, pc.contractAddress, packedData, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (pc *PaybackCache) getMinStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("minStake()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = pc.evm.StaticCall(sender, pc.contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (pc *PaybackCache) GetStore() Store {
	return pc.store
}

func (pc *PaybackCache) SetEVM(evm *vm.EVM) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.evm = evm
}

func (pc *PaybackCache) getBaseRewardPerSecond(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("baseRewardPerSecond()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = pc.evm.StaticCall(sender, sfcContract.ContractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (pc *PaybackCache) getTotalStake(address common.Address, contractAddress common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("totalStake()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = pc.evm.StaticCall(sender, contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (pc *PaybackCache) getStakesForEpoch(epoch idx.Epoch, address common.Address) []StakeInfo {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.getStakesForEpochLocked(epoch, address)
}

// getStakesForEpochLocked requires pc.mu to be held by the caller.
func (pc *PaybackCache) getStakesForEpochLocked(epoch idx.Epoch, address common.Address) []StakeInfo {
	if epochStakes, ok := pc.StakesMap[epoch]; ok {
		if stakes, ok := epochStakes.StakesByAddress[address]; ok {
			dst := make([]StakeInfo, len(stakes))
			copy(dst, stakes)
			return dst
		}
	}
	return nil
}

func (pc *PaybackCache) getSumStakeByAddress(address common.Address, prevEpoch idx.Epoch, currentEpoch idx.Epoch) *big.Int {
	sum := big.NewInt(0)

	pc.mu.RLock()
	defer pc.mu.RUnlock()

	for i := prevEpoch; i <= currentEpoch; i++ {
		stakes := pc.getStakesForEpochLocked(i, address)
		if stakes != nil {
			for _, stake := range stakes {
				sum.Add(sum, stake.Amount)
			}
		}
	}

	return sum
}

// cleanupOldEpochsLocked removes epoch data older than 2 epochs. Caller must hold pc.mu.
func (pc *PaybackCache) cleanupOldEpochsLocked() {
	currentEpoch := pc.store.GetCurrentEpoch()
	if currentEpoch <= 2 {
		pc.PaybackUsedMap = make(map[common.Address]*big.Int)
		return
	}
	cutoffEpoch := currentEpoch - 2

	for epoch := range pc.StakesMap {
		if epoch < cutoffEpoch {
			delete(pc.StakesMap, epoch)
		}
	}

	pc.PaybackUsedMap = make(map[common.Address]*big.Int)
}
