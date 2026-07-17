package launcher

import (
	"os"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter/ier"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

// TestCurrentTestnetGenesisSatisfiesStoredStateRequirement verifies the
// stored-state guard against the REAL published genesis rather than against
// assumed numbers: applying the current testnet genesis must leave a datadir
// that checkStoredChainState accepts. If the requirement's MinEpoch were ever
// set above the genesis' top epoch, every fresh install would be refused by
// the very guard meant to protect it — this test is the tripwire for that.
//
// The genesis file is ~500 MB and is not committed, so the test is skipped
// unless its path is provided:
//
//	VINUCHAIN_TESTNET_GENESIS=/path/to/vitainu-genesis-testnet-20260711.g \
//	  go test ./cmd/opera/launcher/ -run TestCurrentTestnetGenesisSatisfiesStoredStateRequirement
func TestCurrentTestnetGenesisSatisfiesStoredStateRequirement(t *testing.T) {
	require := require.New(t)

	path := os.Getenv("VINUCHAIN_TESTNET_GENESIS")
	if path == "" {
		t.Skip("set VINUCHAIN_TESTNET_GENESIS to the published testnet genesis to run this check")
	}

	f, err := os.Open(path)
	require.NoError(err)
	defer func() { _ = f.Close() }()

	gs, _, err := genesisstore.OpenGenesisStore(f)
	require.NoError(err)
	defer func() { _ = gs.Close() }()

	g := gs.Genesis()
	require.Equal(vinuChainTestnetHeader.GenesisID, g.GenesisID,
		"the provided file must be a VinuChain testnet genesis")

	var req *StoredStateRequirement
	for i := range StoredStateRequirements {
		if StoredStateRequirements[i].GenesisID == g.GenesisID {
			req = &StoredStateRequirements[i]
			break
		}
	}
	require.NotNil(req, "the testnet lineage must carry a stored-state requirement")

	// The Epochs section is a one-shot stream, so gather everything in a
	// single pass: the top record (ApplyGenesis stores the first one it
	// iterates, which is what a fresh install would carry), the epoch at
	// which each tracked upgrade became active, and the UpgradeHeights a
	// fresh install ends up with (Store.ApplyGenesis feeds every historical
	// record to WriteUpgradeHeight, oldest first).
	type epochRec struct {
		epoch     idx.Epoch
		lastBlock idx.Block
		upgrades  opera.Upgrades
	}
	var (
		topEr    *ier.LlrIdxFullEpochRecord
		history  []epochRec
		observed = make(map[string]idx.Epoch, len(req.Activations))
	)
	g.Epochs.ForEach(func(er ier.LlrIdxFullEpochRecord) bool {
		if topEr == nil {
			rec := er
			topEr = &rec
		}
		history = append(history, epochRec{er.EpochState.Epoch, er.BlockState.LastBlock.Idx,
			er.EpochState.Rules.Upgrades})
		for _, a := range req.Activations {
			if !a.Active(er.EpochState.Rules.Upgrades) {
				continue
			}
			if got, ok := observed[a.Name]; !ok || er.EpochState.Epoch < got {
				observed[a.Name] = er.EpochState.Epoch
			}
		}
		return true
	})
	require.NotNil(topEr, "genesis carries no epoch records")

	// Records stream newest-first; ApplyGenesis writes heights oldest-first.
	for i, j := 0, len(history)-1; i < j; i, j = i+1, j-1 {
		history[i], history[j] = history[j], history[i]
	}

	// The derived activation epochs are only meaningful if the history
	// actually predates every tracked activation — otherwise the oldest
	// record would masquerade as the activation point. Assert it rather than
	// assume it.
	var earliestTracked idx.Epoch
	for _, a := range req.Activations {
		if earliestTracked == 0 || a.ActiveFromEpoch < earliestTracked {
			earliestTracked = a.ActiveFromEpoch
		}
	}
	require.Less(history[0].epoch, earliestTracked,
		"genesis history starts at epoch %d, at or after the earliest tracked activation (%d): the derived "+
			"activation epochs would be an artifact of where the history begins", history[0].epoch, earliestTracked)
	require.False(req.Activations[0].Active(history[0].upgrades),
		"the oldest genesis record already has %s active, so its activation epoch cannot be derived",
		req.Activations[0].Name)
	var (
		freshHeights []opera.UpgradeHeight
		prev         *opera.Upgrades
	)
	for _, r := range history {
		if prev == nil || r.upgrades != *prev {
			freshHeights = append(freshHeights, opera.UpgradeHeight{Upgrades: r.upgrades, Height: r.lastBlock + 1})
			u := r.upgrades
			prev = &u
		}
	}

	storedEpoch := topEr.EpochState.Epoch
	storedUpgrades := topEr.EpochState.Rules.Upgrades
	t.Logf("fresh install from %s would store epoch=%d SfcV2Patch7=%t SfcV2Patch8=%t SfcV2Patch9=%t",
		path, storedEpoch, storedUpgrades.SfcV2Patch7, storedUpgrades.SfcV2Patch8, storedUpgrades.SfcV2Patch9)
	require.Greater(storedEpoch, idx.Epoch(0))

	require.NoError(checkStoredChainState(&g.GenesisID, storedEpoch, storedUpgrades, freshHeights),
		"a fresh install from the current published genesis must not be refused by the stored-state guard")
	require.GreaterOrEqual(storedEpoch, req.minStoredEpoch(),
		"published genesis top epoch %d is below the requirement's floor %d — every fresh install would be refused",
		storedEpoch, req.minStoredEpoch())

	// Compare the hardcoded activation table against the history the genesis
	// actually carries. This is what turns those constants from asserted into
	// evidenced: if an activation epoch is wrong, the guard would refuse (or
	// wrongly accept) real datadirs, and this fails first.
	for _, a := range req.Activations {
		got, ok := observed[a.Name]
		require.True(ok, "%s is never active in the published genesis history, but the requirement tracks it", a.Name)
		gotBlock := recordedActivationHeight(freshHeights, a)
		t.Logf("%s first active at epoch %d block %d (table says epoch %d block %d)",
			a.Name, got, gotBlock, a.ActiveFromEpoch, a.ActiveFromBlock)
		require.Equal(a.ActiveFromEpoch, got,
			"%s activation epoch in the table (%d) disagrees with the published genesis history (%d)",
			a.Name, a.ActiveFromEpoch, got)
		require.Equal(a.ActiveFromBlock, gotBlock,
			"%s activation block in the table (%d) disagrees with the height a fresh install records (%d)",
			a.Name, a.ActiveFromBlock, gotBlock)
	}
}
