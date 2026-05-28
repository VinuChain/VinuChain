package gossip

import (
	"errors"

	"github.com/Fantom-foundation/go-opera/opera"
)

// errUnprotectedTxsOnMainnet is returned when an operator attempts to admit
// pre-EIP-155 (replay-vulnerable) transactions on mainnet.
var errUnprotectedTxsOnMainnet = errors.New("AllowUnprotectedTxs cannot be enabled on mainnet (NetworkID 207)")

// checkUnprotectedTxsPolicy enforces VinuChain's invariant that replay-vulnerable
// (pre-EIP-155) transactions can never be admitted over RPC on mainnet, regardless
// of how AllowUnprotectedTxs is set via CLI, TOML, or env. NewService calls this
// before the RPC backend is constructed, so the refusal fires before any RPC
// server can accept a transaction.
//
// The flag is defined in cmd/opera/launcher while this guard lives here: keep the
// NewService call wired so a partial cherry-pick cannot separate the flag from its
// mainnet guard.
func checkUnprotectedTxsPolicy(allowUnprotectedTxs bool, networkID uint64) error {
	if allowUnprotectedTxs && networkID == opera.VinuChainMainNetworkID {
		return errUnprotectedTxsOnMainnet
	}
	return nil
}
