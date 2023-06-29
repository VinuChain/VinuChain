// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfc100

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"justification\",\"type\":\"string\"}],\"name\":\"InflatedFTM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedUpStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundedSlashedLegacyDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"UnlockedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UpdatedOfflinePenaltyThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTokenizerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockedRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getStakes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSFC.Stake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getWrRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSFC.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeDriver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offlinePenaltyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSecond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"updateOfflinePenaltyThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"updateStakeTokenizerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"diff\",\"type\":\"int256\"}],\"name\":\"updateTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"justification\",\"type\":\"string\"}],\"name\":\"mintFTM\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTime\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getUnlockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"relockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b5061609f80620000216000396000f3fe6080604052600436106104315760003560e01c8063854873e111610229578063b88a37e21161012e578063d5a44f86116100b6578063e08d7e661161007a578063e08d7e6614610c80578063e261641a14610ca0578063e2f8c33614610cc0578063ebdf104c14610ce0578063f2fde38b14610d0057610431565b8063d5a44f8614610bdb578063d9a7c1f914610c0b578063dc31e1af14610c20578063de67f21514610c40578063df00c92214610c6057610431565b8063c65ee0e1116100fd578063c65ee0e114610b46578063c7be95de14610b66578063cc8343aa14610b7b578063cfd4766314610b9b578063cfdbb7cd14610bbb57610431565b8063b88a37e214610ac4578063bd14d90714610af1578063c3de580e14610b11578063c5f530af14610b3157610431565b8063a198d229116101b1578063a86a056f11610180578063a86a056f14610a1c578063b5d8962714610a3c578063b6d9edd514610a6f578063b810e41114610a8f578063b82b842714610aaf57610431565b8063a198d229146109b4578063a2f6e6bc146109d4578063a5a470ad146109f4578063a778651514610a0757610431565b80638cddb015116101f85780638cddb0151461091a5780638da5cb5b1461093a5780638f32d59b1461094f57806396c7ee46146109715780639fa6dd35146109a157610431565b8063854873e11461088b5780638b0e9f3f146108b85780638b1a0d11146108cd5780638c3c51d8146108ed57610431565b806339b80c001161033a5780636099ecb2116102c25780636f498663116102865780636f498663146107ff578063702797e31461081f578063715018a61461084c57806376671808146108615780637cacb1d61461087657610431565b80636099ecb21461076a57806361e53fcc1461078a578063634b91e3146107aa578063650acd66146107ca578063670322f8146107df57610431565b806354fd4d501161030957806354fd4d50146106f35780635601fe011461071557806358f95b80146107355780635e2308d2146105f35780635fab23a81461075557610431565b806339b80c0014610660578063441a3e70146106935780634f7c4efb146106b35780634feb92f3146106d357610431565b806318f628d4116103bd5780632265f2841161038c5780632265f284146105de5780632709275e146105f357806328f73148146106085780632cedb0971461061d578063346bdcfb1461064057610431565b806318f628d41461054f5780631d3ac42c1461056f5780631e702f831461058f5780631f270152146105af57610431565b80630d4955e3116104045780630d4955e3146104ce5780630d7b2609146104e35780630e559d82146104f857806312622d0e1461051a57806318160ddd1461053a57610431565b80630135b1db14610436578063019e27291461046c57806308c368741461048e5780630962ef79146104ae575b600080fd5b34801561044257600080fd5b50610456610451366004614c02565b610d20565b6040516104639190615e55565b60405180910390f35b34801561047857600080fd5b5061048c6104873660046150c5565b610d32565b005b34801561049a57600080fd5b5061048c6104a9366004615058565b610f0e565b3480156104ba57600080fd5b5061048c6104c9366004615058565b610fd7565b3480156104da57600080fd5b506104566110dd565b3480156104ef57600080fd5b506104566110e6565b34801561050457600080fd5b5061050d6110ed565b6040516104639190615b77565b34801561052657600080fd5b50610456610535366004614c88565b6110fc565b34801561054657600080fd5b50610456611185565b34801561055b57600080fd5b5061048c61056a366004614e3c565b61118b565b34801561057b57600080fd5b5061045661058a3660046150a6565b6112af565b34801561059b57600080fd5b5061048c6105aa3660046150a6565b6113d4565b3480156105bb57600080fd5b506105cf6105ca366004614d8e565b61142f565b60405161046393929190615eb2565b3480156105ea57600080fd5b50610456611461565b3480156105ff57600080fd5b50610456611473565b34801561061457600080fd5b5061045661148f565b34801561062957600080fd5b50610632611495565b604051610463929190615ea4565b34801561064c57600080fd5b5061048c61065b366004615058565b61149f565b34801561066c57600080fd5b5061068061067b366004615058565b6114e9565b6040516104639796959493929190615f50565b34801561069f57600080fd5b5061048c6106ae3660046150a6565b61152b565b3480156106bf57600080fd5b5061048c6106ce3660046150a6565b61181d565b3480156106df57600080fd5b5061048c6106ee366004614cc2565b6118dc565b3480156106ff57600080fd5b50610708611963565b6040516104639190615c16565b34801561072157600080fd5b50610456610730366004615058565b61196d565b34801561074157600080fd5b506104566107503660046150a6565b6119a3565b34801561076157600080fd5b506104566119c0565b34801561077657600080fd5b50610456610785366004614c88565b6119c6565b34801561079657600080fd5b506104566107a53660046150a6565b611a04565b3480156107b657600080fd5b5061048c6107c53660046150a6565b611a25565b3480156107d657600080fd5b50610456611bca565b3480156107eb57600080fd5b506104566107fa366004614c88565b611bcf565b34801561080b57600080fd5b5061045661081a366004614c88565b611c10565b34801561082b57600080fd5b5061083f61083a366004614ddb565b611c7a565b6040516104639190615be6565b34801561085857600080fd5b5061048c611d41565b34801561086d57600080fd5b50610456611daf565b34801561088257600080fd5b50610456611db8565b34801561089757600080fd5b506108ab6108a6366004615058565b611dbe565b6040516104639190615c24565b3480156108c457600080fd5b50610456611e59565b3480156108d957600080fd5b5061048c6108e83660046150a6565b611e5f565b3480156108f957600080fd5b5061090d6109083660046150a6565b611ecb565b6040516104639190615bd5565b34801561092657600080fd5b5061048c610935366004614c88565b611fac565b34801561094657600080fd5b5061050d611fd2565b34801561095b57600080fd5b50610964611fe1565b6040516104639190615c08565b34801561097d57600080fd5b5061099161098c366004614c88565b611ff2565b6040516104639493929190615eda565b61048c6109af366004615058565b612024565b3480156109c057600080fd5b506104566109cf3660046150a6565b61202f565b3480156109e057600080fd5b5061048c6109ef366004614c02565b612050565b61048c610a02366004615022565b612096565b348015610a1357600080fd5b50610456612127565b348015610a2857600080fd5b50610456610a37366004614c88565b61213d565b348015610a4857600080fd5b50610a5c610a57366004615058565b61215a565b6040516104639796959493929190615ee8565b348015610a7b57600080fd5b5061048c610a8a366004615058565b6121a0565b348015610a9b57600080fd5b506105cf610aaa366004614c88565b61222d565b348015610abb57600080fd5b50610456612259565b348015610ad057600080fd5b50610ae4610adf366004615058565b612260565b6040516104639190615bf7565b348015610afd57600080fd5b5061048c610b0c36600461511a565b6122c5565b348015610b1d57600080fd5b50610964610b2c366004615058565b6122d8565b348015610b3d57600080fd5b506104566122ef565b348015610b5257600080fd5b50610456610b61366004615058565b6122fd565b348015610b7257600080fd5b5061045661230f565b348015610b8757600080fd5b5061048c610b96366004615076565b612315565b348015610ba757600080fd5b50610456610bb6366004614c88565b612445565b348015610bc757600080fd5b50610964610bd6366004614c88565b612462565b348015610be757600080fd5b50610bfb610bf6366004615058565b6124f8565b6040516104639493929190615ba0565b348015610c1757600080fd5b50610456612539565b348015610c2c57600080fd5b50610456610c3b3660046150a6565b61253f565b348015610c4c57600080fd5b5061048c610c5b36600461511a565b612560565b348015610c6c57600080fd5b50610456610c7b3660046150a6565b6125b1565b348015610c8c57600080fd5b5061048c610c9b366004614ef0565b6125d2565b348015610cac57600080fd5b50610456610cbb3660046150a6565b61268d565b348015610ccc57600080fd5b5061048c610cdb366004614c20565b6126ae565b348015610cec57600080fd5b5061048c610cfb366004614f32565b61275d565b348015610d0c57600080fd5b5061048c610d1b366004614c02565b61291a565b60696020526000908152604090205481565b600054610100900460ff1680610d4b5750610d4b612947565b80610d59575060005460ff16155b610d7e5760405162461bcd60e51b8152600401610d7590615d55565b60405180910390fd5b600054610100900460ff16158015610da9576000805460ff1961ff0019909116610100171660011790555b610db28261294d565b6067859055606680546001600160a01b0319166001600160a01b03851617905560768490556755cfe697852e904c6075556103e86078556203f480607955610df8612a1f565b600086815260776020908152604080832060070193909355825160808101845282815290810182815292810182815260608201838152607c8054600181018255945291517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd0600490940293840180546001600160a01b0319166001600160a01b0390921691909117905592517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd183015591517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd282015590517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd3909101558015610f07576000805461ff00191690555b5050505050565b33610f176149f0565b610f218284612a23565b60208101518151919250600091610f3d9163ffffffff612af316565b9050610f608385610f5b856040015185612af390919063ffffffff16565b612b18565b6001600160a01b03831660008181526073602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610fc9939092909190615eb2565b60405180910390a350505050565b33610fe06149f0565b610fea8284612a23565b90506000826001600160a01b0316611027836040015161101b85602001518660000151612af390919063ffffffff16565b9063ffffffff612af316565b60405161103390615b6c565b60006040518083038185875af1925050503d8060008114611070576040519150601f19603f3d011682016040523d82523d6000602084013e611075565b606091505b50509050806110965760405162461bcd60e51b8152600401610d7590615d95565b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610fc993615eb2565b6301e133805b90565b6212750090565b607b546001600160a01b031681565b60006111088383612462565b61113657506001600160a01b038216600090815260726020908152604080832084845290915290205461117f565b6001600160a01b03831660008181526073602090815260408083208684528252808320549383526072825280832086845290915290205461117c9163ffffffff612b9916565b90505b92915050565b60765481565b61119433612bdb565b6111b05760405162461bcd60e51b8152600401610d7590615c85565b6111bb898989612bef565b6001600160a01b0389166000908152606f602090815260408083208b845290915290206002018190556111ed87612edd565b85156112a457868611156112135760405162461bcd60e51b8152600401610d7590615e45565b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e789061129a9088908c90615ea4565b60405180910390a3505b505050505050505050565b3360008181526073602090815260408083208684529091528120909190836112e95760405162461bcd60e51b8152600401610d7590615c45565b6112f38286612462565b61130f5760405162461bcd60e51b8152600401610d7590615cf5565b80548411156113305760405162461bcd60e51b8152600401610d7590615dc5565b61133a8286612f5b565b6113565760405162461bcd60e51b8152600401610d7590615dd5565b6113608286612ff8565b50600061137383878785600001546131c3565b8254869003835590506113878387836132f4565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b987846040516113c3929190615ea4565b60405180910390a395945050505050565b6113dd33612bdb565b6113f95760405162461bcd60e51b8152600401610d7590615c85565b806114165760405162461bcd60e51b8152600401610d7590615cc5565b61142082826134a0565b61142b826000612315565b5050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b600061146b6135bf565b601002905090565b6000606461147f6135bf565b601e028161148957fe5b04905090565b606d5481565b6078546079549091565b6114a7611fe1565b6114c35760405162461bcd60e51b8152600401610d7590615d45565b600081126114d85760768054820190556114e6565b607680546000839003900390555b50565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b336115346149f0565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526115a35760405162461bcd60e51b8152600401610d7590615da5565b6115ad8285612f5b565b6115c95760405162461bcd60e51b8152600401610d7590615dd5565b60208082015182516000878152606890935260409092206001015490919015801590611605575060008681526068602052604090206001015482115b15611626575050600084815260686020526040902060018101546002909101545b61162e612259565b8201611638612a1f565b10156116565760405162461bcd60e51b8152600401610d7590615cd5565b61165e611bca565b8101611668611daf565b10156116865760405162461bcd60e51b8152600401610d7590615e15565b6001600160a01b03841660009081526071602090815260408083208984528252808320888452909152812060020154906116bf886122d8565b905060006116e18383607a60008d8152602001908152602001600020546135cb565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116117435760405162461bcd60e51b8152600401610d7590615e35565b60006001600160a01b03881661175f858463ffffffff612b9916565b60405161176b90615b6c565b60006040518083038185875af1925050503d80600081146117a8576040519150601f19603f3d011682016040523d82523d6000602084013e6117ad565b606091505b50509050806117ce5760405162461bcd60e51b8152600401610d7590615d95565b888a896001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21876040516118099190615e55565b60405180910390a450505050505050505050565b611825611fe1565b6118415760405162461bcd60e51b8152600401610d7590615d45565b61184a826122d8565b6118665760405162461bcd60e51b8152600401610d7590615c35565b61186e6135bf565b81111561188d5760405162461bcd60e51b8152600401610d7590615d65565b6000828152607a6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917906118d0908490615e55565b60405180910390a25050565b6118e533612bdb565b6119015760405162461bcd60e51b8152600401610d7590615c85565b611949898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a9150899050888861362d565b606b548811156112a457606b889055505050505050505050565b620ccc0d60ea1b90565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b606e5481565b60006119d06149f0565b6119da84846137a3565b8051602082015160408301519293506119fc9261101b9163ffffffff612af316565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b33611a308184612ff8565b5060008211611a515760405162461bcd60e51b8152600401610d7590615c45565b611a5b81846110fc565b821115611a7a5760405162461bcd60e51b8152600401610d7590615d75565b611a848184612f5b565b611aa05760405162461bcd60e51b8152600401610d7590615dd5565b6001600160a01b0381166000908152607e602090815260408083208684529091529020805460018101909155611ad78285856132f4565b6001600160a01b038216600090815260716020908152604080832087845282528083208484529091529020600201839055611b10611daf565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152902055611b44612a1f565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152812060010191909155611b81908590612315565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d5786604051611bbc9190615e55565b60405180910390a450505050565b600390565b6000611bdb8383612462565b611be75750600061117f565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b6000611c1a6149f0565b506001600160a01b0383166000908152606f60209081526040808320858452825291829020825160608101845281548082526001830154938201849052600290920154938101849052926119fc92909161101b919063ffffffff612af316565b60608082604051908082528060200260200182016040528015611cb757816020015b611ca46149f0565b815260200190600190039081611c9c5790505b50905060005b83811015611d37576001600160a01b03871660009081526071602090815260408083208984528252808320888501845282529182902082516060810184528154815260018201549281019290925260020154918101919091528251839083908110611d2457fe5b6020908102919091010152600101611cbd565b5095945050505050565b611d49611fe1565b611d655760405162461bcd60e51b8152600401610d7590615d45565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b60675460010190565b60675481565b606a6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015611e515780601f10611e2657610100808354040283529160200191611e51565b820191906000526020600020905b815481529060010190602001808311611e3457829003601f168201915b505050505081565b606c5481565b611e67611fe1565b611e835760405162461bcd60e51b8152600401610d7590615d45565b607981905560788290556040517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a3490611ebf9084908490615ea4565b60405180910390a15050565b607c54604080518381526020808502820101909152606091908290848015611f0d57816020015b611efa614a11565b815260200190600190039081611ef25790505b50905060005b84811015611fa3578281870110611f2957611fa3565b607c81870181548110611f3857fe5b600091825260209182902060408051608081018252600490930290910180546001600160a01b0316835260018101549383019390935260028301549082015260039091015460608201528251839083908110611f9057fe5b6020908102919091010152600101611f13565b50949350505050565b611fb68282612ff8565b61142b5760405162461bcd60e51b8152600401610d7590615c95565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6114e6338234612b18565b60009182526077602090815260408084209284526005909201905290205490565b612058611fe1565b6120745760405162461bcd60e51b8152600401610d7590615d45565b607b80546001600160a01b0319166001600160a01b0392909216919091179055565b61209e6122ef565b3410156120bd5760405162461bcd60e51b8152600401610d7590615e25565b806120da5760405162461bcd60e51b8152600401610d7590615d85565b61211a3383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061381192505050565b61142b33606b5434612b18565b600060646121336135bf565b600f028161148957fe5b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b6121a8611fe1565b6121c45760405162461bcd60e51b8152600401610d7590615d45565b6801c985c8903591eb208111156121ed5760405162461bcd60e51b8152600401610d7590615d05565b60758190556040517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae239690612222908390615e55565b60405180910390a150565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b62093a8090565b6000818152607760209081526040918290206006018054835181840281018401909452808452606093928301828280156122b957602002820191906000526020600020905b8154815260200190600101908083116122a5575b50505050509050919050565b336122d28185858561383c565b50505050565b600090815260686020526040902054608016151590565b6969e10de76676d080000090565b607a6020526000908152604090205481565b606b5481565b61231e82613a19565b61233a5760405162461bcd60e51b8152600401610d7590615e05565b60008281526068602052604090206003810154905415612358575060005b60665460405163520337df60e11b81526001600160a01b039091169063a4066fbe9061238a9086908590600401615ea4565b600060405180830381600087803b1580156123a457600080fd5b505af11580156123b8573d6000803e3d6000fd5b505050508180156123c857508015155b15612440576066546000848152606a602052604090819020905163242a6e3f60e01b81526001600160a01b039092169163242a6e3f9161240d91879190600401615e63565b600060405180830381600087803b15801561242757600080fd5b505af115801561243b573d6000803e3d6000fd5b505050505b505050565b607260209081526000928352604080842090915290825290205481565b6001600160a01b0382166000908152607360209081526040808320848452909152812060020154158015906124b957506001600160a01b038316600090815260736020908152604080832085845290915290205415155b801561117c57506001600160a01b03831660009081526073602090815260408083208584529091529020600201546124ef612a1f565b11159392505050565b607c818154811061250557fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b039092169350919084565b60755481565b60009182526077602090815260408084209284526003909201905290205490565b338161257e5760405162461bcd60e51b8152600401610d7590615c45565b6125888185612462565b156125a55760405162461bcd60e51b8152600401610d7590615c65565b6122d28185858561383c565b60009182526077602090815260408084209284526002909201905290205490565b6125db33612bdb565b6125f75760405162461bcd60e51b8152600401610d7590615c85565b600060776000612605611daf565b8152602001908152602001600020905060008090505b8281101561267e57600084848381811061263157fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c86015490935061266f91508263ffffffff612af316565b600c850155505060010161261b565b506122d2600682018484614a42565b60009182526077602090815260408084209284526004909201905290205490565b6126b6611fe1565b6126d25760405162461bcd60e51b8152600401610d7590615d45565b6126db83612edd565b6040516001600160a01b0385169084156108fc029085906000818181858888f19350505050158015612711573d6000803e3d6000fd5b50836001600160a01b03167f9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad84848460405161274f93929190615e83565b60405180910390a250505050565b61276633612bdb565b6127825760405162461bcd60e51b8152600401610d7590615c85565b600060776000612790611daf565b815260200190815260200160002090506060816006018054806020026020016040519081016040528092919081815260200182805480156127f057602002820191906000526020600020905b8154815260200190600101908083116127dc575b5050505050905061287782828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613a3092505050565b6128e6828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a918291850190849080828437600092019190915250613b3f92505050565b6128ee611daf565b6067556128f9612a1f565b600783015550607554600b820155607654600d909101555050505050505050565b612922611fe1565b61293e5760405162461bcd60e51b8152600401610d7590615d45565b6114e681614185565b303b1590565b600054610100900460ff16806129665750612966612947565b80612974575060005460ff16155b6129905760405162461bcd60e51b8152600401610d7590615d55565b600054610100900460ff161580156129bb576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3801561142b576000805461ff00191690555050565b4290565b612a2b6149f0565b612a358383612ff8565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392612a939261101b9163ffffffff612af316565b905080612ab25760405162461bcd60e51b8152600401610d7590615d25565b6001600160a01b0384166000908152606f6020908152604080832086845290915281208181556001810182905560020155612aec81612edd565b5092915050565b60008282018381101561117c5760405162461bcd60e51b8152600401610d7590615c75565b612b2182613a19565b612b3d5760405162461bcd60e51b8152600401610d7590615e05565b60008281526068602052604090205415612b695760405162461bcd60e51b8152600401610d7590615cb5565b612b74838383612bef565b612b7d82614207565b6124405760405162461bcd60e51b8152600401610d7590615df5565b600061117c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061424f565b6066546001600160a01b0390811691161490565b60008111612c0f5760405162461bcd60e51b8152600401610d7590615c45565b612c198383612ff8565b506001600160a01b0383166000908152607d6020908152604080832085845290915290205480612d4c57607c80546001600160a01b038681166000818152607d602090815260408083208a84528252808320869055805160808101825293845290830189815290830188815242606085019081526001870188559690925291517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd0600490950294850180546001600160a01b0319169190941617909255517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd1830155517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd282015590517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd390910155612dc3565b612d7d82607c8381548110612d5d57fe5b906000526020600020906004020160020154612af390919063ffffffff16565b607c8281548110612d8a57fe5b90600052602060002090600402016002018190555042607c8281548110612dad57fe5b9060005260206000209060040201600301819055505b6001600160a01b0384166000908152607260209081526040808320868452909152902054612df7908363ffffffff612af316565b6001600160a01b0385166000908152607260209081526040808320878452825280832093909355606890522060030154612e37818463ffffffff612af316565b600085815260686020526040902060030155606c54612e5c908463ffffffff612af316565b606c55600084815260686020526040902054612e8957606d54612e85908463ffffffff612af316565b606d555b612e94848215612315565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb85604051612ece9190615e55565b60405180910390a35050505050565b6066546040516366e7ea0f60e01b81526001600160a01b03909116906366e7ea0f90612f0f9030908590600401615b85565b600060405180830381600087803b158015612f2957600080fd5b505af1158015612f3d573d6000803e3d6000fd5b5050607654612f55925090508263ffffffff612af316565b60765550565b607b546000906001600160a01b0316612f765750600161117f565b607b546040516321d585c360e01b81526001600160a01b03909116906321d585c390612fa89086908690600401615b85565b60206040518083038186803b158015612fc057600080fd5b505afa158015612fd4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061117c9190810190615004565b60006130026149f0565b61300c848461427b565b9050613017836143b3565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f82528281208682528252829020825160608101845281548152600182015492810192909252600201549181019190915261307d908261440e565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260748252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526130f7908261440e565b6001600160a01b03851660009081526074602090815260408083208784528252918290208351815590830151600182015591015160029091015561313b8484612462565b61319e576001600160a01b0384166000818152607360209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260748352818420888552909252822082815592830182905591909101555b60208101511515806131b05750805115155b806119fc57506040015115159392505050565b6001600160a01b0384166000908152607460209081526040808320868452909152812054819061320b9084906131ff908763ffffffff61448016565b9063ffffffff6144ba16565b6001600160a01b03871660009081526074602090815260408083208984529091528120600101549192509061324c9085906131ff908863ffffffff61448016565b6001600160a01b03881660009081526074602090815260408083208a8452909152902054909150600282048301906132849084612b99565b6001600160a01b03891660009081526074602090815260408083208b84529091529020908155600101546132b89083612b99565b6001600160a01b03891660009081526074602090815260408083208b84529091529020600101558581106132e95750845b979650505050505050565b6001600160a01b0383166000908152607d60209081526040808320858452909152902054607c805461334c9184918490811061332c57fe5b906000526020600020906004020160020154612b9990919063ffffffff16565b607c828154811061335957fe5b906000526020600020906004020160020181905550607c818154811061337b57fe5b9060005260206000209060040201600201546000141561339e5761339e816144fc565b6001600160a01b038416600090815260726020908152604080832086845282528083208054869003905560689091529020600301546133e3908363ffffffff612b9916565b600084815260686020526040902060030155606c54613408908363ffffffff612b9916565b606c5560008381526068602052604090205461343557606d54613431908363ffffffff612b9916565b606d555b60006134408461196d565b90508015613499576134506122ef565b81101561346f5760405162461bcd60e51b8152600401610d7590615e25565b61347884614207565b6134945760405162461bcd60e51b8152600401610d7590615df5565b610f07565b610f078460015b6000828152606860205260409020541580156134bb57508015155b156134e857600082815260686020526040902060030154606d546134e49163ffffffff612b9916565b606d555b60008281526068602052604090205481111561142b5760008281526068602052604090208181556002015461358f5761351f611daf565b600083815260686020526040902060020155613539612a1f565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e479261358692909190615ea4565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516118d09190615e55565b670de0b6b3a764000090565b60008215806135e157506135dd6135bf565b8210155b156135ee57506000613626565b613619600161101b6135fe6135bf565b6131ff8661360a6135bf565b8a91900363ffffffff61448016565b9050838111156136265750825b9392505050565b6001600160a01b038816600090815260696020526040902054156136635760405162461bcd60e51b8152600401610d7590615ce5565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180546001600160a01b031916909417909355606a815291902087516136d492890190614a8d565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051613711929190615ea4565b60405180910390a3811561375a57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613751929190615ea4565b60405180910390a25b841561379957867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e866040516137909190615e55565b60405180910390a25b5050505050505050565b6137ab6149f0565b6137b36149f0565b6137bd848461427b565b6001600160a01b0385166000908152606f6020908152604080832087845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506119fc908261440e565b606b805460010190819055612440838284600061382c611daf565b613834612a1f565b60008061362d565b61384684846110fc565b8111156138655760405162461bcd60e51b8152600401610d7590615de5565b600083815260686020526040902054156138915760405162461bcd60e51b8152600401610d7590615cb5565b6138996110e6565b82101580156138af57506138ab6110dd565b8211155b6138cb5760405162461bcd60e51b8152600401610d7590615ca5565b60006138d98361101b612a1f565b6000858152606860205260409020600601549091506001600160a01b039081169086168114613948576001600160a01b03811660009081526073602090815260408083208884529091529020600201548211156139485760405162461bcd60e51b8152600401610d7590615db5565b6139528686612ff8565b506001600160a01b03861660009081526073602090815260408083208884529091529020600381015485101561399a5760405162461bcd60e51b8152600401610d7590615d35565b80546139ac908563ffffffff612af316565b81556139b6611daf565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7890613a089089908990615ea4565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b60005b8351811015610f0757607854828281518110613a4b57fe5b6020026020010151118015613a755750607954838281518110613a6a57fe5b602002602001015110155b15613ab657613a98848281518110613a8957fe5b602002602001015160086134a0565b613ab6848281518110613aa757fe5b60200260200101516000612315565b828181518110613ac257fe5b6020026020010151856004016000868481518110613adc57fe5b6020026020010151815260200190815260200160002081905550818181518110613b0257fe5b6020026020010151856005016000868481518110613b1c57fe5b602090810291909101810151825281019190915260400160002055600101613a33565b613b47614afb565b6040518060c001604052808551604051908082528060200260200182016040528015613b7d578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015613bb5578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060776000613bef6001613be3611daf565b9063ffffffff612b9916565b81526020810191909152604001600020600160808401526007810154909150613c16612a1f565b1115613c30578060070154613c29612a1f565b0360808301525b60005b8551811015613d38576000826003016000888481518110613c5057fe5b60200260200101518152602001908152602001600020549050600080905081868481518110613c7b57fe5b60200260200101511115613ca25781868481518110613c9657fe5b60200260200101510390505b8460800151878481518110613cb357fe5b6020026020010151820281613cc457fe5b0485604001518481518110613cd557fe5b602002602001018181525050613d0f85604001518481518110613cf457fe5b60200260200101518660600151612af390919063ffffffff16565b606086015260a0850151613d29908263ffffffff612af316565b60a08601525050600101613c33565b5060005b8551811015613e09578260800151858281518110613d5657fe5b60200260200101518460800151878481518110613d6f57fe5b60200260200101518a60000160008b8781518110613d8957fe5b60200260200101518152602001908152602001600020540281613da857fe5b040281613db157fe5b0483600001518281518110613dc257fe5b602002602001018181525050613dfc83600001518281518110613de157fe5b60200260200101518460200151612af390919063ffffffff16565b6020840152600101613d3c565b5060005b855181101561415d576000613e45846080015160755486600001518581518110613e3357fe5b60200260200101518760200151614650565b9050613e81613e748560a0015186604001518581518110613e6257fe5b60200260200101518760600151614691565b829063ffffffff612af316565b90506000878381518110613e9157fe5b6020908102919091018101516000818152606890925260408220600601549092506001600160a01b031690613ecd84613ec8612127565b6146ee565b6001600160a01b0383166000908152607260209081526040808320878452909152902054909150801561407457600081613f078587611bcf565b840281613f1057fe5b049050808303613f1e6149f0565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154613f5090849061470b565b9050613f5a6149f0565b613f6583600061470b565b6001600160a01b0388166000908152606f602090815260408083208c84528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150613fba9083836147fc565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d8452825291829020825160608101845281548152948101549185019190915290910154908201526140359083836147fc565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156140a657816140996135bf565b8402816140a257fe5b0490505b808a600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508b89815181106140e357fe5b60200260200101518f6003016000898152602001908152602001600020819055508c898151811061411057fe5b60200260200101518a600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050613e0d565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b0381166141ab5760405162461bcd60e51b8152600401610d7590615c55565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b60006142346142146135bf565b6131ff61421f611461565b6142288661196d565b9063ffffffff61448016565b60008381526068602052604090206003015411159050919050565b600081848411156142735760405162461bcd60e51b8152600401610d759190615c24565b505050900390565b6142836149f0565b6001600160a01b0383166000908152607060209081526040808320858452909152812054906142b1846143b3565b905060006142bf8686614817565b9050818111156142cc5750805b828110156142d75750815b6001600160a01b0386166000818152607360209081526040808320898452825280832093835260728252808320898452909152812054825490919061432390839063ffffffff612b9916565b9050600061433784600001548a89886148d6565b90506143416149f0565b61434f82866003015461470b565b905061435d838b8a896148d6565b91506143676149f0565b61437283600061470b565b9050614380858c898b6148d6565b925061438a6149f0565b61439584600061470b565b90506143a28383836147fc565b9d9c50505050505050505050505050565b600081815260686020526040812060020154156144065760008281526068602052604090206002015460675410156143ee575060675461199e565b5060008181526068602052604090206002015461199e565b505060675490565b6144166149f0565b6040805160608101909152825184518291614437919063ffffffff612af316565b815260200161445784602001518660200151612af390919063ffffffff16565b815260200161447784604001518660400151612af390919063ffffffff16565b90529392505050565b60008261448f5750600061117f565b8282028284828161449c57fe5b041461117c5760405162461bcd60e51b8152600401610d7590615d15565b600061117c83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250614939565b607c5460009061451390600163ffffffff612b9916565b905080821461460a57607c818154811061452957fe5b9060005260206000209060040201607c838154811061454457fe5b60009182526020909120825460049092020180546001600160a01b0319166001600160a01b039092169190911781556001808301549082015560028083015490820155600391820154910155614598614a11565b607c82815481106145a557fe5b6000918252602080832060408051608081018252600490940290910180546001600160a01b03168085526001820154858501908152600283015486850152600390920154606090950194909452928452607d8252808420925184529190529020839055505b607c80548061461557fe5b60008281526020812060046000199093019283020180546001600160a01b031916815560018101829055600281018290556003015590555050565b60008261465f575060006119fc565b6000614671868663ffffffff61448016565b9050614687836131ff838763ffffffff61448016565b9695505050505050565b6000826146a057506000613626565b60006146b6836131ff878763ffffffff61448016565b90506146e56146c36135bf565b6131ff6146ce611473565b6146d66135bf565b8591900363ffffffff61448016565b95945050505050565b600061117c6146fb6135bf565b6131ff858563ffffffff61448016565b6147136149f0565b6040518060600160405280600081526020016000815260200160008152509050816000146147ce576000614745611473565b61474d6135bf565b039050600061476d61475d6110dd565b6131ff848763ffffffff61448016565b9050600061479661477c6135bf565b6131ff84614788611473565b8a910163ffffffff61448016565b90506147bb6147a36135bf565b6131ff6147ae611473565b899063ffffffff61448016565b60208501819052900383525061117f9050565b6147f16147d96135bf565b6131ff6147e4611473565b869063ffffffff61448016565b604082015292915050565b6148046149f0565b6119fc614811858561440e565b8361440e565b6001600160a01b038216600090815260736020908152604080832084845290915281206001015460675461484c858583614970565b1561485a57915061117f9050565b614865858584614970565b6148745760009250505061117f565b808211156148875760009250505061117f565b808210156148ba576002818301046148a0868683614970565b156148b0578060010192506148b4565b8091505b50614887565b806148ca5760009250505061117f565b60001901949350505050565b60008183106148e7575060006119fc565b6000838152607760208181526040808420888552600190810183528185205487865293835281852089865201909152909120546132e96149256135bf565b6131ff89614228858763ffffffff612b9916565b6000818361495a5760405162461bcd60e51b8152600401610d759190615c24565b50600083858161496657fe5b0495945050505050565b6001600160a01b038316600090815260736020908152604080832085845290915281206001015482108015906119fc57506001600160a01b03841660009081526073602090815260408083208684529091529020600201546149d1836149db565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b604051806080016040528060006001600160a01b031681526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215614a7d579160200282015b82811115614a7d578235825591602001919060010190614a62565b50614a89929150614b31565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614ace57805160ff1916838001178555614a7d565b82800160010185558215614a7d579182015b82811115614a7d578251825591602001919060010190614ae0565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b6110e391905b80821115614a895760008155600101614b37565b803561117f81616036565b60008083601f840112614b6857600080fd5b50813567ffffffffffffffff811115614b8057600080fd5b602083019150836020820283011115614b9857600080fd5b9250929050565b803561117f8161604a565b805161117f8161604a565b60008083601f840112614bc757600080fd5b50813567ffffffffffffffff811115614bdf57600080fd5b602083019150836001820283011115614b9857600080fd5b803561117f81616053565b600060208284031215614c1457600080fd5b60006119fc8484614b4b565b60008060008060608587031215614c3657600080fd5b6000614c428787614b4b565b9450506020614c5387828801614bf7565b935050604085013567ffffffffffffffff811115614c7057600080fd5b614c7c87828801614bb5565b95989497509550505050565b60008060408385031215614c9b57600080fd5b6000614ca78585614b4b565b9250506020614cb885828601614bf7565b9150509250929050565b60008060008060008060008060006101008a8c031215614ce157600080fd5b6000614ced8c8c614b4b565b9950506020614cfe8c828d01614bf7565b98505060408a013567ffffffffffffffff811115614d1b57600080fd5b614d278c828d01614bb5565b97509750506060614d3a8c828d01614bf7565b9550506080614d4b8c828d01614bf7565b94505060a0614d5c8c828d01614bf7565b93505060c0614d6d8c828d01614bf7565b92505060e0614d7e8c828d01614bf7565b9150509295985092959850929598565b600080600060608486031215614da357600080fd5b6000614daf8686614b4b565b9350506020614dc086828701614bf7565b9250506040614dd186828701614bf7565b9150509250925092565b60008060008060808587031215614df157600080fd5b6000614dfd8787614b4b565b9450506020614e0e87828801614bf7565b9350506040614e1f87828801614bf7565b9250506060614e3087828801614bf7565b91505092959194509250565b60008060008060008060008060006101208a8c031215614e5b57600080fd5b6000614e678c8c614b4b565b9950506020614e788c828d01614bf7565b9850506040614e898c828d01614bf7565b9750506060614e9a8c828d01614bf7565b9650506080614eab8c828d01614bf7565b95505060a0614ebc8c828d01614bf7565b94505060c0614ecd8c828d01614bf7565b93505060e0614ede8c828d01614bf7565b925050610100614d7e8c828d01614bf7565b60008060208385031215614f0357600080fd5b823567ffffffffffffffff811115614f1a57600080fd5b614f2685828601614b56565b92509250509250929050565b6000806000806000806000806080898b031215614f4e57600080fd5b883567ffffffffffffffff811115614f6557600080fd5b614f718b828c01614b56565b9850985050602089013567ffffffffffffffff811115614f9057600080fd5b614f9c8b828c01614b56565b9650965050604089013567ffffffffffffffff811115614fbb57600080fd5b614fc78b828c01614b56565b9450945050606089013567ffffffffffffffff811115614fe657600080fd5b614ff28b828c01614b56565b92509250509295985092959890939650565b60006020828403121561501657600080fd5b60006119fc8484614baa565b6000806020838503121561503557600080fd5b823567ffffffffffffffff81111561504c57600080fd5b614f2685828601614bb5565b60006020828403121561506a57600080fd5b60006119fc8484614bf7565b6000806040838503121561508957600080fd5b60006150958585614bf7565b9250506020614cb885828601614b9f565b600080604083850312156150b957600080fd5b6000614ca78585614bf7565b600080600080608085870312156150db57600080fd5b60006150e78787614bf7565b94505060206150f887828801614bf7565b935050604061510987828801614b4b565b9250506060614e3087828801614b4b565b60008060006060848603121561512f57600080fd5b6000614daf8686614bf7565b60006151478383615ae6565b505060800190565b600061515b8383615b30565b505060600190565b600061516f8383615b63565b505060200190565b61518081615fcb565b82525050565b600061519182615fbe565b61519b8185615fc2565b93506151a683615fac565b8060005b838110156151d45781516151be888261513b565b97506151c983615fac565b9250506001016151aa565b509495945050505050565b60006151ea82615fbe565b6151f48185615fc2565b93506151ff83615fac565b8060005b838110156151d4578151615217888261514f565b975061522283615fac565b925050600101615203565b600061523882615fbe565b6152428185615fc2565b935061524d83615fac565b8060005b838110156151d45781516152658882615163565b975061527083615fac565b925050600101615251565b61518081615fd6565b61518081615fdb565b600061529882615fbe565b6152a28185615fc2565b93506152b2818560208601616000565b6152bb8161602c565b9093019392505050565b6000815460018116600081146152e2576001811461530857615347565b607f60028304166152f38187615fc2565b60ff1984168152955050602085019250615347565b600282046153168187615fc2565b955061532185615fb2565b60005b8281101561534057815488820152600190910190602001615324565b8701945050505b505092915050565b600061535b8385615fc2565b9350615368838584615ff4565b6152bb8361602c565b600061537e601783615fc2565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b60006153b7600b83615fc2565b6a1e995c9bc8185b5bdd5b9d60aa1b815260200192915050565b60006153de602683615fc2565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000615426601183615fc2565b700616c7265616479206c6f636b656420757607c1b815260200192915050565b6000615453601b83615fc2565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b600061548c602983615fc2565b7f63616c6c6572206973206e6f7420746865204e6f6465447269766572417574688152680818dbdb9d1c9858dd60ba1b602082015260400192915050565b60006154d7601083615fc2565b6f0dcdee8d0d2dcce40e8de40e6e8c2e6d60831b815260200192915050565b6000615503601283615fc2565b7134b731b7b93932b1ba10323ab930ba34b7b760711b815260200192915050565b6000615531601683615fc2565b7576616c696461746f722069736e27742061637469766560501b815260200192915050565b6000615563600c83615fc2565b6b77726f6e672073746174757360a01b815260200192915050565b600061558b601683615fc2565b751b9bdd08195b9bdd59da081d1a5b59481c185cdcd95960521b815260200192915050565b60006155bd601883615fc2565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b60006155f6600d83615fc2565b6c06e6f74206c6f636b656420757609c1b815260200192915050565b600061561f601b83615fc2565b7f746f6f206c617267652072657761726420706572207365636f6e640000000000815260200192915050565b6000615658602183615fc2565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f8152607760f81b602082015260400192915050565b600061569b600c83615fc2565b6b7a65726f207265776172647360a01b815260200192915050565b60006156c3601f83615fc2565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b60006156fc602083615fc2565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000615735602e83615fc2565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000615785602183615fc2565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e8152600360fc1b602082015260400192915050565b60006157c8601983615fc2565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b6000615801600c83615fc2565b6b656d707479207075626b657960a01b815260200192915050565b6000615829601283615fc2565b714661696c656420746f2073656e642046544d60701b815260200192915050565b6000615857601583615fc2565b741c995c5d595cdd08191bd95cdb89dd08195e1a5cdd605a1b815260200192915050565b6000615888602883615fc2565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e648152671032b0b93634b2b960c11b602082015260400192915050565b60006158d2601783615fc2565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b600061590b601883615fc2565b7f6f75747374616e64696e67207346544d2062616c616e63650000000000000000815260200192915050565b600061117f60008361199e565b6000615951601083615fc2565b6f6e6f7420656e6f756768207374616b6560801b815260200192915050565b600061597d602983615fc2565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381526808195e18d95959195960ba1b602082015260400192915050565b60006159c8601783615fc2565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b6000615a01601883615fc2565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b6000615a3a601783615fc2565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000615a73601683615fc2565b751cdd185ad9481a5cc8199d5b1b1e481cdb185cda195960521b815260200192915050565b6000615aa5602c83615fc2565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581526b2077686f6c65207374616b6560a01b602082015260400192915050565b80516080830190615af78482615177565b506020820151615b0a6020850182615b63565b506040820151615b1d6040850182615b63565b5060608201516122d26060850182615b63565b80516060830190615b418482615b63565b506020820151615b546020850182615b63565b5060408201516122d260408501825b615180816110e3565b600061117f82615937565b6020810161117f8284615177565b60408101615b938285615177565b6136266020830184615b63565b60808101615bae8287615177565b615bbb6020830186615b63565b615bc86040830185615b63565b6146e56060830184615b63565b6020808252810161117c8184615186565b6020808252810161117c81846151df565b6020808252810161117c818461522d565b6020810161117f828461527b565b6020810161117f8284615284565b6020808252810161117c818461528d565b6020808252810161117f81615371565b6020808252810161117f816153aa565b6020808252810161117f816153d1565b6020808252810161117f81615419565b6020808252810161117f81615446565b6020808252810161117f8161547f565b6020808252810161117f816154ca565b6020808252810161117f816154f6565b6020808252810161117f81615524565b6020808252810161117f81615556565b6020808252810161117f8161557e565b6020808252810161117f816155b0565b6020808252810161117f816155e9565b6020808252810161117f81615612565b6020808252810161117f8161564b565b6020808252810161117f8161568e565b6020808252810161117f816156b6565b6020808252810161117f816156ef565b6020808252810161117f81615728565b6020808252810161117f81615778565b6020808252810161117f816157bb565b6020808252810161117f816157f4565b6020808252810161117f8161581c565b6020808252810161117f8161584a565b6020808252810161117f8161587b565b6020808252810161117f816158c5565b6020808252810161117f816158fe565b6020808252810161117f81615944565b6020808252810161117f81615970565b6020808252810161117f816159bb565b6020808252810161117f816159f4565b6020808252810161117f81615a2d565b6020808252810161117f81615a66565b6020808252810161117f81615a98565b6020810161117f8284615b63565b60408101615e718285615b63565b81810360208301526119fc81846152c5565b60408101615e918286615b63565b81810360208301526146e581848661534f565b60408101615b938285615b63565b60608101615ec08286615b63565b615ecd6020830185615b63565b6119fc6040830184615b63565b60808101615bae8287615b63565b60e08101615ef6828a615b63565b615f036020830189615b63565b615f106040830188615b63565b615f1d6060830187615b63565b615f2a6080830186615b63565b615f3760a0830185615b63565b615f4460c0830184615177565b98975050505050505050565b60e08101615f5e828a615b63565b615f6b6020830189615b63565b615f786040830188615b63565b615f856060830187615b63565b615f926080830186615b63565b615f9f60a0830185615b63565b615f4460c0830184615b63565b60200190565b60009081526020902090565b5190565b90815260200190565b600061117f82615fe8565b151590565b6001600160e81b03191690565b6001600160a01b031690565b82818337506000910152565b60005b8381101561601b578181015183820152602001616003565b838111156122d25750506000910152565b601f01601f191690565b61603f81615fcb565b81146114e657600080fd5b61603f81615fd6565b61603f816110e356fea365627a7a723158202804d16a0eeb5ca6138bb77576e00d41e3e20db418ba1a2f9294e1ce2dfd5a9d6c6578706572696d656e74616cf564736f6c63430005110040"

