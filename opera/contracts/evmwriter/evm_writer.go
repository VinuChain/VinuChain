package evmwriter

import (
	"bytes"
	"math"
	"math/big"
	"strings"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera/contracts/driver"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driverauth"
	"github.com/Fantom-foundation/go-opera/opera/contracts/netinit"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
)

var (
	// ContractAddress is the EvmWriter pre-compiled contract address
	ContractAddress = common.HexToAddress("0xd100ec0000000000000000000000000000000000")
	// ContractABI is the input ABI used to generate the binding from
	ContractABI string = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"copyCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"with\",\"type\":\"address\"}],\"name\":\"swapCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"diff\",\"type\":\"uint256\"}],\"name\":\"incNonce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

var (
	setBalanceMethodID     []byte
	copyCodeMethodID       []byte
	swapCodeMethodID       []byte
	setStorageMethodID     []byte
	incNonceMethodID       []byte
	balanceWarningThreshold = new(big.Int).Exp(big.NewInt(10), big.NewInt(24), nil)
)

func init() {
	abi, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		panic(err)
	}

	for name, constID := range map[string]*[]byte{
		"setBalance": &setBalanceMethodID,
		"copyCode":   &copyCodeMethodID,
		"swapCode":   &swapCodeMethodID,
		"setStorage": &setStorageMethodID,
		"incNonce":   &incNonceMethodID,
	} {
		method, exist := abi.Methods[name]
		if !exist {
			panic("unknown EvmWriter method")
		}

		*constID = make([]byte, len(method.ID))
		copy(*constID, method.ID)
	}
}

type PreCompiledContract struct {
	// paybackProxyAddr stores the payback proxy contract address (common.Address)
	// set from Rules.Economy.QuotaCacheAddress at block processing time.
	// Accessed atomically since RPC and block processing run concurrently.
	paybackProxyAddr atomic.Value
	// elemont stores whether the Elemont upgrade is active (bool).
	// Used to gate consensus-affecting gas cost changes.
	elemont atomic.Value
}

func (c *PreCompiledContract) SetElemont(active bool) {
	c.elemont.Store(active)
}

func (c *PreCompiledContract) isElemont() bool {
	v := c.elemont.Load()
	if v == nil {
		return false
	}
	return v.(bool)
}

// SetPaybackProxyAddr updates the payback proxy address used by isSystemContract.
// Called when rules are loaded or change (epoch boundaries, governance).
func (c *PreCompiledContract) SetPaybackProxyAddr(addr common.Address) {
	c.paybackProxyAddr.Store(addr)
}

func (c *PreCompiledContract) getPaybackProxyAddr() common.Address {
	v := c.paybackProxyAddr.Load()
	if v == nil {
		return common.Address{}
	}
	return v.(common.Address)
}

