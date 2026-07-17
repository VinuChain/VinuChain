package launcher

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
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
// epoch hash"), while initialized datadirs are accepted only after their
// stored epoch rules prove that all three patch seals already activated.
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

// TestCheckStoredTestnetUpgradeSeals pins the initialized-datadir half of the
// stale-genesis guard. It is driven from the loaded store rather than a
// --genesis argument, so ordinary restarts and starts with the replacement
// genesis cannot bypass it.
func TestCheckStoredTestnetUpgradeSeals(t *testing.T) {
	assert := require.New(t)
	current := opera.VinuChainTestNetRules()
	publicGenesisID := vinuChainTestnetHeader.GenesisID
	privateGenesisID := hash.HexToHash("0x010203")
	canonicalHistory := map[idx.Epoch]*iblockproc.EpochState{
		6016: {Epoch: 6016, Rules: rulesWithTestnetPatches(false, false, false)},
		6017: {Epoch: 6017, Rules: rulesWithTestnetPatches(true, false, false)},
		6117: {Epoch: 6117, Rules: rulesWithTestnetPatches(true, false, false)},
		6118: {Epoch: 6118, Rules: rulesWithTestnetPatches(true, true, false)},
		6119: {Epoch: 6119, Rules: rulesWithTestnetPatches(true, true, true)},
	}
	canonicalBlocks := map[idx.Epoch]*iblockproc.BlockState{
		6017: {LastBlock: iblockproc.BlockCtx{Idx: 1508211}},
		6118: {LastBlock: iblockproc.BlockCtx{Idx: 1529200}},
		6119: {LastBlock: iblockproc.BlockCtx{Idx: 1529442}},
	}

	nonTestnet := &testnetUpgradeSealStoreStub{
		rules: opera.VinuChainMainNetRules(), epoch: 1, genesisID: publicGenesisID,
	}
	assert.NoError(checkStoredTestnetUpgradeSeals(nonTestnet))

	generatedPrivateNetwork := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 2, genesisID: privateGenesisID,
	}
	assert.NoError(checkStoredTestnetUpgradeSeals(generatedPrivateNetwork),
		"a private testnet generated by `opera network new` remains restartable")

	beforeMinimumEpoch := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6118, genesisID: publicGenesisID,
		history: canonicalHistory, historyBlocks: canonicalBlocks,
	}
	err := checkStoredTestnetUpgradeSeals(beforeMinimumEpoch)
	assert.Error(err)
	assert.Contains(err.Error(), "epoch 6119")

	for _, missing := range []struct {
		name     string
		upgrades opera.Upgrades
	}{
		{"SfcV2Patch7", opera.Upgrades{SfcV2Patch8: true, SfcV2Patch9: true}},
		{"SfcV2Patch8", opera.Upgrades{SfcV2Patch7: true, SfcV2Patch9: true}},
		{"SfcV2Patch9", opera.Upgrades{SfcV2Patch7: true, SfcV2Patch8: true}},
	} {
		t.Run(missing.name, func(t *testing.T) {
			require := require.New(t)
			rules := current
			rules.Upgrades = missing.upgrades
			store := &testnetUpgradeSealStoreStub{
				rules: rules, epoch: 6119, genesisID: publicGenesisID,
				history: canonicalHistory, historyBlocks: canonicalBlocks,
			}
			err := checkStoredTestnetUpgradeSeals(store)
			require.Error(err)
			require.Contains(err.Error(), missing.name)
		})
	}

	wrongPatch7Seal := map[idx.Epoch]*iblockproc.EpochState{
		6016: {Epoch: 6016, Rules: rulesWithTestnetPatches(false, false, false)},
		6017: {Epoch: 6017, Rules: rulesWithTestnetPatches(false, false, false)},
		6117: canonicalHistory[6117],
		6118: canonicalHistory[6118],
		6119: canonicalHistory[6119],
	}
	replayedSupersededStore := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6200, genesisID: publicGenesisID,
		history: wrongPatch7Seal, historyBlocks: canonicalBlocks,
	}
	err = checkStoredTestnetUpgradeSeals(replayedSupersededStore)
	assert.Error(err)
	assert.Contains(err.Error(), "SfcV2Patch7")
	assert.Contains(err.Error(), "epoch 6016")

	validPublicStore := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6119, genesisID: publicGenesisID,
		history: canonicalHistory, historyBlocks: canonicalBlocks,
	}
	assert.NoError(checkStoredTestnetUpgradeSeals(validPublicStore))

	wrongPatch7Block := map[idx.Epoch]*iblockproc.BlockState{
		6017: {LastBlock: iblockproc.BlockCtx{Idx: 1508212}},
		6118: canonicalBlocks[6118],
		6119: canonicalBlocks[6119],
	}
	wrongBlockHistoryStore := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6200, genesisID: publicGenesisID,
		history: canonicalHistory, historyBlocks: wrongPatch7Block,
	}
	err = checkStoredTestnetUpgradeSeals(wrongBlockHistoryStore)
	assert.Error(err)
	assert.Contains(err.Error(), "SfcV2Patch7")
	assert.Contains(err.Error(), "1508212")

	canonicalUpgradeHeights := []opera.UpgradeHeight{
		{Upgrades: rulesWithTestnetPatches(false, false, false).Upgrades, Height: 0},
		{Upgrades: rulesWithTestnetPatches(true, false, false).Upgrades, Height: 1508212},
		{Upgrades: rulesWithTestnetPatches(true, true, false).Upgrades, Height: 1529201},
		{Upgrades: rulesWithTestnetPatches(true, true, true).Upgrades, Height: 1529443},
	}
	validPrunedPublicStore := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6200, genesisID: publicGenesisID, upgradeHeights: canonicalUpgradeHeights,
	}
	assert.NoError(checkStoredTestnetUpgradeSeals(validPrunedPublicStore),
		"durable upgrade heights must keep valid pruned public stores restartable")

	wrongHeight := append([]opera.UpgradeHeight(nil), canonicalUpgradeHeights...)
	wrongHeight[1].Height++
	replayedAndPrunedStore := &testnetUpgradeSealStoreStub{
		rules: current, epoch: 6200, genesisID: publicGenesisID,
		history: canonicalHistory, historyBlocks: canonicalBlocks,
		upgradeHeights: wrongHeight,
	}
	err = checkStoredTestnetUpgradeSeals(replayedAndPrunedStore)
	assert.Error(err)
	assert.Contains(err.Error(), "SfcV2Patch7")
	assert.Contains(err.Error(), "1508212")
}

