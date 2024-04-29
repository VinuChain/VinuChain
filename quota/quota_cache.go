package quota

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Fantom-foundation/go-opera/opera"
	sfcContract "github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/quota/contract/quotaProxy"
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

type QuotaCache struct {
	// QuotaUsedMap contains last quota used
	QuotaUsedMap map[common.Address]*big.Int

	// StakesMap contains stakes and unstakes
	StakesMap map[idx.Epoch]*EpochStakes

	// ABI to get names of called methods
	ContractABI *abi.ABI

	store Store

	evm *vm.EVM

	contractAddress common.Address
}

func (qc *QuotaCache) deleteCurrentBlock() error {
	return nil
}

func (qc *QuotaCache) AddTransaction(tx *types.Transaction, receipt *types.Receipt) error {
	if receipt.Status == types.ReceiptStatusSuccessful {
		epoch := qc.store.GetCurrentEpoch()
		txtype := getTxType(tx, *qc.ContractABI)
		if txtype == TxTypeStake {
			stakes := qc.StakesMap[epoch].StakesByAddress[tx.From()]
			stakes = append(stakes, StakeInfo{
				Amount:    tx.Value(),
				Timestamp: qc.store.GetBlock(idx.Block(receipt.BlockNumber.Uint64())).Time.Time(),
			})

			qc.StakesMap[epoch].StakesByAddress[tx.From()] = stakes
		}

		if receipt.FeeRefund != nil {
			if _, ok := qc.QuotaUsedMap[tx.From()]; !ok {
				qc.QuotaUsedMap[tx.From()] = big.NewInt(0)
			}
			qc.QuotaUsedMap[tx.From()].Add(qc.QuotaUsedMap[tx.From()], receipt.FeeRefund)
		}

	}
	return nil
}

func (qc *QuotaCache) GetQuotaUsed(address common.Address) *big.Int {
	if _, ok := qc.QuotaUsedMap[address]; ok {
		return qc.QuotaUsedMap[address]
	}
	return big.NewInt(0)
}

func NewQuotaCache(store Store) *QuotaCache {
	qc := QuotaCache{
		QuotaUsedMap: make(map[common.Address]*big.Int),
		StakesMap:    make(map[idx.Epoch]*EpochStakes),
		store:        store,
	}

	if store.GetRules().Upgrades.Podgorica {
		qc.contractAddress = store.GetRules().Economy.QuotaCacheAddress
		log.Info("NewQuotaCache: QuotaCacheAddress is set", "address", qc.contractAddress.String())
	} else {
		log.Info("HardFork Podgorica is not activated", "status", store.GetRules().Upgrades.Podgorica)
		log.Info("NewQuotaCache:", "Rules", store.GetRules())
	}

	abiQuotaProxy, err := abi.JSON(strings.NewReader(quotaProxy.QuotaProxyABI))
	if err != nil {
		panic(err)
	}
	qc.ContractABI = &abiQuotaProxy

	return &qc
}

// String() prints all data in cache
func (qc *QuotaCache) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("QuotaUsedMap: %v\n", qc.QuotaUsedMap))
	sb.WriteString(fmt.Sprintf("StakesMap: %v\n", qc.StakesMap))
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

