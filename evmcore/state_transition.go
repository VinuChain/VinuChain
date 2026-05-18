// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package evmcore

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var emptyCodeHash = crypto.Keccak256Hash(nil)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay gas
3) Create a new state object if the recipient is \0*32
4) Value transfer
== If contract creation ==

	4a) Attempt to run transaction data
	4b) If valid, use result as code for the new state object

== end ==
5) Run Script section
6) Derive new state root
*/
type StateTransition struct {
	gp             *GasPool
	msg            Message
	gas            uint64
	gasPrice       *big.Int
	initialGas     uint64
	value          *big.Int
	data           []byte
	state          vm.StateDB
	evm            *vm.EVM
	availableQuota *big.Int
	feeRefund      *big.Int
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	To() *common.Address

	GasPrice() *big.Int
	GasFeeCap() *big.Int
	GasTipCap() *big.Int
	Gas() uint64
	Value() *big.Int

	Nonce() uint64
	IsFake() bool
	Data() []byte
	AccessList() types.AccessList
	SetCodeAuthorizations() []types.SetCodeAuthorization
}

// ExecutionResult includes all output after executing given evm
// message no matter the execution itself is successful or not.
type ExecutionResult struct {
	FeeRefund  *big.Int // Refund for the fee of the transaction
	UsedGas    uint64   // Total used gas but include the refunded gas
	Err        error    // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData []byte   // Returned data from evm(function result or data supplied with revert opcode)
}

// Unwrap returns the internal evm error which allows us for further
// analysis outside.
func (result *ExecutionResult) Unwrap() error {
	return result.Err
}