func (c *PreCompiledContract) Run(stateDB vm.StateDB, _ vm.BlockContext, txCtx vm.TxContext, caller common.Address, input []byte, suppliedGas uint64) ([]byte, uint64, error) {
	if caller != driver.ContractAddress {
		return nil, 0, vm.ErrExecutionReverted
	}
	if len(input) < 4 {
		return nil, 0, vm.ErrExecutionReverted
	}
	if bytes.Equal(input[:4], setBalanceMethodID) {
		input = input[4:]
		// setBalance
		if suppliedGas < params.CallValueTransferGas {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= params.CallValueTransferGas
		if len(input) != 64 {
			return nil, 0, vm.ErrExecutionReverted
		}

		acc := common.BytesToAddress(input[12:32])
		input = input[32:]
		value := new(big.Int).SetBytes(input[:32])

		if value.Cmp(balanceWarningThreshold) > 0 {
			log.Warn("EvmWriter setBalance: unusually large balance", "addr", acc, "value", value)
		}

		if acc == txCtx.Origin {
			// Origin balance shouldn't decrease during his transaction
			return nil, 0, vm.ErrExecutionReverted
		}

		balance := stateDB.GetBalance(acc)
		if balance.Cmp(value) >= 0 {
			diff := new(big.Int).Sub(balance, value)
			stateDB.SubBalance(acc, diff)
		} else {
			diff := new(big.Int).Sub(value, balance)
			stateDB.AddBalance(acc, diff)
		}
	} else if bytes.Equal(input[:4], copyCodeMethodID) {
		input = input[4:]
		// copyCode
		if suppliedGas < params.CreateGas {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= params.CreateGas
		if len(input) != 64 {
			return nil, 0, vm.ErrExecutionReverted
		}

		accTo := common.BytesToAddress(input[12:32])
		input = input[32:]
		accFrom := common.BytesToAddress(input[12:32])

		if accTo == txCtx.Origin {
			return nil, 0, vm.ErrExecutionReverted
		}

		code := stateDB.GetCode(accFrom)
		if code == nil {
			code = []byte{}
		}
		perByte := params.CreateDataGas + params.MemoryGas
		if uint64(len(code)) > math.MaxUint64/perByte {
			return nil, 0, vm.ErrOutOfGas
		}
		cost := uint64(len(code)) * perByte
		if suppliedGas < cost {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= cost
		if accTo != accFrom {
			stateDB.SetCode(accTo, code)
		}
	} else if bytes.Equal(input[:4], swapCodeMethodID) {
		input = input[4:]
		// swapCode
		cost := 2 * params.CreateGas
		if suppliedGas < cost {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= cost
		if len(input) != 64 {
			return nil, 0, vm.ErrExecutionReverted
		}

		acc0 := common.BytesToAddress(input[12:32])
		input = input[32:]
		acc1 := common.BytesToAddress(input[12:32])

		if acc0 == txCtx.Origin || acc1 == txCtx.Origin {
			return nil, 0, vm.ErrExecutionReverted
		}
		if c.isSystemContract(acc0) || c.isSystemContract(acc1) {
			return nil, 0, vm.ErrExecutionReverted
		}
		code0 := stateDB.GetCode(acc0)
		if code0 == nil {
			code0 = []byte{}
		}
		code1 := stateDB.GetCode(acc1)
		if code1 == nil {
			code1 = []byte{}
		}
		cost0 := uint64(len(code0)) * (params.CreateDataGas + params.MemoryGas)
		cost1 := uint64(len(code1)) * (params.CreateDataGas + params.MemoryGas)
		if cost0 > math.MaxUint64-cost1 {
			return nil, 0, vm.ErrOutOfGas
		}
		if c.isElemont() {
			cost = cost0 + cost1
		} else {
			cost = (cost0 + cost1) / 2 // pre-Elemont 50% discount
		}
		if suppliedGas < cost {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= cost
		if acc0 != acc1 {
			stateDB.SetCode(acc0, code1)
			stateDB.SetCode(acc1, code0)
		}
	} else if bytes.Equal(input[:4], setStorageMethodID) {
		input = input[4:]
		// setStorage
		if suppliedGas < params.SstoreSetGasEIP2200 {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= params.SstoreSetGasEIP2200
		if len(input) != 96 {
			return nil, 0, vm.ErrExecutionReverted
		}
		acc := common.BytesToAddress(input[12:32])
		input = input[32:]
		key := common.BytesToHash(input[:32])
		input = input[32:]
		value := common.BytesToHash(input[:32])

		if acc == txCtx.Origin || c.isSystemContract(acc) {
			return nil, 0, vm.ErrExecutionReverted
		}

		stateDB.SetState(acc, key, value)
	} else if bytes.Equal(input[:4], incNonceMethodID) {
		input = input[4:]
		// incNonce
		if suppliedGas < params.CallValueTransferGas {
			return nil, 0, vm.ErrOutOfGas
		}
		suppliedGas -= params.CallValueTransferGas
		if len(input) != 64 {
			return nil, 0, vm.ErrExecutionReverted
		}

		acc := common.BytesToAddress(input[12:32])
		input = input[32:]
		value := new(big.Int).SetBytes(input[:32])

		if acc == txCtx.Origin || acc == (common.Address{}) {
			// Origin nonce shouldn't change during his transaction.
			// Zero address is the internal tx sequencer — nonce inflation halts the chain.
			return nil, 0, vm.ErrExecutionReverted
		}

		if value.Cmp(common.Big256) >= 0 {
			// Don't allow large nonce increasing to prevent a nonce overflow
			return nil, 0, vm.ErrExecutionReverted
		}
		if value.Sign() <= 0 {
			return nil, 0, vm.ErrExecutionReverted
		}

		currentNonce := stateDB.GetNonce(acc)
		increment := value.Uint64()
		if currentNonce > math.MaxUint64-increment {
			return nil, 0, vm.ErrExecutionReverted
		}
		stateDB.SetNonce(acc, currentNonce+increment)
	} else {
		return nil, 0, vm.ErrExecutionReverted
	}
	return nil, suppliedGas, nil
}

func (c *PreCompiledContract) isSystemContract(addr common.Address) bool {
	proxyAddr := c.getPaybackProxyAddr()
	return addr == ContractAddress ||
		addr == driver.ContractAddress ||
		addr == driverauth.ContractAddress ||
		addr == netinit.ContractAddress ||
		addr == sfc.ContractAddress ||
		(proxyAddr != (common.Address{}) && addr == proxyAddr)
}
