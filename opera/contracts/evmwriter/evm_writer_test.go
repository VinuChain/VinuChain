package evmwriter

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/opera/contracts/driver"
)

func TestSign(t *testing.T) {
	require := require.New(t)

	require.Equal([]byte{0xe3, 0x04, 0x43, 0xbc}, setBalanceMethodID)
	require.Equal([]byte{0xd6, 0xa0, 0xc7, 0xaf}, copyCodeMethodID)
	require.Equal([]byte{0x07, 0x69, 0x0b, 0x2a}, swapCodeMethodID)
	require.Equal([]byte{0x39, 0xe5, 0x03, 0xab}, setStorageMethodID)
	require.Equal([]byte{0x79, 0xbe, 0xad, 0x38}, incNonceMethodID)

}

// stubStateDB implements vm.StateDB with minimal stubs for testing the
// EvmWriter precompile. Only methods exercised by EvmWriter carry logic;
// the rest are no-ops.
type stubStateDB struct {
	balances map[common.Address]*big.Int
	nonces   map[common.Address]uint64
	code     map[common.Address][]byte
	storage  map[common.Address]map[common.Hash]common.Hash
}

func newStubStateDB() *stubStateDB {
	return &stubStateDB{
		balances: make(map[common.Address]*big.Int),
		nonces:   make(map[common.Address]uint64),
		code:     make(map[common.Address][]byte),
		storage:  make(map[common.Address]map[common.Hash]common.Hash),
	}
}

func (s *stubStateDB) GetBalance(addr common.Address) *big.Int {
	if b, ok := s.balances[addr]; ok {
		return new(big.Int).Set(b)
	}
	return new(big.Int)
}

func (s *stubStateDB) AddBalance(addr common.Address, amount *big.Int) {
	if _, ok := s.balances[addr]; !ok {
		s.balances[addr] = new(big.Int)
	}
	s.balances[addr].Add(s.balances[addr], amount)
}

func (s *stubStateDB) SubBalance(addr common.Address, amount *big.Int) {
	if _, ok := s.balances[addr]; !ok {
		s.balances[addr] = new(big.Int)
	}
	s.balances[addr].Sub(s.balances[addr], amount)
}

func (s *stubStateDB) GetNonce(addr common.Address) uint64 {
	return s.nonces[addr]
}

func (s *stubStateDB) SetNonce(addr common.Address, nonce uint64) {
	s.nonces[addr] = nonce
}

func (s *stubStateDB) GetCode(addr common.Address) []byte {
	return s.code[addr]
}

func (s *stubStateDB) SetCode(addr common.Address, code []byte) {
	s.code[addr] = code
}

func (s *stubStateDB) GetCodeSize(addr common.Address) int {
	return len(s.code[addr])
}

func (s *stubStateDB) GetCodeHash(addr common.Address) common.Hash {
	return common.Hash{}
}

func (s *stubStateDB) GetState(addr common.Address, key common.Hash) common.Hash {
	if m, ok := s.storage[addr]; ok {
		return m[key]
	}
	return common.Hash{}
}

func (s *stubStateDB) SetState(addr common.Address, key common.Hash, value common.Hash) {
	if _, ok := s.storage[addr]; !ok {
		s.storage[addr] = make(map[common.Hash]common.Hash)
	}
	s.storage[addr][key] = value
}

func (s *stubStateDB) GetCommittedState(common.Address, common.Hash) common.Hash {
	return common.Hash{}
}

func (s *stubStateDB) CreateAccount(common.Address)                          {}
func (s *stubStateDB) AddRefund(uint64)                                      {}
func (s *stubStateDB) SubRefund(uint64)                                      {}
func (s *stubStateDB) GetRefund() uint64                                     { return 0 }
func (s *stubStateDB) Suicide(common.Address) bool                           { return false }
func (s *stubStateDB) HasSuicided(common.Address) bool                       { return false }
func (s *stubStateDB) Exist(common.Address) bool                             { return true }
func (s *stubStateDB) Empty(common.Address) bool                             { return false }
func (s *stubStateDB) RevertToSnapshot(int)                                  {}
func (s *stubStateDB) Snapshot() int                                         { return 0 }
func (s *stubStateDB) AddLog(*types.Log)                                     {}
func (s *stubStateDB) AddPreimage(common.Hash, []byte)                       {}
func (s *stubStateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error {
	return nil
}
func (s *stubStateDB) PrepareAccessList(common.Address, *common.Address, []common.Address, types.AccessList) {
}
func (s *stubStateDB) AddressInAccessList(common.Address) bool                        { return false }
func (s *stubStateDB) SlotInAccessList(common.Address, common.Hash) (bool, bool)      { return false, false }
func (s *stubStateDB) AddAddressToAccessList(common.Address)                          {}
func (s *stubStateDB) AddSlotToAccessList(common.Address, common.Hash)                {}

var (
	driverAddr = driver.ContractAddress
	origin     = common.HexToAddress("0x1111111111111111111111111111111111111111")
	otherAddr  = common.HexToAddress("0x2222222222222222222222222222222222222222")
	otherAddr2 = common.HexToAddress("0x3333333333333333333333333333333333333333")
	highGas    = uint64(1e8)
)

func padAddress(addr common.Address) []byte {
	padded := make([]byte, 32)
	copy(padded[12:], addr.Bytes())
	return padded
}

func padUint256(val *big.Int) []byte {
	b := val.Bytes()
	padded := make([]byte, 32)
	copy(padded[32-len(b):], b)
	return padded
}

func buildInput(methodID []byte, args ...[]byte) []byte {
	input := make([]byte, 0, 4+len(args)*32)
	input = append(input, methodID...)
	for _, arg := range args {
		input = append(input, arg...)
	}
	return input
}

func runPrecompile(stateDB vm.StateDB, originAddr common.Address, input []byte) ([]byte, uint64, error) {
	txCtx := vm.TxContext{Origin: originAddr, GasPrice: big.NewInt(1)}
	pc := PreCompiledContract{}
	return pc.Run(stateDB, vm.BlockContext{}, txCtx, driverAddr, input, highGas)
}

// --- setBalance Origin check tests ---

func TestSetBalance_RejectsOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(setBalanceMethodID, padAddress(origin), padUint256(big.NewInt(100)))
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestSetBalance_AllowsNonOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(setBalanceMethodID, padAddress(otherAddr), padUint256(big.NewInt(100)))
	_, _, err := runPrecompile(sdb, origin, input)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(100), sdb.GetBalance(otherAddr))
}

