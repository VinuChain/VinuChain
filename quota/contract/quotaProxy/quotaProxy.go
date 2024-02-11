// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quotaProxy

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

// QuotaProxyMetaData contains all meta data concerning the QuotaProxy contract.
var QuotaProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFeeRefundBlockCount\",\"type\":\"uint16\"}],\"name\":\"FeeRefundBlockCountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinStake\",\"type\":\"uint256\"}],\"name\":\"MinStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addressTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataStorage\",\"outputs\":[{\"internalType\":\"contractDataStorageContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRefundBlockCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dataStorageAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_feeRefundBlockCount\",\"type\":\"uint16\"}],\"name\":\"setFeeRefundBlockCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStake\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quotaContractAddress_\",\"type\":\"address\"}],\"name\":\"setQuotaContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// QuotaProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use QuotaProxyMetaData.ABI instead.
var QuotaProxyABI = QuotaProxyMetaData.ABI

// QuotaProxy is an auto generated Go binding around an Ethereum contract.
type QuotaProxy struct {
	QuotaProxyCaller     // Read-only binding to the contract
	QuotaProxyTransactor // Write-only binding to the contract
	QuotaProxyFilterer   // Log filterer for contract events
}

// QuotaProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuotaProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuotaProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuotaProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuotaProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuotaProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuotaProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuotaProxySession struct {
	Contract     *QuotaProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuotaProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuotaProxyCallerSession struct {
	Contract *QuotaProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// QuotaProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuotaProxyTransactorSession struct {
	Contract     *QuotaProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// QuotaProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuotaProxyRaw struct {
	Contract *QuotaProxy // Generic contract binding to access the raw methods on
}

// QuotaProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuotaProxyCallerRaw struct {
	Contract *QuotaProxyCaller // Generic read-only contract binding to access the raw methods on
}

// QuotaProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuotaProxyTransactorRaw struct {
	Contract *QuotaProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuotaProxy creates a new instance of QuotaProxy, bound to a specific deployed contract.
func NewQuotaProxy(address common.Address, backend bind.ContractBackend) (*QuotaProxy, error) {
	contract, err := bindQuotaProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuotaProxy{QuotaProxyCaller: QuotaProxyCaller{contract: contract}, QuotaProxyTransactor: QuotaProxyTransactor{contract: contract}, QuotaProxyFilterer: QuotaProxyFilterer{contract: contract}}, nil
}

// NewQuotaProxyCaller creates a new read-only instance of QuotaProxy, bound to a specific deployed contract.
func NewQuotaProxyCaller(address common.Address, caller bind.ContractCaller) (*QuotaProxyCaller, error) {
	contract, err := bindQuotaProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyCaller{contract: contract}, nil
}

// NewQuotaProxyTransactor creates a new write-only instance of QuotaProxy, bound to a specific deployed contract.
func NewQuotaProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*QuotaProxyTransactor, error) {
	contract, err := bindQuotaProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyTransactor{contract: contract}, nil
}

// NewQuotaProxyFilterer creates a new log filterer instance of QuotaProxy, bound to a specific deployed contract.
func NewQuotaProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*QuotaProxyFilterer, error) {
	contract, err := bindQuotaProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyFilterer{contract: contract}, nil
}

// bindQuotaProxy binds a generic wrapper to an already deployed contract.
func bindQuotaProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuotaProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuotaProxy *QuotaProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuotaProxy.Contract.QuotaProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuotaProxy *QuotaProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuotaProxy.Contract.QuotaProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuotaProxy *QuotaProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuotaProxy.Contract.QuotaProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuotaProxy *QuotaProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuotaProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuotaProxy *QuotaProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuotaProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuotaProxy *QuotaProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuotaProxy.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_QuotaProxy *QuotaProxyCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_QuotaProxy *QuotaProxySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _QuotaProxy.Contract.UPGRADEINTERFACEVERSION(&_QuotaProxy.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_QuotaProxy *QuotaProxyCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _QuotaProxy.Contract.UPGRADEINTERFACEVERSION(&_QuotaProxy.CallOpts)
}

