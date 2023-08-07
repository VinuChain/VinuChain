//// Code generated - DO NOT EDIT.
//// This file is a generated binding and any manual changes will be lost.
//
package sfc100

//
//import (
//	"errors"
//	"math/big"
//	"strings"
//
//	ethereum "github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/event"
//)
//
//// Reference imports to suppress errors if they are not otherwise used.
//var (
//	_ = errors.New
//	_ = big.NewInt
//	_ = strings.NewReader
//	_ = ethereum.NotFound
//	_ = bind.Bind
//	_ = common.Big1
//	_ = types.BloomLookup
//	_ = event.NewSubscription
//	_ = abi.ConvertType
//)
//
//// ContractMetaData contains all meta data concerning the Contract contract.
//var ContractMetaData = &bind.MetaData{
//	ABI: "[\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"ChangedValidatorStatus\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"ClaimedRewards\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"CreatedValidator\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"DeactivatedValidator\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Delegated\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"receiver\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"string\",\n          \"name\": \"justification\",\n          \"type\": \"string\"\n        }\n      ],\n      \"name\": \"InflatedFTM\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"duration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"LockedUpStake\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"previousOwner\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"newOwner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"OwnershipTransferred\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"RefundedSlashedLegacyDelegation\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"RestakedRewards\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Undelegated\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"penalty\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UnlockedStake\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"value\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedBaseRewardPerSec\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"blocksNum\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"period\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedOfflinePenaltyThreshold\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"refundRatio\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"UpdatedSlashingRefundRatio\",\n      \"type\": \"event\"\n    },\n    {\n      \"anonymous\": false,\n      \"inputs\": [\n        {\n          \"indexed\": true,\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": true,\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"indexed\": false,\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"Withdrawn\",\n      \"type\": \"event\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"syncPubkey\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"name\": \"_syncValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"baseRewardPerSecond\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"claimRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"contractCommission\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"pubkey\",\n          \"type\": \"bytes\"\n        }\n      ],\n      \"name\": \"createValidator\",\n      \"outputs\": [],\n      \"payable\": true,\n      \"stateMutability\": \"payable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"currentEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"currentSealedEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"deactivateValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"delegate\",\n      \"outputs\": [],\n      \"payable\": true,\n      \"stateMutability\": \"payable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedOriginatedTxsFee\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedRewardPerToken\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochAccumulatedUptime\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochOfflineBlocks\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochOfflineTime\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochReceivedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochSnapshot\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"endTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epochFee\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalBaseRewardWeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalTxRewardWeight\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"baseRewardPerSecond\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"totalSupply\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getEpochValidatorIDs\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"\",\n          \"type\": \"uint256[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getLockedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getLockupInfo\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"fromEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"endTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"duration\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getSelfStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"offset\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"limit\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStakes\",\n      \"outputs\": [\n        {\n          \"components\": [\n            {\n              \"internalType\": \"address\",\n              \"name\": \"delegator\",\n              \"type\": \"address\"\n            },\n            {\n              \"internalType\": \"uint96\",\n              \"name\": \"timestamp\",\n              \"type\": \"uint96\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"validatorId\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"amount\",\n              \"type\": \"uint256\"\n            }\n          ],\n          \"internalType\": \"struct SFC.Stake[]\",\n          \"name\": \"\",\n          \"type\": \"tuple[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getStashedLockupRewards\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupExtraReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupBaseReward\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"unlockedReward\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getUnlockedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getValidator\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"receivedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"getValidatorID\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getValidatorPubkey\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"\",\n          \"type\": \"bytes\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getWithdrawalRequest\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"epoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"time\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"offset\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"limit\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"getWrRequests\",\n      \"outputs\": [\n        {\n          \"components\": [\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"epoch\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"time\",\n              \"type\": \"uint256\"\n            },\n            {\n              \"internalType\": \"uint256\",\n              \"name\": \"amount\",\n              \"type\": \"uint256\"\n            }\n          ],\n          \"internalType\": \"struct SFC.WithdrawalRequest[]\",\n          \"name\": \"\",\n          \"type\": \"tuple[]\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"sealedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"_totalSupply\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"nodeDriver\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"address\",\n          \"name\": \"owner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"initialize\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"isLockedUp\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"isOwner\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"isSlashed\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bool\",\n          \"name\": \"\",\n          \"type\": \"bool\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"lastValidatorID\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"lockStake\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"maxDelegatedRatio\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"maxLockupDuration\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"minLockupDuration\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"minSelfStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address payable\",\n          \"name\": \"receiver\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"string\",\n          \"name\": \"justification\",\n          \"type\": \"string\"\n        }\n      ],\n      \"name\": \"mintFTM\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"offlinePenaltyThreshold\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"blocksNum\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"time\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"owner\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"pendingRewards\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"relockStake\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [],\n      \"name\": \"renounceOwnership\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"restakeRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"rewardsStash\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"offlineTime\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"offlineBlocks\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"uptimes\",\n          \"type\": \"uint256[]\"\n        },\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"originatedTxsFee\",\n          \"type\": \"uint256[]\"\n        }\n      ],\n      \"name\": \"sealEpoch\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256[]\",\n          \"name\": \"nextValidatorIDs\",\n          \"type\": \"uint256[]\"\n        }\n      ],\n      \"name\": \"sealEpochValidators\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"stake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockedStake\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupFromEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupEndTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"lockupDuration\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"earlyUnlockPenalty\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"rewards\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"setGenesisDelegation\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"auth\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"bytes\",\n          \"name\": \"pubkey\",\n          \"type\": \"bytes\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"status\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"createdTime\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedEpoch\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"deactivatedTime\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"setGenesisValidator\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"slashingRefundRatio\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"stakeTokenizerAddress\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stakes\",\n      \"outputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint96\",\n          \"name\": \"timestamp\",\n          \"type\": \"uint96\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorId\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"delegator\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stashRewards\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"\",\n          \"type\": \"address\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"stashedRewardsUntilEpoch\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalActiveStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalSlashedStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"totalSupply\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"view\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"newOwner\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"transferOwnership\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"undelegate\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"amount\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"unlockStake\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"unlockedRewardRatio\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"value\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"updateBaseRewardPerSecond\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"blocksNum\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"time\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"updateOfflinePenaltyThreshold\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"validatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"refundRatio\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"updateSlashingRefundRatio\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"address\",\n          \"name\": \"addr\",\n          \"type\": \"address\"\n        }\n      ],\n      \"name\": \"updateStakeTokenizerAddress\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"int256\",\n          \"name\": \"diff\",\n          \"type\": \"int256\"\n        }\n      ],\n      \"name\": \"updateTotalSupply\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"validatorCommission\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"version\",\n      \"outputs\": [\n        {\n          \"internalType\": \"bytes3\",\n          \"name\": \"\",\n          \"type\": \"bytes3\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": false,\n      \"inputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"toValidatorID\",\n          \"type\": \"uint256\"\n        },\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"wrID\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"name\": \"withdraw\",\n      \"outputs\": [],\n      \"payable\": false,\n      \"stateMutability\": \"nonpayable\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"withdrawalPeriodEpochs\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    },\n    {\n      \"constant\": true,\n      \"inputs\": [],\n      \"name\": \"withdrawalPeriodTime\",\n      \"outputs\": [\n        {\n          \"internalType\": \"uint256\",\n          \"name\": \"\",\n          \"type\": \"uint256\"\n        }\n      ],\n      \"payable\": false,\n      \"stateMutability\": \"pure\",\n      \"type\": \"function\"\n    }\n  ]",
//}
//
//// ContractBin is the compiled bytecode used for deploying new contracts.
//var ContractBin = "0x608060405234801561001057600080fd5b5061607580620000216000396000f3fe6080604052600436106104315760003560e01c8063854873e111610229578063b88a37e21161012e578063d5a44f86116100b6578063e08d7e661161007a578063e08d7e6614610c7f578063e261641a14610c9f578063e2f8c33614610cbf578063ebdf104c14610cdf578063f2fde38b14610cff57610431565b8063d5a44f8614610bdb578063d9a7c1f914610c0a578063dc31e1af14610c1f578063de67f21514610c3f578063df00c92214610c5f57610431565b8063c65ee0e1116100fd578063c65ee0e114610b46578063c7be95de14610b66578063cc8343aa14610b7b578063cfd4766314610b9b578063cfdbb7cd14610bbb57610431565b8063b88a37e214610ac4578063bd14d90714610af1578063c3de580e14610b11578063c5f530af14610b3157610431565b8063a198d229116101b1578063a86a056f11610180578063a86a056f14610a1c578063b5d8962714610a3c578063b6d9edd514610a6f578063b810e41114610a8f578063b82b842714610aaf57610431565b8063a198d229146109b4578063a2f6e6bc146109d4578063a5a470ad146109f4578063a778651514610a0757610431565b80638cddb015116101f85780638cddb0151461091a5780638da5cb5b1461093a5780638f32d59b1461094f57806396c7ee46146109715780639fa6dd35146109a157610431565b8063854873e11461088b5780638b0e9f3f146108b85780638b1a0d11146108cd5780638c3c51d8146108ed57610431565b806339b80c001161033a5780636099ecb2116102c25780636f498663116102865780636f498663146107ff578063702797e31461081f578063715018a61461084c57806376671808146108615780637cacb1d61461087657610431565b80636099ecb21461076a57806361e53fcc1461078a578063634b91e3146107aa578063650acd66146107ca578063670322f8146107df57610431565b806354fd4d501161030957806354fd4d50146106f35780635601fe011461071557806358f95b80146107355780635e2308d2146105f35780635fab23a81461075557610431565b806339b80c0014610660578063441a3e70146106935780634f7c4efb146106b35780634feb92f3146106d357610431565b806318f628d4116103bd5780632265f2841161038c5780632265f284146105de5780632709275e146105f357806328f73148146106085780632cedb0971461061d578063346bdcfb1461064057610431565b806318f628d41461054f5780631d3ac42c1461056f5780631e702f831461058f5780631f270152146105af57610431565b80630d4955e3116104045780630d4955e3146104ce5780630d7b2609146104e35780630e559d82146104f857806312622d0e1461051a57806318160ddd1461053a57610431565b80630135b1db14610436578063019e27291461046c57806308c368741461048e5780630962ef79146104ae575b600080fd5b34801561044257600080fd5b50610456610451366004614bb6565b610d1f565b6040516104639190615e05565b60405180910390f35b34801561047857600080fd5b5061048c610487366004615079565b610d31565b005b34801561049a57600080fd5b5061048c6104a936600461500c565b610ed3565b3480156104ba57600080fd5b5061048c6104c936600461500c565b610f9c565b3480156104da57600080fd5b506104566110a2565b3480156104ef57600080fd5b506104566110ab565b34801561050457600080fd5b5061050d6110b2565b6040516104639190615b34565b34801561052657600080fd5b50610456610535366004614c3c565b6110c1565b34801561054657600080fd5b5061045661114a565b34801561055b57600080fd5b5061048c61056a366004614df0565b611150565b34801561057b57600080fd5b5061045661058a36600461505a565b611274565b34801561059b57600080fd5b5061048c6105aa36600461505a565b611399565b3480156105bb57600080fd5b506105cf6105ca366004614d42565b6113f4565b60405161046393929190615e62565b3480156105ea57600080fd5b50610456611426565b3480156105ff57600080fd5b50610456611438565b34801561061457600080fd5b50610456611454565b34801561062957600080fd5b5061063261145a565b604051610463929190615e54565b34801561064c57600080fd5b5061048c61065b36600461500c565b611464565b34801561066c57600080fd5b5061068061067b36600461500c565b6114ae565b6040516104639796959493929190615f1a565b34801561069f57600080fd5b5061048c6106ae36600461505a565b6114f0565b3480156106bf57600080fd5b5061048c6106ce36600461505a565b6117e2565b3480156106df57600080fd5b5061048c6106ee366004614c76565b6118a1565b3480156106ff57600080fd5b50610708611928565b6040516104639190615bc6565b34801561072157600080fd5b5061045661073036600461500c565b611932565b34801561074157600080fd5b5061045661075036600461505a565b611968565b34801561076157600080fd5b50610456611985565b34801561077657600080fd5b50610456610785366004614c3c565b61198b565b34801561079657600080fd5b506104566107a536600461505a565b6119c9565b3480156107b657600080fd5b5061048c6107c536600461505a565b6119ea565b3480156107d657600080fd5b50610456611b8f565b3480156107eb57600080fd5b506104566107fa366004614c3c565b611b94565b34801561080b57600080fd5b5061045661081a366004614c3c565b611bd5565b34801561082b57600080fd5b5061083f61083a366004614d8f565b611c3f565b6040516104639190615b96565b34801561085857600080fd5b5061048c611d2e565b34801561086d57600080fd5b50610456611d9c565b34801561088257600080fd5b50610456611da5565b34801561089757600080fd5b506108ab6108a636600461500c565b611dab565b6040516104639190615bd4565b3480156108c457600080fd5b50610456611e46565b3480156108d957600080fd5b5061048c6108e836600461505a565b611e4c565b3480156108f957600080fd5b5061090d61090836600461505a565b611eb8565b6040516104639190615b85565b34801561092657600080fd5b5061048c610935366004614c3c565b61202a565b34801561094657600080fd5b5061050d612050565b34801561095b57600080fd5b5061096461205f565b6040516104639190615bb8565b34801561097d57600080fd5b5061099161098c366004614c3c565b612070565b6040516104639493929190615e7d565b61048c6109af36600461500c565b6120a2565b3480156109c057600080fd5b506104566109cf36600461505a565b6120ad565b3480156109e057600080fd5b5061048c6109ef366004614bb6565b6120ce565b61048c610a02366004614fd6565b612114565b348015610a1357600080fd5b506104566121a5565b348015610a2857600080fd5b50610456610a37366004614c3c565b6121b4565b348015610a4857600080fd5b50610a5c610a5736600461500c565b6121d1565b6040516104639796959493929190615eb2565b348015610a7b57600080fd5b5061048c610a8a36600461500c565b612217565b348015610a9b57600080fd5b506105cf610aaa366004614c3c565b6122a4565b348015610abb57600080fd5b506104566122d0565b348015610ad057600080fd5b50610ae4610adf36600461500c565b6122d7565b6040516104639190615ba7565b348015610afd57600080fd5b5061048c610b0c3660046150ce565b61233c565b348015610b1d57600080fd5b50610964610b2c36600461500c565b61234f565b348015610b3d57600080fd5b50610456612366565b348015610b5257600080fd5b50610456610b6136600461500c565b612372565b348015610b7257600080fd5b50610456612384565b348015610b8757600080fd5b5061048c610b9636600461502a565b61238a565b348015610ba757600080fd5b50610456610bb6366004614c3c565b6124ba565b348015610bc757600080fd5b50610964610bd6366004614c3c565b6124d7565b348015610be757600080fd5b50610bfb610bf636600461500c565b61256d565b60405161046393929190615b5d565b348015610c1657600080fd5b506104566125b3565b348015610c2b57600080fd5b50610456610c3a36600461505a565b6125b9565b348015610c4b57600080fd5b5061048c610c5a3660046150ce565b6125da565b348015610c6b57600080fd5b50610456610c7a36600461505a565b61262b565b348015610c8b57600080fd5b5061048c610c9a366004614ea4565b61264c565b348015610cab57600080fd5b50610456610cba36600461505a565b612707565b348015610ccb57600080fd5b5061048c610cda366004614bd4565b612728565b348015610ceb57600080fd5b5061048c610cfa366004614ee6565b6127d7565b348015610d0b57600080fd5b5061048c610d1a366004614bb6565b612994565b60696020526000908152604090205481565b600054610100900460ff1680610d4a5750610d4a6129c1565b80610d58575060005460ff16155b610d7d5760405162461bcd60e51b8152600401610d7490615d05565b60405180910390fd5b600054610100900460ff16158015610da8576000805460ff1961ff0019909116610100171660011790555b610db1826129c7565b6067859055606680546001600160a01b0319166001600160a01b03851617905560768490556755cfe697852e904c6075556103e86078556203f480607955610df7612a99565b6000868152607760209081526040808320600701939093558251606081018452828152908101828152928101828152607c80546001810182559352905160029092027f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd08101805494516001600160601b0316600160a01b026001600160a01b039485166001600160a01b03199096169590951790931693909317909155517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd1909101558015610ecc576000805461ff00191690555b5050505050565b33610edc61499b565b610ee68284612a9d565b60208101518151919250600091610f029163ffffffff612b6d16565b9050610f258385610f20856040015185612b6d90919063ffffffff16565b612b92565b6001600160a01b03831660008181526073602090815260408083208884528252918290208054850190558451908501518583015192518894937f4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e2693610f8e939092909190615e62565b60405180910390a350505050565b33610fa561499b565b610faf8284612a9d565b90506000826001600160a01b0316610fec8360400151610fe085602001518660000151612b6d90919063ffffffff16565b9063ffffffff612b6d16565b604051610ff890615b29565b60006040518083038185875af1925050503d8060008114611035576040519150601f19603f3d011682016040523d82523d6000602084013e61103a565b606091505b505090508061105b5760405162461bcd60e51b8152600401610d7490615d45565b81516020830151604080850151905187936001600160a01b038816937fc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf4593610f8e93615e62565b6301e133805b90565b6212750090565b607b546001600160a01b031681565b60006110cd83836124d7565b6110fb57506001600160a01b0382166000908152607260209081526040808320848452909152902054611144565b6001600160a01b0383166000818152607360209081526040808320868452825280832054938352607282528083208684529091529020546111419163ffffffff612c1316565b90505b92915050565b60765481565b61115933612c55565b6111755760405162461bcd60e51b8152600401610d7490615c35565b611180898989612c69565b6001600160a01b0389166000908152606f602090815260408083208b845290915290206002018190556111b287612ef1565b851561126957868611156111d85760405162461bcd60e51b8152600401610d7490615df5565b6001600160a01b03891660008181526073602090815260408083208c845282528083208a8155600181018a90556002810189905560038101889055848452607483528184208d855290925291829020859055905190918a917f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e789061125f9088908c90615e54565b60405180910390a3505b505050505050505050565b3360008181526073602090815260408083208684529091528120909190836112ae5760405162461bcd60e51b8152600401610d7490615bf5565b6112b882866124d7565b6112d45760405162461bcd60e51b8152600401610d7490615ca5565b80548411156112f55760405162461bcd60e51b8152600401610d7490615d75565b6112ff8286612f6f565b61131b5760405162461bcd60e51b8152600401610d7490615d85565b611325828661300c565b50600061133883878785600001546131d7565b82548690038355905061134c838783613308565b85836001600160a01b03167fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b98784604051611388929190615e54565b60405180910390a395945050505050565b6113a233612c55565b6113be5760405162461bcd60e51b8152600401610d7490615c35565b806113db5760405162461bcd60e51b8152600401610d7490615c75565b6113e58282613461565b6113f082600061238a565b5050565b607160209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b6000611430613580565b601002905090565b60006064611444613580565b601e028161144e57fe5b04905090565b606d5481565b6078546079549091565b61146c61205f565b6114885760405162461bcd60e51b8152600401610d7490615cf5565b6000811261149d5760768054820190556114ab565b607680546000839003900390555b50565b607760205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b336114f961499b565b506001600160a01b0381166000908152607160209081526040808320868452825280832085845282529182902082516060810184528154808252600183015493820193909352600290910154928101929092526115685760405162461bcd60e51b8152600401610d7490615d55565b6115728285612f6f565b61158e5760405162461bcd60e51b8152600401610d7490615d85565b602080820151825160008781526068909352604090922060010154909190158015906115ca575060008681526068602052604090206001015482115b156115eb575050600084815260686020526040902060018101546002909101545b6115f36122d0565b82016115fd612a99565b101561161b5760405162461bcd60e51b8152600401610d7490615c85565b611623611b8f565b810161162d611d9c565b101561164b5760405162461bcd60e51b8152600401610d7490615dc5565b6001600160a01b03841660009081526071602090815260408083208984528252808320888452909152812060020154906116848861234f565b905060006116a68383607a60008d81526020019081526020016000205461358c565b6001600160a01b03881660009081526071602090815260408083208d845282528083208c845290915281208181556001810182905560020155606e80548201905590508083116117085760405162461bcd60e51b8152600401610d7490615de5565b60006001600160a01b038816611724858463ffffffff612c1316565b60405161173090615b29565b60006040518083038185875af1925050503d806000811461176d576040519150601f19603f3d011682016040523d82523d6000602084013e611772565b606091505b50509050806117935760405162461bcd60e51b8152600401610d7490615d45565b888a896001600160a01b03167f75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21876040516117ce9190615e05565b60405180910390a450505050505050505050565b6117ea61205f565b6118065760405162461bcd60e51b8152600401610d7490615cf5565b61180f8261234f565b61182b5760405162461bcd60e51b8152600401610d7490615be5565b611833613580565b8111156118525760405162461bcd60e51b8152600401610d7490615d15565b6000828152607a6020526040908190208290555182907f047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb91790611895908490615e05565b60405180910390a25050565b6118aa33612c55565b6118c65760405162461bcd60e51b8152600401610d7490615c35565b61190e898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a915089905088886135ee565b606b5488111561126957606b889055505050505050505050565b620ccc0d60ea1b90565b6000818152606860209081526040808320600601546001600160a01b03168352607282528083208484529091529020545b919050565b600091825260776020908152604080842092845291905290205490565b606e5481565b600061199561499b565b61199f8484613764565b8051602082015160408301519293506119c192610fe09163ffffffff612b6d16565b949350505050565b60009182526077602090815260408084209284526001909201905290205490565b336119f5818461300c565b5060008211611a165760405162461bcd60e51b8152600401610d7490615bf5565b611a2081846110c1565b821115611a3f5760405162461bcd60e51b8152600401610d7490615d25565b611a498184612f6f565b611a655760405162461bcd60e51b8152600401610d7490615d85565b6001600160a01b0381166000908152607e602090815260408083208684529091529020805460018101909155611a9c828585613308565b6001600160a01b038216600090815260716020908152604080832087845282528083208484529091529020600201839055611ad5611d9c565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152902055611b09612a99565b6001600160a01b03831660009081526071602090815260408083208884528252808320858452909152812060010191909155611b4690859061238a565b8084836001600160a01b03167fd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d5786604051611b819190615e05565b60405180910390a450505050565b600390565b6000611ba083836124d7565b611bac57506000611144565b506001600160a01b03919091166000908152607360209081526040808320938352929052205490565b6000611bdf61499b565b506001600160a01b0383166000908152606f60209081526040808320858452825291829020825160608101845281548082526001830154938201849052600290920154938101849052926119c1929091610fe0919063ffffffff612b6d16565b60608082604051908082528060200260200182016040528015611c7c57816020015b611c6961499b565b815260200190600190039081611c615790505b50905060005b83811015611d24576001600160a01b0387166000908152607160209081526040808320898452909152812090611cbe878463ffffffff612b6d16565b81526020019081526020016000206040518060600160405290816000820154815260200160018201548152602001600282015481525050828281518110611d0157fe5b6020908102919091010152611d1d81600163ffffffff612b6d16565b9050611c82565b5095945050505050565b611d3661205f565b611d525760405162461bcd60e51b8152600401610d7490615cf5565b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380546001600160a01b0319169055565b60675460010190565b60675481565b606a6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015611e3e5780601f10611e1357610100808354040283529160200191611e3e565b820191906000526020600020905b815481529060010190602001808311611e2157829003601f168201915b505050505081565b606c5481565b611e5461205f565b611e705760405162461bcd60e51b8152600401610d7490615cf5565b607981905560788290556040517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a3490611eac9084908490615e54565b60405180910390a15050565b607c54604080518381526020808502820101909152606091908290848015611efa57816020015b611ee76149bc565b815260200190600190039081611edf5790505b50905060005b848110156120215782611f19878363ffffffff612b6d16565b10611f2357612021565b6000607c82880181548110611f3457fe5b60009182526020822060029091020154607c80546001600160a01b03909216935090898501908110611f6257fe5b90600052602060002090600202016001015490506040518060800160405280836001600160a01b03168152602001607c858b0181548110611f9f57fe5b600091825260208083206002909202909101546001600160601b03600160a01b9091041683528281018590526001600160a01b038616825260728152604080832086845290915290819020549101528451859085908110611ffc57fe5b602090810291909101015261201883600163ffffffff612b6d16565b92505050611f00565b50949350505050565b612034828261300c565b6113f05760405162461bcd60e51b8152600401610d7490615c45565b6033546001600160a01b031690565b6033546001600160a01b0316331490565b607360209081526000928352604080842090915290825290208054600182015460028301546003909301549192909184565b6114ab338234612b92565b60009182526077602090815260408084209284526005909201905290205490565b6120d661205f565b6120f25760405162461bcd60e51b8152600401610d7490615cf5565b607b80546001600160a01b0319166001600160a01b0392909216919091179055565b61211c612366565b34101561213b5760405162461bcd60e51b8152600401610d7490615dd5565b806121585760405162461bcd60e51b8152600401610d7490615d35565b6121983383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506137d292505050565b6113f033606b5434612b92565b60006121af613580565b905090565b607060209081526000928352604080842090915290825290205481565b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b61221f61205f565b61223b5760405162461bcd60e51b8152600401610d7490615cf5565b6801c985c8903591eb208111156122645760405162461bcd60e51b8152600401610d7490615cb5565b60758190556040517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae239690612299908390615e05565b60405180910390a150565b607460209081526000928352604080842090915290825290208054600182015460029092015490919083565b62093a8090565b60008181526077602090815260409182902060060180548351818402810184019094528084526060939283018282801561233057602002820191906000526020600020905b81548152602001906001019080831161231c575b50505050509050919050565b33612349818585856137fd565b50505050565b600090815260686020526040902054608016151590565b670467fc915c2fc00090565b607a6020526000908152604090205481565b606b5481565b612393826139da565b6123af5760405162461bcd60e51b8152600401610d7490615db5565b600082815260686020526040902060038101549054156123cd575060005b60665460405163520337df60e11b81526001600160a01b039091169063a4066fbe906123ff9086908590600401615e54565b600060405180830381600087803b15801561241957600080fd5b505af115801561242d573d6000803e3d6000fd5b5050505081801561243d57508015155b156124b5576066546000848152606a602052604090819020905163242a6e3f60e01b81526001600160a01b039092169163242a6e3f9161248291879190600401615e13565b600060405180830381600087803b15801561249c57600080fd5b505af11580156124b0573d6000803e3d6000fd5b505050505b505050565b607260209081526000928352604080842090915290825290205481565b6001600160a01b03821660009081526073602090815260408083208484529091528120600201541580159061252e57506001600160a01b038316600090815260736020908152604080832085845290915290205415155b801561114157506001600160a01b0383166000908152607360209081526040808320858452909152902060020154612564612a99565b11159392505050565b607c818154811061257a57fe5b6000918252602090912060029091020180546001909101546001600160a01b0382169250600160a01b9091046001600160601b03169083565b60755481565b60009182526077602090815260408084209284526003909201905290205490565b33816125f85760405162461bcd60e51b8152600401610d7490615bf5565b61260281856124d7565b1561261f5760405162461bcd60e51b8152600401610d7490615c15565b612349818585856137fd565b60009182526077602090815260408084209284526002909201905290205490565b61265533612c55565b6126715760405162461bcd60e51b8152600401610d7490615c35565b60006077600061267f611d9c565b8152602001908152602001600020905060008090505b828110156126f85760008484838181106126ab57fe5b60209081029290920135600081815260688452604080822060030154948890529020839055600c8601549093506126e991508263ffffffff612b6d16565b600c8501555050600101612695565b506123496006820184846149f6565b60009182526077602090815260408084209284526004909201905290205490565b61273061205f565b61274c5760405162461bcd60e51b8152600401610d7490615cf5565b61275583612ef1565b6040516001600160a01b0385169084156108fc029085906000818181858888f1935050505015801561278b573d6000803e3d6000fd5b50836001600160a01b03167f9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad8484846040516127c993929190615e33565b60405180910390a250505050565b6127e033612c55565b6127fc5760405162461bcd60e51b8152600401610d7490615c35565b60006077600061280a611d9c565b8152602001908152602001600020905060608160060180548060200260200160405190810160405280929190818152602001828054801561286a57602002820191906000526020600020905b815481526020019060010190808311612856575b505050505090506128f182828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506139f192505050565b612960828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a918291850190849080828437600092019190915250613b0092505050565b612968611d9c565b606755612973612a99565b600783015550607554600b820155607654600d909101555050505050505050565b61299c61205f565b6129b85760405162461bcd60e51b8152600401610d7490615cf5565b6114ab81614146565b303b1590565b600054610100900460ff16806129e057506129e06129c1565b806129ee575060005460ff16155b612a0a5760405162461bcd60e51b8152600401610d7490615d05565b600054610100900460ff16158015612a35576000805460ff1961ff0019909116610100171660011790555b603380546001600160a01b0319166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a380156113f0576000805461ff00191690555050565b4290565b612aa561499b565b612aaf838361300c565b50506001600160a01b0382166000908152606f6020908152604080832084845282528083208151606081018352815480825260018301549482018590526002909201549281018390529392612b0d92610fe09163ffffffff612b6d16565b905080612b2c5760405162461bcd60e51b8152600401610d7490615cd5565b6001600160a01b0384166000908152606f6020908152604080832086845290915281208181556001810182905560020155612b6681612ef1565b5092915050565b6000828201838110156111415760405162461bcd60e51b8152600401610d7490615c25565b612b9b826139da565b612bb75760405162461bcd60e51b8152600401610d7490615db5565b60008281526068602052604090205415612be35760405162461bcd60e51b8152600401610d7490615c65565b612bee838383612c69565b612bf7826141c8565b6124b55760405162461bcd60e51b8152600401610d7490615da5565b600061114183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250614210565b6066546001600160a01b0390811691161490565b60008111612c895760405162461bcd60e51b8152600401610d7490615bf5565b612c93838361300c565b506001600160a01b0383166000908152607d6020908152604080832085845290915290205480612d9357607c80546001600160a01b038087166000818152607d602090815260408083208a8452825280832086905580516060810182529384526001600160601b034281169285019283529084018a815260018701885596909252915160029094027f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd0810180549351909216600160a01b029484166001600160a01b03199093169290921790921692909217905590517f9222cbf5d0ddc505a6f2f04716e22c226cee16a955fef88c618922096dae2fd190910155612dd7565b42607c8281548110612da157fe5b906000526020600020906002020160000160146101000a8154816001600160601b0302191690836001600160601b031602179055505b6001600160a01b0384166000908152607260209081526040808320868452909152902054612e0b908363ffffffff612b6d16565b6001600160a01b0385166000908152607260209081526040808320878452825280832093909355606890522060030154612e4b818463ffffffff612b6d16565b600085815260686020526040902060030155606c54612e70908463ffffffff612b6d16565b606c55600084815260686020526040902054612e9d57606d54612e99908463ffffffff612b6d16565b606d555b612ea884821561238a565b83856001600160a01b03167f9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb85604051612ee29190615e05565b60405180910390a35050505050565b6066546040516366e7ea0f60e01b81526001600160a01b03909116906366e7ea0f90612f239030908590600401615b42565b600060405180830381600087803b158015612f3d57600080fd5b505af1158015612f51573d6000803e3d6000fd5b5050607654612f69925090508263ffffffff612b6d16565b60765550565b607b546000906001600160a01b0316612f8a57506001611144565b607b546040516321d585c360e01b81526001600160a01b03909116906321d585c390612fbc9086908690600401615b42565b60206040518083038186803b158015612fd457600080fd5b505afa158015612fe8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506111419190810190614fb8565b600061301661499b565b613020848461423c565b905061302b83614374565b6001600160a01b0385166000818152607060209081526040808320888452825280832094909455918152606f82528281208682528252829020825160608101845281548152600182015492810192909252600201549181019190915261309190826143cf565b6001600160a01b0385166000818152606f602090815260408083208884528252808320855181558583015160018083019190915595820151600291820155938352607482528083208884528252918290208251606081018452815481529481015491850191909152909101549082015261310b90826143cf565b6001600160a01b03851660009081526074602090815260408083208784528252918290208351815590830151600182015591015160029091015561314f84846124d7565b6131b2576001600160a01b0384166000818152607360209081526040808320878452825280832083815560018082018590556002808301869055600390920185905594845260748352818420888552909252822082815592830182905591909101555b60208101511515806131c45750805115155b806119c157506040015115159392505050565b6001600160a01b0384166000908152607460209081526040808320868452909152812054819061321f908490613213908763ffffffff61444116565b9063ffffffff61447b16565b6001600160a01b038716600090815260746020908152604080832089845290915281206001015491925090613260908590613213908863ffffffff61444116565b6001600160a01b03881660009081526074602090815260408083208a8452909152902054909150600282048301906132989084612c13565b6001600160a01b03891660009081526074602090815260408083208b84529091529020908155600101546132cc9083612c13565b6001600160a01b03891660009081526074602090815260408083208b84529091529020600101558581106132fd5750845b979650505050505050565b6001600160a01b0383166000908152607260209081526040808320858452825280832080548590039055606890915290206003015461334d908263ffffffff612c1316565b600083815260686020526040902060030155606c54613372908263ffffffff612c1316565b606c556001600160a01b03831660009081526072602090815260408083208584529091529020546133cc576001600160a01b0383166000908152607d602090815260408083208584529091529020546133ca816144bd565b505b6000828152606860205260409020546133f657606d546133f2908263ffffffff612c1316565b606d555b600061340183611932565b9050801561345a57613411612366565b8110156134305760405162461bcd60e51b8152600401610d7490615dd5565b613439836141c8565b6134555760405162461bcd60e51b8152600401610d7490615da5565b612349565b6123498360015b60008281526068602052604090205415801561347c57508015155b156134a957600082815260686020526040902060030154606d546134a59163ffffffff612c1316565b606d555b6000828152606860205260409020548111156113f057600082815260686020526040902081815560020154613550576134e0611d9c565b6000838152606860205260409020600201556134fa612a99565b600083815260686020526040908190206001810183905560020154905184927fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e479261354792909190615e54565b60405180910390a25b817fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e826040516118959190615e05565b670de0b6b3a764000090565b60008215806135a2575061359e613580565b8210155b156135af575060006135e7565b6135da6001610fe06135bf613580565b613213866135cb613580565b8a91900363ffffffff61444116565b9050838111156135e75750825b9392505050565b6001600160a01b038816600090815260696020526040902054156136245760405162461bcd60e51b8152600401610d7490615c95565b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180546001600160a01b031916909417909355606a8152919020875161369592890190614a41565b50876001600160a01b0316877f49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf86866040516136d2929190615e54565b60405180910390a3811561371b57867fac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e478383604051613712929190615e54565b60405180910390a25b841561375a57867fcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e866040516137519190615e05565b60405180910390a25b5050505050505050565b61376c61499b565b61377461499b565b61377e848461423c565b6001600160a01b0385166000908152606f6020908152604080832087845282529182902082516060810184528154815260018201549281019290925260020154918101919091529091506119c190826143cf565b606b8054600101908190556124b583828460006137ed611d9c565b6137f5612a99565b6000806135ee565b61380784846110c1565b8111156138265760405162461bcd60e51b8152600401610d7490615d95565b600083815260686020526040902054156138525760405162461bcd60e51b8152600401610d7490615c65565b61385a6110ab565b8210158015613870575061386c6110a2565b8211155b61388c5760405162461bcd60e51b8152600401610d7490615c55565b600061389a83610fe0612a99565b6000858152606860205260409020600601549091506001600160a01b039081169086168114613909576001600160a01b03811660009081526073602090815260408083208884529091529020600201548211156139095760405162461bcd60e51b8152600401610d7490615d65565b613913868661300c565b506001600160a01b03861660009081526073602090815260408083208884529091529020600381015485101561395b5760405162461bcd60e51b8152600401610d7490615ce5565b805461396d908563ffffffff612b6d16565b8155613977611d9c565b6001820155600281018390556003810185905560405186906001600160a01b038916907f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78906139c99089908990615e54565b60405180910390a350505050505050565b600090815260686020526040902060050154151590565b60005b8351811015610ecc57607854828281518110613a0c57fe5b6020026020010151118015613a365750607954838281518110613a2b57fe5b602002602001015110155b15613a7757613a59848281518110613a4a57fe5b60200260200101516008613461565b613a77848281518110613a6857fe5b6020026020010151600061238a565b828181518110613a8357fe5b6020026020010151856004016000868481518110613a9d57fe5b6020026020010151815260200190815260200160002081905550818181518110613ac357fe5b6020026020010151856005016000868481518110613add57fe5b6020908102919091018101518252810191909152604001600020556001016139f4565b613b08614aaf565b6040518060c001604052808551604051908082528060200260200182016040528015613b3e578160200160208202803883390190505b508152602001600081526020018551604051908082528060200260200182016040528015613b76578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060776000613bb06001613ba4611d9c565b9063ffffffff612c1316565b81526020810191909152604001600020600160808401526007810154909150613bd7612a99565b1115613bf1578060070154613bea612a99565b0360808301525b60005b8551811015613cf9576000826003016000888481518110613c1157fe5b60200260200101518152602001908152602001600020549050600080905081868481518110613c3c57fe5b60200260200101511115613c635781868481518110613c5757fe5b60200260200101510390505b8460800151878481518110613c7457fe5b6020026020010151820281613c8557fe5b0485604001518481518110613c9657fe5b602002602001018181525050613cd085604001518481518110613cb557fe5b60200260200101518660600151612b6d90919063ffffffff16565b606086015260a0850151613cea908263ffffffff612b6d16565b60a08601525050600101613bf4565b5060005b8551811015613dca578260800151858281518110613d1757fe5b60200260200101518460800151878481518110613d3057fe5b60200260200101518a60000160008b8781518110613d4a57fe5b60200260200101518152602001908152602001600020540281613d6957fe5b040281613d7257fe5b0483600001518281518110613d8357fe5b602002602001018181525050613dbd83600001518281518110613da257fe5b60200260200101518460200151612b6d90919063ffffffff16565b6020840152600101613cfd565b5060005b855181101561411e576000613e06846080015160755486600001518581518110613df457fe5b602002602001015187602001516145fb565b9050613e42613e358560a0015186604001518581518110613e2357fe5b6020026020010151876060015161463c565b829063ffffffff612b6d16565b90506000878381518110613e5257fe5b6020908102919091018101516000818152606890925260408220600601549092506001600160a01b031690613e8e84613e896121a5565b614699565b6001600160a01b0383166000908152607260209081526040808320878452909152902054909150801561403557600081613ec88587611b94565b840281613ed157fe5b049050808303613edf61499b565b6001600160a01b03861660009081526073602090815260408083208a8452909152902060030154613f119084906146b6565b9050613f1b61499b565b613f268360006146b6565b6001600160a01b0388166000908152606f602090815260408083208c84528252918290208251606081018452815481526001820154928101929092526002015491810191909152909150613f7b9083836147a7565b6001600160a01b0388166000818152606f602090815260408083208d84528252808320855181558583015160018083019190915595820151600291820155938352607482528083208d845282529182902082516060810184528154815294810154918501919091529091015490820152613ff69083836147a7565b6001600160a01b03881660009081526074602090815260408083208c845282529182902083518155908301516001820155910151600290910155505050505b600084815260686020526040812060030154838703918115614067578161405a613580565b84028161406357fe5b0490505b808a600101600089815260200190815260200160002054018f6001016000898152602001908152602001600020819055508b89815181106140a457fe5b60200260200101518f6003016000898152602001908152602001600020819055508c89815181106140d157fe5b60200260200101518a600201600089815260200190815260200160002054018f60020160008981526020019081526020016000208190555050505050505050508080600101915050613dce565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b03811661416c5760405162461bcd60e51b8152600401610d7490615c05565b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b60006141f56141d5613580565b6132136141e0611426565b6141e986611932565b9063ffffffff61444116565b60008381526068602052604090206003015411159050919050565b600081848411156142345760405162461bcd60e51b8152600401610d749190615bd4565b505050900390565b61424461499b565b6001600160a01b03831660009081526070602090815260408083208584529091528120549061427284614374565b9050600061428086866147c2565b90508181111561428d5750805b828110156142985750815b6001600160a01b038616600081815260736020908152604080832089845282528083209383526072825280832089845290915281205482549091906142e490839063ffffffff612c1316565b905060006142f884600001548a8988614881565b905061430261499b565b6143108286600301546146b6565b905061431e838b8a89614881565b915061432861499b565b6143338360006146b6565b9050614341858c898b614881565b925061434b61499b565b6143568460006146b6565b90506143638383836147a7565b9d9c50505050505050505050505050565b600081815260686020526040812060020154156143c75760008281526068602052604090206002015460675410156143af5750606754611963565b50600081815260686020526040902060020154611963565b505060675490565b6143d761499b565b60408051606081019091528251845182916143f8919063ffffffff612b6d16565b815260200161441884602001518660200151612b6d90919063ffffffff16565b815260200161443884604001518660400151612b6d90919063ffffffff16565b90529392505050565b60008261445057506000611144565b8282028284828161445d57fe5b04146111415760405162461bcd60e51b8152600401610d7490615cc5565b600061114183836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506148e4565b607c546000906144d490600163ffffffff612c1316565b90508082146145ce57607c81815481106144ea57fe5b9060005260206000209060020201607c838154811061450557fe5b600091825260208220835460029092020180546001600160a01b0319166001600160a01b039283161780825584546001600160601b03600160a01b9182900416029216919091178155600192830154920191909155607c80548492607d9290918590811061456f57fe5b600091825260208083206002909202909101546001600160a01b031683528201929092526040018120607c8054919291859081106145a957fe5b9060005260206000209060020201600101548152602001908152602001600020819055505b607c8054806145d957fe5b6000828152602081206002600019909301928302018181556001015590555050565b60008261460a575060006119c1565b600061461c868663ffffffff61444116565b905061463283613213838763ffffffff61444116565b9695505050505050565b60008261464b575060006135e7565b600061466183613213878763ffffffff61444116565b905061469061466e613580565b613213614679611438565b614681613580565b8591900363ffffffff61444116565b95945050505050565b60006111416146a6613580565b613213858563ffffffff61444116565b6146be61499b565b6040518060600160405280600081526020016000815260200160008152509050816000146147795760006146f0611438565b6146f8613580565b03905060006147186147086110a2565b613213848763ffffffff61444116565b90506000614741614727613580565b61321384614733611438565b8a910163ffffffff61444116565b905061476661474e613580565b613213614759611438565b899063ffffffff61444116565b6020850181905290038352506111449050565b61479c614784613580565b61321361478f611438565b869063ffffffff61444116565b604082015292915050565b6147af61499b565b6119c16147bc85856143cf565b836143cf565b6001600160a01b03821660009081526073602090815260408083208484529091528120600101546067546147f785858361491b565b156148055791506111449050565b61481085858461491b565b61481f57600092505050611144565b8082111561483257600092505050611144565b808210156148655760028183010461484b86868361491b565b1561485b5780600101925061485f565b8091505b50614832565b8061487557600092505050611144565b60001901949350505050565b6000818310614892575060006119c1565b6000838152607760208181526040808420888552600190810183528185205487865293835281852089865201909152909120546132fd6148d0613580565b613213896141e9858763ffffffff612c1316565b600081836149055760405162461bcd60e51b8152600401610d749190615bd4565b50600083858161491157fe5b0495945050505050565b6001600160a01b038316600090815260736020908152604080832085845290915281206001015482108015906119c157506001600160a01b038416600090815260736020908152604080832086845290915290206002015461497c83614986565b1115949350505050565b60009081526077602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b604051806080016040528060006001600160a01b0316815260200160006001600160601b0316815260200160008152602001600081525090565b828054828255906000526020600020908101928215614a31579160200282015b82811115614a31578235825591602001919060010190614a16565b50614a3d929150614ae5565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614a8257805160ff1916838001178555614a31565b82800160010185558215614a31579182015b82811115614a31578251825591602001919060010190614a94565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b6110a891905b80821115614a3d5760008155600101614aeb565b80356111448161600c565b60008083601f840112614b1c57600080fd5b50813567ffffffffffffffff811115614b3457600080fd5b602083019150836020820283011115614b4c57600080fd5b9250929050565b803561114481616020565b805161114481616020565b60008083601f840112614b7b57600080fd5b50813567ffffffffffffffff811115614b9357600080fd5b602083019150836001820283011115614b4c57600080fd5b803561114481616029565b600060208284031215614bc857600080fd5b60006119c18484614aff565b60008060008060608587031215614bea57600080fd5b6000614bf68787614aff565b9450506020614c0787828801614bab565b935050604085013567ffffffffffffffff811115614c2457600080fd5b614c3087828801614b69565b95989497509550505050565b60008060408385031215614c4f57600080fd5b6000614c5b8585614aff565b9250506020614c6c85828601614bab565b9150509250929050565b60008060008060008060008060006101008a8c031215614c9557600080fd5b6000614ca18c8c614aff565b9950506020614cb28c828d01614bab565b98505060408a013567ffffffffffffffff811115614ccf57600080fd5b614cdb8c828d01614b69565b97509750506060614cee8c828d01614bab565b9550506080614cff8c828d01614bab565b94505060a0614d108c828d01614bab565b93505060c0614d218c828d01614bab565b92505060e0614d328c828d01614bab565b9150509295985092959850929598565b600080600060608486031215614d5757600080fd5b6000614d638686614aff565b9350506020614d7486828701614bab565b9250506040614d8586828701614bab565b9150509250925092565b60008060008060808587031215614da557600080fd5b6000614db18787614aff565b9450506020614dc287828801614bab565b9350506040614dd387828801614bab565b9250506060614de487828801614bab565b91505092959194509250565b60008060008060008060008060006101208a8c031215614e0f57600080fd5b6000614e1b8c8c614aff565b9950506020614e2c8c828d01614bab565b9850506040614e3d8c828d01614bab565b9750506060614e4e8c828d01614bab565b9650506080614e5f8c828d01614bab565b95505060a0614e708c828d01614bab565b94505060c0614e818c828d01614bab565b93505060e0614e928c828d01614bab565b925050610100614d328c828d01614bab565b60008060208385031215614eb757600080fd5b823567ffffffffffffffff811115614ece57600080fd5b614eda85828601614b0a565b92509250509250929050565b6000806000806000806000806080898b031215614f0257600080fd5b883567ffffffffffffffff811115614f1957600080fd5b614f258b828c01614b0a565b9850985050602089013567ffffffffffffffff811115614f4457600080fd5b614f508b828c01614b0a565b9650965050604089013567ffffffffffffffff811115614f6f57600080fd5b614f7b8b828c01614b0a565b9450945050606089013567ffffffffffffffff811115614f9a57600080fd5b614fa68b828c01614b0a565b92509250509295985092959890939650565b600060208284031215614fca57600080fd5b60006119c18484614b5e565b60008060208385031215614fe957600080fd5b823567ffffffffffffffff81111561500057600080fd5b614eda85828601614b69565b60006020828403121561501e57600080fd5b60006119c18484614bab565b6000806040838503121561503d57600080fd5b60006150498585614bab565b9250506020614c6c85828601614b53565b6000806040838503121561506d57600080fd5b6000614c5b8585614bab565b6000806000806080858703121561508f57600080fd5b600061509b8787614bab565b94505060206150ac87828801614bab565b93505060406150bd87828801614aff565b9250506060614de487828801614aff565b6000806000606084860312156150e357600080fd5b6000614d638686614bab565b60006150fb8383615a9a565b505060800190565b600061510f8383615ae4565b505060600190565b60006151238383615b17565b505060200190565b61513481615f95565b82525050565b600061514582615f88565b61514f8185615f8c565b935061515a83615f76565b8060005b8381101561518857815161517288826150ef565b975061517d83615f76565b92505060010161515e565b509495945050505050565b600061519e82615f88565b6151a88185615f8c565b93506151b383615f76565b8060005b838110156151885781516151cb8882615103565b97506151d683615f76565b9250506001016151b7565b60006151ec82615f88565b6151f68185615f8c565b935061520183615f76565b8060005b838110156151885781516152198882615117565b975061522483615f76565b925050600101615205565b61513481615fa0565b61513481615fa5565b600061524c82615f88565b6152568185615f8c565b9350615266818560208601615fd6565b61526f81616002565b9093019392505050565b60008154600181166000811461529657600181146152bc576152fb565b607f60028304166152a78187615f8c565b60ff19841681529550506020850192506152fb565b600282046152ca8187615f8c565b95506152d585615f7c565b60005b828110156152f4578154888201526001909101906020016152d8565b8701945050505b505092915050565b600061530f8385615f8c565b935061531c838584615fca565b61526f83616002565b6000615332601783615f8c565b7f76616c696461746f722069736e277420736c6173686564000000000000000000815260200192915050565b600061536b600b83615f8c565b6a1e995c9bc8185b5bdd5b9d60aa1b815260200192915050565b6000615392602683615f8c565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b60006153da601183615f8c565b700616c7265616479206c6f636b656420757607c1b815260200192915050565b6000615407601b83615f8c565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b6000615440602983615f8c565b7f63616c6c6572206973206e6f7420746865204e6f6465447269766572417574688152680818dbdb9d1c9858dd60ba1b602082015260400192915050565b600061548b601083615f8c565b6f0dcdee8d0d2dcce40e8de40e6e8c2e6d60831b815260200192915050565b60006154b7601283615f8c565b7134b731b7b93932b1ba10323ab930ba34b7b760711b815260200192915050565b60006154e5601683615f8c565b7576616c696461746f722069736e27742061637469766560501b815260200192915050565b6000615517600c83615f8c565b6b77726f6e672073746174757360a01b815260200192915050565b600061553f601683615f8c565b751b9bdd08195b9bdd59da081d1a5b59481c185cdcd95960521b815260200192915050565b6000615571601883615f8c565b7f76616c696461746f7220616c7265616479206578697374730000000000000000815260200192915050565b60006155aa600d83615f8c565b6c06e6f74206c6f636b656420757609c1b815260200192915050565b60006155d3601b83615f8c565b7f746f6f206c617267652072657761726420706572207365636f6e640000000000815260200192915050565b600061560c602183615f8c565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f8152607760f81b602082015260400192915050565b600061564f600c83615f8c565b6b7a65726f207265776172647360a01b815260200192915050565b6000615677601f83615f8c565b7f6c6f636b7570206475726174696f6e2063616e6e6f7420646563726561736500815260200192915050565b60006156b0602083615f8c565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b60006156e9602e83615f8c565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000615739602183615f8c565b7f6d757374206265206c657373207468616e206f7220657175616c20746f20312e8152600360fc1b602082015260400192915050565b600061577c601983615f8c565b7f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000815260200192915050565b60006157b5600c83615f8c565b6b656d707479207075626b657960a01b815260200192915050565b60006157dd601283615f8c565b714661696c656420746f2073656e642046544d60701b815260200192915050565b600061580b601583615f8c565b741c995c5d595cdd08191bd95cdb89dd08195e1a5cdd605a1b815260200192915050565b600061583c602883615f8c565b7f76616c696461746f72206c6f636b757020706572696f642077696c6c20656e648152671032b0b93634b2b960c11b602082015260400192915050565b6000615886601783615f8c565b7f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000815260200192915050565b60006158bf601883615f8c565b7f6f75747374616e64696e67207346544d2062616c616e63650000000000000000815260200192915050565b6000611144600083611963565b6000615905601083615f8c565b6f6e6f7420656e6f756768207374616b6560801b815260200192915050565b6000615931602983615f8c565b7f76616c696461746f7227732064656c65676174696f6e73206c696d697420697381526808195e18d95959195960ba1b602082015260400192915050565b600061597c601783615f8c565b7f76616c696461746f7220646f65736e2774206578697374000000000000000000815260200192915050565b60006159b5601883615f8c565b7f6e6f7420656e6f7567682065706f636873207061737365640000000000000000815260200192915050565b60006159ee601783615f8c565b7f696e73756666696369656e742073656c662d7374616b65000000000000000000815260200192915050565b6000615a27601683615f8c565b751cdd185ad9481a5cc8199d5b1b1e481cdb185cda195960521b815260200192915050565b6000615a59602c83615f8c565b7f6c6f636b6564207374616b652069732067726561746572207468616e2074686581526b2077686f6c65207374616b6560a01b602082015260400192915050565b80516080830190615aab848261512b565b506020820151615abe6020850182615b20565b506040820151615ad16040850182615b17565b5060608201516123496060850182615b17565b80516060830190615af58482615b17565b506020820151615b086020850182615b17565b50604082015161234960408501825b615134816110a8565b61513481615fbe565b6000611144826158eb565b60208101611144828461512b565b60408101615b50828561512b565b6135e76020830184615b17565b60608101615b6b828661512b565b615b786020830185615b20565b6119c16040830184615b17565b60208082528101611141818461513a565b602080825281016111418184615193565b6020808252810161114181846151e1565b60208101611144828461522f565b602081016111448284615238565b602080825281016111418184615241565b6020808252810161114481615325565b602080825281016111448161535e565b6020808252810161114481615385565b60208082528101611144816153cd565b60208082528101611144816153fa565b6020808252810161114481615433565b602080825281016111448161547e565b60208082528101611144816154aa565b60208082528101611144816154d8565b602080825281016111448161550a565b6020808252810161114481615532565b6020808252810161114481615564565b602080825281016111448161559d565b60208082528101611144816155c6565b60208082528101611144816155ff565b6020808252810161114481615642565b602080825281016111448161566a565b60208082528101611144816156a3565b60208082528101611144816156dc565b602080825281016111448161572c565b602080825281016111448161576f565b60208082528101611144816157a8565b60208082528101611144816157d0565b60208082528101611144816157fe565b602080825281016111448161582f565b6020808252810161114481615879565b60208082528101611144816158b2565b60208082528101611144816158f8565b6020808252810161114481615924565b602080825281016111448161596f565b60208082528101611144816159a8565b60208082528101611144816159e1565b6020808252810161114481615a1a565b6020808252810161114481615a4c565b602081016111448284615b17565b60408101615e218285615b17565b81810360208301526119c18184615279565b60408101615e418286615b17565b8181036020830152614690818486615303565b60408101615b508285615b17565b60608101615e708286615b17565b615b786020830185615b17565b60808101615e8b8287615b17565b615e986020830186615b17565b615ea56040830185615b17565b6146906060830184615b17565b60e08101615ec0828a615b17565b615ecd6020830189615b17565b615eda6040830188615b17565b615ee76060830187615b17565b615ef46080830186615b17565b615f0160a0830185615b17565b615f0e60c083018461512b565b98975050505050505050565b60e08101615f28828a615b17565b615f356020830189615b17565b615f426040830188615b17565b615f4f6060830187615b17565b615f5c6080830186615b17565b615f6960a0830185615b17565b615f0e60c0830184615b17565b60200190565b60009081526020902090565b5190565b90815260200190565b600061114482615fb2565b151590565b6001600160e81b03191690565b6001600160a01b031690565b6001600160601b031690565b82818337506000910152565b60005b83811015615ff1578181015183820152602001615fd9565b838111156123495750506000910152565b601f01601f191690565b61601581615f95565b81146114ab57600080fd5b61601581615fa0565b616015816110a856fea365627a7a72315820c16cb1611d08c40d5dd778a4947a6881ec3b8728f3a08d86d7d90bce0a6bb8696c6578706572696d656e74616cf564736f6c63430005110040"
//
//// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
//func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
//	parsed, err := abi.JSON(strings.NewReader(ContractABI))
//	if err != nil {
//		return common.Address{}, nil, nil, err
//	}
//
//	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
//	if err != nil {
//		return common.Address{}, nil, nil, err
//	}
//	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
//}
//
//// SFCStake is an auto generated low-level Go binding around an user-defined struct.
//type SFCStake struct {
//	Delegator   common.Address
//	Timestamp   *big.Int
//	ValidatorId *big.Int
//	Amount      *big.Int
//}
//
//// SFCWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
//type SFCWithdrawalRequest struct {
//	Epoch  *big.Int
//	Time   *big.Int
//	Amount *big.Int
//}
//
//// ContractABI is the input ABI used to generate the binding from.
//// Deprecated: Use ContractMetaData.ABI instead.
//var ContractABI = ContractMetaData.ABI
//
//// Contract is an auto generated Go binding around an Ethereum contract.
//type Contract struct {
//	ContractCaller     // Read-only binding to the contract
//	ContractTransactor // Write-only binding to the contract
//	ContractFilterer   // Log filterer for contract events
//}
//
//// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
//type ContractCaller struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
//type ContractTransactor struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
//type ContractFilterer struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// ContractSession is an auto generated Go binding around an Ethereum contract,
//// with pre-set call and transact options.
//type ContractSession struct {
//	Contract     *Contract         // Generic contract binding to set the session for
//	CallOpts     bind.CallOpts     // Call options to use throughout this session
//	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
//}
//
//// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
//// with pre-set call options.
//type ContractCallerSession struct {
//	Contract *ContractCaller // Generic contract caller binding to set the session for
//	CallOpts bind.CallOpts   // Call options to use throughout this session
//}
//
//// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
//// with pre-set transact options.
//type ContractTransactorSession struct {
//	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
//	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
//}
//
//// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
//type ContractRaw struct {
//	Contract *Contract // Generic contract binding to access the raw methods on
//}
//
//// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
//type ContractCallerRaw struct {
//	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
//}
//
//// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
//type ContractTransactorRaw struct {
//	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
//}
//
//// NewContract creates a new instance of Contract, bound to a specific deployed contract.
//func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
//	contract, err := bindContract(address, backend, backend, backend)
//	if err != nil {
//		return nil, err
//	}
//	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
//}
//
//// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
//func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
//	contract, err := bindContract(address, caller, nil, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractCaller{contract: contract}, nil
//}
//
//// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
//func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
//	contract, err := bindContract(address, nil, transactor, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractTransactor{contract: contract}, nil
//}
//
//// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
//func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
//	contract, err := bindContract(address, nil, nil, filterer)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractFilterer{contract: contract}, nil
//}
//
//// bindContract binds a generic wrapper to an already deployed contract.
//func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
//	parsed, err := ContractMetaData.GetAbi()
//	if err != nil {
//		return nil, err
//	}
//	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
//	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
//	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
//	return _Contract.Contract.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.Contract.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
//	return _Contract.Contract.contract.Transact(opts, method, params...)
//}
//
//// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
////
//// Solidity: function baseRewardPerSecond() view returns(uint256)
//func (_Contract *ContractCaller) BaseRewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "baseRewardPerSecond")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
////
//// Solidity: function baseRewardPerSecond() view returns(uint256)
//func (_Contract *ContractSession) BaseRewardPerSecond() (*big.Int, error) {
//	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
//}
//
//// BaseRewardPerSecond is a free data retrieval call binding the contract method 0xd9a7c1f9.
////
//// Solidity: function baseRewardPerSecond() view returns(uint256)
//func (_Contract *ContractCallerSession) BaseRewardPerSecond() (*big.Int, error) {
//	return _Contract.Contract.BaseRewardPerSecond(&_Contract.CallOpts)
//}
//
//// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
////
//// Solidity: function contractCommission() pure returns(uint256)
//func (_Contract *ContractCaller) ContractCommission(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "contractCommission")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
////
//// Solidity: function contractCommission() pure returns(uint256)
//func (_Contract *ContractSession) ContractCommission() (*big.Int, error) {
//	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
//}
//
//// ContractCommission is a free data retrieval call binding the contract method 0x2709275e.
////
//// Solidity: function contractCommission() pure returns(uint256)
//func (_Contract *ContractCallerSession) ContractCommission() (*big.Int, error) {
//	return _Contract.Contract.ContractCommission(&_Contract.CallOpts)
//}
//
//// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
////
//// Solidity: function currentEpoch() view returns(uint256)
//func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "currentEpoch")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
////
//// Solidity: function currentEpoch() view returns(uint256)
//func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
//	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
//}
//
//// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
////
//// Solidity: function currentEpoch() view returns(uint256)
//func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
//	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
//}
//
//// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
////
//// Solidity: function currentSealedEpoch() view returns(uint256)
//func (_Contract *ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "currentSealedEpoch")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
////
//// Solidity: function currentSealedEpoch() view returns(uint256)
//func (_Contract *ContractSession) CurrentSealedEpoch() (*big.Int, error) {
//	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
//}
//
//// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
////
//// Solidity: function currentSealedEpoch() view returns(uint256)
//func (_Contract *ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
//	return _Contract.Contract.CurrentSealedEpoch(&_Contract.CallOpts)
//}
//
//// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
////
//// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochAccumulatedOriginatedTxsFee(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedOriginatedTxsFee", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
////
//// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
////
//// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
////
//// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochAccumulatedRewardPerToken(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedRewardPerToken", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
////
//// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
////
//// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedRewardPerToken(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
////
//// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochAccumulatedUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochAccumulatedUptime", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
////
//// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
////
//// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochAccumulatedUptime(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
////
//// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochOfflineBlocks(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochOfflineBlocks", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
////
//// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
////
//// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochOfflineBlocks(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
////
//// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochOfflineTime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochOfflineTime", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
////
//// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
////
//// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochOfflineTime(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
////
//// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetEpochReceivedStake(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochReceivedStake", epoch, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
////
//// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
////
//// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetEpochReceivedStake(&_Contract.CallOpts, epoch, validatorID)
//}
//
//// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
////
//// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
//func (_Contract *ContractCaller) GetEpochSnapshot(opts *bind.CallOpts, arg0 *big.Int) (struct {
//	EndTime               *big.Int
//	EpochFee              *big.Int
//	TotalBaseRewardWeight *big.Int
//	TotalTxRewardWeight   *big.Int
//	BaseRewardPerSecond   *big.Int
//	TotalStake            *big.Int
//	TotalSupply           *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochSnapshot", arg0)
//
//	outstruct := new(struct {
//		EndTime               *big.Int
//		EpochFee              *big.Int
//		TotalBaseRewardWeight *big.Int
//		TotalTxRewardWeight   *big.Int
//		BaseRewardPerSecond   *big.Int
//		TotalStake            *big.Int
//		TotalSupply           *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.EpochFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.TotalBaseRewardWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//	outstruct.TotalTxRewardWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
//	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
//	outstruct.TotalStake = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
//	outstruct.TotalSupply = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
////
//// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
//func (_Contract *ContractSession) GetEpochSnapshot(arg0 *big.Int) (struct {
//	EndTime               *big.Int
//	EpochFee              *big.Int
//	TotalBaseRewardWeight *big.Int
//	TotalTxRewardWeight   *big.Int
//	BaseRewardPerSecond   *big.Int
//	TotalStake            *big.Int
//	TotalSupply           *big.Int
//}, error) {
//	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
//}
//
//// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
////
//// Solidity: function getEpochSnapshot(uint256 ) view returns(uint256 endTime, uint256 epochFee, uint256 totalBaseRewardWeight, uint256 totalTxRewardWeight, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
//func (_Contract *ContractCallerSession) GetEpochSnapshot(arg0 *big.Int) (struct {
//	EndTime               *big.Int
//	EpochFee              *big.Int
//	TotalBaseRewardWeight *big.Int
//	TotalTxRewardWeight   *big.Int
//	BaseRewardPerSecond   *big.Int
//	TotalStake            *big.Int
//	TotalSupply           *big.Int
//}, error) {
//	return _Contract.Contract.GetEpochSnapshot(&_Contract.CallOpts, arg0)
//}
//
//// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
////
//// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
//func (_Contract *ContractCaller) GetEpochValidatorIDs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getEpochValidatorIDs", epoch)
//
//	if err != nil {
//		return *new([]*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
//
//	return out0, err
//
//}
//
//// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
////
//// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
//func (_Contract *ContractSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
//	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
//}
//
//// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
////
//// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
//func (_Contract *ContractCallerSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
//	return _Contract.Contract.GetEpochValidatorIDs(&_Contract.CallOpts, epoch)
//}
//
//// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
////
//// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetLockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getLockedStake", delegator, toValidatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
////
//// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// GetLockedStake is a free data retrieval call binding the contract method 0x670322f8.
////
//// Solidity: function getLockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetLockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetLockedStake(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
////
//// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
//func (_Contract *ContractCaller) GetLockupInfo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
//	LockedStake *big.Int
//	FromEpoch   *big.Int
//	EndTime     *big.Int
//	Duration    *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getLockupInfo", arg0, arg1)
//
//	outstruct := new(struct {
//		LockedStake *big.Int
//		FromEpoch   *big.Int
//		EndTime     *big.Int
//		Duration    *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.LockedStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.FromEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
////
//// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
//func (_Contract *ContractSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
//	LockedStake *big.Int
//	FromEpoch   *big.Int
//	EndTime     *big.Int
//	Duration    *big.Int
//}, error) {
//	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetLockupInfo is a free data retrieval call binding the contract method 0x96c7ee46.
////
//// Solidity: function getLockupInfo(address , uint256 ) view returns(uint256 lockedStake, uint256 fromEpoch, uint256 endTime, uint256 duration)
//func (_Contract *ContractCallerSession) GetLockupInfo(arg0 common.Address, arg1 *big.Int) (struct {
//	LockedStake *big.Int
//	FromEpoch   *big.Int
//	EndTime     *big.Int
//	Duration    *big.Int
//}, error) {
//	return _Contract.Contract.GetLockupInfo(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
////
//// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetSelfStake(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getSelfStake", validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
////
//// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
//}
//
//// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
////
//// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetSelfStake(&_Contract.CallOpts, validatorID)
//}
//
//// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
////
//// Solidity: function getStake(address , uint256 ) view returns(uint256)
//func (_Contract *ContractCaller) GetStake(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getStake", arg0, arg1)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
////
//// Solidity: function getStake(address , uint256 ) view returns(uint256)
//func (_Contract *ContractSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
////
//// Solidity: function getStake(address , uint256 ) view returns(uint256)
//func (_Contract *ContractCallerSession) GetStake(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetStake(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
////
//// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint96,uint256,uint256)[])
//func (_Contract *ContractCaller) GetStakes(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]SFCStake, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getStakes", offset, limit)
//
//	if err != nil {
//		return *new([]SFCStake), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new([]SFCStake)).(*[]SFCStake)
//
//	return out0, err
//
//}
//
//// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
////
//// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint96,uint256,uint256)[])
//func (_Contract *ContractSession) GetStakes(offset *big.Int, limit *big.Int) ([]SFCStake, error) {
//	return _Contract.Contract.GetStakes(&_Contract.CallOpts, offset, limit)
//}
//
//// GetStakes is a free data retrieval call binding the contract method 0x8c3c51d8.
////
//// Solidity: function getStakes(uint256 offset, uint256 limit) view returns((address,uint96,uint256,uint256)[])
//func (_Contract *ContractCallerSession) GetStakes(offset *big.Int, limit *big.Int) ([]SFCStake, error) {
//	return _Contract.Contract.GetStakes(&_Contract.CallOpts, offset, limit)
//}
//
//// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
////
//// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractCaller) GetStashedLockupRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
//	LockupExtraReward *big.Int
//	LockupBaseReward  *big.Int
//	UnlockedReward    *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getStashedLockupRewards", arg0, arg1)
//
//	outstruct := new(struct {
//		LockupExtraReward *big.Int
//		LockupBaseReward  *big.Int
//		UnlockedReward    *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.LockupExtraReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.LockupBaseReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.UnlockedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
////
//// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
//	LockupExtraReward *big.Int
//	LockupBaseReward  *big.Int
//	UnlockedReward    *big.Int
//}, error) {
//	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetStashedLockupRewards is a free data retrieval call binding the contract method 0xb810e411.
////
//// Solidity: function getStashedLockupRewards(address , uint256 ) view returns(uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractCallerSession) GetStashedLockupRewards(arg0 common.Address, arg1 *big.Int) (struct {
//	LockupExtraReward *big.Int
//	LockupBaseReward  *big.Int
//	UnlockedReward    *big.Int
//}, error) {
//	return _Contract.Contract.GetStashedLockupRewards(&_Contract.CallOpts, arg0, arg1)
//}
//
//// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
////
//// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCaller) GetUnlockedStake(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getUnlockedStake", delegator, toValidatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
////
//// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// GetUnlockedStake is a free data retrieval call binding the contract method 0x12622d0e.
////
//// Solidity: function getUnlockedStake(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) GetUnlockedStake(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.GetUnlockedStake(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
////
//// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
//func (_Contract *ContractCaller) GetValidator(opts *bind.CallOpts, arg0 *big.Int) (struct {
//	Status           *big.Int
//	DeactivatedTime  *big.Int
//	DeactivatedEpoch *big.Int
//	ReceivedStake    *big.Int
//	CreatedEpoch     *big.Int
//	CreatedTime      *big.Int
//	Auth             common.Address
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getValidator", arg0)
//
//	outstruct := new(struct {
//		Status           *big.Int
//		DeactivatedTime  *big.Int
//		DeactivatedEpoch *big.Int
//		ReceivedStake    *big.Int
//		CreatedEpoch     *big.Int
//		CreatedTime      *big.Int
//		Auth             common.Address
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.DeactivatedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.DeactivatedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//	outstruct.ReceivedStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
//	outstruct.CreatedEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
//	outstruct.CreatedTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
//	outstruct.Auth = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
//
//	return *outstruct, err
//
//}
//
//// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
////
//// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
//func (_Contract *ContractSession) GetValidator(arg0 *big.Int) (struct {
//	Status           *big.Int
//	DeactivatedTime  *big.Int
//	DeactivatedEpoch *big.Int
//	ReceivedStake    *big.Int
//	CreatedEpoch     *big.Int
//	CreatedTime      *big.Int
//	Auth             common.Address
//}, error) {
//	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
//}
//
//// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
////
//// Solidity: function getValidator(uint256 ) view returns(uint256 status, uint256 deactivatedTime, uint256 deactivatedEpoch, uint256 receivedStake, uint256 createdEpoch, uint256 createdTime, address auth)
//func (_Contract *ContractCallerSession) GetValidator(arg0 *big.Int) (struct {
//	Status           *big.Int
//	DeactivatedTime  *big.Int
//	DeactivatedEpoch *big.Int
//	ReceivedStake    *big.Int
//	CreatedEpoch     *big.Int
//	CreatedTime      *big.Int
//	Auth             common.Address
//}, error) {
//	return _Contract.Contract.GetValidator(&_Contract.CallOpts, arg0)
//}
//
//// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
////
//// Solidity: function getValidatorID(address ) view returns(uint256)
//func (_Contract *ContractCaller) GetValidatorID(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getValidatorID", arg0)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
////
//// Solidity: function getValidatorID(address ) view returns(uint256)
//func (_Contract *ContractSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
//	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
//}
//
//// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
////
//// Solidity: function getValidatorID(address ) view returns(uint256)
//func (_Contract *ContractCallerSession) GetValidatorID(arg0 common.Address) (*big.Int, error) {
//	return _Contract.Contract.GetValidatorID(&_Contract.CallOpts, arg0)
//}
//
//// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
////
//// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
//func (_Contract *ContractCaller) GetValidatorPubkey(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getValidatorPubkey", arg0)
//
//	if err != nil {
//		return *new([]byte), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
//
//	return out0, err
//
//}
//
//// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
////
//// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
//func (_Contract *ContractSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
//	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
//}
//
//// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
////
//// Solidity: function getValidatorPubkey(uint256 ) view returns(bytes)
//func (_Contract *ContractCallerSession) GetValidatorPubkey(arg0 *big.Int) ([]byte, error) {
//	return _Contract.Contract.GetValidatorPubkey(&_Contract.CallOpts, arg0)
//}
//
//// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
////
//// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
//func (_Contract *ContractCaller) GetWithdrawalRequest(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
//	Epoch  *big.Int
//	Time   *big.Int
//	Amount *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getWithdrawalRequest", arg0, arg1, arg2)
//
//	outstruct := new(struct {
//		Epoch  *big.Int
//		Time   *big.Int
//		Amount *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
////
//// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
//func (_Contract *ContractSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
//	Epoch  *big.Int
//	Time   *big.Int
//	Amount *big.Int
//}, error) {
//	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
//}
//
//// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
////
//// Solidity: function getWithdrawalRequest(address , uint256 , uint256 ) view returns(uint256 epoch, uint256 time, uint256 amount)
//func (_Contract *ContractCallerSession) GetWithdrawalRequest(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
//	Epoch  *big.Int
//	Time   *big.Int
//	Amount *big.Int
//}, error) {
//	return _Contract.Contract.GetWithdrawalRequest(&_Contract.CallOpts, arg0, arg1, arg2)
//}
//
//// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
////
//// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
//func (_Contract *ContractCaller) GetWrRequests(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "getWrRequests", delegator, validatorID, offset, limit)
//
//	if err != nil {
//		return *new([]SFCWithdrawalRequest), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new([]SFCWithdrawalRequest)).(*[]SFCWithdrawalRequest)
//
//	return out0, err
//
//}
//
//// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
////
//// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
//func (_Contract *ContractSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
//	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
//}
//
//// GetWrRequests is a free data retrieval call binding the contract method 0x702797e3.
////
//// Solidity: function getWrRequests(address delegator, uint256 validatorID, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256)[])
//func (_Contract *ContractCallerSession) GetWrRequests(delegator common.Address, validatorID *big.Int, offset *big.Int, limit *big.Int) ([]SFCWithdrawalRequest, error) {
//	return _Contract.Contract.GetWrRequests(&_Contract.CallOpts, delegator, validatorID, offset, limit)
//}
//
//// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
////
//// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
//func (_Contract *ContractCaller) IsLockedUp(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (bool, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "isLockedUp", delegator, toValidatorID)
//
//	if err != nil {
//		return *new(bool), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
//
//	return out0, err
//
//}
//
//// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
////
//// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
//func (_Contract *ContractSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
//	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// IsLockedUp is a free data retrieval call binding the contract method 0xcfdbb7cd.
////
//// Solidity: function isLockedUp(address delegator, uint256 toValidatorID) view returns(bool)
//func (_Contract *ContractCallerSession) IsLockedUp(delegator common.Address, toValidatorID *big.Int) (bool, error) {
//	return _Contract.Contract.IsLockedUp(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "isOwner")
//
//	if err != nil {
//		return *new(bool), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
//
//	return out0, err
//
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractSession) IsOwner() (bool, error) {
//	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
//}
//
//// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
////
//// Solidity: function isOwner() view returns(bool)
//func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
//	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
//}
//
//// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
////
//// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
//func (_Contract *ContractCaller) IsSlashed(opts *bind.CallOpts, validatorID *big.Int) (bool, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "isSlashed", validatorID)
//
//	if err != nil {
//		return *new(bool), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
//
//	return out0, err
//
//}
//
//// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
////
//// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
//func (_Contract *ContractSession) IsSlashed(validatorID *big.Int) (bool, error) {
//	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
//}
//
//// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
////
//// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
//func (_Contract *ContractCallerSession) IsSlashed(validatorID *big.Int) (bool, error) {
//	return _Contract.Contract.IsSlashed(&_Contract.CallOpts, validatorID)
//}
//
//// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
////
//// Solidity: function lastValidatorID() view returns(uint256)
//func (_Contract *ContractCaller) LastValidatorID(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "lastValidatorID")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
////
//// Solidity: function lastValidatorID() view returns(uint256)
//func (_Contract *ContractSession) LastValidatorID() (*big.Int, error) {
//	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
//}
//
//// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
////
//// Solidity: function lastValidatorID() view returns(uint256)
//func (_Contract *ContractCallerSession) LastValidatorID() (*big.Int, error) {
//	return _Contract.Contract.LastValidatorID(&_Contract.CallOpts)
//}
//
//// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
////
//// Solidity: function maxDelegatedRatio() pure returns(uint256)
//func (_Contract *ContractCaller) MaxDelegatedRatio(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "maxDelegatedRatio")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
////
//// Solidity: function maxDelegatedRatio() pure returns(uint256)
//func (_Contract *ContractSession) MaxDelegatedRatio() (*big.Int, error) {
//	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
//}
//
//// MaxDelegatedRatio is a free data retrieval call binding the contract method 0x2265f284.
////
//// Solidity: function maxDelegatedRatio() pure returns(uint256)
//func (_Contract *ContractCallerSession) MaxDelegatedRatio() (*big.Int, error) {
//	return _Contract.Contract.MaxDelegatedRatio(&_Contract.CallOpts)
//}
//
//// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
////
//// Solidity: function maxLockupDuration() pure returns(uint256)
//func (_Contract *ContractCaller) MaxLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "maxLockupDuration")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
////
//// Solidity: function maxLockupDuration() pure returns(uint256)
//func (_Contract *ContractSession) MaxLockupDuration() (*big.Int, error) {
//	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
//}
//
//// MaxLockupDuration is a free data retrieval call binding the contract method 0x0d4955e3.
////
//// Solidity: function maxLockupDuration() pure returns(uint256)
//func (_Contract *ContractCallerSession) MaxLockupDuration() (*big.Int, error) {
//	return _Contract.Contract.MaxLockupDuration(&_Contract.CallOpts)
//}
//
//// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
////
//// Solidity: function minLockupDuration() pure returns(uint256)
//func (_Contract *ContractCaller) MinLockupDuration(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "minLockupDuration")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
////
//// Solidity: function minLockupDuration() pure returns(uint256)
//func (_Contract *ContractSession) MinLockupDuration() (*big.Int, error) {
//	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
//}
//
//// MinLockupDuration is a free data retrieval call binding the contract method 0x0d7b2609.
////
//// Solidity: function minLockupDuration() pure returns(uint256)
//func (_Contract *ContractCallerSession) MinLockupDuration() (*big.Int, error) {
//	return _Contract.Contract.MinLockupDuration(&_Contract.CallOpts)
//}
//
//// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
////
//// Solidity: function minSelfStake() pure returns(uint256)
//func (_Contract *ContractCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "minSelfStake")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
////
//// Solidity: function minSelfStake() pure returns(uint256)
//func (_Contract *ContractSession) MinSelfStake() (*big.Int, error) {
//	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
//}
//
//// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
////
//// Solidity: function minSelfStake() pure returns(uint256)
//func (_Contract *ContractCallerSession) MinSelfStake() (*big.Int, error) {
//	return _Contract.Contract.MinSelfStake(&_Contract.CallOpts)
//}
//
//// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
////
//// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
//func (_Contract *ContractCaller) OfflinePenaltyThreshold(opts *bind.CallOpts) (struct {
//	BlocksNum *big.Int
//	Time      *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "offlinePenaltyThreshold")
//
//	outstruct := new(struct {
//		BlocksNum *big.Int
//		Time      *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.BlocksNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
////
//// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
//func (_Contract *ContractSession) OfflinePenaltyThreshold() (struct {
//	BlocksNum *big.Int
//	Time      *big.Int
//}, error) {
//	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
//}
//
//// OfflinePenaltyThreshold is a free data retrieval call binding the contract method 0x2cedb097.
////
//// Solidity: function offlinePenaltyThreshold() view returns(uint256 blocksNum, uint256 time)
//func (_Contract *ContractCallerSession) OfflinePenaltyThreshold() (struct {
//	BlocksNum *big.Int
//	Time      *big.Int
//}, error) {
//	return _Contract.Contract.OfflinePenaltyThreshold(&_Contract.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "owner")
//
//	if err != nil {
//		return *new(common.Address), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
//
//	return out0, err
//
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractSession) Owner() (common.Address, error) {
//	return _Contract.Contract.Owner(&_Contract.CallOpts)
//}
//
//// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
////
//// Solidity: function owner() view returns(address)
//func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
//	return _Contract.Contract.Owner(&_Contract.CallOpts)
//}
//
//// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
////
//// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCaller) PendingRewards(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "pendingRewards", delegator, toValidatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
////
//// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
////
//// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.PendingRewards(&_Contract.CallOpts, delegator, toValidatorID)
//}
//
//// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
////
//// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCaller) RewardsStash(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "rewardsStash", delegator, validatorID)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
////
//// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
//}
//
//// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
////
//// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
//func (_Contract *ContractCallerSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
//	return _Contract.Contract.RewardsStash(&_Contract.CallOpts, delegator, validatorID)
//}
//
//// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
////
//// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
//func (_Contract *ContractCaller) SlashingRefundRatio(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "slashingRefundRatio", arg0)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
////
//// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
//func (_Contract *ContractSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
//}
//
//// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
////
//// Solidity: function slashingRefundRatio(uint256 ) view returns(uint256)
//func (_Contract *ContractCallerSession) SlashingRefundRatio(arg0 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.SlashingRefundRatio(&_Contract.CallOpts, arg0)
//}
//
//// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
////
//// Solidity: function stakeTokenizerAddress() view returns(address)
//func (_Contract *ContractCaller) StakeTokenizerAddress(opts *bind.CallOpts) (common.Address, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "stakeTokenizerAddress")
//
//	if err != nil {
//		return *new(common.Address), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
//
//	return out0, err
//
//}
//
//// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
////
//// Solidity: function stakeTokenizerAddress() view returns(address)
//func (_Contract *ContractSession) StakeTokenizerAddress() (common.Address, error) {
//	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
//}
//
//// StakeTokenizerAddress is a free data retrieval call binding the contract method 0x0e559d82.
////
//// Solidity: function stakeTokenizerAddress() view returns(address)
//func (_Contract *ContractCallerSession) StakeTokenizerAddress() (common.Address, error) {
//	return _Contract.Contract.StakeTokenizerAddress(&_Contract.CallOpts)
//}
//
//// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
////
//// Solidity: function stakes(uint256 ) view returns(address delegator, uint96 timestamp, uint256 validatorId)
//func (_Contract *ContractCaller) Stakes(opts *bind.CallOpts, arg0 *big.Int) (struct {
//	Delegator   common.Address
//	Timestamp   *big.Int
//	ValidatorId *big.Int
//}, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "stakes", arg0)
//
//	outstruct := new(struct {
//		Delegator   common.Address
//		Timestamp   *big.Int
//		ValidatorId *big.Int
//	})
//	if err != nil {
//		return *outstruct, err
//	}
//
//	outstruct.Delegator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
//	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
//	outstruct.ValidatorId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
//
//	return *outstruct, err
//
//}
//
//// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
////
//// Solidity: function stakes(uint256 ) view returns(address delegator, uint96 timestamp, uint256 validatorId)
//func (_Contract *ContractSession) Stakes(arg0 *big.Int) (struct {
//	Delegator   common.Address
//	Timestamp   *big.Int
//	ValidatorId *big.Int
//}, error) {
//	return _Contract.Contract.Stakes(&_Contract.CallOpts, arg0)
//}
//
//// Stakes is a free data retrieval call binding the contract method 0xd5a44f86.
////
//// Solidity: function stakes(uint256 ) view returns(address delegator, uint96 timestamp, uint256 validatorId)
//func (_Contract *ContractCallerSession) Stakes(arg0 *big.Int) (struct {
//	Delegator   common.Address
//	Timestamp   *big.Int
//	ValidatorId *big.Int
//}, error) {
//	return _Contract.Contract.Stakes(&_Contract.CallOpts, arg0)
//}
//
//// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
////
//// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
//func (_Contract *ContractCaller) StashedRewardsUntilEpoch(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "stashedRewardsUntilEpoch", arg0, arg1)
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
////
//// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
//func (_Contract *ContractSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
//}
//
//// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
////
//// Solidity: function stashedRewardsUntilEpoch(address , uint256 ) view returns(uint256)
//func (_Contract *ContractCallerSession) StashedRewardsUntilEpoch(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
//	return _Contract.Contract.StashedRewardsUntilEpoch(&_Contract.CallOpts, arg0, arg1)
//}
//
//// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
////
//// Solidity: function totalActiveStake() view returns(uint256)
//func (_Contract *ContractCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "totalActiveStake")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
////
//// Solidity: function totalActiveStake() view returns(uint256)
//func (_Contract *ContractSession) TotalActiveStake() (*big.Int, error) {
//	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
//}
//
//// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
////
//// Solidity: function totalActiveStake() view returns(uint256)
//func (_Contract *ContractCallerSession) TotalActiveStake() (*big.Int, error) {
//	return _Contract.Contract.TotalActiveStake(&_Contract.CallOpts)
//}
//
//// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
////
//// Solidity: function totalSlashedStake() view returns(uint256)
//func (_Contract *ContractCaller) TotalSlashedStake(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "totalSlashedStake")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
////
//// Solidity: function totalSlashedStake() view returns(uint256)
//func (_Contract *ContractSession) TotalSlashedStake() (*big.Int, error) {
//	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
//}
//
//// TotalSlashedStake is a free data retrieval call binding the contract method 0x5fab23a8.
////
//// Solidity: function totalSlashedStake() view returns(uint256)
//func (_Contract *ContractCallerSession) TotalSlashedStake() (*big.Int, error) {
//	return _Contract.Contract.TotalSlashedStake(&_Contract.CallOpts)
//}
//
//// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
////
//// Solidity: function totalStake() view returns(uint256)
//func (_Contract *ContractCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "totalStake")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
////
//// Solidity: function totalStake() view returns(uint256)
//func (_Contract *ContractSession) TotalStake() (*big.Int, error) {
//	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
//}
//
//// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
////
//// Solidity: function totalStake() view returns(uint256)
//func (_Contract *ContractCallerSession) TotalStake() (*big.Int, error) {
//	return _Contract.Contract.TotalStake(&_Contract.CallOpts)
//}
//
//// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
////
//// Solidity: function totalSupply() view returns(uint256)
//func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "totalSupply")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
////
//// Solidity: function totalSupply() view returns(uint256)
//func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
//	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
//}
//
//// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
////
//// Solidity: function totalSupply() view returns(uint256)
//func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
//	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
//}
//
//// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
////
//// Solidity: function unlockedRewardRatio() pure returns(uint256)
//func (_Contract *ContractCaller) UnlockedRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "unlockedRewardRatio")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
////
//// Solidity: function unlockedRewardRatio() pure returns(uint256)
//func (_Contract *ContractSession) UnlockedRewardRatio() (*big.Int, error) {
//	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
//}
//
//// UnlockedRewardRatio is a free data retrieval call binding the contract method 0x5e2308d2.
////
//// Solidity: function unlockedRewardRatio() pure returns(uint256)
//func (_Contract *ContractCallerSession) UnlockedRewardRatio() (*big.Int, error) {
//	return _Contract.Contract.UnlockedRewardRatio(&_Contract.CallOpts)
//}
//
//// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
////
//// Solidity: function validatorCommission() pure returns(uint256)
//func (_Contract *ContractCaller) ValidatorCommission(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "validatorCommission")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
////
//// Solidity: function validatorCommission() pure returns(uint256)
//func (_Contract *ContractSession) ValidatorCommission() (*big.Int, error) {
//	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
//}
//
//// ValidatorCommission is a free data retrieval call binding the contract method 0xa7786515.
////
//// Solidity: function validatorCommission() pure returns(uint256)
//func (_Contract *ContractCallerSession) ValidatorCommission() (*big.Int, error) {
//	return _Contract.Contract.ValidatorCommission(&_Contract.CallOpts)
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() pure returns(bytes3)
//func (_Contract *ContractCaller) Version(opts *bind.CallOpts) ([3]byte, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "version")
//
//	if err != nil {
//		return *new([3]byte), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new([3]byte)).(*[3]byte)
//
//	return out0, err
//
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() pure returns(bytes3)
//func (_Contract *ContractSession) Version() ([3]byte, error) {
//	return _Contract.Contract.Version(&_Contract.CallOpts)
//}
//
//// Version is a free data retrieval call binding the contract method 0x54fd4d50.
////
//// Solidity: function version() pure returns(bytes3)
//func (_Contract *ContractCallerSession) Version() ([3]byte, error) {
//	return _Contract.Contract.Version(&_Contract.CallOpts)
//}
//
//// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
////
//// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
//func (_Contract *ContractCaller) WithdrawalPeriodEpochs(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodEpochs")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
////
//// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
//func (_Contract *ContractSession) WithdrawalPeriodEpochs() (*big.Int, error) {
//	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
//}
//
//// WithdrawalPeriodEpochs is a free data retrieval call binding the contract method 0x650acd66.
////
//// Solidity: function withdrawalPeriodEpochs() pure returns(uint256)
//func (_Contract *ContractCallerSession) WithdrawalPeriodEpochs() (*big.Int, error) {
//	return _Contract.Contract.WithdrawalPeriodEpochs(&_Contract.CallOpts)
//}
//
//// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
////
//// Solidity: function withdrawalPeriodTime() pure returns(uint256)
//func (_Contract *ContractCaller) WithdrawalPeriodTime(opts *bind.CallOpts) (*big.Int, error) {
//	var out []interface{}
//	err := _Contract.contract.Call(opts, &out, "withdrawalPeriodTime")
//
//	if err != nil {
//		return *new(*big.Int), err
//	}
//
//	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
//
//	return out0, err
//
//}
//
//// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
////
//// Solidity: function withdrawalPeriodTime() pure returns(uint256)
//func (_Contract *ContractSession) WithdrawalPeriodTime() (*big.Int, error) {
//	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
//}
//
//// WithdrawalPeriodTime is a free data retrieval call binding the contract method 0xb82b8427.
////
//// Solidity: function withdrawalPeriodTime() pure returns(uint256)
//func (_Contract *ContractCallerSession) WithdrawalPeriodTime() (*big.Int, error) {
//	return _Contract.Contract.WithdrawalPeriodTime(&_Contract.CallOpts)
//}
//
//// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
////
//// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
//func (_Contract *ContractTransactor) SyncValidator(opts *bind.TransactOpts, validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "_syncValidator", validatorID, syncPubkey)
//}
//
//// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
////
//// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
//func (_Contract *ContractSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
//	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
//}
//
//// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
////
//// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
//func (_Contract *ContractTransactorSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
//	return _Contract.Contract.SyncValidator(&_Contract.TransactOpts, validatorID, syncPubkey)
//}
//
//// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
////
//// Solidity: function claimRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "claimRewards", toValidatorID)
//}
//
//// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
////
//// Solidity: function claimRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
//}
//
//// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
////
//// Solidity: function claimRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractTransactorSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, toValidatorID)
//}
//
//// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
////
//// Solidity: function createValidator(bytes pubkey) payable returns()
//func (_Contract *ContractTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "createValidator", pubkey)
//}
//
//// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
////
//// Solidity: function createValidator(bytes pubkey) payable returns()
//func (_Contract *ContractSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
//}
//
//// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
////
//// Solidity: function createValidator(bytes pubkey) payable returns()
//func (_Contract *ContractTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.CreateValidator(&_Contract.TransactOpts, pubkey)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
//}
//
//// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
////
//// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
//func (_Contract *ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.DeactivateValidator(&_Contract.TransactOpts, validatorID, status)
//}
//
//// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
////
//// Solidity: function delegate(uint256 toValidatorID) payable returns()
//func (_Contract *ContractTransactor) Delegate(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "delegate", toValidatorID)
//}
//
//// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
////
//// Solidity: function delegate(uint256 toValidatorID) payable returns()
//func (_Contract *ContractSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
//}
//
//// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
////
//// Solidity: function delegate(uint256 toValidatorID) payable returns()
//func (_Contract *ContractTransactorSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Delegate(&_Contract.TransactOpts, toValidatorID)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
////
//// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
//func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "initialize", sealedEpoch, _totalSupply, nodeDriver, owner)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
////
//// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
//func (_Contract *ContractSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
////
//// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address owner) returns()
//func (_Contract *ContractTransactorSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, owner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, owner)
//}
//
//// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
////
//// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractTransactor) LockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "lockStake", toValidatorID, lockupDuration, amount)
//}
//
//// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
////
//// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
//}
//
//// LockStake is a paid mutator transaction binding the contract method 0xde67f215.
////
//// Solidity: function lockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractTransactorSession) LockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.LockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
//}
//
//// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
////
//// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
//func (_Contract *ContractTransactor) MintFTM(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "mintFTM", receiver, amount, justification)
//}
//
//// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
////
//// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
//func (_Contract *ContractSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
//	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
//}
//
//// MintFTM is a paid mutator transaction binding the contract method 0xe2f8c336.
////
//// Solidity: function mintFTM(address receiver, uint256 amount, string justification) returns()
//func (_Contract *ContractTransactorSession) MintFTM(receiver common.Address, amount *big.Int, justification string) (*types.Transaction, error) {
//	return _Contract.Contract.MintFTM(&_Contract.TransactOpts, receiver, amount, justification)
//}
//
//// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
////
//// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractTransactor) RelockStake(opts *bind.TransactOpts, toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "relockStake", toValidatorID, lockupDuration, amount)
//}
//
//// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
////
//// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
//}
//
//// RelockStake is a paid mutator transaction binding the contract method 0xbd14d907.
////
//// Solidity: function relockStake(uint256 toValidatorID, uint256 lockupDuration, uint256 amount) returns()
//func (_Contract *ContractTransactorSession) RelockStake(toValidatorID *big.Int, lockupDuration *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.RelockStake(&_Contract.TransactOpts, toValidatorID, lockupDuration, amount)
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "renounceOwnership")
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
//	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
//}
//
//// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
////
//// Solidity: function renounceOwnership() returns()
//func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
//	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
//}
//
//// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
////
//// Solidity: function restakeRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractTransactor) RestakeRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "restakeRewards", toValidatorID)
//}
//
//// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
////
//// Solidity: function restakeRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
//}
//
//// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
////
//// Solidity: function restakeRewards(uint256 toValidatorID) returns()
//func (_Contract *ContractTransactorSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.RestakeRewards(&_Contract.TransactOpts, toValidatorID)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "sealEpoch", offlineTime, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactorSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
//}
//
//// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
////
//// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
//func (_Contract *ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpochValidators(&_Contract.TransactOpts, nextValidatorIDs)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisDelegation is a paid mutator transaction binding the contract method 0x18f628d4.
////
//// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake, uint256 lockedStake, uint256 lockupFromEpoch, uint256 lockupEndTime, uint256 lockupDuration, uint256 earlyUnlockPenalty, uint256 rewards) returns()
//func (_Contract *ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int, lockedStake *big.Int, lockupFromEpoch *big.Int, lockupEndTime *big.Int, lockupDuration *big.Int, earlyUnlockPenalty *big.Int, rewards *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisDelegation(&_Contract.TransactOpts, delegator, toValidatorID, stake, lockedStake, lockupFromEpoch, lockupEndTime, lockupDuration, earlyUnlockPenalty, rewards)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setGenesisValidator", auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactorSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
////
//// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
//func (_Contract *ContractTransactor) StashRewards(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "stashRewards", delegator, toValidatorID)
//}
//
//// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
////
//// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
//func (_Contract *ContractSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
//}
//
//// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
////
//// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
//func (_Contract *ContractTransactorSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.StashRewards(&_Contract.TransactOpts, delegator, toValidatorID)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
//}
//
//// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
////
//// Solidity: function transferOwnership(address newOwner) returns()
//func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
//}
//
//// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
////
//// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
//func (_Contract *ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "undelegate", toValidatorID, amount)
//}
//
//// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
////
//// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
//func (_Contract *ContractSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
//}
//
//// Undelegate is a paid mutator transaction binding the contract method 0x634b91e3.
////
//// Solidity: function undelegate(uint256 toValidatorID, uint256 amount) returns()
//func (_Contract *ContractTransactorSession) Undelegate(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Undelegate(&_Contract.TransactOpts, toValidatorID, amount)
//}
//
//// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
////
//// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
//func (_Contract *ContractTransactor) UnlockStake(opts *bind.TransactOpts, toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "unlockStake", toValidatorID, amount)
//}
//
//// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
////
//// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
//func (_Contract *ContractSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
//}
//
//// UnlockStake is a paid mutator transaction binding the contract method 0x1d3ac42c.
////
//// Solidity: function unlockStake(uint256 toValidatorID, uint256 amount) returns(uint256)
//func (_Contract *ContractTransactorSession) UnlockStake(toValidatorID *big.Int, amount *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UnlockStake(&_Contract.TransactOpts, toValidatorID, amount)
//}
//
//// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
////
//// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
//func (_Contract *ContractTransactor) UpdateBaseRewardPerSecond(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateBaseRewardPerSecond", value)
//}
//
//// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
////
//// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
//func (_Contract *ContractSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
//}
//
//// UpdateBaseRewardPerSecond is a paid mutator transaction binding the contract method 0xb6d9edd5.
////
//// Solidity: function updateBaseRewardPerSecond(uint256 value) returns()
//func (_Contract *ContractTransactorSession) UpdateBaseRewardPerSecond(value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateBaseRewardPerSecond(&_Contract.TransactOpts, value)
//}
//
//// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
////
//// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
//func (_Contract *ContractTransactor) UpdateOfflinePenaltyThreshold(opts *bind.TransactOpts, blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateOfflinePenaltyThreshold", blocksNum, time)
//}
//
//// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
////
//// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
//func (_Contract *ContractSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
//}
//
//// UpdateOfflinePenaltyThreshold is a paid mutator transaction binding the contract method 0x8b1a0d11.
////
//// Solidity: function updateOfflinePenaltyThreshold(uint256 blocksNum, uint256 time) returns()
//func (_Contract *ContractTransactorSession) UpdateOfflinePenaltyThreshold(blocksNum *big.Int, time *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateOfflinePenaltyThreshold(&_Contract.TransactOpts, blocksNum, time)
//}
//
//// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
////
//// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
//func (_Contract *ContractTransactor) UpdateSlashingRefundRatio(opts *bind.TransactOpts, validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateSlashingRefundRatio", validatorID, refundRatio)
//}
//
//// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
////
//// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
//func (_Contract *ContractSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
//}
//
//// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
////
//// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
//func (_Contract *ContractTransactorSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateSlashingRefundRatio(&_Contract.TransactOpts, validatorID, refundRatio)
//}
//
//// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
////
//// Solidity: function updateStakeTokenizerAddress(address addr) returns()
//func (_Contract *ContractTransactor) UpdateStakeTokenizerAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateStakeTokenizerAddress", addr)
//}
//
//// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
////
//// Solidity: function updateStakeTokenizerAddress(address addr) returns()
//func (_Contract *ContractSession) UpdateStakeTokenizerAddress(addr common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateStakeTokenizerAddress(&_Contract.TransactOpts, addr)
//}
//
//// UpdateStakeTokenizerAddress is a paid mutator transaction binding the contract method 0xa2f6e6bc.
////
//// Solidity: function updateStakeTokenizerAddress(address addr) returns()
//func (_Contract *ContractTransactorSession) UpdateStakeTokenizerAddress(addr common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateStakeTokenizerAddress(&_Contract.TransactOpts, addr)
//}
//
//// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
////
//// Solidity: function updateTotalSupply(int256 diff) returns()
//func (_Contract *ContractTransactor) UpdateTotalSupply(opts *bind.TransactOpts, diff *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateTotalSupply", diff)
//}
//
//// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
////
//// Solidity: function updateTotalSupply(int256 diff) returns()
//func (_Contract *ContractSession) UpdateTotalSupply(diff *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateTotalSupply(&_Contract.TransactOpts, diff)
//}
//
//// UpdateTotalSupply is a paid mutator transaction binding the contract method 0x346bdcfb.
////
//// Solidity: function updateTotalSupply(int256 diff) returns()
//func (_Contract *ContractTransactorSession) UpdateTotalSupply(diff *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateTotalSupply(&_Contract.TransactOpts, diff)
//}
//
//// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
////
//// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
//func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "withdraw", toValidatorID, wrID)
//}
//
//// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
////
//// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
//func (_Contract *ContractSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
//}
//
//// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
////
//// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
//func (_Contract *ContractTransactorSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, toValidatorID, wrID)
//}
//
//// ContractChangedValidatorStatusIterator is returned from FilterChangedValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangedValidatorStatus events raised by the Contract contract.
//type ContractChangedValidatorStatusIterator struct {
//	Event *ContractChangedValidatorStatus // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractChangedValidatorStatusIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractChangedValidatorStatus)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractChangedValidatorStatus)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractChangedValidatorStatusIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractChangedValidatorStatusIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractChangedValidatorStatus represents a ChangedValidatorStatus event raised by the Contract contract.
//type ContractChangedValidatorStatus struct {
//	ValidatorID *big.Int
//	Status      *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterChangedValidatorStatus is a free log retrieval operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
////
//// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
//func (_Contract *ContractFilterer) FilterChangedValidatorStatus(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractChangedValidatorStatusIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChangedValidatorStatus", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractChangedValidatorStatusIterator{contract: _Contract.contract, event: "ChangedValidatorStatus", logs: logs, sub: sub}, nil
//}
//
//// WatchChangedValidatorStatus is a free log subscription operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
////
//// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
//func (_Contract *ContractFilterer) WatchChangedValidatorStatus(opts *bind.WatchOpts, sink chan<- *ContractChangedValidatorStatus, validatorID []*big.Int) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChangedValidatorStatus", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractChangedValidatorStatus)
//				if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseChangedValidatorStatus is a log parse operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
////
//// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
//func (_Contract *ContractFilterer) ParseChangedValidatorStatus(log types.Log) (*ContractChangedValidatorStatus, error) {
//	event := new(ContractChangedValidatorStatus)
//	if err := _Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the Contract contract.
//type ContractClaimedRewardsIterator struct {
//	Event *ContractClaimedRewards // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractClaimedRewardsIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractClaimedRewards)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractClaimedRewards)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractClaimedRewardsIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractClaimedRewardsIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractClaimedRewards represents a ClaimedRewards event raised by the Contract contract.
//type ContractClaimedRewards struct {
//	Delegator         common.Address
//	ToValidatorID     *big.Int
//	LockupExtraReward *big.Int
//	LockupBaseReward  *big.Int
//	UnlockedReward    *big.Int
//	Raw               types.Log // Blockchain specific contextual infos
//}
//
//// FilterClaimedRewards is a free log retrieval operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
////
//// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) FilterClaimedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractClaimedRewardsIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractClaimedRewardsIterator{contract: _Contract.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
//}
//
//// WatchClaimedRewards is a free log subscription operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
////
//// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *ContractClaimedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractClaimedRewards)
//				if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseClaimedRewards is a log parse operation binding the contract event 0xc1d8eb6e444b89fb8ff0991c19311c070df704ccb009e210d1462d5b2410bf45.
////
//// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) ParseClaimedRewards(log types.Log) (*ContractClaimedRewards, error) {
//	event := new(ContractClaimedRewards)
//	if err := _Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractCreatedValidatorIterator is returned from FilterCreatedValidator and is used to iterate over the raw logs and unpacked data for CreatedValidator events raised by the Contract contract.
//type ContractCreatedValidatorIterator struct {
//	Event *ContractCreatedValidator // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractCreatedValidatorIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractCreatedValidator)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractCreatedValidator)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractCreatedValidatorIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractCreatedValidatorIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractCreatedValidator represents a CreatedValidator event raised by the Contract contract.
//type ContractCreatedValidator struct {
//	ValidatorID  *big.Int
//	Auth         common.Address
//	CreatedEpoch *big.Int
//	CreatedTime  *big.Int
//	Raw          types.Log // Blockchain specific contextual infos
//}
//
//// FilterCreatedValidator is a free log retrieval operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
////
//// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
//func (_Contract *ContractFilterer) FilterCreatedValidator(opts *bind.FilterOpts, validatorID []*big.Int, auth []common.Address) (*ContractCreatedValidatorIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//	var authRule []interface{}
//	for _, authItem := range auth {
//		authRule = append(authRule, authItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "CreatedValidator", validatorIDRule, authRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractCreatedValidatorIterator{contract: _Contract.contract, event: "CreatedValidator", logs: logs, sub: sub}, nil
//}
//
//// WatchCreatedValidator is a free log subscription operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
////
//// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
//func (_Contract *ContractFilterer) WatchCreatedValidator(opts *bind.WatchOpts, sink chan<- *ContractCreatedValidator, validatorID []*big.Int, auth []common.Address) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//	var authRule []interface{}
//	for _, authItem := range auth {
//		authRule = append(authRule, authItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "CreatedValidator", validatorIDRule, authRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractCreatedValidator)
//				if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseCreatedValidator is a log parse operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
////
//// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
//func (_Contract *ContractFilterer) ParseCreatedValidator(log types.Log) (*ContractCreatedValidator, error) {
//	event := new(ContractCreatedValidator)
//	if err := _Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractDeactivatedValidatorIterator is returned from FilterDeactivatedValidator and is used to iterate over the raw logs and unpacked data for DeactivatedValidator events raised by the Contract contract.
//type ContractDeactivatedValidatorIterator struct {
//	Event *ContractDeactivatedValidator // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractDeactivatedValidatorIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractDeactivatedValidator)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractDeactivatedValidator)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractDeactivatedValidatorIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractDeactivatedValidatorIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractDeactivatedValidator represents a DeactivatedValidator event raised by the Contract contract.
//type ContractDeactivatedValidator struct {
//	ValidatorID      *big.Int
//	DeactivatedEpoch *big.Int
//	DeactivatedTime  *big.Int
//	Raw              types.Log // Blockchain specific contextual infos
//}
//
//// FilterDeactivatedValidator is a free log retrieval operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
////
//// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
//func (_Contract *ContractFilterer) FilterDeactivatedValidator(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractDeactivatedValidatorIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "DeactivatedValidator", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractDeactivatedValidatorIterator{contract: _Contract.contract, event: "DeactivatedValidator", logs: logs, sub: sub}, nil
//}
//
//// WatchDeactivatedValidator is a free log subscription operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
////
//// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
//func (_Contract *ContractFilterer) WatchDeactivatedValidator(opts *bind.WatchOpts, sink chan<- *ContractDeactivatedValidator, validatorID []*big.Int) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "DeactivatedValidator", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractDeactivatedValidator)
//				if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseDeactivatedValidator is a log parse operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
////
//// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
//func (_Contract *ContractFilterer) ParseDeactivatedValidator(log types.Log) (*ContractDeactivatedValidator, error) {
//	event := new(ContractDeactivatedValidator)
//	if err := _Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Contract contract.
//type ContractDelegatedIterator struct {
//	Event *ContractDelegated // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractDelegatedIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractDelegated)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractDelegated)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractDelegatedIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractDelegatedIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractDelegated represents a Delegated event raised by the Contract contract.
//type ContractDelegated struct {
//	Delegator     common.Address
//	ToValidatorID *big.Int
//	Amount        *big.Int
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
////
//// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
//func (_Contract *ContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractDelegatedIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractDelegatedIterator{contract: _Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
//}
//
//// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
////
//// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
//func (_Contract *ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegated, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractDelegated)
//				if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
////
//// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
//func (_Contract *ContractFilterer) ParseDelegated(log types.Log) (*ContractDelegated, error) {
//	event := new(ContractDelegated)
//	if err := _Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractInflatedFTMIterator is returned from FilterInflatedFTM and is used to iterate over the raw logs and unpacked data for InflatedFTM events raised by the Contract contract.
//type ContractInflatedFTMIterator struct {
//	Event *ContractInflatedFTM // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractInflatedFTMIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractInflatedFTM)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractInflatedFTM)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractInflatedFTMIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractInflatedFTMIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractInflatedFTM represents a InflatedFTM event raised by the Contract contract.
//type ContractInflatedFTM struct {
//	Receiver      common.Address
//	Amount        *big.Int
//	Justification string
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterInflatedFTM is a free log retrieval operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
////
//// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
//func (_Contract *ContractFilterer) FilterInflatedFTM(opts *bind.FilterOpts, receiver []common.Address) (*ContractInflatedFTMIterator, error) {
//
//	var receiverRule []interface{}
//	for _, receiverItem := range receiver {
//		receiverRule = append(receiverRule, receiverItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "InflatedFTM", receiverRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractInflatedFTMIterator{contract: _Contract.contract, event: "InflatedFTM", logs: logs, sub: sub}, nil
//}
//
//// WatchInflatedFTM is a free log subscription operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
////
//// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
//func (_Contract *ContractFilterer) WatchInflatedFTM(opts *bind.WatchOpts, sink chan<- *ContractInflatedFTM, receiver []common.Address) (event.Subscription, error) {
//
//	var receiverRule []interface{}
//	for _, receiverItem := range receiver {
//		receiverRule = append(receiverRule, receiverItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "InflatedFTM", receiverRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractInflatedFTM)
//				if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseInflatedFTM is a log parse operation binding the contract event 0x9eec469b348bcf64bbfb60e46ce7b160e2e09bf5421496a2cdbc43714c28b8ad.
////
//// Solidity: event InflatedFTM(address indexed receiver, uint256 amount, string justification)
//func (_Contract *ContractFilterer) ParseInflatedFTM(log types.Log) (*ContractInflatedFTM, error) {
//	event := new(ContractInflatedFTM)
//	if err := _Contract.contract.UnpackLog(event, "InflatedFTM", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractLockedUpStakeIterator is returned from FilterLockedUpStake and is used to iterate over the raw logs and unpacked data for LockedUpStake events raised by the Contract contract.
//type ContractLockedUpStakeIterator struct {
//	Event *ContractLockedUpStake // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractLockedUpStakeIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractLockedUpStake)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractLockedUpStake)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractLockedUpStakeIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractLockedUpStakeIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractLockedUpStake represents a LockedUpStake event raised by the Contract contract.
//type ContractLockedUpStake struct {
//	Delegator   common.Address
//	ValidatorID *big.Int
//	Duration    *big.Int
//	Amount      *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterLockedUpStake is a free log retrieval operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
////
//// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
//func (_Contract *ContractFilterer) FilterLockedUpStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractLockedUpStakeIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractLockedUpStakeIterator{contract: _Contract.contract, event: "LockedUpStake", logs: logs, sub: sub}, nil
//}
//
//// WatchLockedUpStake is a free log subscription operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
////
//// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
//func (_Contract *ContractFilterer) WatchLockedUpStake(opts *bind.WatchOpts, sink chan<- *ContractLockedUpStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "LockedUpStake", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractLockedUpStake)
//				if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseLockedUpStake is a log parse operation binding the contract event 0x138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78.
////
//// Solidity: event LockedUpStake(address indexed delegator, uint256 indexed validatorID, uint256 duration, uint256 amount)
//func (_Contract *ContractFilterer) ParseLockedUpStake(log types.Log) (*ContractLockedUpStake, error) {
//	event := new(ContractLockedUpStake)
//	if err := _Contract.contract.UnpackLog(event, "LockedUpStake", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
//type ContractOwnershipTransferredIterator struct {
//	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractOwnershipTransferredIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractOwnershipTransferred)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractOwnershipTransferred)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractOwnershipTransferredIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractOwnershipTransferredIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
//type ContractOwnershipTransferred struct {
//	PreviousOwner common.Address
//	NewOwner      common.Address
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {
//
//	var previousOwnerRule []interface{}
//	for _, previousOwnerItem := range previousOwner {
//		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
//	}
//	var newOwnerRule []interface{}
//	for _, newOwnerItem := range newOwner {
//		newOwnerRule = append(newOwnerRule, newOwnerItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
//}
//
//// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
//
//	var previousOwnerRule []interface{}
//	for _, previousOwnerItem := range previousOwner {
//		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
//	}
//	var newOwnerRule []interface{}
//	for _, newOwnerItem := range newOwner {
//		newOwnerRule = append(newOwnerRule, newOwnerItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractOwnershipTransferred)
//				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
////
//// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
//func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
//	event := new(ContractOwnershipTransferred)
//	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractRefundedSlashedLegacyDelegationIterator is returned from FilterRefundedSlashedLegacyDelegation and is used to iterate over the raw logs and unpacked data for RefundedSlashedLegacyDelegation events raised by the Contract contract.
//type ContractRefundedSlashedLegacyDelegationIterator struct {
//	Event *ContractRefundedSlashedLegacyDelegation // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractRefundedSlashedLegacyDelegationIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractRefundedSlashedLegacyDelegation)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractRefundedSlashedLegacyDelegation)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractRefundedSlashedLegacyDelegationIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractRefundedSlashedLegacyDelegationIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractRefundedSlashedLegacyDelegation represents a RefundedSlashedLegacyDelegation event raised by the Contract contract.
//type ContractRefundedSlashedLegacyDelegation struct {
//	Delegator   common.Address
//	ValidatorID *big.Int
//	Amount      *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterRefundedSlashedLegacyDelegation is a free log retrieval operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
////
//// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
//func (_Contract *ContractFilterer) FilterRefundedSlashedLegacyDelegation(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractRefundedSlashedLegacyDelegationIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractRefundedSlashedLegacyDelegationIterator{contract: _Contract.contract, event: "RefundedSlashedLegacyDelegation", logs: logs, sub: sub}, nil
//}
//
//// WatchRefundedSlashedLegacyDelegation is a free log subscription operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
////
//// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
//func (_Contract *ContractFilterer) WatchRefundedSlashedLegacyDelegation(opts *bind.WatchOpts, sink chan<- *ContractRefundedSlashedLegacyDelegation, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractRefundedSlashedLegacyDelegation)
//				if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseRefundedSlashedLegacyDelegation is a log parse operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
////
//// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
//func (_Contract *ContractFilterer) ParseRefundedSlashedLegacyDelegation(log types.Log) (*ContractRefundedSlashedLegacyDelegation, error) {
//	event := new(ContractRefundedSlashedLegacyDelegation)
//	if err := _Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractRestakedRewardsIterator is returned from FilterRestakedRewards and is used to iterate over the raw logs and unpacked data for RestakedRewards events raised by the Contract contract.
//type ContractRestakedRewardsIterator struct {
//	Event *ContractRestakedRewards // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractRestakedRewardsIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractRestakedRewards)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractRestakedRewards)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractRestakedRewardsIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractRestakedRewardsIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractRestakedRewards represents a RestakedRewards event raised by the Contract contract.
//type ContractRestakedRewards struct {
//	Delegator         common.Address
//	ToValidatorID     *big.Int
//	LockupExtraReward *big.Int
//	LockupBaseReward  *big.Int
//	UnlockedReward    *big.Int
//	Raw               types.Log // Blockchain specific contextual infos
//}
//
//// FilterRestakedRewards is a free log retrieval operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
////
//// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) FilterRestakedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*ContractRestakedRewardsIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractRestakedRewardsIterator{contract: _Contract.contract, event: "RestakedRewards", logs: logs, sub: sub}, nil
//}
//
//// WatchRestakedRewards is a free log subscription operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
////
//// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) WatchRestakedRewards(opts *bind.WatchOpts, sink chan<- *ContractRestakedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractRestakedRewards)
//				if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseRestakedRewards is a log parse operation binding the contract event 0x4119153d17a36f9597d40e3ab4148d03261a439dddbec4e91799ab7159608e26.
////
//// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 lockupExtraReward, uint256 lockupBaseReward, uint256 unlockedReward)
//func (_Contract *ContractFilterer) ParseRestakedRewards(log types.Log) (*ContractRestakedRewards, error) {
//	event := new(ContractRestakedRewards)
//	if err := _Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Contract contract.
//type ContractUndelegatedIterator struct {
//	Event *ContractUndelegated // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractUndelegatedIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUndelegated)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractUndelegated)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractUndelegatedIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUndelegatedIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUndelegated represents a Undelegated event raised by the Contract contract.
//type ContractUndelegated struct {
//	Delegator     common.Address
//	ToValidatorID *big.Int
//	WrID          *big.Int
//	Amount        *big.Int
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterUndelegated is a free log retrieval operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
////
//// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractUndelegatedIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//	var wrIDRule []interface{}
//	for _, wrIDItem := range wrID {
//		wrIDRule = append(wrIDRule, wrIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUndelegatedIterator{contract: _Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
//}
//
//// WatchUndelegated is a free log subscription operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
////
//// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *ContractUndelegated, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//	var wrIDRule []interface{}
//	for _, wrIDItem := range wrID {
//		wrIDRule = append(wrIDRule, wrIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUndelegated)
//				if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseUndelegated is a log parse operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
////
//// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) ParseUndelegated(log types.Log) (*ContractUndelegated, error) {
//	event := new(ContractUndelegated)
//	if err := _Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUnlockedStakeIterator is returned from FilterUnlockedStake and is used to iterate over the raw logs and unpacked data for UnlockedStake events raised by the Contract contract.
//type ContractUnlockedStakeIterator struct {
//	Event *ContractUnlockedStake // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractUnlockedStakeIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUnlockedStake)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractUnlockedStake)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractUnlockedStakeIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUnlockedStakeIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUnlockedStake represents a UnlockedStake event raised by the Contract contract.
//type ContractUnlockedStake struct {
//	Delegator   common.Address
//	ValidatorID *big.Int
//	Amount      *big.Int
//	Penalty     *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterUnlockedStake is a free log retrieval operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
////
//// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
//func (_Contract *ContractFilterer) FilterUnlockedStake(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*ContractUnlockedStakeIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUnlockedStakeIterator{contract: _Contract.contract, event: "UnlockedStake", logs: logs, sub: sub}, nil
//}
//
//// WatchUnlockedStake is a free log subscription operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
////
//// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
//func (_Contract *ContractFilterer) WatchUnlockedStake(opts *bind.WatchOpts, sink chan<- *ContractUnlockedStake, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UnlockedStake", delegatorRule, validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUnlockedStake)
//				if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseUnlockedStake is a log parse operation binding the contract event 0xef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9.
////
//// Solidity: event UnlockedStake(address indexed delegator, uint256 indexed validatorID, uint256 amount, uint256 penalty)
//func (_Contract *ContractFilterer) ParseUnlockedStake(log types.Log) (*ContractUnlockedStake, error) {
//	event := new(ContractUnlockedStake)
//	if err := _Contract.contract.UnpackLog(event, "UnlockedStake", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdatedBaseRewardPerSecIterator is returned from FilterUpdatedBaseRewardPerSec and is used to iterate over the raw logs and unpacked data for UpdatedBaseRewardPerSec events raised by the Contract contract.
//type ContractUpdatedBaseRewardPerSecIterator struct {
//	Event *ContractUpdatedBaseRewardPerSec // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractUpdatedBaseRewardPerSecIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdatedBaseRewardPerSec)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractUpdatedBaseRewardPerSec)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractUpdatedBaseRewardPerSecIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdatedBaseRewardPerSecIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdatedBaseRewardPerSec represents a UpdatedBaseRewardPerSec event raised by the Contract contract.
//type ContractUpdatedBaseRewardPerSec struct {
//	Value *big.Int
//	Raw   types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdatedBaseRewardPerSec is a free log retrieval operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
////
//// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
//func (_Contract *ContractFilterer) FilterUpdatedBaseRewardPerSec(opts *bind.FilterOpts) (*ContractUpdatedBaseRewardPerSecIterator, error) {
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBaseRewardPerSec")
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdatedBaseRewardPerSecIterator{contract: _Contract.contract, event: "UpdatedBaseRewardPerSec", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdatedBaseRewardPerSec is a free log subscription operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
////
//// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
//func (_Contract *ContractFilterer) WatchUpdatedBaseRewardPerSec(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBaseRewardPerSec) (event.Subscription, error) {
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBaseRewardPerSec")
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdatedBaseRewardPerSec)
//				if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseUpdatedBaseRewardPerSec is a log parse operation binding the contract event 0x8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae2396.
////
//// Solidity: event UpdatedBaseRewardPerSec(uint256 value)
//func (_Contract *ContractFilterer) ParseUpdatedBaseRewardPerSec(log types.Log) (*ContractUpdatedBaseRewardPerSec, error) {
//	event := new(ContractUpdatedBaseRewardPerSec)
//	if err := _Contract.contract.UnpackLog(event, "UpdatedBaseRewardPerSec", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdatedOfflinePenaltyThresholdIterator is returned from FilterUpdatedOfflinePenaltyThreshold and is used to iterate over the raw logs and unpacked data for UpdatedOfflinePenaltyThreshold events raised by the Contract contract.
//type ContractUpdatedOfflinePenaltyThresholdIterator struct {
//	Event *ContractUpdatedOfflinePenaltyThreshold // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractUpdatedOfflinePenaltyThreshold)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdatedOfflinePenaltyThresholdIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdatedOfflinePenaltyThreshold represents a UpdatedOfflinePenaltyThreshold event raised by the Contract contract.
//type ContractUpdatedOfflinePenaltyThreshold struct {
//	BlocksNum *big.Int
//	Period    *big.Int
//	Raw       types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdatedOfflinePenaltyThreshold is a free log retrieval operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
////
//// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
//func (_Contract *ContractFilterer) FilterUpdatedOfflinePenaltyThreshold(opts *bind.FilterOpts) (*ContractUpdatedOfflinePenaltyThresholdIterator, error) {
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedOfflinePenaltyThreshold")
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdatedOfflinePenaltyThresholdIterator{contract: _Contract.contract, event: "UpdatedOfflinePenaltyThreshold", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdatedOfflinePenaltyThreshold is a free log subscription operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
////
//// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
//func (_Contract *ContractFilterer) WatchUpdatedOfflinePenaltyThreshold(opts *bind.WatchOpts, sink chan<- *ContractUpdatedOfflinePenaltyThreshold) (event.Subscription, error) {
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedOfflinePenaltyThreshold")
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdatedOfflinePenaltyThreshold)
//				if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseUpdatedOfflinePenaltyThreshold is a log parse operation binding the contract event 0x702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34.
////
//// Solidity: event UpdatedOfflinePenaltyThreshold(uint256 blocksNum, uint256 period)
//func (_Contract *ContractFilterer) ParseUpdatedOfflinePenaltyThreshold(log types.Log) (*ContractUpdatedOfflinePenaltyThreshold, error) {
//	event := new(ContractUpdatedOfflinePenaltyThreshold)
//	if err := _Contract.contract.UnpackLog(event, "UpdatedOfflinePenaltyThreshold", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdatedSlashingRefundRatioIterator is returned from FilterUpdatedSlashingRefundRatio and is used to iterate over the raw logs and unpacked data for UpdatedSlashingRefundRatio events raised by the Contract contract.
//type ContractUpdatedSlashingRefundRatioIterator struct {
//	Event *ContractUpdatedSlashingRefundRatio // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractUpdatedSlashingRefundRatioIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdatedSlashingRefundRatio)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractUpdatedSlashingRefundRatio)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractUpdatedSlashingRefundRatioIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdatedSlashingRefundRatioIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdatedSlashingRefundRatio represents a UpdatedSlashingRefundRatio event raised by the Contract contract.
//type ContractUpdatedSlashingRefundRatio struct {
//	ValidatorID *big.Int
//	RefundRatio *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdatedSlashingRefundRatio is a free log retrieval operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
////
//// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
//func (_Contract *ContractFilterer) FilterUpdatedSlashingRefundRatio(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdatedSlashingRefundRatioIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdatedSlashingRefundRatioIterator{contract: _Contract.contract, event: "UpdatedSlashingRefundRatio", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdatedSlashingRefundRatio is a free log subscription operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
////
//// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
//func (_Contract *ContractFilterer) WatchUpdatedSlashingRefundRatio(opts *bind.WatchOpts, sink chan<- *ContractUpdatedSlashingRefundRatio, validatorID []*big.Int) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdatedSlashingRefundRatio)
//				if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseUpdatedSlashingRefundRatio is a log parse operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
////
//// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
//func (_Contract *ContractFilterer) ParseUpdatedSlashingRefundRatio(log types.Log) (*ContractUpdatedSlashingRefundRatio, error) {
//	event := new(ContractUpdatedSlashingRefundRatio)
//	if err := _Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contract contract.
//type ContractWithdrawnIterator struct {
//	Event *ContractWithdrawn // Event containing the contract specifics and raw log
//
//	contract *bind.BoundContract // Generic contract to use for unpacking event data
//	event    string              // Event name to use for unpacking event data
//
//	logs chan types.Log        // Log channel receiving the found contract events
//	sub  ethereum.Subscription // Subscription for errors, completion and termination
//	done bool                  // Whether the subscription completed delivering logs
//	fail error                 // Occurred error to stop iteration
//}
//
//// Next advances the iterator to the subsequent event, returning whether there
//// are any more events found. In case of a retrieval or parsing error, false is
//// returned and Error() can be queried for the exact failure.
//func (it *ContractWithdrawnIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractWithdrawn)
//			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//				it.fail = err
//				return false
//			}
//			it.Event.Raw = log
//			return true
//
//		default:
//			return false
//		}
//	}
//	// Iterator still in progress, wait for either a data or an error event
//	select {
//	case log := <-it.logs:
//		it.Event = new(ContractWithdrawn)
//		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
//			it.fail = err
//			return false
//		}
//		it.Event.Raw = log
//		return true
//
//	case err := <-it.sub.Err():
//		it.done = true
//		it.fail = err
//		return it.Next()
//	}
//}
//
//// Error returns any retrieval or parsing error occurred during filtering.
//func (it *ContractWithdrawnIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractWithdrawnIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractWithdrawn represents a Withdrawn event raised by the Contract contract.
//type ContractWithdrawn struct {
//	Delegator     common.Address
//	ToValidatorID *big.Int
//	WrID          *big.Int
//	Amount        *big.Int
//	Raw           types.Log // Blockchain specific contextual infos
//}
//
//// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
////
//// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*ContractWithdrawnIterator, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//	var wrIDRule []interface{}
//	for _, wrIDItem := range wrID {
//		wrIDRule = append(wrIDRule, wrIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
//}
//
//// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
////
//// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {
//
//	var delegatorRule []interface{}
//	for _, delegatorItem := range delegator {
//		delegatorRule = append(delegatorRule, delegatorItem)
//	}
//	var toValidatorIDRule []interface{}
//	for _, toValidatorIDItem := range toValidatorID {
//		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
//	}
//	var wrIDRule []interface{}
//	for _, wrIDItem := range wrID {
//		wrIDRule = append(wrIDRule, wrIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractWithdrawn)
//				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
//					return err
//				}
//				event.Raw = log
//
//				select {
//				case sink <- event:
//				case err := <-sub.Err():
//					return err
//				case <-quit:
//					return nil
//				}
//			case err := <-sub.Err():
//				return err
//			case <-quit:
//				return nil
//			}
//		}
//	}), nil
//}
//
//// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
////
//// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
//func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
//	event := new(ContractWithdrawn)
//	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
