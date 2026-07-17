package launcher

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/Fantom-foundation/lachesis-base/abft"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/naoina/toml"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/gossip"
	"github.com/Fantom-foundation/go-opera/gossip/emitter"
	"github.com/Fantom-foundation/go-opera/integration"
	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	futils "github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/go-opera/vecmt"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(nodeFlags, testFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}
	checkConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(checkConfig),
		Name:        "checkconfig",
		Usage:       "Checks configuration file",
		ArgsUsage:   "",
		Flags:       append(nodeFlags, testFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The checkconfig checks configuration file.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}

	// DataDirFlag defines directory to store VinuChain state and user's wallets
	DataDirFlag = utils.DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: utils.DirectoryString(DefaultDataDir()),
	}

	ValidatorsFileFlag = cli.StringFlag{
		Name:  "validatorsfile",
		Usage: "Path to validators file",
	}

	CacheFlag = cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching",
		Value: DefaultCacheSize,
	}
	// GenesisFlag specifies network genesis configuration
	GenesisFlag = cli.StringFlag{
		Name:  "genesis",
		Usage: "'path to genesis file' - sets the network genesis configuration.",
	}
	ExperimentalGenesisFlag = cli.BoolFlag{
		Name:  "genesis.allowExperimental",
		Usage: "Allow to use experimental genesis file.",
	}

	FastEmitFlag = cli.BoolFlag{
		Name:  "fastemit",
		Usage: "Enable fast emit mode (1 block per 10 seconds)",
	}

	NetworkMainnetFlag = cli.BoolFlag{
		Name:  "network.mainnet",
		Usage: "Use mainnet network configuration",
	}

	RPCGlobalGasCapFlag = cli.Uint64Flag{
		Name:  "rpc.gascap",
		Usage: "Sets a cap on gas that can be used in vc_call/estimateGas (0=infinite)",
		Value: gossip.DefaultConfig(cachescale.Identity).RPCGasCap,
	}
	RPCGlobalTxFeeCapFlag = cli.Float64Flag{
		Name:  "rpc.txfeecap",
		Usage: "Sets a cap on transaction fee (in VC) that can be sent via the RPC APIs (0 = no cap)",
		Value: gossip.DefaultConfig(cachescale.Identity).RPCTxFeeCap,
	}
	RPCGlobalTimeoutFlag = cli.DurationFlag{
		Name:  "rpc.timeout",
		Usage: "Time limit for RPC calls execution",
		Value: gossip.DefaultConfig(cachescale.Identity).RPCTimeout,
	}

	// RPCAllowUnprotectedTxsFlag exposes the gossip.Config AllowUnprotectedTxs
	// knob via the CLI. Required to land Arachnid's deterministic deployer
	// (Nick's method) which signs with v=27/28 and no chain ID — needed
	// to deploy ERC-4337 EntryPoint at the canonical cross-chain address
	// 0x0000000071727De22E5E9d8BAf0edAc6f37da032. The gossip-layer guard
	// in service.go refuses to honour this on mainnet (NetworkID 207)
	// regardless of how the flag is set — replay-vulnerable txs stay
	// blocked on mainnet.
	RPCAllowUnprotectedTxsFlag = cli.BoolFlag{
		Name: "rpc.allow-unprotected-txs",
		Usage: "Allow pre-EIP-155 (replay-vulnerable) transactions over RPC. " +
			"Refused on mainnet (NetworkID 207) by the gossip-layer guard; for non-mainnet networks only.",
	}

	SyncModeFlag = cli.StringFlag{
		Name:  "syncmode",
		Usage: `Blockchain sync mode ("full" or "snap")`,
		Value: "full",
	}

	GCModeFlag = cli.StringFlag{
		Name:  "gcmode",
		Usage: `Blockchain garbage collection mode ("light", "full", "archive")`,
		Value: "full",
	}

	ExitWhenAgeFlag = cli.DurationFlag{
		Name:  "exitwhensynced.age",
		Usage: "Exits after synchronisation reaches the required age",
	}
	ExitWhenEpochFlag = cli.Uint64Flag{
		Name:  "exitwhensynced.epoch",
		Usage: "Exits after synchronisation reaches the required epoch",
	}

	DBMigrationModeFlag = cli.StringFlag{
		Name:  "db.migration.mode",
		Usage: "MultiDB migration mode ('reformat' or 'rebuild')",
	}
	DBPresetFlag = cli.StringFlag{
		Name:  "db.preset",
		Usage: "DBs layout preset ('pbl-1' or 'ldb-1' or 'legacy-ldb' or 'legacy-pbl')",
	}
	TraceNodeFlag = cli.BoolFlag{
		Name:  "tracenode",
		Usage: "If present, this node records inner transaction traces",
	}
)

