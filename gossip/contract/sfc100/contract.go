// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfc100

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const oldContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedUpStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"UnlockedStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UpdatedOfflinePenaltyThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getUnlockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeDriver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDelegatedRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minLockupDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offlinePenaltyThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTime\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockedRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateBaseRewardPerSecond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"updateOfflinePenaltyThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStakeDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegationDecrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationLockPeriodEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toStakerID\",\"type\":\"uint256\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakerID\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidUntilEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedMe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dagAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getStakerID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsTotalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"isDelegationLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"isStakeLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersLastID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegationsNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"lockedDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"createDelegation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calcDelegationRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calcValidatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"claimDelegationCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimValidatorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimValidatorCompoundRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"prepareToWithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawStakePartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"prepareToWithdrawDelegationPartial\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"partialWithdrawByRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"}],\"name\":\"lockUpStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toStakerID\",\"type\":\"uint256\"}],\"name\":\"lockUpDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const ContractABI = "[\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"BurntFTM\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"ChangedValidatorStatus\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"ClaimedRewards\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"CreatedValidator\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"DeactivatedValidator\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Delegated\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"receiver\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"string\",\n          \"name\": \"justification\",\n          \"type\": \"string\"\n        }\n      ],\n      \"name\": \"InflatedFTM\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"duration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"LockedUpStake\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"previousOwner\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"newOwner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"OwnershipTransferred\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"RefundedSlashedLegacyDelegation\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"RestakedRewards\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Undelegated\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"penalty\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UnlockedStake\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"value\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedBaseRewardPerSec\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"blocksNum\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"period\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedOfflinePenaltyThreshold\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"refundRatio\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedSlashingRefundRatio\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Withdrawn\",\n      \"type\": \"event\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"syncPubkey\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"name\": \"_syncValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"currentEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"currentSealedEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochSnapshot\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"endTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epochFee\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalBaseRewardWeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalTxRewardWeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"baseRewardPerSecond\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalSupply\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getLockedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getLockupInfo\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"fromEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"endTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"duration\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStashedLockupRewards\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getValidator\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"receivedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"getValidatorID\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getValidatorPubkey\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"\",\n          \"type\": \"bytes\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getWithdrawalRequest\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"time\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"isLockedUp\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"isOwner\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"lastValidatorID\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"minGasPrice\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"owner\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [],\n      \"name\": \"renounceOwnership\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"slashingRefundRatio\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"stakeTokenizerAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stakes\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorId\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"timestamp\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stashedRewardsUntilEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalActiveStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalSlashedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalSupply\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"newOwner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"transferOwnership\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"treasuryAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"version\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bytes3\",\n          \"name\": \"\",\n          \"type\": \"bytes3\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"voteBookAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"sealedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"_totalSupply\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"nodeDriver\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"_c\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"owner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"initialize\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"addr\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"updateStakeTokenizerAddress\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"v\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"updateTreasuryAddress\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"v\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"updateConstsAddress\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"constsAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"v\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"updateVoteBookAddress\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"offlineTime\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"offlineBlocks\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"uptimes\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"originatedTxsFee\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epochGas\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"sealEpoch\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"nextValidatorIDs\",\n          \"type\": \"uint256[]\"\n        }\n      ],\n      \"name\": \"sealEpochValidators\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochValidatorIDs\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"\",\n          \"type\": \"uint256[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochReceivedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedRewardPerToken\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedUptime\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedOriginatedTxsFee\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochOfflineTime\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochOfflineBlocks\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"rewardsStash\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"offset\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"limit\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStakes\",\n      \"outputs\": [\n        {\n          \"components\": [\n            {\n              \"internalType\": \"address\",\n              \"name\": \"delegator\",\n              \"type\": \"address\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"validatorId\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"amount\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"timestamp\",\n              \"type\": \"uint256\"\n            }\n          ],\n          \"internalType\": \"struct SFCState.Stake[]\",\n          \"name\": \"\",\n          \"type\": \"tuple[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"offset\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"limit\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getWrRequests\",\n      \"outputs\": [\n        {\n          \"components\": [\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"epoch\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"time\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"amount\",\n              \"type\": \"uint256\"\n            }\n          ],\n          \"internalType\": \"struct SFCState.WithdrawalRequest[]\",\n          \"name\": \"\",\n          \"type\": \"tuple[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"pubkey\",\n          \"type\": \"bytes\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"setGenesisValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"stake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupFromEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupEndTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"earlyUnlockPenalty\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"rewards\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"setGenesisDelegation\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"pubkey\",\n          \"type\": \"bytes\"\n        }\n      ],\n      \"name\": \"createValidator\",\n      \"outputs\": [],\n      \"payable\": true,\n      \"stateMutability\": \"payable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getSelfStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"delegate\",\n      \"outputs\": [],\n      \"payable\": true,\n      \"stateMutability\": \"payable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"validatorAuth\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"strict\",\n          \"type\": \"bool\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"gas\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"recountVotes\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"undelegate\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"isSlashed\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"withdraw\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"deactivateValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"pendingRewards\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stashRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"claimRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"restakeRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address payable\",\n          \"name\": \"receiver\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"string\",\n          \"name\": \"justification\",\n          \"type\": \"string\"\n        }\n      ],\n      \"name\": \"mintFTM\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"burnFTM\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getUnlockedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"lockStake\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"relockStake\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"unlockStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"refundRatio\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"updateSlashingRefundRatio\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    }\n  ]"


// ContractBin is the compiled bytecode used for deploying new contracts.
var oldContractBin = "0x608060405234801561001057600080fd5b50615caa80620000216000396000f3fe6080604052600436106105fb5760003560e01c80638b0e9f3f1161030e578063c65ee0e11161019b578063de67f215116100e7578063ebdf104c116100a0578063f3ae5b1a1161007a578063f3ae5b1a14611899578063f8b18d8a146115e7578063f99837e6146118c3578063fd5e6dd1146118f3576105fb565b8063ebdf104c146116e6578063ec6a7f1c14611851578063f2fde38b14611866576105fb565b8063de67f21514611581578063df00c922146115b7578063df0e307a146115e7578063df4f49d414611611578063e08d7e661461163b578063e261641a146116b6576105fb565b8063cfd5fa0c11610154578063d9a7c1f91161012e578063d9a7c1f9146114d3578063dc31e1af146114e8578063dc599bb114611518578063dd099bb614611548576105fb565b8063cfd5fa0c1461141c578063cfdbb7cd14611455578063d845fc901461148e576105fb565b8063c65ee0e114611348578063c7be95de14611372578063cb1c4e671461068e578063cc8343aa14611387578063cda5826a146113b9578063cfd47663146113e3576105fb565b8063b1e643391161025a578063bb03a4bd11610213578063c3de580e116101ed578063c3de580e146112f4578063c41b64051461131e578063c4b5dd7e1461068e578063c5f530af14611333576105fb565b8063bb03a4bd146112a9578063bed9d861146112df578063c312eb0714610fe9576105fb565b8063b1e6433914611122578063b5d896271461114c578063b6d9edd5146111b7578063b810e411146111e1578063b82b84271461121a578063b88a37e21461122f576105fb565b806396c7ee46116102c7578063a4b89fab116102a1578063a4b89fab14611036578063a5a470ad14611066578063a7786515146110d4578063a86a056f146110e9576105fb565b806396c7ee4614610f8a5780639fa6dd3514610fe9578063a198d22914611006576105fb565b80638b0e9f3f14610e905780638b1a0d1114610ea55780638cddb01514610ed55780638da5cb5b14610f0e5780638f32d59b14610f3f57806396060e7114610f54576105fb565b80633d0317fe1161048c5780636099ecb2116103d85780636f498663116103915780637cacb1d61161036b5780637cacb1d614610d9e5780637f664d8714610db357806381d9dc7a146106a3578063854873e114610df1576105fb565b80636f49866314610d3b578063715018a614610d745780637667180814610d89576105fb565b80636099ecb214610c5157806360c7e37f1461068e57806361e53fcc14610c8a57806363321e2714610cba578063650acd6614610ced578063670322f814610d02576105fb565b80634feb92f3116104455780635601fe011161041f5780635601fe0114610be257806358f95b8014610c0c5780635e2308d2146109715780635fab23a814610c3c576105fb565b80634feb92f314610b0757806354d77ed21461081957806354fd4d5014610bb0576105fb565b80633d0317fe14610a475780633fee10a814610819578063441a3e7014610a5c5780634bd202dc14610a8c5780634f7c4efb14610aa15780634f864df414610ad1576105fb565b80631d58179c1161054b5780632709275e116105045780632cedb097116104de5780632cedb097146109c557806330fa9929146109f3578063375b3c0a14610a0857806339b80c0014610a1d576105fb565b80632709275e1461097157806328f7314814610986578063295cccba1461099b576105fb565b80631d58179c146108195780631e702f831461082e5780631f2701521461085e578063223fae09146108bb5780632265f2841461092c57806326682c7114610941576105fb565b80630d4955e3116105b857806318160ddd1161059257806318160ddd1461076f57806318f628d41461078457806319ddb54f1461068e5780631d3ac42c146107e9576105fb565b80630d4955e31461070c5780630d7b26091461072157806312622d0e14610736576105fb565b80630135b1db14610600578063019e272914610645578063029859921461068e57806308728f6e146106a357806308c36874146106b85780630962ef79146106e2575b600080fd5b34801561060c57600080fd5b506106336004803603602081101561062357600080fd5b50356001600160a01b0316611979565b60408051918252519081900360200190f35b34801561065157600080fd5b5061068c6004803603608081101561066857600080fd5b508035906020810135906001600160a01b036040820135811691606001351661198b565b005b34801561069a57600080fd5b50610633611a92565b3480156106af57600080fd5b50610633611a98565b3480156106c457600080fd5b5061068c600480360360208110156106db57600080fd5b5035611a9e565b3480156106ee57600080fd5b5061068c6004803603602081101561070557600080fd5b5035611b6a565b34801561071857600080fd5b50610633611c47565b34801561072d57600080fd5b50610633611c4f565b34801561074257600080fd5b506106336004803603604081101561075957600080fd5b506001600160a01b038135169060200135611c56565b34801561077b57600080fd5b50610633611cdf565b34801561079057600080fd5b5061068c60048036036101208110156107a857600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e0810135906101000135611ce5565b3480156107f557600080fd5b506106336004803603604081101561080c57600080fd5b5080359060200135611e45565b34801561082557600080fd5b50610633611fd8565b34801561083a57600080fd5b5061068c6004803603604081101561085157600080fd5b5080359060200135611fe7565b34801561086a57600080fd5b5061089d6004803603606081101561088157600080fd5b506001600160a01b038135169060208101359060400135612085565b60408051938452602084019290925282820152519081900360600190f35b3480156108c757600080fd5b506108f4600480360360408110156108de57600080fd5b506001600160a01b0381351690602001356120b7565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b34801561093857600080fd5b5061063361212b565b34801561094d57600080fd5b5061068c6004803603604081101561096457600080fd5b508035906020013561213d565b34801561097d57600080fd5b5061063361215d565b34801561099257600080fd5b50610633612179565b3480156109a757600080fd5b5061068c600480360360208110156109be57600080fd5b503561217f565b3480156109d157600080fd5b506109da612198565b6040805192835260208301919091528051918290030190f35b3480156109ff57600080fd5b506106336121a2565b348015610a1457600080fd5b506106336121b5565b348015610a2957600080fd5b506108f460048036036020811015610a4057600080fd5b50356121bf565b348015610a5357600080fd5b50610633612201565b348015610a6857600080fd5b5061068c60048036036040811015610a7f57600080fd5b5080359060200135612212565b348015610a9857600080fd5b50610633612556565b348015610aad57600080fd5b5061068c60048036036040811015610ac457600080fd5b508035906020013561255b565b348015610add57600080fd5b5061068c60048036036060811015610af457600080fd5b508035906020810135906040013561268d565b348015610b1357600080fd5b5061068c6004803603610100811015610b2b57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b811115610b5a57600080fd5b820183602082011115610b6c57600080fd5b803590602001918460018302840111600160201b83111715610b8d57600080fd5b9193509150803590602081013590604081013590606081013590608001356129f9565b348015610bbc57600080fd5b50610bc5612a9f565b604080516001600160e81b03199092168252519081900360200190f35b348015610bee57600080fd5b5061063360048036036020811015610c0557600080fd5b5035612aa9565b348015610c1857600080fd5b5061063360048036036040811015610c2f57600080fd5b5080359060200135612adf565b348015610c4857600080fd5b50610633612afc565b348015610c5d57600080fd5b5061063360048036036040811015610c7457600080fd5b506001600160a01b038135169060200135612b02565b348015610c9657600080fd5b5061063360048036036040811015610cad57600080fd5b5080359060200135612b40565b348015610cc657600080fd5b5061063360048036036020811015610cdd57600080fd5b50356001600160a01b0316612b61565b348015610cf957600080fd5b50610633612b7c565b348015610d0e57600080fd5b5061063360048036036040811015610d2557600080fd5b506001600160a01b038135169060200135612b81565b348015610d4757600080fd5b5061063360048036036040811015610d5e57600080fd5b506001600160a01b038135169060200135612bc2565b348015610d8057600080fd5b5061068c612c2c565b348015610d9557600080fd5b50610633612cbd565b348015610daa57600080fd5b50610633612cc6565b348015610dbf57600080fd5b50610ddd60048036036020811015610dd657600080fd5b5035612ccc565b604080519115158252519081900360200190f35b348015610dfd57600080fd5b50610e1b60048036036020811015610e1457600080fd5b5035612cf1565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610e55578181015183820152602001610e3d565b50505050905090810190601f168015610e825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610e9c57600080fd5b50610633612d8c565b348015610eb157600080fd5b5061068c60048036036040811015610ec857600080fd5b5080359060200135612d92565b348015610ee157600080fd5b5061068c60048036036040811015610ef857600080fd5b506001600160a01b038135169060200135612e22565b348015610f1a57600080fd5b50610f23612e70565b604080516001600160a01b039092168252519081900360200190f35b348015610f4b57600080fd5b50610ddd612e7f565b348015610f6057600080fd5b5061089d60048036036060811015610f7757600080fd5b5080359060208101359060400135612e90565b348015610f9657600080fd5b50610fc360048036036040811015610fad57600080fd5b506001600160a01b038135169060200135612ee8565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61068c60048036036020811015610fff57600080fd5b5035612f1a565b34801561101257600080fd5b506106336004803603604081101561102957600080fd5b5080359060200135612f28565b34801561104257600080fd5b5061068c6004803603604081101561105957600080fd5b5080359060200135612f49565b61068c6004803603602081101561107c57600080fd5b810190602081018135600160201b81111561109657600080fd5b8201836020820111156110a857600080fd5b803590602001918460018302840111600160201b831117156110c957600080fd5b509092509050612f71565b3480156110e057600080fd5b50610633613055565b3480156110f557600080fd5b506106336004803603604081101561110c57600080fd5b506001600160a01b03813516906020013561306b565b34801561112e57600080fd5b5061068c6004803603602081101561114557600080fd5b5035613088565b34801561115857600080fd5b506111766004803603602081101561116f57600080fd5b50356130d5565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b3480156111c357600080fd5b5061068c600480360360208110156111da57600080fd5b503561311b565b3480156111ed57600080fd5b5061089d6004803603604081101561120457600080fd5b506001600160a01b0381351690602001356131fb565b34801561122657600080fd5b50610633613227565b34801561123b57600080fd5b506112596004803603602081101561125257600080fd5b503561322e565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561129557818101518382015260200161127d565b505050509050019250505060405180910390f35b3480156112b557600080fd5b5061068c600480360360608110156112cc57600080fd5b5080359060208101359060400135613293565b3480156112eb57600080fd5b5061068c61329e565b34801561130057600080fd5b50610ddd6004803603602081101561131757600080fd5b50356132eb565b34801561132a57600080fd5b5061068c613088565b34801561133f57600080fd5b50610633613302565b34801561135457600080fd5b506106336004803603602081101561136b57600080fd5b5035613311565b34801561137e57600080fd5b50610633613323565b34801561139357600080fd5b5061068c600480360360408110156113aa57600080fd5b50803590602001351515613329565b3480156113c557600080fd5b5061068c600480360360208110156113dc57600080fd5b503561350b565b3480156113ef57600080fd5b506106336004803603604081101561140657600080fd5b506001600160a01b038135169060200135613524565b34801561142857600080fd5b50610ddd6004803603604081101561143f57600080fd5b506001600160a01b038135169060200135613541565b34801561146157600080fd5b50610ddd6004803603604081101561147857600080fd5b506001600160a01b038135169060200135613549565b34801561149a57600080fd5b5061089d600480360360808110156114b157600080fd5b506001600160a01b0381351690602081013590604081013590606001356135b1565b3480156114df57600080fd5b506106336135ef565b3480156114f457600080fd5b506106336004803603604081101561150b57600080fd5b50803590602001356135f5565b34801561152457600080fd5b5061068c6004803603604081101561153b57600080fd5b5080359060200135613616565b34801561155457600080fd5b5061089d6004803603604081101561156b57600080fd5b506001600160a01b03813516906020013561361f565b34801561158d57600080fd5b5061068c600480360360608110156115a457600080fd5b508035906020810135906040013561368b565b3480156115c357600080fd5b50610633600480360360408110156115da57600080fd5b508035906020013561398c565b3480156115f357600080fd5b5061068c6004803603602081101561160a57600080fd5b503561329e565b34801561161d57600080fd5b5061089d6004803603602081101561163457600080fd5b50356139ad565b34801561164757600080fd5b5061068c6004803603602081101561165e57600080fd5b810190602081018135600160201b81111561167857600080fd5b82018360208201111561168a57600080fd5b803590602001918460208302840111600160201b831117156116ab57600080fd5b5090925090506139e3565b3480156116c257600080fd5b50610633600480360360408110156116d957600080fd5b5080359060200135613ac3565b3480156116f257600080fd5b5061068c6004803603608081101561170957600080fd5b810190602081018135600160201b81111561172357600080fd5b82018360208201111561173557600080fd5b803590602001918460208302840111600160201b8311171561175657600080fd5b919390929091602081019035600160201b81111561177357600080fd5b82018360208201111561178557600080fd5b803590602001918460208302840111600160201b831117156117a657600080fd5b919390929091602081019035600160201b8111156117c357600080fd5b8201836020820111156117d557600080fd5b803590602001918460208302840111600160201b831117156117f657600080fd5b919390929091602081019035600160201b81111561181357600080fd5b82018360208201111561182557600080fd5b803590602001918460208302840111600160201b8311171561184657600080fd5b509092509050613ae4565b34801561185d57600080fd5b50610633613cc0565b34801561187257600080fd5b5061068c6004803603602081101561188957600080fd5b50356001600160a01b0316613cca565b3480156118a557600080fd5b5061068c600480360360208110156118bc57600080fd5b5035613d1a565b3480156118cf57600080fd5b5061068c600480360360408110156118e657600080fd5b5080359060200135613d3d565b3480156118ff57600080fd5b5061191d6004803603602081101561191657600080fd5b5035613d46565b604080519a8b5260208b0199909952898901979097526060890195909552608088019390935260a087019190915260c086015260e08501526001600160a01b039081166101008501521661012083015251908190036101400190f35b60696020526000908152604090205481565b600054610100900460ff16806119a457506119a4613e69565b806119b2575060005460ff16155b6119ed5760405162461bcd60e51b815260040180806020018281038252602e815260200180615baa602e913960400191505060405180910390fd5b600054610100900460ff16158015611a18576000805460ff1961ff0019909116610100171660011790555b611a2182613e6f565b6067859055606680546001600160a01b0319166001600160a01b03851617905560768490556755cfe697852e904c6075556103e86078556203f480607955611a67613f60565b6000868152607760205260409020600701558015611a8b576000805461ff00191690555b5050505050565b60015b90565b606b5490565b33611aa7615981565b611ab18284613f64565b60208101518151919250600091611acd9163ffffffff61405816565b9050611af08385611aeb85604001518561405890919063ffffffff16565b6140b2565b6001600160a01b0383166000818152607360209081526040808320888452825291829020805485019055845185820151868401518451928352928201528083019190915290518692917f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26919081900360600190a350505050565b33611b73615981565b611b7d8284613f64565b9050816001600160a01b03166108fc611bbb8360400151611baf8560200151866000015161405890919063ffffffff16565b9063ffffffff61405816565b6040518115909202916000818181858888f19350505050158015611be3573d6000803e3d6000fd5b5082826001600160a01b03167fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4583600001518460200151856040015160405180848152602001838152602001828152602001935050505060405180910390a3505050565b6301e1338090565b6212750090565b6000611c628383613549565b611c9057506001600160a01b0382166000908152607260209081526040808320848452909152902054611cd9565b6001600160a01b038316600081815260736020908152604080832086845282528083205493835260728252808320868452909152902054611cd69163ffffffff6141af16565b90505b92915050565b60765481565b611cee336141f1565b611d295760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b611d34898989614205565b6001600160a01b0389166000908152606f602090815260408083208b84529091529020600201819055611d668761436a565b8515611e3a5786861115611dab5760405162461bcd60e51b815260040180806020018281038252602c815260200180615c4a602c913960400191505060405180910390fd5b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855283529281902086905580518781529182018a9052805192938c9390927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7892908290030190a3505b505050505050505050565b336000818152607360209081526040808320868452909152812090919083611ea2576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b611eac8286613549565b611eed576040805162461bcd60e51b815260206004820152600d60248201526c06e6f74206c6f636b656420757609c1b604482015290519081900360640190fd5b8054841115611f43576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000604482015290519081900360640190fd5b611f4d82866143d1565b506000611f608387878560000154614519565b825486900383556001600160a01b03841660008181526072602090815260408083208b8452825291829020805485900390558151898152908101849052815193945089937fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9929181900390910190a395945050505050565b6000611fe2612b7c565b905090565b611ff0336141f1565b61202b5760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b8061206c576040805162461bcd60e51b815260206004820152600c60248201526b77726f6e672073746174757360a01b604482015290519081900360640190fd5b6120768282614662565b612081826000613329565b5050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6001600160a01b03821660009081526072602090815260408083208484529091528120548190819081908190819081908061210857506000965086955085945084935083925082915081905061211f565b600197508796506000955085945092508591508790505b92959891949750929550565b600061213561478c565b601002905090565b3360009081526069602052604090205461215881848461268d565b505050565b6000606461216961478c565b601e028161217357fe5b04905090565b606d5481565b3360009081526069602052604090205461208181611b6a565b6078546079549091565b60006121ac612201565b606c5403905090565b6000611fe2613302565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b60006064606c546018028161217357fe5b3361221b615981565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526122b7576040805162461bcd60e51b81526020600482015260156024820152741c995c5d595cdd08191bd95cdb89dd08195e1a5cdd605a1b604482015290519081900360640190fd5b602080820151825160008781526068909352604090922060010154909190158015906122f3575060008681526068602052604090206001015482115b15612314575050600084815260686020526040902060018101546002909101545b61231c613227565b8201612326613f60565b1015612372576040805162461bcd60e51b81526020600482015260166024820152751b9bdd08195b9bdd59da081d1a5b59481c185cdcd95960521b604482015290519081900360640190fd5b61237a612b7c565b8101612384612cbd565b10156123d7576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152607160209081526040808320898452825280832088845290915281206002015490612410886132eb565b905060006124328383607a60008d815260200190815260200160002054614798565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116124c2576040805162461bcd60e51b81526020600482015260166024820152751cdd185ad9481a5cc8199d5b1b1e481cdb185cda195960521b604482015290519081900360640190fd5b6001600160a01b0387166108fc6124df858463ffffffff6141af16565b6040518115909202916000818181858888f19350505050158015612507573d6000803e3d6000fd5b508789886001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21866040518082815260200191505060405180910390a4505050505050505050565b600090565b612563612e7f565b6125a2576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6125ab826132eb565b6125fc576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f722069736e277420736c6173686564000000000000000000604482015290519081900360640190fd5b61260461478c565b8111156126425760405162461bcd60e51b8152600401808060200182810382526021815260200180615bd86021913960400191505060405180910390fd5b6000828152607a60209081526040918290208390558151838152915184927f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91792908290030190a25050565b3361269881856143d1565b50600082116126dc576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b6126e68185611c56565b82111561273a576040805162461bcd60e51b815260206004820152601960248201527f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526071602090815260408083208784528252808320868452909152902060020154156127b1576040805162461bcd60e51b81526020600482015260136024820152727772494420616c72656164792065786973747360681b604482015290519081900360640190fd5b6001600160a01b038116600090815260726020908152604080832087845282528083208054869003905560689091529020600301546127f6908363ffffffff6141af16565b600085815260686020526040902060030155606c5461281b908363ffffffff6141af16565b606c5560008481526068602052604090205461284857606d54612844908363ffffffff6141af16565b606d555b600061285385612aa9565b905080156128fa57612863613302565b8110156128b1576040805162461bcd60e51b8152602060048201526017602482015276696e73756666696369656e742073656c662d7374616b6560481b604482015290519081900360640190fd5b6128ba856147fa565b6128f55760405162461bcd60e51b8152600401808060200182810382526029815260200180615c216029913960400191505060405180910390fd5b612905565b612905856001614662565b6001600160a01b03821660009081526071602090815260408083208884528252808320878452909152902060020183905561293e612cbd565b6001600160a01b03831660009081526071602090815260408083208984528252808320888452909152902055612972613f60565b6001600160a01b038316600090815260716020908152604080832089845282528083208884529091528120600101919091556129af908690613329565b8385836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040518082815260200191505060405180910390a45050505050565b612a02336141f1565b612a3d5760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b612a85898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888614842565b606b54881115611e3a57606b889055505050505050505050565b6203330360ec1b90565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b606e5481565b6000612b0c615981565b612b1684846149f1565b805160208201516040830151929350612b3892611baf9163ffffffff61405816565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b6001600160a01b031660009081526069602052604090205490565b600390565b6000612b8d8383613549565b612b9957506000611cd9565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b6000612bcc615981565b506001600160a01b0383166000908152606f6020908152604080832085845282529182902082516060810184528154808252600183015493820184905260029092015493810184905292612b38929091611baf919063ffffffff61405816565b612c34612e7f565b612c73576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b60675460010190565b60675481565b600081815260686020526040812060060154611cd9906001600160a01b031683613549565b606a6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015612d845780601f10612d5957610100808354040283529160200191612d84565b820191906000526020600020905b815481529060010190602001808311612d6757829003601f168201915b505050505081565b606c5481565b612d9a612e7f565b612dd9576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b60798190556078829055604080518381526020810183905281517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34929181900390910190a15050565b612e2c82826143d1565b612081576040805162461bcd60e51b815260206004820152601060248201526f0dcdee8d0d2dcce40e8de40e6e8c2e6d60831b604482015290519081900360640190fd5b6033546001600160a01b031690565b6033546001600160a01b0316331490565b600083815260686020526040812060060154819081908190612ebb906001600160a01b031688612b02565b905080612ed357506000925060019150829050612edf565b60675490935091508190505b93509350939050565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b612f253382346140b2565b50565b60009182526077602090815260408084209284526005909201905290205490565b336000908152607260209081526040808320848452909152902054612081908290849061368b565b612f79613302565b341015612fc7576040805162461bcd60e51b8152602060048201526017602482015276696e73756666696369656e742073656c662d7374616b6560481b604482015290519081900360640190fd5b80613008576040805162461bcd60e51b815260206004820152600c60248201526b656d707479207075626b657960a01b604482015290519081900360640190fd5b6130483383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250614a5f92505050565b61208133606b54346140b2565b6000606461306161478c565b600f028161217357fe5b607060209081526000928352604080842090915290825290205481565b6040805162461bcd60e51b815260206004820152601f60248201527f75736520534643763320756e64656c656761746528292066756e6374696f6e00604482015290519081900360640190fd5b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b613123612e7f565b613162576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6801c985c8903591eb208111156131c0576040805162461bcd60e51b815260206004820152601b60248201527f746f6f206c617267652072657761726420706572207365636f6e640000000000604482015290519081900360640190fd5b60758190556040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b62093a8090565b60008181526077602090815260409182902060060180548351818402810184019094528084526060939283018282801561328757602002820191906000526020600020905b815481526020019060010190808311613273575b50505050509050919050565b61215882848361268d565b6040805162461bcd60e51b815260206004820152601d60248201527f75736520534643763320776974686472617728292066756e6374696f6e000000604482015290519081900360640190fd5b600090815260686020526040902054608016151590565b6a02a055184a310c1260000090565b607a6020526000908152604090205481565b606b5481565b61333282614a8a565b61337d576040805162461bcd60e51b81526020600482015260176024820152761d985b1a59185d1bdc88191bd95cdb89dd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000828152606860205260409020600381015490541561339b575060005b6066546040805163520337df60e11b8152600481018690526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b1580156133ef57600080fd5b505af1158015613403573d6000803e3d6000fd5b5050505081801561341357508015155b15612158576066546000848152606a602052604090819020815163242a6e3f60e01b81526004810187815260248201938452825460026000196001831615610100020190911604604483018190526001600160a01b039095169463242a6e3f948994939091606490910190849080156134cd5780601f106134a2576101008083540402835291602001916134cd565b820191906000526020600020905b8154815290600101906020018083116134b057829003601f168201915b50509350505050600060405180830381600087803b1580156134ee57600080fd5b505af1158015613502573d6000803e3d6000fd5b50505050505050565b3360009081526069602052604090205461208181611a9e565b607260209081526000928352604080842090915290825290205481565b6000611cd683835b6001600160a01b038216600090815260736020908152604080832084845290915281206002015415801590611cd657506001600160a01b03831660009081526073602090815260408083208584529091529020600201546135a8613f60565b11159392505050565b6000806000806135c18888612b02565b9050806135d9575060009250600191508290506135e5565b60675490935091508190505b9450945094915050565b60755481565b60009182526077602090815260408084209284526003909201905290205490565b61208181611a9e565b600080600061362c6159a2565b505050506001600160a01b03919091166000908152607360209081526040808320938352928152908290208251608081018452815481526001820154928101839052600282015493810184905260039091015460609091018190529092565b33816136cc576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b6136d68185613549565b1561371c576040805162461bcd60e51b81526020600482015260116024820152700616c7265616479206c6f636b656420757607c1b604482015290519081900360640190fd5b6137268185611c56565b82111561376d576040805162461bcd60e51b815260206004820152601060248201526f6e6f7420656e6f756768207374616b6560801b604482015290519081900360640190fd5b600084815260686020526040902054156137c7576040805162461bcd60e51b815260206004820152601660248201527576616c696461746f722069736e27742061637469766560501b604482015290519081900360640190fd5b6137cf611c4f565b83101580156137e557506137e1611c47565b8311155b61382b576040805162461bcd60e51b815260206004820152601260248201527134b731b7b93932b1ba10323ab930ba34b7b760711b604482015290519081900360640190fd5b600061383984611baf613f60565b6000868152606860205260409020600601549091506001600160a01b0390811690831681146138c7576001600160a01b03811660009081526073602090815260408083208984529091529020600201548211156138c75760405162461bcd60e51b8152600401808060200182810382526028815260200180615bf96028913960400191505060405180910390fd5b6138d183876143d1565b506001600160a01b03831660009081526073602090815260408083208984529091529020848155613900612cbd565b6001808301919091556002808301859055600383018890556001600160a01b03861660008181526074602090815260408083208d845282528083208381559586018390559490930155825189815291820188905282518a9391927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7892908290030190a350505050505050565b60009182526077602090815260408084209284526002909201905290205490565b600081815260686020526040812060060154819081906139d6906001600160a01b03168561361f565b9250925092509193909250565b6139ec336141f1565b613a275760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b600060776000613a35612cbd565b8152602001908152602001600020905060008090505b82811015613aae576000848483818110613a6157fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c860154909350613a9f91508263ffffffff61405816565b600c8501555050600101613a4b565b50613abd6006820184846159ca565b50505050565b60009182526077602090815260408084209284526004909201905290205490565b613aed336141f1565b613b285760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b600060776000613b36612cbd565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015613b9657602002820191906000526020600020905b815481526020019060010190808311613b82575b50505050509050613c1d82828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250614aa192505050565b613c8c828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a918291850190849080828437600092019190915250614bb092505050565b613c94612cbd565b606755613c9f613f60565b600783015550607554600b820155607654600d909101555050505050505050565b6000611fe2613227565b613cd2612e7f565b613d11576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b612f25816151d0565b336000908152606960205260409020546120818183613d3882612aa9565b61368b565b61208181611b6a565b600080600080600080600080600080613d5d615a15565b5060008b815260686020908152604091829020825160e08101845281548082526001830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a0840152600601546001600160a01b031660c083015260081415613dd7576101008152613df9565b805160801415613dea5760018152613df9565b805160011415613df957600081525b6000613e048d612aa9565b9050816000015182608001518360a0015184604001518560200151856001613e39888a606001516141af90919063ffffffff16565b8960c001518a60c001518393509b509b509b509b509b509b509b509b509b509b5050509193959799509193959799565b303b1590565b600054610100900460ff1680613e885750613e88613e69565b80613e96575060005460ff16155b613ed15760405162461bcd60e51b815260040180806020018281038252602e815260200180615baa602e913960400191505060405180910390fd5b600054610100900460ff16158015613efc576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015612081576000805461ff00191690555050565b4290565b613f6c615981565b613f7683836143d1565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392613fd492611baf9163ffffffff61405816565b905080614017576040805162461bcd60e51b815260206004820152600c60248201526b7a65726f207265776172647360a01b604482015290519081900360640190fd5b6001600160a01b0384166000908152606f60209081526040808320868452909152812081815560018101829055600201556140518161436a565b5092915050565b600082820183811015611cd6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6140bb82614a8a565b614106576040805162461bcd60e51b81526020600482015260176024820152761d985b1a59185d1bdc88191bd95cdb89dd08195e1a5cdd604a1b604482015290519081900360640190fd5b60008281526068602052604090205415614160576040805162461bcd60e51b815260206004820152601660248201527576616c696461746f722069736e27742061637469766560501b604482015290519081900360640190fd5b61416b838383614205565b614174826147fa565b6121585760405162461bcd60e51b8152600401808060200182810382526029815260200180615c216029913960400191505060405180910390fd5b6000611cd683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615271565b6066546001600160a01b0390811691161490565b60008111614248576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b61425283836143d1565b506001600160a01b0383166000908152607260209081526040808320858452909152902054614287908263ffffffff61405816565b6001600160a01b03841660009081526072602090815260408083208684528252808320939093556068905220600301546142c7818363ffffffff61405816565b600084815260686020526040902060030155606c546142ec908363ffffffff61405816565b606c5560008381526068602052604090205461431957606d54614315908363ffffffff61405816565b606d555b614324838215613329565b60408051838152905184916001600160a01b038716917f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb9181900360200190a350505050565b606654604080516366e7ea0f60e01b81523060048201526024810184905290516001600160a01b03909216916366e7ea0f9160448082019260009290919082900301818387803b1580156143bd57600080fd5b505af1158015611a8b573d6000803e3d6000fd5b60006143db615981565b6143e58484615308565b90506143f083615440565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f825282812086825282528290208251606081018452815481526001820154928101929092526002015491810191909152614456908261549b565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260748252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526144d0908261549b565b6001600160a01b03851660009081526074602090815260408083208784528252918290208351815590830151600180830191909155929091015160029091015591505092915050565b6001600160a01b03841660009081526074602090815260408083208684529091528120548190614561908490614555908763ffffffff61550d16565b9063ffffffff61556616565b6001600160a01b0387166000908152607460209081526040808320898452909152812060010154919250906145a2908590614555908863ffffffff61550d16565b905060028104820160006145ba86614555848a61550d565b6001600160a01b038a1660009081526074602090815260408083208c84529091529020549091506145f1908563ffffffff6141af16565b6001600160a01b038a1660009081526074602090815260408083208c845290915290209081556001015461462590846141af565b6001600160a01b038a1660009081526074602090815260408083208c84529091529020600101558681106146565750855b98975050505050505050565b60008281526068602052604090205415801561467d57508015155b156146aa57600082815260686020526040902060030154606d546146a69163ffffffff6141af16565b606d555b60008281526068602052604090205481111561208157600082815260686020526040902081815560020154614752576146e1612cbd565b6000838152606860205260409020600201556146fb613f60565b6000838152606860209081526040918290206001810184905560020154825190815290810192909252805184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792908290030190a25b60408051828152905183917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25050565b670de0b6b3a764000090565b60008215806147ae57506147aa61478c565b8210155b156147bb575060006147f3565b6147e66001611baf6147cb61478c565b614555866147d761478c565b8a91900363ffffffff61550d16565b9050838111156147f35750825b9392505050565b600061482761480761478c565b61455561481261212b565b61481b86612aa9565b9063ffffffff61550d16565b60008381526068602052604090206003015411159050919050565b6001600160a01b038816600090815260696020526040902054156148ad576040805162461bcd60e51b815260206004820152601860248201527f76616c696461746f7220616c7265616479206578697374730000000000000000604482015290519081900360640190fd5b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180546001600160a01b031916909417909355606a8152919020875161491e92890190615a5b565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051808381526020018281526020019250505060405180910390a381156149aa576040805183815260208101839052815189927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47928290030190a25b84156149e75760408051868152905188917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25b5050505050505050565b6149f9615981565b614a01615981565b614a0b8484615308565b6001600160a01b0385166000908152606f602090815260408083208784528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150612b38908261549b565b606b8054600101908190556121588382846000614a7a612cbd565b614a82613f60565b600080614842565b600090815260686020526040902060050154151590565b60005b8351811015611a8b57607854828281518110614abc57fe5b6020026020010151118015614ae65750607954838281518110614adb57fe5b602002602001015110155b15614b2757614b09848281518110614afa57fe5b60200260200101516008614662565b614b27848281518110614b1857fe5b60200260200101516000613329565b828181518110614b3357fe5b6020026020010151856004016000868481518110614b4d57fe5b6020026020010151815260200190815260200160002081905550818181518110614b7357fe5b6020026020010151856005016000868481518110614b8d57fe5b602090810291909101810151825281019190915260400160002055600101614aa4565b614bb8615ac9565b6040518060c001604052808551604051908082528060200260200182016040528015614bee578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015614c26578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060776000614c606001614c54612cbd565b9063ffffffff6141af16565b81526020810191909152604001600020600160808401526007810154909150614c87613f60565b1115614ca1578060070154614c9a613f60565b0360808301525b60005b8551811015614d6c578260800151858281518110614cbe57fe5b6020026020010151858381518110614cd257fe5b60200260200101510281614ce257fe5b0483604001518281518110614cf357fe5b602002602001018181525050614d2d83604001518281518110614d1257fe5b6020026020010151846060015161405890919063ffffffff16565b60608401528351614d5f90859083908110614d4457fe5b60200260200101518460a0015161405890919063ffffffff16565b60a0840152600101614ca4565b5060005b8551811015614e3d578260800151858281518110614d8a57fe5b60200260200101518460800151878481518110614da357fe5b60200260200101518a60000160008b8781518110614dbd57fe5b60200260200101518152602001908152602001600020540281614ddc57fe5b040281614de557fe5b0483600001518281518110614df657fe5b602002602001018181525050614e3083600001518281518110614e1557fe5b6020026020010151846020015161405890919063ffffffff16565b6020840152600101614d70565b5060005b85518110156151a8576000614e79846080015160755486600001518581518110614e6757fe5b602002602001015187602001516155a8565b9050614eb5614ea88560a0015186604001518581518110614e9657fe5b602002602001015187606001516155e9565b829063ffffffff61405816565b90506000878381518110614ec557fe5b6020908102919091018101516000818152606890925260408220600601549092506001600160a01b031690614f0184614efc613055565b615646565b6001600160a01b038316600090815260726020908152604080832087845290915290205490915080156150a857600081614f3b8587612b81565b840281614f4457fe5b049050808303614f52615981565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154614f84908490615663565b9050614f8e615981565b614f99836000615663565b6001600160a01b0388166000908152606f602090815260408083208c84528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150614fee908383615754565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d845282529182902082516060810184528154815294810154918501919091529091015490820152615069908383615754565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156150da57816150cd61478c565b8402816150d657fe5b0490505b808a600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508b898151811061511757fe5b60200260200101518a600301600089815260200190815260200160002054018f6003016000898152602001908152602001600020819055508c898151811061515b57fe5b60200260200101518a600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050614e41565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b0381166152155760405162461bcd60e51b8152600401808060200182810382526026815260200180615b1a6026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b600081848411156153005760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156152c55781810151838201526020016152ad565b50505050905090810190601f1680156152f25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b615310615981565b6001600160a01b03831660009081526070602090815260408083208584529091528120549061533e84615440565b9050600061534c868661576f565b9050818111156153595750805b828110156153645750815b6001600160a01b038616600081815260736020908152604080832089845282528083209383526072825280832089845290915281205482549091906153b090839063ffffffff6141af16565b905060006153c484600001548a898861582e565b90506153ce615981565b6153dc828660030154615663565b90506153ea838b8a8961582e565b91506153f4615981565b6153ff836000615663565b905061540d858c898b61582e565b9250615417615981565b615422846000615663565b905061542f838383615754565b9d9c50505050505050505050505050565b6000818152606860205260408120600201541561549357600082815260686020526040902060020154606754101561547b5750606754612ada565b50600081815260686020526040902060020154612ada565b505060675490565b6154a3615981565b60408051606081019091528251845182916154c4919063ffffffff61405816565b81526020016154e48460200151866020015161405890919063ffffffff16565b81526020016155048460400151866040015161405890919063ffffffff16565b90529392505050565b60008261551c57506000611cd9565b8282028284828161552957fe5b0414611cd65760405162461bcd60e51b8152600401808060200182810382526021815260200180615b696021913960400191505060405180910390fd5b6000611cd683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061589c565b6000826155b757506000612b38565b60006155c9868663ffffffff61550d16565b90506155df83614555838763ffffffff61550d16565b9695505050505050565b6000826155f8575060006147f3565b600061560e83614555878763ffffffff61550d16565b905061563d61561b61478c565b61455561562661215d565b61562e61478c565b8591900363ffffffff61550d16565b95945050505050565b6000611cd661565361478c565b614555858563ffffffff61550d16565b61566b615981565b60405180606001604052806000815260200160008152602001600081525090508160001461572657600061569d61215d565b6156a561478c565b03905060006156c56156b5611c47565b614555848763ffffffff61550d16565b905060006156ee6156d461478c565b614555846156e061215d565b8a910163ffffffff61550d16565b90506157136156fb61478c565b61455561570661215d565b899063ffffffff61550d16565b602085018190529003835250611cd99050565b61574961573161478c565b61455561573c61215d565b869063ffffffff61550d16565b604082015292915050565b61575c615981565b612b38615769858561549b565b8361549b565b6001600160a01b03821660009081526073602090815260408083208484529091528120600101546067546157a4858583615901565b156157b2579150611cd99050565b6157bd858584615901565b6157cc57600092505050611cd9565b808211156157df57600092505050611cd9565b80821015615812576002818301046157f8868683615901565b156158085780600101925061580c565b8091505b506157df565b8061582257600092505050611cd9565b60001901949350505050565b600081831061583f57506000612b38565b60008381526077602081815260408084208885526001908101835281852054878652938352818520898652019091529091205461589161587d61478c565b6145558961481b858763ffffffff6141af16565b979650505050505050565b600081836158eb5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156152c55781810151838201526020016152ad565b5060008385816158f757fe5b0495945050505050565b6001600160a01b03831660009081526073602090815260408083208584529091528120600101548210801590612b3857506001600160a01b03841660009081526073602090815260408083208684529091529020600201546159628361596c565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215615a05579160200282015b82811115615a055782358255916020019190600101906159ea565b50615a11929150615aff565b5090565b6040518060e0016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10615a9c57805160ff1916838001178555615a05565b82800160010185558215615a05579182015b82811115615a05578251825591602001919060010190615aae565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b611a9591905b80821115615a115760008155600101615b0556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616c6c6572206973206e6f7420746865204e6f64654472697665724175746820636f6e7472616374536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65646d757374206265206c657373207468616e206f7220657175616c20746f20312e3076616c696461746f72206c6f636b757020706572696f642077696c6c20656e64206561726c69657276616c696461746f7227732064656c65676174696f6e73206c696d69742069732065786365656465646c6f636b6564207374616b652069732067726561746572207468616e207468652077686f6c65207374616b65a265627a7a7231582068a2eef0a6cf0c5b39dc3f21daab05a006fbd4ba3e09bc65ee4da456a539535664736f6c63430005110032"
var ContractBin = "0x608060405234801561001057600080fd5b5061745480620000216000396000f3fe60806040526004361061044a5760003560e01c8063860c275011610243578063bd14d90711610143578063d5a44f86116100bb578063df00c9221161008a578063e261641a1161006f578063e261641a14610c8d578063e2f8c33614610cad578063f2fde38b14610ccd5761044a565b8063df00c92214610c4d578063e08d7e6614610c6d5761044a565b8063d5a44f8614610bc8578063d96ed50514610bf8578063dc31e1af14610c0d578063de67f21514610c2d5761044a565b8063c7be95de11610112578063cfd47663116100f7578063cfd4766314610b73578063cfdbb7cd14610b93578063d46fa51814610bb35761044a565b8063c7be95de14610b3e578063cc8343aa14610b535761044a565b8063bd14d90714610ac9578063c3de580e14610ae9578063c5f956af14610b09578063c65ee0e114610b1e5761044a565b806396c7ee46116101d6578063a5a470ad116101a5578063b5d896271161018a578063b5d8962714610a49578063b810e41114610a7c578063b88a37e214610a9c5761044a565b8063a5a470ad14610a16578063a86a056f14610a295761044a565b806396c7ee46146109935780639fa6dd35146109c3578063a198d229146109d6578063a2f6e6bc146109f65761044a565b80638cddb015116102125780638cddb0151461091c5780638da5cb5b1461093c5780638f32d59b1461095157806390a6c475146109735761044a565b8063860c2750146108a5578063893675c6146108c55780638b0e9f3f146108da5780638c3c51d8146108ef5761044a565b806354fd4d501161034e578063634b91e3116102e1578063715018a6116102b05780637cacb1d6116102955780637cacb1d614610843578063841e456114610858578063854873e1146108785761044a565b8063715018a614610819578063766718081461082e5761044a565b8063634b91e31461078c578063670322f8146107ac5780636f498663146107cc578063702797e3146107ec5761044a565b8063592fe0c01161031d578063592fe0c0146107175780635fab23a8146107375780636099ecb21461074c57806361e53fcc1461076c5761044a565b806354fd4d5014610695578063550359a0146106b75780635601fe01146106d757806358f95b80146106f75761044a565b80631e702f83116103e157806339b80c00116103b0578063441a3e7011610395578063441a3e70146106355780634f7c4efb146106555780634feb92f3146106755761044a565b806339b80c00146105e25780633fbfd4df146106155761044a565b80631e702f831461055e5780631f2701521461057e57806320c0849d146105ad57806328f73148146105cd5761044a565b806312622d0e1161041d57806312622d0e146104e957806318160ddd1461050957806318f628d41461051e5780631d3ac42c1461053e5761044a565b80630135b1db1461044f57806308c36874146104855780630962ef79146104a75780630e559d82146104c7575b600080fd5b34801561045b57600080fd5b5061046f61046a366004615d2e565b610ced565b60405161047c91906171d4565b60405180910390f35b34801561049157600080fd5b506104a56104a03660046161cb565b610cff565b005b3480156104b357600080fd5b506104a56104c23660046161cb565b610dc8565b3480156104d357600080fd5b506104dc610ed7565b60405161047c9190616edb565b3480156104f557600080fd5b5061046f610504366004615e15565b610ee6565b34801561051557600080fd5b5061046f610f6f565b34801561052a57600080fd5b506104a5610539366004615fac565b610f75565b34801561054a57600080fd5b5061046f610559366004616237565b61109b565b34801561056a57600080fd5b506104a5610579366004616237565b6111cb565b34801561058a57600080fd5b5061059e610599366004615f1b565b611250565b60405161047c93929190617231565b3480156105b957600080fd5b506104a56105c8366004615db4565b611282565b3480156105d957600080fd5b5061046f61139c565b3480156105ee57600080fd5b506106026105fd3660046161cb565b6113a2565b60405161047c97969594939291906172cf565b34801561062157600080fd5b506104a5610630366004616256565b6113e4565b34801561064157600080fd5b506104a5610650366004616237565b611647565b34801561066157600080fd5b506104a5610670366004616237565b611a3e565b34801561068157600080fd5b506104a5610690366004615e4f565b611afd565b3480156106a157600080fd5b506106aa611b84565b60405161047c9190616f95565b3480156106c357600080fd5b506104a56106d2366004615d2e565b611ba9565b3480156106e357600080fd5b5061046f6106f23660046161cb565b611c07565b34801561070357600080fd5b5061046f610712366004616237565b611c3d565b34801561072357600080fd5b506104a56107323660046160a2565b611c5a565b34801561074357600080fd5b5061046f611f02565b34801561075857600080fd5b5061046f610767366004615e15565b611f08565b34801561077857600080fd5b5061046f610787366004616237565b611f46565b34801561079857600080fd5b506104a56107a7366004616237565b611f67565b3480156107b857600080fd5b5061046f6107c7366004615e15565b612112565b3480156107d857600080fd5b5061046f6107e7366004615e15565b612153565b3480156107f857600080fd5b5061080c610807366004615f68565b6121bd565b60405161047c9190616f65565b34801561082557600080fd5b506104a5612284565b34801561083a57600080fd5b5061046f61230a565b34801561084f57600080fd5b5061046f612313565b34801561086457600080fd5b506104a5610873366004615d2e565b612319565b34801561088457600080fd5b506108986108933660046161cb565b612377565b60405161047c9190616fa3565b3480156108b157600080fd5b506104a56108c0366004615d2e565b612430565b3480156108d157600080fd5b506104dc61248e565b3480156108e657600080fd5b5061046f61249d565b3480156108fb57600080fd5b5061090f61090a366004616237565b6124a3565b60405161047c9190616f54565b34801561092857600080fd5b506104a5610937366004615e15565b612584565b34801561094857600080fd5b506104dc6125ae565b34801561095d57600080fd5b506109666125bd565b60405161047c9190616f87565b34801561097f57600080fd5b506104a561098e3660046161cb565b6125ce565b34801561099f57600080fd5b506109b36109ae366004615e15565b6125fe565b60405161047c9493929190617259565b6104a56109d13660046161cb565b612630565b3480156109e257600080fd5b5061046f6109f1366004616237565b61263b565b348015610a0257600080fd5b506104a5610a11366004615d2e565b61265c565b6104a5610a24366004616195565b6126ba565b348015610a3557600080fd5b5061046f610a44366004615e15565b6127c9565b348015610a5557600080fd5b50610a69610a643660046161cb565b6127e6565b60405161047c9796959493929190617267565b348015610a8857600080fd5b5061059e610a97366004615e15565b61282c565b348015610aa857600080fd5b50610abc610ab73660046161cb565b612858565b60405161047c9190616f76565b348015610ad557600080fd5b506104a5610ae43660046162cb565b6128bd565b348015610af557600080fd5b50610966610b043660046161cb565b6128d0565b348015610b1557600080fd5b506104dc6128e7565b348015610b2a57600080fd5b5061046f610b393660046161cb565b6128f6565b348015610b4a57600080fd5b5061046f612908565b348015610b5f57600080fd5b506104a5610b6e366004616207565b61290e565b348015610b7f57600080fd5b5061046f610b8e366004615e15565b612a6f565b348015610b9f57600080fd5b50610966610bae366004615e15565b612a8c565b348015610bbf57600080fd5b506104dc612b22565b348015610bd457600080fd5b50610be8610be33660046161cb565b612b31565b60405161047c9493929190616f1f565b348015610c0457600080fd5b5061046f612b72565b348015610c1957600080fd5b5061046f610c28366004616237565b612b78565b348015610c3957600080fd5b506104a5610c483660046162cb565b612b99565b348015610c5957600080fd5b5061046f610c68366004616237565b612bea565b348015610c7957600080fd5b506104a5610c88366004616060565b612c0b565b348015610c9957600080fd5b5061046f610ca8366004616237565b612d11565b348015610cb957600080fd5b506104a5610cc8366004615d4c565b612d32565b348015610cd957600080fd5b506104a5610ce8366004615d2e565b612de1565b60696020526000908152604090205481565b33610d08615b1c565b610d128284612e0e565b60208101518151919250600091610d2e9163ffffffff612ede16565b9050610d518385610d4c856040015185612ede90919063ffffffff16565b612f03565b6001600160a01b03831660008181526073602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610dba939092909190617231565b60405180910390a350505050565b33610dd1615b1c565b610ddb8284612e0e565b90506000826001600160a01b0316610e188360400151610e0c85602001518660000151612ede90919063ffffffff16565b9063ffffffff612ede16565b604051610e2490616ed0565b60006040518083038185875af1925050503d8060008114610e61576040519150601f19603f3d011682016040523d82523d6000602084013e610e66565b606091505b5050905080610e905760405162461bcd60e51b8152600401610e8790617104565b60405180910390fd5b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610dba93617231565b607b546001600160a01b031681565b6000610ef28383612a8c565b610f2057506001600160a01b0382166000908152607260209081526040808320848452909152902054610f69565b6001600160a01b038316600081815260736020908152604080832086845282528083205493835260728252808320868452909152902054610f669163ffffffff612f8616565b90505b92915050565b60765481565b610f7e33612fc8565b610f9a5760405162461bcd60e51b8152600401610e8790617004565b610fa78989896000612fdc565b6001600160a01b0389166000908152606f602090815260408083208b84529091529020600201819055610fd987613302565b85156110905786861115610fff5760405162461bcd60e51b8152600401610e87906171c4565b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78906110869088908c90617223565b60405180910390a3505b505050505050505050565b3360008181526073602090815260408083208684529091528120909190836110d55760405162461bcd60e51b8152600401610e8790616fc4565b6110df8286612a8c565b6110fb5760405162461bcd60e51b8152600401610e8790617074565b805484111561111c5760405162461bcd60e51b8152600401610e8790617134565b6111268286613399565b6111425760405162461bcd60e51b8152600401610e8790617144565b61114c828661344f565b50600061115f838787856000015461361a565b825486900383559050611175838783600161374b565b61117e816139b3565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b987846040516111ba929190617223565b60405180910390a395945050505050565b6111d433612fc8565b6111f05760405162461bcd60e51b8152600401610e8790617004565b8061120d5760405162461bcd60e51b8152600401610e8790617044565b6112178282613a21565b61122282600061290e565b6000828152606860205260408120600601546001600160a01b03169061124b9082908190613b40565b505050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6082546040516000916001600160a01b03169083906112a79088908890602401616ee9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb00000000000000000000000000000000000000000000000000000000179052516113289190616ec4565b60006040518083038160008787f1925050503d8060008114611366576040519150601f19603f3d011682016040523d82523d6000602084013e61136b565b606091505b505090508080611379575082155b6113955760405162461bcd60e51b8152600401610e8790617174565b5050505050565b606d5481565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b600054610100900460ff16806113fd57506113fd613c67565b8061140b575060005460ff16155b6114275760405162461bcd60e51b8152600401610e87906170c4565b600054610100900460ff1615801561148d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61149682613c6d565b6067869055606680546001600160a01b038087167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255608180549286169290911691909117905560768590556114f0613daf565b607e556114fb613db8565b60008781526077602090815260408083206007019390935582516080810184528281529081018281529281018281526060820183815260838054600181018255945291517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a863600490940293840180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0390921691909117905592517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a86483015591517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a86582015590517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a86690910155801561163f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b505050505050565b33611650615b1c565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526116bf5760405162461bcd60e51b8152600401610e8790617114565b6116c98285613399565b6116e55760405162461bcd60e51b8152600401610e8790617144565b60208082015182516000878152606890935260409092206001015490919015801590611721575060008681526068602052604090206001015482115b15611742575050600084815260686020526040902060018101546002909101545b608160009054906101000a90046001600160a01b03166001600160a01b031663b82b84276040518163ffffffff1660e01b815260040160206040518083038186803b15801561179057600080fd5b505afa1580156117a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117c891908101906161e9565b82016117d2613db8565b10156117f05760405162461bcd60e51b8152600401610e8790617054565b608160009054906101000a90046001600160a01b03166001600160a01b031663650acd666040518163ffffffff1660e01b815260040160206040518083038186803b15801561183e57600080fd5b505afa158015611852573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061187691908101906161e9565b810161188061230a565b101561189e5760405162461bcd60e51b8152600401610e8790617194565b6001600160a01b03841660009081526071602090815260408083208984528252808320888452909152812060020154906118d7886128d0565b905060006118f98383607a60008d815260200190815260200160002054613dbc565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e805482019055905080831161195b5760405162461bcd60e51b8152600401610e87906171b4565b60006001600160a01b038816611977858463ffffffff612f8616565b60405161198390616ed0565b60006040518083038185875af1925050503d80600081146119c0576040519150601f19603f3d011682016040523d82523d6000602084013e6119c5565b606091505b50509050806119e65760405162461bcd60e51b8152600401610e8790617104565b6119ef826139b3565b888a896001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a2187604051611a2a91906171d4565b60405180910390a450505050505050505050565b611a466125bd565b611a625760405162461bcd60e51b8152600401610e87906170b4565b611a6b826128d0565b611a875760405162461bcd60e51b8152600401610e8790616fb4565b611a8f613e1e565b811115611aae5760405162461bcd60e51b8152600401610e87906170d4565b6000828152607a6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91790611af19084906171d4565b60405180910390a25050565b611b0633612fc8565b611b225760405162461bcd60e51b8152600401610e8790617004565b611b6a898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888613e2a565b606b5488111561109057606b889055505050505050505050565b7f33303400000000000000000000000000000000000000000000000000000000005b90565b611bb16125bd565b611bcd5760405162461bcd60e51b8152600401610e87906170b4565b608280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b611c6333612fc8565b611c7f5760405162461bcd60e51b8152600401610e8790617004565b600060776000611c8d61230a565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015611ced57602002820191906000526020600020905b815481526020019060010190808311611cd9575b50505050509050611d7482828d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250613fb892505050565b60675460009081526077602052604090206007810154600190611d95613db8565b1115611dac578160070154611da8613db8565b0390505b611e2e818584868d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506141cc92505050565b611e3881866149e7565b5050611e4261230a565b606755611e4d613db8565b6007830155608154604080517fd9a7c1f900000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163d9a7c1f991600480820192602092909190829003018186803b158015611eb057600080fd5b505afa158015611ec4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611ee891908101906161e9565b600b83015550607654600d90910155505050505050505050565b606e5481565b6000611f12615b1c565b611f1c8484614b78565b805160208201516040830151929350611f3e92610e0c9163ffffffff612ede16565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b33611f72818461344f565b5060008211611f935760405162461bcd60e51b8152600401610e8790616fc4565b611f9d8184610ee6565b821115611fbc5760405162461bcd60e51b8152600401610e87906170e4565b611fc68184613399565b611fe25760405162461bcd60e51b8152600401610e8790617144565b6001600160a01b03811660009081526085602090815260408083208684529091529020805460018082019092559061201f9083908690869061374b565b6001600160a01b03821660009081526071602090815260408083208784528252808320848452909152902060020183905561205861230a565b6001600160a01b0383166000908152607160209081526040808320888452825280832085845290915290205561208c613db8565b6001600160a01b038316600090815260716020908152604080832088845282528083208584529091528120600101919091556120c990859061290e565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d578660405161210491906171d4565b60405180910390a450505050565b600061211e8383612a8c565b61212a57506000610f69565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b600061215d615b1c565b506001600160a01b0383166000908152606f6020908152604080832085845282529182902082516060810184528154808252600183015493820184905260029092015493810184905292611f3e929091610e0c919063ffffffff612ede16565b606080826040519080825280602002602001820160405280156121fa57816020015b6121e7615b1c565b8152602001906001900390816121df5790505b50905060005b8381101561227a576001600160a01b0387166000908152607160209081526040808320898452825280832088850184528252918290208251606081018452815481526001820154928101929092526002015491810191909152825183908390811061226757fe5b6020908102919091010152600101612200565b5095945050505050565b61228c6125bd565b6122a85760405162461bcd60e51b8152600401610e87906170b4565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60675460010190565b60675481565b6123216125bd565b61233d5760405162461bcd60e51b8152600401610e87906170b4565b607f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156124285780601f106123fd57610100808354040283529160200191612428565b820191906000526020600020905b81548152906001019060200180831161240b57829003601f168201915b505050505081565b6124386125bd565b6124545760405162461bcd60e51b8152600401610e87906170b4565b608180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6082546001600160a01b031681565b606c5481565b6083546040805183815260208085028201019091526060919082908480156124e557816020015b6124d2615b3d565b8152602001906001900390816124ca5790505b50905060005b8481101561257b5782818701106125015761257b565b60838187018154811061251057fe5b600091825260209182902060408051608081018252600490930290910180546001600160a01b031683526001810154938301939093526002830154908201526003909101546060820152825183908390811061256857fe5b60209081029190910101526001016124eb565b50949350505050565b61258e828261344f565b6125aa5760405162461bcd60e51b8152600401610e8790617014565b5050565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b6125d66125bd565b6125f25760405162461bcd60e51b8152600401610e87906170b4565b6125fb816139b3565b50565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6125fb338234612f03565b60009182526077602090815260408084209284526005909201905290205490565b6126646125bd565b6126805760405162461bcd60e51b8152600401610e87906170b4565b607b80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b608160009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b15801561270857600080fd5b505afa15801561271c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061274091908101906161e9565b34101561275f5760405162461bcd60e51b8152600401610e87906171a4565b8061277c5760405162461bcd60e51b8152600401610e87906170f4565b6127bc3383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250614be692505050565b6125aa33606b5434612f03565b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b6000818152607760209081526040918290206006018054835181840281018401909452808452606093928301828280156128b157602002820191906000526020600020905b81548152602001906001019080831161289d575b50505050509050919050565b336128ca81858585614c11565b50505050565b600090815260686020526040902054608016151590565b607f546001600160a01b031681565b607a6020526000908152604090205481565b606b5481565b61291782614eea565b6129335760405162461bcd60e51b8152600401610e8790617184565b60008281526068602052604090206003810154905415612951575060005b6066546040517fa4066fbe0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a4066fbe9061299c9086908590600401617223565b600060405180830381600087803b1580156129b657600080fd5b505af11580156129ca573d6000803e3d6000fd5b505050508180156129da57508015155b1561124b576066546000848152606a60205260409081902090517f242a6e3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163242a6e3f91612a38918791906004016171e2565b600060405180830381600087803b158015612a5257600080fd5b505af1158015612a66573d6000803e3d6000fd5b50505050505050565b607260209081526000928352604080842090915290825290205481565b6001600160a01b038216600090815260736020908152604080832084845290915281206002015415801590612ae357506001600160a01b038316600090815260736020908152604080832085845290915290205415155b8015610f6657506001600160a01b0383166000908152607360209081526040808320858452909152902060020154612b19613db8565b11159392505050565b6081546001600160a01b031690565b60838181548110612b3e57fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b039092169350919084565b607e5481565b60009182526077602090815260408084209284526003909201905290205490565b3381612bb75760405162461bcd60e51b8152600401610e8790616fc4565b612bc18185612a8c565b15612bde5760405162461bcd60e51b8152600401610e8790616fe4565b6128ca81858585614c11565b60009182526077602090815260408084209284526002909201905290205490565b612c1433612fc8565b612c305760405162461bcd60e51b8152600401610e8790617004565b600060776000612c3e61230a565b8152602001908152602001600020905060008090505b82811015612cb7576000848483818110612c6a57fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c860154909350612ca891508263ffffffff612ede16565b600c8501555050600101612c54565b50612cc6600682018484615b6e565b50606654607e546040517f07aaf3440000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916307aaf34491612a38916004016171d4565b60009182526077602090815260408084209284526004909201905290205490565b612d3a6125bd565b612d565760405162461bcd60e51b8152600401610e87906170b4565b612d5f83613302565b6040516001600160a01b0385169084156108fc029085906000818181858888f19350505050158015612d95573d6000803e3d6000fd5b50836001600160a01b03167f9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad848484604051612dd393929190617202565b60405180910390a250505050565b612de96125bd565b612e055760405162461bcd60e51b8152600401610e87906170b4565b6125fb81614f01565b612e16615b1c565b612e20838361344f565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392612e7e92610e0c9163ffffffff612ede16565b905080612e9d5760405162461bcd60e51b8152600401610e8790617094565b6001600160a01b0384166000908152606f6020908152604080832086845290915281208181556001810182905560020155612ed781613302565b5092915050565b600082820183811015610f665760405162461bcd60e51b8152600401610e8790616ff4565b612f0c82614eea565b612f285760405162461bcd60e51b8152600401610e8790617184565b60008281526068602052604090205415612f545760405162461bcd60e51b8152600401610e8790617034565b612f618383836001612fdc565b612f6a82614f9b565b61124b5760405162461bcd60e51b8152600401610e8790617164565b6000610f6683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061506f565b6066546001600160a01b0390811691161490565b60008211612ffc5760405162461bcd60e51b8152600401610e8790616fc4565b613006848461344f565b506001600160a01b03841660009081526084602090815260408083208684529091529020548061315157608380546001600160a01b0387811660008181526084602090815260408083208b8452825280832086905580516080810182529384529083018a815290830189815242606085019081526001870188559690925291517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a863600490950294850180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169190941617909255517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a864830155517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a86582015590517f1397b88f412a83a7f1c0d834c533e486ff1f24f42a31819e91b624931060a866909101556131c8565b613182836083838154811061316257fe5b906000526020600020906004020160020154612ede90919063ffffffff16565b6083828154811061318f57fe5b90600052602060002090600402016002018190555042608382815481106131b257fe5b9060005260206000209060040201600301819055505b6001600160a01b03851660009081526072602090815260408083208784529091529020546131fc908463ffffffff612ede16565b6001600160a01b038616600090815260726020908152604080832088845282528083209390935560689052206003015461323c818563ffffffff612ede16565b600086815260686020526040902060030155606c54613261908563ffffffff612ede16565b606c5560008581526068602052604090205461328e57606d5461328a908563ffffffff612ede16565b606d555b61329985821561290e565b84866001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb866040516132d391906171d4565b60405180910390a360008581526068602052604090206006015461163f9087906001600160a01b031685613b40565b6066546040517f66e7ea0f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906366e7ea0f9061334d9030908590600401616f04565b600060405180830381600087803b15801561336757600080fd5b505af115801561337b573d6000803e3d6000fd5b5050607654613393925090508263ffffffff612ede16565b60765550565b607b546000906001600160a01b03166133b457506001610f69565b607b546040517f21d585c30000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906321d585c3906133ff9086908690600401616f04565b60206040518083038186803b15801561341757600080fd5b505afa15801561342b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f669190810190616177565b6000613459615b1c565b613463848461509b565b905061346e836151d3565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f8252828120868252825282902082516060810184528154815260018201549281019290925260020154918101919091526134d4908261522e565b6001600160a01b0385166000818152606f602090815260408083208884528252808320855181558583015160018083019190915595820151600291820155938352607482528083208884528252918290208251606081018452815481529481015491850191909152909101549082015261354e908261522e565b6001600160a01b0385166000908152607460209081526040808320878452825291829020835181559083015160018201559101516002909101556135928484612a8c565b6135f5576001600160a01b0384166000818152607360209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260748352818420888552909252822082815592830182905591909101555b60208101511515806136075750805115155b80611f3e57506040015115159392505050565b6001600160a01b03841660009081526074602090815260408083208684529091528120548190613662908490613656908763ffffffff6152a016565b9063ffffffff6152da16565b6001600160a01b0387166000908152607460209081526040808320898452909152812060010154919250906136a3908590613656908863ffffffff6152a016565b6001600160a01b03881660009081526074602090815260408083208a8452909152902054909150600282048301906136db9084612f86565b6001600160a01b03891660009081526074602090815260408083208b845290915290209081556001015461370f9083612f86565b6001600160a01b03891660009081526074602090815260408083208b84529091529020600101558581106137405750845b979650505050505050565b6001600160a01b0384166000908152608460209081526040808320868452909152902054608380546137a39185918490811061378357fe5b906000526020600020906004020160020154612f8690919063ffffffff16565b608382815481106137b057fe5b906000526020600020906004020160020181905550608381815481106137d257fe5b906000526020600020906004020160020154600014156137f5576137f58161531c565b6001600160a01b0385166000908152607260209081526040808320878452825280832080548790039055606890915290206003015461383a908463ffffffff612f8616565b600085815260686020526040902060030155606c5461385f908463ffffffff612f8616565b606c5560008481526068602052604090205461388c57606d54613888908463ffffffff612f8616565b606d555b600061389785611c07565b905080156139815760008581526068602052604090205461397c57608160009054906101000a90046001600160a01b03166001600160a01b031663c5f530af6040518163ffffffff1660e01b815260040160206040518083038186803b15801561390057600080fd5b505afa158015613914573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061393891908101906161e9565b8110156139575760405162461bcd60e51b8152600401610e87906171a4565b61396085614f9b565b61397c5760405162461bcd60e51b8152600401610e8790617164565b61398c565b61398c856001613a21565b60008581526068602052604090206006015461163f9087906001600160a01b031685613b40565b80156125fb5760405160009082156108fc0290839083818181858288f193505050501580156139e6573d6000803e3d6000fd5b507f8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae881604051613a1691906171d4565b60405180910390a150565b600082815260686020526040902054158015613a3c57508015155b15613a6957600082815260686020526040902060030154606d54613a659163ffffffff612f8616565b606d555b6000828152606860205260409020548111156125aa57600082815260686020526040902081815560020154613b1057613aa061230a565b600083815260686020526040902060020155613aba613db8565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792613b0792909190617223565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e82604051611af191906171d4565b6082546001600160a01b03161561124b576082546040516000916001600160a01b031690627a120090613b799087908790602401616ee9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4a7702bb0000000000000000000000000000000000000000000000000000000017905251613bfa9190616ec4565b60006040518083038160008787f1925050503d8060008114613c38576040519150601f19603f3d011682016040523d82523d6000602084013e613c3d565b606091505b505090508080613c4b575081155b6128ca5760405162461bcd60e51b8152600401610e8790617174565b303b1590565b600054610100900460ff1680613c865750613c86613c67565b80613c94575060005460ff16155b613cb05760405162461bcd60e51b8152600401610e87906170c4565b600054610100900460ff16158015613d1657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a380156125aa57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555050565b64174876e80090565b4290565b6000821580613dd25750613dce613e1e565b8210155b15613ddf57506000613e17565b613e0a6001610e0c613def613e1e565b61365686613dfb613e1e565b8a91900363ffffffff6152a016565b905083811115613e175750825b9392505050565b670de0b6b3a764000090565b6001600160a01b03881660009081526069602052604090205415613e605760405162461bcd60e51b8152600401610e8790617064565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909417909355606a81529190208751613ee992890190615bb5565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051613f26929190617223565b60405180910390a38115613f6f57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613f66929190617223565b60405180910390a25b8415613fae57867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e86604051613fa591906171d4565b60405180910390a25b5050505050505050565b60005b835181101561139557608160009054906101000a90046001600160a01b03166001600160a01b0316635a68f01a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561401257600080fd5b505afa158015614026573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061404a91908101906161e9565b82828151811061405657fe5b60200260200101511180156141025750608160009054906101000a90046001600160a01b03166001600160a01b031662cc7f836040518163ffffffff1660e01b815260040160206040518083038186803b1580156140b357600080fd5b505afa1580156140c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506140eb91908101906161e9565b8382815181106140f757fe5b602002602001015110155b156141435761412584828151811061411657fe5b60200260200101516008613a21565b61414384828151811061413457fe5b6020026020010151600061290e565b82818151811061414f57fe5b602002602001015185600401600086848151811061416957fe5b602002602001015181526020019081526020016000208190555081818151811061418f57fe5b60200260200101518560050160008684815181106141a957fe5b602090810291909101810151825281019190915260400160002055600101613fbb565b6141d4615c23565b6040518060a00160405280855160405190808252806020026020018201604052801561420a578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015614242578160200160208202803883390190505b508152602001600081526020016000815250905060008090505b845181101561435d57600086600301600087848151811061427957fe5b602002602001015181526020019081526020016000205490506000809050818584815181106142a457fe5b602002602001015111156142cb57818584815181106142bf57fe5b60200260200101510390505b898684815181106142d857fe5b60200260200101518202816142e957fe5b04846040015184815181106142fa57fe5b6020026020010181815250506143348460400151848151811061431957fe5b60200260200101518560600151612ede90919063ffffffff16565b6060850152608084015161434e908263ffffffff612ede16565b6080850152505060010161425c565b5060005b8451811015614426578784828151811061437757fe5b60200260200101518986848151811061438c57fe5b60200260200101518a60000160008a87815181106143a657fe5b602002602001015181526020019081526020016000205402816143c557fe5b0402816143ce57fe5b04826000015182815181106143df57fe5b602002602001018181525050614419826000015182815181106143fe57fe5b60200260200101518360200151612ede90919063ffffffff16565b6020830152600101614361565b5060005b845181101561487d5760006144df89608160009054906101000a90046001600160a01b03166001600160a01b031663d9a7c1f96040518163ffffffff1660e01b815260040160206040518083038186803b15801561448757600080fd5b505afa15801561449b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506144bf91908101906161e9565b85518051869081106144cd57fe5b602002602001015186602001516154be565b905061451b61450e8460800151856040015185815181106144fc57fe5b602002602001015186606001516154ff565b829063ffffffff612ede16565b9050600086838151811061452b57fe5b60209081029190910181015160008181526068835260408082206006015460815482517fa778651500000000000000000000000000000000000000000000000000000000815292519496506001600160a01b039182169593946145ed948994929093169263a77865159260048082019391829003018186803b1580156145b057600080fd5b505afa1580156145c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506145e891908101906161e9565b615670565b6001600160a01b03831660009081526072602090815260408083208784529091529020549091508015614794576000816146278587612112565b84028161463057fe5b04905080830361463e615b1c565b6001600160a01b03861660009081526073602090815260408083208a845290915290206003015461467090849061568d565b905061467a615b1c565b61468583600061568d565b6001600160a01b0388166000908152606f602090815260408083208c845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506146da908383615866565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d845282529182902082516060810184528154815294810154918501919091529091015490820152614755908383615866565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156147c657816147b9613e1e565b8402816147c257fe5b0490505b808e600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508a898151811061480357fe5b60200260200101518f6003016000898152602001908152602001600020819055508b898151811061483057fe5b60200260200101518e600201600089815260200190815260200160002054018f6002016000898152602001908152602001600020819055505050505050505050808060010191505061442a565b50608081015160088701819055602082015160098801556060820151600a88015560765411156148bb576008860154607680549190910390556148c1565b60006076555b607f546001600160a01b031615612a665760006148dc613e1e565b608160009054906101000a90046001600160a01b03166001600160a01b03166394c3e9146040518163ffffffff1660e01b815260040160206040518083038186803b15801561492a57600080fd5b505afa15801561493e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061496291908101906161e9565b8360800151028161496f57fe5b04905061497b81613302565b607f546040516001600160a01b0390911690829061499890616ed0565b60006040518083038185875af1925050503d80600081146149d5576040519150601f19603f3d011682016040523d82523d6000602084013e6149da565b606091505b5050505050505050505050565b608154604080517f3a3ef66c00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691633a3ef66c916004808301926020929190829003018186803b158015614a4557600080fd5b505afa158015614a59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614a7d91908101906161e9565b83026001019050600081614a8f613e1e565b840281614a9857fe5b0490506000608160009054906101000a90046001600160a01b03166001600160a01b0316632c8c36a56040518163ffffffff1660e01b815260040160206040518083038186803b158015614aeb57600080fd5b505afa158015614aff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614b2391908101906161e9565b9050808501614b30613e1e565b82028387020181614b3d57fe5b049150614b4982615881565b91506000614b55613e1e565b83607e540281614b6157fe5b049050614b6d816158ef565b607e55505050505050565b614b80615b1c565b614b88615b1c565b614b92848461509b565b6001600160a01b0385166000908152606f602090815260408083208784528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150611f3e908261522e565b606b80546001019081905561124b8382846000614c0161230a565b614c09613db8565b600080613e2a565b614c1b8484610ee6565b811115614c3a5760405162461bcd60e51b8152600401610e8790617154565b60008381526068602052604090205415614c665760405162461bcd60e51b8152600401610e8790617034565b608160009054906101000a90046001600160a01b03166001600160a01b0316630d7b26096040518163ffffffff1660e01b815260040160206040518083038186803b158015614cb457600080fd5b505afa158015614cc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614cec91908101906161e9565b8210158015614d805750608160009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b158015614d4457600080fd5b505afa158015614d58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614d7c91908101906161e9565b8211155b614d9c5760405162461bcd60e51b8152600401610e8790617024565b6000614daa83610e0c613db8565b6000858152606860205260409020600601549091506001600160a01b039081169086168114614e19576001600160a01b0381166000908152607360209081526040808320888452909152902060020154821115614e195760405162461bcd60e51b8152600401610e8790617124565b614e23868661344f565b506001600160a01b038616600090815260736020908152604080832088845290915290206003810154851015614e6b5760405162461bcd60e51b8152600401610e87906170a4565b8054614e7d908563ffffffff612ede16565b8155614e8761230a565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7890614ed99089908990617223565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b6001600160a01b038116614f275760405162461bcd60e51b8152600401610e8790616fd4565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000615054614fa8613e1e565b608154604080517f2265f2840000000000000000000000000000000000000000000000000000000081529051613656926001600160a01b031691632265f284916004808301926020929190829003018186803b15801561500757600080fd5b505afa15801561501b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061503f91908101906161e9565b61504886611c07565b9063ffffffff6152a016565b60008381526068602052604090206003015411159050919050565b600081848411156150935760405162461bcd60e51b8152600401610e879190616fa3565b505050900390565b6150a3615b1c565b6001600160a01b0383166000908152607060209081526040808320858452909152812054906150d1846151d3565b905060006150df8686615925565b9050818111156150ec5750805b828110156150f75750815b6001600160a01b0386166000818152607360209081526040808320898452825280832093835260728252808320898452909152812054825490919061514390839063ffffffff612f8616565b9050600061515784600001548a8988615a02565b9050615161615b1c565b61516f82866003015461568d565b905061517d838b8a89615a02565b9150615187615b1c565b61519283600061568d565b90506151a0858c898b615a02565b92506151aa615b1c565b6151b584600061568d565b90506151c2838383615866565b9d9c50505050505050505050505050565b6000818152606860205260408120600201541561522657600082815260686020526040902060020154606754101561520e5750606754611c38565b50600081815260686020526040902060020154611c38565b505060675490565b615236615b1c565b6040805160608101909152825184518291615257919063ffffffff612ede16565b815260200161527784602001518660200151612ede90919063ffffffff16565b815260200161529784604001518660400151612ede90919063ffffffff16565b90529392505050565b6000826152af57506000610f69565b828202828482816152bc57fe5b0414610f665760405162461bcd60e51b8152600401610e8790617084565b6000610f6683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250615a65565b60835460009061533390600163ffffffff612f8616565b9050808214615442576083818154811061534957fe5b90600052602060002090600402016083838154811061536457fe5b60009182526020909120825460049092020180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0390921691909117815560018083015490820155600280830154908201556003918201549101556153d0615b3d565b608382815481106153dd57fe5b6000918252602080832060408051608081018252600490940290910180546001600160a01b0316808552600182015485850190815260028301548685015260039092015460609095019490945292845260848252808420925184529190529020839055505b608380548061544d57fe5b60008281526020812060047fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9093019283020180547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560018101829055600281018290556003015590555050565b6000826154cd57506000611f3e565b60006154df868663ffffffff6152a016565b90506154f583613656838763ffffffff6152a016565b9695505050505050565b60008261550e57506000613e17565b600061552483613656878763ffffffff6152a016565b9050615667615531613e1e565b608154604080517f94c3e9140000000000000000000000000000000000000000000000000000000081529051613656926001600160a01b0316916394c3e914916004808301926020929190829003018186803b15801561559057600080fd5b505afa1580156155a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506155c891908101906161e9565b608160009054906101000a90046001600160a01b03166001600160a01b031663c74dd6216040518163ffffffff1660e01b815260040160206040518083038186803b15801561561657600080fd5b505afa15801561562a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061564e91908101906161e9565b615656613e1e565b0303846152a090919063ffffffff16565b95945050505050565b6000610f6661567d613e1e565b613656858563ffffffff6152a016565b615695615b1c565b60405180606001604052806000815260200160008152602001600081525090506000608160009054906101000a90046001600160a01b03166001600160a01b0316635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b15801561570557600080fd5b505afa158015615719573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061573d91908101906161e9565b9050821561583f57600081615750613e1e565b03905060006157ee608160009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b1580156157a657600080fd5b505afa1580156157ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506157de91908101906161e9565b613656848863ffffffff6152a016565b9050600061580f6157fd613e1e565b6136568987860163ffffffff6152a016565b905061582c61581c613e1e565b613656898763ffffffff6152a016565b602086018190529003845250612ed79050565b61585a61584a613e1e565b613656868463ffffffff6152a016565b60408301525092915050565b61586e615b1c565b611f3e61587b858561522e565b8361522e565b6000606461588d613e1e565b6069028161589757fe5b048211156158bb5760646158a9613e1e565b606902816158b357fe5b049050611c38565b60646158c5613e1e565b605f02816158cf57fe5b048210156158eb5760646158e1613e1e565b605f02816158b357fe5b5090565b600066038d7ea4c6800082111561590e575066038d7ea4c68000611c38565b633b9aca008210156158eb5750633b9aca00611c38565b6001600160a01b038216600090815260736020908152604080832084845290915281206001015460675461595a858583615a9c565b15615968579150610f699050565b615973858584615a9c565b61598257600092505050610f69565b8082111561599557600092505050610f69565b808210156159c8576002818301046159ae868683615a9c565b156159be578060010192506159c2565b8091505b50615995565b806159d857600092505050610f69565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01949350505050565b6000818310615a1357506000611f3e565b600083815260776020818152604080842088855260019081018352818520548786529383528185208986520190915290912054613740615a51613e1e565b61365689615048858763ffffffff612f8616565b60008183615a865760405162461bcd60e51b8152600401610e879190616fa3565b506000838581615a9257fe5b0495945050505050565b6001600160a01b03831660009081526073602090815260408083208584529091528120600101548210801590611f3e57506001600160a01b0384166000908152607360209081526040808320868452909152902060020154615afd83615b07565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b604051806080016040528060006001600160a01b031681526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215615ba9579160200282015b82811115615ba9578235825591602001919060010190615b8e565b506158eb929150615c52565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10615bf657805160ff1916838001178555615ba9565b82800160010185558215615ba9579182015b82811115615ba9578251825591602001919060010190615c08565b6040518060a0016040528060608152602001600081526020016060815260200160008152602001600081525090565b611ba691905b808211156158eb5760008155600101615c58565b8035610f69816173eb565b60008083601f840112615c8957600080fd5b50813567ffffffffffffffff811115615ca157600080fd5b602083019150836020820283011115615cb957600080fd5b9250929050565b8035610f69816173ff565b8051610f69816173ff565b60008083601f840112615ce857600080fd5b50813567ffffffffffffffff811115615d0057600080fd5b602083019150836001820283011115615cb957600080fd5b8035610f6981617408565b8051610f6981617408565b600060208284031215615d4057600080fd5b6000611f3e8484615c6c565b60008060008060608587031215615d6257600080fd5b6000615d6e8787615c6c565b9450506020615d7f87828801615d18565b935050604085013567ffffffffffffffff811115615d9c57600080fd5b615da887828801615cd6565b95989497509550505050565b60008060008060808587031215615dca57600080fd5b6000615dd68787615c6c565b9450506020615de787828801615c6c565b9350506040615df887828801615cc0565b9250506060615e0987828801615d18565b91505092959194509250565b60008060408385031215615e2857600080fd5b6000615e348585615c6c565b9250506020615e4585828601615d18565b9150509250929050565b60008060008060008060008060006101008a8c031215615e6e57600080fd5b6000615e7a8c8c615c6c565b9950506020615e8b8c828d01615d18565b98505060408a013567ffffffffffffffff811115615ea857600080fd5b615eb48c828d01615cd6565b97509750506060615ec78c828d01615d18565b9550506080615ed88c828d01615d18565b94505060a0615ee98c828d01615d18565b93505060c0615efa8c828d01615d18565b92505060e0615f0b8c828d01615d18565b9150509295985092959850929598565b600080600060608486031215615f3057600080fd5b6000615f3c8686615c6c565b9350506020615f4d86828701615d18565b9250506040615f5e86828701615d18565b9150509250925092565b60008060008060808587031215615f7e57600080fd5b6000615f8a8787615c6c565b9450506020615f9b87828801615d18565b9350506040615df887828801615d18565b60008060008060008060008060006101208a8c031215615fcb57600080fd5b6000615fd78c8c615c6c565b9950506020615fe88c828d01615d18565b9850506040615ff98c828d01615d18565b975050606061600a8c828d01615d18565b965050608061601b8c828d01615d18565b95505060a061602c8c828d01615d18565b94505060c061603d8c828d01615d18565b93505060e061604e8c828d01615d18565b925050610100615f0b8c828d01615d18565b6000806020838503121561607357600080fd5b823567ffffffffffffffff81111561608a57600080fd5b61609685828601615c77565b92509250509250929050565b600080600080600080600080600060a08a8c0312156160c057600080fd5b893567ffffffffffffffff8111156160d757600080fd5b6160e38c828d01615c77565b995099505060208a013567ffffffffffffffff81111561610257600080fd5b61610e8c828d01615c77565b975097505060408a013567ffffffffffffffff81111561612d57600080fd5b6161398c828d01615c77565b955095505060608a013567ffffffffffffffff81111561615857600080fd5b6161648c828d01615c77565b93509350506080615f0b8c828d01615d18565b60006020828403121561618957600080fd5b6000611f3e8484615ccb565b600080602083850312156161a857600080fd5b823567ffffffffffffffff8111156161bf57600080fd5b61609685828601615cd6565b6000602082840312156161dd57600080fd5b6000611f3e8484615d18565b6000602082840312156161fb57600080fd5b6000611f3e8484615d23565b6000806040838503121561621a57600080fd5b60006162268585615d18565b9250506020615e4585828601615cc0565b6000806040838503121561624a57600080fd5b6000615e348585615d18565b600080600080600060a0868803121561626e57600080fd5b600061627a8888615d18565b955050602061628b88828901615d18565b945050604061629c88828901615c6c565b93505060606162ad88828901615c6c565b92505060806162be88828901615c6c565b9150509295509295909350565b6000806000606084860312156162e057600080fd5b6000615f3c8686615d18565b60006162f88383616e3e565b505060800190565b600061630c8383616e88565b505060600190565b60006163208383616ebb565b505060200190565b6163318161734a565b82525050565b60006163428261733d565b61634c8185617341565b93506163578361732b565b8060005b8381101561638557815161636f88826162ec565b975061637a8361732b565b92505060010161635b565b509495945050505050565b600061639b8261733d565b6163a58185617341565b93506163b08361732b565b8060005b838110156163855781516163c88882616300565b97506163d38361732b565b9250506001016163b4565b60006163e98261733d565b6163f38185617341565b93506163fe8361732b565b8060005b838110156163855781516164168882616314565b97506164218361732b565b925050600101616402565b61633181617355565b6163318161735a565b60006164498261733d565b6164538185611c38565b9350616463818560208601617397565b9290920192915050565b60006164788261733d565b6164828185617341565b9350616492818560208601617397565b61649b816173c3565b9093019392505050565b6000815460018116600081146164c2576001811461650657616545565b607f60028304166164d38187617341565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0084168152955050602085019250616545565b600282046165148187617341565b955061651f85617331565b60005b8281101561653e57815488820152600190910190602001616522565b8701945050505b505092915050565b60006165598385617341565b935061656683858461738b565b61649b836173c3565b600061657c601783617341565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b60006165b5600b83617341565b7f7a65726f20616d6f756e74000000000000000000000000000000000000000000815260200192915050565b60006165ee602683617341565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181527f6464726573730000000000000000000000000000000000000000000000000000602082015260400192915050565b600061664d601183617341565b7f616c7265616479206c6f636b6564207570000000000000000000000000000000815260200192915050565b6000616686601b83617341565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b60006166bf602983617341565b7f63616c6c6572206973206e6f7420746865204e6f64654472697665724175746881527f20636f6e74726163740000000000000000000000000000000000000000000000602082015260400192915050565b600061671e601083617341565b7f6e6f7468696e6720746f20737461736800000000000000000000000000000000815260200192915050565b6000616757601283617341565b7f696e636f7272656374206475726174696f6e0000000000000000000000000000815260200192915050565b6000616790601683617341565b7f76616c696461746f722069736e27742061637469766500000000000000000000815260200192915050565b60006167c9600c83617341565b7f77726f6e67207374617475730000000000000000000000000000000000000000815260200192915050565b6000616802601683617341565b7f6e6f7420656e6f7567682074696d652070617373656400000000000000000000815260200192915050565b600061683b601883617341565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b6000616874600d83617341565b7f6e6f74206c6f636b656420757000000000000000000000000000000000000000815260200192915050565b60006168ad602183617341565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81527f7700000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b600061690c600c83617341565b7f7a65726f20726577617264730000000000000000000000000000000000000000815260200192915050565b6000616945601f83617341565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b600061697e602083617341565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b60006169b7602e83617341565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581527f656e20696e697469616c697a6564000000000000000000000000000000000000602082015260400192915050565b6000616a16602183617341565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e81527f3000000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b6000616a75601983617341565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b6000616aae600c83617341565b7f656d707479207075626b65790000000000000000000000000000000000000000815260200192915050565b6000616ae7601283617341565b7f4661696c656420746f2073656e642046544d0000000000000000000000000000815260200192915050565b6000616b20601583617341565b7f7265717565737420646f65736e27742065786973740000000000000000000000815260200192915050565b6000616b59602883617341565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e6481527f206561726c696572000000000000000000000000000000000000000000000000602082015260400192915050565b6000616bb8601783617341565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b6000616bf1601883617341565b7f6f75747374616e64696e67207346544d2062616c616e63650000000000000000815260200192915050565b6000610f69600083611c38565b6000616c37601083617341565b7f6e6f7420656e6f756768207374616b6500000000000000000000000000000000815260200192915050565b6000616c70602983617341565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381527f2065786365656465640000000000000000000000000000000000000000000000602082015260400192915050565b6000616ccf601b83617341565b7f676f7620766f746573207265636f756e74696e67206661696c65640000000000815260200192915050565b6000616d08601783617341565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b6000616d41601883617341565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b6000616d7a601783617341565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000616db3601683617341565b7f7374616b652069732066756c6c7920736c617368656400000000000000000000815260200192915050565b6000616dec602c83617341565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581527f2077686f6c65207374616b650000000000000000000000000000000000000000602082015260400192915050565b80516080830190616e4f8482616328565b506020820151616e626020850182616ebb565b506040820151616e756040850182616ebb565b5060608201516128ca6060850182616ebb565b80516060830190616e998482616ebb565b506020820151616eac6020850182616ebb565b5060408201516128ca60408501825b61633181611ba6565b6000613e17828461643e565b6000610f6982616c1d565b60208101610f698284616328565b60408101616ef78285616328565b613e176020830184616328565b60408101616f128285616328565b613e176020830184616ebb565b60808101616f2d8287616328565b616f3a6020830186616ebb565b616f476040830185616ebb565b6156676060830184616ebb565b60208082528101610f668184616337565b60208082528101610f668184616390565b60208082528101610f6681846163de565b60208101610f69828461642c565b60208101610f698284616435565b60208082528101610f66818461646d565b60208082528101610f698161656f565b60208082528101610f69816165a8565b60208082528101610f69816165e1565b60208082528101610f6981616640565b60208082528101610f6981616679565b60208082528101610f69816166b2565b60208082528101610f6981616711565b60208082528101610f698161674a565b60208082528101610f6981616783565b60208082528101610f69816167bc565b60208082528101610f69816167f5565b60208082528101610f698161682e565b60208082528101610f6981616867565b60208082528101610f69816168a0565b60208082528101610f69816168ff565b60208082528101610f6981616938565b60208082528101610f6981616971565b60208082528101610f69816169aa565b60208082528101610f6981616a09565b60208082528101610f6981616a68565b60208082528101610f6981616aa1565b60208082528101610f6981616ada565b60208082528101610f6981616b13565b60208082528101610f6981616b4c565b60208082528101610f6981616bab565b60208082528101610f6981616be4565b60208082528101610f6981616c2a565b60208082528101610f6981616c63565b60208082528101610f6981616cc2565b60208082528101610f6981616cfb565b60208082528101610f6981616d34565b60208082528101610f6981616d6d565b60208082528101610f6981616da6565b60208082528101610f6981616ddf565b60208101610f698284616ebb565b604081016171f08285616ebb565b8181036020830152611f3e81846164a5565b604081016172108286616ebb565b818103602083015261566781848661654d565b60408101616f128285616ebb565b6060810161723f8286616ebb565b61724c6020830185616ebb565b611f3e6040830184616ebb565b60808101616f2d8287616ebb565b60e08101617275828a616ebb565b6172826020830189616ebb565b61728f6040830188616ebb565b61729c6060830187616ebb565b6172a96080830186616ebb565b6172b660a0830185616ebb565b6172c360c0830184616328565b98975050505050505050565b60e081016172dd828a616ebb565b6172ea6020830189616ebb565b6172f76040830188616ebb565b6173046060830187616ebb565b6173116080830186616ebb565b61731e60a0830185616ebb565b6172c360c0830184616ebb565b60200190565b60009081526020902090565b5190565b90815260200190565b6000610f698261737f565b151590565b7fffffff00000000000000000000000000000000000000000000000000000000001690565b6001600160a01b031690565b82818337506000910152565b60005b838110156173b257818101518382015260200161739a565b838111156128ca5750506000910152565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b6173f48161734a565b81146125fb57600080fd5b6173f481617355565b6173f481611ba656fea365627a7a72315820dd8c91db5ff0493f52039960637d61bbf0e7908ecf45196d8babeb77a40014fd6c6578706572696d656e74616cf564736f6c63430005110040"

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
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcDelegationRewards(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calcDelegationRewards", delegator, toStakerID, arg2, arg3)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, arg2, arg3)
}

// CalcDelegationRewards is a free data retrieval call binding the contract method 0xd845fc90.
//
// Solidity: function calcDelegationRewards(address delegator, uint256 toStakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcDelegationRewards(delegator common.Address, toStakerID *big.Int, arg2 *big.Int, arg3 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcDelegationRewards(&_Contract.CallOpts, delegator, toStakerID, arg2, arg3)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) CalcValidatorRewards(opts *bind.CallOpts, stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "calcValidatorRewards", stakerID, arg1, arg2)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractSession) CalcValidatorRewards(stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, arg1, arg2)
}

