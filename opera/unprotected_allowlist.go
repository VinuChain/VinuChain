package opera

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// arachnidDeployerTxHash is the keccak256 of the canonical Arachnid
// deterministic-deployment-proxy transaction (Nick's method). This single
// pre-signed, pre-EIP-155 transaction deploys the deterministic deployer at
// 0x4e59b44847b379578588920cA78FbF26c0B4956C on every EVM chain, which in
// turn lets ERC-4337 EntryPoint v0.7 land at the canonical cross-chain
// address 0x0000000071727De22E5E9d8BAf0edAc6f37da032 via CREATE2.
//
// The hash is chain-independent precisely because the transaction is
// unprotected (no chain ID in the signed payload). Pinning the exact 32-byte
// transaction hash — not just the sender/nonce — means no other unprotected
// transaction can satisfy the mainnet allowlist, even one crafted by the same
// throwaway signer.
//
// This transaction is replay-benign: it deploys a stateless, fund-less factory
// contract, and its signer (0x3fab184622dc19b6109349b94811493bf2a45362,
// recovered from the magic r=s=0x2222… signature) holds no funds on any chain.
// Replaying it elsewhere merely re-creates the same factory — the entire point
// of deterministic deployment.
var arachnidDeployerTxHash = common.HexToHash(
	"0xeddf9e61fb9d8f5111840daef55e5fde0041f5702856532cdbb5a02998033d26",
)

// UnprotectedTxAllowlistedOnMainnet reports whether the given pre-EIP-155
// transaction is the one specific transaction VinuChain mainnet admits despite
// its lack of replay protection: the canonical Arachnid deterministic deployer.
//
// This is intentionally the ONLY exception to mainnet's blanket rejection of
// unprotected transactions (ethapi/api.go SubmitTransaction). It does NOT
// depend on, and must not be confused with, the gossip.Config.AllowUnprotectedTxs
// flag — that flag remains permanently refused on mainnet by the startup guard
// in gossip/service.go. This allowlist is a separate, always-on, single-tx
// carve-out that adds effectively zero replay surface.
func UnprotectedTxAllowlistedOnMainnet(tx *types.Transaction) bool {
	if tx == nil {
		return false
	}
	return tx.Hash() == arachnidDeployerTxHash
}

// ArachnidDeployerTxHash returns the pinned canonical deployer transaction
// hash. Exposed for tests and operational tooling.
func ArachnidDeployerTxHash() common.Hash {
	return arachnidDeployerTxHash
}
