package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/evmcore"
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
