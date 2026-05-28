package opera

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// rawArachnidDeployerTx is the canonical Arachnid deterministic-deployment-proxy
// transaction (Nick's method). Decoding it must yield a tx whose hash equals the
// pinned allowlist hash.
const rawArachnidDeployerTx = "f8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222"

func decodeRaw(t *testing.T, hexStr string) *types.Transaction {
	t.Helper()
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(common.FromHex(hexStr), tx); err != nil {
		t.Fatalf("failed to decode raw tx: %v", err)
	}
	return tx
}

func TestArachnidDeployerTxHashMatchesCanonical(t *testing.T) {
	tx := decodeRaw(t, rawArachnidDeployerTx)
	if tx.Protected() {
		t.Fatalf("canonical Arachnid deployer tx must be unprotected (pre-EIP-155)")
	}
	if got := tx.Hash(); got != ArachnidDeployerTxHash() {
		t.Fatalf("decoded tx hash %s != pinned allowlist hash %s", got.Hex(), ArachnidDeployerTxHash().Hex())
	}
}

func TestUnprotectedTxAllowlistedOnMainnet_AcceptsCanonicalDeployer(t *testing.T) {
	tx := decodeRaw(t, rawArachnidDeployerTx)
	if !UnprotectedTxAllowlistedOnMainnet(tx) {
		t.Fatalf("canonical Arachnid deployer tx should be allowlisted on mainnet")
	}
}

func TestUnprotectedTxAllowlistedOnMainnet_RejectsOtherUnprotectedTx(t *testing.T) {
	// A value-bearing unprotected legacy tx (the dangerous kind) must NOT be
	// allowlisted, even if crafted to look innocuous.
	to := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	other := types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: big.NewInt(1e9),
		Gas:      21000,
		To:       &to,
		Value:    big.NewInt(1e18),
		Data:     nil,
	})
	if UnprotectedTxAllowlistedOnMainnet(other) {
		t.Fatalf("a value-bearing unprotected tx must NOT be allowlisted on mainnet")
	}
}

func TestUnprotectedTxAllowlistedOnMainnet_RejectsNil(t *testing.T) {
	if UnprotectedTxAllowlistedOnMainnet(nil) {
		t.Fatalf("nil tx must not be allowlisted")
	}
}
