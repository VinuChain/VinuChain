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
	// Superseded under THIS binary: it stages the not-yet-sealed SfcV2Patch10,
	// so a fresh replay of the 2026-07-11 history would stage Patch10 at the
	// wrong local seal and diverge. Fresh installs are pointed at the
	// post-Patch10 snapshot until the follow-up release ships a regenerated
	// genesis (which becomes the new un-superseded current preset).
	require.Equal(testnetPostPatch10BootstrapPointer, preset.SupersededBy,
		"the 2026-07-11 genesis must be fresh-install-refused while SfcV2Patch10 is staged but unsealed")
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
		// Superseded by the 2026-08-29 ELEMONT activation: both distributed
		// mainnet genesis files pre-date SfcV2/Elemont/Shanghai/Cancun/
		// Prague/VinuBLS12381/VinuLatestEVM/PaybackV2, so a fresh replay
		// stages all of them at the wrong local seals and diverges.
		"VinuChain mainnet without history":         false,
		"VinuChain mainnet with deployed contracts": false,
		// Superseded as of the SfcV2Patch10 release: its history pre-dates
		// Patch10, so a fresh replay stages Patch10 at the wrong local seal.
		// Un-supersede in the follow-up release that ships a regenerated
		// post-Patch10 genesis.
		"VinuChain testnet with history (2026-07-11)": false,
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
// unaffected. The staging lineage must have no requirement until its own
// activation rollout ships one; mainnet's shipped with the 2026-08-29
// ELEMONT activation and is pinned by TestStoredStateRequirementsMainnet.
func TestStoredStateRequirementsTestnet(t *testing.T) {
	require := require.New(t)

	var testnetReq *StoredStateRequirement
	for i := range StoredStateRequirements {
		req := &StoredStateRequirements[i]
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
	// 6119, not 6118: this binary stages SfcV2Patch10, whose live activation
	// is not pinned yet (StagesUnpinned), so the floor is the newest pin
	// itself — a datadir inside epoch 6118 would co-stage Patch9+Patch10 at
	// one local seal and fork. Reverts to newest-1 semantics when the
	// follow-up release pins Patch10 and clears StagesUnpinned.
	require.Equal(idx.Epoch(6119), testnetReq.minStoredEpoch())

	// Every activation the requirement tracks must be one this binary
	// actually hardcodes — otherwise the guard demands state the binary
	// would never produce.
	hardcoded := opera.VinuChainTestNetRules().Upgrades
	for _, a := range testnetReq.Activations {
		require.True(a.Active(hardcoded),
			"requirement tracks %q but the binary's testnet rules do not activate it", a.Name)
	}

	// Every superseded preset must point fresh installs at the same
	// replacement genesis its OWN lineage's stored-state requirement names.
	// Matching on GenesisID matters now that both mainnet and testnet have
	// superseded presets with different bootstrap pointers — comparing every
	// preset against the testnet requirement would demand mainnet operators
	// bootstrap from a testnet snapshot.
	for i := range AllowedOperaGenesis {
		preset := &AllowedOperaGenesis[i]
		if preset.SupersededBy == "" {
			continue
		}
		var req *StoredStateRequirement
		for j := range StoredStateRequirements {
			if StoredStateRequirements[j].GenesisID == preset.Header.GenesisID {
				req = &StoredStateRequirements[j]
				break
			}
		}
		require.NotNil(req,
			"superseded preset %q has no stored-state requirement for its lineage", preset.Name)
		require.Equal(req.Bootstrap, preset.SupersededBy,
			"superseded preset %q must name its own requirement's bootstrap genesis", preset.Name)
	}
}

// TestStoredStateRequirementsMainnet pins the mainnet half of the stale-state
// guard, shipped with the 2026-08-29/30 ELEMONT activation. The epochs are
// observed fact (vc_getRules across the five seals) and each block is the
// FIRST executed under the new rule set, derived by binary search on the epoch
// prefix in every block id and corroborated for 7892 by eth_config's
// activationBlock. If a future release changes these, it is rewriting history.
func TestStoredStateRequirementsMainnet(t *testing.T) {
	require := require.New(t)

	var req *StoredStateRequirement
	for i := range StoredStateRequirements {
		if StoredStateRequirements[i].GenesisID == vinuChainMainnetHeader.GenesisID {
			req = &StoredStateRequirements[i]
			break
		}
	}
	require.NotNil(req, "mainnet lineage must carry a stored-state requirement after ELEMONT")

	epochs := map[string]idx.Epoch{}
	blocks := map[string]idx.Block{}
	for _, a := range req.Activations {
		epochs[a.Name] = a.ActiveFromEpoch
		blocks[a.Name] = a.ActiveFromBlock
	}
	require.Equal(map[string]idx.Epoch{
		"SfcV2": 7889, "Elemont": 7889, "ElemontPubkeyValidation": 7889,
		"Shanghai": 7889, "PaybackV2": 7889,
		"Cancun": 7890, "Prague": 7891,
		"VinuBLS12381": 7892, "VinuLatestEVM": 7893,
	}, epochs)
	require.Equal(map[string]idx.Block{
		"SfcV2": 14701168, "Elemont": 14701168, "ElemontPubkeyValidation": 14701168,
		"Shanghai": 14701168, "PaybackV2": 14701168,
		"Cancun": 14702730, "Prague": 14704227,
		"VinuBLS12381": 14705762, "VinuLatestEVM": 14707397,
	}, blocks)

	// Nothing is staged-but-unpinned on mainnet any more, so the floor is the
	// usual newest-1: a datadir inside 7892 still has only VinuLatestEVM left
	// to stage, at the canonical 7892->7893 seal.
	require.False(req.StagesUnpinned, "no mainnet upgrade is staged-but-unpinned after seal 5")
	require.Equal(idx.Epoch(7892), req.minStoredEpoch())

	// Every activation tracked must be one this binary actually hardcodes.
	hardcoded := opera.VinuChainMainNetRules().Upgrades
	for _, a := range req.Activations {
		require.True(a.Active(hardcoded),
			"requirement tracks %q but the binary's mainnet rules do not activate it", a.Name)
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
	defer func() { _ = gs.Close() }()
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

	// Patch9 activation boundary: on v2.0.44-46 a node stopped inside epoch
	// 6118 was resumable (only Patch9 remained). THIS binary also stages the
	// unpinned SfcV2Patch10, so the same datadir would co-stage Patch9+Patch10
	// at one local seal the live chain never performed — it must now be
	// refused (StagesUnpinned floors the datadir at the newest pin, 6119).
	err6118 := checkStoredChainState(&testnetID, 6118, preSealPatch9, preSealHeights)
	require.Error(err6118,
		"a datadir inside epoch 6118 would co-stage Patch9+Patch10 at one local seal under this binary and must be refused")
	require.Contains(err6118.Error(), "6119")

	// Same epoch, but Patch9 already active: the live chain did not have it
	// at 6118, so this datadir sealed it early and is forked.
	err := checkStoredChainState(&testnetID, 6118, allActive, liveHeights)
	require.Error(err, "Patch9 active at epoch 6118 contradicts the live chain's history")
	require.Contains(err.Error(), "SfcV2Patch9")
	require.Contains(err.Error(), "post-SfcV2Patch10 chaindata snapshot")

	// Stopped one epoch too early: Patch8 AND Patch9 would both activate at
	// the 6117→6118 seal, but the live chain activated only Patch8 there.
	require.Error(checkStoredChainState(&testnetID, 6117, opera.Upgrades{SfcV2Patch7: true},
		liveHeights[:2]), "a node below the boundary would activate several upgrades at one seal and fork")

	// Legitimately stranded long before the seals (flags match history at
	// that epoch, but resuming would stage all three at once).
	err = checkStoredChainState(&testnetID, 5700, opera.Upgrades{}, liveHeights[:1])
	require.Error(err)
	require.Contains(err.Error(), "VinuChain Testnet")
	// Floor is 6119 under this binary (StagesUnpinned: Patch10 staged but
	// unpinned), not the newest-1 = 6118 of v2.0.44-46.
	require.Contains(err.Error(), "6119")

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

	// Mainnet carries a requirement since the 2026-08-29 ELEMONT activation.
	// A pre-activation mainnet datadir must now be REFUSED: resuming it would
	// co-stage SfcV2/Elemont/Shanghai/Cancun/Prague/BLS12381/LatestEVM/
	// PaybackV2 at local seals the live chain never performed.
	mainnetID := vinuChainMainnetHeader.GenesisID
	err = checkStoredChainState(&mainnetID, 100, opera.Upgrades{}, nil)
	require.Error(err, "a pre-ELEMONT mainnet datadir must not resume under this binary")
	require.Contains(err.Error(), "7892")

	// A mainnet datadir one epoch below the newest pin (VinuLatestEVM at
	// 7893) is still resumable: only that newest upgrade is left to stage.
	// Its recorded history must be the real five-seal ladder — the guard
	// compares each tracked upgrade's recorded block against the requirement.
	mainnetFull := opera.VinuChainMainNetRules().Upgrades
	mainnetSeal1 := mainnetFull
	mainnetSeal1.Cancun, mainnetSeal1.Prague = false, false
	mainnetSeal1.VinuBLS12381, mainnetSeal1.VinuLatestEVM = false, false
	mainnetSeal2 := mainnetSeal1
	mainnetSeal2.Cancun = true
	mainnetSeal3 := mainnetSeal2
	mainnetSeal3.Prague = true
	mainnetSeal4 := mainnetSeal3
	mainnetSeal4.VinuBLS12381 = true // == every ELEMONT flag but VinuLatestEVM

	mainnetHeights := []opera.UpgradeHeight{
		{Upgrades: mainnetSeal1, Height: mainnetElemontSeal1ActiveFromBlock},
		{Upgrades: mainnetSeal2, Height: mainnetCancunActiveFromBlock},
		{Upgrades: mainnetSeal3, Height: mainnetPragueActiveFromBlock},
		{Upgrades: mainnetSeal4, Height: mainnetBLS12381ActiveFromBlock},
	}
	require.NoError(checkStoredChainState(&mainnetID, 7892, mainnetSeal4, mainnetHeights),
		"a mainnet datadir at the newest-1 floor must still resume")

	// A mainnet fork that re-staged every ELEMONT flag at one local seal must
	// be refused on its recorded activation block, even once its epoch and
	// flags look current.
	mainnetForkHeights := []opera.UpgradeHeight{
		{Upgrades: mainnetFull, Height: mainnetElemontSeal1ActiveFromBlock + 1},
	}
	err = checkStoredChainState(&mainnetID, 7900, mainnetFull, mainnetForkHeights)
	require.Error(err, "a mainnet wrong-seal replay must be refused")
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
