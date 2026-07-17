package launcher

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
	futils "github.com/Fantom-foundation/go-opera/utils"
)

func TestDefaultBootnodesUseNetworkNames(t *testing.T) {
	require := require.New(t)

	require.Equal(vinuChainTestnetNetworkName, vinuChainTestnetHeader.NetworkName)
	require.Equal(vinuChainStagingMainnetNetworkName, vinuChainTestMainnetHeader.NetworkName)
	require.Equal(vinuChainMainnetNetworkName, vinuChainMainnetHeader.NetworkName)

	require.Equal(Bootnodes["test"], Bootnodes[vinuChainTestnetHeader.NetworkName])
	require.NotEmpty(Bootnodes[vinuChainTestnetHeader.NetworkName])

	require.Equal(Bootnodes["main"], Bootnodes[vinuChainNetworkName])
	require.NotEmpty(Bootnodes[vinuChainNetworkName])

	require.Equal(Bootnodes["main"], Bootnodes[vinuChainMainnetHeader.NetworkName])
	require.NotEmpty(Bootnodes[vinuChainMainnetHeader.NetworkName])

	_, ok := Bootnodes[vinuChainTestMainnetHeader.NetworkName]
	require.True(ok)
}

// TestAllowedOperaGenesisTestnet20260711Preset pins the post-SfcV2Patch9
// regenerated testnet genesis (history through epoch 6119 / block 1,529,442)
// as a trusted preset, so fresh installs under v2.0.44+ rules can bootstrap
// without --genesis.allowExperimental and without re-staging Patch7/8/9 at a
// wrong replay seal.
func TestAllowedOperaGenesisTestnet20260711Preset(t *testing.T) {
	require := require.New(t)

	var preset *GenesisTemplate
	for i := range AllowedOperaGenesis {
		if AllowedOperaGenesis[i].Name == "VinuChain testnet with history (2026-07-11)" {
			preset = &AllowedOperaGenesis[i]
			break
		}
	}
	require.NotNil(preset, "post-SfcV2Patch9 testnet genesis preset missing from AllowedOperaGenesis")

	require.Equal(vinuChainTestnetHeader, preset.Header)
	require.Equal(hash.HexToHash("0xf7493a6a546a7cac432864c9ddd92f3d4cead6ecf9cecf416c84669683cde15b"),
		preset.Hashes[genesisstore.EpochsSection(0)])
	require.Equal(hash.HexToHash("0x21b5b18943bb2d7ccfc8fa6a09781146d8322a2dab5280db2d5ab87860f9ff0f"),
		preset.Hashes[genesisstore.BlocksSection(0)])
	require.Equal(hash.HexToHash("0xeb82e4cf63b20c0655cc9514c3f5d87774d796c0b09afc01179f899bbdf4168b"),
		preset.Hashes[genesisstore.EvmSection(0)])
	require.Empty(preset.SupersededBy, "the current testnet genesis must not be marked superseded")
}

// TestStaleTestnetGenesisPresetsAreSuperseded pins that every testnet genesis
// preset pre-dating the SfcV2Patch7/8/9 activations is marked superseded, so
// mayGetGenesisStore refuses to initialize a FRESH datadir from it (a replay
// re-stages the patches at the wrong seal and diverges with "wrong event
// epoch hash"), while nodes with already-initialized datadirs stay accepted.
func TestStaleTestnetGenesisPresetsAreSuperseded(t *testing.T) {
	require := require.New(t)

	stale := map[string]bool{
		"VinuChain testnet without history":           false,
		"VinuChain testnet with history (2026-04-19)": false,
	}
	for i := range AllowedOperaGenesis {
		preset := &AllowedOperaGenesis[i]
		if _, ok := stale[preset.Name]; ok {
			stale[preset.Name] = true
			require.NotEmpty(preset.SupersededBy, "stale preset %q must be marked superseded", preset.Name)
		} else {
			require.Empty(preset.SupersededBy, "preset %q must not be marked superseded", preset.Name)
		}
	}
	for name, seen := range stale {
		require.True(seen, "expected stale preset %q in AllowedOperaGenesis", name)
	}
}