type GenesisTemplate struct {
	Name   string
	Header genesis.Header
	Hashes genesis.Hashes
	// SupersededBy, when non-empty, names the replacement genesis (URL or
	// file) for a preset that pre-dates staged upgrade activations. Such a
	// preset is refused for fresh installs: a replay from it re-stages the
	// missing upgrades at a different seal than the live chain's historical
	// activations and diverges with "wrong event epoch hash". Datadirs that
	// already carry chain state are verified independently at startup via
	// checkStoredChainState, which is keyed on the persisted GenesisID and
	// therefore also covers boots without any --genesis flag.
	SupersededBy string
}

// UpgradeActivation records when an upgrade became active on a live public
// network. These are historical facts, not policy — they are read back from
// the published genesis' epoch history by
// TestCurrentTestnetGenesisSatisfiesStoredStateRequirement.
type UpgradeActivation struct {
	// Name identifies the upgrade in refusal messages.
	Name string
	// ActiveFromEpoch is the first epoch in which Active reports true on
	// the live chain, i.e. the seal that ended ActiveFromEpoch-1 performed
	// the activation.
	ActiveFromEpoch idx.Epoch
	// ActiveFromBlock is the first block executed under the new rules — the
	// Height a store records for this activation. Both writers agree on it:
	// block_processor.sealEpochIfNeeded records blockCtx.Idx+1 at the seal,
	// and Store.ApplyGenesis records LastBlock.Idx+1 replaying history, so
	// a live node and a fresh install carry the same value.
	ActiveFromBlock idx.Block
	// Active reports whether the upgrade is active in a rules set.
	Active func(opera.Upgrades) bool
}

// recordedActivationHeight returns the height at which upgradeHeights shows a
// first becoming active, or 0 if it never does.
func recordedActivationHeight(upgradeHeights []opera.UpgradeHeight, a UpgradeActivation) idx.Block {
	var found idx.Block
	for _, h := range upgradeHeights {
		if !a.Active(h.Upgrades) {
			continue
		}
		if found == 0 || h.Height < found {
			found = h.Height
		}
	}
	return found
}

// StoredStateRequirement pins the activation history a datadir of a known
// public-network genesis lineage must agree with before this binary may run
// it. Matching is on GenesisID — the lineage identity every published genesis
// file of a network shares — so generated private networks and fakenets
// (content-derived GenesisIDs) never match.
//
// Three independent conditions are checked (see checkStoredChainState):
//
//   - Consistency: the stored flags must equal the flags the live chain had
//     at the stored epoch.
//   - Provenance: every active upgrade must have been recorded at the block
//     the live chain activated it. Flags and epoch alone are a snapshot and
//     cannot establish history — a fork descended from a stale genesis that
//     kept sealing on its own would eventually present a current-looking
//     epoch with every flag set. Its stored UpgradeHeights still name the
//     local seals it actually performed, so this catches it.
//   - Currency: at most the newest activation may still be pending, so that
//     stageHardcodedUpgrades activates it at the canonical seal. A datadir
//     further behind would activate several upgrades at once at its own next
//     seal — a seal the live chain never performed — and diverge with "wrong
//     event epoch hash".
type StoredStateRequirement struct {
	// GenesisID identifies the public-network genesis lineage.
	GenesisID hash.Hash
	// NetworkName is used in the refusal message.
	NetworkName string
	// Activations is the lineage's historical activation record for every
	// upgrade this binary hardcodes as active.
	Activations []UpgradeActivation
	// Bootstrap names the current genesis (URL) a refused operator should
	// bootstrap a fresh datadir from.
	Bootstrap string
}

