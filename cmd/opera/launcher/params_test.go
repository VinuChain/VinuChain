package launcher

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/stretchr/testify/require"

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
