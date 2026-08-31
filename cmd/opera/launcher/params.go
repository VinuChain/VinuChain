package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

const (
	vinuChainNetworkName               = "VinuChain"
	vinuChainTestnetNetworkName        = "VinuChain Testnet"
	vinuChainStagingMainnetNetworkName = "VinuChain Staging Mainnet"
	vinuChainMainnetNetworkName        = "VinuChain Mainnet"

	// testnetPostPatch10BootstrapPointer is what refused operators and
	// superseded presets are pointed at while this binary stages the not-yet-
	// sealed SfcV2Patch10: a fresh replay of ANY published genesis would
	// stage Patch10 at the wrong local seal, so the only safe bootstrap is a
	// post-Patch10 chaindata snapshot. The follow-up release (which pins
	// Patch10 and ships a regenerated genesis) replaces this with the new
	// genesis URL.
	testnetPostPatch10BootstrapPointer = "the current post-SfcV2Patch10 chaindata snapshot at https://vinu-blockchain-genesis.s3.amazonaws.com/chaindata-snapshots/"

	// mainnetPostElemontBootstrapPointer is what refused mainnet presets are
	// pointed at after the 2026-08-29 ELEMONT activation. Both distributed
	// mainnet genesis files pre-date SfcV2/Elemont/Shanghai/Cancun/Prague/
	// VinuBLS12381/VinuLatestEVM/PaybackV2, so a fresh replay under current
	// binary rules stages every one of them at the wrong local seal and
	// diverges with "wrong event epoch hash". Existing datadirs are
	// unaffected - SupersededBy only refuses FIRST launch.
	mainnetPostElemontBootstrapPointer = "the post-activation chaindata snapshot at https://vinu-blockchain-mainnet-genesis.s3.amazonaws.com/chaindata-snapshots/elemont-20260829/seal-5/"

	// Epoch and block at which the live testnet activated each SFC patch this
	// binary hardcodes. These are historical facts read back from the
	// published 2026-07-11 genesis' epoch history (sha256
	// a31e5100c0bf72deeab924ce0bed41955e2bc2b64b7a766827ee7603e68efaeb; its
	// top record is epoch 6119 with all three active), not policy. Re-verify
	// with TestCurrentTestnetGenesisSatisfiesStoredStateRequirement before
	// changing them — see StoredStateRequirements.
	//
	// The blocks are the first executed under each new rule set, which is
	// what a store records as UpgradeHeight.Height. They are deliberately
	// one past the deployment log's "activated at block" values (1,508,211 /
	// 1,529,200 / 1,529,442), which name the seal block itself.
	// Epoch and block at which live MAINNET activated each ELEMONT-era flag
	// during the 2026-08-29/30 window. Historical fact, not policy: the epochs
	// were observed from vc_getRules across the five seals, and each block is
	// the FIRST block executed under the new rule set — derived by binary
	// search on the epoch prefix encoded in every block id, and independently
	// corroborated for epoch 7892 by eth_config's activationBlock
	// (14,705,762). Note the seal-1 SetCode landed in block 14,701,167, which
	// is still epoch 7888; 14,701,168 is the first block under the new rules,
	// matching the "one past the deployment log" convention used above.
	mainnetElemontSeal1ActiveFromEpoch = idx.Epoch(7889)
	mainnetElemontSeal1ActiveFromBlock = idx.Block(14701168)
	mainnetCancunActiveFromEpoch       = idx.Epoch(7890)
	mainnetCancunActiveFromBlock       = idx.Block(14702730)
	mainnetPragueActiveFromEpoch       = idx.Epoch(7891)
	mainnetPragueActiveFromBlock       = idx.Block(14704227)
	mainnetBLS12381ActiveFromEpoch     = idx.Epoch(7892)
	mainnetBLS12381ActiveFromBlock     = idx.Block(14705762)
	mainnetLatestEVMActiveFromEpoch    = idx.Epoch(7893)
	mainnetLatestEVMActiveFromBlock    = idx.Block(14707397)

	testnetSfcV2Patch7ActiveFromEpoch = idx.Epoch(6017)
	testnetSfcV2Patch7ActiveFromBlock = idx.Block(1508212)
	testnetSfcV2Patch8ActiveFromEpoch = idx.Epoch(6118)
	testnetSfcV2Patch8ActiveFromBlock = idx.Block(1529201)
	testnetSfcV2Patch9ActiveFromEpoch = idx.Epoch(6119)
	testnetSfcV2Patch9ActiveFromBlock = idx.Block(1529443)
)