// Failed returns the indicator whether the execution is successful or not
func (result *ExecutionResult) Failed() bool { return result.Err != nil }

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (result *ExecutionResult) Return() []byte {
	if result.Err != nil {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (result *ExecutionResult) Revert() []byte {
	if result.Err != vm.ErrExecutionReverted {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// IntrinsicGas computes the 'intrinsic gas' for a message with the given data.
func IntrinsicGas(data []byte, accessList types.AccessList, authList []types.SetCodeAuthorization, isContractCreation bool, isHomestead, isEIP2028, isEIP3860 bool) (uint64, error) {
	// Set the starting gas for the raw transaction
	var gas uint64
	if isContractCreation && isHomestead {
		gas = params.TxGasContractCreation
	} else {
		gas = params.TxGas
	}
	// Bump the required gas by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		nonZeroGas := params.TxDataNonZeroGasFrontier
		if isEIP2028 {
			nonZeroGas = params.TxDataNonZeroGasEIP2028
		}
		if (math.MaxUint64-gas)/nonZeroGas < nz {
			return 0, ErrGasUintOverflow
		}
		gas += nz * nonZeroGas

		z := uint64(len(data)) - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return 0, ErrGasUintOverflow
		}
		gas += z * params.TxDataZeroGas

		if isContractCreation && isEIP3860 {
			lenWords := (uint64(len(data)) + 31) / 32
			if (math.MaxUint64-gas)/params.InitCodeWordGas < lenWords {
				return 0, ErrGasUintOverflow
			}
			gas += lenWords * params.InitCodeWordGas
		}
	}
	if accessList != nil {
		gas += uint64(len(accessList)) * params.TxAccessListAddressGas
		gas += uint64(accessList.StorageKeys()) * params.TxAccessListStorageKeyGas
	}
	if authList != nil {
		if (math.MaxUint64-gas)/params.CallNewAccountGas < uint64(len(authList)) {
			return 0, ErrGasUintOverflow
		}
		gas += uint64(len(authList)) * params.CallNewAccountGas
	}
	return gas, nil
}

// NewStateTransition initialises and returns a new state transition object.
//
// The congestion guard in refundGas reads its floor from evm.Context.BaseFeeFloor;
// contexts built for read-only simulations leave it nil, which disables the guard.
func NewStateTransition(evm *vm.EVM, msg Message, gp *GasPool, availableQuota *big.Int) *StateTransition {
	return &StateTransition{
		gp:             gp,
		evm:            evm,
		msg:            msg,
		gasPrice:       msg.GasPrice(),
		value:          msg.Value(),
		data:           msg.Data(),
		state:          evm.StateDB,
		availableQuota: availableQuota,
		feeRefund:      big.NewInt(0),
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg Message, gp *GasPool, availableQuota *big.Int) (*ExecutionResult, error) {
	res, err := NewStateTransition(evm, msg, gp, availableQuota).TransitionDb()
	if err != nil {
		log.Debug("Tx skipped", "err", err)
	}
	return res, err
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}

func (st *StateTransition) buyGas() error {
	mgval := new(big.Int).SetUint64(st.msg.Gas())
	mgval = mgval.Mul(mgval, st.gasPrice)
	// Note: Opera doesn't need to check against gasFeeCap instead of gasPrice, as it's too aggressive in the asynchronous environment
	if have, want := st.state.GetBalance(st.msg.From()), mgval; have.Cmp(want) < 0 {
		return fmt.Errorf("%w: address %v have %v want %v", ErrInsufficientFunds, st.msg.From().Hex(), have, want)
	}
	if err := st.gp.SubGas(st.msg.Gas()); err != nil {
		return err
	}
	st.gas += st.msg.Gas()

	st.initialGas = st.msg.Gas()
	st.state.SubBalance(st.msg.From(), mgval)
	return nil
}

func (st *StateTransition) preCheck() error {
	// Only check transactions that are not fake
	if !st.msg.IsFake() {
		// Make sure this transaction's nonce is correct.
		stNonce := st.state.GetNonce(st.msg.From())
		if msgNonce := st.msg.Nonce(); stNonce < msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooHigh,
				st.msg.From().Hex(), msgNonce, stNonce)
		} else if stNonce > msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooLow,
				st.msg.From().Hex(), msgNonce, stNonce)
		}
		// Make sure the sender is an EOA. EIP-7702 delegation designators are
		// allowed after Prague because they represent delegated EOAs, not
		// deployed contracts.
		if codeHash := st.state.GetCodeHash(st.msg.From()); codeHash != emptyCodeHash && codeHash != (common.Hash{}) {
			prague := st.evm.ChainConfig().IsPrague(st.evm.Context.BlockNumber)
			if !prague {
				return fmt.Errorf("%w: address %v, codehash: %s", ErrSenderNoEOA,
					st.msg.From().Hex(), codeHash)
			}
			if _, ok := types.ParseDelegation(st.state.GetCode(st.msg.From())); !ok {
				return fmt.Errorf("%w: address %v, codehash: %s", ErrSenderNoEOA,
					st.msg.From().Hex(), codeHash)
			}
		}
	}
	if authList := st.msg.SetCodeAuthorizations(); authList != nil {
		if err := types.ValidateSetCodeAuthorizations(authList); err != nil {
			return err
		}
		if st.msg.To() == nil {
			return fmt.Errorf("%w: address %v", ErrSetCodeTxCreate, st.msg.From().Hex())
		}
		if len(authList) == 0 {
			return fmt.Errorf("%w: address %v", ErrEmptyAuthList, st.msg.From().Hex())
		}
	}
	// Note: Opera doesn't need to check gasFeeCap >= BaseFee, because it's already checked by epochcheck
	return st.buyGas()
}

func (st *StateTransition) internal() bool {
	zeroAddr := common.Address{}
	return st.msg.From() == zeroAddr
}