// minStoredEpoch is the oldest epoch a datadir may sit at and still activate
// every pending upgrade at the seal the live chain used: one epoch below the
// newest activation, so only that newest upgrade is left to stage.
func (req *StoredStateRequirement) minStoredEpoch() idx.Epoch {
	var newest idx.Epoch
	for _, a := range req.Activations {
		if a.ActiveFromEpoch > newest {
			newest = a.ActiveFromEpoch
		}
	}
	if newest == 0 {
		return 0
	}
	return newest - 1
}

// checkStoredChainState refuses to run a datadir of a known public-network
// genesis lineage whose persisted epoch/rules disagree with the lineage's
// upgrade activation history. It is enforced on every node boot (makeNode),
// independent of the CLI genesis input: restarts without --genesis, stale
// datadirs pointed at the current genesis file, and datadirs already bricked
// by a wrong-seal replay are all covered.
func checkStoredChainState(genesisID *hash.Hash, storedEpoch idx.Epoch, storedUpgrades opera.Upgrades,
	upgradeHeights []opera.UpgradeHeight) error {
	if genesisID == nil {
		return nil
	}
	for i := range StoredStateRequirements {
		req := &StoredStateRequirements[i]
		if req.GenesisID != *genesisID {
			continue
		}
		for _, a := range req.Activations {
			// Consistency with the live chain's history at this epoch.
			live := storedEpoch >= a.ActiveFromEpoch
			if a.Active(storedUpgrades) != live {
				state, expected := "inactive", "active"
				if !live {
					state, expected = "active", "inactive"
				}
				return fmt.Errorf("this datadir belongs to the %s network but disagrees with its upgrade "+
					"activation history: at epoch %d the live chain has %s %s, while this datadir has it %s "+
					"(%s activated at epoch %d). The datadir activated upgrades at a local epoch seal the live "+
					"chain never performed and has diverged with \"wrong event epoch hash\"; bootstrap a fresh "+
					"datadir from %s, or restore a current chaindata snapshot",
					req.NetworkName, storedEpoch, a.Name, expected, state, a.Name, a.ActiveFromEpoch,
					req.Bootstrap)
			}
			if !live {
				continue
			}
			// Provenance: it must have activated at the live chain's seal.
			if got := recordedActivationHeight(upgradeHeights, a); got != a.ActiveFromBlock {
				return fmt.Errorf("this datadir belongs to the %s network but activated %s at block %d, while "+
					"the live chain activated it at block %d: the datadir sealed the upgrade at a local block "+
					"the live chain never performed and has diverged with \"wrong event epoch hash\"; bootstrap "+
					"a fresh datadir from %s, or restore a current chaindata snapshot",
					req.NetworkName, a.Name, got, a.ActiveFromBlock, req.Bootstrap)
			}
		}
		// Currency: only the newest activation may still be pending.
		if min := req.minStoredEpoch(); storedEpoch < min {
			return fmt.Errorf("this datadir belongs to the %s network but stopped syncing at epoch %d, before its "+
				"upgrade activation seals (this binary may only resume a datadir at epoch %d or later): resuming "+
				"it would activate several upgrades at once at its next local epoch seal — a seal the live chain "+
				"never performed — and diverge with \"wrong event epoch hash\"; bootstrap a fresh datadir from "+
				"%s, or restore a current chaindata snapshot",
				req.NetworkName, storedEpoch, min, req.Bootstrap)
		}
		return nil
	}
	return nil
}

// checkGenesisPresetFreshness refuses a superseded genesis preset that is
// about to initialize a fresh (or interrupted-genesis) datadir, before the
// replay is even attempted. Datadirs that already carry chain state are
// verified at startup by checkStoredChainState instead.
func checkGenesisPresetFreshness(preset GenesisTemplate, firstLaunchPending bool) error {
	if preset.SupersededBy == "" || !firstLaunchPending {
		return nil
	}
	return fmt.Errorf("genesis preset %q is superseded and unsafe for fresh installs under current binary rules: "+
		"replaying it re-stages later upgrade activations at the wrong epoch seal and diverges from the live chain "+
		"with \"wrong event epoch hash\"; bootstrap from %s instead, or restore a current chaindata snapshot",
		preset.Name, preset.SupersededBy)
}

const (
	// DefaultCacheSize is calculated as memory consumption in a worst case scenario with default configuration
	// Average memory consumption might be 3-5 times lower than the maximum
	DefaultCacheSize  = 3600
	ConstantCacheSize = 600
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		return fmt.Errorf("field '%s' is not defined in %s", field, rt.String())
	},
}

