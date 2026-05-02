package opera

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	ethparams "github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera/contracts/evmwriter"
)

const (
	MainNetworkID              uint64 = 0xfa
	TestNetworkID              uint64 = 0xfa2
	VinuChainStagingNetworkID         = 0xcd // 205 (test mainnet / staging)
	VinuChainTestNetworkID            = 0xce // 206
	VinuChainMainNetworkID            = 0xcf // 207
	VinuChainNewNetworkID             = 0x1b
	DefaultEventGas            uint64 = 28000
	berlinBit                         = 1 << 0
	londonBit                         = 1 << 1
	llrBit                            = 1 << 2
	podgoricaBit                      = 1 << 3
	sfcV2Bit                          = 1 << 4
	elemontBit                        = 1 << 5
	sfcV2PatchBit                     = 1 << 6
	sfcV2Patch2Bit                    = 1 << 7
	sfcV2Patch3Bit                    = 1 << 8
	sfcV2Patch4Bit                    = 1 << 9
	elemontPubkeyValidationBit        = 1 << 10
	sfcV2Patch5Bit                    = 1 << 11
)

var DefaultVMConfig = vm.Config{
	StatePrecompiles: map[common.Address]vm.PrecompiledStateContract{
		evmwriter.ContractAddress: &evmwriter.PreCompiledContract{},
	},
}

type RulesRLP struct {
	Name      string
	NetworkID uint64

	// Graph options
	Dag DagRules

	// Epochs options
	Epochs EpochsRules

	// Blockchain options
	Blocks BlocksRules

	// Economy options
	Economy EconomyRules

	Upgrades Upgrades `rlp:"-"`
}

// Rules describes opera net.
// Note keep track of all the non-copiable variables in Copy()
type Rules RulesRLP

// GasPowerRules defines gas power rules in the consensus.
type GasPowerRules struct {
	AllocPerSec        uint64
	MaxAllocPeriod     inter.Timestamp
	StartupAllocPeriod inter.Timestamp
	MinStartupGas      uint64
}

type GasRulesRLPV1 struct {
	MaxEventGas  uint64
	EventGas     uint64
	ParentGas    uint64
	ExtraDataGas uint64
	// Post-LLR fields
	BlockVotesBaseGas    uint64
	BlockVoteGas         uint64
	EpochVoteGas         uint64
	MisbehaviourProofGas uint64
}

type GasRules GasRulesRLPV1

type EpochsRules struct {
	MaxEpochGas      uint64
	MaxEpochDuration inter.Timestamp
}

// DagRules of VinuChain DAG (directed acyclic graph).
type DagRules struct {
	MaxParents     idx.Event
	MaxFreeParents idx.Event // maximum number of parents with no gas cost
	MaxExtraData   uint32
}

// BlocksMissed is information about missed blocks from a staker
type BlocksMissed struct {
	BlocksNum idx.Block
	Period    inter.Timestamp
}

// EconomyRules contains economy constants
type EconomyRules struct {
	BlockMissedSlack idx.Block

	Gas GasRules

	MinGasPrice *big.Int

	ShortGasPower GasPowerRules
	LongGasPower  GasPowerRules

	QuotaCacheAddress      common.Address `rlp:"optional"`
	QuotaCacheMaxAddresses uint64         `rlp:"optional"`
}

// BlocksRules contains blocks constants
type BlocksRules struct {
	MaxBlockGas             uint64 // technical hard limit, gas is mostly governed by gas power allocation
	MaxEmptyBlockSkipPeriod inter.Timestamp
}