// TransitionDb will transition the state by applying the current message and
// returning the evm execution result with following fields.
//
//   - used gas:
//     total gas used (including gas being refunded)
//   - returndata:
//     the returned data from evm
//   - concrete execution error:
//     various **EVM** error which aborts the execution,
//     e.g. ErrOutOfGas, ErrExecutionReverted
//
// However if any consensus issue encountered, return the error directly with
// nil evm execution result.
func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {
	// First check this message satisfies all consensus rules before
	// applying the message. The rules include these clauses
	//
	// 1. the nonce of the message caller is correct
	// 2. caller has enough balance to cover transaction fee(gaslimit * gasprice)
	// 3. the amount of gas required is available in the block
	// 4. the purchased gas is enough to cover intrinsic usage
	// 5. there is no overflow when calculating intrinsic gas

	// Note: insufficient balance for **topmost** call isn't a consensus error in Opera, unlike Ethereum
	// Such transaction will revert and consume sender's gas

	msg := st.msg
	sender := vm.AccountRef(msg.From())
	contractCreation := msg.To() == nil

	homestead := st.evm.ChainConfig().IsHomestead(st.evm.Context.BlockNumber)
	istanbul := st.evm.ChainConfig().IsIstanbul(st.evm.Context.BlockNumber)
	london := st.evm.ChainConfig().IsLondon(st.evm.Context.BlockNumber)
	shanghai := st.evm.ChainConfig().IsShanghai(st.evm.Context.BlockNumber)
	prague := st.evm.ChainConfig().IsPrague(st.evm.Context.BlockNumber)
	authList := msg.SetCodeAuthorizations()

	if authList != nil && !prague {
		return nil, ErrTxTypeNotSupported
	}

	// Check clauses 4-5 before buying gas; skipped invalid transactions must
	// not debit balances or block gas.
	gas, err := IntrinsicGas(st.data, st.msg.AccessList(), authList, contractCreation, homestead, istanbul, shanghai)
	if err != nil {
		return nil, err
	}
	if shanghai && contractCreation && len(st.data) > params.MaxInitCodeSize {
		return nil, fmt.Errorf("%w: code size %v limit %v", ErrMaxInitCodeSizeExceeded, len(st.data), params.MaxInitCodeSize)
	}
	if msg.Gas() < gas {
		return nil, fmt.Errorf("%w: have %d, want %d", ErrIntrinsicGas, msg.Gas(), gas)
	}

	// Buy gas only after fork-dependent transaction validation succeeds.
	if err := st.preCheck(); err != nil {
		return nil, err
	}
	st.gas -= gas

	// Set up the initial access list.
	if rules := st.evm.ChainConfig().Rules(st.evm.Context.BlockNumber); rules.IsBerlin {
		st.state.PrepareAccessList(msg.From(), msg.To(), vm.ActivePrecompiles(rules), msg.AccessList())
		if rules.IsShanghai {
			st.state.AddAddressToAccessList(st.evm.Context.Coinbase)
		}
	}

	var (
		ret   []byte
		vmerr error // vm errors do not effect consensus and are therefore not assigned to err
	)
	if contractCreation {
		ret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)
		if authList != nil {
			for _, auth := range authList {
				_ = st.applyAuthorization(&auth)
			}
		}
		if prague {
			if addr, ok := types.ParseDelegation(st.state.GetCode(*msg.To())); ok {
				st.state.AddAddressToAccessList(addr)
			}
		}
		ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)
	}
	// Penalize 10% of unused gas to discourage over-estimation.
	if !st.internal() {
		st.gas -= st.gas / 10
	}

	if !london {
		// Before EIP-3529: refunds were capped to gasUsed / 2
		st.refundGas(params.RefundQuotient)
	} else {
		// After EIP-3529: refunds are capped to gasUsed / 5
		st.refundGas(params.RefundQuotientEIP3529)
	}

	return &ExecutionResult{
		FeeRefund:  st.feeRefund,
		UsedGas:    st.gasUsed(),
		Err:        vmerr,
		ReturnData: ret,
	}, nil
}