// TestStoredStateRequirementsTestnet pins the datadir-keyed half of the
// stale-state guard: the public testnet genesis lineage (matched by the
// persisted GenesisID, which every published testnet genesis file shares)
// must require the SfcV2Patch7/8/9 activations covered by the 2026-07-11
// genesis AND the first epoch at which all of them are live (6119). The
// requirement is keyed on GenesisID — not NetworkID — so generated private
// networks (`opera network new`, content-derived GenesisID) and fakenets are
// unaffected. Mainnet and staging lineages must have no requirement until
// their own activation rollouts ship one.
func TestStoredStateRequirementsTestnet(t *testing.T) {
	require := require.New(t)

	var testnetReq *StoredStateRequirement
	for i := range StoredStateRequirements {
		req := &StoredStateRequirements[i]
		require.NotEqual(vinuChainMainnetHeader.GenesisID, req.GenesisID,
			"mainnet lineage must not carry a stored-state requirement yet")
		require.NotEqual(vinuChainTestMainnetHeader.GenesisID, req.GenesisID,
			"staging lineage must not carry a stored-state requirement yet")
		if req.GenesisID == vinuChainTestnetHeader.GenesisID {
			testnetReq = req
		}
	}
	require.NotNil(testnetReq, "public testnet lineage must carry a stored-state requirement")
	require.NotEmpty(testnetReq.Bootstrap)

	// The activation epochs are the live chain's history, read back from the
	// published genesis (see the constants' comment and
	// TestCurrentTestnetGenesisSatisfiesStoredStateRequirement).
	activations := map[string]idx.Epoch{}
	blocks := map[string]idx.Block{}
	for _, a := range testnetReq.Activations {
		require.NotNil(a.Active, "activation %q needs a flag getter", a.Name)
		activations[a.Name] = a.ActiveFromEpoch
		blocks[a.Name] = a.ActiveFromBlock
	}
	require.Equal(map[string]idx.Epoch{
		"SfcV2Patch7": 6017,
		"SfcV2Patch8": 6118,
		"SfcV2Patch9": 6119,
	}, activations)
	require.Equal(map[string]idx.Block{
		"SfcV2Patch7": 1508212,
		"SfcV2Patch8": 1529201,
		"SfcV2Patch9": 1529443,
	}, blocks)

	// Only the newest activation may be pending, so a node inside epoch 6118
	// (SfcV2Patch9 staged, activating at the canonical 6118→6119 seal) is
	// still resumable.
	require.Equal(idx.Epoch(6118), testnetReq.minStoredEpoch())

	// Every activation the requirement tracks must be one this binary
	// actually hardcodes — otherwise the guard demands state the binary
	// would never produce.
	hardcoded := opera.VinuChainTestNetRules().Upgrades
	for _, a := range testnetReq.Activations {
		require.True(a.Active(hardcoded),
			"requirement tracks %q but the binary's testnet rules do not activate it", a.Name)
	}

	// Every superseded testnet preset must point fresh installs at the same
	// replacement genesis the stored-state requirement names.
	for i := range AllowedOperaGenesis {
		preset := &AllowedOperaGenesis[i]
		if preset.SupersededBy != "" {
			require.Equal(testnetReq.Bootstrap, preset.SupersededBy,
				"superseded preset %q must name the requirement's bootstrap genesis", preset.Name)
		}
	}
}

// TestGeneratedNetworkGenesisIDIsNotPublicTestnet pins the carve-out that
// keeps `opera network new` (and fakenet) datadirs runnable: a generated
// network uses testnet RULES and starts at epoch 2 — far below the public
// testnet's activation seals — but its genesis is content-derived, so its
// GenesisID never collides with the published testnet lineage and
// checkStoredChainState never matches it. If a future refactor keyed the
// requirement on NetworkID instead, generated private networks would be
// refused on restart; this test fails first.
func TestGeneratedNetworkGenesisIDIsNotPublicTestnet(t *testing.T) {
	require := require.New(t)

	rules := opera.VinuChainTestNetRules()
	require.EqualValues(opera.VinuChainTestNetworkID, rules.NetworkID,
		"generated networks inherit the testnet NetworkID — the reason GenesisID keying matters")

	gs := makefakegenesis.FakeGenesisStoreWithRulesAndStart(
		1, futils.ToVC(1000), futils.ToVC(10), rules, idx.Epoch(2), idx.Block(1))
	defer gs.Close()
	generatedID := gs.Genesis().GenesisID

	require.NotEqual(vinuChainTestnetHeader.GenesisID, generatedID,
		"a generated network must not claim the published testnet GenesisID")
	require.NoError(checkStoredChainState(&generatedID, 2, rules.Upgrades, nil),
		"a freshly generated private network must restart at its start epoch")
}

