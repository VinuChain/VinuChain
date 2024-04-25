package quota

import (
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

func (qc *QuotaCache) GetAvailableQuotaByAddress(address common.Address, tx *types.Transaction) *big.Int {
	quota := big.NewInt(0)

	if qc.contractAddress == (common.Address{}) {
		if qc.store != nil {
			if qc.store.GetRules() != (opera.Rules{}) {
				log.Info("GetAvailableQuotaByAddress: HardFork Podgorica", "status", qc.store.GetRules().Upgrades.Podgorica)
			}
		}
	}

	if address == (common.Address{}) {
		return quota
	}

	if qc.contractAddress == (common.Address{}) {
		if qc.store != nil {
			if qc.store.GetRules() != (opera.Rules{}) {
				if qc.store.GetRules().Upgrades.Podgorica {
					qc.contractAddress = qc.store.GetRules().Economy.QuotaCacheAddress
					log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is set", "address", qc.contractAddress.String())
				} else {
					return quota
				}
			} else {
				return quota
			}
		} else {
			log.Warn("GetAvailableQuotaByAddress: store is nil")
			return quota
		}
	} else {
		if qc.store.GetRules().Economy.QuotaCacheAddress != qc.contractAddress {
			qc.contractAddress = qc.store.GetRules().Economy.QuotaCacheAddress
			log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress has been changed", "address", qc.contractAddress.String())
		}
	}

	addressTotalStake, err := qc.getAddressTotalStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	minStake, err := qc.getMinStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	if addressTotalStake.Cmp(minStake) < 0 {
		return quota
	}

	currentEpoch := qc.store.GetCurrentEpoch()
	prevEpoch := currentEpoch - 1

	if qc.StakesMap[currentEpoch] == nil {
		qc.StakesMap[currentEpoch] = &EpochStakes{
			StakesByAddress: make(map[common.Address][]StakeInfo),
		}

		qc.cleanupOldEpochs()
	}

	baseRewardPerSecond, err := qc.getBaseRewardPerSecond(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	log.Info("GetAvailableQuotaByAddress:", "baseRewardPerSecond", baseRewardPerSecond)

	totalStakeSFI, err := qc.getTotalStake(address, sfcContract.ContractAddress)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	log.Info("GetAvailableQuotaByAddress:", "totalStakeSFI", totalStakeSFI)

	totalStakeQuota, err := qc.getTotalStake(address, qc.contractAddress)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	log.Info("GetAvailableQuotaByAddress:", "totalStakeQuota", totalStakeQuota)

	sumTotalStake := totalStakeSFI.Add(totalStakeSFI, totalStakeQuota)
	if sumTotalStake.Cmp(big.NewInt(0)) == 0 {
		log.Warn("GetAvailableQuotaByAddress", "sumTotalStake", "is zero")
		return big.NewInt(0)
	}

	log.Info("GetAvailableQuotaByAddress:", "sumTotalStake", sumTotalStake)

	addressTotalStakeFloat := big.NewFloat(0).SetInt(addressTotalStake)

	log.Info("GetAvailableQuotaByAddress:", "addressTotalStakeFloat", addressTotalStakeFloat)

	sumTotalStakeFloat := big.NewFloat(0).SetInt(sumTotalStake)

	log.Info("GetAvailableQuotaByAddress:", "sumTotalStakeFloat", sumTotalStakeFloat)

	sliceOfBaseRewardPerSecondPercent := big.NewFloat(0).Quo(addressTotalStakeFloat, sumTotalStakeFloat)

	log.Info("GetAvailableQuotaByAddress:", "sliceOfBaseRewardPerSecondPercent", sliceOfBaseRewardPerSecondPercent)

	baseRewardPerSecondFloat := new(big.Float).SetInt(baseRewardPerSecond)
	sliceOfBaseRewardPerSecondValueFloat := new(big.Float).Mul(baseRewardPerSecondFloat, sliceOfBaseRewardPerSecondPercent)
	sliceOfBaseRewardPerSecondValue, _ := sliceOfBaseRewardPerSecondValueFloat.Int(nil)

	log.Info("GetAvailableQuotaByAddress:", "sliceOfBaseRewardPerSecondValue", sliceOfBaseRewardPerSecondValue)

	var fullDuration int64

	sumStakeByAddress := qc.getSumStakeByAddress(address, prevEpoch, currentEpoch)
	if sumStakeByAddress.Cmp(big.NewInt(0)) != 0 {
		sumStakeByAddressCurrentEpoch := qc.getSumStakeByAddress(address, currentEpoch, currentEpoch)
		sumStakeByAddressPrevEpoch := qc.getSumStakeByAddress(address, prevEpoch, prevEpoch)

		switch {
		case sumStakeByAddressCurrentEpoch.Cmp(big.NewInt(0)) > 0 &&
			sumStakeByAddressPrevEpoch.Cmp(big.NewInt(0)) == 0 &&
			addressTotalStake.Cmp(sumStakeByAddressCurrentEpoch) == 0:

			lastBlockTime := qc.store.GetBlock(idx.Block(qc.store.GetLatestBlockIndex())).Time.Time()
			stakes := qc.getStakesForEpoch(currentEpoch, address)
			lastTimeStake := stakes[len(stakes)-1].Timestamp
			fullDuration = int64(lastBlockTime.Sub(lastTimeStake) / 1e9)

		}
	}

	if fullDuration == 0 {
		durationPrevEpochBySec := qc.store.GetHistoryEpochState(prevEpoch).Duration() / 1e9
		fullDuration = int64(durationPrevEpochBySec)
	}

	log.Info("GetAvailableQuotaByAddress:", "fullDuration", fullDuration)

	quotaSum := big.NewInt(0)
	quotaSum = quotaSum.Mul(sliceOfBaseRewardPerSecondValue, big.NewInt(fullDuration))

	log.Info("GetAvailableQuotaByAddress:", "quotaSum", quotaSum)

	quotaSum = quotaSum.Sub(quotaSum, qc.GetQuotaUsed(address))

	log.Info("GetAvailableQuotaByAddress:", "quotaSum after subtracting used quota", quotaSum)

	if quotaSum.Cmp(big.NewInt(0)) < 0 {
		log.Warn("GetAvailableQuotaByAddress: quotaSum is negative", "quotaSum", quotaSum)
		return big.NewInt(0)
	}

	return quotaSum
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