// AddressTotalStake is a free data retrieval call binding the contract method 0x65fdda5e.
//
// Solidity: function addressTotalStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) AddressTotalStake(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "addressTotalStake", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressTotalStake is a free data retrieval call binding the contract method 0x65fdda5e.
//
// Solidity: function addressTotalStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxySession) AddressTotalStake(sender common.Address) (*big.Int, error) {
	return _QuotaProxy.Contract.AddressTotalStake(&_QuotaProxy.CallOpts, sender)
}

// AddressTotalStake is a free data retrieval call binding the contract method 0x65fdda5e.
//
// Solidity: function addressTotalStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) AddressTotalStake(sender common.Address) (*big.Int, error) {
	return _QuotaProxy.Contract.AddressTotalStake(&_QuotaProxy.CallOpts, sender)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_QuotaProxy *QuotaProxyCaller) DataStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "dataStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_QuotaProxy *QuotaProxySession) DataStorage() (common.Address, error) {
	return _QuotaProxy.Contract.DataStorage(&_QuotaProxy.CallOpts)
}

// DataStorage is a free data retrieval call binding the contract method 0x8870455f.
//
// Solidity: function dataStorage() view returns(address)
func (_QuotaProxy *QuotaProxyCallerSession) DataStorage() (common.Address, error) {
	return _QuotaProxy.Contract.DataStorage(&_QuotaProxy.CallOpts)
}

// GetFeeRefundBlockCount is a free data retrieval call binding the contract method 0xe2e96253.
//
// Solidity: function getFeeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxyCaller) GetFeeRefundBlockCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "getFeeRefundBlockCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFeeRefundBlockCount is a free data retrieval call binding the contract method 0xe2e96253.
//
// Solidity: function getFeeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxySession) GetFeeRefundBlockCount() (uint16, error) {
	return _QuotaProxy.Contract.GetFeeRefundBlockCount(&_QuotaProxy.CallOpts)
}

// GetFeeRefundBlockCount is a free data retrieval call binding the contract method 0xe2e96253.
//
// Solidity: function getFeeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxyCallerSession) GetFeeRefundBlockCount() (uint16, error) {
	return _QuotaProxy.Contract.GetFeeRefundBlockCount(&_QuotaProxy.CallOpts)
}

// GetMinStake is a free data retrieval call binding the contract method 0x56a3b5fa.
//
// Solidity: function getMinStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) GetMinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "getMinStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStake is a free data retrieval call binding the contract method 0x56a3b5fa.
//
// Solidity: function getMinStake() view returns(uint256)
func (_QuotaProxy *QuotaProxySession) GetMinStake() (*big.Int, error) {
	return _QuotaProxy.Contract.GetMinStake(&_QuotaProxy.CallOpts)
}