// CalcValidatorRewards is a free data retrieval call binding the contract method 0x96060e71.
//
// Solidity: function calcValidatorRewards(uint256 stakerID, uint256 , uint256 ) view returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) CalcValidatorRewards(stakerID *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.CalcValidatorRewards(&_Contract.CallOpts, stakerID, arg1, arg2)
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

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodEpochs is a free data retrieval call binding the contract method 0x1d58179c.
//
// Solidity: function delegationLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodEpochs(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) DelegationLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// DelegationLockPeriodTime is a free data retrieval call binding the contract method 0xec6a7f1c.
//
// Solidity: function delegationLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.DelegationLockPeriodTime(&_Contract.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCaller) Delegations(opts *bind.CallOpts, _from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegations", _from, _toStakerID)

	outstruct := new(struct {
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		Amount           *big.Int
		PaidUntilEpoch   *big.Int
		ToStakerID       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CreatedEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ToStakerID = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractSession) Delegations(_from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, _from, _toStakerID)
}

// Delegations is a free data retrieval call binding the contract method 0x223fae09.
//
// Solidity: function delegations(address _from, uint256 _toStakerID) view returns(uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 amount, uint256 paidUntilEpoch, uint256 toStakerID)
func (_Contract *ContractCallerSession) Delegations(_from common.Address, _toStakerID *big.Int) (struct {
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Amount           *big.Int
	PaidUntilEpoch   *big.Int
	ToStakerID       *big.Int
}, error) {
	return _Contract.Contract.Delegations(&_Contract.CallOpts, _from, _toStakerID)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractCaller) DelegationsNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationsNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsNum is a free data retrieval call binding the contract method 0x4bd202dc.
