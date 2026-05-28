package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/opera"
)

type submitTxBoundaryBackend struct {
	stubBackend

	config *params.ChainConfig
	block  *evmcore.EvmBlock
	sent   *types.Transaction
}

func (b *submitTxBoundaryBackend) ChainConfig() *params.ChainConfig {
	return b.config
}

func (b *submitTxBoundaryBackend) CurrentBlock() *evmcore.EvmBlock {
	return b.block
}

func (b *submitTxBoundaryBackend) SendTx(_ context.Context, tx *types.Transaction) error {
	b.sent = tx
	return nil
}

func TestSubmitTransactionLogsSetCodeTxAtPragueBoundary(t *testing.T) {
	chainID := big.NewInt(206)
	cfg := submitTxChainConfig(chainID)
	cfg.ShanghaiBlock = big.NewInt(1)
	cfg.CancunBlock = big.NewInt(1)
	cfg.PragueBlock = big.NewInt(1)

	currentBlock := submitTxBlock(common.Big0)

	tx := signedSetCodeSubmitTx(t, chainID)
	if _, err := types.Sender(types.MakeSigner(&cfg, currentBlock.Number), tx); err != types.ErrTxTypeNotSupported {
		t.Fatalf("current-block signer error = %v, want %v", err, types.ErrTxTypeNotSupported)
	}

	backend := &submitTxBoundaryBackend{
		config: &cfg,
		block:  currentBlock,
	}
	got, err := SubmitTransaction(context.Background(), backend, tx)
	if err != nil {
		t.Fatalf("SubmitTransaction returned error after backend accepted tx: %v", err)
	}
	if got != tx.Hash() {
		t.Fatalf("SubmitTransaction hash = %s, want %s", got, tx.Hash())
	}
	if backend.sent != tx {
		t.Fatal("SubmitTransaction did not submit transaction to backend")
	}
}

func TestSubmitTransactionDoesNotReturnPostSubmitSenderRecoveryError(t *testing.T) {
	chainID := big.NewInt(206)
	cfg := submitTxChainConfig(chainID)
	cfg.PragueBlock = nil

	tx := signedSetCodeSubmitTx(t, chainID)
	if _, err := types.Sender(types.LatestSigner(&cfg), tx); err != types.ErrTxTypeNotSupported {
		t.Fatalf("latest signer error = %v, want %v", err, types.ErrTxTypeNotSupported)
	}

	backend := &submitTxBoundaryBackend{
		config: &cfg,
		block:  submitTxBlock(common.Big0),
	}
	got, err := SubmitTransaction(context.Background(), backend, tx)
	if err != nil {
		t.Fatalf("SubmitTransaction returned post-submit sender recovery error: %v", err)
	}
	if got != tx.Hash() {
		t.Fatalf("SubmitTransaction hash = %s, want %s", got, tx.Hash())
	}
	if backend.sent != tx {
		t.Fatal("SubmitTransaction did not submit transaction to backend")
	}
}

// rawArachnidDeployerTx is the canonical Arachnid deterministic-deployment-proxy
// transaction (Nick's method). Duplicated from the opera package's allowlist test
// so the ethapi call-site is pinned independently of the opera-level unit test.
const rawArachnidDeployerTx = "f8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222"

func decodeArachnidDeployerTx(t *testing.T) *types.Transaction {
	t.Helper()
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(common.FromHex(rawArachnidDeployerTx), tx); err != nil {
		t.Fatalf("decode Arachnid deployer tx: %v", err)
	}
	if tx.Hash() != opera.ArachnidDeployerTxHash() {
		t.Fatalf("decoded tx hash %s != pinned allowlist hash %s", tx.Hash(), opera.ArachnidDeployerTxHash())
	}
	return tx
}

// TestSubmitTransactionAdmitsArachnidDeployerOnlyOnMainnet pins the call-site
// scoping of the unprotected-tx carve-out: the canonical Arachnid deployer is
// admitted only when the chain is mainnet (NetworkID 207). On any other network
// the carve-out is inert and an unprotected tx is refused unless the operator
// has separately enabled AllowUnprotectedTxs (the stub backend keeps it off).
// This guards against a refactor silently widening or inverting the
// `mainnet && allowlisted` gate, which the pure opera-level test cannot catch.
func TestSubmitTransactionAdmitsArachnidDeployerOnlyOnMainnet(t *testing.T) {
	tx := decodeArachnidDeployerTx(t)
	if tx.Protected() {
		t.Fatal("Arachnid deployer tx must be unprotected (pre-EIP-155)")
	}

	t.Run("mainnet admits the canonical deployer", func(t *testing.T) {
		cfg := submitTxChainConfig(big.NewInt(int64(opera.VinuChainMainNetworkID)))
		backend := &submitTxBoundaryBackend{config: &cfg, block: submitTxBlock(common.Big0)}
		got, err := SubmitTransaction(context.Background(), backend, tx)
		if err != nil {
			t.Fatalf("mainnet must admit the canonical Arachnid deployer: %v", err)
		}
		if got != tx.Hash() {
			t.Fatalf("returned hash = %s, want %s", got, tx.Hash())
		}
		if backend.sent != tx {
			t.Fatal("admitted tx was not forwarded to the backend")
		}
	})

	t.Run("testnet refuses the same tx", func(t *testing.T) {
		cfg := submitTxChainConfig(big.NewInt(206))
		backend := &submitTxBoundaryBackend{config: &cfg, block: submitTxBlock(common.Big0)}
		got, err := SubmitTransaction(context.Background(), backend, tx)
		if err == nil {
			t.Fatal("testnet must refuse the unprotected tx (carve-out is mainnet-only; the flag is off)")
		}
		if got != (common.Hash{}) {
			t.Fatalf("refused tx must return zero hash, got %s", got)
		}
		if backend.sent != nil {
			t.Fatal("refused tx must not be forwarded to the backend")
		}
	})
}

func submitTxChainConfig(chainID *big.Int) params.ChainConfig {
	cfg := *params.TestChainConfig
	cfg.ChainID = chainID
	cfg.HomesteadBlock = common.Big0
	cfg.EIP155Block = common.Big0
	cfg.BerlinBlock = common.Big0
	cfg.LondonBlock = common.Big0
	return cfg
}

func submitTxBlock(number *big.Int) *evmcore.EvmBlock {
	return evmcore.NewEvmBlock(&evmcore.EvmHeader{
		Number:   number,
		GasLimit: 1_000_000,
		BaseFee:  big.NewInt(1),
	}, nil)
}

func signedSetCodeSubmitTx(t *testing.T, chainID *big.Int) *types.Transaction {
	t.Helper()

	senderKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	authorityKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	auth, err := types.SignSetCode(authorityKey, types.SetCodeAuthorization{
		ChainID: chainID,
		Address: common.HexToAddress("0x4000000000000000000000000000000000000000"),
		Nonce:   0,
	})
	if err != nil {
		t.Fatal(err)
	}
	tx, err := types.SignNewTx(senderKey, types.NewPragueSigner(chainID), &types.SetCodeTx{
		ChainID:   chainID,
		Nonce:     1,
		GasTipCap: big.NewInt(1),
		GasFeeCap: big.NewInt(1),
		Gas:       params.TxGas + params.CallNewAccountGas,
		To:        common.HexToAddress("0x3000000000000000000000000000000000000000"),
		Value:     new(big.Int),
		AuthList:  []types.SetCodeAuthorization{auth},
	})
	if err != nil {
		t.Fatal(err)
	}
	return tx
}