// Upgrades tracks which hard fork features are active on the network.
// Each boolean flag gates consensus-critical behavior changes. Adding a new
// upgrade requires updating this struct, the bitfield encoding in Copy/RLP,
// the EvmChainConfig mapping, and the network rule constructors. A fork
// registry pattern would reduce this coupling, but for now all four sites
// must be kept in sync when introducing new upgrades.
//
// Upgrade flags are protected from on-chain governance changes by UpdateRules.
// Activating a new upgrade on an existing network requires shipping a new
// binary with the flag set in the corresponding hardcoded rule constructor
// (e.g. VinuChainMainNetRules, VinuChainTestNetRules).
type Upgrades struct {
	Berlin    bool
	London    bool
	Llr       bool
	Podgorica bool
	// SfcV2 enables the V2 SFC bytecode upgrade and the 30% fee burn mechanism.
	// Activation requires a new binary release with SfcV2 set to true in the
	// network's hardcoded rule constructor; governance cannot toggle it.
	SfcV2 bool
	// Elemont gates consensus-critical behavioral fixes: NoCheaters merged view,
	// AdvanceEpochs full 32-byte ABI decode, cheater fee zeroing, vecmt
	// GatherFrom tie-breaking, MedianTime stable sort, and empty-pubkey
	// validator skip at epoch seal. All changes activate atomically.
	Elemont bool
	// SfcV2Patch re-flashes the SFC V2 bytecode at sfc.ContractAddress on
	// activation, overwriting whatever bytecode the prior SfcV2 transition
	// installed. It exists so networks that activated SfcV2 with an earlier
	// (buggy) version of the V2 bytecode can upgrade to a corrected version
	// without needing a second hard fork or chain reset. On networks that
	// have not yet activated SfcV2, this flag is a no-op — the initial
	// SfcV2 activation already installs whatever GetContractBin() currently
	// returns, so mainnet operators do not need to enable SfcV2Patch when
	// performing their first SfcV2 activation.
	SfcV2Patch bool
	// SfcV2Patch2 re-flashes the SFC V2 bytecode a second time. It exists
	// because testnet activated SfcV2Patch while GetContractBin() still
	// returned the b7ab5b5-era 43,743-byte bytecode (missing the
	// reentrancyguard fix and all Cycle 158 hardening). SfcV2Patch is now
	// exhausted on testnet and cannot fire again. SfcV2Patch2 provides a
	// new one-shot transition so the current 45,240-byte Cycle-158 bytecode
	// is installed on testnet without a chain reset. Mainnet does not need
	// this flag — its first SfcV2 activation installs the latest bytecode
	// directly.
	SfcV2Patch2 bool
	// SfcV2Patch3 re-flashes the SFC V2 bytecode a third time to install the
	// Cycle-159 bytecode. Required because SfcV2Patch2 installed bytecode
	// whose inline reentrancy guard required _reentrancyGuardCounter == 1,
	// but the storage slot was appended after genesis and is 0 on-chain —
	// bricking every nonReentrant function (delegate, undelegate, withdraw,
	// claimRewards, restakeRewards, stashRewards, createValidator). Cycle-159
	// relaxes the guard to < 2 so the 0-initialised slot is accepted as
	// "not entered" while still failing closed on corruption. Testnet-only:
	// mainnet has not yet activated SfcV2/SfcV2Patch/SfcV2Patch2 and will
	// consume the Cycle-159 bytecode directly on its first SfcV2 activation.
	SfcV2Patch3 bool
	// SfcV2Patch4 re-flashes the SFC V2 bytecode a fourth time to install the
	// Cycle-160 bytecode sourced from VinuChain/vinuchain-lists PR #2. The
	// Cycle-160 delta fixes the `relock`/`extendLock` lock-end-time bug where
	// the existing-vs-new lock comparison checked `newDuration >= oldDuration`
	// instead of `newEndTime >= oldEndTime`, letting a staker silently shorten
	// their effective lock period by relocking with a shorter duration from a
	// point later in time. Testnet-only: mainnet has not yet activated any
	// SfcV2* flag and will consume the latest available bytecode directly on
	// its first SfcV2 activation.
	//
	// The patch4 bytecode is re-flashed via sfc.GetPatch4ContractBin(), which
	// is a separate source from GetContractBin() so the Cycle-160 asset can
	// be compiled and dropped in as a scaffolded placeholder until PR #2
	// merges. A validation guard in sfc_patch4_bytecode.go log.Crits at the
	// first attempted re-flash if the placeholder bytes have not yet been
	// replaced with a real compiled SFC, preventing a release from silently
	// shipping the sentinel.
	SfcV2Patch4 bool
	// ElemontPubkeyValidation rejects validators whose pubkey deviates from
	// the canonical secp256k1 shape (1 type byte 0xc0 + 65 raw bytes) at
	// epoch seal. Required because inter/validatorpk/pubkey.FromBytes accepts
	// any non-empty byte slice without validation, which let a malformed
	// 65-byte 0x04-prefixed pubkey enter the testnet validator set at epoch
	// 5682. The flag is intentionally separate from Elemont so existing
	// chaindata replay (which already admitted that malformed validator)
	// stays bit-for-bit identical: only NEW admissions evaluated AFTER the
	// flag is set in a hardcoded rule constructor are subject to the check.
	// Defaulted to false on every current network constructor; flipping it
	// to true is a hard fork delivered via a new binary release, identical
	// in shape to how Elemont and the SfcV2Patch* flags are activated.
	ElemontPubkeyValidation bool
	// SfcV2Patch5 re-flashes the SFC V2 bytecode a fifth time to install the
	// Cycle-161 bytecode sourced from VinuChain/vinuchain-lists. The Cycle-161
	// delta tightens validator-pubkey ingress: createValidator,
	// _rawCreateValidator (which also covers setGenesisValidator), and
	// NodeDriverAuth.updateValidatorPubkey now require pubkey.length == 66
	// AND pubkey[0] == 0xc0. This rejects the malformed-pubkey shape that
	// admitted testnet validator 16 (a 65-byte 0x04-prefixed pubkey lacking
	// the canonical secp256k1 type-byte prefix), which lachesis-base could
	// not verify, leaving 270k VC of delegator stake earning zero rewards.
	// Testnet-only at activation time: mainnet has not yet activated any
	// SfcV2* flag and will consume the latest available bytecode directly
	// on its first SfcV2 activation.
	//
	// The patch5 bytecode is re-flashed via sfc.GetPatch5ContractBin(),
	// which is a separate source from GetContractBin() and
	// GetPatch4ContractBin() so the Cycle-161 asset can be compiled and
	// dropped in as a scaffolded placeholder. A validation guard in
	// sfc_patch5_bytecode.go log.Crits at the first attempted re-flash if
	// the placeholder bytes have not yet been replaced with a real compiled
	// SFC, preventing a release from silently shipping the sentinel.
	SfcV2Patch5 bool
}