//
// Solidity: function delegationsNum() pure returns(uint256)
func (_Contract *ContractCallerSession) DelegationsNum() (*big.Int, error) {
	return _Contract.Contract.DelegationsNum(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) DelegationsTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "delegationsTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
}

// DelegationsTotalAmount is a free data retrieval call binding the contract method 0x30fa9929.
//
// Solidity: function delegationsTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) DelegationsTotalAmount() (*big.Int, error) {
	return _Contract.Contract.DelegationsTotalAmount(&_Contract.CallOpts)
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

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractCaller) GetStakerID(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStakerID", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractSession) GetStakerID(_addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, _addr)
}

// GetStakerID is a free data retrieval call binding the contract method 0x63321e27.
//
// Solidity: function getStakerID(address _addr) view returns(uint256)
func (_Contract *ContractCallerSession) GetStakerID(_addr common.Address) (*big.Int, error) {
	return _Contract.Contract.GetStakerID(&_Contract.CallOpts, _addr)
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

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCaller) IsDelegationLockedUp(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isDelegationLockedUp", delegator, toStakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
}

// IsDelegationLockedUp is a free data retrieval call binding the contract method 0xcfd5fa0c.
//
// Solidity: function isDelegationLockedUp(address delegator, uint256 toStakerID) view returns(bool)
func (_Contract *ContractCallerSession) IsDelegationLockedUp(delegator common.Address, toStakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsDelegationLockedUp(&_Contract.CallOpts, delegator, toStakerID)
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

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractCaller) IsStakeLockedUp(opts *bind.CallOpts, stakerID *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isStakeLockedUp", stakerID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractSession) IsStakeLockedUp(stakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, stakerID)
}