type config struct {
	Node          node.Config
	Opera         gossip.Config
	Emitter       emitter.Config
	TxPool        evmcore.TxPoolConfig
	OperaStore    gossip.StoreConfig
	Lachesis      abft.Config
	LachesisStore abft.StoreConfig
	VectorClock   vecmt.IndexConfig
	DBs           integration.DBsConfig
}

func (c *config) AppConfigs() integration.Configs {
	return integration.Configs{
		Opera:         c.Opera,
		OperaStore:    c.OperaStore,
		Lachesis:      c.Lachesis,
		LachesisStore: c.LachesisStore,
		VectorClock:   c.VectorClock,
		DBs:           c.DBs,
	}
}

const maxConfigFileSize = 10 << 20 // 10 MB

func loadAllConfigs(file string, cfg *config) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	if info.Size() > maxConfigFileSize {
		return fmt.Errorf("config file %s exceeds maximum size of %d bytes", file, maxConfigFileSize)
	}

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	if err != nil {
		return fmt.Errorf("TOML config file error: %v\n"+ //nolint:staticcheck // ST1005: legacy multi-line error format
			"Use 'dumpconfig' command to get an example config file.\n"+
			"If node was recently upgraded and a previous network config file is used, then check updates for the config file.", err)
	}
	return nil
}

func mayGetGenesisStore(ctx *cli.Context, dataDir string) *genesisstore.Store {
	switch {
	case ctx.GlobalIsSet(FakeNetFlag.Name):
		_, num, err := parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
		if err != nil {
			log.Crit("Invalid flag", "flag", FakeNetFlag.Name, "err", err)
		}
		return makefakegenesis.FakeGenesisStore(num, futils.ToVC(333_333_333), futils.ToVC(200_000))
	case ctx.GlobalIsSet(GenesisFlag.Name):
		genesisPath := ctx.GlobalString(GenesisFlag.Name)

		f, err := os.Open(genesisPath)
		if err != nil {
			utils.Fatalf("Failed to open genesis file: %v", err)
		}
		genesisStore, genesisHashes, err := genesisstore.OpenGenesisStore(f)
		if err != nil {
			utils.Fatalf("Failed to read genesis file: %v", err)
		}

		// check if it's a trusted preset
		{
			g := genesisStore.Genesis()
			gHeader := genesis.Header{
				GenesisID:   g.GenesisID,
				NetworkID:   g.NetworkID,
				NetworkName: g.NetworkName,
			}
			var matched *GenesisTemplate
			for i := range AllowedOperaGenesis {
				allowed := &AllowedOperaGenesis[i]
				if allowed.Hashes.Equal(genesisHashes) && allowed.Header.Equal(gHeader) {
					log.Info("Genesis file is a known preset", "name", allowed.Name)
					matched = allowed
					break
				}
			}
			switch {
			case matched != nil && matched.SupersededBy != "":
				firstLaunchPending := integration.FirstLaunchPending(path.Join(dataDir, "chaindata"))
				if err := checkGenesisPresetFreshness(*matched, firstLaunchPending); err != nil {
					utils.Fatalf("%v", err)
				}
				log.Warn("Genesis preset is superseded — this datadir already carries chain state, which is verified separately against the binary's upgrade activation seals; fresh installs must use the replacement",
					"name", matched.Name, "replacement", matched.SupersededBy)
			case matched != nil:
				// current trusted preset
			case ctx.GlobalBool(ExperimentalGenesisFlag.Name):
				log.Warn("SECURITY WARNING: Genesis file doesn't refer to any trusted preset — node may join a different network")
			default:
				utils.Fatalf("Genesis file doesn't refer to any trusted preset. Enable experimental genesis with --genesis.allowExperimental")
			}
		}
		return genesisStore
	}
	return nil
}

func setBootnodes(ctx *cli.Context, urls []string, cfg *node.Config) {
	cfg.P2P.BootstrapNodesV5 = []*enode.Node{}
	for _, url := range urls {
		if url != "" {
			node, err := enode.Parse(enode.ValidSchemes, url)
			if err != nil {
				log.Error("Bootstrap URL invalid", "enode", url, "err", err)
				continue
			}
			cfg.P2P.BootstrapNodesV5 = append(cfg.P2P.BootstrapNodesV5, node)
		}
	}
	cfg.P2P.BootstrapNodes = cfg.P2P.BootstrapNodesV5
}

