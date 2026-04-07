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
	MainNetworkID     uint64 = 0xfa
	TestNetworkID     uint64 = 0xfa2
	VinuChainStagingNetworkID     = 0xcd // 205 (test mainnet / staging)
	VinuChainTestNetworkID        = 0xce // 206
	VinuChainMainNetworkID        = 0xcf // 207
	VinuChainNewNetworkID         = 0x1b
	DefaultEventGas   uint64 = 28000
	berlinBit                = 1 << 0
	londonBit                = 1 << 1
	llrBit                   = 1 << 2
	podgoricaBit             = 1 << 3
	sfcV2Bit                 = 1 << 4
	elemontBit               = 1 << 5
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

	QuotaCacheAddress    common.Address `rlp:"optional"`
	QuotaCacheMaxAddresses uint64       `rlp:"optional"`
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
			Berlin:  true,
			London:  true,
			Llr:     true,
			SfcV2:   true,
			Elemont: true,
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
			Berlin:  true,
			London:  true,
			Llr:     true,
			SfcV2:   true,
			Elemont: true,
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
			Berlin:  true,
			London:  true,
			Llr:     true,
			SfcV2:   true,
			Elemont: true,
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
			Berlin: true,
			London: true,
			Llr:    true,
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