// GetAvailableQuotaByAddress calculates the available quota for a given address.
func (qc *QuotaCache) GetAvailableQuotaByAddress(address common.Address) *big.Int {
	// Initialize quota to zero.
	quota := big.NewInt(0)

	// Return zero quota if the address is empty.
	if address == (common.Address{}) {
		return quota
	}

	// Check and update the contract address.
	if err := qc.checkAndUpdateContractAddress(); err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
		return quota
	}

	// Retrieve total stake for the address and handle errors.
	addressTotalStake, err := qc.getAddressTotalStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	// Get minimum stake required and handle errors.
	minStake, err := qc.getMinStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	// Return zero quota if the address's total stake is below the minimum stake.
	if addressTotalStake.Cmp(minStake) < 0 {
		return quota
	}

	// Retrieve the current and previous epochs.
	currentEpoch := qc.store.GetCurrentEpoch()
	prevEpoch := currentEpoch - 1

	// Initialize stakes map for the current epoch if not already done.
	if qc.StakesMap[currentEpoch] == nil {
		qc.StakesMap[currentEpoch] = &EpochStakes{
			StakesByAddress: make(map[common.Address][]StakeInfo),
		}

		// Cleanup old epochs.
		qc.cleanupOldEpochs()
	}

	// Calculate the base reward per second and the total stake from SFC contract and Quota contract.
	baseRewardPerSecond, sumTotalStake, err := qc.calculateStakeDetails(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
		return quota
	}

	log.Info("GetAvailableQuotaByAddress:", "baseRewardPerSecond", baseRewardPerSecond)
	log.Info("GetAvailableQuotaByAddress:", "sumTotalStake", sumTotalStake)

	// Apply a multiplier for calculation precision.
	multiplier := big.NewInt(1e10)
	addressTotalStakeMultiplied := new(big.Int).Mul(addressTotalStake, multiplier)
	log.Info("GetAvailableQuotaByAddress:", "addressTotalStakeMultiplied", addressTotalStakeMultiplied)

	// Calculate the slice of base reward per second based on the stake percentage.
	sliceOfBaseRewardPerSecondPercent := new(big.Int).Div(addressTotalStakeMultiplied, sumTotalStake)
	log.Info("GetAvailableQuotaByAddress:", "sliceOfBaseRewardPerSecondPercent * multiplier(1e10)", sliceOfBaseRewardPerSecondPercent)

	// Calculate full duration for quota calculations
	fullDuration := qc.calculateFullDuration(address, currentEpoch, prevEpoch)
	log.Info("GetAvailableQuotaByAddress:", "fullDuration", fullDuration)

	// Calculate total available quota based on base reward and duration.
	quotaSum := qc.calculateQuota(multiplier, baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent, fullDuration)

	// Subtract used quota from the total available quota.
	quotaSum = quotaSum.Sub(quotaSum, qc.GetQuotaUsed(address))
	log.Info("GetAvailableQuotaByAddress:", "quotaSum after subtracting used quota", quotaSum)

	// Return zero if the total quota is negative after adjustments.
	if quotaSum.Cmp(big.NewInt(0)) < 0 {
		log.Warn("GetAvailableQuotaByAddress: quotaSum is negative", "quotaSum", quotaSum)
		return big.NewInt(0)
	}

	// Return the calculated quota.
	return quotaSum
}

// calculateQuota calculates the available quota based on the base reward per second, the percentage of total stake, and the full duration.
func (qc *QuotaCache) calculateQuota(multiplier, baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent *big.Int, fullDuration int64) *big.Int {
	// Calculate actual base reward per second value using the calculated percentage.
	sliceOfBaseRewardPerSecondValueMultiplied := new(big.Int).Mul(baseRewardPerSecond, sliceOfBaseRewardPerSecondPercent)
	sliceOfBaseRewardPerSecondValue := new(big.Int).Div(sliceOfBaseRewardPerSecondValueMultiplied, multiplier)

	// Calculate total available quota based on base reward and duration.
	quotaSum := big.NewInt(0)
	quotaSum = quotaSum.Mul(sliceOfBaseRewardPerSecondValue, big.NewInt(fullDuration))

	return quotaSum
}

// checkAndUpdateContractAddress ensures the contract address is current and updates it if necessary.
func (qc *QuotaCache) checkAndUpdateContractAddress() error {
	if qc.store == nil {
		return errors.New("store is nil")
	}

	if qc.store.GetRules() == (opera.Rules{}) {
		return errors.New("rules are empty")
	}

	newContractAddress := qc.store.GetRules().Economy.QuotaCacheAddress
	if qc.contractAddress == (common.Address{}) || qc.contractAddress != newContractAddress {
		if qc.store.GetRules().Upgrades.Podgorica {
			qc.contractAddress = newContractAddress
			log.Info("checkAndUpdateContractAddress: QuotaCacheAddress is updated", "address", qc.contractAddress.String())
		} else {
			log.Info("checkAndUpdateContractAddress: HardFork Podgorica not activated", "status", qc.store.GetRules().Upgrades.Podgorica)
			return errors.New("Podgorica upgrade  not activated")
		}
	}

	return nil
}

