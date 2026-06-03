package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

type estimateGasBlockBackend struct {
	stubBackend
	resolved idx.Block
}

func (b *estimateGasBlockBackend) ResolveRpcBlockNumberOrHash(context.Context, rpc.BlockNumberOrHash) (idx.Block, error) {
	return b.resolved, nil
}

func TestVinuLatestEVMEstimateGasCap(t *testing.T) {
	cfg := &params.ChainConfig{VinuLatestEVMBlock: big.NewInt(10)}

	if got := vinuLatestEVMEstimateGasCap(cfg, big.NewInt(9), params.MaxTxGasLimit+1); got != params.MaxTxGasLimit+1 {
		t.Fatalf("pre-fork cap = %d, want unchanged", got)
	}
	if got := vinuLatestEVMEstimateGasCap(cfg, big.NewInt(10), params.MaxTxGasLimit+1); got != params.MaxTxGasLimit {
		t.Fatalf("post-fork cap = %d, want %d", got, params.MaxTxGasLimit)
	}
	if got := vinuLatestEVMEstimateGasCap(cfg, big.NewInt(10), params.TxGas); got != params.TxGas {
		t.Fatalf("below-cap gas = %d, want unchanged", got)
	}
}

func TestVinuLatestEVMEstimateGasMayNeedCap(t *testing.T) {
	cfg := &params.ChainConfig{VinuLatestEVMBlock: big.NewInt(10)}
	if !vinuLatestEVMEstimateGasMayNeedCap(cfg, params.MaxTxGasLimit+1) {
		t.Fatal("configured fork with over-cap gas should require block lookup")
	}
	if vinuLatestEVMEstimateGasMayNeedCap(cfg, params.MaxTxGasLimit) {
		t.Fatal("at-cap gas should not require block lookup")
	}
	if vinuLatestEVMEstimateGasMayNeedCap(&params.ChainConfig{}, params.MaxTxGasLimit+1) {
		t.Fatal("unconfigured fork should not require block lookup")
	}
}

func TestVinuLatestEVMEstimateGasBlockNumber(t *testing.T) {
	backend := &estimateGasBlockBackend{resolved: idx.Block(9)}
	latest := rpc.BlockNumberOrHashWithNumber(rpc.LatestBlockNumber)
	if got, err := vinuLatestEVMEstimateGasBlockNumber(context.Background(), backend, latest); err != nil || got.Cmp(big.NewInt(10)) != 0 {
		t.Fatalf("latest estimate block number = %v, err = %v, want 10", got, err)
	}

	backend.resolved = idx.Block(10)
	pending := rpc.BlockNumberOrHashWithNumber(rpc.PendingBlockNumber)
	if got, err := vinuLatestEVMEstimateGasBlockNumber(context.Background(), backend, pending); err != nil || got.Cmp(big.NewInt(11)) != 0 {
		t.Fatalf("pending estimate block number = %v, err = %v, want 11", got, err)
	}

	backend.resolved = idx.Block(9)
	exact := rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(9))
	if got, err := vinuLatestEVMEstimateGasBlockNumber(context.Background(), backend, exact); err != nil || got.Cmp(big.NewInt(9)) != 0 {
		t.Fatalf("exact estimate block number = %v, err = %v, want unchanged", got, err)
	}
}