type testnetUpgradeSealStoreStub struct {
	rules          opera.Rules
	epoch          idx.Epoch
	genesisID      hash.Hash
	history        map[idx.Epoch]*iblockproc.EpochState
	historyBlocks  map[idx.Epoch]*iblockproc.BlockState
	upgradeHeights []opera.UpgradeHeight
}

func (s *testnetUpgradeSealStoreStub) GetRules() opera.Rules {
	return s.rules
}

func (s *testnetUpgradeSealStoreStub) GetEpoch() idx.Epoch {
	return s.epoch
}

func (s *testnetUpgradeSealStoreStub) GetGenesisID() *hash.Hash {
	return &s.genesisID
}

func (s *testnetUpgradeSealStoreStub) GetHistoryBlockEpochState(epoch idx.Epoch) (*iblockproc.BlockState, *iblockproc.EpochState) {
	return s.historyBlocks[epoch], s.history[epoch]
}

func (s *testnetUpgradeSealStoreStub) GetUpgradeHeights() []opera.UpgradeHeight {
	return s.upgradeHeights
}

func rulesWithTestnetPatches(patch7, patch8, patch9 bool) opera.Rules {
	rules := opera.VinuChainTestNetRules()
	rules.Upgrades.SfcV2Patch7 = patch7
	rules.Upgrades.SfcV2Patch8 = patch8
	rules.Upgrades.SfcV2Patch9 = patch9
	return rules
}
