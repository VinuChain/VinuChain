// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package driverauth100

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
const odlContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sfc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_driver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDriverAuth\",\"type\":\"address\"}],\"name\":\"migrateTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"updateNetworkRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateNetworkVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"advanceEpochs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateValidatorWeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"updateValidatorPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
const ContractABI = "" +
	"[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"UpdatedBaseRewardPerSec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocksNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UpdatedOfflinePenaltyThreshold\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBaseRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTxRewardWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"getLockedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getLockupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getStashedLockupRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockupExtraReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupBaseReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockedReward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"isLockedUp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeTokenizerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSlashedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteBookAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeDriver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"updateStakeTokenizerAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateLibAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateTreasuryAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateConstsAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"constsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateVoteBookAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTime\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"epochGas\",\"type\":\"uint256\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]" +
	""



// ContractBin is the compiled bytecode used for deploying new contracts.
var oldContractBin = "0x608060405234801561001057600080fd5b50611c06806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80638da5cb5b116100b2578063c0c53b8b11610081578063e08d7e6611610066578063e08d7e66146104f5578063ebdf104c14610565578063f2fde38b146106cb57610136565b8063c0c53b8b14610475578063d6a0c7af146104ba57610136565b80638da5cb5b146103955780638f32d59b146103c6578063a4066fbe146103e2578063b9cc6b1c1461040557610136565b8063267ab446116101095780634feb92f3116100ee5780634feb92f3146102a957806366e7ea0f14610354578063715018a61461038d57610136565b8063267ab446146102595780634ddaf8f21461027657610136565b80630aeeca001461013b57806318f628d41461015a5780631e702f83146101bf578063242a6e3f146101e2575b600080fd5b6101586004803603602081101561015157600080fd5b50356106fe565b005b610158600480360361012081101561017157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001356107e5565b610158600480360360408110156101d557600080fd5b508035906020013561090c565b610158600480360360408110156101f857600080fd5b8135919081019060408101602082013564010000000081111561021a57600080fd5b82018360208201111561022c57600080fd5b8035906020019184600183028401116401000000008311171561024e57600080fd5b5090925090506109f8565b6101586004803603602081101561026f57600080fd5b5035610b27565b6101586004803603602081101561028c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bf3565b61015860048036036101008110156102c057600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516916020810135918101906060810160408201356401000000008111156102fd57600080fd5b82018360208201111561030f57600080fd5b8035906020019184600183028401116401000000008311171561033157600080fd5b919350915080359060208101359060408101359060608101359060800135610cc0565b6101586004803603604081101561036a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610e1b565b610158610f80565b61039d611048565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6103ce611064565b604080519115158252519081900360200190f35b610158600480360360408110156103f857600080fd5b5080359060200135611082565b6101586004803603602081101561041b57600080fd5b81019060208101813564010000000081111561043657600080fd5b82018360208201111561044857600080fd5b8035906020019184600183028401116401000000008311171561046a57600080fd5b509092509050611168565b6101586004803603606081101561048b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101358216916040909101351661125e565b610158600480360360408110156104d057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166113b9565b6101586004803603602081101561050b57600080fd5b81019060208101813564010000000081111561052657600080fd5b82018360208201111561053857600080fd5b8035906020019184602083028401116401000000008311171561055a57600080fd5b50909250905061151d565b6101586004803603608081101561057b57600080fd5b81019060208101813564010000000081111561059657600080fd5b8201836020820111156105a857600080fd5b803590602001918460208302840111640100000000831117156105ca57600080fd5b9193909290916020810190356401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b91939092909160208101903564010000000081111561063a57600080fd5b82018360208201111561064c57600080fd5b8035906020019184602083028401116401000000008311171561066e57600080fd5b91939092909160208101903564010000000081111561068c57600080fd5b82018360208201111561069e57600080fd5b803590602001918460208302840111640100000000831117156106c057600080fd5b509092509050611616565b610158600480360360208110156106e157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661181c565b610706611064565b610757576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f0aeeca0000000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff90921691630aeeca009160248082019260009290919082900301818387803b1580156107ca57600080fd5b505af11580156107de573d6000803e3d6000fd5b5050505050565b60675473ffffffffffffffffffffffffffffffffffffffff16331461083b5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b50505050505050505050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146109625760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b1580156109dc57600080fd5b505af11580156109f0573d6000803e3d6000fd5b505050505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610a64576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517f242a6e3f00000000000000000000000000000000000000000000000000000000815260048101868152602482019283526044820185905273ffffffffffffffffffffffffffffffffffffffff9093169263242a6e3f928792879287929091606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b50505050505050565b610b2f611064565b610b80576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f267ab44600000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163267ab4469160248082019260009290919082900301818387803b1580156107ca57600080fd5b610bfb611064565b610c4c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fda7fc24f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151919092169163da7fc24f91602480830192600092919082900301818387803b1580156107ca57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff163314610d165760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156108e957600080fd5b60665473ffffffffffffffffffffffffffffffffffffffff163314610e87576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff838116911614610ee05760405162461bcd60e51b8152600401808060200182810382526021815260200180611b8c6021913960400191505060405180910390fd5b60675473ffffffffffffffffffffffffffffffffffffffff9081169063e30443bc908490610f17908216318563ffffffff61188116565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156109dc57600080fd5b610f88611064565b610fd9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60335473ffffffffffffffffffffffffffffffffffffffff16331490565b60665473ffffffffffffffffffffffffffffffffffffffff1633146110ee576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517fa4066fbe0000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163a4066fbe9160448082019260009290919082900301818387803b1580156109dc57600080fd5b611170611064565b6111c1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6067546040517fb9cc6b1c0000000000000000000000000000000000000000000000000000000081526020600482019081526024820184905273ffffffffffffffffffffffffffffffffffffffff9092169163b9cc6b1c91859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b600054610100900460ff168061127757506112776118e2565b80611285575060005460ff16155b6112c05760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff1615801561132657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61132f826118e8565b6067805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680549287169290911691909117905580156113b357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b6113c1611064565b611412576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff83811691161480611450575073ffffffffffffffffffffffffffffffffffffffff821630145b6114a1576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420534643206f722073656c662061646472657373000000000000000000604482015290519081900360640190fd5b606754604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1633146115735760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b6066546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff16331461166c5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b1580156117fa57600080fd5b505af115801561180e573d6000803e3d6000fd5b505050505050505050505050565b611824611064565b611875576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61187e81611a57565b50565b6000828201838110156118db576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b303b1590565b600054610100900460ff168061190157506119016118e2565b8061190f575060005460ff16155b61194a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff161580156119b057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611a5357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b73ffffffffffffffffffffffffffffffffffffffff8116611aa95760405162461bcd60e51b8152600401808060200182810382526026815260200180611b386026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564726563697069656e74206973206e6f74207468652053464320636f6e747261637463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e7472616374a265627a7a7231582085f0b421e25c5c41c27b69f82881c0fbe44729bdd97b1b955484e0541fc3768f64736f6c63430005110032"
var ContractBin = "0x608060405234801561001057600080fd5b5061336a806100206000396000f3fe6080604052600436106102855760003560e01c80638b0e9f3f11610153578063c65ee0e1116100cb578063d46fa5181161007f578063e08d7e6611610064578063e08d7e6614610acd578063e6f45adf14610b4a578063f2fde38b14610b7d57610285565b8063d46fa51814610aa3578063d96ed50514610ab857610285565b8063cc8343aa116100b0578063cc8343aa146109ff578063cfd4766314610a31578063cfdbb7cd14610a6a57610285565b8063c65ee0e1146109c0578063c7be95de146109ea57610285565b8063a2f6e6bc11610122578063b5d8962711610107578063b5d8962714610907578063b810e41114610972578063c5f956af146109ab57610285565b8063a2f6e6bc1461089b578063a86a056f146108ce57610285565b80638b0e9f3f146107e95780638da5cb5b146107fe5780638f32d59b1461081357806396c7ee461461083c57610285565b8063592fe0c0116102015780637cacb1d6116101b5578063854873e11161019a578063854873e114610702578063860c2750146107a1578063893675c6146107d457610285565b80637cacb1d6146106ba578063841e4561146106cf57610285565b8063670322f8116101e6578063670322f814610657578063715018a61461069057806376671808146106a557610285565b8063592fe0c0146104cf5780635fab23a81461064257610285565b80631f2701521161025857806339b80c001161023d57806339b80c00146103f057806354fd4d5014610452578063550359a01461049c57610285565b80631f2701521461037e57806328f73148146103db57610285565b80630135b1db1461029c5780630e559d82146102e157806310e51e141461031257806318160ddd14610369575b60805461029a906001600160a01b0316610bb0565b005b3480156102a857600080fd5b506102cf600480360360208110156102bf57600080fd5b50356001600160a01b0316610bd9565b60408051918252519081900360200190f35b3480156102ed57600080fd5b506102f6610beb565b604080516001600160a01b039092168252519081900360200190f35b34801561031e57600080fd5b5061029a600480360360c081101561033557600080fd5b508035906020810135906001600160a01b0360408201358116916060810135821691608082013581169160a0013516610bfa565b34801561037557600080fd5b506102cf610d87565b34801561038a57600080fd5b506103bd600480360360608110156103a157600080fd5b506001600160a01b038135169060208101359060400135610d8d565b60408051938452602084019290925282820152519081900360600190f35b3480156103e757600080fd5b506102cf610dbf565b3480156103fc57600080fd5b5061041a6004803603602081101561041357600080fd5b5035610dc5565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b34801561045e57600080fd5b50610467610e07565b604080517fffffff00000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b3480156104a857600080fd5b5061029a600480360360208110156104bf57600080fd5b50356001600160a01b0316610e2c565b3480156104db57600080fd5b5061029a600480360360a08110156104f257600080fd5b81019060208101813564010000000081111561050d57600080fd5b82018360208201111561051f57600080fd5b8035906020019184602083028401116401000000008311171561054157600080fd5b91939092909160208101903564010000000081111561055f57600080fd5b82018360208201111561057157600080fd5b8035906020019184602083028401116401000000008311171561059357600080fd5b9193909290916020810190356401000000008111156105b157600080fd5b8201836020820111156105c357600080fd5b803590602001918460208302840111640100000000831117156105e557600080fd5b91939092909160208101903564010000000081111561060357600080fd5b82018360208201111561061557600080fd5b8035906020019184602083028401116401000000008311171561063757600080fd5b919350915035610ebf565b34801561064e57600080fd5b506102cf61117a565b34801561066357600080fd5b506102cf6004803603604081101561067a57600080fd5b506001600160a01b038135169060200135611180565b34801561069c57600080fd5b5061029a6111c4565b3480156106b157600080fd5b506102cf61127f565b3480156106c657600080fd5b506102cf611288565b3480156106db57600080fd5b5061029a600480360360208110156106f257600080fd5b50356001600160a01b031661128e565b34801561070e57600080fd5b5061072c6004803603602081101561072557600080fd5b5035611321565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561076657818101518382015260200161074e565b50505050905090810190601f1680156107935780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156107ad57600080fd5b5061029a600480360360208110156107c457600080fd5b50356001600160a01b03166113da565b3480156107e057600080fd5b506102f661146d565b3480156107f557600080fd5b506102cf61147c565b34801561080a57600080fd5b506102f6611482565b34801561081f57600080fd5b50610828611491565b604080519115158252519081900360200190f35b34801561084857600080fd5b506108756004803603604081101561085f57600080fd5b506001600160a01b0381351690602001356114a2565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156108a757600080fd5b5061029a600480360360208110156108be57600080fd5b50356001600160a01b03166114d4565b3480156108da57600080fd5b506102cf600480360360408110156108f157600080fd5b506001600160a01b038135169060200135611567565b34801561091357600080fd5b506109316004803603602081101561092a57600080fd5b5035611584565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b34801561097e57600080fd5b506103bd6004803603604081101561099557600080fd5b506001600160a01b0381351690602001356115ca565b3480156109b757600080fd5b506102f66115f6565b3480156109cc57600080fd5b506102cf600480360360208110156109e357600080fd5b5035611605565b3480156109f657600080fd5b506102cf611617565b348015610a0b57600080fd5b5061029a60048036036040811015610a2257600080fd5b5080359060200135151561161d565b348015610a3d57600080fd5b506102cf60048036036040811015610a5457600080fd5b506001600160a01b03813516906020013561184c565b348015610a7657600080fd5b5061082860048036036040811015610a8d57600080fd5b506001600160a01b038135169060200135611869565b348015610aaf57600080fd5b506102f6611900565b348015610ac457600080fd5b506102cf61190f565b348015610ad957600080fd5b5061029a60048036036020811015610af057600080fd5b810190602081018135640100000000811115610b0b57600080fd5b820183602082011115610b1d57600080fd5b80359060200191846020830284011164010000000083111715610b3f57600080fd5b509092509050611915565b348015610b5657600080fd5b5061029a60048036036020811015610b6d57600080fd5b50356001600160a01b0316611a59565b348015610b8957600080fd5b5061029a60048036036020811015610ba057600080fd5b50356001600160a01b0316611aec565b3660008037600080366000845af43d6000803e808015610bcf573d6000f35b3d6000fd5b505050565b60696020526000908152604090205481565b607b546001600160a01b031681565b600054610100900460ff1680610c135750610c13611b51565b80610c21575060005460ff16155b610c5c5760405162461bcd60e51b815260040180806020018281038252602e815260200180613308602e913960400191505060405180910390fd5b600054610100900460ff16158015610cc257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b610ccb82611b57565b6067879055606680546001600160a01b038088167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556080805487841690831617905560818054928616929091169190911790556076869055610d32611cb9565b607e55610d3d611cc2565b6000888152607760205260409020600701558015610d7e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050505050565b60765481565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b606d5481565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b7f33303400000000000000000000000000000000000000000000000000000000005b90565b610e34611491565b610e85576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610ec833611cc6565b610f035760405162461bcd60e51b81526004018080602001828103825260298152602001806132be6029913960400191505060405180910390fd5b600060776000610f1161127f565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015610f7157602002820191906000526020600020905b815481526020019060010190808311610f5d575b50505050509050610ff882828d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611cdd92505050565b60675460009081526077602052604090206007810154600190611019611cc2565b111561103057816007015461102c611cc2565b0390505b6110b2818584868d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611ee492505050565b6110bc81866126ce565b50506110c661127f565b6067556110d1611cc2565b6007830155608154604080517fd9a7c1f900000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163d9a7c1f991600480820192602092909190829003018186803b15801561113457600080fd5b505afa158015611148573d6000803e3d6000fd5b505050506040513d602081101561115e57600080fd5b5051600b83015550607654600d90910155505050505050505050565b606e5481565b600061118c8383611869565b611198575060006111be565b506001600160a01b03821660009081526073602090815260408083208484529091529020545b92915050565b6111cc611491565b61121d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60675460010190565b60675481565b611296611491565b6112e7576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b607f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156113d25780601f106113a7576101008083540402835291602001916113d2565b820191906000526020600020905b8154815290600101906020018083116113b557829003601f168201915b505050505081565b6113e2611491565b611433576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6082546001600160a01b031681565b606c5481565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6114dc611491565b61152d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b607b80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b607f546001600160a01b031681565b607a6020526000908152604090205481565b606b5481565b61162682612847565b611677576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b60008281526068602052604090206003810154905415611695575060005b606654604080517fa4066fbe000000000000000000000000000000000000000000000000000000008152600481018690526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b15801561170257600080fd5b505af1158015611716573d6000803e3d6000fd5b5050505081801561172657508015155b15610bd4576066546000848152606a60205260409081902081517f242a6e3f0000000000000000000000000000000000000000000000000000000081526004810187815260248201938452825460027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001831615610100020190911604604483018190526001600160a01b039095169463242a6e3f948994939091606490910190849080156118175780601f106117ec57610100808354040283529160200191611817565b820191906000526020600020905b8154815290600101906020018083116117fa57829003601f168201915b50509350505050600060405180830381600087803b15801561183857600080fd5b505af1158015610d7e573d6000803e3d6000fd5b607260209081526000928352604080842090915290825290205481565b6001600160a01b0382166000908152607360209081526040808320848452909152812060020154158015906118c057506001600160a01b038316600090815260736020908152604080832085845290915290205415155b80156118f957506001600160a01b03831660009081526073602090815260408083208584529091529020600201546118f6611cc2565b11155b9392505050565b6081546001600160a01b031690565b607e5481565b61191e33611cc6565b6119595760405162461bcd60e51b81526004018080602001828103825260298152602001806132be6029913960400191505060405180910390fd5b60006077600061196761127f565b8152602001908152602001600020905060008090505b828110156119e057600084848381811061199357fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c8601549093506119d191508263ffffffff61285e16565b600c850155505060010161197d565b506119ef6006820184846131e6565b50606654607e54604080517f07aaf3440000000000000000000000000000000000000000000000000000000081526004810192909252516001600160a01b03909216916307aaf3449160248082019260009290919082900301818387803b15801561183857600080fd5b611a61611491565b611ab2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b608080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b611af4611491565b611b45576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611b4e816128b8565b50565b303b1590565b600054610100900460ff1680611b705750611b70611b51565b80611b7e575060005460ff16155b611bb95760405162461bcd60e51b815260040180806020018281038252602e815260200180613308602e913960400191505060405180910390fd5b600054610100900460ff16158015611c1f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611cb557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b64174876e80090565b4290565b6066546001600160a01b038281169116145b919050565b60005b8351811015611edd57608160009054906101000a90046001600160a01b03166001600160a01b0316635a68f01a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d3757600080fd5b505afa158015611d4b573d6000803e3d6000fd5b505050506040513d6020811015611d6157600080fd5b50518251839083908110611d7157fe5b6020026020010151118015611e135750608160009054906101000a90046001600160a01b03166001600160a01b031662cc7f836040518163ffffffff1660e01b815260040160206040518083038186803b158015611dce57600080fd5b505afa158015611de2573d6000803e3d6000fd5b505050506040513d6020811015611df857600080fd5b50518351849083908110611e0857fe5b602002602001015110155b15611e5457611e36848281518110611e2757fe5b60200260200101516008612971565b611e54848281518110611e4557fe5b6020026020010151600061161d565b828181518110611e6057fe5b6020026020010151856004016000868481518110611e7a57fe5b6020026020010151815260200190815260200160002081905550818181518110611ea057fe5b6020026020010151856005016000868481518110611eba57fe5b602090810291909101810151825281019190915260400160002055600101611ce0565b5050505050565b611eec61322d565b6040518060a001604052808551604051908082528060200260200182016040528015611f22578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015611f5a578160200160208202803883390190505b508152602001600081526020016000815250905060008090505b8451811015612075576000866003016000878481518110611f9157fe5b60200260200101518152602001908152602001600020549050600080905081858481518110611fbc57fe5b60200260200101511115611fe35781858481518110611fd757fe5b60200260200101510390505b89868481518110611ff057fe5b602002602001015182028161200157fe5b048460400151848151811061201257fe5b60200260200101818152505061204c8460400151848151811061203157fe5b6020026020010151856060015161285e90919063ffffffff16565b60608501526080840151612066908263ffffffff61285e16565b60808501525050600101611f74565b5060005b845181101561213e578784828151811061208f57fe5b6020026020010151898684815181106120a457fe5b60200260200101518a60000160008a87815181106120be57fe5b602002602001015181526020019081526020016000205402816120dd57fe5b0402816120e657fe5b04826000015182815181106120f757fe5b6020026020010181815250506121318260000151828151811061211657fe5b6020026020010151836020015161285e90919063ffffffff16565b6020830152600101612079565b5060005b845181101561257d5760006121eb89608160009054906101000a90046001600160a01b03166001600160a01b031663d9a7c1f96040518163ffffffff1660e01b815260040160206040518083038186803b15801561219f57600080fd5b505afa1580156121b3573d6000803e3d6000fd5b505050506040513d60208110156121c957600080fd5b505185518051869081106121d957fe5b60200260200101518660200151612a9b565b905061222761221a84608001518560400151858151811061220857fe5b60200260200101518660600151612aea565b829063ffffffff61285e16565b9050600086838151811061223757fe5b60209081029190910181015160008181526068835260408082206006015460815482517fa778651500000000000000000000000000000000000000000000000000000000815292519496506001600160a01b039182169593946122ed948994929093169263a77865159260048082019391829003018186803b1580156122bc57600080fd5b505afa1580156122d0573d6000803e3d6000fd5b505050506040513d60208110156122e657600080fd5b5051612c53565b6001600160a01b03831660009081526072602090815260408083208784529091529020549091508015612494576000816123278587611180565b84028161233057fe5b04905080830361233e61325c565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154612370908490612c70565b905061237a61325c565b612385836000612c70565b6001600160a01b0388166000908152606f602090815260408083208c845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506123da908383612e32565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d845282529182902082516060810184528154815294810154918501919091529091015490820152612455908383612e32565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b6000848152606860205260408120600301548387039181156124c657816124b9612e4d565b8402816124c257fe5b0490505b808e600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508a898151811061250357fe5b60200260200101518f6003016000898152602001908152602001600020819055508b898151811061253057fe5b60200260200101518e600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050612142565b50608081015160088701819055602082015160098801556060820151600a88015560765411156125bb576008860154607680549190910390556125c1565b60006076555b607f546001600160a01b031615610d7e5760006125dc612e4d565b608160009054906101000a90046001600160a01b03166001600160a01b03166394c3e9146040518163ffffffff1660e01b815260040160206040518083038186803b15801561262a57600080fd5b505afa15801561263e573d6000803e3d6000fd5b505050506040513d602081101561265457600080fd5b50516080840151028161266357fe5b04905061266f81612e59565b607f546040516001600160a01b03909116908290600081818185875af1925050503d80600081146126bc576040519150601f19603f3d011682016040523d82523d6000602084013e6126c1565b606091505b5050505050505050505050565b608154604080517f3a3ef66c00000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691633a3ef66c916004808301926020929190829003018186803b15801561272c57600080fd5b505afa158015612740573d6000803e3d6000fd5b505050506040513d602081101561275657600080fd5b50518302600101905060008161276a612e4d565b84028161277357fe5b0490506000608160009054906101000a90046001600160a01b03166001600160a01b0316632c8c36a56040518163ffffffff1660e01b815260040160206040518083038186803b1580156127c657600080fd5b505afa1580156127da573d6000803e3d6000fd5b505050506040513d60208110156127f057600080fd5b505190508481016127ff612e4d565b8202838702018161280c57fe5b04915061281882612ef7565b91506000612824612e4d565b83607e54028161283057fe5b04905061283c81612f65565b607e55505050505050565b600090815260686020526040902060050154151590565b6000828201838110156118f9576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6001600160a01b0381166128fd5760405162461bcd60e51b81526004018080602001828103825260268152602001806132986026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008281526068602052604090205415801561298c57508015155b156129b957600082815260686020526040902060030154606d546129b59163ffffffff612f9b16565b606d555b600082815260686020526040902054811115611cb557600082815260686020526040902081815560020154612a61576129f061127f565b600083815260686020526040902060020155612a0a611cc2565b6000838152606860209081526040918290206001810184905560020154825190815290810192909252805184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e4792908290030190a25b60408051828152905183917fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e919081900360200190a25050565b600082612aaa57506000612ae2565b6000612abc868663ffffffff612fdd16565b9050612ade83612ad2838763ffffffff612fdd16565b9063ffffffff61303616565b9150505b949350505050565b600082612af9575060006118f9565b6000612b0f83612ad2878763ffffffff612fdd16565b9050612c4a612b1c612e4d565b608154604080517f94c3e9140000000000000000000000000000000000000000000000000000000081529051612ad2926001600160a01b0316916394c3e914916004808301926020929190829003018186803b158015612b7b57600080fd5b505afa158015612b8f573d6000803e3d6000fd5b505050506040513d6020811015612ba557600080fd5b5051608154604080517fc74dd62100000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163c74dd62191600480820192602092909190829003018186803b158015612c0557600080fd5b505afa158015612c19573d6000803e3d6000fd5b505050506040513d6020811015612c2f57600080fd5b5051612c39612e4d565b030384612fdd90919063ffffffff16565b95945050505050565b60006118f9612c60612e4d565b612ad2858563ffffffff612fdd16565b612c7861325c565b60405180606001604052806000815260200160008152602001600081525090506000608160009054906101000a90046001600160a01b03166001600160a01b0316635e2308d26040518163ffffffff1660e01b815260040160206040518083038186803b158015612ce857600080fd5b505afa158015612cfc573d6000803e3d6000fd5b505050506040513d6020811015612d1257600080fd5b505190508215612e0a57600081612d27612e4d565b0390506000612db9608160009054906101000a90046001600160a01b03166001600160a01b0316630d4955e36040518163ffffffff1660e01b815260040160206040518083038186803b158015612d7d57600080fd5b505afa158015612d91573d6000803e3d6000fd5b505050506040513d6020811015612da757600080fd5b5051612ad2848863ffffffff612fdd16565b90506000612dda612dc8612e4d565b612ad28987860163ffffffff612fdd16565b9050612df7612de7612e4d565b612ad2898763ffffffff612fdd16565b602086018190529003845250612e2b9050565b612e25612e15612e4d565b612ad2868463ffffffff612fdd16565b60408301525b5092915050565b612e3a61325c565b612ae2612e478585613078565b83613078565b670de0b6b3a764000090565b606654604080517f66e7ea0f0000000000000000000000000000000000000000000000000000000081523060048201526024810184905290516001600160a01b03909216916366e7ea0f9160448082019260009290919082900301818387803b158015612ec557600080fd5b505af1158015612ed9573d6000803e3d6000fd5b5050607654612ef1925090508263ffffffff61285e16565b60765550565b60006064612f03612e4d565b60690281612f0d57fe5b04821115612f31576064612f1f612e4d565b60690281612f2957fe5b049050611cd8565b6064612f3b612e4d565b605f0281612f4557fe5b04821015612f61576064612f57612e4d565b605f0281612f2957fe5b5090565b600066038d7ea4c68000821115612f84575066038d7ea4c68000611cd8565b633b9aca00821015612f615750633b9aca00611cd8565b60006118f983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506130ea565b600082612fec575060006111be565b82820282848281612ff957fe5b04146118f95760405162461bcd60e51b81526004018080602001828103825260218152602001806132e76021913960400191505060405180910390fd5b60006118f983836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613181565b61308061325c565b60408051606081019091528251845182916130a1919063ffffffff61285e16565b81526020016130c18460200151866020015161285e90919063ffffffff16565b81526020016130e18460400151866040015161285e90919063ffffffff16565b90529392505050565b600081848411156131795760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561313e578181015183820152602001613126565b50505050905090810190601f16801561316b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600081836131d05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561313e578181015183820152602001613126565b5060008385816131dc57fe5b0495945050505050565b828054828255906000526020600020908101928215613221579160200282015b82811115613221578235825591602001919060010190613206565b50612f6192915061327d565b6040518060a0016040528060608152602001600081526020016060815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b610e2991905b80821115612f61576000815560010161328356fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737363616c6c6572206973206e6f7420746865204e6f64654472697665724175746820636f6e7472616374536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a72315820b6cbc530854e975384a8519774c0438027faf38c6dab8e5c94a931c0018bcde564736f6c63430005110032"


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

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactor) AdvanceEpochs(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "advanceEpochs", num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
//
// Solidity: function advanceEpochs(uint256 num) returns()
func (_Contract *ContractTransactorSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactor) CopyCode(opts *bind.TransactOpts, acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "copyCode", acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
}

// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
//
// Solidity: function copyCode(address acc, address from) returns()
func (_Contract *ContractTransactorSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
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

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractTransactor) IncBalance(opts *bind.TransactOpts, acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "incBalance", acc, diff)
}

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
}