// validateAuthorization validates an EIP-7702 authorization against the state.
func (st *StateTransition) validateAuthorization(auth *types.SetCodeAuthorization) (common.Address, error) {
	chainID := st.evm.ChainConfig().ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	if auth.ChainID != nil && auth.ChainID.Sign() != 0 && auth.ChainID.Cmp(chainID) != 0 {
		return common.Address{}, ErrAuthorizationWrongChainID
	}
	if auth.Nonce == math.MaxUint64 {
		return common.Address{}, ErrAuthorizationNonceOverflow
	}
	authority, err := auth.Authority()
	if err != nil {
		return common.Address{}, fmt.Errorf("%w: %v", ErrAuthorizationInvalidSignature, err)
	}
	st.state.AddAddressToAccessList(authority)
	if code := st.state.GetCode(authority); len(code) != 0 {
		if _, ok := types.ParseDelegation(code); !ok {
			return authority, ErrAuthorizationDestinationHasCode
		}
	}
	if have := st.state.GetNonce(authority); have != auth.Nonce {
		return authority, ErrAuthorizationNonceMismatch
	}
	return authority, nil
}

// applyAuthorization applies an EIP-7702 code delegation to the state.
func (st *StateTransition) applyAuthorization(auth *types.SetCodeAuthorization) error {
	authority, err := st.validateAuthorization(auth)
	if err != nil {
		return err
	}
	if st.state.Exist(authority) {
		st.state.AddRefund(params.CallNewAccountGas - params.TxAuthTupleGas)
	}
	st.state.SetNonce(authority, auth.Nonce+1)
	if auth.Address == (common.Address{}) {
		st.state.SetCode(authority, nil)
	} else {
		st.state.SetCode(authority, types.AddressToDelegation(auth.Address))
	}
	return nil
}

func (st *StateTransition) refundGas(refundQuotient uint64) {
	// Apply refund counter, capped to a refund quotient
	refund := min(st.gasUsed()/refundQuotient, st.state.GetRefund())
	st.gas += refund

	// Return wei for remaining gas, exchanged at the original rate.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)

	fee := new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice)
	feeRefund := big.NewInt(0)

	if st.availableQuota == nil {
		st.state.AddBalance(st.msg.From(), remaining)
		st.gp.AddGas(st.gas)
		return
	}

	// When the network is congested (base fee above the chain-configured floor),
	// suppress quota refunds so EIP-1559 fee escalation can deter spam.
	//
	// Simulation paths (eth_call, estimateGas, tracers) also populate BaseFeeFloor
	// since they fetch headers via ToEvmHeader. The guard may fire during simulation,
	// but those callers pass availableQuota = 0, so the refund branch is already a
	// no-op and the guard's effect is not observable. If a future simulation passes
	// a non-zero quota, it will hit this guard under congestion — that's the correct
	// behavior, not a bug.
	baseFee, floor := st.evm.Context.BaseFee, st.evm.Context.BaseFeeFloor
	if floor != nil && baseFee != nil && baseFee.Cmp(floor) > 0 {
		st.state.AddBalance(st.msg.From(), remaining)
		st.gp.AddGas(st.gas)
		return
	}

	if st.msg.From() != (common.Address{}) && st.availableQuota.Cmp(big.NewInt(0)) != 0 {
		log.Debug("refundGas: QuotaValue", "address", st.msg.From().String(), "available", st.availableQuota.String())
		log.Debug("refundGas: Fee", "address", st.msg.From().String(), "fee", fee.String())
	}

	if fee.Cmp(st.availableQuota) > 0 {
		feeRefund = new(big.Int).Set(st.availableQuota)
	} else {
		feeRefund = fee
	}

	if st.msg.From() != (common.Address{}) && st.availableQuota.Cmp(big.NewInt(0)) != 0 {
		log.Debug("refundGas: FeeRefund", "address", st.msg.From().String(), "feeRefund", feeRefund.String())
	}

	if feeRefund.Sign() > 0 {
		st.feeRefund = feeRefund
		remaining = remaining.Add(remaining, feeRefund)
	}

	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(st.gas)
}

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gas
}
