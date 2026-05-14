package opera

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// PaybackV2 contract addresses per network. These are the addresses that
// Economy.QuotaCacheAddress is overwritten to at PaybackV2 activation
// (see gossip/block_processor.go sealEpochIfNeeded). Each address must be
// a freshly-deployed QuotaContractV2 (vinu-quotacontract repo) whose owner
// is a recoverable EOA — the whole point of the PaybackV2 upgrade is to
// escape the original proxy's lost-ProxyAdmin-key trap.
//
// Until the contract is deployed and its address is recorded here, the
// sentinel zero value remains. The startup check in
// EnforcePaybackV2StartupCheck() refuses to boot a node where PaybackV2
// is set true in any rule constructor while the matching network's
// address is still the sentinel.
var (
	// paybackV2TestnetAddress is the QuotaContractV2 address on VinuChain
	// testnet (chain 206). Replace common.Address{} with the
	// `deployed addr` printed by scripts/deploy-quotacontract-v2.ts after
	// running it on testnet. The same value must also be referenced in
	// the testnet finalizer / runbook so explorers can be re-pointed.
	paybackV2TestnetAddress = common.Address{}

	// paybackV2MainnetAddress is the QuotaContractV2 address on VinuChain
	// mainnet (chain 207). Stays sentinel until the mainnet rollout. See
	// .claude/rules/deployment-log.md -> "Mainnet PaybackV2 Upgrade" for
	// the mainnet-only prerequisite checklist that must complete before
	// this constant is filled in and PaybackV2 is enabled on
	// VinuChainMainNetRules().
	paybackV2MainnetAddress = common.Address{}

	// paybackV2StagingAddress mirrors mainnet for the staging network
	// (chain 205). Kept separate so a staging rehearsal can deploy its
	// own contract without entangling the mainnet address slot.
	paybackV2StagingAddress = common.Address{}
)

// PaybackV2ContractAddress returns the V2 contract address for the given
// network ID. Returns an error for unknown networks. Returns the zero
// address (sentinel) for networks where the V2 address has not yet been
// recorded — callers MUST treat the zero address as a "not yet deployed"
// signal and refuse to use it (see the activation gate in
// gossip/block_processor.go sealEpochIfNeeded, which log.Crits rather
// than swap to a zero address).
func PaybackV2ContractAddress(networkID uint64) (common.Address, error) {
	switch networkID {
	case VinuChainTestNetworkID:
		return paybackV2TestnetAddress, nil
	case VinuChainMainNetworkID:
		return paybackV2MainnetAddress, nil
	case VinuChainStagingNetworkID:
		return paybackV2StagingAddress, nil
	case VinuChainNewNetworkID, TestNetworkID, MainNetworkID:
		// Fakenet / legacy: PaybackV2 may be enabled in FakeNetRules() for
		// in-process exercising of the activation path. Tests inject the
		// address they want via SetPaybackV2ContractAddressForTesting.
		return common.Address{}, nil
	default:
		return common.Address{}, fmt.Errorf("PaybackV2: no V2 contract address configured for network %d (0x%x)", networkID, networkID)
	}
}

// PaybackV2AddressIsSentinel reports whether the given address is the zero
// sentinel that signals "V2 contract not yet deployed and recorded". Used
// by the startup check and by the activation gate in the block processor.
func PaybackV2AddressIsSentinel(addr common.Address) bool {
	return addr == (common.Address{})
}

// SetPaybackV2ContractAddressForTesting overrides the per-network V2
// address. Test-only helper — production code must rely on the
// build-time-baked addresses. Returns a cleanup function that restores
// the previous value. Safe to call only from tests.
func SetPaybackV2ContractAddressForTesting(networkID uint64, addr common.Address) (restore func()) {
	switch networkID {
	case VinuChainTestNetworkID:
		prev := paybackV2TestnetAddress
		paybackV2TestnetAddress = addr
		return func() { paybackV2TestnetAddress = prev }
	case VinuChainMainNetworkID:
		prev := paybackV2MainnetAddress
		paybackV2MainnetAddress = addr
		return func() { paybackV2MainnetAddress = prev }
	case VinuChainStagingNetworkID:
		prev := paybackV2StagingAddress
		paybackV2StagingAddress = addr
		return func() { paybackV2StagingAddress = prev }
	default:
		return func() {}
	}
}

// EnforcePaybackV2StartupCheck fails fast if any of the hardcoded rule
// constructors enable PaybackV2 while the V2 contract address for that
// network is still the zero sentinel. Wired into the opera binary via
// cmd/opera/launcher/payback_v2_startup_check.go's init().
//
// The check is per-network: enabling PaybackV2 on testnet rules with a
// real testnet address while mainnet rules leave PaybackV2 false (and
// mainnet address still sentinel) is the expected staged-rollout state.
// The check ONLY refuses to start when a network's hardcoded rules say
// "activate PaybackV2" but its address slot is empty.
func EnforcePaybackV2StartupCheck() {
	checks := []struct {
		network string
		rules   Rules
		addr    common.Address
	}{
		{"VinuChain Testnet", VinuChainTestNetRules(), paybackV2TestnetAddress},
		{"VinuChain Mainnet", VinuChainMainNetRules(), paybackV2MainnetAddress},
	}
	for _, c := range checks {
		if c.rules.Upgrades.PaybackV2 && PaybackV2AddressIsSentinel(c.addr) {
			panic(errors.New(
				"PaybackV2 startup check failed: " + c.network +
					" rule constructor has PaybackV2=true but the V2 contract address is still the zero sentinel. " +
					"Deploy QuotaContractV2 via scripts/deploy-quotacontract-v2.ts and record its address in opera/payback_v2_address.go before shipping this binary.",
			))
		}
	}
}
