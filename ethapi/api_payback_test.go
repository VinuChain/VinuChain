package ethapi

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// paybackStubBackend embeds stubBackend and adds a configurable
// GetPaybackBalance implementation so we can exercise the RPC layer without
// a real EVM.
type paybackStubBackend struct {
	stubBackend
	mu        sync.Mutex
	returns   map[common.Address]*big.Int
	err       error
	enterCh   chan struct{} // notified on entry to simulate concurrent calls
	releaseCh chan struct{} // signalled when the test wants handlers to return
	calls     int
}

func newPaybackStubBackend() *paybackStubBackend {
	return &paybackStubBackend{
		returns: make(map[common.Address]*big.Int),
	}
}

func (b *paybackStubBackend) GetPaybackBalance(ctx context.Context, addr common.Address, blockNrOrHash *rpc.BlockNumberOrHash) (*big.Int, error) {
	b.mu.Lock()
	b.calls++
	enter := b.enterCh
	release := b.releaseCh
	v := b.returns[addr]
	err := b.err
	b.mu.Unlock()

	if enter != nil {
		enter <- struct{}{}
	}
	if release != nil {
		select {
		case <-release:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if err != nil {
		return nil, err
	}
	if v == nil {
		return big.NewInt(0), nil
	}
	return new(big.Int).Set(v), nil
}

// TestPublicPaybackAPI_KnownAddress verifies a non-zero payback balance is
// returned as a hexutil.Big for an address with staking.
func TestPublicPaybackAPI_KnownAddress(t *testing.T) {
	b := newPaybackStubBackend()
	addr := common.HexToAddress("0xAbCdEF")
	b.returns[addr] = big.NewInt(1_234_567)

	api := NewPublicPaybackAPI(b)
	got, err := api.GetPaybackBalance(context.Background(), addr, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("expected non-nil result")
	}
	if (*big.Int)(got).Cmp(big.NewInt(1_234_567)) != 0 {
		t.Fatalf("expected 1234567, got %v", got)
	}
}

// TestPublicPaybackAPI_UnknownAddressReturnsZero verifies that an unknown
// address yields "0x0" rather than an error.
func TestPublicPaybackAPI_UnknownAddressReturnsZero(t *testing.T) {
	b := newPaybackStubBackend()

	api := NewPublicPaybackAPI(b)
	got, err := api.GetPaybackBalance(context.Background(), common.HexToAddress("0xDEADBEEF"), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || (*big.Int)(got).Sign() != 0 {
		t.Fatalf("expected 0, got %v", got)
	}
}

// TestPublicPaybackAPI_BackendErrorPropagates verifies backend errors
// surface to the caller.
func TestPublicPaybackAPI_BackendErrorPropagates(t *testing.T) {
	b := newPaybackStubBackend()
	b.err = errors.New("state not available")

	api := NewPublicPaybackAPI(b)
	_, err := api.GetPaybackBalance(context.Background(), common.HexToAddress("0x1"), nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

// TestPublicPaybackAPI_ConcurrentLimit verifies that when the internal
// semaphore is saturated, additional callers are rejected with the
// documented rpc error code -32005.
func TestPublicPaybackAPI_ConcurrentLimit(t *testing.T) {
	b := newPaybackStubBackend()
	b.enterCh = make(chan struct{}, paybackConcurrentCap+1)
	b.releaseCh = make(chan struct{})
	api := NewPublicPaybackAPI(b)

	addr := common.HexToAddress("0xBEEF")

	// Launch paybackConcurrentCap in-flight calls that stall inside the
	// backend, each holding one permit.
	var wg sync.WaitGroup
	wg.Add(paybackConcurrentCap)
	for i := 0; i < paybackConcurrentCap; i++ {
		go func() {
			defer wg.Done()
			_, err := api.GetPaybackBalance(context.Background(), addr, nil)
			if err != nil {
				t.Errorf("in-flight call got error: %v", err)
			}
		}()
	}

	// Wait for all permits to be taken.
	for i := 0; i < paybackConcurrentCap; i++ {
		select {
		case <-b.enterCh:
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out waiting for in-flight calls to enter (%d of %d entered)", i, paybackConcurrentCap)
		}
	}

	// The (cap+1)th call must be rejected with -32005 after the acquire
	// timeout elapses. Use a tight ctx so the test does not rely on the
	// full 2s default timeout.
	ctx, cancel := context.WithTimeout(context.Background(), paybackAcquireTimeout+500*time.Millisecond)
	defer cancel()

	_, err := api.GetPaybackBalance(ctx, addr, nil)
	if err == nil {
		t.Fatal("expected rejection from cap+1th call, got nil error")
	}
	var rpcErr rpc.Error
	if !errors.As(err, &rpcErr) {
		t.Fatalf("expected rpc.Error, got %T: %v", err, err)
	}
	if rpcErr.ErrorCode() != paybackRateLimitErrorCode {
		t.Fatalf("expected code %d, got %d", paybackRateLimitErrorCode, rpcErr.ErrorCode())
	}

	// Release in-flight callers.
	close(b.releaseCh)
	wg.Wait()
}

// TestPublicPaybackAPI_ReleasesPermitOnReturn verifies that the semaphore is
// released after each call so the same goroutine can make many calls in
// series without exhausting capacity.
func TestPublicPaybackAPI_ReleasesPermitOnReturn(t *testing.T) {
	b := newPaybackStubBackend()
	api := NewPublicPaybackAPI(b)
	addr := common.HexToAddress("0xC0DE")
	for i := 0; i < paybackConcurrentCap*3; i++ {
		_, err := api.GetPaybackBalance(context.Background(), addr, nil)
		if err != nil {
			t.Fatalf("call %d failed: %v", i, err)
		}
	}
}

// TestPaybackAPI_RegisteredUnderVCNamespace verifies the payback API is
// exposed under the "vc" namespace via ethapi.GetAPIs. The test looks at the
// returned []rpc.API and asserts that at least one entry has Namespace=="vc"
// with a Service that exposes the GetPaybackBalance method.
func TestPaybackAPI_RegisteredUnderVCNamespace(t *testing.T) {
	b := newPaybackStubBackend()
	apis := GetAPIs(b)

	var foundVC bool
	var foundEth bool
	for _, a := range apis {
		if _, ok := a.Service.(*PublicPaybackAPI); !ok {
			continue
		}
		switch a.Namespace {
		case "vc":
			foundVC = true
		case "eth":
			foundEth = true
		}
	}
	if !foundVC {
		t.Fatal("PublicPaybackAPI not registered under vc namespace")
	}
	if foundEth {
		t.Fatal("PublicPaybackAPI must NOT be registered under eth namespace")
	}
}