// GetMinStake is a free data retrieval call binding the contract method 0x56a3b5fa.
//
// Solidity: function getMinStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) GetMinStake() (*big.Int, error) {
	return _QuotaProxy.Contract.GetMinStake(&_QuotaProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuotaProxy *QuotaProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuotaProxy *QuotaProxySession) Owner() (common.Address, error) {
	return _QuotaProxy.Contract.Owner(&_QuotaProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuotaProxy *QuotaProxyCallerSession) Owner() (common.Address, error) {
	return _QuotaProxy.Contract.Owner(&_QuotaProxy.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_QuotaProxy *QuotaProxyCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_QuotaProxy *QuotaProxySession) ProxiableUUID() ([32]byte, error) {
	return _QuotaProxy.Contract.ProxiableUUID(&_QuotaProxy.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_QuotaProxy *QuotaProxyCallerSession) ProxiableUUID() ([32]byte, error) {
	return _QuotaProxy.Contract.ProxiableUUID(&_QuotaProxy.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_QuotaProxy *QuotaProxySession) TotalStake() (*big.Int, error) {
	return _QuotaProxy.Contract.TotalStake(&_QuotaProxy.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) TotalStake() (*big.Int, error) {
	return _QuotaProxy.Contract.TotalStake(&_QuotaProxy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorageAddress) returns()
func (_QuotaProxy *QuotaProxyTransactor) Initialize(opts *bind.TransactOpts, _dataStorageAddress common.Address) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "initialize", _dataStorageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorageAddress) returns()
func (_QuotaProxy *QuotaProxySession) Initialize(_dataStorageAddress common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.Initialize(&_QuotaProxy.TransactOpts, _dataStorageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dataStorageAddress) returns()
func (_QuotaProxy *QuotaProxyTransactorSession) Initialize(_dataStorageAddress common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.Initialize(&_QuotaProxy.TransactOpts, _dataStorageAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuotaProxy *QuotaProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuotaProxy *QuotaProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _QuotaProxy.Contract.RenounceOwnership(&_QuotaProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuotaProxy *QuotaProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuotaProxy.Contract.RenounceOwnership(&_QuotaProxy.TransactOpts)
}

// SetFeeRefundBlockCount is a paid mutator transaction binding the contract method 0xd62b032c.
//
// Solidity: function setFeeRefundBlockCount(uint16 _feeRefundBlockCount) returns()
func (_QuotaProxy *QuotaProxyTransactor) SetFeeRefundBlockCount(opts *bind.TransactOpts, _feeRefundBlockCount uint16) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "setFeeRefundBlockCount", _feeRefundBlockCount)
}

// SetFeeRefundBlockCount is a paid mutator transaction binding the contract method 0xd62b032c.
//
// Solidity: function setFeeRefundBlockCount(uint16 _feeRefundBlockCount) returns()
func (_QuotaProxy *QuotaProxySession) SetFeeRefundBlockCount(_feeRefundBlockCount uint16) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetFeeRefundBlockCount(&_QuotaProxy.TransactOpts, _feeRefundBlockCount)
}

// SetFeeRefundBlockCount is a paid mutator transaction binding the contract method 0xd62b032c.
//
// Solidity: function setFeeRefundBlockCount(uint16 _feeRefundBlockCount) returns()
func (_QuotaProxy *QuotaProxyTransactorSession) SetFeeRefundBlockCount(_feeRefundBlockCount uint16) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetFeeRefundBlockCount(&_QuotaProxy.TransactOpts, _feeRefundBlockCount)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_QuotaProxy *QuotaProxyTransactor) SetMinStake(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "setMinStake", _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_QuotaProxy *QuotaProxySession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetMinStake(&_QuotaProxy.TransactOpts, _minStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 _minStake) returns()
func (_QuotaProxy *QuotaProxyTransactorSession) SetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetMinStake(&_QuotaProxy.TransactOpts, _minStake)
}

// SetQuotaContractAddress is a paid mutator transaction binding the contract method 0x2fa5c39e.
//
// Solidity: function setQuotaContractAddress(address quotaContractAddress_) returns()
func (_QuotaProxy *QuotaProxyTransactor) SetQuotaContractAddress(opts *bind.TransactOpts, quotaContractAddress_ common.Address) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "setQuotaContractAddress", quotaContractAddress_)
}

// SetQuotaContractAddress is a paid mutator transaction binding the contract method 0x2fa5c39e.
//
// Solidity: function setQuotaContractAddress(address quotaContractAddress_) returns()
func (_QuotaProxy *QuotaProxySession) SetQuotaContractAddress(quotaContractAddress_ common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetQuotaContractAddress(&_QuotaProxy.TransactOpts, quotaContractAddress_)
}

// SetQuotaContractAddress is a paid mutator transaction binding the contract method 0x2fa5c39e.
//
// Solidity: function setQuotaContractAddress(address quotaContractAddress_) returns()
func (_QuotaProxy *QuotaProxyTransactorSession) SetQuotaContractAddress(quotaContractAddress_ common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.SetQuotaContractAddress(&_QuotaProxy.TransactOpts, quotaContractAddress_)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_QuotaProxy *QuotaProxyTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_QuotaProxy *QuotaProxySession) Stake() (*types.Transaction, error) {
	return _QuotaProxy.Contract.Stake(&_QuotaProxy.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_QuotaProxy *QuotaProxyTransactorSession) Stake() (*types.Transaction, error) {
	return _QuotaProxy.Contract.Stake(&_QuotaProxy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuotaProxy *QuotaProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuotaProxy *QuotaProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.TransferOwnership(&_QuotaProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuotaProxy *QuotaProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuotaProxy.Contract.TransferOwnership(&_QuotaProxy.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_QuotaProxy *QuotaProxyTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_QuotaProxy *QuotaProxySession) Unstake() (*types.Transaction, error) {
	return _QuotaProxy.Contract.Unstake(&_QuotaProxy.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_QuotaProxy *QuotaProxyTransactorSession) Unstake() (*types.Transaction, error) {
	return _QuotaProxy.Contract.Unstake(&_QuotaProxy.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_QuotaProxy *QuotaProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _QuotaProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_QuotaProxy *QuotaProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _QuotaProxy.Contract.UpgradeToAndCall(&_QuotaProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_QuotaProxy *QuotaProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _QuotaProxy.Contract.UpgradeToAndCall(&_QuotaProxy.TransactOpts, newImplementation, data)
}

// QuotaProxyDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the QuotaProxy contract.
type QuotaProxyDelegatedIterator struct {
	Event *QuotaProxyDelegated // Event containing the contract specifics and raw log

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
func (it *QuotaProxyDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyDelegated)
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
		it.Event = new(QuotaProxyDelegated)
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
func (it *QuotaProxyDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyDelegated represents a Delegated event raised by the QuotaProxy contract.
type QuotaProxyDelegated struct {
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed delegator, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address) (*QuotaProxyDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "Delegated", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyDelegatedIterator{contract: _QuotaProxy.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed delegator, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *QuotaProxyDelegated, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "Delegated", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyDelegated)
				if err := _QuotaProxy.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x83b3f5ce88736f0128f880f5cac19836da52ea5c5ca7704c7b38f3b06fffd7ab.
//
// Solidity: event Delegated(address indexed delegator, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) ParseDelegated(log types.Log) (*QuotaProxyDelegated, error) {
	event := new(QuotaProxyDelegated)
	if err := _QuotaProxy.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyFeeRefundBlockCountUpdatedIterator is returned from FilterFeeRefundBlockCountUpdated and is used to iterate over the raw logs and unpacked data for FeeRefundBlockCountUpdated events raised by the QuotaProxy contract.
type QuotaProxyFeeRefundBlockCountUpdatedIterator struct {
	Event *QuotaProxyFeeRefundBlockCountUpdated // Event containing the contract specifics and raw log

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
func (it *QuotaProxyFeeRefundBlockCountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyFeeRefundBlockCountUpdated)
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
		it.Event = new(QuotaProxyFeeRefundBlockCountUpdated)
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
func (it *QuotaProxyFeeRefundBlockCountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyFeeRefundBlockCountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyFeeRefundBlockCountUpdated represents a FeeRefundBlockCountUpdated event raised by the QuotaProxy contract.
type QuotaProxyFeeRefundBlockCountUpdated struct {
	NewFeeRefundBlockCount uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFeeRefundBlockCountUpdated is a free log retrieval operation binding the contract event 0x80feeeebed629d55c83e2c34697608c5d4ac58116124ed3951595e5f0d5340ca.
//
// Solidity: event FeeRefundBlockCountUpdated(uint16 newFeeRefundBlockCount)
func (_QuotaProxy *QuotaProxyFilterer) FilterFeeRefundBlockCountUpdated(opts *bind.FilterOpts) (*QuotaProxyFeeRefundBlockCountUpdatedIterator, error) {

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "FeeRefundBlockCountUpdated")
	if err != nil {
		return nil, err
	}
	return &QuotaProxyFeeRefundBlockCountUpdatedIterator{contract: _QuotaProxy.contract, event: "FeeRefundBlockCountUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRefundBlockCountUpdated is a free log subscription operation binding the contract event 0x80feeeebed629d55c83e2c34697608c5d4ac58116124ed3951595e5f0d5340ca.
//
// Solidity: event FeeRefundBlockCountUpdated(uint16 newFeeRefundBlockCount)
func (_QuotaProxy *QuotaProxyFilterer) WatchFeeRefundBlockCountUpdated(opts *bind.WatchOpts, sink chan<- *QuotaProxyFeeRefundBlockCountUpdated) (event.Subscription, error) {

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "FeeRefundBlockCountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyFeeRefundBlockCountUpdated)
				if err := _QuotaProxy.contract.UnpackLog(event, "FeeRefundBlockCountUpdated", log); err != nil {
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

// ParseFeeRefundBlockCountUpdated is a log parse operation binding the contract event 0x80feeeebed629d55c83e2c34697608c5d4ac58116124ed3951595e5f0d5340ca.
//
// Solidity: event FeeRefundBlockCountUpdated(uint16 newFeeRefundBlockCount)
func (_QuotaProxy *QuotaProxyFilterer) ParseFeeRefundBlockCountUpdated(log types.Log) (*QuotaProxyFeeRefundBlockCountUpdated, error) {
	event := new(QuotaProxyFeeRefundBlockCountUpdated)
	if err := _QuotaProxy.contract.UnpackLog(event, "FeeRefundBlockCountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the QuotaProxy contract.
type QuotaProxyInitializedIterator struct {
	Event *QuotaProxyInitialized // Event containing the contract specifics and raw log

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
func (it *QuotaProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyInitialized)
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
		it.Event = new(QuotaProxyInitialized)
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
func (it *QuotaProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyInitialized represents a Initialized event raised by the QuotaProxy contract.
type QuotaProxyInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_QuotaProxy *QuotaProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*QuotaProxyInitializedIterator, error) {

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QuotaProxyInitializedIterator{contract: _QuotaProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_QuotaProxy *QuotaProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QuotaProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyInitialized)
				if err := _QuotaProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_QuotaProxy *QuotaProxyFilterer) ParseInitialized(log types.Log) (*QuotaProxyInitialized, error) {
	event := new(QuotaProxyInitialized)
	if err := _QuotaProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyMinStakeUpdatedIterator is returned from FilterMinStakeUpdated and is used to iterate over the raw logs and unpacked data for MinStakeUpdated events raised by the QuotaProxy contract.
type QuotaProxyMinStakeUpdatedIterator struct {
	Event *QuotaProxyMinStakeUpdated // Event containing the contract specifics and raw log

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
func (it *QuotaProxyMinStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyMinStakeUpdated)
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
		it.Event = new(QuotaProxyMinStakeUpdated)
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
func (it *QuotaProxyMinStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyMinStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyMinStakeUpdated represents a MinStakeUpdated event raised by the QuotaProxy contract.
type QuotaProxyMinStakeUpdated struct {
	NewMinStake *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinStakeUpdated is a free log retrieval operation binding the contract event 0x47ab46f2c8d4258304a2f5551c1cbdb6981be49631365d1ba7191288a73f39ef.
//
// Solidity: event MinStakeUpdated(uint256 newMinStake)
func (_QuotaProxy *QuotaProxyFilterer) FilterMinStakeUpdated(opts *bind.FilterOpts) (*QuotaProxyMinStakeUpdatedIterator, error) {

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "MinStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &QuotaProxyMinStakeUpdatedIterator{contract: _QuotaProxy.contract, event: "MinStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinStakeUpdated is a free log subscription operation binding the contract event 0x47ab46f2c8d4258304a2f5551c1cbdb6981be49631365d1ba7191288a73f39ef.
//
// Solidity: event MinStakeUpdated(uint256 newMinStake)
func (_QuotaProxy *QuotaProxyFilterer) WatchMinStakeUpdated(opts *bind.WatchOpts, sink chan<- *QuotaProxyMinStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "MinStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyMinStakeUpdated)
				if err := _QuotaProxy.contract.UnpackLog(event, "MinStakeUpdated", log); err != nil {
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

// ParseMinStakeUpdated is a log parse operation binding the contract event 0x47ab46f2c8d4258304a2f5551c1cbdb6981be49631365d1ba7191288a73f39ef.
//
// Solidity: event MinStakeUpdated(uint256 newMinStake)
func (_QuotaProxy *QuotaProxyFilterer) ParseMinStakeUpdated(log types.Log) (*QuotaProxyMinStakeUpdated, error) {
	event := new(QuotaProxyMinStakeUpdated)
	if err := _QuotaProxy.contract.UnpackLog(event, "MinStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuotaProxy contract.
type QuotaProxyOwnershipTransferredIterator struct {
	Event *QuotaProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuotaProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyOwnershipTransferred)
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
		it.Event = new(QuotaProxyOwnershipTransferred)
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
func (it *QuotaProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyOwnershipTransferred represents a OwnershipTransferred event raised by the QuotaProxy contract.
type QuotaProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuotaProxy *QuotaProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuotaProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyOwnershipTransferredIterator{contract: _QuotaProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuotaProxy *QuotaProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuotaProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyOwnershipTransferred)
				if err := _QuotaProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuotaProxy *QuotaProxyFilterer) ParseOwnershipTransferred(log types.Log) (*QuotaProxyOwnershipTransferred, error) {
	event := new(QuotaProxyOwnershipTransferred)
	if err := _QuotaProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the QuotaProxy contract.
type QuotaProxyUndelegatedIterator struct {
	Event *QuotaProxyUndelegated // Event containing the contract specifics and raw log

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
func (it *QuotaProxyUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyUndelegated)
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
		it.Event = new(QuotaProxyUndelegated)
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
func (it *QuotaProxyUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyUndelegated represents a Undelegated event raised by the QuotaProxy contract.
type QuotaProxyUndelegated struct {
	Delegator common.Address
	WrID      *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xb09c0be2bd2dc9e2a544469e3ab9922421f53f6546f0a032ab6f952cb45312a8.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed wrID, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, wrID []*big.Int) (*QuotaProxyUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "Undelegated", delegatorRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyUndelegatedIterator{contract: _QuotaProxy.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xb09c0be2bd2dc9e2a544469e3ab9922421f53f6546f0a032ab6f952cb45312a8.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed wrID, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *QuotaProxyUndelegated, delegator []common.Address, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "Undelegated", delegatorRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyUndelegated)
				if err := _QuotaProxy.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0xb09c0be2bd2dc9e2a544469e3ab9922421f53f6546f0a032ab6f952cb45312a8.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed wrID, uint256 amount)
func (_QuotaProxy *QuotaProxyFilterer) ParseUndelegated(log types.Log) (*QuotaProxyUndelegated, error) {
	event := new(QuotaProxyUndelegated)
	if err := _QuotaProxy.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuotaProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the QuotaProxy contract.
type QuotaProxyUpgradedIterator struct {
	Event *QuotaProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *QuotaProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuotaProxyUpgraded)
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
		it.Event = new(QuotaProxyUpgraded)
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
func (it *QuotaProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuotaProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuotaProxyUpgraded represents a Upgraded event raised by the QuotaProxy contract.
type QuotaProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_QuotaProxy *QuotaProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*QuotaProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _QuotaProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &QuotaProxyUpgradedIterator{contract: _QuotaProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_QuotaProxy *QuotaProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *QuotaProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _QuotaProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuotaProxyUpgraded)
				if err := _QuotaProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_QuotaProxy *QuotaProxyFilterer) ParseUpgraded(log types.Log) (*QuotaProxyUpgraded, error) {
	event := new(QuotaProxyUpgraded)
	if err := _QuotaProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
