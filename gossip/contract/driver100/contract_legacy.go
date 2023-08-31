//// Code generated - DO NOT EDIT.
//// This file is a generated binding and any manual changes will be lost.
//
package driver100

//
//import (
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
//	_ = big.NewInt
//	_ = strings.NewReader
//	_ = ethereum.NotFound
//	_ = bind.Bind
//	_ = common.Big1
//	_ = types.BloomLookup
//	_ = event.NewSubscription
//)
//
//// ContractABI is the input ABI used to generate the binding from.
//const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"AdvanceEpochs\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"UpdateNetworkRules\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"UpdateNetworkVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"UpdateValidatorPubkey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"UpdateValidatorWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backend\",\"type\":\"address\"}],\"name\":\"UpdatedBackend\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"}],\"name\":\"setBackend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_backend\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_evmWriterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"with\",\"type\":\"address\"}],\"name\":\"swapCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"diff\",\"type\":\"bytes\"}],\"name\":\"updateNetworkRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"updateNetworkVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"advanceEpochs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"updateValidatorWeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"updateValidatorPubkey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupFromEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockupDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earlyUnlockPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
//
//// ContractBin is the compiled bytecode used for deploying new contracts.
//var ContractBin = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c80634feb92f3116100a2578063d6a0c7af11610071578063d6a0c7af146101f6578063da7fc24f14610209578063e08d7e661461021c578063e30443bc1461022f578063ebdf104c146102425761010b565b80634feb92f3146101aa57806379bead38146101bd578063a4066fbe146101d0578063b9cc6b1c146101e35761010b565b8063242a6e3f116100de578063242a6e3f1461015e578063267ab4461461017157806339e503ab14610184578063485cc955146101975761010b565b806307690b2a146101105780630aeeca001461012557806318f628d4146101385780631e702f831461014b575b600080fd5b61012361011e366004610af7565b610255565b005b610123610133366004610e78565b6102f0565b610123610146366004610c7a565b610354565b610123610159366004610eec565b6103ef565b61012361016c366004610e96565b61043f565b61012361017f366004610e78565b6104a8565b610123610192366004610b31565b610501565b6101236101a5366004610af7565b610596565b6101236101b8366004610bae565b61067d565b6101236101cb366004610b7e565b6106db565b6101236101de366004610eec565b610737565b6101236101f1366004610e42565b61079d565b610123610204366004610af7565b610804565b610123610217366004610ad1565b610860565b61012361022a366004610d2e565b6108e0565b61012361023d366004610b7e565b610930565b610123610250366004610d70565b61098c565b6034546001600160a01b031633146102885760405162461bcd60e51b815260040161027f9061121b565b60405180910390fd5b6035546040516303b4859560e11b81526001600160a01b03909116906307690b2a906102ba9085908590600401611039565b600060405180830381600087803b1580156102d457600080fd5b505af11580156102e8573d6000803e3d6000fd5b505050505050565b6034546001600160a01b0316331461031a5760405162461bcd60e51b815260040161027f9061121b565b7f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd181604051610349919061123b565b60405180910390a150565b33156103725760405162461bcd60e51b815260040161027f9061122b565b60345460405163063d8a3560e21b81526001600160a01b03909116906318f628d4906103b2908c908c908c908c908c908c908c908c908c9060040161111e565b600060405180830381600087803b1580156103cc57600080fd5b505af11580156103e0573d6000803e3d6000fd5b50505050505050505050505050565b331561040d5760405162461bcd60e51b815260040161027f9061122b565b603454604051631e702f8360e01b81526001600160a01b0390911690631e702f83906102ba9085908590600401611249565b6034546001600160a01b031633146104695760405162461bcd60e51b815260040161027f9061121b565b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405161049b9291906111f9565b60405180910390a2505050565b6034546001600160a01b031633146104d25760405162461bcd60e51b815260040161027f9061121b565b7f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb81604051610349919061123b565b6034546001600160a01b0316331461052b5760405162461bcd60e51b815260040161027f9061121b565b6035546040516339e503ab60e01b81526001600160a01b03909116906339e503ab9061055f9086908690869060040161105b565b600060405180830381600087803b15801561057957600080fd5b505af115801561058d573d6000803e3d6000fd5b50505050505050565b600054610100900460ff16806105af57506105af610a24565b806105bd575060005460ff16155b6105d95760405162461bcd60e51b815260040161027f9061120b565b600054610100900460ff16158015610604576000805460ff1961ff0019909116610100171660011790555b603480546001600160a01b0319166001600160a01b0385169081179091556040517f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603580546001600160a01b0319166001600160a01b0384161790558015610678576000805461ff00191690555b505050565b331561069b5760405162461bcd60e51b815260040161027f9061122b565b603454604051634feb92f360e01b81526001600160a01b0390911690634feb92f3906103b2908c908c908c908c908c908c908c908c908c9060040161109e565b6034546001600160a01b031633146107055760405162461bcd60e51b815260040161027f9061121b565b603554604051630f37d5a760e31b81526001600160a01b03909116906379bead38906102ba9085908590600401611083565b6034546001600160a01b031633146107615760405162461bcd60e51b815260040161027f9061121b565b817fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a93582604051610791919061123b565b60405180910390a25050565b6034546001600160a01b031633146107c75760405162461bcd60e51b815260040161027f9061121b565b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb9982826040516107f89291906111f9565b60405180910390a15050565b6034546001600160a01b0316331461082e5760405162461bcd60e51b815260040161027f9061121b565b60355460405163d6a0c7af60e01b81526001600160a01b039091169063d6a0c7af906102ba9085908590600401611039565b6034546001600160a01b0316331461088a5760405162461bcd60e51b815260040161027f9061121b565b6040516001600160a01b038216907f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603480546001600160a01b0319166001600160a01b0392909216919091179055565b33156108fe5760405162461bcd60e51b815260040161027f9061122b565b603454604051637046bf3360e11b81526001600160a01b039091169063e08d7e66906102ba9085908590600401611196565b6034546001600160a01b0316331461095a5760405162461bcd60e51b815260040161027f9061121b565b6035546040516338c110ef60e21b81526001600160a01b039091169063e30443bc906102ba9085908590600401611083565b33156109aa5760405162461bcd60e51b815260040161027f9061122b565b603454604051633af7c41360e21b81526001600160a01b039091169063ebdf104c906109e8908b908b908b908b908b908b908b908b906004016111a8565b600060405180830381600087803b158015610a0257600080fd5b505af1158015610a16573d6000803e3d6000fd5b505050505050505050505050565b303b1590565b8035610a3581611290565b92915050565b60008083601f840112610a4d57600080fd5b50813567ffffffffffffffff811115610a6557600080fd5b602083019150836020820283011115610a7d57600080fd5b9250929050565b8035610a35816112a7565b60008083601f840112610aa157600080fd5b50813567ffffffffffffffff811115610ab957600080fd5b602083019150836001820283011115610a7d57600080fd5b600060208284031215610ae357600080fd5b6000610aef8484610a2a565b949350505050565b60008060408385031215610b0a57600080fd5b6000610b168585610a2a565b9250506020610b2785828601610a2a565b9150509250929050565b600080600060608486031215610b4657600080fd5b6000610b528686610a2a565b9350506020610b6386828701610a84565b9250506040610b7486828701610a84565b9150509250925092565b60008060408385031215610b9157600080fd5b6000610b9d8585610a2a565b9250506020610b2785828601610a84565b60008060008060008060008060006101008a8c031215610bcd57600080fd5b6000610bd98c8c610a2a565b9950506020610bea8c828d01610a84565b98505060408a013567ffffffffffffffff811115610c0757600080fd5b610c138c828d01610a8f565b97509750506060610c268c828d01610a84565b9550506080610c378c828d01610a84565b94505060a0610c488c828d01610a84565b93505060c0610c598c828d01610a84565b92505060e0610c6a8c828d01610a84565b9150509295985092959850929598565b60008060008060008060008060006101208a8c031215610c9957600080fd5b6000610ca58c8c610a2a565b9950506020610cb68c828d01610a84565b9850506040610cc78c828d01610a84565b9750506060610cd88c828d01610a84565b9650506080610ce98c828d01610a84565b95505060a0610cfa8c828d01610a84565b94505060c0610d0b8c828d01610a84565b93505060e0610d1c8c828d01610a84565b925050610100610c6a8c828d01610a84565b60008060208385031215610d4157600080fd5b823567ffffffffffffffff811115610d5857600080fd5b610d6485828601610a3b565b92509250509250929050565b6000806000806000806000806080898b031215610d8c57600080fd5b883567ffffffffffffffff811115610da357600080fd5b610daf8b828c01610a3b565b9850985050602089013567ffffffffffffffff811115610dce57600080fd5b610dda8b828c01610a3b565b9650965050604089013567ffffffffffffffff811115610df957600080fd5b610e058b828c01610a3b565b9450945050606089013567ffffffffffffffff811115610e2457600080fd5b610e308b828c01610a3b565b92509250509295985092959890939650565b60008060208385031215610e5557600080fd5b823567ffffffffffffffff811115610e6c57600080fd5b610d6485828601610a8f565b600060208284031215610e8a57600080fd5b6000610aef8484610a84565b600080600060408486031215610eab57600080fd5b6000610eb78686610a84565b935050602084013567ffffffffffffffff811115610ed457600080fd5b610ee086828701610a8f565b92509250509250925092565b60008060408385031215610eff57600080fd5b6000610b9d8585610a84565b610f1481611260565b82525050565b6000610f268385611257565b93506001600160fb1b03831115610f3c57600080fd5b602083029250610f4d83858461127a565b50500190565b610f148161126b565b6000610f688385611257565b9350610f7583858461127a565b610f7e83611286565b9093019392505050565b6000610f95602e83611257565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000610fe5601983611257565b7f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000815260200192915050565b600061101e600c83611257565b6b6e6f742063616c6c61626c6560a01b815260200192915050565b604081016110478285610f0b565b6110546020830184610f0b565b9392505050565b606081016110698286610f0b565b6110766020830185610f53565b610aef6040830184610f53565b604081016110918285610f0b565b6110546020830184610f53565b61010081016110ad828c610f0b565b6110ba602083018b610f53565b81810360408301526110cd81898b610f5c565b90506110dc6060830188610f53565b6110e96080830187610f53565b6110f660a0830186610f53565b61110360c0830185610f53565b61111060e0830184610f53565b9a9950505050505050505050565b610120810161112d828c610f0b565b61113a602083018b610f53565b611147604083018a610f53565b6111546060830189610f53565b6111616080830188610f53565b61116e60a0830187610f53565b61117b60c0830186610f53565b61118860e0830185610f53565b611110610100830184610f53565b60208082528101610aef818486610f1a565b608080825281016111ba818a8c610f1a565b905081810360208301526111cf81888a610f1a565b905081810360408301526111e4818688610f1a565b90508181036060830152611110818486610f1a565b60208082528101610aef818486610f5c565b60208082528101610a3581610f88565b60208082528101610a3581610fd8565b60208082528101610a3581611011565b60208101610a358284610f53565b604081016110918285610f53565b90815260200190565b6000610a358261126e565b90565b6001600160a01b031690565b82818337506000910152565b601f01601f191690565b61129981611260565b81146112a457600080fd5b50565b6112998161126b56fea365627a7a7231582021a36812ff819247a0e389ca6ea90cce1b92938519ec5ecd93e0eb2068d7820d6c6578706572696d656e74616cf564736f6c63430005110040"
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
//	parsed, err := abi.JSON(strings.NewReader(ContractABI))
//	if err != nil {
//		return nil, err
//	}
//	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractTransactor) AdvanceEpochs(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "advanceEpochs", num)
//}
//
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
//}
//
//// AdvanceEpochs is a paid mutator transaction binding the contract method 0x0aeeca00.
////
//// Solidity: function advanceEpochs(uint256 num) returns()
//func (_Contract *ContractTransactorSession) AdvanceEpochs(num *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.AdvanceEpochs(&_Contract.TransactOpts, num)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractTransactor) CopyCode(opts *bind.TransactOpts, acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "copyCode", acc, from)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
//}
//
//// CopyCode is a paid mutator transaction binding the contract method 0xd6a0c7af.
////
//// Solidity: function copyCode(address acc, address from) returns()
//func (_Contract *ContractTransactorSession) CopyCode(acc common.Address, from common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.CopyCode(&_Contract.TransactOpts, acc, from)
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
//// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
////
//// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
//func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "initialize", _backend, _evmWriterAddress)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
////
//// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
//func (_Contract *ContractSession) Initialize(_backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _backend, _evmWriterAddress)
//}
//
//// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
////
//// Solidity: function initialize(address _backend, address _evmWriterAddress) returns()
//func (_Contract *ContractTransactorSession) Initialize(_backend common.Address, _evmWriterAddress common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _backend, _evmWriterAddress)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "sealEpoch", offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
//}
//
//// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
////
//// Solidity: function sealEpoch(uint256[] offlineTimes, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
//func (_Contract *ContractTransactorSession) SealEpoch(offlineTimes []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SealEpoch(&_Contract.TransactOpts, offlineTimes, offlineBlocks, uptimes, originatedTxsFee)
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
//// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
////
//// Solidity: function setBackend(address _backend) returns()
//func (_Contract *ContractTransactor) SetBackend(opts *bind.TransactOpts, _backend common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setBackend", _backend)
//}
//
//// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
////
//// Solidity: function setBackend(address _backend) returns()
//func (_Contract *ContractSession) SetBackend(_backend common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.SetBackend(&_Contract.TransactOpts, _backend)
//}
//
//// SetBackend is a paid mutator transaction binding the contract method 0xda7fc24f.
////
//// Solidity: function setBackend(address _backend) returns()
//func (_Contract *ContractTransactorSession) SetBackend(_backend common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.SetBackend(&_Contract.TransactOpts, _backend)
//}
//
//// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
////
//// Solidity: function setBalance(address acc, uint256 value) returns()
//func (_Contract *ContractTransactor) SetBalance(opts *bind.TransactOpts, acc common.Address, value *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setBalance", acc, value)
//}
//
//// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
////
//// Solidity: function setBalance(address acc, uint256 value) returns()
//func (_Contract *ContractSession) SetBalance(acc common.Address, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, acc, value)
//}
//
//// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
////
//// Solidity: function setBalance(address acc, uint256 value) returns()
//func (_Contract *ContractTransactorSession) SetBalance(acc common.Address, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetBalance(&_Contract.TransactOpts, acc, value)
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
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, _auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setGenesisValidator", _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetGenesisValidator is a paid mutator transaction binding the contract method 0x4feb92f3.
////
//// Solidity: function setGenesisValidator(address _auth, uint256 validatorID, bytes pubkey, uint256 status, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedEpoch, uint256 deactivatedTime) returns()
//func (_Contract *ContractTransactorSession) SetGenesisValidator(_auth common.Address, validatorID *big.Int, pubkey []byte, status *big.Int, createdEpoch *big.Int, createdTime *big.Int, deactivatedEpoch *big.Int, deactivatedTime *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.SetGenesisValidator(&_Contract.TransactOpts, _auth, validatorID, pubkey, status, createdEpoch, createdTime, deactivatedEpoch, deactivatedTime)
//}
//
//// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
////
//// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
//func (_Contract *ContractTransactor) SetStorage(opts *bind.TransactOpts, acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "setStorage", acc, key, value)
//}
//
//// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
////
//// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
//func (_Contract *ContractSession) SetStorage(acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
//	return _Contract.Contract.SetStorage(&_Contract.TransactOpts, acc, key, value)
//}
//
//// SetStorage is a paid mutator transaction binding the contract method 0x39e503ab.
////
//// Solidity: function setStorage(address acc, bytes32 key, bytes32 value) returns()
//func (_Contract *ContractTransactorSession) SetStorage(acc common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
//	return _Contract.Contract.SetStorage(&_Contract.TransactOpts, acc, key, value)
//}
//
//// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
////
//// Solidity: function swapCode(address acc, address with) returns()
//func (_Contract *ContractTransactor) SwapCode(opts *bind.TransactOpts, acc common.Address, with common.Address) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "swapCode", acc, with)
//}
//
//// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
////
//// Solidity: function swapCode(address acc, address with) returns()
//func (_Contract *ContractSession) SwapCode(acc common.Address, with common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.SwapCode(&_Contract.TransactOpts, acc, with)
//}
//
//// SwapCode is a paid mutator transaction binding the contract method 0x07690b2a.
////
//// Solidity: function swapCode(address acc, address with) returns()
//func (_Contract *ContractTransactorSession) SwapCode(acc common.Address, with common.Address) (*types.Transaction, error) {
//	return _Contract.Contract.SwapCode(&_Contract.TransactOpts, acc, with)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractTransactor) UpdateNetworkRules(opts *bind.TransactOpts, diff []byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateNetworkRules", diff)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
//}
//
//// UpdateNetworkRules is a paid mutator transaction binding the contract method 0xb9cc6b1c.
////
//// Solidity: function updateNetworkRules(bytes diff) returns()
//func (_Contract *ContractTransactorSession) UpdateNetworkRules(diff []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkRules(&_Contract.TransactOpts, diff)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractTransactor) UpdateNetworkVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateNetworkVersion", version)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
//}
//
//// UpdateNetworkVersion is a paid mutator transaction binding the contract method 0x267ab446.
////
//// Solidity: function updateNetworkVersion(uint256 version) returns()
//func (_Contract *ContractTransactorSession) UpdateNetworkVersion(version *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateNetworkVersion(&_Contract.TransactOpts, version)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractTransactor) UpdateValidatorPubkey(opts *bind.TransactOpts, validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateValidatorPubkey", validatorID, pubkey)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
//}
//
//// UpdateValidatorPubkey is a paid mutator transaction binding the contract method 0x242a6e3f.
////
//// Solidity: function updateValidatorPubkey(uint256 validatorID, bytes pubkey) returns()
//func (_Contract *ContractTransactorSession) UpdateValidatorPubkey(validatorID *big.Int, pubkey []byte) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorPubkey(&_Contract.TransactOpts, validatorID, pubkey)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractTransactor) UpdateValidatorWeight(opts *bind.TransactOpts, validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.contract.Transact(opts, "updateValidatorWeight", validatorID, value)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
//}
//
//// UpdateValidatorWeight is a paid mutator transaction binding the contract method 0xa4066fbe.
////
//// Solidity: function updateValidatorWeight(uint256 validatorID, uint256 value) returns()
//func (_Contract *ContractTransactorSession) UpdateValidatorWeight(validatorID *big.Int, value *big.Int) (*types.Transaction, error) {
//	return _Contract.Contract.UpdateValidatorWeight(&_Contract.TransactOpts, validatorID, value)
//}
//
//// ContractAdvanceEpochsIterator is returned from FilterAdvanceEpochs and is used to iterate over the raw logs and unpacked data for AdvanceEpochs events raised by the Contract contract.
//type ContractAdvanceEpochsIterator struct {
//	Event *ContractAdvanceEpochs // Event containing the contract specifics and raw log
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
//func (it *ContractAdvanceEpochsIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractAdvanceEpochs)
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
//		it.Event = new(ContractAdvanceEpochs)
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
//func (it *ContractAdvanceEpochsIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractAdvanceEpochsIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractAdvanceEpochs represents a AdvanceEpochs event raised by the Contract contract.
//type ContractAdvanceEpochs struct {
//	Num *big.Int
//	Raw types.Log // Blockchain specific contextual infos
//}
//
//// FilterAdvanceEpochs is a free log retrieval operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
////
//// Solidity: event AdvanceEpochs(uint256 num)
//func (_Contract *ContractFilterer) FilterAdvanceEpochs(opts *bind.FilterOpts) (*ContractAdvanceEpochsIterator, error) {
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdvanceEpochs")
//	if err != nil {
//		return nil, err
//	}
//	return &ContractAdvanceEpochsIterator{contract: _Contract.contract, event: "AdvanceEpochs", logs: logs, sub: sub}, nil
//}
//
//// WatchAdvanceEpochs is a free log subscription operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
////
//// Solidity: event AdvanceEpochs(uint256 num)
//func (_Contract *ContractFilterer) WatchAdvanceEpochs(opts *bind.WatchOpts, sink chan<- *ContractAdvanceEpochs) (event.Subscription, error) {
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdvanceEpochs")
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractAdvanceEpochs)
//				if err := _Contract.contract.UnpackLog(event, "AdvanceEpochs", log); err != nil {
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
//// ParseAdvanceEpochs is a log parse operation binding the contract event 0x0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd1.
////
//// Solidity: event AdvanceEpochs(uint256 num)
//func (_Contract *ContractFilterer) ParseAdvanceEpochs(log types.Log) (*ContractAdvanceEpochs, error) {
//	event := new(ContractAdvanceEpochs)
//	if err := _Contract.contract.UnpackLog(event, "AdvanceEpochs", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdateNetworkRulesIterator is returned from FilterUpdateNetworkRules and is used to iterate over the raw logs and unpacked data for UpdateNetworkRules events raised by the Contract contract.
//type ContractUpdateNetworkRulesIterator struct {
//	Event *ContractUpdateNetworkRules // Event containing the contract specifics and raw log
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
//func (it *ContractUpdateNetworkRulesIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdateNetworkRules)
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
//		it.Event = new(ContractUpdateNetworkRules)
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
//func (it *ContractUpdateNetworkRulesIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdateNetworkRulesIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdateNetworkRules represents a UpdateNetworkRules event raised by the Contract contract.
//type ContractUpdateNetworkRules struct {
//	Diff []byte
//	Raw  types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdateNetworkRules is a free log retrieval operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
////
//// Solidity: event UpdateNetworkRules(bytes diff)
//func (_Contract *ContractFilterer) FilterUpdateNetworkRules(opts *bind.FilterOpts) (*ContractUpdateNetworkRulesIterator, error) {
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateNetworkRules")
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdateNetworkRulesIterator{contract: _Contract.contract, event: "UpdateNetworkRules", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdateNetworkRules is a free log subscription operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
////
//// Solidity: event UpdateNetworkRules(bytes diff)
//func (_Contract *ContractFilterer) WatchUpdateNetworkRules(opts *bind.WatchOpts, sink chan<- *ContractUpdateNetworkRules) (event.Subscription, error) {
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateNetworkRules")
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdateNetworkRules)
//				if err := _Contract.contract.UnpackLog(event, "UpdateNetworkRules", log); err != nil {
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
//// ParseUpdateNetworkRules is a log parse operation binding the contract event 0x47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99.
////
//// Solidity: event UpdateNetworkRules(bytes diff)
//func (_Contract *ContractFilterer) ParseUpdateNetworkRules(log types.Log) (*ContractUpdateNetworkRules, error) {
//	event := new(ContractUpdateNetworkRules)
//	if err := _Contract.contract.UnpackLog(event, "UpdateNetworkRules", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdateNetworkVersionIterator is returned from FilterUpdateNetworkVersion and is used to iterate over the raw logs and unpacked data for UpdateNetworkVersion events raised by the Contract contract.
//type ContractUpdateNetworkVersionIterator struct {
//	Event *ContractUpdateNetworkVersion // Event containing the contract specifics and raw log
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
//func (it *ContractUpdateNetworkVersionIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdateNetworkVersion)
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
//		it.Event = new(ContractUpdateNetworkVersion)
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
//func (it *ContractUpdateNetworkVersionIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdateNetworkVersionIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdateNetworkVersion represents a UpdateNetworkVersion event raised by the Contract contract.
//type ContractUpdateNetworkVersion struct {
//	Version *big.Int
//	Raw     types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdateNetworkVersion is a free log retrieval operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
////
//// Solidity: event UpdateNetworkVersion(uint256 version)
//func (_Contract *ContractFilterer) FilterUpdateNetworkVersion(opts *bind.FilterOpts) (*ContractUpdateNetworkVersionIterator, error) {
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateNetworkVersion")
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdateNetworkVersionIterator{contract: _Contract.contract, event: "UpdateNetworkVersion", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdateNetworkVersion is a free log subscription operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
////
//// Solidity: event UpdateNetworkVersion(uint256 version)
//func (_Contract *ContractFilterer) WatchUpdateNetworkVersion(opts *bind.WatchOpts, sink chan<- *ContractUpdateNetworkVersion) (event.Subscription, error) {
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateNetworkVersion")
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdateNetworkVersion)
//				if err := _Contract.contract.UnpackLog(event, "UpdateNetworkVersion", log); err != nil {
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
//// ParseUpdateNetworkVersion is a log parse operation binding the contract event 0x2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb.
////
//// Solidity: event UpdateNetworkVersion(uint256 version)
//func (_Contract *ContractFilterer) ParseUpdateNetworkVersion(log types.Log) (*ContractUpdateNetworkVersion, error) {
//	event := new(ContractUpdateNetworkVersion)
//	if err := _Contract.contract.UnpackLog(event, "UpdateNetworkVersion", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdateValidatorPubkeyIterator is returned from FilterUpdateValidatorPubkey and is used to iterate over the raw logs and unpacked data for UpdateValidatorPubkey events raised by the Contract contract.
//type ContractUpdateValidatorPubkeyIterator struct {
//	Event *ContractUpdateValidatorPubkey // Event containing the contract specifics and raw log
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
//func (it *ContractUpdateValidatorPubkeyIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdateValidatorPubkey)
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
//		it.Event = new(ContractUpdateValidatorPubkey)
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
//func (it *ContractUpdateValidatorPubkeyIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdateValidatorPubkeyIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdateValidatorPubkey represents a UpdateValidatorPubkey event raised by the Contract contract.
//type ContractUpdateValidatorPubkey struct {
//	ValidatorID *big.Int
//	Pubkey      []byte
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdateValidatorPubkey is a free log retrieval operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
////
//// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
//func (_Contract *ContractFilterer) FilterUpdateValidatorPubkey(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdateValidatorPubkeyIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateValidatorPubkey", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdateValidatorPubkeyIterator{contract: _Contract.contract, event: "UpdateValidatorPubkey", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdateValidatorPubkey is a free log subscription operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
////
//// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
//func (_Contract *ContractFilterer) WatchUpdateValidatorPubkey(opts *bind.WatchOpts, sink chan<- *ContractUpdateValidatorPubkey, validatorID []*big.Int) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateValidatorPubkey", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdateValidatorPubkey)
//				if err := _Contract.contract.UnpackLog(event, "UpdateValidatorPubkey", log); err != nil {
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
//// ParseUpdateValidatorPubkey is a log parse operation binding the contract event 0x0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e.
////
//// Solidity: event UpdateValidatorPubkey(uint256 indexed validatorID, bytes pubkey)
//func (_Contract *ContractFilterer) ParseUpdateValidatorPubkey(log types.Log) (*ContractUpdateValidatorPubkey, error) {
//	event := new(ContractUpdateValidatorPubkey)
//	if err := _Contract.contract.UnpackLog(event, "UpdateValidatorPubkey", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdateValidatorWeightIterator is returned from FilterUpdateValidatorWeight and is used to iterate over the raw logs and unpacked data for UpdateValidatorWeight events raised by the Contract contract.
//type ContractUpdateValidatorWeightIterator struct {
//	Event *ContractUpdateValidatorWeight // Event containing the contract specifics and raw log
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
//func (it *ContractUpdateValidatorWeightIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdateValidatorWeight)
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
//		it.Event = new(ContractUpdateValidatorWeight)
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
//func (it *ContractUpdateValidatorWeightIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdateValidatorWeightIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdateValidatorWeight represents a UpdateValidatorWeight event raised by the Contract contract.
//type ContractUpdateValidatorWeight struct {
//	ValidatorID *big.Int
//	Weight      *big.Int
//	Raw         types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdateValidatorWeight is a free log retrieval operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
////
//// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
//func (_Contract *ContractFilterer) FilterUpdateValidatorWeight(opts *bind.FilterOpts, validatorID []*big.Int) (*ContractUpdateValidatorWeightIterator, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateValidatorWeight", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdateValidatorWeightIterator{contract: _Contract.contract, event: "UpdateValidatorWeight", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdateValidatorWeight is a free log subscription operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
////
//// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
//func (_Contract *ContractFilterer) WatchUpdateValidatorWeight(opts *bind.WatchOpts, sink chan<- *ContractUpdateValidatorWeight, validatorID []*big.Int) (event.Subscription, error) {
//
//	var validatorIDRule []interface{}
//	for _, validatorIDItem := range validatorID {
//		validatorIDRule = append(validatorIDRule, validatorIDItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateValidatorWeight", validatorIDRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdateValidatorWeight)
//				if err := _Contract.contract.UnpackLog(event, "UpdateValidatorWeight", log); err != nil {
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
//// ParseUpdateValidatorWeight is a log parse operation binding the contract event 0xb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935.
////
//// Solidity: event UpdateValidatorWeight(uint256 indexed validatorID, uint256 weight)
//func (_Contract *ContractFilterer) ParseUpdateValidatorWeight(log types.Log) (*ContractUpdateValidatorWeight, error) {
//	event := new(ContractUpdateValidatorWeight)
//	if err := _Contract.contract.UnpackLog(event, "UpdateValidatorWeight", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//// ContractUpdatedBackendIterator is returned from FilterUpdatedBackend and is used to iterate over the raw logs and unpacked data for UpdatedBackend events raised by the Contract contract.
//type ContractUpdatedBackendIterator struct {
//	Event *ContractUpdatedBackend // Event containing the contract specifics and raw log
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
//func (it *ContractUpdatedBackendIterator) Next() bool {
//	// If the iterator failed, stop iterating
//	if it.fail != nil {
//		return false
//	}
//	// If the iterator completed, deliver directly whatever's available
//	if it.done {
//		select {
//		case log := <-it.logs:
//			it.Event = new(ContractUpdatedBackend)
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
//		it.Event = new(ContractUpdatedBackend)
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
//func (it *ContractUpdatedBackendIterator) Error() error {
//	return it.fail
//}
//
//// Close terminates the iteration process, releasing any pending underlying
//// resources.
//func (it *ContractUpdatedBackendIterator) Close() error {
//	it.sub.Unsubscribe()
//	return nil
//}
//
//// ContractUpdatedBackend represents a UpdatedBackend event raised by the Contract contract.
//type ContractUpdatedBackend struct {
//	Backend common.Address
//	Raw     types.Log // Blockchain specific contextual infos
//}
//
//// FilterUpdatedBackend is a free log retrieval operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
////
//// Solidity: event UpdatedBackend(address indexed backend)
//func (_Contract *ContractFilterer) FilterUpdatedBackend(opts *bind.FilterOpts, backend []common.Address) (*ContractUpdatedBackendIterator, error) {
//
//	var backendRule []interface{}
//	for _, backendItem := range backend {
//		backendRule = append(backendRule, backendItem)
//	}
//
//	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdatedBackend", backendRule)
//	if err != nil {
//		return nil, err
//	}
//	return &ContractUpdatedBackendIterator{contract: _Contract.contract, event: "UpdatedBackend", logs: logs, sub: sub}, nil
//}
//
//// WatchUpdatedBackend is a free log subscription operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
////
//// Solidity: event UpdatedBackend(address indexed backend)
//func (_Contract *ContractFilterer) WatchUpdatedBackend(opts *bind.WatchOpts, sink chan<- *ContractUpdatedBackend, backend []common.Address) (event.Subscription, error) {
//
//	var backendRule []interface{}
//	for _, backendItem := range backend {
//		backendRule = append(backendRule, backendItem)
//	}
//
//	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdatedBackend", backendRule)
//	if err != nil {
//		return nil, err
//	}
//	return event.NewSubscription(func(quit <-chan struct{}) error {
//		defer sub.Unsubscribe()
//		for {
//			select {
//			case log := <-logs:
//				// New log arrived, parse the event and forward to the user
//				event := new(ContractUpdatedBackend)
//				if err := _Contract.contract.UnpackLog(event, "UpdatedBackend", log); err != nil {
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
//// ParseUpdatedBackend is a log parse operation binding the contract event 0x64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef245035830693.
////
//// Solidity: event UpdatedBackend(address indexed backend)
//func (_Contract *ContractFilterer) ParseUpdatedBackend(log types.Log) (*ContractUpdatedBackend, error) {
//	event := new(ContractUpdatedBackend)
//	if err := _Contract.contract.UnpackLog(event, "UpdatedBackend", log); err != nil {
//		return nil, err
//	}
//	event.Raw = log
//	return event, nil
//}
//
//var ContractBinRuntime = "0x608060405234801561001057600080fd5b506004361061010b5760003560e01c80634feb92f3116100a2578063d6a0c7af11610071578063d6a0c7af146101f6578063da7fc24f14610209578063e08d7e661461021c578063e30443bc1461022f578063ebdf104c146102425761010b565b80634feb92f3146101aa57806379bead38146101bd578063a4066fbe146101d0578063b9cc6b1c146101e35761010b565b8063242a6e3f116100de578063242a6e3f1461015e578063267ab4461461017157806339e503ab14610184578063485cc955146101975761010b565b806307690b2a146101105780630aeeca001461012557806318f628d4146101385780631e702f831461014b575b600080fd5b61012361011e366004610af7565b610255565b005b610123610133366004610e78565b6102f0565b610123610146366004610c7a565b610354565b610123610159366004610eec565b6103ef565b61012361016c366004610e96565b61043f565b61012361017f366004610e78565b6104a8565b610123610192366004610b31565b610501565b6101236101a5366004610af7565b610596565b6101236101b8366004610bae565b61067d565b6101236101cb366004610b7e565b6106db565b6101236101de366004610eec565b610737565b6101236101f1366004610e42565b61079d565b610123610204366004610af7565b610804565b610123610217366004610ad1565b610860565b61012361022a366004610d2e565b6108e0565b61012361023d366004610b7e565b610930565b610123610250366004610d70565b61098c565b6034546001600160a01b031633146102885760405162461bcd60e51b815260040161027f9061121b565b60405180910390fd5b6035546040516303b4859560e11b81526001600160a01b03909116906307690b2a906102ba9085908590600401611039565b600060405180830381600087803b1580156102d457600080fd5b505af11580156102e8573d6000803e3d6000fd5b505050505050565b6034546001600160a01b0316331461031a5760405162461bcd60e51b815260040161027f9061121b565b7f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd181604051610349919061123b565b60405180910390a150565b33156103725760405162461bcd60e51b815260040161027f9061122b565b60345460405163063d8a3560e21b81526001600160a01b03909116906318f628d4906103b2908c908c908c908c908c908c908c908c908c9060040161111e565b600060405180830381600087803b1580156103cc57600080fd5b505af11580156103e0573d6000803e3d6000fd5b50505050505050505050505050565b331561040d5760405162461bcd60e51b815260040161027f9061122b565b603454604051631e702f8360e01b81526001600160a01b0390911690631e702f83906102ba9085908590600401611249565b6034546001600160a01b031633146104695760405162461bcd60e51b815260040161027f9061121b565b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405161049b9291906111f9565b60405180910390a2505050565b6034546001600160a01b031633146104d25760405162461bcd60e51b815260040161027f9061121b565b7f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb81604051610349919061123b565b6034546001600160a01b0316331461052b5760405162461bcd60e51b815260040161027f9061121b565b6035546040516339e503ab60e01b81526001600160a01b03909116906339e503ab9061055f9086908690869060040161105b565b600060405180830381600087803b15801561057957600080fd5b505af115801561058d573d6000803e3d6000fd5b50505050505050565b600054610100900460ff16806105af57506105af610a24565b806105bd575060005460ff16155b6105d95760405162461bcd60e51b815260040161027f9061120b565b600054610100900460ff16158015610604576000805460ff1961ff0019909116610100171660011790555b603480546001600160a01b0319166001600160a01b0385169081179091556040517f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603580546001600160a01b0319166001600160a01b0384161790558015610678576000805461ff00191690555b505050565b331561069b5760405162461bcd60e51b815260040161027f9061122b565b603454604051634feb92f360e01b81526001600160a01b0390911690634feb92f3906103b2908c908c908c908c908c908c908c908c908c9060040161109e565b6034546001600160a01b031633146107055760405162461bcd60e51b815260040161027f9061121b565b603554604051630f37d5a760e31b81526001600160a01b03909116906379bead38906102ba9085908590600401611083565b6034546001600160a01b031633146107615760405162461bcd60e51b815260040161027f9061121b565b817fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a93582604051610791919061123b565b60405180910390a25050565b6034546001600160a01b031633146107c75760405162461bcd60e51b815260040161027f9061121b565b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb9982826040516107f89291906111f9565b60405180910390a15050565b6034546001600160a01b0316331461082e5760405162461bcd60e51b815260040161027f9061121b565b60355460405163d6a0c7af60e01b81526001600160a01b039091169063d6a0c7af906102ba9085908590600401611039565b6034546001600160a01b0316331461088a5760405162461bcd60e51b815260040161027f9061121b565b6040516001600160a01b038216907f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603480546001600160a01b0319166001600160a01b0392909216919091179055565b33156108fe5760405162461bcd60e51b815260040161027f9061122b565b603454604051637046bf3360e11b81526001600160a01b039091169063e08d7e66906102ba9085908590600401611196565b6034546001600160a01b0316331461095a5760405162461bcd60e51b815260040161027f9061121b565b6035546040516338c110ef60e21b81526001600160a01b039091169063e30443bc906102ba9085908590600401611083565b33156109aa5760405162461bcd60e51b815260040161027f9061122b565b603454604051633af7c41360e21b81526001600160a01b039091169063ebdf104c906109e8908b908b908b908b908b908b908b908b906004016111a8565b600060405180830381600087803b158015610a0257600080fd5b505af1158015610a16573d6000803e3d6000fd5b505050505050505050505050565b303b1590565b8035610a3581611290565b92915050565b60008083601f840112610a4d57600080fd5b50813567ffffffffffffffff811115610a6557600080fd5b602083019150836020820283011115610a7d57600080fd5b9250929050565b8035610a35816112a7565b60008083601f840112610aa157600080fd5b50813567ffffffffffffffff811115610ab957600080fd5b602083019150836001820283011115610a7d57600080fd5b600060208284031215610ae357600080fd5b6000610aef8484610a2a565b949350505050565b60008060408385031215610b0a57600080fd5b6000610b168585610a2a565b9250506020610b2785828601610a2a565b9150509250929050565b600080600060608486031215610b4657600080fd5b6000610b528686610a2a565b9350506020610b6386828701610a84565b9250506040610b7486828701610a84565b9150509250925092565b60008060408385031215610b9157600080fd5b6000610b9d8585610a2a565b9250506020610b2785828601610a84565b60008060008060008060008060006101008a8c031215610bcd57600080fd5b6000610bd98c8c610a2a565b9950506020610bea8c828d01610a84565b98505060408a013567ffffffffffffffff811115610c0757600080fd5b610c138c828d01610a8f565b97509750506060610c268c828d01610a84565b9550506080610c378c828d01610a84565b94505060a0610c488c828d01610a84565b93505060c0610c598c828d01610a84565b92505060e0610c6a8c828d01610a84565b9150509295985092959850929598565b60008060008060008060008060006101208a8c031215610c9957600080fd5b6000610ca58c8c610a2a565b9950506020610cb68c828d01610a84565b9850506040610cc78c828d01610a84565b9750506060610cd88c828d01610a84565b9650506080610ce98c828d01610a84565b95505060a0610cfa8c828d01610a84565b94505060c0610d0b8c828d01610a84565b93505060e0610d1c8c828d01610a84565b925050610100610c6a8c828d01610a84565b60008060208385031215610d4157600080fd5b823567ffffffffffffffff811115610d5857600080fd5b610d6485828601610a3b565b92509250509250929050565b6000806000806000806000806080898b031215610d8c57600080fd5b883567ffffffffffffffff811115610da357600080fd5b610daf8b828c01610a3b565b9850985050602089013567ffffffffffffffff811115610dce57600080fd5b610dda8b828c01610a3b565b9650965050604089013567ffffffffffffffff811115610df957600080fd5b610e058b828c01610a3b565b9450945050606089013567ffffffffffffffff811115610e2457600080fd5b610e308b828c01610a3b565b92509250509295985092959890939650565b60008060208385031215610e5557600080fd5b823567ffffffffffffffff811115610e6c57600080fd5b610d6485828601610a8f565b600060208284031215610e8a57600080fd5b6000610aef8484610a84565b600080600060408486031215610eab57600080fd5b6000610eb78686610a84565b935050602084013567ffffffffffffffff811115610ed457600080fd5b610ee086828701610a8f565b92509250509250925092565b60008060408385031215610eff57600080fd5b6000610b9d8585610a84565b610f1481611260565b82525050565b6000610f268385611257565b93506001600160fb1b03831115610f3c57600080fd5b602083029250610f4d83858461127a565b50500190565b610f148161126b565b6000610f688385611257565b9350610f7583858461127a565b610f7e83611286565b9093019392505050565b6000610f95602e83611257565b7f436f6e747261637420696e7374616e63652068617320616c726561647920626581526d195b881a5b9a5d1a585b1a5e995960921b602082015260400192915050565b6000610fe5601983611257565b7f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000815260200192915050565b600061101e600c83611257565b6b6e6f742063616c6c61626c6560a01b815260200192915050565b604081016110478285610f0b565b6110546020830184610f0b565b9392505050565b606081016110698286610f0b565b6110766020830185610f53565b610aef6040830184610f53565b604081016110918285610f0b565b6110546020830184610f53565b61010081016110ad828c610f0b565b6110ba602083018b610f53565b81810360408301526110cd81898b610f5c565b90506110dc6060830188610f53565b6110e96080830187610f53565b6110f660a0830186610f53565b61110360c0830185610f53565b61111060e0830184610f53565b9a9950505050505050505050565b610120810161112d828c610f0b565b61113a602083018b610f53565b611147604083018a610f53565b6111546060830189610f53565b6111616080830188610f53565b61116e60a0830187610f53565b61117b60c0830186610f53565b61118860e0830185610f53565b611110610100830184610f53565b60208082528101610aef818486610f1a565b608080825281016111ba818a8c610f1a565b905081810360208301526111cf81888a610f1a565b905081810360408301526111e4818688610f1a565b90508181036060830152611110818486610f1a565b60208082528101610aef818486610f5c565b60208082528101610a3581610f88565b60208082528101610a3581610fd8565b60208082528101610a3581611011565b60208101610a358284610f53565b604081016110918285610f53565b90815260200190565b6000610a358261126e565b90565b6001600160a01b031690565b82818337506000910152565b601f01601f191690565b61129981611260565b81146112a457600080fd5b50565b6112998161126b56fea365627a7a7231582021a36812ff819247a0e389ca6ea90cce1b92938519ec5ecd93e0eb2068d7820d6c6578706572696d656e74616cf564736f6c63430005110040"