type UpgradeHeight struct {
	Upgrades Upgrades
	Height   idx.Block
}

// EvmChainConfig returns ChainConfig for transactions signing and execution
func (r Rules) EvmChainConfig(hh []UpgradeHeight) *ethparams.ChainConfig {
	cfg := *ethparams.AllEthashProtocolChanges
	cfg.ChainID = new(big.Int).SetUint64(r.NetworkID)
	cfg.BerlinBlock = nil
	cfg.LondonBlock = nil
	for i, h := range hh {
		height := new(big.Int)
		if i > 0 {
			height.SetUint64(uint64(h.Height))
		}
		if cfg.BerlinBlock == nil && h.Upgrades.Berlin {
			cfg.BerlinBlock = height
		}
		if !h.Upgrades.Berlin {
			cfg.BerlinBlock = nil
		}

		if cfg.LondonBlock == nil && h.Upgrades.London {
			cfg.LondonBlock = height
		}
		if !h.Upgrades.London {
			cfg.LondonBlock = nil
		}
	}
	return &cfg
}

func MainNetRules() Rules {
	return Rules{
		Name:      "main",
		NetworkID: MainNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    DefaultEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(1 * time.Minute),
		},
	}
}

func TestNetRules() Rules {
	return Rules{
		Name:      "test",
		NetworkID: TestNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    DefaultEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(1 * time.Minute),
		},
	}
}