// IncBalance is a paid mutator transaction binding the contract method 0x66e7ea0f.
//
// Solidity: function incBalance(address acc, uint256 diff) returns()
func (_Contract *ContractTransactorSession) IncBalance(acc common.Address, diff *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncBalance(&_Contract.TransactOpts, acc, diff)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _sfc, _driver, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _sfc, address _driver, address _owner) returns()
func (_Contract *ContractTransactorSession) Initialize(_sfc common.Address, _driver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _sfc, _driver, _owner)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractTransactor) MigrateTo(opts *bind.TransactOpts, newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "migrateTo", newDriverAuth)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
}

// MigrateTo is a paid mutator transaction binding the contract method 0x4ddaf8f2.
//
// Solidity: function migrateTo(address newDriverAuth) returns()
func (_Contract *ContractTransactorSession) MigrateTo(newDriverAuth common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MigrateTo(&_Contract.TransactOpts, newDriverAuth)
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

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sealEpoch", offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_Contract *ContractTransactorSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
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
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, _auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGenesisValidator", _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
//
// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
func (_Contract *ContractTransactorSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
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

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactor) UpdateNetworkRules(opts *bind.TransactOpts, diff []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkRules", diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
//
// Solidity: function updateNetworkRules(bytes diff) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactor) UpdateNetworkVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateNetworkVersion", version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
//
// Solidity: function updateNetworkVersion(uint256 version) returns()
func (_Contract *ContractTransactorSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactor) UpdateValidatorPubkey(opts *bind.TransactOpts, validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorPubkey", validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
//
// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactor) UpdateValidatorWeight(opts *bind.TransactOpts, validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateValidatorWeight", validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
}

// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
//
// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
func (_Contract *ContractTransactorSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
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

var ContractBinRuntime = "0x608060405234801561001057600080fd5b50600436106101365760003560e01c80638da5cb5b116100b2578063c0c53b8b11610081578063e08d7e6611610066578063e08d7e66146104f5578063ebdf104c14610565578063f2fde38b146106cb57610136565b8063c0c53b8b14610475578063d6a0c7af146104ba57610136565b80638da5cb5b146103955780638f32d59b146103c6578063a4066fbe146103e2578063b9cc6b1c1461040557610136565b8063267ab446116101095780634feb92f3116100ee5780634feb92f3146102a957806366e7ea0f14610354578063715018a61461038d57610136565b8063267ab446146102595780634ddaf8f21461027657610136565b80630aeeca001461013b57806318f628d41461015a5780631e702f83146101bf578063242a6e3f146101e2575b600080fd5b6101586004803603602081101561015157600080fd5b50356106fe565b005b610158600480360361012081101561017157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001356107e5565b610158600480360360408110156101d557600080fd5b508035906020013561090c565b610158600480360360408110156101f857600080fd5b8135919081019060408101602082013564010000000081111561021a57600080fd5b82018360208201111561022c57600080fd5b8035906020019184600183028401116401000000008311171561024e57600080fd5b5090925090506109f8565b6101586004803603602081101561026f57600080fd5b5035610b27565b6101586004803603602081101561028c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bf3565b61015860048036036101008110156102c057600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516916020810135918101906060810160408201356401000000008111156102fd57600080fd5b82018360208201111561030f57600080fd5b8035906020019184600183028401116401000000008311171561033157600080fd5b919350915080359060208101359060408101359060608101359060800135610cc0565b6101586004803603604081101561036a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610e1b565b610158610f80565b61039d611048565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6103ce611064565b604080519115158252519081900360200190f35b610158600480360360408110156103f857600080fd5b5080359060200135611082565b6101586004803603602081101561041b57600080fd5b81019060208101813564010000000081111561043657600080fd5b82018360208201111561044857600080fd5b8035906020019184600183028401116401000000008311171561046a57600080fd5b509092509050611168565b6101586004803603606081101561048b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101358216916040909101351661125e565b610158600480360360408110156104d057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166113b9565b6101586004803603602081101561050b57600080fd5b81019060208101813564010000000081111561052657600080fd5b82018360208201111561053857600080fd5b8035906020019184602083028401116401000000008311171561055a57600080fd5b50909250905061151d565b6101586004803603608081101561057b57600080fd5b81019060208101813564010000000081111561059657600080fd5b8201836020820111156105a857600080fd5b803590602001918460208302840111640100000000831117156105ca57600080fd5b9193909290916020810190356401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b91939092909160208101903564010000000081111561063a57600080fd5b82018360208201111561064c57600080fd5b8035906020019184602083028401116401000000008311171561066e57600080fd5b91939092909160208101903564010000000081111561068c57600080fd5b82018360208201111561069e57600080fd5b803590602001918460208302840111640100000000831117156106c057600080fd5b509092509050611616565b610158600480360360208110156106e157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661181c565b610706611064565b610757576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f0aeeca0000000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff90921691630aeeca009160248082019260009290919082900301818387803b1580156107ca57600080fd5b505af11580156107de573d6000803e3d6000fd5b5050505050565b60675473ffffffffffffffffffffffffffffffffffffffff16331461083b5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b50505050505050505050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146109625760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b1580156109dc57600080fd5b505af11580156109f0573d6000803e3d6000fd5b505050505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610a64576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517f242a6e3f00000000000000000000000000000000000000000000000000000000815260048101868152602482019283526044820185905273ffffffffffffffffffffffffffffffffffffffff9093169263242a6e3f928792879287929091606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b50505050505050565b610b2f611064565b610b80576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f267ab44600000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163267ab4469160248082019260009290919082900301818387803b1580156107ca57600080fd5b610bfb611064565b610c4c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fda7fc24f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151919092169163da7fc24f91602480830192600092919082900301818387803b1580156107ca57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff163314610d165760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156108e957600080fd5b60665473ffffffffffffffffffffffffffffffffffffffff163314610e87576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff838116911614610ee05760405162461bcd60e51b8152600401808060200182810382526021815260200180611b8c6021913960400191505060405180910390fd5b60675473ffffffffffffffffffffffffffffffffffffffff9081169063e30443bc908490610f17908216318563ffffffff61188116565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156109dc57600080fd5b610f88611064565b610fd9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60335473ffffffffffffffffffffffffffffffffffffffff16331490565b60665473ffffffffffffffffffffffffffffffffffffffff1633146110ee576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517fa4066fbe0000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163a4066fbe9160448082019260009290919082900301818387803b1580156109dc57600080fd5b611170611064565b6111c1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6067546040517fb9cc6b1c0000000000000000000000000000000000000000000000000000000081526020600482019081526024820184905273ffffffffffffffffffffffffffffffffffffffff9092169163b9cc6b1c91859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b600054610100900460ff168061127757506112776118e2565b80611285575060005460ff16155b6112c05760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff1615801561132657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61132f826118e8565b6067805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680549287169290911691909117905580156113b357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b6113c1611064565b611412576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff83811691161480611450575073ffffffffffffffffffffffffffffffffffffffff821630145b6114a1576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420534643206f722073656c662061646472657373000000000000000000604482015290519081900360640190fd5b606754604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1633146115735760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b6066546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff16331461166c5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b1580156117fa57600080fd5b505af115801561180e573d6000803e3d6000fd5b505050505050505050505050565b611824611064565b611875576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61187e81611a57565b50565b6000828201838110156118db576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b303b1590565b600054610100900460ff168061190157506119016118e2565b8061190f575060005460ff16155b61194a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff161580156119b057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611a5357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b73ffffffffffffffffffffffffffffffffffffffff8116611aa95760405162461bcd60e51b8152600401808060200182810382526026815260200180611b386026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564726563697069656e74206973206e6f74207468652053464320636f6e747261637463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e7472616374a265627a7a7231582085f0b421e25c5c41c27b69f82881c0fbe44729bdd97b1b955484e0541fc3768f64736f6c63430005110032"