// TestCheckStoredChainState pins the boot-time enforcement: a datadir whose
// persisted GenesisID belongs to a known public-network lineage is refused
// when its stored epoch/rules have not genuinely crossed the lineage's
// historical activation seals. Stored upgrade flags alone are not enough: a
// datadir that replayed a stale genesis under a v2.0.41+ binary has all
// flags set — re-staged at a wrong local seal — while being stuck far below
// the live seals, already diverged with "wrong event epoch hash".
func TestCheckStoredChainState(t *testing.T) {
	require := require.New(t)

	testnetID := vinuChainTestnetHeader.GenesisID
	allActive := opera.Upgrades{SfcV2Patch7: true, SfcV2Patch8: true, SfcV2Patch9: true}
	preSealPatch9 := opera.Upgrades{SfcV2Patch7: true, SfcV2Patch8: true}

	// The activation heights the live chain recorded, as a healthy datadir
	// carries them (first block under each new rule set).
	liveHeights := []opera.UpgradeHeight{
		{Upgrades: opera.Upgrades{}, Height: 2},
		{Upgrades: opera.Upgrades{SfcV2Patch7: true}, Height: 1508212},
		{Upgrades: preSealPatch9, Height: 1529201},
		{Upgrades: allActive, Height: 1529443},
	}
	preSealHeights := liveHeights[:3]

	// No genesis ID stored yet — nothing to match.
	require.NoError(checkStoredChainState(nil, 0, opera.Upgrades{}, nil))

	// Generated private network (`opera network new`): testnet rules and a
	// low epoch, but a content-derived GenesisID no requirement matches —
	// it must restart freely.
	privateID := hash.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	require.NoError(checkStoredChainState(&privateID, 2, allActive, nil))

	// Healthy public-testnet datadirs: a fresh install from the current
	// genesis (epoch 6119) and a live fleet node.
	require.NoError(checkStoredChainState(&testnetID, 6119, allActive, liveHeights))
	require.NoError(checkStoredChainState(&testnetID, 6200, allActive, liveHeights))

	// Resumable boundary: a node stopped inside epoch 6118 has Patch7/8
	// sealed and only Patch9 left to stage, which activates at the canonical
	// 6118→6119 seal. It must NOT be forced into a needless bootstrap.
	require.NoError(checkStoredChainState(&testnetID, 6118, preSealPatch9, preSealHeights),
		"a node at the Patch9 activation boundary activates it at the canonical seal and must be resumable")

	// Same epoch, but Patch9 already active: the live chain did not have it
	// at 6118, so this datadir sealed it early and is forked.
	err := checkStoredChainState(&testnetID, 6118, allActive, liveHeights)
	require.Error(err, "Patch9 active at epoch 6118 contradicts the live chain's history")
	require.Contains(err.Error(), "SfcV2Patch9")
	require.Contains(err.Error(), "vitainu-genesis-testnet-20260711.g")

	// Stopped one epoch too early: Patch8 AND Patch9 would both activate at
	// the 6117→6118 seal, but the live chain activated only Patch8 there.
	require.Error(checkStoredChainState(&testnetID, 6117, opera.Upgrades{SfcV2Patch7: true},
		liveHeights[:2]), "a node below the boundary would activate several upgrades at one seal and fork")

	// Legitimately stranded long before the seals (flags match history at
	// that epoch, but resuming would stage all three at once).
	err = checkStoredChainState(&testnetID, 5700, opera.Upgrades{}, liveHeights[:1])
	require.Error(err)
	require.Contains(err.Error(), "VinuChain Testnet")
	require.Contains(err.Error(), "6118")

	// Wrong-seal-bricked datadir: a stale-genesis replay under a
	// v2.0.41..v2.0.44 binary re-staged Patch7/8/9 at its first local seal,
	// so all flags are set at an epoch where the live chain had none.
	forkHeights := []opera.UpgradeHeight{
		{Upgrades: opera.Upgrades{}, Height: 2},
		{Upgrades: allActive, Height: 1423702},
	}
	err = checkStoredChainState(&testnetID, 5639, allActive, forkHeights)
	require.Error(err, "a wrong-seal replay must be refused")
	require.Contains(err.Error(), "SfcV2Patch7")

	// The same fork, kept sealing on its own until its epoch looks current.
	// Flags and epoch now match the live chain exactly, so only the recorded
	// activation blocks reveal that it never crossed the canonical seals.
	err = checkStoredChainState(&testnetID, 6200, allActive, forkHeights)
	require.Error(err, "a fork with current-looking flags must be caught by its recorded activation blocks")
	require.Contains(err.Error(), "1423702")
	require.Contains(err.Error(), "1508212")

	// A datadir with no recorded activation history at all cannot prove it
	// crossed the seals either.
	require.Error(checkStoredChainState(&testnetID, 6200, allActive, nil),
		"missing activation history must not pass as proof of a canonical seal")

	// Mainnet lineage carries no requirement yet: a pre-activation mainnet
	// datadir must keep booting under this binary.
	mainnetID := vinuChainMainnetHeader.GenesisID
	require.NoError(checkStoredChainState(&mainnetID, 100, opera.Upgrades{}, nil))
}

// TestCheckGenesisPresetFreshness pins the fresh-install guard: a superseded
// preset must be rejected when it is about to initialize an empty (or
// interrupted-genesis) datadir, and accepted in every other combination.
func TestCheckGenesisPresetFreshness(t *testing.T) {
	require := require.New(t)

	current := GenesisTemplate{Name: "current"}
	superseded := GenesisTemplate{
		Name:         "old",
		SupersededBy: "https://example.invalid/new.g",
	}

	require.NoError(checkGenesisPresetFreshness(current, false))
	require.NoError(checkGenesisPresetFreshness(current, true))
	require.NoError(checkGenesisPresetFreshness(superseded, false))

	err := checkGenesisPresetFreshness(superseded, true)
	require.Error(err)
	require.Contains(err.Error(), "old")
	require.Contains(err.Error(), "https://example.invalid/new.g")
}