// --- copyCode Origin check tests ---

func TestCopyCode_RejectsOriginAsTarget(t *testing.T) {
	sdb := newStubStateDB()
	sdb.code[otherAddr] = []byte{0x60, 0x00}
	input := buildInput(copyCodeMethodID, padAddress(origin), padAddress(otherAddr))
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestCopyCode_AllowsNonOriginTarget(t *testing.T) {
	sdb := newStubStateDB()
	code := []byte{0x60, 0x00}
	sdb.code[otherAddr2] = code
	input := buildInput(copyCodeMethodID, padAddress(otherAddr), padAddress(otherAddr2))
	_, _, err := runPrecompile(sdb, origin, input)
	require.NoError(t, err)
	require.Equal(t, code, sdb.GetCode(otherAddr))
}

// --- swapCode Origin check tests ---

func TestSwapCode_RejectsOriginAsFirstAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(swapCodeMethodID, padAddress(origin), padAddress(otherAddr))
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestSwapCode_RejectsOriginAsSecondAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(swapCodeMethodID, padAddress(otherAddr), padAddress(origin))
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestSwapCode_AllowsNonOriginAccounts(t *testing.T) {
	sdb := newStubStateDB()
	code1 := []byte{0x60, 0x01}
	code2 := []byte{0x60, 0x02}
	sdb.code[otherAddr] = code1
	sdb.code[otherAddr2] = code2
	input := buildInput(swapCodeMethodID, padAddress(otherAddr), padAddress(otherAddr2))
	_, _, err := runPrecompile(sdb, origin, input)
	require.NoError(t, err)
	require.Equal(t, code2, sdb.GetCode(otherAddr))
	require.Equal(t, code1, sdb.GetCode(otherAddr2))
}

// --- setStorage Origin check tests ---

func TestSetStorage_RejectsOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	key := padUint256(big.NewInt(1))
	value := padUint256(big.NewInt(42))
	input := buildInput(setStorageMethodID, padAddress(origin), key, value)
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestSetStorage_AllowsNonOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	key := padUint256(big.NewInt(1))
	value := padUint256(big.NewInt(42))
	input := buildInput(setStorageMethodID, padAddress(otherAddr), key, value)
	_, _, err := runPrecompile(sdb, origin, input)
	require.NoError(t, err)
	stored := sdb.GetState(otherAddr, common.BytesToHash(key))
	require.Equal(t, common.BytesToHash(value), stored)
}

// --- incNonce Origin check tests ---

func TestIncNonce_RejectsOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(incNonceMethodID, padAddress(origin), padUint256(big.NewInt(1)))
	_, _, err := runPrecompile(sdb, origin, input)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}

func TestIncNonce_AllowsNonOriginAccount(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(incNonceMethodID, padAddress(otherAddr), padUint256(big.NewInt(1)))
	_, _, err := runPrecompile(sdb, origin, input)
	require.NoError(t, err)
	require.Equal(t, uint64(1), sdb.GetNonce(otherAddr))
}

// --- caller check (not driver) ---

func TestRun_RejectsNonDriverCaller(t *testing.T) {
	sdb := newStubStateDB()
	input := buildInput(setBalanceMethodID, padAddress(otherAddr), padUint256(big.NewInt(100)))
	txCtx := vm.TxContext{Origin: origin, GasPrice: big.NewInt(1)}
	pc := PreCompiledContract{}
	_, _, err := pc.Run(sdb, vm.BlockContext{}, txCtx, otherAddr, input, highGas)
	require.ErrorIs(t, err, vm.ErrExecutionReverted)
}