// calculateStakeDetails computes the total stakes and base reward for an address.
func (qc *QuotaCache) calculateStakeDetails(address common.Address) (*big.Int, *big.Int, error) {
	baseRewardPerSecond, err := qc.getBaseRewardPerSecond(address)
	if err != nil {
		log.Warn("calculateStakeDetails:", "error", err)
		return nil, nil, err
	}

	totalStakeSFC, err := qc.getTotalStake(address, sfcContract.ContractAddress)
	if err != nil {
		log.Warn("calculateStakeDetails:", "error", err)
		return nil, nil, err
	}

	totalStakeQuota, err := qc.getTotalStake(address, qc.contractAddress)
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
func (qc *QuotaCache) calculateFullDuration(address common.Address, currentEpoch, prevEpoch idx.Epoch) int64 {
	sumStakeByAddress := qc.getSumStakeByAddress(address, prevEpoch, currentEpoch)
	var fullDuration int64

	if sumStakeByAddress.Cmp(big.NewInt(0)) != 0 {
		sumStakeByAddressCurrentEpoch := qc.getSumStakeByAddress(address, currentEpoch, currentEpoch)
		sumStakeByAddressPrevEpoch := qc.getSumStakeByAddress(address, prevEpoch, prevEpoch)

		// Calculate the duration if stakes exist in the current epoch but not in the previous.
		if sumStakeByAddressCurrentEpoch.Cmp(big.NewInt(0)) > 0 && sumStakeByAddressPrevEpoch.Cmp(big.NewInt(0)) == 0 {
			lastBlockTime := qc.store.GetBlock(idx.Block(qc.store.GetLatestBlockIndex())).Time.Time()
			stakes := qc.getStakesForEpoch(currentEpoch, address)
			lastTimeStake := stakes[len(stakes)-1].Timestamp
			fullDuration = int64(lastBlockTime.Sub(lastTimeStake) / 1e9)
		}
	}

	// Use previous epoch duration as full duration if no current stakes.
	if fullDuration == 0 {
		durationPrevEpochBySec := qc.store.GetHistoryEpochState(prevEpoch).Duration() / 1e9
		fullDuration = int64(durationPrevEpochBySec)
	}

	return fullDuration
}

func (qc *QuotaCache) getAddressTotalStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	packedData, err := qc.ContractABI.Pack("getStake", address)
	if err != nil {
		return big.NewInt(0), err
	}

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, packedData, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) getMinStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("minStake()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) GetStore() Store {
	return qc.store
}

func (qc *QuotaCache) SetEVM(evm *vm.EVM) {
	qc.evm = evm
}

func (qc *QuotaCache) getBaseRewardPerSecond(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("baseRewardPerSecond()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = qc.evm.StaticCall(sender, sfcContract.ContractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) getTotalStake(address common.Address, contractAddress common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("totalStake()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = qc.evm.StaticCall(sender, contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) getStakesForEpoch(epoch idx.Epoch, address common.Address) []StakeInfo {
	if epochStakes, ok := qc.StakesMap[epoch]; ok {
		if stakes, ok := epochStakes.StakesByAddress[address]; ok {
			return stakes
		}
	}
	return nil
}

func (qc *QuotaCache) getSumStakeByAddress(address common.Address, prevEpoch idx.Epoch, currentEpoch idx.Epoch) *big.Int {
	sum := big.NewInt(0)

	for i := prevEpoch; i <= currentEpoch; i++ {
		stakes := qc.getStakesForEpoch(i, address)
		if stakes != nil {
			for _, stake := range stakes {
				sum.Add(sum, stake.Amount)
			}
		}
	}

	return sum
}

func (qc *QuotaCache) cleanupOldEpochs() {
	currentEpoch := qc.store.GetCurrentEpoch()
	cutoffEpoch := currentEpoch - 2

	for epoch := range qc.StakesMap {
		if epoch < cutoffEpoch {
			delete(qc.StakesMap, epoch)
		}
	}

	qc.QuotaUsedMap = make(map[common.Address]*big.Int)
}