var ContractBinRuntime = "0x6080604052600436106104315760003560e01c8063854873e111610229578063b88a37e21161012e578063d5a44f86116100b6578063e08d7e661161007a578063e08d7e6614610c80578063e261641a14610ca0578063e2f8c33614610cc0578063ebdf104c14610ce0578063f2fde38b14610d0057610431565b8063d5a44f8614610bdb578063d9a7c1f914610c0b578063dc31e1af14610c20578063de67f21514610c40578063df00c92214610c6057610431565b8063c65ee0e1116100fd578063c65ee0e114610b46578063c7be95de14610b66578063cc8343aa14610b7b578063cfd4766314610b9b578063cfdbb7cd14610bbb57610431565b8063b88a37e214610ac4578063bd14d90714610af1578063c3de580e14610b11578063c5f530af14610b3157610431565b8063a198d229116101b1578063a86a056f11610180578063a86a056f14610a1c578063b5d8962714610a3c578063b6d9edd514610a6f578063b810e41114610a8f578063b82b842714610aaf57610431565b8063a198d229146109b4578063a2f6e6bc146109d4578063a5a470ad146109f4578063a778651514610a0757610431565b80638cddb015116101f85780638cddb0151461091a5780638da5cb5b1461093a5780638f32d59b1461094f57806396c7ee46146109715780639fa6dd35146109a157610431565b8063854873e11461088b5780638b0e9f3f146108b85780638b1a0d11146108cd5780638c3c51d8146108ed57610431565b806339b80c001161033a5780636099ecb2116102c25780636f498663116102865780636f498663146107ff578063702797e31461081f578063715018a61461084c57806376671808146108615780637cacb1d61461087657610431565b80636099ecb21461076a57806361e53fcc1461078a578063634b91e3146107aa578063650acd66146107ca578063670322f8146107df57610431565b806354fd4d501161030957806354fd4d50146106f35780635601fe011461071557806358f95b80146107355780635e2308d2146105f35780635fab23a81461075557610431565b806339b80c0014610660578063441a3e70146106935780634f7c4efb146106b35780634feb92f3146106d357610431565b806318f628d4116103bd5780632265f2841161038c5780632265f284146105de5780632709275e146105f357806328f73148146106085780632cedb0971461061d578063346bdcfb1461064057610431565b806318f628d41461054f5780631d3ac42c1461056f5780631e702f831461058f5780631f270152146105af57610431565b80630d4955e3116104045780630d4955e3146104ce5780630d7b2609146104e35780630e559d82146104f857806312622d0e1461051a57806318160ddd1461053a57610431565b80630135b1db14610436578063019e27291461046c57806308c368741461048e5780630962ef79146104ae575b600080fd5b34801561044257600080fd5b50610456610451366004614c02565b610d20565b6040516104639190615e55565b60405180910390f35b34801561047857600080fd5b5061048c6104873660046150c5565b610d32565b005b34801561049a57600080fd5b5061048c6104a9366004615058565b610f0e565b3480156104ba57600080fd5b5061048c6104c9366004615058565b610fd7565b3480156104da57600080fd5b506104566110dd565b3480156104ef57600080fd5b506104566110e6565b34801561050457600080fd5b5061050d6110ed565b6040516104639190615b77565b34801561052657600080fd5b50610456610535366004614c88565b6110fc565b34801561054657600080fd5b50610456611185565b34801561055b57600080fd5b5061048c61056a366004614e3c565b61118b565b34801561057b57600080fd5b5061045661058a3660046150a6565b6112af565b34801561059b57600080fd5b5061048c6105aa3660046150a6565b6113d4565b3480156105bb57600080fd5b506105cf6105ca366004614d8e565b61142f565b60405161046393929190615eb2565b3480156105ea57600080fd5b50610456611461565b3480156105ff57600080fd5b50610456611473565b34801561061457600080fd5b5061045661148f565b34801561062957600080fd5b50610632611495565b604051610463929190615ea4565b34801561064c57600080fd5b5061048c61065b366004615058565b61149f565b34801561066c57600080fd5b5061068061067b366004615058565b6114e9565b6040516104639796959493929190615f50565b34801561069f57600080fd5b5061048c6106ae3660046150a6565b61152b565b3480156106bf57600080fd5b5061048c6106ce3660046150a6565b61181d565b3480156106df57600080fd5b5061048c6106ee366004614cc2565b6118dc565b3480156106ff57600080fd5b50610708611963565b6040516104639190615c16565b34801561072157600080fd5b50610456610730366004615058565b61196d565b34801561074157600080fd5b506104566107503660046150a6565b6119a3565b34801561076157600080fd5b506104566119c0565b34801561077657600080fd5b50610456610785366004614c88565b6119c6565b34801561079657600080fd5b506104566107a53660046150a6565b611a04565b3480156107b657600080fd5b5061048c6107c53660046150a6565b611a25565b3480156107d657600080fd5b50610456611bca565b3480156107eb57600080fd5b506104566107fa366004614c88565b611bcf565b34801561080b57600080fd5b5061045661081a366004614c88565b611c10565b34801561082b57600080fd5b5061083f61083a366004614ddb565b611c7a565b6040516104639190615be6565b34801561085857600080fd5b5061048c611d41565b34801561086d57600080fd5b50610456611daf565b34801561088257600080fd5b50610456611db8565b34801561089757600080fd5b506108ab6108a6366004615058565b611dbe565b6040516104639190615c24565b3480156108c457600080fd5b50610456611e59565b3480156108d957600080fd5b5061048c6108e83660046150a6565b611e5f565b3480156108f957600080fd5b5061090d6109083660046150a6565b611ecb565b6040516104639190615bd5565b34801561092657600080fd5b5061048c610935366004614c88565b611fac565b34801561094657600080fd5b5061050d611fd2565b34801561095b57600080fd5b50610964611fe1565b6040516104639190615c08565b34801561097d57600080fd5b5061099161098c366004614c88565b611ff2565b6040516104639493929190615eda565b61048c6109af366004615058565b612024565b3480156109c057600080fd5b506104566109cf3660046150a6565b61202f565b3480156109e057600080fd5b5061048c6109ef366004614c02565b612050565b61048c610a02366004615022565b612096565b348015610a1357600080fd5b50610456612127565b348015610a2857600080fd5b50610456610a37366004614c88565b61213d565b348015610a4857600080fd5b50610a5c610a57366004615058565b61215a565b6040516104639796959493929190615ee8565b348015610a7b57600080fd5b5061048c610a8a366004615058565b6121a0565b348015610a9b57600080fd5b506105cf610aaa366004614c88565b61222d565b348015610abb57600080fd5b50610456612259565b348015610ad057600080fd5b50610ae4610adf366004615058565b612260565b6040516104639190615bf7565b348015610afd57600080fd5b5061048c610b0c36600461511a565b6122c5565b348015610b1d57600080fd5b50610964610b2c366004615058565b6122d8565b348015610b3d57600080fd5b506104566122ef565b348015610b5257600080fd5b50610456610b61366004615058565b6122fd565b348015610b7257600080fd5b5061045661230f565b348015610b8757600080fd5b5061048c610b96366004615076565b612315565b348015610ba757600080fd5b50610456610bb6366004614c88565b612445565b348015610bc757600080fd5b50610964610bd6366004614c88565b612462565b348015610be757600080fd5b50610bfb610bf6366004615058565b6124f8565b6040516104639493929190615ba0565b348015610c1757600080fd5b50610456612539565b348015610c2c57600080fd5b50610456610c3b3660046150a6565b61253f565b348015610c4c57600080fd5b5061048c610c5b36600461511a565b612560565b348015610c6c57600080fd5b50610456610c7b3660046150a6565b6125b1565b348015610c8c57600080fd5b5061048c610c9b366004614ef0565b6125d2565b348015610cac57600080fd5b50610456610cbb3660046150a6565b61268d565b348015610ccc57600080fd5b5061048c610cdb366004614c20565b6126ae565b348015610cec57600080fd5b5061048c610cfb366004614f32565b61275d565b348015610d0c57600080fd5b5061048c610d1b366004614c02565b61291a565b60696020526000908152604090205481565b600054610100900460ff1680610d4b5750610d4b612947565b80610d59575060005460ff16155b610d7e5760405162461bcd60e51b8152600401610d7590615d55565b60405180910390fd5b600054610100900460ff16158015610da9576000805460ff1961ff0019909116610100171660011790555b610db28261294d565b6067859055606680546001600160a01b0319166001600160a01b03851617905560768490556755cfe697852e904c6075556103e86078556203f480607955610df8612a1f565b600086815260776020908152604080832060070193909355825160808101845282815290810182815292810182815260608201838152607c8054600181018255945291517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd0600490940293840180546001600160a01b0319166001600160a01b0390921691909117905592517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd183015591517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd282015590517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd3909101558015610f07576000805461ff00191690555b5050505050565b33610f176149f0565b610f218284612a23565b60208101518151919250600091610f3d9163ffffffff612af316565b9050610f608385610f5b856040015185612af390919063ffffffff16565b612b18565b6001600160a01b03831660008181526073602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610fc9939092909190615eb2565b60405180910390a350505050565b33610fe06149f0565b610fea8284612a23565b90506000826001600160a01b0316611027836040015161101b85602001518660000151612af390919063ffffffff16565b9063ffffffff612af316565b60405161103390615b6c565b60006040518083038185875af1925050503d8060008114611070576040519150601f19603f3d011682016040523d82523d6000602084013e611075565b606091505b50509050806110965760405162461bcd60e51b8152600401610d7590615d95565b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610fc993615eb2565b6301e133805b90565b6212750090565b607b546001600160a01b031681565b60006111088383612462565b61113657506001600160a01b038216600090815260726020908152604080832084845290915290205461117f565b6001600160a01b03831660008181526073602090815260408083208684528252808320549383526072825280832086845290915290205461117c9163ffffffff612b9916565b90505b92915050565b60765481565b61119433612bdb565b6111b05760405162461bcd60e51b8152600401610d7590615c85565b6111bb898989612bef565b6001600160a01b0389166000908152606f602090815260408083208b845290915290206002018190556111ed87612edd565b85156112a457868611156112135760405162461bcd60e51b8152600401610d7590615e45565b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e789061129a9088908c90615ea4565b60405180910390a3505b505050505050505050565b3360008181526073602090815260408083208684529091528120909190836112e95760405162461bcd60e51b8152600401610d7590615c45565b6112f38286612462565b61130f5760405162461bcd60e51b8152600401610d7590615cf5565b80548411156113305760405162461bcd60e51b8152600401610d7590615dc5565b61133a8286612f5b565b6113565760405162461bcd60e51b8152600401610d7590615dd5565b6113608286612ff8565b50600061137383878785600001546131c3565b8254869003835590506113878387836132f4565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b987846040516113c3929190615ea4565b60405180910390a395945050505050565b6113dd33612bdb565b6113f95760405162461bcd60e51b8152600401610d7590615c85565b806114165760405162461bcd60e51b8152600401610d7590615cc5565b61142082826134a0565b61142b826000612315565b5050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b600061146b6135bf565b601002905090565b6000606461147f6135bf565b601e028161148957fe5b04905090565b606d5481565b6078546079549091565b6114a7611fe1565b6114c35760405162461bcd60e51b8152600401610d7590615d45565b600081126114d85760768054820190556114e6565b607680546000839003900390555b50565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b336115346149f0565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526115a35760405162461bcd60e51b8152600401610d7590615da5565b6115ad8285612f5b565b6115c95760405162461bcd60e51b8152600401610d7590615dd5565b60208082015182516000878152606890935260409092206001015490919015801590611605575060008681526068602052604090206001015482115b15611626575050600084815260686020526040902060018101546002909101545b61162e612259565b8201611638612a1f565b10156116565760405162461bcd60e51b8152600401610d7590615cd5565b61165e611bca565b8101611668611daf565b10156116865760405162461bcd60e51b8152600401610d7590615e15565b6001600160a01b03841660009081526071602090815260408083208984528252808320888452909152812060020154906116bf886122d8565b905060006116e18383607a60008d8152602001908152602001600020546135cb565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116117435760405162461bcd60e51b8152600401610d7590615e35565b60006001600160a01b03881661175f858463ffffffff612b9916565b60405161176b90615b6c565b60006040518083038185875af1925050503d80600081146117a8576040519150601f19603f3d011682016040523d82523d6000602084013e6117ad565b606091505b50509050806117ce5760405162461bcd60e51b8152600401610d7590615d95565b888a896001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21876040516118099190615e55565b60405180910390a450505050505050505050565b611825611fe1565b6118415760405162461bcd60e51b8152600401610d7590615d45565b61184a826122d8565b6118665760405162461bcd60e51b8152600401610d7590615c35565b61186e6135bf565b81111561188d5760405162461bcd60e51b8152600401610d7590615d65565b6000828152607a6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917906118d0908490615e55565b60405180910390a25050565b6118e533612bdb565b6119015760405162461bcd60e51b8152600401610d7590615c85565b611949898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a9150899050888861362d565b606b548811156112a457606b889055505050505050505050565b620ccc0d60ea1b90565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b606e5481565b60006119d06149f0565b6119da84846137a3565b8051602082015160408301519293506119fc9261101b9163ffffffff612af316565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b33611a308184612ff8565b5060008211611a515760405162461bcd60e51b8152600401610d7590615c45565b611a5b81846110fc565b821115611a7a5760405162461bcd60e51b8152600401610d7590615d75565b611a848184612f5b565b611aa05760405162461bcd60e51b8152600401610d7590615dd5565b6001600160a01b0381166000908152607e602090815260408083208684529091529020805460018101909155611ad78285856132f4565b6001600160a01b038216600090815260716020908152604080832087845282528083208484529091529020600201839055611b10611daf565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152902055611b44612a1f565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152812060010191909155611b81908590612315565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d5786604051611bbc9190615e55565b60405180910390a450505050565b600390565b6000611bdb8383612462565b611be75750600061117f565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b6000611c1a6149f0565b506001600160a01b0383166000908152606f60209081526040808320858452825291829020825160608101845281548082526001830154938201849052600290920154938101849052926119fc92909161101b919063ffffffff612af316565b60608082604051908082528060200260200182016040528015611cb757816020015b611ca46149f0565b815260200190600190039081611c9c5790505b50905060005b83811015611d37576001600160a01b03871660009081526071602090815260408083208984528252808320888501845282529182902082516060810184528154815260018201549281019290925260020154918101919091528251839083908110611d2457fe5b6020908102919091010152600101611cbd565b5095945050505050565b611d49611fe1565b611d655760405162461bcd60e51b8152600401610d7590615d45565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b60675460010190565b60675481565b606a6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015611e515780601f10611e2657610100808354040283529160200191611e51565b820191906000526020600020905b815481529060010190602001808311611e3457829003601f168201915b505050505081565b606c5481565b611e67611fe1565b611e835760405162461bcd60e51b8152600401610d7590615d45565b607981905560788290556040517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a3490611ebf9084908490615ea4565b60405180910390a15050565b607c54604080518381526020808502820101909152606091908290848015611f0d57816020015b611efa614a11565b815260200190600190039081611ef25790505b50905060005b84811015611fa3578281870110611f2957611fa3565b607c81870181548110611f3857fe5b600091825260209182902060408051608081018252600490930290910180546001600160a01b0316835260018101549383019390935260028301549082015260039091015460608201528251839083908110611f9057fe5b6020908102919091010152600101611f13565b50949350505050565b611fb68282612ff8565b61142b5760405162461bcd60e51b8152600401610d7590615c95565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6114e6338234612b18565b60009182526077602090815260408084209284526005909201905290205490565b612058611fe1565b6120745760405162461bcd60e51b8152600401610d7590615d45565b607b80546001600160a01b0319166001600160a01b0392909216919091179055565b61209e6122ef565b3410156120bd5760405162461bcd60e51b8152600401610d7590615e25565b806120da5760405162461bcd60e51b8152600401610d7590615d85565b61211a3383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061381192505050565b61142b33606b5434612b18565b600060646121336135bf565b600f028161148957fe5b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b6121a8611fe1565b6121c45760405162461bcd60e51b8152600401610d7590615d45565b6801c985c8903591eb208111156121ed5760405162461bcd60e51b8152600401610d7590615d05565b60758190556040517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae239690612222908390615e55565b60405180910390a150565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b62093a8090565b6000818152607760209081526040918290206006018054835181840281018401909452808452606093928301828280156122b957602002820191906000526020600020905b8154815260200190600101908083116122a5575b50505050509050919050565b336122d28185858561383c565b50505050565b600090815260686020526040902054608016151590565b6969e10de76676d080000090565b607a6020526000908152604090205481565b606b5481565b61231e82613a19565b61233a5760405162461bcd60e51b8152600401610d7590615e05565b60008281526068602052604090206003810154905415612358575060005b60665460405163520337df60e11b81526001600160a01b039091169063a4066fbe9061238a9086908590600401615ea4565b600060405180830381600087803b1580156123a457600080fd5b505af11580156123b8573d6000803e3d6000fd5b505050508180156123c857508015155b15612440576066546000848152606a602052604090819020905163242a6e3f60e01b81526001600160a01b039092169163242a6e3f9161240d91879190600401615e63565b600060405180830381600087803b15801561242757600080fd5b505af115801561243b573d6000803e3d6000fd5b505050505b505050565b607260209081526000928352604080842090915290825290205481565b6001600160a01b0382166000908152607360209081526040808320848452909152812060020154158015906124b957506001600160a01b038316600090815260736020908152604080832085845290915290205415155b801561117c57506001600160a01b03831660009081526073602090815260408083208584529091529020600201546124ef612a1f565b11159392505050565b607c818154811061250557fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b039092169350919084565b60755481565b60009182526077602090815260408084209284526003909201905290205490565b338161257e5760405162461bcd60e51b8152600401610d7590615c45565b6125888185612462565b156125a55760405162461bcd60e51b8152600401610d7590615c65565b6122d28185858561383c565b60009182526077602090815260408084209284526002909201905290205490565b6125db33612bdb565b6125f75760405162461bcd60e51b8152600401610d7590615c85565b600060776000612605611daf565b8152602001908152602001600020905060008090505b8281101561267e57600084848381811061263157fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c86015490935061266f91508263ffffffff612af316565b600c850155505060010161261b565b506122d2600682018484614a42565b60009182526077602090815260408084209284526004909201905290205490565b6126b6611fe1565b6126d25760405162461bcd60e51b8152600401610d7590615d45565b6126db83612edd565b6040516001600160a01b0385169084156108fc029085906000818181858888f19350505050158015612711573d6000803e3d6000fd5b50836001600160a01b03167f9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad84848460405161274f93929190615e83565b60405180910390a250505050565b61276633612bdb565b6127825760405162461bcd60e51b8152600401610d7590615c85565b600060776000612790611daf565b815260200190815260200160002090506060816006018054806020026020016040519081016040528092919081815260200182805480156127f057602002820191906000526020600020905b8154815260200190600101908083116127dc575b5050505050905061287782828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613a3092505050565b6128e6828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a918291850190849080828437600092019190915250613b3f92505050565b6128ee611daf565b6067556128f9612a1f565b600783015550607554600b820155607654600d909101555050505050505050565b612922611fe1565b61293e5760405162461bcd60e51b8152600401610d7590615d45565b6114e681614185565b303b1590565b600054610100900460ff16806129665750612966612947565b80612974575060005460ff16155b6129905760405162461bcd60e51b8152600401610d7590615d55565b600054610100900460ff161580156129bb576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3801561142b576000805461ff00191690555050565b4290565b612a2b6149f0565b612a358383612ff8565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392612a939261101b9163ffffffff612af316565b905080612ab25760405162461bcd60e51b8152600401610d7590615d25565b6001600160a01b0384166000908152606f6020908152604080832086845290915281208181556001810182905560020155612aec81612edd565b5092915050565b60008282018381101561117c5760405162461bcd60e51b8152600401610d7590615c75565b612b2182613a19565b612b3d5760405162461bcd60e51b8152600401610d7590615e05565b60008281526068602052604090205415612b695760405162461bcd60e51b8152600401610d7590615cb5565b612b74838383612bef565b612b7d82614207565b6124405760405162461bcd60e51b8152600401610d7590615df5565b600061117c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061424f565b6066546001600160a01b0390811691161490565b60008111612c0f5760405162461bcd60e51b8152600401610d7590615c45565b612c198383612ff8565b506001600160a01b0383166000908152607d6020908152604080832085845290915290205480612d4c57607c80546001600160a01b038681166000818152607d602090815260408083208a84528252808320869055805160808101825293845290830189815290830188815242606085019081526001870188559690925291517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd0600490950294850180546001600160a01b0319169190941617909255517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd1830155517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd282015590517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd390910155612dc3565b612d7d82607c8381548110612d5d57fe5b906000526020600020906004020160020154612af390919063ffffffff16565b607c8281548110612d8a57fe5b90600052602060002090600402016002018190555042607c8281548110612dad57fe5b9060005260206000209060040201600301819055505b6001600160a01b0384166000908152607260209081526040808320868452909152902054612df7908363ffffffff612af316565b6001600160a01b0385166000908152607260209081526040808320878452825280832093909355606890522060030154612e37818463ffffffff612af316565b600085815260686020526040902060030155606c54612e5c908463ffffffff612af316565b606c55600084815260686020526040902054612e8957606d54612e85908463ffffffff612af316565b606d555b612e94848215612315565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb85604051612ece9190615e55565b60405180910390a35050505050565b6066546040516366e7ea0f60e01b81526001600160a01b03909116906366e7ea0f90612f0f9030908590600401615b85565b600060405180830381600087803b158015612f2957600080fd5b505af1158015612f3d573d6000803e3d6000fd5b5050607654612f55925090508263ffffffff612af316565b60765550565b607b546000906001600160a01b0316612f765750600161117f565b607b546040516321d585c360e01b81526001600160a01b03909116906321d585c390612fa89086908690600401615b85565b60206040518083038186803b158015612fc057600080fd5b505afa158015612fd4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061117c9190810190615004565b60006130026149f0565b61300c848461427b565b9050613017836143b3565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f82528281208682528252829020825160608101845281548152600182015492810192909252600201549181019190915261307d908261440e565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260748252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526130f7908261440e565b6001600160a01b03851660009081526074602090815260408083208784528252918290208351815590830151600182015591015160029091015561313b8484612462565b61319e576001600160a01b0384166000818152607360209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260748352818420888552909252822082815592830182905591909101555b60208101511515806131b05750805115155b806119fc57506040015115159392505050565b6001600160a01b0384166000908152607460209081526040808320868452909152812054819061320b9084906131ff908763ffffffff61448016565b9063ffffffff6144ba16565b6001600160a01b03871660009081526074602090815260408083208984529091528120600101549192509061324c9085906131ff908863ffffffff61448016565b6001600160a01b03881660009081526074602090815260408083208a8452909152902054909150600282048301906132849084612b99565b6001600160a01b03891660009081526074602090815260408083208b84529091529020908155600101546132b89083612b99565b6001600160a01b03891660009081526074602090815260408083208b84529091529020600101558581106132e95750845b979650505050505050565b6001600160a01b0383166000908152607d60209081526040808320858452909152902054607c805461334c9184918490811061332c57fe5b906000526020600020906004020160020154612b9990919063ffffffff16565b607c828154811061335957fe5b906000526020600020906004020160020181905550607c818154811061337b57fe5b9060005260206000209060040201600201546000141561339e5761339e816144fc565b6001600160a01b038416600090815260726020908152604080832086845282528083208054869003905560689091529020600301546133e3908363ffffffff612b9916565b600084815260686020526040902060030155606c54613408908363ffffffff612b9916565b606c5560008381526068602052604090205461343557606d54613431908363ffffffff612b9916565b606d555b60006134408461196d565b90508015613499576134506122ef565b81101561346f5760405162461bcd60e51b8152600401610d7590615e25565b61347884614207565b6134945760405162461bcd60e51b8152600401610d7590615df5565b610f07565b610f078460015b6000828152606860205260409020541580156134bb57508015155b156134e857600082815260686020526040902060030154606d546134e49163ffffffff612b9916565b606d555b60008281526068602052604090205481111561142b5760008281526068602052604090208181556002015461358f5761351f611daf565b600083815260686020526040902060020155613539612a1f565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e479261358692909190615ea4565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516118d09190615e55565b670de0b6b3a764000090565b60008215806135e157506135dd6135bf565b8210155b156135ee57506000613626565b613619600161101b6135fe6135bf565b6131ff8661360a6135bf565b8a91900363ffffffff61448016565b9050838111156136265750825b9392505050565b6001600160a01b038816600090815260696020526040902054156136635760405162461bcd60e51b8152600401610d7590615ce5565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180546001600160a01b031916909417909355606a815291902087516136d492890190614a8d565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051613711929190615ea4565b60405180910390a3811561375a57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613751929190615ea4565b60405180910390a25b841561379957867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e866040516137909190615e55565b60405180910390a25b5050505050505050565b6137ab6149f0565b6137b36149f0565b6137bd848461427b565b6001600160a01b0385166000908152606f6020908152604080832087845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506119fc908261440e565b606b805460010190819055612440838284600061382c611daf565b613834612a1f565b60008061362d565b61384684846110fc565b8111156138655760405162461bcd60e51b8152600401610d7590615de5565b600083815260686020526040902054156138915760405162461bcd60e51b8152600401610d7590615cb5565b6138996110e6565b82101580156138af57506138ab6110dd565b8211155b6138cb5760405162461bcd60e51b8152600401610d7590615ca5565b60006138d98361101b612a1f565b6000858152606860205260409020600601549091506001600160a01b039081169086168114613948576001600160a01b03811660009081526073602090815260408083208884529091529020600201548211156139485760405162461bcd60e51b8152600401610d7590615db5565b6139528686612ff8565b506001600160a01b03861660009081526073602090815260408083208884529091529020600381015485101561399a5760405162461bcd60e51b8152600401610d7590615d35565b80546139ac908563ffffffff612af316565b81556139b6611daf565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7890613a089089908990615ea4565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b60005b8351811015610f0757607854828281518110613a4b57fe5b6020026020010151118015613a755750607954838281518110613a6a57fe5b602002602001015110155b15613ab657613a98848281518110613a8957fe5b602002602001015160086134a0565b613ab6848281518110613aa757fe5b60200260200101516000612315565b828181518110613ac257fe5b6020026020010151856004016000868481518110613adc57fe5b6020026020010151815260200190815260200160002081905550818181518110613b0257fe5b6020026020010151856005016000868481518110613b1c57fe5b602090810291909101810151825281019190915260400160002055600101613a33565b613b47614afb565b6040518060c001604052808551604051908082528060200260200182016040528015613b7d578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015613bb5578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060776000613bef6001613be3611daf565b9063ffffffff612b9916565b81526020810191909152604001600020600160808401526007810154909150613c16612a1f565b1115613c30578060070154613c29612a1f565b0360808301525b60005b8551811015613d38576000826003016000888481518110613c5057fe5b60200260200101518152602001908152602001600020549050600080905081868481518110613c7b57fe5b60200260200101511115613ca25781868481518110613c9657fe5b60200260200101510390505b8460800151878481518110613cb357fe5b6020026020010151820281613cc457fe5b0485604001518481518110613cd557fe5b602002602001018181525050613d0f85604001518481518110613cf457fe5b60200260200101518660600151612af390919063ffffffff16565b606086015260a0850151613d29908263ffffffff612af316565b60a08601525050600101613c33565b5060005b8551811015613e09578260800151858281518110613d5657fe5b60200260200101518460800151878481518110613d6f57fe5b60200260200101518a60000160008b8781518110613d8957fe5b60200260200101518152602001908152602001600020540281613da857fe5b040281613db157fe5b0483600001518281518110613dc257fe5b602002602001018181525050613dfc83600001518281518110613de157fe5b60200260200101518460200151612af390919063ffffffff16565b6020840152600101613d3c565b5060005b855181101561415d576000613e45846080015160755486600001518581518110613e3357fe5b60200260200101518760200151614650565b9050613e81613e748560a0015186604001518581518110613e6257fe5b60200260200101518760600151614691565b829063ffffffff612af316565b90506000878381518110613e9157fe5b6020908102919091018101516000818152606890925260408220600601549092506001600160a01b031690613ecd84613ec8612127565b6146ee565b6001600160a01b0383166000908152607260209081526040808320878452909152902054909150801561407457600081613f078587611bcf565b840281613f1057fe5b049050808303613f1e6149f0565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154613f5090849061470b565b9050613f5a6149f0565b613f6583600061470b565b6001600160a01b0388166000908152606f602090815260408083208c84528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150613fba9083836147fc565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d8452825291829020825160608101845281548152948101549185019190915290910154908201526140359083836147fc565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156140a657816140996135bf565b8402816140a257fe5b0490505b808a600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508b89815181106140e357fe5b60200260200101518f6003016000898152602001908152602001600020819055508c898151811061411057fe5b60200260200101518a600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050613e0d565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b0381166141ab5760405162461bcd60e51b8152600401610d7590615c55565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b60006142346142146135bf565b6131ff61421f611461565b6142288661196d565b9063ffffffff61448016565b60008381526068602052604090206003015411159050919050565b600081848411156142735760405162461bcd60e51b8152600401610d759190615c24565b505050900390565b6142836149f0565b6001600160a01b0383166000908152607060209081526040808320858452909152812054906142b1846143b3565b905060006142bf8686614817565b9050818111156142cc5750805b828110156142d75750815b6001600160a01b0386166000818152607360209081526040808320898452825280832093835260728252808320898452909152812054825490919061432390839063ffffffff612b9916565b9050600061433784600001548a89886148d6565b90506143416149f0565b61434f82866003015461470b565b905061435d838b8a896148d6565b91506143676149f0565b61437283600061470b565b9050614380858c898b6148d6565b925061438a6149f0565b61439584600061470b565b90506143a28383836147fc565b9d9c50505050505050505050505050565b600081815260686020526040812060020154156144065760008281526068602052604090206002015460675410156143ee575060675461199e565b5060008181526068602052604090206002015461199e565b505060675490565b6144166149f0565b6040805160608101909152825184518291614437919063ffffffff612af316565b815260200161445784602001518660200151612af390919063ffffffff16565b815260200161447784604001518660400151612af390919063ffffffff16565b90529392505050565b60008261448f5750600061117f565b8282028284828161449c57fe5b041461117c5760405162461bcd60e51b8152600401610d7590615d15565b600061117c83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250614939565b607c5460009061451390600163ffffffff612b9916565b905080821461460a57607c818154811061452957fe5b9060005260206000209060040201607c838154811061454457fe5b60009182526020909120825460049092020180546001600160a01b0319166001600160a01b039092169190911781556001808301549082015560028083015490820155600391820154910155614598614a11565b607c82815481106145a557fe5b6000918252602080832060408051608081018252600490940290910180546001600160a01b03168085526001820154858501908152600283015486850152600390920154606090950194909452928452607d8252808420925184529190529020839055505b607c80548061461557fe5b60008281526020812060046000199093019283020180546001600160a01b031916815560018101829055600281018290556003015590555050565b60008261465f575060006119fc565b6000614671868663ffffffff61448016565b9050614687836131ff838763ffffffff61448016565b9695505050505050565b6000826146a057506000613626565b60006146b6836131ff878763ffffffff61448016565b90506146e56146c36135bf565b6131ff6146ce611473565b6146d66135bf565b8591900363ffffffff61448016565b95945050505050565b600061117c6146fb6135bf565b6131ff858563ffffffff61448016565b6147136149f0565b6040518060600160405280600081526020016000815260200160008152509050816000146147ce576000614745611473565b61474d6135bf565b039050600061476d61475d6110dd565b6131ff848763ffffffff61448016565b9050600061479661477c6135bf565b6131ff84614788611473565b8a910163ffffffff61448016565b90506147bb6147a36135bf565b6131ff6147ae611473565b899063ffffffff61448016565b60208501819052900383525061117f9050565b6147f16147d96135bf565b6131ff6147e4611473565b869063ffffffff61448016565b604082015292915050565b6148046149f0565b6119fc614811858561440e565b8361440e565b6001600160a01b038216600090815260736020908152604080832084845290915281206001015460675461484c858583614970565b1561485a57915061117f9050565b614865858584614970565b6148745760009250505061117f565b808211156148875760009250505061117f565b808210156148ba576002818301046148a0868683614970565b156148b0578060010192506148b4565b8091505b50614887565b806148ca5760009250505061117f565b60001901949350505050565b60008183106148e7575060006119fc565b6000838152607760208181526040808420888552600190810183528185205487865293835281852089865201909152909120546132e96149256135bf565b6131ff89614228858763ffffffff612b9916565b6000818361495a5760405162461bcd60e51b8152600401610d759190615c24565b50600083858161496657fe5b0495945050505050565b6001600160a01b038316600090815260736020908152604080832085845290915281206001015482108015906119fc57506001600160a01b03841660009081526073602090815260408083208684529091529020600201546149d1836149db565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b604051806080016040528060006001600160a01b031681526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215614a7d579160200282015b82811115614a7d578235825591602001919060010190614a62565b50614a89929150614b31565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614ace57805160ff1916838001178555614a7d565b82800160010185558215614a7d579182015b82811115614a7d578251825591602001919060010190614ae0565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b6110e391905b80821115614a895760008155600101614b37565b803561117f81616036565b60008083601f840112614b6857600080fd5b50813567ffffffffffffffff811115614b8057600080fd5b602083019150836020820283011115614b9857600080fd5b9250929050565b803561117f8161604a565b805161117f8161604a565b60008083601f840112614bc757600080fd5b50813567ffffffffffffffff811115614bdf57600080fd5b602083019150836001820283011115614b9857600080fd5b803561117f81616053565b600060208284031215614c1457600080fd5b60006119fc8484614b4b565b60008060008060608587031215614c3657600080fd5b6000614c428787614b4b565b9450506020614c5387828801614bf7565b935050604085013567ffffffffffffffff811115614c7057600080fd5b614c7c87828801614bb5565b95989497509550505050565b60008060408385031215614c9b57600080fd5b6000614ca78585614b4b565b9250506020614cb885828601614bf7565b9150509250929050565b60008060008060008060008060006101008a8c031215614ce157600080fd5b6000614ced8c8c614b4b565b9950506020614cfe8c828d01614bf7565b98505060408a013567ffffffffffffffff811115614d1b57600080fd5b614d278c828d01614bb5565b97509750506060614d3a8c828d01614bf7565b9550506080614d4b8c828d01614bf7565b94505060a0614d5c8c828d01614bf7565b93505060c0614d6d8c828d01614bf7565b92505060e0614d7e8c828d01614bf7565b9150509295985092959850929598565b600080600060608486031215614da357600080fd5b6000614daf8686614b4b565b9350506020614dc086828701614bf7565b9250506040614dd186828701614bf7565b9150509250925092565b60008060008060808587031215614df157600080fd5b6000614dfd8787614b4b565b9450506020614e0e87828801614bf7565b9350506040614e1f87828801614bf7565b9250506060614e3087828801614bf7565b91505092959194509250565b60008060008060008060008060006101208a8c031215614e5b57600080fd5b6000614e678c8c614b4b565b9950506020614e788c828d01614bf7565b9850506040614e898c828d01614bf7565b9750506060614e9a8c828d01614bf7565b9650506080614eab8c828d01614bf7565b95505060a0614ebc8c828d01614bf7565b94505060c0614ecd8c828d01614bf7565b93505060e0614ede8c828d01614bf7565b925050610100614d7e8c828d01614bf7565b60008060208385031215614f0357600080fd5b823567ffffffffffffffff811115614f1a57600080fd5b614f2685828601614b56565b92509250509250929050565b6000806000806000806000806080898b031215614f4e57600080fd5b883567ffffffffffffffff811115614f6557600080fd5b614f718b828c01614b56565b9850985050602089013567ffffffffffffffff811115614f9057600080fd5b614f9c8b828c01614b56565b9650965050604089013567ffffffffffffffff811115614fbb57600080fd5b614fc78b828c01614b56565b9450945050606089013567ffffffffffffffff811115614fe657600080fd5b614ff28b828c01614b56565b92509250509295985092959890939650565b60006020828403121561501657600080fd5b60006119fc8484614baa565b6000806020838503121561503557600080fd5b823567ffffffffffffffff81111561504c57600080fd5b614f2685828601614bb5565b60006020828403121561506a57600080fd5b60006119fc8484614bf7565b6000806040838503121561508957600080fd5b60006150958585614bf7565b9250506020614cb885828601614b9f565b600080604083850312156150b957600080fd5b6000614ca78585614bf7565b600080600080608085870312156150db57600080fd5b60006150e78787614bf7565b94505060206150f887828801614bf7565b935050604061510987828801614b4b565b9250506060614e3087828801614b4b565b60008060006060848603121561512f57600080fd5b6000614daf8686614bf7565b60006151478383615ae6565b505060800190565b600061515b8383615b30565b505060600190565b600061516f8383615b63565b505060200190565b61518081615fcb565b82525050565b600061519182615fbe565b61519b8185615fc2565b93506151a683615fac565b8060005b838110156151d45781516151be888261513b565b97506151c983615fac565b9250506001016151aa565b509495945050505050565b60006151ea82615fbe565b6151f48185615fc2565b93506151ff83615fac565b8060005b838110156151d4578151615217888261514f565b975061522283615fac565b925050600101615203565b600061523882615fbe565b6152428185615fc2565b935061524d83615fac565b8060005b838110156151d45781516152658882615163565b975061527083615fac565b925050600101615251565b61518081615fd6565b61518081615fdb565b600061529882615fbe565b6152a28185615fc2565b93506152b2818560208601616000565b6152bb8161602c565b9093019392505050565b6000815460018116600081146152e2576001811461530857615347565b607f60028304166152f38187615fc2565b60ff1984168152955050602085019250615347565b600282046153168187615fc2565b955061532185615fb2565b60005b8281101561534057815488820152600190910190602001615324565b8701945050505b505092915050565b600061535b8385615fc2565b9350615368838584615ff4565b6152bb8361602c565b600061537e601783615fc2565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b60006153b7600b83615fc2565b6a1e995c9bc8185b5bdd5b9d60aa1b815260200192915050565b60006153de602683615fc2565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000615426601183615fc2565b700616c7265616479206c6f636b656420757607c1b815260200192915050565b6000615453601b83615fc2565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b600061548c602983615fc2565b7f63616c6c6572206973206e6f7420746865204e6f6465447269766572417574688152680818dbdb9d1c9858dd60ba1b602082015260400192915050565b60006154d7601083615fc2565b6f0dcdee8d0d2dcce40e8de40e6e8c2e6d60831b815260200192915050565b6000615503601283615fc2565b7134b731b7b93932b1ba10323ab930ba34b7b760711b815260200192915050565b6000615531601683615fc2565b7576616c696461746f722069736e27742061637469766560501b815260200192915050565b6000615563600c83615fc2565b6b77726f6e672073746174757360a01b815260200192915050565b600061558b601683615fc2565b751b9bdd08195b9bdd59da081d1a5b59481c185cdcd95960521b815260200192915050565b60006155bd601883615fc2565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b60006155f6600d83615fc2565b6c06e6f74206c6f636b656420757609c1b815260200192915050565b600061561f601b83615fc2565b7f746f6f206c617267652072657761726420706572207365636f6e640000000000815260200192915050565b6000615658602183615fc2565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f8152607760f81b602082015260400192915050565b600061569b600c83615fc2565b6b7a65726f207265776172647360a01b815260200192915050565b60006156c3601f83615fc2565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b60006156fc602083615fc2565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000615735602e83615fc2565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000615785602183615fc2565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e8152600360fc1b602082015260400192915050565b60006157c8601983615fc2565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b6000615801600c83615fc2565b6b656d707479207075626b657960a01b815260200192915050565b6000615829601283615fc2565b714661696c656420746f2073656e642046544d60701b815260200192915050565b6000615857601583615fc2565b741c995c5d595cdd08191bd95cdb89dd08195e1a5cdd605a1b815260200192915050565b6000615888602883615fc2565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e648152671032b0b93634b2b960c11b602082015260400192915050565b60006158d2601783615fc2565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b600061590b601883615fc2565b7f6f75747374616e64696e67207346544d2062616c616e63650000000000000000815260200192915050565b600061117f60008361199e565b6000615951601083615fc2565b6f6e6f7420656e6f756768207374616b6560801b815260200192915050565b600061597d602983615fc2565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381526808195e18d95959195960ba1b602082015260400192915050565b60006159c8601783615fc2565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b6000615a01601883615fc2565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b6000615a3a601783615fc2565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000615a73601683615fc2565b751cdd185ad9481a5cc8199d5b1b1e481cdb185cda195960521b815260200192915050565b6000615aa5602c83615fc2565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581526b2077686f6c65207374616b6560a01b602082015260400192915050565b80516080830190615af78482615177565b506020820151615b0a6020850182615b63565b506040820151615b1d6040850182615b63565b5060608201516122d26060850182615b63565b80516060830190615b418482615b63565b506020820151615b546020850182615b63565b5060408201516122d260408501825b615180816110e3565b600061117f82615937565b6020810161117f8284615177565b60408101615b938285615177565b6136266020830184615b63565b60808101615bae8287615177565b615bbb6020830186615b63565b615bc86040830185615b63565b6146e56060830184615b63565b6020808252810161117c8184615186565b6020808252810161117c81846151df565b6020808252810161117c818461522d565b6020810161117f828461527b565b6020810161117f8284615284565b6020808252810161117c818461528d565b6020808252810161117f81615371565b6020808252810161117f816153aa565b6020808252810161117f816153d1565b6020808252810161117f81615419565b6020808252810161117f81615446565b6020808252810161117f8161547f565b6020808252810161117f816154ca565b6020808252810161117f816154f6565b6020808252810161117f81615524565b6020808252810161117f81615556565b6020808252810161117f8161557e565b6020808252810161117f816155b0565b6020808252810161117f816155e9565b6020808252810161117f81615612565b6020808252810161117f8161564b565b6020808252810161117f8161568e565b6020808252810161117f816156b6565b6020808252810161117f816156ef565b6020808252810161117f81615728565b6020808252810161117f81615778565b6020808252810161117f816157bb565b6020808252810161117f816157f4565b6020808252810161117f8161581c565b6020808252810161117f8161584a565b6020808252810161117f8161587b565b6020808252810161117f816158c5565b6020808252810161117f816158fe565b6020808252810161117f81615944565b6020808252810161117f81615970565b6020808252810161117f816159bb565b6020808252810161117f816159f4565b6020808252810161117f81615a2d565b6020808252810161117f81615a66565b6020808252810161117f81615a98565b6020810161117f8284615b63565b60408101615e718285615b63565b81810360208301526119fc81846152c5565b60408101615e918286615b63565b81810360208301526146e581848661534f565b60408101615b938285615b63565b60608101615ec08286615b63565b615ecd6020830185615b63565b6119fc6040830184615b63565b60808101615bae8287615b63565b60e08101615ef6828a615b63565b615f036020830189615b63565b615f106040830188615b63565b615f1d6060830187615b63565b615f2a6080830186615b63565b615f3760a0830185615b63565b615f4460c0830184615177565b98975050505050505050565b60e08101615f5e828a615b63565b615f6b6020830189615b63565b615f786040830188615b63565b615f856060830187615b63565b615f926080830186615b63565b615f9f60a0830185615b63565b615f4460c0830184615b63565b60200190565b60009081526020902090565b5190565b90815260200190565b600061117f82615fe8565b151590565b6001600160e81b03191690565b6001600160a01b031690565b82818337506000910152565b60005b8381101561601b578181015183820152602001616003565b838111156122d25750506000910152565b601f01601f191690565b61603f81615fcb565b81146114e657600080fd5b61603f81615fd6565b61603f816110e356fea365627a7a723158202804d16a0eeb5ca6138bb77576e00d41e3e20db418ba1a2f9294e1ce2dfd5a9d6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// SFCStake is an auto generated low-level Go binding around an user-defined struct.
type SFCStake struct {
	Delegator   common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Timestamp   *big.Int
}

// SFCWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type SFCWithdrawalRequest struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractCaller) BaseRewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "baseRewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractSession) BaseRewardPerSecond() (*big.Int, error) {
	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
}

// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
//
// Solidity: function baseRewardPerSecond() view returns(uint256)
func (_Contract *ContractCallerSession) BaseRewardPerSecond() (*big.Int, error) {
	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "contractCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
//
// Solidity: function contractCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ContractCommission() (*big.Int, error) {
	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentSealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedOriginatedTxsFee(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedOriginatedTxsFee", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedRewardPerToken(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedRewardPerToken", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochAccumulatedUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedUptime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineBlocks(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineBlocks", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochOfflineTime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochOfflineTime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetEpochReceivedStake(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochReceivedStake", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCaller) GetEpochSnapshot(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochSnapshot", arg0)

	outstruct := new(struct {
		EndTime               *big.Int
		EpochFee              *big.Int
		TotalBaseRewardWeight *big.Int
		TotalTxRewardWeight   *big.Int
		BaseRewardPerSecond   *big.Int
		TotalStake            *big.Int
		TotalSupply           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBaseRewardWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalTxRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_Contract *ContractCallerSession) GetEpochSnapshot(arg0 *big.Int) (struct {
	EndTime               *big.Int
	EpochFee              *big.Int
	TotalBaseRewardWeight *big.Int
	TotalTxRewardWeight   *big.Int
	BaseRewardPerSecond   *big.Int
	TotalStake            *big.Int
	TotalSupply           *big.Int
}, error) {
	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCaller) GetEpochValidatorIDs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochValidatorIDs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_Contract *ContractCallerSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetLockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
//
// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) GetLockupInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLockupInfo", arg0, arg1)

	outstruct := new(struct {
		LockedStake *big.Int
		FromEpoch   *big.Int
		EndTime     *big.Int
		Duration    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FromEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
//
// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
	LockedStake *big.Int
	FromEpoch   *big.Int
	EndTime     *big.Int
	Duration    *big.Int
}, error) {
	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) GetSelfStake(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSelfStake", validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) GetStake(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStake", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
}

// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
//
// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint256,uint256,uint256)[])
func (_Contract *ContractCaller) GetStakes(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]SFCStake, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStakes", offset, limit)

	if err != nil {
		return *new([]SFCStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]SFCStake)).(*[]SFCStake)

	return out0, err

}

// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
//
// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint256,uint256,uint256)[])
func (_Contract *ContractSession) GetStakes(offset *big.Int, limit *big.Int) ([]SFCStake, error) {
	return _Contract.Contract.GetStakes(&_Contract.CallOpts, offset, limit)
}

// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
//
// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint256,uint256,uint256)[])
func (_Contract *ContractCallerSession) GetStakes(offset *big.Int, limit *big.Int) ([]SFCStake, error) {
	return _Contract.Contract.GetStakes(&_Contract.CallOpts, offset, limit)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCaller) GetStashedLockupRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStashedLockupRewards", arg0, arg1)

	outstruct := new(struct {
		LockupExtraReward *big.Int
		LockupBaseReward  *big.Int
		UnlockedReward    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockupExtraReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockupBaseReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnlockedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
//
// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractCallerSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
}, error) {
	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) GetUnlockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUnlockedStake", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
//
// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCaller) GetValidator(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidator", arg0)

	outstruct := new(struct {
		Status           *big.Int
		DeactivatedTime  *big.Int
		DeactivatedEpoch *big.Int
		ReceivedStake    *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		Auth             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReceivedStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
func (_Contract *ContractCallerSession) GetValidator(arg0 *big.Int) (struct {
	Status           *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
	ReceivedStake    *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	Auth             common.Address
}, error) {
	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCaller) GetValidatorID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorID", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address ) view returns(uint256)
func (_Contract *ContractCallerSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCaller) GetValidatorPubkey(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidatorPubkey", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
func (_Contract *ContractCallerSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCaller) GetWithdrawalRequest(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawalRequest", arg0, arg1, arg2)

	outstruct := new(struct {
		Epoch  *big.Int
		Time   *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_Contract *ContractCallerSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractCaller) GetWrRequests(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWrRequests", delegator, validatorID, offset, limit)

	if err != nil {
		return *new([]SFCWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]SFCWithdrawalRequest)).(*[]SFCWithdrawalRequest)

	return out0, err

}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
}

// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
//
// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
func (_Contract *ContractCallerSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCaller) IsLockedUp(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isLockedUp", delegator, toValidatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
//
// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCaller) IsSlashed(opts *bind.CallOpts, validatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isSlashed", validatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_Contract *ContractCallerSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCaller) LastValidatorID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastValidatorID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_Contract *ContractCallerSession) LastValidatorID() (*big.Int, error) {
	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxDelegatedRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
//
// Solidity: function maxDelegatedRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MaxLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxLockupDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
//
// Solidity: function maxLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MaxLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCaller) MinLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minLockupDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
//
// Solidity: function minLockupDuration() pure returns(uint256)
func (_Contract *ContractCallerSession) MinLockupDuration() (*big.Int, error) {
	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minSelfStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractSession) MinSelfStake() (*big.Int, error) {
	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() pure returns(uint256)
func (_Contract *ContractCallerSession) MinSelfStake() (*big.Int, error) {
	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractCaller) OfflinePenaltyThreshold(opts *bind.CallOpts) (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "offlinePenaltyThreshold")

	outstruct := new(struct {
		BlocksNum *big.Int
		Time      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlocksNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractSession) OfflinePenaltyThreshold() (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
}

// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
//
// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
func (_Contract *ContractCallerSession) OfflinePenaltyThreshold() (struct {
	BlocksNum *big.Int
	Time      *big.Int
}, error) {
	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCaller) PendingRewards(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pendingRewards", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_Contract *ContractCallerSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCaller) RewardsStash(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsStash", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_Contract *ContractCallerSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SlashingRefundRatio(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slashingRefundRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractCaller) StakeTokenizerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeTokenizerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractSession) StakeTokenizerAddress() (common.Address, error) {
	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
}

// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
//
// Solidity: function stakeTokenizerAddress() view returns(address)
func (_Contract *ContractCallerSession) StakeTokenizerAddress() (common.Address, error) {
	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address delegator, uint256 validatorId, uint256 amount, uint256 timestamp)
func (_Contract *ContractCaller) Stakes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Delegator   common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Delegator   common.Address
		ValidatorId *big.Int
		Amount      *big.Int
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidatorId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address delegator, uint256 validatorId, uint256 amount, uint256 timestamp)
func (_Contract *ContractSession) Stakes(arg0 *big.Int) (struct {
	Delegator   common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.Stakes(&_Contract.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
//
// Solidity: function stakes(uint256 ) view returns(address delegator, uint256 validatorId, uint256 amount, uint256 timestamp)
func (_Contract *ContractCallerSession) Stakes(arg0 *big.Int) (struct {
	Delegator   common.Address
	ValidatorId *big.Int
	Amount      *big.Int
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.Stakes(&_Contract.CallOpts, arg0)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCaller) StashedRewardsUntilEpoch(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stashedRewardsUntilEpoch", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalActiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalActiveStake() (*big.Int, error) {
	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCaller) TotalSlashedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSlashedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
//
// Solidity: function totalSlashedStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSlashedStake() (*big.Int, error) {
	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Contract *ContractCallerSession) TotalStake() (*big.Int, error) {
	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCaller) UnlockedRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "unlockedRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
//
// Solidity: function unlockedRewardRatio() pure returns(uint256)
func (_Contract *ContractCallerSession) UnlockedRewardRatio() (*big.Int, error) {
	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
//
// Solidity: function validatorCommission() pure returns(uint256)
func (_Contract *ContractCallerSession) ValidatorCommission() (*big.Int, error) {
	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) ([3]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([3]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([3]byte)).(*[3]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_Contract *ContractCallerSession) Version() ([3]byte, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) WithdrawalPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
}

// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
//
// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) WithdrawalPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) WithdrawalPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
}

// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
//
// Solidity: function withdrawalPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) WithdrawalPeriodTime() (*big.Int, error) {
	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactor) SyncValidator(opts *bind.TransactOpts, validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "_syncValidator", validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_Contract *ContractTransactorSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimRewards", toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createValidator", pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Contract *ContractTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "delegate", toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactorSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", sealedEpoch, _totalSupply, nodeDriver, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
func (_Contract *ContractTransactorSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactor) LockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockStake", toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
//
// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactorSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractTransactor) MintFTM(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintFTM", receiver, amount, justification)
}

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
}

// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
//
// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
func (_Contract *ContractTransactorSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactor) RelockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "relockStake", toValidatorID, lockupDuration, amount)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
//
// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
func (_Contract *ContractTransactorSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) RestakeRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "restakeRewards", toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpoch", offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactorSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactor) StashRewards(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stashRewards", delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_Contract *ContractTransactorSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) UnlockStake(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unlockStake", toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
//
// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractTransactor) UpdateBaseRewardPerSecond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateBaseRewardPerSecond", value)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
}

// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
//
// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractTransactor) UpdateOfflinePenaltyThreshold(opts *bind.TransactOpts, blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateOfflinePenaltyThreshold", blocksNum, time)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
}

// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
//
// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
func (_Contract *ContractTransactorSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactor) UpdateSlashingRefundRatio(opts *bind.TransactOpts, validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSlashingRefundRatio", validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_Contract *ContractTransactorSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
//
// Solidity: function updateStakeTokenizerAddress(address addr) returns()
func (_Contract *ContractTransactor) UpdateStakeTokenizerAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateStakeTokenizerAddress", addr)
}

// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
//
// Solidity: function updateStakeTokenizerAddress(address addr) returns()
func (_Contract *ContractSession) UpdateStakeTokenizerAddress(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakeTokenizerAddress(&_Contract.TransactOpts, addr)
}

// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
//
// Solidity: function updateStakeTokenizerAddress(address addr) returns()
func (_Contract *ContractTransactorSession) UpdateStakeTokenizerAddress(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakeTokenizerAddress(&_Contract.TransactOpts, addr)
}

// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
//
// Solidity: function updateTotalSupply(int256 diff) returns()
func (_Contract *ContractTransactor) UpdateTotalSupply(opts *bind.TransactOpts, diff *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateTotalSupply", diff)
}

// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
//
// Solidity: function updateTotalSupply(int256 diff) returns()
func (_Contract *ContractSession) UpdateTotalSupply(diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTotalSupply(&_Contract.TransactOpts, diff)
}

// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
//
// Solidity: function updateTotalSupply(int256 diff) returns()
func (_Contract *ContractTransactorSession) UpdateTotalSupply(diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateTotalSupply(&_Contract.TransactOpts, diff)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_Contract *ContractTransactorSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
}

// ContractChangedValidatorStatusIterator is returned from FilterChangedValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangedValidatorStatus events raised by the Contract contract.
type ContractChangedValidatorStatusIterator struct {
	Event *ContractChangedValidatorStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractChangedValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractChangedValidatorStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractChangedValidatorStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractChangedValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractChangedValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractChangedValidatorStatus represents a ChangedValidatorStatus event raised by the Contract contract.
type ContractChangedValidatorStatus struct {
	ValidatorID *big.Int
	Status      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedValidatorStatus is a free log retrieval operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) FilterChangedValidatorStatus(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractChangedValidatorStatusIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractChangedValidatorStatusIterator{contract: _Contract.contract, event: "ChangedValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchChangedValidatorStatus is a free log subscription operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) WatchChangedValidatorStatus(opts *bind.WatchOpts, sink chan<- *ContractChangedValidatorStatus, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractChangedValidatorStatus)
				if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedValidatorStatus is a log parse operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_Contract *ContractFilterer) ParseChangedValidatorStatus(log types.Log) (*ContractChangedValidatorStatus, error) {
	event := new(ContractChangedValidatorStatus)
	if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the Contract contract.
type ContractClaimedRewardsIterator struct {
	Event *ContractClaimedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractClaimedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimedRewards represents a ClaimedRewards event raised by the Contract contract.
type ContractClaimedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterClaimedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractClaimedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimedRewardsIterator{contract: _Contract.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimedRewards)
				if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedRewards is a log parse operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseClaimedRewards(log types.Log) (*ContractClaimedRewards, error) {
	event := new(ContractClaimedRewards)
	if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractCreatedValidatorIterator is returned from FilterCreatedValidator and is used to iterate over the raw logs and unpacked data for CreatedValidator events raised by the Contract contract.
type ContractCreatedValidatorIterator struct {
	Event *ContractCreatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractCreatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCreatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractCreatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractCreatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractCreatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractCreatedValidator represents a CreatedValidator event raised by the Contract contract.
type ContractCreatedValidator struct {
	ValidatorID  *big.Int
	Auth         common.Address
	CreatedEpoch *big.Int
	CreatedTime  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedValidator is a free log retrieval operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) FilterCreatedValidator(opts *bind.FilterOpts, validatorID []*big.Int, auth []common.Address) (*ContractCreatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return &ContractCreatedValidatorIterator{contract: _Contract.contract, event: "CreatedValidator", logs: logs, sub: sub}, nil
}

// WatchCreatedValidator is a free log subscription operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) WatchCreatedValidator(opts *bind.WatchOpts, sink chan<- *ContractCreatedValidator, validatorID []*big.Int, auth []common.Address) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractCreatedValidator)
				if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedValidator is a log parse operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_Contract *ContractFilterer) ParseCreatedValidator(log types.Log) (*ContractCreatedValidator, error) {
	event := new(ContractCreatedValidator)
	if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeactivatedValidatorIterator is returned from FilterDeactivatedValidator and is used to iterate over the raw logs and unpacked data for DeactivatedValidator events raised by the Contract contract.
type ContractDeactivatedValidatorIterator struct {
	Event *ContractDeactivatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDeactivatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeactivatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDeactivatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDeactivatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeactivatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeactivatedValidator represents a DeactivatedValidator event raised by the Contract contract.
type ContractDeactivatedValidator struct {
	ValidatorID      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedValidator is a free log retrieval operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) FilterDeactivatedValidator(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractDeactivatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeactivatedValidatorIterator{contract: _Contract.contract, event: "DeactivatedValidator", logs: logs, sub: sub}, nil
}

// WatchDeactivatedValidator is a free log subscription operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) WatchDeactivatedValidator(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedValidator, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeactivatedValidator)
				if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedValidator is a log parse operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_Contract *ContractFilterer) ParseDeactivatedValidator(log types.Log) (*ContractDeactivatedValidator, error) {
	event := new(ContractDeactivatedValidator)
	if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Contract contract.
type ContractDelegatedIterator struct {
	Event *ContractDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegated represents a Delegated event raised by the Contract contract.
type ContractDelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegatedIterator{contract: _Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegated, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegated)
				if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_Contract *ContractFilterer) ParseDelegated(log types.Log) (*ContractDelegated, error) {
	event := new(ContractDelegated)
	if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInflatedFTMIterator is returned from FilterInflatedFTM and is used to iterate over the raw logs and unpacked data for InflatedFTM events raised by the Contract contract.
type ContractInflatedFTMIterator struct {
	Event *ContractInflatedFTM // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInflatedFTMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInflatedFTM)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInflatedFTM)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInflatedFTMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInflatedFTMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInflatedFTM represents a InflatedFTM event raised by the Contract contract.
type ContractInflatedFTM struct {
	Receiver      common.Address
	Amount        *big.Int
	Justification string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInflatedFTM is a free log retrieval operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) FilterInflatedFTM(opts *bind.FilterOpts, receiver []common.Address) (*ContractInflatedFTMIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InflatedFTM", receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractInflatedFTMIterator{contract: _Contract.contract, event: "InflatedFTM", logs: logs, sub: sub}, nil
}

// WatchInflatedFTM is a free log subscription operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) WatchInflatedFTM(opts *bind.WatchOpts, sink chan<- *ContractInflatedFTM, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InflatedFTM", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInflatedFTM)
				if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInflatedFTM is a log parse operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
//
// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
func (_Contract *ContractFilterer) ParseInflatedFTM(log types.Log) (*ContractInflatedFTM, error) {
	event := new(ContractInflatedFTM)
	if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLockedUpStakeIterator is returned from FilterLockedUpStake and is used to iterate over the raw logs and unpacked data for LockedUpStake events raised by the Contract contract.
type ContractLockedUpStakeIterator struct {
	Event *ContractLockedUpStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractLockedUpStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLockedUpStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractLockedUpStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractLockedUpStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLockedUpStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLockedUpStake represents a LockedUpStake event raised by the Contract contract.
type ContractLockedUpStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Duration    *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockedUpStake is a free log retrieval operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) FilterLockedUpStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractLockedUpStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractLockedUpStakeIterator{contract: _Contract.contract, event: "LockedUpStake", logs: logs, sub: sub}, nil
}

// WatchLockedUpStake is a free log subscription operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) WatchLockedUpStake(opts *bind.WatchOpts, sink chan<- *ContractLockedUpStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLockedUpStake)
				if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockedUpStake is a log parse operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
//
// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
func (_Contract *ContractFilterer) ParseLockedUpStake(log types.Log) (*ContractLockedUpStake, error) {
	event := new(ContractLockedUpStake)
	if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRefundedSlashedLegacyDelegationIterator is returned from FilterRefundedSlashedLegacyDelegation and is used to iterate over the raw logs and unpacked data for RefundedSlashedLegacyDelegation events raised by the Contract contract.
type ContractRefundedSlashedLegacyDelegationIterator struct {
	Event *ContractRefundedSlashedLegacyDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRefundedSlashedLegacyDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRefundedSlashedLegacyDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRefundedSlashedLegacyDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRefundedSlashedLegacyDelegation represents a RefundedSlashedLegacyDelegation event raised by the Contract contract.
type ContractRefundedSlashedLegacyDelegation struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefundedSlashedLegacyDelegation is a free log retrieval operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) FilterRefundedSlashedLegacyDelegation(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractRefundedSlashedLegacyDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractRefundedSlashedLegacyDelegationIterator{contract: _Contract.contract, event: "RefundedSlashedLegacyDelegation", logs: logs, sub: sub}, nil
}

// WatchRefundedSlashedLegacyDelegation is a free log subscription operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) WatchRefundedSlashedLegacyDelegation(opts *bind.WatchOpts, sink chan<- *ContractRefundedSlashedLegacyDelegation, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRefundedSlashedLegacyDelegation)
				if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundedSlashedLegacyDelegation is a log parse operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_Contract *ContractFilterer) ParseRefundedSlashedLegacyDelegation(log types.Log) (*ContractRefundedSlashedLegacyDelegation, error) {
	event := new(ContractRefundedSlashedLegacyDelegation)
	if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRestakedRewardsIterator is returned from FilterRestakedRewards and is used to iterate over the raw logs and unpacked data for RestakedRewards events raised by the Contract contract.
type ContractRestakedRewardsIterator struct {
	Event *ContractRestakedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRestakedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRestakedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRestakedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRestakedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRestakedRewards represents a RestakedRewards event raised by the Contract contract.
type ContractRestakedRewards struct {
	Delegator         common.Address
	ToValidatorID     *big.Int
	LockupExtraReward *big.Int
	LockupBaseReward  *big.Int
	UnlockedReward    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRestakedRewards is a free log retrieval operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) FilterRestakedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractRestakedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakedRewardsIterator{contract: _Contract.contract, event: "RestakedRewards", logs: logs, sub: sub}, nil
}

// WatchRestakedRewards is a free log subscription operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) WatchRestakedRewards(opts *bind.WatchOpts, sink chan<- *ContractRestakedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRestakedRewards)
				if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakedRewards is a log parse operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
func (_Contract *ContractFilterer) ParseRestakedRewards(log types.Log) (*ContractRestakedRewards, error) {
	event := new(ContractRestakedRewards)
	if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Contract contract.
type ContractUndelegatedIterator struct {
	Event *ContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUndelegated represents a Undelegated event raised by the Contract contract.
type ContractUndelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUndelegatedIterator{contract: _Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *ContractUndelegated, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUndelegated)
				if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseUndelegated(log types.Log) (*ContractUndelegated, error) {
	event := new(ContractUndelegated)
	if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnlockedStakeIterator is returned from FilterUnlockedStake and is used to iterate over the raw logs and unpacked data for UnlockedStake events raised by the Contract contract.
type ContractUnlockedStakeIterator struct {
	Event *ContractUnlockedStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUnlockedStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnlockedStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUnlockedStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUnlockedStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnlockedStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnlockedStake represents a UnlockedStake event raised by the Contract contract.
type ContractUnlockedStake struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Penalty     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockedStake is a free log retrieval operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) FilterUnlockedStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractUnlockedStakeIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnlockedStakeIterator{contract: _Contract.contract, event: "UnlockedStake", logs: logs, sub: sub}, nil
}

// WatchUnlockedStake is a free log subscription operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) WatchUnlockedStake(opts *bind.WatchOpts, sink chan<- *ContractUnlockedStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnlockedStake)
				if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockedStake is a log parse operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
//
// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
func (_Contract *ContractFilterer) ParseUnlockedStake(log types.Log) (*ContractUnlockedStake, error) {
	event := new(ContractUnlockedStake)
	if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the Contract contract.
type ContractUpdatedBaseRewardPerSecIterator struct {
	Event *ContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedBaseRewardPerSecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedBaseRewardPerSec)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedBaseRewardPerSec)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedBaseRewardPerSecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedBaseRewardPerSecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the Contract contract.
type ContractUpdatedBaseRewardPerSec struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*ContractUpdatedBaseRewardPerSecIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedBaseRewardPerSecIterator{contract: _Contract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
}

// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBaseRewardPerSec) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedBaseRewardPerSec)
				if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
//
// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
func (_Contract *ContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*ContractUpdatedBaseRewardPerSec, error) {
	event := new(ContractUpdatedBaseRewardPerSec)
	if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedOfflinePenaltyThresholdIterator is returned from FilterUpdatedOfflinePenaltyThreshold and is used to iterate over the raw logs and unpacked data for UpdatedOfflinePenaltyThreshold events raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThresholdIterator struct {
	Event *ContractUpdatedOfflinePenaltyThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedOfflinePenaltyThreshold represents a UpdatedOfflinePenaltyThreshold event raised by the Contract contract.
type ContractUpdatedOfflinePenaltyThreshold struct {
	BlocksNum *big.Int
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOfflinePenaltyThreshold is a free log retrieval operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) FilterUpdatedOfflinePenaltyThreshold(opts *bind.FilterOpts) (*ContractUpdatedOfflinePenaltyThresholdIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedOfflinePenaltyThresholdIterator{contract: _Contract.contract, event: "UpdatedOfflinePenaltyThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatedOfflinePenaltyThreshold is a free log subscription operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) WatchUpdatedOfflinePenaltyThreshold(opts *bind.WatchOpts, sink chan<- *ContractUpdatedOfflinePenaltyThreshold) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedOfflinePenaltyThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedOfflinePenaltyThreshold)
				if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedOfflinePenaltyThreshold is a log parse operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
//
// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
func (_Contract *ContractFilterer) ParseUpdatedOfflinePenaltyThreshold(log types.Log) (*ContractUpdatedOfflinePenaltyThreshold, error) {
	event := new(ContractUpdatedOfflinePenaltyThreshold)
	if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdatedSlashingRefundRatioIterator is returned from FilterUpdatedSlashingRefundRatio and is used to iterate over the raw logs and unpacked data for UpdatedSlashingRefundRatio events raised by the Contract contract.
type ContractUpdatedSlashingRefundRatioIterator struct {
	Event *ContractUpdatedSlashingRefundRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpdatedSlashingRefundRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdatedSlashingRefundRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpdatedSlashingRefundRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpdatedSlashingRefundRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdatedSlashingRefundRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdatedSlashingRefundRatio represents a UpdatedSlashingRefundRatio event raised by the Contract contract.
type ContractUpdatedSlashingRefundRatio struct {
	ValidatorID *big.Int
	RefundRatio *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSlashingRefundRatio is a free log retrieval operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) FilterUpdatedSlashingRefundRatio(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdatedSlashingRefundRatioIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdatedSlashingRefundRatioIterator{contract: _Contract.contract, event: "UpdatedSlashingRefundRatio", logs: logs, sub: sub}, nil
}

// WatchUpdatedSlashingRefundRatio is a free log subscription operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) WatchUpdatedSlashingRefundRatio(opts *bind.WatchOpts, sink chan<- *ContractUpdatedSlashingRefundRatio, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdatedSlashingRefundRatio)
				if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedSlashingRefundRatio is a log parse operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_Contract *ContractFilterer) ParseUpdatedSlashingRefundRatio(log types.Log) (*ContractUpdatedSlashingRefundRatio, error) {
	event := new(ContractUpdatedSlashingRefundRatio)
	if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contract contract.
type ContractWithdrawnIterator struct {
	Event *ContractWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdrawn represents a Withdrawn event raised by the Contract contract.
type ContractWithdrawn struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractWithdrawnIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
	event := new(ContractWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
