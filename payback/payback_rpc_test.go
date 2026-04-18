package payback

import (
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// rpcStubStore is a Store that returns rules with Podgorica inactive so that
// GetAvailablePaybackByAddressRPC returns zero before ever touching the EVM.
// This lets us exercise the RPC entrypoint's preconditions without setting up
// a full EVM + deployed contracts.
type rpcStubStore struct {
	stubStore
	epoch idx.Epoch
	rules opera.Rules
}

func (s *rpcStubStore) GetRules() opera.Rules                                 { return s.rules }
func (s *rpcStubStore) GetCurrentEpoch() idx.Epoch                            { return s.epoch }
func (s *rpcStubStore) GetHistoryEpochState(idx.Epoch) *iblockproc.EpochState { return nil }

// TestGetAvailablePaybackByAddressRPC_ZeroAddress ensures the RPC entrypoint
// returns zero for the zero address without touching the EVM, mirroring the
// behavior of the block-processing path.
func TestGetAvailablePaybackByAddressRPC_ZeroAddress(t *testing.T) {
	store := &rpcStubStore{
		epoch: idx.Epoch(5),
		rules: opera.FakeNetRules(),
	}
	pc, err := NewPaybackCache(store, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}

	got, err := pc.GetAvailablePaybackByAddressRPC(context.Background(), common.Address{}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || got.Sign() != 0 {
		t.Fatalf("expected zero for zero address, got %v", got)
	}
}

// TestGetAvailablePaybackByAddressRPC_PodgoricaInactive verifies the RPC path
// short-circuits to zero when the Podgorica upgrade is not active on the
// network, without needing an EVM.
func TestGetAvailablePaybackByAddressRPC_PodgoricaInactive(t *testing.T) {
	rules := opera.FakeNetRules()
	rules.Upgrades.Podgorica = false
	store := &rpcStubStore{
		epoch: idx.Epoch(5),
		rules: rules,
	}
	pc, err := NewPaybackCache(store, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}

	got, err := pc.GetAvailablePaybackByAddressRPC(context.Background(), common.HexToAddress("0x1234"), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || got.Sign() != 0 {
		t.Fatalf("expected zero when Podgorica inactive, got %v", got)
	}
}

// TestGetAvailablePaybackByAddressRPC_DoesNotMutateBlkCtx ensures the RPC path
// never writes to blkCtx. Concurrent block processing relies on blkCtx being
// owned exclusively by the block-processing goroutine; an RPC that mutates it
// would corrupt in-flight block computations.
func TestGetAvailablePaybackByAddressRPC_DoesNotMutateBlkCtx(t *testing.T) {
	rules := opera.FakeNetRules()
	rules.Upgrades.Podgorica = true
	store := &rpcStubStore{
		epoch: idx.Epoch(5),
		rules: rules,
	}
	pc, err := NewPaybackCache(store, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}

	// Simulate block processing in flight: an orthogonal call installs a
	// blkCtx with sentinel values that the RPC path must not touch.
	pc.PrepareForBlock(idx.Epoch(5), rules, time.Unix(1700000000, 0))
	defer pc.FinishBlock()
	pc.mu.RLock()
	before := *pc.blkCtx
	pc.mu.RUnlock()

	// Invoke RPC path with a nil EVM and an unknown address. The zero-address
	// shortcut is not taken because the address is non-zero; instead the
	// method must reach the EVM-call path. We supply a nil EVM which will
	// cause the static calls to error out — that is fine, the invariant we
	// care about is "no writes to blkCtx".
	//
	// Note: a nil EVM will panic inside StaticCall. To avoid that, we wrap
	// the call; the test asserts the wrapper sees no mutation regardless.
	func() {
		defer func() { _ = recover() }()
		_, _ = pc.GetAvailablePaybackByAddressRPC(context.Background(), common.HexToAddress("0xABCDEF"), (*vm.EVM)(nil))
	}()

	pc.mu.RLock()
	after := *pc.blkCtx
	pc.mu.RUnlock()

	if before.baseRewardPerSecond != nil || after.baseRewardPerSecond != nil {
		t.Fatalf("RPC path must not write baseRewardPerSecond to blkCtx (before=%v, after=%v)",
			before.baseRewardPerSecond, after.baseRewardPerSecond)
	}
	if before.prevEpochStateLoaded || after.prevEpochStateLoaded {
		t.Fatalf("RPC path must not set prevEpochStateLoaded on blkCtx (before=%v, after=%v)",
			before.prevEpochStateLoaded, after.prevEpochStateLoaded)
	}
}

// TestGetAvailablePaybackByAddressRPC_DoesNotMutateStakesMap verifies the RPC
// path never inserts an empty map entry into StakesMap. The production path
// auto-creates pc.StakesMap[currentEpoch] as a side effect of getPaybackData;
// RPC traffic must not slowly bloat StakesMap with empty entries for every
// address that has ever been queried.
func TestGetAvailablePaybackByAddressRPC_DoesNotMutateStakesMap(t *testing.T) {
	rules := opera.FakeNetRules()
	rules.Upgrades.Podgorica = true
	store := &rpcStubStore{
		epoch: idx.Epoch(5),
		rules: rules,
	}
	pc, err := NewPaybackCache(store, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}

	before := len(pc.StakesMap)

	func() {
		defer func() { _ = recover() }()
		_, _ = pc.GetAvailablePaybackByAddressRPC(context.Background(), common.HexToAddress("0xABCDEF"), (*vm.EVM)(nil))
	}()

	if got := len(pc.StakesMap); got != before {
		t.Fatalf("RPC path must not insert into StakesMap (before=%d, after=%d)", before, got)
	}
}

// TestGetAvailablePaybackByAddressRPC_ConcurrentWithBlockProc runs the RPC path
// concurrently with repeated PrepareForBlock/FinishBlock cycles. Under -race,
// any shared-state mutation by the RPC path that races with block processing
// will be flagged.
func TestGetAvailablePaybackByAddressRPC_ConcurrentWithBlockProc(t *testing.T) {
	rules := opera.FakeNetRules()
	rules.Upgrades.Podgorica = false // keep it cheap; short-circuits before EVM
	store := &rpcStubStore{
		epoch: idx.Epoch(5),
		rules: rules,
	}
	pc, err := NewPaybackCache(store, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}

	const duration = 50 * time.Millisecond
	stop := time.Now().Add(duration)

	var wg sync.WaitGroup

	// Block processing simulator.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for time.Now().Before(stop) {
			pc.PrepareForBlock(idx.Epoch(5), rules, time.Now())
			pc.FinishBlock()
		}
	}()

	// RPC callers.
	const rpcGoroutines = 8
	for g := 0; g < rpcGoroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			addr := common.BigToAddress(big.NewInt(int64(g + 1)))
			for time.Now().Before(stop) {
				if _, err := pc.GetAvailablePaybackByAddressRPC(context.Background(), addr, nil); err != nil {
					t.Errorf("rpc call errored: %v", err)
					return
				}
			}
		}(g)
	}

	wg.Wait()
}