// IsStakeLockedUp is a free data retrieval call binding the contract method 0x7f664d87.
//
// Solidity: function isStakeLockedUp(uint256 stakerID) view returns(bool)
func (_Contract *ContractCallerSession) IsStakeLockedUp(stakerID *big.Int) (bool, error) {
	return _Contract.Contract.IsStakeLockedUp(&_Contract.CallOpts, stakerID)
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

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedDelegations(opts *bind.CallOpts, delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lockedDelegations", delegator, toStakerID)

	outstruct := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedDelegations(delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, delegator, toStakerID)
}

// LockedDelegations is a free data retrieval call binding the contract method 0xdd099bb6.
//
// Solidity: function lockedDelegations(address delegator, uint256 toStakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedDelegations(delegator common.Address, toStakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedDelegations(&_Contract.CallOpts, delegator, toStakerID)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCaller) LockedStakes(opts *bind.CallOpts, stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lockedStakes", stakerID)

	outstruct := new(struct {
		FromEpoch *big.Int
		EndTime   *big.Int
		Duration  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromEpoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractSession) LockedStakes(stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, stakerID)
}

// LockedStakes is a free data retrieval call binding the contract method 0xdf4f49d4.
//
// Solidity: function lockedStakes(uint256 stakerID) view returns(uint256 fromEpoch, uint256 endTime, uint256 duration)
func (_Contract *ContractCallerSession) LockedStakes(stakerID *big.Int) (struct {
	FromEpoch *big.Int
	EndTime   *big.Int
	Duration  *big.Int
}, error) {
	return _Contract.Contract.LockedStakes(&_Contract.CallOpts, stakerID)
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

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegation() (*big.Int, error) {
	return _Contract.Contract.MinDelegation(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegationDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationDecrease is a free data retrieval call binding the contract method 0xcb1c4e67.
//
// Solidity: function minDelegationDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationDecrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationDecrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinDelegationIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minDelegationIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
}

// MinDelegationIncrease is a free data retrieval call binding the contract method 0x60c7e37f.
//
// Solidity: function minDelegationIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinDelegationIncrease() (*big.Int, error) {
	return _Contract.Contract.MinDelegationIncrease(&_Contract.CallOpts)
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

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStake() (*big.Int, error) {
	return _Contract.Contract.MinStake(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeDecrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStakeDecrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeDecrease is a free data retrieval call binding the contract method 0x19ddb54f.
//
// Solidity: function minStakeDecrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeDecrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeDecrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCaller) MinStakeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minStakeIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
}

// MinStakeIncrease is a free data retrieval call binding the contract method 0xc4b5dd7e.
//
// Solidity: function minStakeIncrease() pure returns(uint256)
func (_Contract *ContractCallerSession) MinStakeIncrease() (*big.Int, error) {
	return _Contract.Contract.MinStakeIncrease(&_Contract.CallOpts)
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

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeLockPeriodEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodEpochs is a free data retrieval call binding the contract method 0x54d77ed2.
//
// Solidity: function stakeLockPeriodEpochs() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodEpochs() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodEpochs(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCaller) StakeLockPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeLockPeriodTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeLockPeriodTime is a free data retrieval call binding the contract method 0x3fee10a8.
//
// Solidity: function stakeLockPeriodTime() pure returns(uint256)
func (_Contract *ContractCallerSession) StakeLockPeriodTime() (*big.Int, error) {
	return _Contract.Contract.StakeLockPeriodTime(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCaller) StakeTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakeTotalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// StakeTotalAmount is a free data retrieval call binding the contract method 0x3d0317fe.
//
// Solidity: function stakeTotalAmount() view returns(uint256)
func (_Contract *ContractCallerSession) StakeTotalAmount() (*big.Int, error) {
	return _Contract.Contract.StakeTotalAmount(&_Contract.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCaller) Stakers(opts *bind.CallOpts, _stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakers", _stakerID)

	outstruct := new(struct {
		Status           *big.Int
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedEpoch *big.Int
		DeactivatedTime  *big.Int
		StakeAmount      *big.Int
		PaidUntilEpoch   *big.Int
		DelegatedMe      *big.Int
		DagAddress       common.Address
		SfcAddress       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreatedEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StakeAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PaidUntilEpoch = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DelegatedMe = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.DagAddress = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.SfcAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractSession) Stakers(_stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, _stakerID)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 _stakerID) view returns(uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime, uint256 stakeAmount, uint256 paidUntilEpoch, uint256 delegatedMe, address dagAddress, address sfcAddress)
func (_Contract *ContractCallerSession) Stakers(_stakerID *big.Int) (struct {
	Status           *big.Int
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	StakeAmount      *big.Int
	PaidUntilEpoch   *big.Int
	DelegatedMe      *big.Int
	DagAddress       common.Address
	SfcAddress       common.Address
}, error) {
	return _Contract.Contract.Stakers(&_Contract.CallOpts, _stakerID)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCaller) StakersLastID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakersLastID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersLastID is a free data retrieval call binding the contract method 0x81d9dc7a.
//
// Solidity: function stakersLastID() view returns(uint256)
func (_Contract *ContractCallerSession) StakersLastID() (*big.Int, error) {
	return _Contract.Contract.StakersLastID(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCaller) StakersNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "stakersNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
}

// StakersNum is a free data retrieval call binding the contract method 0x08728f6e.
//
// Solidity: function stakersNum() view returns(uint256)
func (_Contract *ContractCallerSession) StakersNum() (*big.Int, error) {
	return _Contract.Contract.StakersNum(&_Contract.CallOpts)
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

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationCompoundRewards(opts *bind.TransactOpts, arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationCompoundRewards", arg0, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationCompoundRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationCompoundRewards is a paid mutator transaction binding the contract method 0xdc599bb1.
//
// Solidity: function claimDelegationCompoundRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationCompoundRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationCompoundRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactor) ClaimDelegationRewards(opts *bind.TransactOpts, arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDelegationRewards", arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractSession) ClaimDelegationRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, arg0, toStakerID)
}

// ClaimDelegationRewards is a paid mutator transaction binding the contract method 0xf99837e6.
//
// Solidity: function claimDelegationRewards(uint256 , uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) ClaimDelegationRewards(arg0 *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDelegationRewards(&_Contract.TransactOpts, arg0, toStakerID)
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

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractTransactor) ClaimValidatorCompoundRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorCompoundRewards", arg0)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractSession) ClaimValidatorCompoundRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorCompoundRewards is a paid mutator transaction binding the contract method 0xcda5826a.
//
// Solidity: function claimValidatorCompoundRewards(uint256 ) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorCompoundRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorCompoundRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractTransactor) ClaimValidatorRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimValidatorRewards", arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractSession) ClaimValidatorRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, arg0)
}

// ClaimValidatorRewards is a paid mutator transaction binding the contract method 0x295cccba.
//
// Solidity: function claimValidatorRewards(uint256 ) returns()
func (_Contract *ContractTransactorSession) ClaimValidatorRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ClaimValidatorRewards(&_Contract.TransactOpts, arg0)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactor) CreateDelegation(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createDelegation", toValidatorID)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractSession) CreateDelegation(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, toValidatorID)
}

// CreateDelegation is a paid mutator transaction binding the contract method 0xc312eb07.
//
// Solidity: function createDelegation(uint256 toValidatorID) payable returns()
func (_Contract *ContractTransactorSession) CreateDelegation(toValidatorID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateDelegation(&_Contract.TransactOpts, toValidatorID)
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

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactor) LockUpDelegation(opts *bind.TransactOpts, lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpDelegation", lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpDelegation is a paid mutator transaction binding the contract method 0xa4b89fab.
//
// Solidity: function lockUpDelegation(uint256 lockDuration, uint256 toStakerID) returns()
func (_Contract *ContractTransactorSession) LockUpDelegation(lockDuration *big.Int, toStakerID *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpDelegation(&_Contract.TransactOpts, lockDuration, toStakerID)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactor) LockUpStake(opts *bind.TransactOpts, lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockUpStake", lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// LockUpStake is a paid mutator transaction binding the contract method 0xf3ae5b1a.
//
// Solidity: function lockUpStake(uint256 lockDuration) returns()
func (_Contract *ContractTransactorSession) LockUpStake(lockDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LockUpStake(&_Contract.TransactOpts, lockDuration)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractTransactor) PartialWithdrawByRequest(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "partialWithdrawByRequest", arg0)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractSession) PartialWithdrawByRequest(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, arg0)
}

// PartialWithdrawByRequest is a paid mutator transaction binding the contract method 0xf8b18d8a.
//
// Solidity: function partialWithdrawByRequest(uint256 ) returns()
func (_Contract *ContractTransactorSession) PartialWithdrawByRequest(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PartialWithdrawByRequest(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegation(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegation", arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegation is a paid mutator transaction binding the contract method 0xb1e64339.
//
// Solidity: function prepareToWithdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawDelegationPartial(opts *bind.TransactOpts, wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawDelegationPartial", wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawDelegationPartial is a paid mutator transaction binding the contract method 0xbb03a4bd.
//
// Solidity: function prepareToWithdrawDelegationPartial(uint256 wrID, uint256 toStakerID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawDelegationPartial(wrID *big.Int, toStakerID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawDelegationPartial(&_Contract.TransactOpts, wrID, toStakerID, amount)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStake")
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStake is a paid mutator transaction binding the contract method 0xc41b6405.
//
// Solidity: function prepareToWithdrawStake() returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStake(&_Contract.TransactOpts)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) PrepareToWithdrawStakePartial(opts *bind.TransactOpts, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "prepareToWithdrawStakePartial", wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
}

// PrepareToWithdrawStakePartial is a paid mutator transaction binding the contract method 0x26682c71.
//
// Solidity: function prepareToWithdrawStakePartial(uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) PrepareToWithdrawStakePartial(wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PrepareToWithdrawStakePartial(&_Contract.TransactOpts, wrID, amount)
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

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, wrID, amount)
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

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactor) WithdrawDelegation(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawDelegation", arg0)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractSession) WithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 ) returns()
func (_Contract *ContractTransactorSession) WithdrawDelegation(arg0 *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawDelegation(&_Contract.TransactOpts, arg0)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
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

var ContractBinRuntime = "0x6080604052600436106105fb5760003560e01c80638b0e9f3f1161030e578063c65ee0e11161019b578063de67f215116100e7578063ebdf104c116100a0578063f3ae5b1a1161007a578063f3ae5b1a14611899578063f8b18d8a146115e7578063f99837e6146118c3578063fd5e6dd1146118f3576105fb565b8063ebdf104c146116e6578063ec6a7f1c14611851578063f2fde38b14611866576105fb565b8063de67f21514611581578063df00c922146115b7578063df0e307a146115e7578063df4f49d414611611578063e08d7e661461163b578063e261641a146116b6576105fb565b8063cfd5fa0c11610154578063d9a7c1f91161012e578063d9a7c1f9146114d3578063dc31e1af146114e8578063dc599bb114611518578063dd099bb614611548576105fb565b8063cfd5fa0c1461141c578063cfdbb7cd14611455578063d845fc901461148e576105fb565b8063c65ee0e114611348578063c7be95de14611372578063cb1c4e671461068e578063cc8343aa14611387578063cda5826a146113b9578063cfd47663146113e3576105fb565b8063b1e643391161025a578063bb03a4bd11610213578063c3de580e116101ed578063c3de580e146112f4578063c41b64051461131e578063c4b5dd7e1461068e578063c5f530af14611333576105fb565b8063bb03a4bd146112a9578063bed9d861146112df578063c312eb0714610fe9576105fb565b8063b1e6433914611122578063b5d896271461114c578063b6d9edd5146111b7578063b810e411146111e1578063b82b84271461121a578063b88a37e21461122f576105fb565b806396c7ee46116102c7578063a4b89fab116102a1578063a4b89fab14611036578063a5a470ad14611066578063a7786515146110d4578063a86a056f146110e9576105fb565b806396c7ee4614610f8a5780639fa6dd3514610fe9578063a198d22914611006576105fb565b80638b0e9f3f14610e905780638b1a0d1114610ea55780638cddb01514610ed55780638da5cb5b14610f0e5780638f32d59b14610f3f57806396060e7114610f54576105fb565b80633d0317fe1161048c5780636099ecb2116103d85780636f498663116103915780637cacb1d61161036b5780637cacb1d614610d9e5780637f664d8714610db357806381d9dc7a146106a3578063854873e114610df1576105fb565b80636f49866314610d3b578063715018a614610d745780637667180814610d89576105fb565b80636099ecb214610c5157806360c7e37f1461068e57806361e53fcc14610c8a57806363321e2714610cba578063650acd6614610ced578063670322f814610d02576105fb565b80634feb92f3116104455780635601fe011161041f5780635601fe0114610be257806358f95b8014610c0c5780635e2308d2146109715780635fab23a814610c3c576105fb565b80634feb92f314610b0757806354d77ed21461081957806354fd4d5014610bb0576105fb565b80633d0317fe14610a475780633fee10a814610819578063441a3e7014610a5c5780634bd202dc14610a8c5780634f7c4efb14610aa15780634f864df414610ad1576105fb565b80631d58179c1161054b5780632709275e116105045780632cedb097116104de5780632cedb097146109c557806330fa9929146109f3578063375b3c0a14610a0857806339b80c0014610a1d576105fb565b80632709275e1461097157806328f7314814610986578063295cccba1461099b576105fb565b80631d58179c146108195780631e702f831461082e5780631f2701521461085e578063223fae09146108bb5780632265f2841461092c57806326682c7114610941576105fb565b80630d4955e3116105b857806318160ddd1161059257806318160ddd1461076f57806318f628d41461078457806319ddb54f1461068e5780631d3ac42c146107e9576105fb565b80630d4955e31461070c5780630d7b26091461072157806312622d0e14610736576105fb565b80630135b1db14610600578063019e272914610645578063029859921461068e57806308728f6e146106a357806308c36874146106b85780630962ef79146106e2575b600080fd5b34801561060c57600080fd5b506106336004803603602081101561062357600080fd5b50356001600160a01b0316611979565b60408051918252519081900360200190f35b34801561065157600080fd5b5061068c6004803603608081101561066857600080fd5b508035906020810135906001600160a01b036040820135811691606001351661198b565b005b34801561069a57600080fd5b50610633611a92565b3480156106af57600080fd5b50610633611a98565b3480156106c457600080fd5b5061068c600480360360208110156106db57600080fd5b5035611a9e565b3480156106ee57600080fd5b5061068c6004803603602081101561070557600080fd5b5035611b6a565b34801561071857600080fd5b50610633611c47565b34801561072d57600080fd5b50610633611c4f565b34801561074257600080fd5b506106336004803603604081101561075957600080fd5b506001600160a01b038135169060200135611c56565b34801561077b57600080fd5b50610633611cdf565b34801561079057600080fd5b5061068c60048036036101208110156107a857600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e0810135906101000135611ce5565b3480156107f557600080fd5b506106336004803603604081101561080c57600080fd5b5080359060200135611e45565b34801561082557600080fd5b50610633611fd8565b34801561083a57600080fd5b5061068c6004803603604081101561085157600080fd5b5080359060200135611fe7565b34801561086a57600080fd5b5061089d6004803603606081101561088157600080fd5b506001600160a01b038135169060208101359060400135612085565b60408051938452602084019290925282820152519081900360600190f35b3480156108c757600080fd5b506108f4600480360360408110156108de57600080fd5b506001600160a01b0381351690602001356120b7565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b34801561093857600080fd5b5061063361212b565b34801561094d57600080fd5b5061068c6004803603604081101561096457600080fd5b508035906020013561213d565b34801561097d57600080fd5b5061063361215d565b34801561099257600080fd5b50610633612179565b3480156109a757600080fd5b5061068c600480360360208110156109be57600080fd5b503561217f565b3480156109d157600080fd5b506109da612198565b6040805192835260208301919091528051918290030190f35b3480156109ff57600080fd5b506106336121a2565b348015610a1457600080fd5b506106336121b5565b348015610a2957600080fd5b506108f460048036036020811015610a4057600080fd5b50356121bf565b348015610a5357600080fd5b50610633612201565b348015610a6857600080fd5b5061068c60048036036040811015610a7f57600080fd5b5080359060200135612212565b348015610a9857600080fd5b50610633612556565b348015610aad57600080fd5b5061068c60048036036040811015610ac457600080fd5b508035906020013561255b565b348015610add57600080fd5b5061068c60048036036060811015610af457600080fd5b508035906020810135906040013561268d565b348015610b1357600080fd5b5061068c6004803603610100811015610b2b57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b811115610b5a57600080fd5b820183602082011115610b6c57600080fd5b803590602001918460018302840111600160201b83111715610b8d57600080fd5b9193509150803590602081013590604081013590606081013590608001356129f9565b348015610bbc57600080fd5b50610bc5612a9f565b604080516001600160e81b03199092168252519081900360200190f35b348015610bee57600080fd5b5061063360048036036020811015610c0557600080fd5b5035612aa9565b348015610c1857600080fd5b5061063360048036036040811015610c2f57600080fd5b5080359060200135612adf565b348015610c4857600080fd5b50610633612afc565b348015610c5d57600080fd5b5061063360048036036040811015610c7457600080fd5b506001600160a01b038135169060200135612b02565b348015610c9657600080fd5b5061063360048036036040811015610cad57600080fd5b5080359060200135612b40565b348015610cc657600080fd5b5061063360048036036020811015610cdd57600080fd5b50356001600160a01b0316612b61565b348015610cf957600080fd5b50610633612b7c565b348015610d0e57600080fd5b5061063360048036036040811015610d2557600080fd5b506001600160a01b038135169060200135612b81565b348015610d4757600080fd5b5061063360048036036040811015610d5e57600080fd5b506001600160a01b038135169060200135612bc2565b348015610d8057600080fd5b5061068c612c2c565b348015610d9557600080fd5b50610633612cbd565b348015610daa57600080fd5b50610633612cc6565b348015610dbf57600080fd5b50610ddd60048036036020811015610dd657600080fd5b5035612ccc565b604080519115158252519081900360200190f35b348015610dfd57600080fd5b50610e1b60048036036020811015610e1457600080fd5b5035612cf1565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610e55578181015183820152602001610e3d565b50505050905090810190601f168015610e825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610e9c57600080fd5b50610633612d8c565b348015610eb157600080fd5b5061068c60048036036040811015610ec857600080fd5b5080359060200135612d92565b348015610ee157600080fd5b5061068c60048036036040811015610ef857600080fd5b506001600160a01b038135169060200135612e22565b348015610f1a57600080fd5b50610f23612e70565b604080516001600160a01b039092168252519081900360200190f35b348015610f4b57600080fd5b50610ddd612e7f565b348015610f6057600080fd5b5061089d60048036036060811015610f7757600080fd5b5080359060208101359060400135612e90565b348015610f9657600080fd5b50610fc360048036036040811015610fad57600080fd5b506001600160a01b038135169060200135612ee8565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61068c60048036036020811015610fff57600080fd5b5035612f1a565b34801561101257600080fd5b506106336004803603604081101561102957600080fd5b5080359060200135612f28565b34801561104257600080fd5b5061068c6004803603604081101561105957600080fd5b5080359060200135612f49565b61068c6004803603602081101561107c57600080fd5b810190602081018135600160201b81111561109657600080fd5b8201836020820111156110a857600080fd5b803590602001918460018302840111600160201b831117156110c957600080fd5b509092509050612f71565b3480156110e057600080fd5b50610633613055565b3480156110f557600080fd5b506106336004803603604081101561110c57600080fd5b506001600160a01b03813516906020013561306b565b34801561112e57600080fd5b5061068c6004803603602081101561114557600080fd5b5035613088565b34801561115857600080fd5b506111766004803603602081101561116f57600080fd5b50356130d5565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b3480156111c357600080fd5b5061068c600480360360208110156111da57600080fd5b503561311b565b3480156111ed57600080fd5b5061089d6004803603604081101561120457600080fd5b506001600160a01b0381351690602001356131fb565b34801561122657600080fd5b50610633613227565b34801561123b57600080fd5b506112596004803603602081101561125257600080fd5b503561322e565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561129557818101518382015260200161127d565b505050509050019250505060405180910390f35b3480156112b557600080fd5b5061068c600480360360608110156112cc57600080fd5b5080359060208101359060400135613293565b3480156112eb57600080fd5b5061068c61329e565b34801561130057600080fd5b50610ddd6004803603602081101561131757600080fd5b50356132eb565b34801561132a57600080fd5b5061068c613088565b34801561133f57600080fd5b50610633613302565b34801561135457600080fd5b506106336004803603602081101561136b57600080fd5b5035613311565b34801561137e57600080fd5b50610633613323565b34801561139357600080fd5b5061068c600480360360408110156113aa57600080fd5b50803590602001351515613329565b3480156113c557600080fd5b5061068c600480360360208110156113dc57600080fd5b503561350b565b3480156113ef57600080fd5b506106336004803603604081101561140657600080fd5b506001600160a01b038135169060200135613524565b34801561142857600080fd5b50610ddd6004803603604081101561143f57600080fd5b506001600160a01b038135169060200135613541565b34801561146157600080fd5b50610ddd6004803603604081101561147857600080fd5b506001600160a01b038135169060200135613549565b34801561149a57600080fd5b5061089d600480360360808110156114b157600080fd5b506001600160a01b0381351690602081013590604081013590606001356135b1565b3480156114df57600080fd5b506106336135ef565b3480156114f457600080fd5b506106336004803603604081101561150b57600080fd5b50803590602001356135f5565b34801561152457600080fd5b5061068c6004803603604081101561153b57600080fd5b5080359060200135613616565b34801561155457600080fd5b5061089d6004803603604081101561156b57600080fd5b506001600160a01b03813516906020013561361f565b34801561158d57600080fd5b5061068c600480360360608110156115a457600080fd5b508035906020810135906040013561368b565b3480156115c357600080fd5b50610633600480360360408110156115da57600080fd5b508035906020013561398c565b3480156115f357600080fd5b5061068c6004803603602081101561160a57600080fd5b503561329e565b34801561161d57600080fd5b5061089d6004803603602081101561163457600080fd5b50356139ad565b34801561164757600080fd5b5061068c6004803603602081101561165e57600080fd5b810190602081018135600160201b81111561167857600080fd5b82018360208201111561168a57600080fd5b803590602001918460208302840111600160201b831117156116ab57600080fd5b5090925090506139e3565b3480156116c257600080fd5b50610633600480360360408110156116d957600080fd5b5080359060200135613ac3565b3480156116f257600080fd5b5061068c6004803603608081101561170957600080fd5b810190602081018135600160201b81111561172357600080fd5b82018360208201111561173557600080fd5b803590602001918460208302840111600160201b8311171561175657600080fd5b919390929091602081019035600160201b81111561177357600080fd5b82018360208201111561178557600080fd5b803590602001918460208302840111600160201b831117156117a657600080fd5b919390929091602081019035600160201b8111156117c357600080fd5b8201836020820111156117d557600080fd5b803590602001918460208302840111600160201b831117156117f657600080fd5b919390929091602081019035600160201b81111561181357600080fd5b82018360208201111561182557600080fd5b803590602001918460208302840111600160201b8311171561184657600080fd5b509092509050613ae4565b34801561185d57600080fd5b50610633613cc0565b34801561187257600080fd5b5061068c6004803603602081101561188957600080fd5b50356001600160a01b0316613cca565b3480156118a557600080fd5b5061068c600480360360208110156118bc57600080fd5b5035613d1a565b3480156118cf57600080fd5b5061068c600480360360408110156118e657600080fd5b5080359060200135613d3d565b3480156118ff57600080fd5b5061191d6004803603602081101561191657600080fd5b5035613d46565b604080519a8b5260208b0199909952898901979097526060890195909552608088019390935260a087019190915260c086015260e08501526001600160a01b039081166101008501521661012083015251908190036101400190f35b60696020526000908152604090205481565b600054610100900460ff16806119a457506119a4613e69565b806119b2575060005460ff16155b6119ed5760405162461bcd60e51b815260040180806020018281038252602e815260200180615baa602e913960400191505060405180910390fd5b600054610100900460ff16158015611a18576000805460ff1961ff0019909116610100171660011790555b611a2182613e6f565b6067859055606680546001600160a01b0319166001600160a01b03851617905560768490556755cfe697852e904c6075556103e86078556203f480607955611a67613f60565b6000868152607760205260409020600701558015611a8b576000805461ff00191690555b5050505050565b60015b90565b606b5490565b33611aa7615981565b611ab18284613f64565b60208101518151919250600091611acd9163ffffffff61405816565b9050611af08385611aeb85604001518561405890919063ffffffff16565b6140b2565b6001600160a01b0383166000818152607360209081526040808320888452825291829020805485019055845185820151868401518451928352928201528083019190915290518692917f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26919081900360600190a350505050565b33611b73615981565b611b7d8284613f64565b9050816001600160a01b03166108fc611bbb8360400151611baf8560200151866000015161405890919063ffffffff16565b9063ffffffff61405816565b6040518115909202916000818181858888f19350505050158015611be3573d6000803e3d6000fd5b5082826001600160a01b03167fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4583600001518460200151856040015160405180848152602001838152602001828152602001935050505060405180910390a3505050565b6301e1338090565b6212750090565b6000611c628383613549565b611c9057506001600160a01b0382166000908152607260209081526040808320848452909152902054611cd9565b6001600160a01b038316600081815260736020908152604080832086845282528083205493835260728252808320868452909152902054611cd69163ffffffff6141af16565b90505b92915050565b60765481565b611cee336141f1565b611d295760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b611d34898989614205565b6001600160a01b0389166000908152606f602090815260408083208b84529091529020600201819055611d668761436a565b8515611e3a5786861115611dab5760405162461bcd60e51b815260040180806020018281038252602c815260200180615c4a602c913960400191505060405180910390fd5b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855283529281902086905580518781529182018a9052805192938c9390927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7892908290030190a3505b505050505050505050565b336000818152607360209081526040808320868452909152812090919083611ea2576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b611eac8286613549565b611eed576040805162461bcd60e51b815260206004820152600d60248201526c06e6f74206c6f636b656420757609c1b604482015290519081900360640190fd5b8054841115611f43576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000604482015290519081900360640190fd5b611f4d82866143d1565b506000611f608387878560000154614519565b825486900383556001600160a01b03841660008181526072602090815260408083208b8452825291829020805485900390558151898152908101849052815193945089937fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9929181900390910190a395945050505050565b6000611fe2612b7c565b905090565b611ff0336141f1565b61202b5760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b8061206c576040805162461bcd60e51b815260206004820152600c60248201526b77726f6e672073746174757360a01b604482015290519081900360640190fd5b6120768282614662565b612081826000613329565b5050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6001600160a01b03821660009081526072602090815260408083208484529091528120548190819081908190819081908061210857506000965086955085945084935083925082915081905061211f565b600197508796506000955085945092508591508790505b92959891949750929550565b600061213561478c565b601002905090565b3360009081526069602052604090205461215881848461268d565b505050565b6000606461216961478c565b601e028161217357fe5b04905090565b606d5481565b3360009081526069602052604090205461208181611b6a565b6078546079549091565b60006121ac612201565b606c5403905090565b6000611fe2613302565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b60006064606c546018028161217357fe5b3361221b615981565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526122b7576040805162461bcd60e51b81526020600482015260156024820152741c995c5d595cdd08191bd95cdb89dd08195e1a5cdd605a1b604482015290519081900360640190fd5b602080820151825160008781526068909352604090922060010154909190158015906122f3575060008681526068602052604090206001015482115b15612314575050600084815260686020526040902060018101546002909101545b61231c613227565b8201612326613f60565b1015612372576040805162461bcd60e51b81526020600482015260166024820152751b9bdd08195b9bdd59da081d1a5b59481c185cdcd95960521b604482015290519081900360640190fd5b61237a612b7c565b8101612384612cbd565b10156123d7576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152607160209081526040808320898452825280832088845290915281206002015490612410886132eb565b905060006124328383607a60008d815260200190815260200160002054614798565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116124c2576040805162461bcd60e51b81526020600482015260166024820152751cdd185ad9481a5cc8199d5b1b1e481cdb185cda195960521b604482015290519081900360640190fd5b6001600160a01b0387166108fc6124df858463ffffffff6141af16565b6040518115909202916000818181858888f19350505050158015612507573d6000803e3d6000fd5b508789886001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21866040518082815260200191505060405180910390a4505050505050505050565b600090565b612563612e7f565b6125a2576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6125ab826132eb565b6125fc576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f722069736e277420736c6173686564000000000000000000604482015290519081900360640190fd5b61260461478c565b8111156126425760405162461bcd60e51b8152600401808060200182810382526021815260200180615bd86021913960400191505060405180910390fd5b6000828152607a60209081526040918290208390558151838152915184927f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91792908290030190a25050565b3361269881856143d1565b50600082116126dc576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b6126e68185611c56565b82111561273a576040805162461bcd60e51b815260206004820152601960248201527f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526071602090815260408083208784528252808320868452909152902060020154156127b1576040805162461bcd60e51b81526020600482015260136024820152727772494420616c72656164792065786973747360681b604482015290519081900360640190fd5b6001600160a01b038116600090815260726020908152604080832087845282528083208054869003905560689091529020600301546127f6908363ffffffff6141af16565b600085815260686020526040902060030155606c5461281b908363ffffffff6141af16565b606c5560008481526068602052604090205461284857606d54612844908363ffffffff6141af16565b606d555b600061285385612aa9565b905080156128fa57612863613302565b8110156128b1576040805162461bcd60e51b8152602060048201526017602482015276696e73756666696369656e742073656c662d7374616b6560481b604482015290519081900360640190fd5b6128ba856147fa565b6128f55760405162461bcd60e51b8152600401808060200182810382526029815260200180615c216029913960400191505060405180910390fd5b612905565b612905856001614662565b6001600160a01b03821660009081526071602090815260408083208884528252808320878452909152902060020183905561293e612cbd565b6001600160a01b03831660009081526071602090815260408083208984528252808320888452909152902055612972613f60565b6001600160a01b038316600090815260716020908152604080832089845282528083208884529091528120600101919091556129af908690613329565b8385836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57866040518082815260200191505060405180910390a45050505050565b612a02336141f1565b612a3d5760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b612a85898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888614842565b606b54881115611e3a57606b889055505050505050505050565b6203330360ec1b90565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b606e5481565b6000612b0c615981565b612b1684846149f1565b805160208201516040830151929350612b3892611baf9163ffffffff61405816565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b6001600160a01b031660009081526069602052604090205490565b600390565b6000612b8d8383613549565b612b9957506000611cd9565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b6000612bcc615981565b506001600160a01b0383166000908152606f6020908152604080832085845282529182902082516060810184528154808252600183015493820184905260029092015493810184905292612b38929091611baf919063ffffffff61405816565b612c34612e7f565b612c73576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b60675460010190565b60675481565b600081815260686020526040812060060154611cd9906001600160a01b031683613549565b606a6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015612d845780601f10612d5957610100808354040283529160200191612d84565b820191906000526020600020905b815481529060010190602001808311612d6757829003601f168201915b505050505081565b606c5481565b612d9a612e7f565b612dd9576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b60798190556078829055604080518381526020810183905281517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34929181900390910190a15050565b612e2c82826143d1565b612081576040805162461bcd60e51b815260206004820152601060248201526f0dcdee8d0d2dcce40e8de40e6e8c2e6d60831b604482015290519081900360640190fd5b6033546001600160a01b031690565b6033546001600160a01b0316331490565b600083815260686020526040812060060154819081908190612ebb906001600160a01b031688612b02565b905080612ed357506000925060019150829050612edf565b60675490935091508190505b93509350939050565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b612f253382346140b2565b50565b60009182526077602090815260408084209284526005909201905290205490565b336000908152607260209081526040808320848452909152902054612081908290849061368b565b612f79613302565b341015612fc7576040805162461bcd60e51b8152602060048201526017602482015276696e73756666696369656e742073656c662d7374616b6560481b604482015290519081900360640190fd5b80613008576040805162461bcd60e51b815260206004820152600c60248201526b656d707479207075626b657960a01b604482015290519081900360640190fd5b6130483383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250614a5f92505050565b61208133606b54346140b2565b6000606461306161478c565b600f028161217357fe5b607060209081526000928352604080842090915290825290205481565b6040805162461bcd60e51b815260206004820152601f60248201527f75736520534643763320756e64656c656761746528292066756e6374696f6e00604482015290519081900360640190fd5b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b613123612e7f565b613162576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b6801c985c8903591eb208111156131c0576040805162461bcd60e51b815260206004820152601b60248201527f746f6f206c617267652072657761726420706572207365636f6e640000000000604482015290519081900360640190fd5b60758190556040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b62093a8090565b60008181526077602090815260409182902060060180548351818402810184019094528084526060939283018282801561328757602002820191906000526020600020905b815481526020019060010190808311613273575b50505050509050919050565b61215882848361268d565b6040805162461bcd60e51b815260206004820152601d60248201527f75736520534643763320776974686472617728292066756e6374696f6e000000604482015290519081900360640190fd5b600090815260686020526040902054608016151590565b6a02a055184a310c1260000090565b607a6020526000908152604090205481565b606b5481565b61333282614a8a565b61337d576040805162461bcd60e51b81526020600482015260176024820152761d985b1a59185d1bdc88191bd95cdb89dd08195e1a5cdd604a1b604482015290519081900360640190fd5b6000828152606860205260409020600381015490541561339b575060005b6066546040805163520337df60e11b8152600481018690526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b1580156133ef57600080fd5b505af1158015613403573d6000803e3d6000fd5b5050505081801561341357508015155b15612158576066546000848152606a602052604090819020815163242a6e3f60e01b81526004810187815260248201938452825460026000196001831615610100020190911604604483018190526001600160a01b039095169463242a6e3f948994939091606490910190849080156134cd5780601f106134a2576101008083540402835291602001916134cd565b820191906000526020600020905b8154815290600101906020018083116134b057829003601f168201915b50509350505050600060405180830381600087803b1580156134ee57600080fd5b505af1158015613502573d6000803e3d6000fd5b50505050505050565b3360009081526069602052604090205461208181611a9e565b607260209081526000928352604080842090915290825290205481565b6000611cd683835b6001600160a01b038216600090815260736020908152604080832084845290915281206002015415801590611cd657506001600160a01b03831660009081526073602090815260408083208584529091529020600201546135a8613f60565b11159392505050565b6000806000806135c18888612b02565b9050806135d9575060009250600191508290506135e5565b60675490935091508190505b9450945094915050565b60755481565b60009182526077602090815260408084209284526003909201905290205490565b61208181611a9e565b600080600061362c6159a2565b505050506001600160a01b03919091166000908152607360209081526040808320938352928152908290208251608081018452815481526001820154928101839052600282015493810184905260039091015460609091018190529092565b33816136cc576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b6136d68185613549565b1561371c576040805162461bcd60e51b81526020600482015260116024820152700616c7265616479206c6f636b656420757607c1b604482015290519081900360640190fd5b6137268185611c56565b82111561376d576040805162461bcd60e51b815260206004820152601060248201526f6e6f7420656e6f756768207374616b6560801b604482015290519081900360640190fd5b600084815260686020526040902054156137c7576040805162461bcd60e51b815260206004820152601660248201527576616c696461746f722069736e27742061637469766560501b604482015290519081900360640190fd5b6137cf611c4f565b83101580156137e557506137e1611c47565b8311155b61382b576040805162461bcd60e51b815260206004820152601260248201527134b731b7b93932b1ba10323ab930ba34b7b760711b604482015290519081900360640190fd5b600061383984611baf613f60565b6000868152606860205260409020600601549091506001600160a01b0390811690831681146138c7576001600160a01b03811660009081526073602090815260408083208984529091529020600201548211156138c75760405162461bcd60e51b8152600401808060200182810382526028815260200180615bf96028913960400191505060405180910390fd5b6138d183876143d1565b506001600160a01b03831660009081526073602090815260408083208984529091529020848155613900612cbd565b6001808301919091556002808301859055600383018890556001600160a01b03861660008181526074602090815260408083208d845282528083208381559586018390559490930155825189815291820188905282518a9391927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e7892908290030190a350505050505050565b60009182526077602090815260408084209284526002909201905290205490565b600081815260686020526040812060060154819081906139d6906001600160a01b03168561361f565b9250925092509193909250565b6139ec336141f1565b613a275760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b600060776000613a35612cbd565b8152602001908152602001600020905060008090505b82811015613aae576000848483818110613a6157fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c860154909350613a9f91508263ffffffff61405816565b600c8501555050600101613a4b565b50613abd6006820184846159ca565b50505050565b60009182526077602090815260408084209284526004909201905290205490565b613aed336141f1565b613b285760405162461bcd60e51b8152600401808060200182810382526029815260200180615b406029913960400191505060405180910390fd5b600060776000613b36612cbd565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015613b9657602002820191906000526020600020905b815481526020019060010190808311613b82575b50505050509050613c1d82828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250614aa192505050565b613c8c828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a918291850190849080828437600092019190915250614bb092505050565b613c94612cbd565b606755613c9f613f60565b600783015550607554600b820155607654600d909101555050505050505050565b6000611fe2613227565b613cd2612e7f565b613d11576040805162461bcd60e51b81526020600482018190526024820152600080516020615b8a833981519152604482015290519081900360640190fd5b612f25816151d0565b336000908152606960205260409020546120818183613d3882612aa9565b61368b565b61208181611b6a565b600080600080600080600080600080613d5d615a15565b5060008b815260686020908152604091829020825160e08101845281548082526001830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a0840152600601546001600160a01b031660c083015260081415613dd7576101008152613df9565b805160801415613dea5760018152613df9565b805160011415613df957600081525b6000613e048d612aa9565b9050816000015182608001518360a0015184604001518560200151856001613e39888a606001516141af90919063ffffffff16565b8960c001518a60c001518393509b509b509b509b509b509b509b509b509b509b5050509193959799509193959799565b303b1590565b600054610100900460ff1680613e885750613e88613e69565b80613e96575060005460ff16155b613ed15760405162461bcd60e51b815260040180806020018281038252602e815260200180615baa602e913960400191505060405180910390fd5b600054610100900460ff16158015613efc576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015612081576000805461ff00191690555050565b4290565b613f6c615981565b613f7683836143d1565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392613fd492611baf9163ffffffff61405816565b905080614017576040805162461bcd60e51b815260206004820152600c60248201526b7a65726f207265776172647360a01b604482015290519081900360640190fd5b6001600160a01b0384166000908152606f60209081526040808320868452909152812081815560018101829055600201556140518161436a565b5092915050565b600082820183811015611cd6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6140bb82614a8a565b614106576040805162461bcd60e51b81526020600482015260176024820152761d985b1a59185d1bdc88191bd95cdb89dd08195e1a5cdd604a1b604482015290519081900360640190fd5b60008281526068602052604090205415614160576040805162461bcd60e51b815260206004820152601660248201527576616c696461746f722069736e27742061637469766560501b604482015290519081900360640190fd5b61416b838383614205565b614174826147fa565b6121585760405162461bcd60e51b8152600401808060200182810382526029815260200180615c216029913960400191505060405180910390fd5b6000611cd683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615271565b6066546001600160a01b0390811691161490565b60008111614248576040805162461bcd60e51b815260206004820152600b60248201526a1e995c9bc8185b5bdd5b9d60aa1b604482015290519081900360640190fd5b61425283836143d1565b506001600160a01b0383166000908152607260209081526040808320858452909152902054614287908263ffffffff61405816565b6001600160a01b03841660009081526072602090815260408083208684528252808320939093556068905220600301546142c7818363ffffffff61405816565b600084815260686020526040902060030155606c546142ec908363ffffffff61405816565b606c5560008381526068602052604090205461431957606d54614315908363ffffffff61405816565b606d555b614324838215613329565b60408051838152905184916001600160a01b038716917f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb9181900360200190a350505050565b606654604080516366e7ea0f60e01b81523060048201526024810184905290516001600160a01b03909216916366e7ea0f9160448082019260009290919082900301818387803b1580156143bd57600080fd5b505af1158015611a8b573d6000803e3d6000fd5b60006143db615981565b6143e58484615308565b90506143f083615440565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f825282812086825282528290208251606081018452815481526001820154928101929092526002015491810191909152614456908261549b565b6001600160a01b0385166000818152606f60209081526040808320888452825280832085518155858301516001808301919091559582015160029182015593835260748252808320888452825291829020825160608101845281548152948101549185019190915290910154908201526144d0908261549b565b6001600160a01b03851660009081526074602090815260408083208784528252918290208351815590830151600180830191909155929091015160029091015591505092915050565b6001600160a01b03841660009081526074602090815260408083208684529091528120548190614561908490614555908763ffffffff61550d16565b9063ffffffff61556616565b6001600160a01b0387166000908152607460209081526040808320898452909152812060010154919250906145a2908590614555908863ffffffff61550d16565b905060028104820160006145ba86614555848a61550d565b6001600160a01b038a1660009081526074602090815260408083208c84529091529020549091506145f1908563ffffffff6141af16565b6001600160a01b038a1660009081526074602090815260408083208c845290915290209081556001015461462590846141af565b6001600160a01b038a1660009081526074602090815260408083208c84529091529020600101558681106146565750855b98975050505050505050565b60008281526068602052604090205415801561467d57508015155b156146aa57600082815260686020526040902060030154606d546146a69163ffffffff6141af16565b606d555b60008281526068602052604090205481111561208157600082815260686020526040902081815560020154614752576146e1612cbd565b6000838152606860205260409020600201556146fb613f60565b6000838152606860209081526040918290206001810184905560020154825190815290810192909252805184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792908290030190a25b60408051828152905183917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25050565b670de0b6b3a764000090565b60008215806147ae57506147aa61478c565b8210155b156147bb575060006147f3565b6147e66001611baf6147cb61478c565b614555866147d761478c565b8a91900363ffffffff61550d16565b9050838111156147f35750825b9392505050565b600061482761480761478c565b61455561481261212b565b61481b86612aa9565b9063ffffffff61550d16565b60008381526068602052604090206003015411159050919050565b6001600160a01b038816600090815260696020526040902054156148ad576040805162461bcd60e51b815260206004820152601860248201527f76616c696461746f7220616c7265616479206578697374730000000000000000604482015290519081900360640190fd5b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180546001600160a01b031916909417909355606a8152919020875161491e92890190615a5b565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf8686604051808381526020018281526020019250505060405180910390a381156149aa576040805183815260208101839052815189927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47928290030190a25b84156149e75760408051868152905188917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25b5050505050505050565b6149f9615981565b614a01615981565b614a0b8484615308565b6001600160a01b0385166000908152606f602090815260408083208784528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150612b38908261549b565b606b8054600101908190556121588382846000614a7a612cbd565b614a82613f60565b600080614842565b600090815260686020526040902060050154151590565b60005b8351811015611a8b57607854828281518110614abc57fe5b6020026020010151118015614ae65750607954838281518110614adb57fe5b602002602001015110155b15614b2757614b09848281518110614afa57fe5b60200260200101516008614662565b614b27848281518110614b1857fe5b60200260200101516000613329565b828181518110614b3357fe5b6020026020010151856004016000868481518110614b4d57fe5b6020026020010151815260200190815260200160002081905550818181518110614b7357fe5b6020026020010151856005016000868481518110614b8d57fe5b602090810291909101810151825281019190915260400160002055600101614aa4565b614bb8615ac9565b6040518060c001604052808551604051908082528060200260200182016040528015614bee578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015614c26578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060776000614c606001614c54612cbd565b9063ffffffff6141af16565b81526020810191909152604001600020600160808401526007810154909150614c87613f60565b1115614ca1578060070154614c9a613f60565b0360808301525b60005b8551811015614d6c578260800151858281518110614cbe57fe5b6020026020010151858381518110614cd257fe5b60200260200101510281614ce257fe5b0483604001518281518110614cf357fe5b602002602001018181525050614d2d83604001518281518110614d1257fe5b6020026020010151846060015161405890919063ffffffff16565b60608401528351614d5f90859083908110614d4457fe5b60200260200101518460a0015161405890919063ffffffff16565b60a0840152600101614ca4565b5060005b8551811015614e3d578260800151858281518110614d8a57fe5b60200260200101518460800151878481518110614da357fe5b60200260200101518a60000160008b8781518110614dbd57fe5b60200260200101518152602001908152602001600020540281614ddc57fe5b040281614de557fe5b0483600001518281518110614df657fe5b602002602001018181525050614e3083600001518281518110614e1557fe5b6020026020010151846020015161405890919063ffffffff16565b6020840152600101614d70565b5060005b85518110156151a8576000614e79846080015160755486600001518581518110614e6757fe5b602002602001015187602001516155a8565b9050614eb5614ea88560a0015186604001518581518110614e9657fe5b602002602001015187606001516155e9565b829063ffffffff61405816565b90506000878381518110614ec557fe5b6020908102919091018101516000818152606890925260408220600601549092506001600160a01b031690614f0184614efc613055565b615646565b6001600160a01b038316600090815260726020908152604080832087845290915290205490915080156150a857600081614f3b8587612b81565b840281614f4457fe5b049050808303614f52615981565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154614f84908490615663565b9050614f8e615981565b614f99836000615663565b6001600160a01b0388166000908152606f602090815260408083208c84528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150614fee908383615754565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d845282529182902082516060810184528154815294810154918501919091529091015490820152615069908383615754565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156150da57816150cd61478c565b8402816150d657fe5b0490505b808a600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508b898151811061511757fe5b60200260200101518a600301600089815260200190815260200160002054018f6003016000898152602001908152602001600020819055508c898151811061515b57fe5b60200260200101518a600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050614e41565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b0381166152155760405162461bcd60e51b8152600401808060200182810382526026815260200180615b1a6026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b600081848411156153005760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156152c55781810151838201526020016152ad565b50505050905090810190601f1680156152f25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b615310615981565b6001600160a01b03831660009081526070602090815260408083208584529091528120549061533e84615440565b9050600061534c868661576f565b9050818111156153595750805b828110156153645750815b6001600160a01b038616600081815260736020908152604080832089845282528083209383526072825280832089845290915281205482549091906153b090839063ffffffff6141af16565b905060006153c484600001548a898861582e565b90506153ce615981565b6153dc828660030154615663565b90506153ea838b8a8961582e565b91506153f4615981565b6153ff836000615663565b905061540d858c898b61582e565b9250615417615981565b615422846000615663565b905061542f838383615754565b9d9c50505050505050505050505050565b6000818152606860205260408120600201541561549357600082815260686020526040902060020154606754101561547b5750606754612ada565b50600081815260686020526040902060020154612ada565b505060675490565b6154a3615981565b60408051606081019091528251845182916154c4919063ffffffff61405816565b81526020016154e48460200151866020015161405890919063ffffffff16565b81526020016155048460400151866040015161405890919063ffffffff16565b90529392505050565b60008261551c57506000611cd9565b8282028284828161552957fe5b0414611cd65760405162461bcd60e51b8152600401808060200182810382526021815260200180615b696021913960400191505060405180910390fd5b6000611cd683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061589c565b6000826155b757506000612b38565b60006155c9868663ffffffff61550d16565b90506155df83614555838763ffffffff61550d16565b9695505050505050565b6000826155f8575060006147f3565b600061560e83614555878763ffffffff61550d16565b905061563d61561b61478c565b61455561562661215d565b61562e61478c565b8591900363ffffffff61550d16565b95945050505050565b6000611cd661565361478c565b614555858563ffffffff61550d16565b61566b615981565b60405180606001604052806000815260200160008152602001600081525090508160001461572657600061569d61215d565b6156a561478c565b03905060006156c56156b5611c47565b614555848763ffffffff61550d16565b905060006156ee6156d461478c565b614555846156e061215d565b8a910163ffffffff61550d16565b90506157136156fb61478c565b61455561570661215d565b899063ffffffff61550d16565b602085018190529003835250611cd99050565b61574961573161478c565b61455561573c61215d565b869063ffffffff61550d16565b604082015292915050565b61575c615981565b612b38615769858561549b565b8361549b565b6001600160a01b03821660009081526073602090815260408083208484529091528120600101546067546157a4858583615901565b156157b2579150611cd99050565b6157bd858584615901565b6157cc57600092505050611cd9565b808211156157df57600092505050611cd9565b80821015615812576002818301046157f8868683615901565b156158085780600101925061580c565b8091505b506157df565b8061582257600092505050611cd9565b60001901949350505050565b600081831061583f57506000612b38565b60008381526077602081815260408084208885526001908101835281852054878652938352818520898652019091529091205461589161587d61478c565b6145558961481b858763ffffffff6141af16565b979650505050505050565b600081836158eb5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156152c55781810151838201526020016152ad565b5060008385816158f757fe5b0495945050505050565b6001600160a01b03831660009081526073602090815260408083208584529091528120600101548210801590612b3857506001600160a01b03841660009081526073602090815260408083208684529091529020600201546159628361596c565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215615a05579160200282015b82811115615a055782358255916020019190600101906159ea565b50615a11929150615aff565b5090565b6040518060e0016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10615a9c57805160ff1916838001178555615a05565b82800160010185558215615a05579182015b82811115615a05578251825591602001919060010190615aae565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b611a9591905b80821115615a115760008155600101615b0556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616c6c6572206973206e6f7420746865204e6f64654472697665724175746820636f6e7472616374536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65646d757374206265206c657373207468616e206f7220657175616c20746f20312e3076616c696461746f72206c6f636b757020706572696f642077696c6c20656e64206561726c69657276616c696461746f7227732064656c65676174696f6e73206c696d69742069732065786365656465646c6f636b6564207374616b652069732067726561746572207468616e207468652077686f6c65207374616b65a265627a7a7231582068a2eef0a6cf0c5b39dc3f21daab05a006fbd4ba3e09bc65ee4da456a539535664736f6c63430005110032"
