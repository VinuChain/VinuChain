// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paybackProxy

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRefundBlockCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quotaFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// FeeRefundBlockCount is a free data retrieval call binding the contract method 0x0fe34e68.
//
// Solidity: function feeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxyCaller) FeeRefundBlockCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "feeRefundBlockCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeRefundBlockCount is a free data retrieval call binding the contract method 0x0fe34e68.
//
// Solidity: function feeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxySession) FeeRefundBlockCount() (uint16, error) {
	return _QuotaProxy.Contract.FeeRefundBlockCount(&_QuotaProxy.CallOpts)
}

// FeeRefundBlockCount is a free data retrieval call binding the contract method 0x0fe34e68.
//
// Solidity: function feeRefundBlockCount() view returns(uint16)
func (_QuotaProxy *QuotaProxyCallerSession) FeeRefundBlockCount() (uint16, error) {
	return _QuotaProxy.Contract.FeeRefundBlockCount(&_QuotaProxy.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) GetStake(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "getStake", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxySession) GetStake(sender common.Address) (*big.Int, error) {
	return _QuotaProxy.Contract.GetStake(&_QuotaProxy.CallOpts, sender)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address sender) view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) GetStake(sender common.Address) (*big.Int, error) {
	return _QuotaProxy.Contract.GetStake(&_QuotaProxy.CallOpts, sender)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_QuotaProxy *QuotaProxySession) MinStake() (*big.Int, error) {
	return _QuotaProxy.Contract.MinStake(&_QuotaProxy.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) MinStake() (*big.Int, error) {
	return _QuotaProxy.Contract.MinStake(&_QuotaProxy.CallOpts)
}

// QuotaFactor is a free data retrieval call binding the contract method 0x976dd021.
//
// Solidity: function quotaFactor() view returns(uint256)
func (_QuotaProxy *QuotaProxyCaller) QuotaFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuotaProxy.contract.Call(opts, &out, "quotaFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaFactor is a free data retrieval call binding the contract method 0x976dd021.
//
// Solidity: function quotaFactor() view returns(uint256)
func (_QuotaProxy *QuotaProxySession) QuotaFactor() (*big.Int, error) {
	return _QuotaProxy.Contract.QuotaFactor(&_QuotaProxy.CallOpts)
}

// QuotaFactor is a free data retrieval call binding the contract method 0x976dd021.
//
// Solidity: function quotaFactor() view returns(uint256)
func (_QuotaProxy *QuotaProxyCallerSession) QuotaFactor() (*big.Int, error) {
	return _QuotaProxy.Contract.QuotaFactor(&_QuotaProxy.CallOpts)
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