var (
	mainnetBootnodes = []string{
		"enode://0281626c7d7fc8696300688cbb19f3781aabd981d74cd16f3f5cd7885a32da4d1d9d64afbb2416b93654935a3088afbe1a4a05d823ff2146e5d1d0c2cbdeca46@188.165.195.122:3000",
	}

	testnetBootnodes = []string{
		"enode://e2a95c1b8d85b018b8e88133bec342801b42e19b59a52e030462d04a5549f02fc57215b4ca97771ec6b3a0d30a78603fdccd2b5091c44f6ac439d6c8be8bc539@44.239.129.39:3000",
		"enode://7a45d086b9c82bd3677a76d36e003b9490066d56b612f33d05cb4d242212acd4e5cab4abbcb15a0df9aa499e41b4b4e868d82ba1c509c1990c9217dfe4607775@44.239.129.39:3001",
		"enode://d8e37eeba79b2c52dcba6e396ff907f27a6a8f7db34528cb8636bc3271291657a01c5649bff53429cea8a23b03fac13a178813c34c6d17d14f7b810a988393b5@44.239.129.39:3002",
		"enode://3f15b5ac22dea3e37a90cd9378cf0cd4ed9ea122851846c8108fcc7d2c7e709ea4a089cf3da93c0d3d3053250417cf0ea9ad9eff0aa77ff07d76b6cf267a2937@44.239.129.39:3003",
	}

	Bootnodes = map[string][]string{
		"main":                             mainnetBootnodes,
		vinuChainNetworkName:               mainnetBootnodes,
		vinuChainMainnetNetworkName:        mainnetBootnodes,
		"test":                             testnetBootnodes,
		vinuChainTestnetNetworkName:        testnetBootnodes,
		vinuChainStagingMainnetNetworkName: {},
	}

	vinuChainTestnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xbf7a3d7f49cd99745acd2aa1c828c81576c41a84fddc9c6ffb9857bab02fe260"),
		NetworkID:   opera.VinuChainTestNetworkID,
		NetworkName: vinuChainTestnetNetworkName,
	}

	vinuChainTestMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xb1b0e08cb0d53d0fb1067658c5af0b3d3ff334d574679f5f74eee2b3448394ce"),
		NetworkID:   opera.VinuChainStagingNetworkID,
		NetworkName: vinuChainStagingMainnetNetworkName,
	}

	vinuChainMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xca7941e04fc93391af59a3a87e2ad386312d6b74922deeaa05068b1c08d9caa4"),
		NetworkID:   opera.VinuChainMainNetworkID,
		NetworkName: vinuChainMainnetNetworkName,
	}

	AllowedOperaGenesis = []GenesisTemplate{

		// Mainnet
		{
			Name:   "VinuChain mainnet without history",
			Header: vinuChainMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xc6640c1b62156a63c3121fba3aaf755ee88b84935d2ebf1497611e4ee7f09144"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x438be95bb65eee5e23d7f78d39773646f2f21a6b18266b4d73d1e723c55fb94e"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x765b90e4674d426b37d05a0e4a35addb3deec44a0cc948391b738e0e815682be"),
			},
			SupersededBy: mainnetPostElemontBootstrapPointer,
		},

		// Mainnet with deployed contracts
		{
			Name:   "VinuChain mainnet with deployed contracts",
			Header: vinuChainMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x482f104dc843b2f86265a3494b1047c65a8568b0578ef1c43ea9aa8c961e6a6f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x9aab452d91d99fe26457feac40c2be7f2b31facf8edf66d815e2b0a184b871de"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x53f30bbcc37b7ba4d705aad4e79b1e1007673d64b6a1ab703e2319776a62bb3d"),
			},
			SupersededBy: mainnetPostElemontBootstrapPointer,
		},

		// Mainnet regenerated after the ELEMONT activation (2026-08-29/30).
		// Exported from a canonical post-seal-5 node at epoch 7896 /
		// block 14,711,847, so it already carries SfcV2, Elemont,
		// ElemontPubkeyValidation, Shanghai, Cancun, Prague, VinuBLS12381,
		// VinuLatestEVM and PaybackV2 as sealed history rather than staging
		// them on replay. This is the only mainnet genesis a fresh install
		// under v2.0.49-elemont+ can replay without diverging.
		{
			Name:   "VinuChain mainnet post-ELEMONT (2026-08-30)",
			Header: vinuChainMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x674f552ca94f2762b5c1a311580835eabc3f3882a042ad505fa98878eb332f51"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x86429387208911fa3348d486abfb02ed3f6f1fcb49402e0567989725385569bd"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x601c8d6fbdc8d7fdd1a0de63787098bfa5ce1cd9b2ff1edd8099f6b4f03b4473"),
			},
		},

		// VinuChain testnet
		{
			Name:   "VinuChain testnet without history",
			Header: vinuChainTestnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x7a74f234769f2285be94ac48c4a97abf98de32b93a05ddfd6cc934027b6d2d4f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xbfe43b2d77e7d672c4b0130d0a43f0710704f53ebb2c39a379c93076a43bddce"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x7c3476d667f7912172df77a5e5804428380541bf98282442689e6b442d16da34"),
			},
			SupersededBy: testnetPostPatch10BootstrapPointer,
		},

		// VinuChain testnet with history through epoch 5637 / block 1,423,701 (2026-04-19)
		// Regenerated from the testnet RPC post-v2.0.8-elemont rollout; replaces the
		// 2024-06-21 genesis which pre-dated SfcV2 / SfcV2Patch / SfcV2Patch2 and
		// caused "wrong event epoch hash" divergence on fresh installs under current
		// binary rules. Distributed at:
		//   https://vinu-blockchain-genesis.s3.amazonaws.com/vitainu-genesis-testnet-20260419.g
		// STALE for fresh installs under v2.0.41+ rules: it pre-dates the
		// SfcV2Patch7/8/9 activations (epochs 6016/6117/6118), so a replay
		// re-stages them at the wrong seal and diverges. SupersededBy keeps
		// it trusted for already-initialized datadirs while refusing fresh
		// installs; use the 2026-07-11 genesis below for new installs.
		{
			Name:   "VinuChain testnet with history (2026-04-19)",
			Header: vinuChainTestnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x72f1b25236876c877f800fa50038870ef36eef0b4c6f3ba0b1d8b67c37b34c22"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xf11619ff578754ce5680982eecc805387dd7dde02fd1d59d43ff4d7ead231fa7"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x459360bfa1fce292e3f9e7c9ea204f91ca89040258126a6aa07c6c0c1e345624"),
			},
			SupersededBy: testnetPostPatch10BootstrapPointer,
		},

		// VinuChain testnet with history through epoch 6119 / block 1,529,442 (2026-07-11)
		// Regenerated post-SfcV2Patch9 (v2.0.44-elemont) from the published
		// testnet-chaindata-v2.0.44-elemont-20260708T163957Z-clean snapshot;
		// covers the SfcV2Patch7 (epoch 6016 seal), SfcV2Patch8 (6117), and
		// SfcV2Patch9 (6118) activations, so fresh installs under v2.0.44+
		// rules stage nothing during replay and stay on the live chain's
		// epoch-state hashes. Replaces the 2026-04-19 genesis, which forked
		// two community validators (17/18) crossing the Patch7 seal.
		// Distributed at:
		//   https://vinu-blockchain-genesis.s3.amazonaws.com/vitainu-genesis-testnet-20260711.g
		{
			Name:   "VinuChain testnet with history (2026-07-11)",
			Header: vinuChainTestnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xf7493a6a546a7cac432864c9ddd92f3d4cead6ecf9cecf416c84669683cde15b"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x21b5b18943bb2d7ccfc8fa6a09781146d8322a2dab5280db2d5ab87860f9ff0f"),
				genesisstore.EvmSection(0):    hash.HexToHash("0xeb82e4cf63b20c0655cc9514c3f5d87774d796c0b09afc01179f899bbdf4168b"),
			},
			// Fresh installs refused under this binary: the 2026-07-11 history
			// pre-dates SfcV2Patch10, so replay would stage it at the wrong seal.
			SupersededBy: testnetPostPatch10BootstrapPointer,
		},

		// VinuChain test mainnet
		{
			Name:   "VinuChain test mainnet without history",
			Header: vinuChainTestMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x4d9b7946e4c2afba6d015e4a4282dd8d2299c1494f13ad6908846e4f09ed11be"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xf51e8759171b4109bbd5d592d62a93d9d882cc0d2846323d20df0ef7b7cd27db"),
				genesisstore.EvmSection(0):    hash.HexToHash("0xe5f319e1c7c064c3f8f2a8226e1481b0102447478089e0602d40eda08055f893"),
			},
		},
	}

	// StoredStateRequirements pins, per public-network genesis lineage, the
	// upgrade activation history a datadir must agree with for this binary
	// to run it (see StoredStateRequirement + checkStoredChainState).
	// Matching is on GenesisID, which every published genesis file of a
	// network shares — including the stale ones — so the check holds for
	// restarts without --genesis and for stale datadirs pointed at the
	// current genesis file, and never matches a generated private network or
	// fakenet (content-derived GenesisID).
	//
	// Testnet: this binary hardcodes SfcV2Patch7/8/9 (live activations at
	// epochs 6017/6118/6119) AND SfcV2Patch10, which has NOT yet activated on
	// the live chain. A datadir that disagrees with the pinned history
	// activated flags at a local seal the live chain never performed — this
	// is how testnet validators 17 and 18 forked on 2026-06-21 — and a
	// datadir stopped below epoch 6118 would activate several at once at its
	// next seal and fork the same way. Both must restore a snapshot or
	// bootstrap from the current genesis.
	//
	// NOTE (SfcV2Patch10 rollout window): a node stopped inside epoch 6118
	// was resumable on the v2.0.44-46 binaries (only Patch9 remained). On
	// THIS binary it is NOT: it would co-stage Patch9+Patch10 at its local
	// 6118→6119 seal while the live chain sealed only Patch9 there, and fork
	// with "wrong event epoch hash". Such long-stopped datadirs must restore
	// the current post-Patch10 snapshot instead. Once Patch10's live
	// activation epoch is historical fact, add its UpgradeActivation entry
	// below (follow-up release, same rule as Patch7/8/9 — never in the
	// release that first stages it, which would refuse the pre-seal fleet).
	//
	// Mainnet and staging carry no requirement: their ELEMONT-era
	// activations have not rolled out yet, so a pre-activation datadir is
	// legitimately below every seal. Add an entry for those lineages as part
	// of the mainnet upgrade rollout, once the activation epochs are
	// historical fact — never at the release that first stages them, which
	// would refuse the whole pre-seal fleet.
	//
	// PaybackV2 is in that same bucket and is worth naming, because unlike
	// the SfcV2 patches it is not only a bytecode reflash: activation
	// rewrites Economy.QuotaCacheAddress into the sealed epoch state
	// (gossip/block_processor.go rebindPaybackV2), so a datadir that crosses
	// the edge locally at the wrong epoch diverges in persisted rules, not
	// just in contract code. The 2026-08-29 release is the one that first
	// stages it on mainnet, so it gets NO entry here. The follow-up release
	// pins all five mainnet activations — SfcV2, Elemont,
	// ElemontPubkeyValidation, PaybackV2 and the EVM forks — at the epochs
	// they actually sealed at, alongside testnet's SfcV2Patch10 pin.
	StoredStateRequirements = []StoredStateRequirement{
		{
			GenesisID:   vinuChainMainnetHeader.GenesisID,
			NetworkName: vinuChainMainnetNetworkName,
			Activations: []UpgradeActivation{
				{
					Name:            "SfcV2",
					ActiveFromEpoch: mainnetElemontSeal1ActiveFromEpoch,
					ActiveFromBlock: mainnetElemontSeal1ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.SfcV2 },
				},
				{
					Name:            "Elemont",
					ActiveFromEpoch: mainnetElemontSeal1ActiveFromEpoch,
					ActiveFromBlock: mainnetElemontSeal1ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.Elemont },
				},
				{
					Name:            "ElemontPubkeyValidation",
					ActiveFromEpoch: mainnetElemontSeal1ActiveFromEpoch,
					ActiveFromBlock: mainnetElemontSeal1ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.ElemontPubkeyValidation },
				},
				{
					Name:            "Shanghai",
					ActiveFromEpoch: mainnetElemontSeal1ActiveFromEpoch,
					ActiveFromBlock: mainnetElemontSeal1ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.Shanghai },
				},
				{
					Name:            "PaybackV2",
					ActiveFromEpoch: mainnetElemontSeal1ActiveFromEpoch,
					ActiveFromBlock: mainnetElemontSeal1ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.PaybackV2 },
				},
				{
					Name:            "Cancun",
					ActiveFromEpoch: mainnetCancunActiveFromEpoch,
					ActiveFromBlock: mainnetCancunActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.Cancun },
				},
				{
					Name:            "Prague",
					ActiveFromEpoch: mainnetPragueActiveFromEpoch,
					ActiveFromBlock: mainnetPragueActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.Prague },
				},
				{
					Name:            "VinuBLS12381",
					ActiveFromEpoch: mainnetBLS12381ActiveFromEpoch,
					ActiveFromBlock: mainnetBLS12381ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.VinuBLS12381 },
				},
				{
					Name:            "VinuLatestEVM",
					ActiveFromEpoch: mainnetLatestEVMActiveFromEpoch,
					ActiveFromBlock: mainnetLatestEVMActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.VinuLatestEVM },
				},
			},
			Bootstrap: mainnetPostElemontBootstrapPointer,
			// Every ELEMONT-era flag this binary hardcodes for mainnet is now
			// pinned historical fact — nothing is staged-but-unpinned — so the
			// usual newest-1 floor applies.
		},
		{
			GenesisID:   vinuChainTestnetHeader.GenesisID,
			NetworkName: vinuChainTestnetNetworkName,
			Activations: []UpgradeActivation{
				{
					Name:            "SfcV2Patch7",
					ActiveFromEpoch: testnetSfcV2Patch7ActiveFromEpoch,
					ActiveFromBlock: testnetSfcV2Patch7ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.SfcV2Patch7 },
				},
				{
					Name:            "SfcV2Patch8",
					ActiveFromEpoch: testnetSfcV2Patch8ActiveFromEpoch,
					ActiveFromBlock: testnetSfcV2Patch8ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.SfcV2Patch8 },
				},
				{
					Name:            "SfcV2Patch9",
					ActiveFromEpoch: testnetSfcV2Patch9ActiveFromEpoch,
					ActiveFromBlock: testnetSfcV2Patch9ActiveFromBlock,
					Active:          func(u opera.Upgrades) bool { return u.SfcV2Patch9 },
				},
			},
			Bootstrap: testnetPostPatch10BootstrapPointer,
			// This binary stages SfcV2Patch10, whose live activation epoch is
			// not historical fact yet. Floor the datadir at the newest pin
			// (6119, Patch9) so Patch10 is the only locally-staged upgrade —
			// a datadir inside epoch 6118 would co-stage Patch9+Patch10 at
			// one local seal and fork. Clear in the follow-up release that
			// pins Patch10's activation.
			StagesUnpinned: true,
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