func FakeNetRules() Rules {
	return Rules{
		Name:      "vinuchain",
		NetworkID: VinuChainNewNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VinuChainNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(3 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin:      true,
			London:      true,
			Llr:         true,
			Podgorica:   true,
			SfcV2:       true,
			Elemont:     true,
			SfcV2Patch:  true,
			SfcV2Patch2: true,
			SfcV2Patch3: true,
			SfcV2Patch4: true,
		},
	}
}

func LegacyFakeNetRules() Rules {
	return Rules{
		Name:      "vinuchain",
		NetworkID: VinuChainNewNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    FakeNetEpochsRules(),
		Economy:   FakeEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(3 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin:      true,
			London:      true,
			Llr:         true,
			Podgorica:   true,
			SfcV2:       true,
			Elemont:     true,
			SfcV2Patch:  true,
			SfcV2Patch2: true,
			SfcV2Patch3: true,
			SfcV2Patch4: true,
		},
	}
}

// VinuChainTestNetRules returns testnet rules
func VinuChainTestNetRules() Rules {
	return Rules{
		Name:      "VinuChain Testnet",
		NetworkID: VinuChainTestNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VinuChainNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(10 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin:      true,
			London:      true,
			Llr:         true,
			Podgorica:   true,
			SfcV2:       true,
			Elemont:     true,
			SfcV2Patch:  true,
			SfcV2Patch2: true,
			SfcV2Patch3: true,
			SfcV2Patch4: true,
		},
	}
}

// VinuChainMainNetRules returns mainnet rules.
// Podgorica activation requires a new binary release that sets Podgorica=true
// in this function. Governance cannot enable Podgorica via UpdateRules because
// upgrade flags are protected (marshal.go strips Upgrades from on-chain updates).
// The QuotaCacheAddress is also protected from governance changes.
//
// SfcV2 activation on mainnet requires a new binary release that sets SfcV2: true
// in this function. Governance cannot enable SfcV2 via UpdateRules because upgrade
// flags are stripped from on-chain rule updates. This is intentional — the SfcV2
// upgrade replaces SFC contract bytecode at startup, which must be coordinated
// across all validators via a code release, not an on-chain governance proposal.
//
// SfcV2Patch and SfcV2Patch2 are intentionally NOT set here. Both flags exist
// only to re-flash the SFC bytecode on networks that activated SfcV2 with
// stale bytecode (testnet only). Mainnet has not yet activated SfcV2, so
// its first SfcV2 transition will install whatever GetContractBin() currently
// returns — which is the latest corrected bytecode — and no re-flash is
// needed. Leave both unset here; adding them would be harmless no-ops but
// would obscure the invariant that mainnet's first SfcV2 activation already
// picks up every subsequent correctness fix to the V2 bytecode automatically.
func VinuChainMainNetRules() Rules {
	return Rules{
		Name:      "VinuChain Mainnet",
		NetworkID: VinuChainMainNetworkID,
		Dag:       DefaultDagRules(),
		Epochs:    VinuChainNetEpochsRules(),
		Economy:   DefaultEconomyRules(),
		Blocks: BlocksRules{
			MaxBlockGas:             20500000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(10 * time.Second),
		},
		Upgrades: Upgrades{
			Berlin:    true,
			London:    true,
			Llr:       true,
			Podgorica: true,
			SfcV2:     true,
			Elemont:   true,
		},
	}
}

// MainNetRulesForNetwork returns the hardcoded rules for a given network ID,
// or nil if the network ID is not a known mainnet/testnet. Used by migrations
// to propagate upgrade flags from the binary into the stored epoch state.
func MainNetRulesForNetwork(networkID uint64) *Rules {
	switch networkID {
	case VinuChainMainNetworkID:
		r := VinuChainMainNetRules()
		return &r
	case VinuChainTestNetworkID:
		r := VinuChainTestNetRules()
		return &r
	case VinuChainStagingNetworkID:
		r := VinuChainMainNetRules()
		r.NetworkID = VinuChainStagingNetworkID
		return &r
	default:
		return nil
	}
}