func setDataDir(ctx *cli.Context, cfg *node.Config) {
	defaultDataDir := DefaultDataDir()

	switch {
	case ctx.GlobalIsSet(DataDirFlag.Name):
		cfg.DataDir = ctx.GlobalString(DataDirFlag.Name)
	case ctx.GlobalIsSet(FakeNetFlag.Name):
		_, num, err := parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
		if err != nil {
			log.Crit("Invalid flag", "flag", FakeNetFlag.Name, "err", err)
		}
		cfg.DataDir = filepath.Join(defaultDataDir, fmt.Sprintf("fakenet-%d", num))
	}
}

func setTxPool(ctx *cli.Context, cfg *evmcore.TxPoolConfig) {
	if ctx.GlobalIsSet(utils.TxPoolLocalsFlag.Name) {
		locals := strings.Split(ctx.GlobalString(utils.TxPoolLocalsFlag.Name), ",")
		for _, account := range locals {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				utils.Fatalf("Invalid account in --txpool.locals: %s", trimmed)
			} else {
				cfg.Locals = append(cfg.Locals, common.HexToAddress(account))
			}
		}
	}
	if ctx.GlobalIsSet(utils.TxPoolNoLocalsFlag.Name) {
		cfg.NoLocals = ctx.GlobalBool(utils.TxPoolNoLocalsFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolJournalFlag.Name) {
		cfg.Journal = ctx.GlobalString(utils.TxPoolJournalFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolRejournalFlag.Name) {
		cfg.Rejournal = ctx.GlobalDuration(utils.TxPoolRejournalFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolPriceLimitFlag.Name) {
		cfg.PriceLimit = ctx.GlobalUint64(utils.TxPoolPriceLimitFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolPriceBumpFlag.Name) {
		cfg.PriceBump = ctx.GlobalUint64(utils.TxPoolPriceBumpFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolAccountSlotsFlag.Name) {
		cfg.AccountSlots = ctx.GlobalUint64(utils.TxPoolAccountSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolGlobalSlotsFlag.Name) {
		cfg.GlobalSlots = ctx.GlobalUint64(utils.TxPoolGlobalSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolAccountQueueFlag.Name) {
		cfg.AccountQueue = ctx.GlobalUint64(utils.TxPoolAccountQueueFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolGlobalQueueFlag.Name) {
		cfg.GlobalQueue = ctx.GlobalUint64(utils.TxPoolGlobalQueueFlag.Name)
	}
	if ctx.GlobalIsSet(utils.TxPoolLifetimeFlag.Name) {
		cfg.Lifetime = ctx.GlobalDuration(utils.TxPoolLifetimeFlag.Name)
	}
}

func gossipConfigWithFlags(ctx *cli.Context, src gossip.Config) (gossip.Config, error) {
	cfg := src

	if ctx.GlobalIsSet(RPCGlobalGasCapFlag.Name) {
		cfg.RPCGasCap = ctx.GlobalUint64(RPCGlobalGasCapFlag.Name)
	}
	if ctx.GlobalIsSet(RPCGlobalTxFeeCapFlag.Name) {
		cfg.RPCTxFeeCap = ctx.GlobalFloat64(RPCGlobalTxFeeCapFlag.Name)
	}
	if ctx.GlobalIsSet(RPCGlobalTimeoutFlag.Name) {
		cfg.RPCTimeout = ctx.GlobalDuration(RPCGlobalTimeoutFlag.Name)
	}
	if ctx.GlobalIsSet(RPCAllowUnprotectedTxsFlag.Name) {
		cfg.AllowUnprotectedTxs = ctx.GlobalBool(RPCAllowUnprotectedTxsFlag.Name)
	}
	if ctx.GlobalIsSet(SyncModeFlag.Name) {
		if syncmode := ctx.GlobalString(SyncModeFlag.Name); syncmode != "full" && syncmode != "snap" {
			utils.Fatalf("--%s must be either 'full' or 'snap'", SyncModeFlag.Name)
		}
		cfg.AllowSnapsync = ctx.GlobalString(SyncModeFlag.Name) == "snap"
	}

	return cfg, nil
}

func gossipStoreConfigWithFlags(ctx *cli.Context, src gossip.StoreConfig) (gossip.StoreConfig, error) {
	cfg := src
	if ctx.GlobalIsSet(utils.GCModeFlag.Name) {
		if gcmode := ctx.GlobalString(utils.GCModeFlag.Name); gcmode != "light" && gcmode != "full" && gcmode != "archive" {
			utils.Fatalf("--%s must be 'light', 'full' or 'archive'", GCModeFlag.Name)
		}
		cfg.EVM.Cache.TrieDirtyDisabled = ctx.GlobalString(utils.GCModeFlag.Name) == "archive"
		cfg.EVM.Cache.GreedyGC = ctx.GlobalString(utils.GCModeFlag.Name) == "full"
	}
	return cfg, nil
}

func setDBConfig(ctx *cli.Context, cfg integration.DBsConfig, cacheRatio cachescale.Func) integration.DBsConfig {
	if ctx.GlobalIsSet(DBPresetFlag.Name) {
		preset := ctx.GlobalString(DBPresetFlag.Name)
		switch preset {
		case "pbl-1":
			cfg = integration.Pbl1DBsConfig(cacheRatio.U64, uint64(utils.MakeDatabaseHandles()))
		case "ldb-1":
			cfg = integration.Ldb1DBsConfig(cacheRatio.U64, uint64(utils.MakeDatabaseHandles()))
		case "legacy-ldb":
			cfg = integration.LdbLegacyDBsConfig(cacheRatio.U64, uint64(utils.MakeDatabaseHandles()))
		case "legacy-pbl":
			cfg = integration.PblLegacyDBsConfig(cacheRatio.U64, uint64(utils.MakeDatabaseHandles()))
		default:
			utils.Fatalf("--%s must be 'pbl-1', 'ldb-1', 'legacy-pbl' or 'legacy-ldb'", DBPresetFlag.Name)
		}
	}
	if ctx.GlobalIsSet(DBMigrationModeFlag.Name) {
		cfg.MigrationMode = ctx.GlobalString(DBMigrationModeFlag.Name)
		switch cfg.MigrationMode {
		case "reformat", "rebuild", "":
		default:
			utils.Fatalf("--%s must be 'reformat' or 'rebuild'", DBMigrationModeFlag.Name)
		}
	}
	return cfg
}

func nodeConfigWithFlags(ctx *cli.Context, cfg node.Config) node.Config {
	utils.SetNodeConfig(ctx, &cfg)

	setDataDir(ctx, &cfg)
	return cfg
}

func cacheScaler(ctx *cli.Context) cachescale.Func {
	if !ctx.GlobalIsSet(CacheFlag.Name) {
		return cachescale.Identity
	}
	targetCache := ctx.GlobalInt(CacheFlag.Name)
	baseSize := DefaultCacheSize
	if targetCache < baseSize {
		log.Crit("Invalid flag", "flag", CacheFlag.Name, "err", fmt.Sprintf("minimum cache size is %d MB", baseSize))
	}
	return cachescale.Ratio{
		Base:   uint64(baseSize - ConstantCacheSize),
		Target: uint64(targetCache - ConstantCacheSize),
	}
}

func mayMakeAllConfigs(ctx *cli.Context) (*config, error) {
	// Defaults (low priority)
	cacheRatio := cacheScaler(ctx)
	cfg := config{
		Node:          defaultNodeConfig(),
		Opera:         gossip.DefaultConfig(cacheRatio),
		Emitter:       emitter.DefaultConfig(),
		TxPool:        evmcore.DefaultTxPoolConfig,
		OperaStore:    gossip.DefaultStoreConfig(cacheRatio),
		Lachesis:      abft.DefaultConfig(),
		LachesisStore: abft.DefaultStoreConfig(cacheRatio),
		VectorClock:   vecmt.DefaultConfig(cacheRatio),
	}

	if ctx.GlobalIsSet(FakeNetFlag.Name) {
		_, num, err := parseFakeGen(ctx.GlobalString(FakeNetFlag.Name))
		if err != nil {
			return nil, fmt.Errorf("invalid --fakenet flag %q: %w", ctx.GlobalString(FakeNetFlag.Name), err)
		}
		cfg.Emitter = emitter.FakeConfig(num)
		setBootnodes(ctx, []string{}, &cfg.Node)
	} else {
		// "asDefault" means set network defaults
		cfg.Node.P2P.BootstrapNodes = asDefault
		cfg.Node.P2P.BootstrapNodesV5 = asDefault
	}

	if ctx.GlobalBool(FastEmitFlag.Name) {
		log.Info("Fast emit mode enabled")
		cfg.Emitter = emitter.FastEmitConfig()
	}

	// Load config file (medium priority)
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadAllConfigs(file, &cfg); err != nil {
			return &cfg, err
		}
	}
	// apply default for DB config if it wasn't touched by config file
	dbDefault := integration.DefaultDBsConfig(cacheRatio.U64, uint64(utils.MakeDatabaseHandles()))
	if len(cfg.DBs.Routing.Table) == 0 {
		cfg.DBs.Routing = dbDefault.Routing
	}
	if len(cfg.DBs.GenesisCache.Table) == 0 {
		cfg.DBs.GenesisCache = dbDefault.GenesisCache
	}
	if len(cfg.DBs.RuntimeCache.Table) == 0 {
		cfg.DBs.RuntimeCache = dbDefault.RuntimeCache
	}

	// Apply flags (high priority)
	var err error
	cfg.Opera, err = gossipConfigWithFlags(ctx, cfg.Opera)
	if err != nil {
		return nil, err
	}
	cfg.OperaStore, err = gossipStoreConfigWithFlags(ctx, cfg.OperaStore)
	if err != nil {
		return nil, err
	}
	cfg.Node = nodeConfigWithFlags(ctx, cfg.Node)
	cfg.Node.MaxConcurrentRPC = cfg.Opera.MaxConcurrentRPC
	cfg.DBs = setDBConfig(ctx, cfg.DBs, cacheRatio)

	err = setValidator(ctx, &cfg.Emitter)
	if err != nil {
		return nil, err
	}
	if cfg.Emitter.Validator.ID != 0 && len(cfg.Emitter.PrevEmittedEventFile.Path) == 0 {
		cfg.Emitter.PrevEmittedEventFile.Path = cfg.Node.ResolvePath(path.Join("emitter", fmt.Sprintf("last-%d", cfg.Emitter.Validator.ID)))
	}
	setTxPool(ctx, &cfg.TxPool)

	if err := cfg.Opera.Validate(); err != nil {
		return nil, err
	}

	if ctx.GlobalIsSet(TraceNodeFlag.Name) {
		cfg.OperaStore.TraceTransactions = true
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *config) error {
	if cfg.Opera.RPCGasCap != 0 && cfg.Opera.RPCGasCap < 21000 {
		return fmt.Errorf("RPCGasCap %d is below minimum transaction gas (21000)", cfg.Opera.RPCGasCap)
	}
	if cfg.Opera.RPCTxFeeCap < 0 {
		return fmt.Errorf("RPCTxFeeCap must not be negative, got %f", cfg.Opera.RPCTxFeeCap)
	}
	if cfg.TxPool.PriceLimit == 0 {
		log.Warn("TxPool PriceLimit is 0, all transactions regardless of gas price will be accepted")
	}
	return nil
}

func makeAllConfigs(ctx *cli.Context) *config {
	cfg, err := mayMakeAllConfigs(ctx)
	if err != nil {
		utils.Fatalf("%v", err)
	}
	return cfg
}

func defaultNodeConfig() node.Config {
	cfg := NodeDefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit, gitDate)
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "vc", "dag", "abft", "web3")
	cfg.WSModules = append(cfg.WSModules, "eth", "vc", "dag", "abft", "web3")
	cfg.IPCPath = "opera.ipc"
	cfg.DataDir = DefaultDataDir()
	return cfg
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	cfg := makeAllConfigs(ctx)
	comment := ""

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}

	dump := os.Stdout
	if ctx.NArg() > 0 {
		dump, err = os.OpenFile(ctx.Args().Get(0), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer func() { _ = dump.Close() }()
	}
	_, _ = dump.WriteString(comment)
	_, _ = dump.Write(out)

	return nil
}

func checkConfig(ctx *cli.Context) error {
	_, err := mayMakeAllConfigs(ctx)
	return err
}
