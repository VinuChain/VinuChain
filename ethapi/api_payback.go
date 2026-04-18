package ethapi

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// Payback RPC concurrency caps.
//
// vc_getPaybackBalance triggers up to five EVM StaticCalls (≈5 × 100k gas)
// against the SFC and payback-proxy contracts: getStake, minStake,
// baseRewardPerSecond, totalStake(SFC), totalStake(Quota). Under RPC load
// this is a non-trivial CPU cost, so we cap concurrent execution across the
// whole process.
//
// paybackConcurrentCap is the maximum number of simultaneous invocations.
// Tuning target: absorb short bursts (< 50ms per call on a modern CPU) without
// starving the RPC worker pool, while providing a hard ceiling against a
// bot-driven flood. 8 is a conservative middle point for an 8-vCPU node
// running alongside block processing.
//
// paybackAcquireTimeout is how long a call will wait for a permit before
// returning the rate-limit error. Kept short so RPC clients receive fast
// feedback instead of piling up in queues.
const (
	paybackConcurrentCap      = 8
	paybackAcquireTimeout     = 2 * time.Second
	paybackRateLimitErrorCode = -32005
)

// paybackSem is a process-global semaphore for vc_getPaybackBalance.
var paybackSem = make(chan struct{}, paybackConcurrentCap)

// paybackLimitError implements rpc.Error so the RPC server reports the
// configured error code to clients.
type paybackLimitError struct{ msg string }

func (e *paybackLimitError) Error() string  { return e.msg }
func (e *paybackLimitError) ErrorCode() int { return paybackRateLimitErrorCode }

// paybackBackend is the subset of Backend needed by PublicPaybackAPI. Keeping
// this narrow makes tests cheap.
type paybackBackend interface {
	GetPaybackBalance(ctx context.Context, addr common.Address, blockNrOrHash *rpc.BlockNumberOrHash) (*big.Int, error)
}

// PublicPaybackAPI exposes VinuChain-specific payback queries under the vc
// namespace. It is a thin wrapper: the semaphore cap is the only behavior
// added on top of the backend.
type PublicPaybackAPI struct {
	b paybackBackend
}

// NewPublicPaybackAPI constructs the API.
func NewPublicPaybackAPI(b paybackBackend) *PublicPaybackAPI {
	return &PublicPaybackAPI{b: b}
}

// GetPaybackBalance returns the available payback (fee-refund quota) for the
// given address at the requested block (defaults to "latest"). Returns "0x0"
// for unknown addresses, for networks where Podgorica is not active, or for
// addresses staking below the minimum.
//
// The method is rate-limited process-wide: if paybackConcurrentCap calls are
// already in flight, additional callers wait up to paybackAcquireTimeout for
// a permit; beyond that they receive an rpc error with code -32005.
func (s *PublicPaybackAPI) GetPaybackBalance(ctx context.Context, addr common.Address, blockNrOrHash *rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	timer := time.NewTimer(paybackAcquireTimeout)
	defer timer.Stop()

	select {
	case paybackSem <- struct{}{}:
		defer func() { <-paybackSem }()
	case <-timer.C:
		return nil, &paybackLimitError{msg: "payback RPC concurrent limit exceeded"}
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	v, err := s.b.GetPaybackBalance(ctx, addr, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return (*hexutil.Big)(new(big.Int)), nil
	}
	return (*hexutil.Big)(v), nil
}