// DefaultEconomyRules returns mainnet economy
func DefaultEconomyRules() EconomyRules {
	return EconomyRules{
		BlockMissedSlack:       50,
		Gas:                    DefaultGasRules(),
		MinGasPrice:            big.NewInt(1e9),
		ShortGasPower:          DefaultShortGasPowerRules(),
		LongGasPower:           DefaultLongGasPowerRules(),
		QuotaCacheAddress:      common.HexToAddress("0x9D6Aa03a8D4AcF7b43c562f349Ee45b3214c3bbF"),
		QuotaCacheMaxAddresses: 10000,
	}
}

// FakeEconomyRules returns fakenet economy
func FakeEconomyRules() EconomyRules {
	cfg := DefaultEconomyRules()
	cfg.ShortGasPower = FakeShortGasPowerRules()
	cfg.LongGasPower = FakeLongGasPowerRules()
	return cfg
}

func DefaultDagRules() DagRules {
	return DagRules{
		MaxParents:     10,
		MaxFreeParents: 3,
		MaxExtraData:   128,
	}
}

func DefaultEpochsRules() EpochsRules {
	return EpochsRules{
		MaxEpochGas:      1500000000,
		MaxEpochDuration: inter.Timestamp(4 * time.Hour),
	}
}

func DefaultGasRules() GasRules {
	return GasRules{
		MaxEventGas:          10000000 + DefaultEventGas,
		EventGas:             DefaultEventGas,
		ParentGas:            2400,
		ExtraDataGas:         25,
		BlockVotesBaseGas:    1024,
		BlockVoteGas:         512,
		EpochVoteGas:         1536,
		MisbehaviourProofGas: 71536,
	}
}

func VinuChainNetEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochDuration = inter.Timestamp(4 * time.Hour)
	return cfg
}

func FakeNetEpochsRules() EpochsRules {
	cfg := DefaultEpochsRules()
	cfg.MaxEpochGas /= 5
	cfg.MaxEpochDuration = inter.Timestamp(4 * time.Hour)
	return cfg
}

// DefaultLongGasPowerRules is long-window config
func DefaultLongGasPowerRules() GasPowerRules {
	return GasPowerRules{
		AllocPerSec:        100 * DefaultEventGas,
		MaxAllocPeriod:     inter.Timestamp(60 * time.Minute),
		StartupAllocPeriod: inter.Timestamp(5 * time.Second),
		MinStartupGas:      DefaultEventGas * 20,
	}
}

// DefaultShortGasPowerRules is short-window config
func DefaultShortGasPowerRules() GasPowerRules {
	// 2x faster allocation rate, 6x lower max accumulated gas power
	cfg := DefaultLongGasPowerRules()
	cfg.AllocPerSec *= 2
	cfg.StartupAllocPeriod /= 2
	cfg.MaxAllocPeriod /= 2 * 6
	return cfg
}

// FakeLongGasPowerRules is fake long-window config
func FakeLongGasPowerRules() GasPowerRules {
	config := DefaultLongGasPowerRules()
	config.AllocPerSec *= 1000
	return config
}

// FakeShortGasPowerRules is fake short-window config
func FakeShortGasPowerRules() GasPowerRules {
	config := DefaultShortGasPowerRules()
	config.AllocPerSec *= 1000
	return config
}

func (r Rules) Copy() Rules {
	cp := r
	if r.Economy.MinGasPrice != nil {
		cp.Economy.MinGasPrice = new(big.Int).Set(r.Economy.MinGasPrice)
	}
	return cp
}

func (r Rules) String() string {
	b, _ := json.Marshal(&r)
	return string(b)
}
