package gossip

import (
	"bytes"
	"context"
	"sync/atomic"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/gossip/emitter"
	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/utils"
	"github.com/Fantom-foundation/go-opera/valkeystore"
)

// newPreElemontTestEnv builds a testEnv whose stored epoch state has SfcV2,
// Podgorica, and Elemont DISABLED, but whose NetworkID matches a network
// recognised by opera.MainNetRulesForNetwork (staging) so that the runtime
// activation logic in newService fires.
//
// This simulates the binary-upgrade scenario for an existing chain that was
// genesised before the V2 upgrade flags were set in the binary rules.
func newPreElemontTestEnv(firstEpoch idx.Epoch, validatorsNum idx.Validator) *testEnv {
	rules := opera.VinuChainMainNetRules()
	rules.NetworkID = opera.VinuChainStagingNetworkID
	rules.Name = "VinuChain Staging"
	rules.Epochs.MaxEpochDuration = inter.Timestamp(maxEpochDuration)
	rules.Blocks.MaxEmptyBlockSkipPeriod = 0
	// Pre-Elemont stored epoch state: V2 flags off. The hardcoded staging rules
	// returned by MainNetRulesForNetwork have all three set to true, so the
	// runtime activation in service.go must propagate them.
	rules.Upgrades.SfcV2 = false
	rules.Upgrades.Podgorica = false
	rules.Upgrades.Elemont = false

	genStore := makefakegenesis.FakeGenesisStoreWithRulesAndStart(
		validatorsNum,
		utils.ToVC(genesisBalance),
		utils.ToVC(genesisStake),
		rules,
		firstEpoch,
		2,
	)
	genesis := genStore.Genesis()

	store := NewMemStore()
	if _, err := store.ApplyGenesis(genesis); err != nil {
		panic(err)
	}

	env := &testEnv{
		t:      store.GetGenesisTime().Time(),
		nonces: make(map[common.Address]uint64),
	}
	blockProc := DefaultBlockProc()
	blockProc.EventsModule = testConfirmedEventsModule{blockProc.EventsModule, env}

	engine, vecClock := makeTestEngine(store)

	txPool := &dummyTxPool{}
	var err error
	env.Service, err = newService(DefaultConfig(cachescale.Identity), store, blockProc, engine, vecClock, func(_ evmcore.StateReader) TxPool {
		return txPool
	})
	if err != nil {
		panic(err)
	}
	txPool.signer = env.EthAPI.signer
	if err := engine.Bootstrap(env.GetConsensusCallbacks()); err != nil {
		panic(err)
	}

	valKeystore := valkeystore.NewDefaultMemKeystore()
	env.signer = valkeystore.NewSigner(valKeystore)

	for i := idx.Validator(0); i < validatorsNum; i++ {
		cfg := emitter.DefaultConfig()
		vid := store.GetValidators().GetID(i)
		pubkey := store.GetEpochState().ValidatorProfiles[vid].PubKey
		cfg.Validator = emitter.ValidatorConfig{
			ID:     vid,
			PubKey: pubkey,
		}
		cfg.EmitIntervals = emitter.EmitIntervals{}
		cfg.MaxParents = idx.Event(validatorsNum/2 + 1)
		cfg.MaxTxsPerAddress = 10000000
		_ = valKeystore.Add(pubkey, crypto.FromECDSA(makefakegenesis.FakeKey(vid)), "fakepassword")
		_ = valKeystore.Unlock(pubkey, "fakepassword")
		world := env.EmitterWorld(env.signer)
		world.External = testEmitterWorldExternal{world.External, env}
		em := emitter.NewEmitter(cfg, world)
		env.RegisterEmitter(em)
		env.pubkeys = append(env.pubkeys, pubkey)
		em.Start()
	}

	_ = env.store.GenerateSnapshotAt(common.Hash(store.GetBlockState().FinalizedStateRoot), false)
	env.blockProcTasks.Start(1)
	env.verWatcher.Start()

	return env
}

// TestSfcV2BytecodeSwapAfterRuntimeActivation pins the runtime activation
// behavior of the SfcV2 upgrade flag. The chain starts with SfcV2 disabled
// in stored epoch state and an SFC contract running V1 bytecode. After the
// service activates the upgrade from binary rules and processes blocks past
// the first epoch seal, the on-chain SFC contract MUST be running V2
// bytecode. Without the fix the bytecode swap in block_processor.go is dead
// code on the binary-upgrade path because service.go pre-flips the stored
// es.Rules.Upgrades.SfcV2 flag, defeating the false-to-true transition guard.
func TestSfcV2BytecodeSwapAfterRuntimeActivation(t *testing.T) {
	logger.SetTestMode(t)

	// Count occurrences of the SFC V2 bytecode upgrade log line. The
	// activation contract requires it to fire exactly once across the test.
	var sfcV2LogCount int64
	prevHandler := log.Root().GetHandler()
	defer log.Root().SetHandler(prevHandler)
	log.Root().SetHandler(log.FuncHandler(func(r *log.Record) error {
		if r.Msg == "Applying SFC V2 bytecode upgrade" {
			atomic.AddInt64(&sfcV2LogCount, 1)
		}
		return prevHandler.Log(r)
	}))

	env := newPreElemontTestEnv(2, 3)
	defer env.Close()

	v1Bin := sfc.GetGenesisContractBin()
	v2Bin := sfc.GetContractBin()
	require.NotEqual(t, v1Bin, v2Bin, "test precondition: V1 and V2 SFC bytecodes must differ")

	// Sanity: at startup the SFC contract holds V1 bytecode (genesis did not
	// pre-install V2 because rules.Upgrades.SfcV2 was false at genesis).
	gotAtStart, err := env.CodeAt(context.TODO(), sfc.ContractAddress, nil)
	require.NoError(t, err)
	require.True(t, bytes.Equal(gotAtStart, v1Bin), "expected V1 SFC bytecode at startup, got %d bytes", len(gotAtStart))

	// Drive blocks across an epoch boundary so that the sealer fires and the
	// block_processor.go SfcV2 false→true transition has the chance to install
	// V2 bytecode.
	admin := idx.ValidatorID(1)
	other := idx.ValidatorID(2)
	_, err = env.ApplyTxs(nextEpoch, env.Transfer(admin, other, utils.ToVC(1)))
	require.NoError(t, err)
	_, err = env.ApplyTxs(nextEpoch, env.Transfer(admin, other, utils.ToVC(1)))
	require.NoError(t, err)

	gotAfterSeal, err := env.CodeAt(context.TODO(), sfc.ContractAddress, nil)
	require.NoError(t, err)
	require.True(t, bytes.Equal(gotAfterSeal, v2Bin),
		"expected V2 SFC bytecode (%d bytes) after epoch seal but got %d bytes; "+
			"the SFC bytecode upgrade in block_processor.go did not fire",
		len(v2Bin), len(gotAfterSeal))

	require.Equal(t, int64(1), atomic.LoadInt64(&sfcV2LogCount),
		"expected the SFC V2 bytecode upgrade log to fire exactly once")

	// Belt-and-suspenders: the bytecode swap is observable on the live state
	// only as long as the new rules are persisted alongside it. If a future
	// regression installed V2 bytecode but failed to persist the rules
	// transition, restart-time staging would re-stage SfcV2 forever.
	postSealEs := env.store.GetEpochState()
	require.True(t, postSealEs.Rules.Upgrades.SfcV2,
		"stored epoch state SfcV2 must be true after the seal")
	require.True(t, postSealEs.Rules.Upgrades.Podgorica,
		"stored epoch state Podgorica must be true after the seal")
	require.True(t, postSealEs.Rules.Upgrades.Elemont,
		"stored epoch state Elemont must be true after the seal")

	// FeeRefundActive must be flipped at seal time so the receipt encoder
	// emits FeeRefund for any post-Podgorica block from this point forward.
	require.True(t, types.FeeRefundActive.Load(),
		"FeeRefundActive must be true after the Podgorica activation seal")
}

// TestRuntimeActivationSurvivesGovernanceUpdate pins the staged-window
// invariant that an UpdateNetworkRules governance proposal landing during
// the staging window cannot strip the staged Upgrades flags. The driver
// listener routes the proposal through opera.UpdateRules, which protects
// the Upgrades field by restoring it from the source rules — and the source
// in this case is bs.DirtyRules (where staging put the new flags), not
// es.Rules. Without this guarantee, a malicious or accidental governance
// update during the activation window could silently revert the staged
// upgrades.
func TestRuntimeActivationSurvivesGovernanceUpdate(t *testing.T) {
	logger.SetTestMode(t)

	env := newPreElemontTestEnv(2, 3)
	defer env.Close()

	// Staging in newService must have written DirtyRules with all three
	// upgrade flags set to true.
	bs := env.store.GetBlockState()
	require.NotNil(t, bs.DirtyRules,
		"newService must have staged DirtyRules from hardcoded binary rules")
	require.True(t, bs.DirtyRules.Upgrades.SfcV2)
	require.True(t, bs.DirtyRules.Upgrades.Podgorica)
	require.True(t, bs.DirtyRules.Upgrades.Elemont)

	// Simulate a governance UpdateNetworkRules tx landing in the staged
	// window. opera.UpdateRules is the choke point used by the driver
	// listener (gossip/blockproc/drivermodule/driver_txs.go uses
	// bs.DirtyRules as the source when present). The diff tries to flip
	// SfcV2 off; UpdateRules must reject that by restoring Upgrades from src.
	diff := []byte(`{"Upgrades":{"SfcV2":false,"Podgorica":false,"Elemont":false}}`)
	updated, err := opera.UpdateRules(*bs.DirtyRules, diff)
	require.NoError(t, err)
	require.True(t, updated.Upgrades.SfcV2,
		"UpdateRules must preserve staged SfcV2 flag")
	require.True(t, updated.Upgrades.Podgorica,
		"UpdateRules must preserve staged Podgorica flag")
	require.True(t, updated.Upgrades.Elemont,
		"UpdateRules must preserve staged Elemont flag")

	// The full pipeline: drive blocks through the seal and assert the
	// bytecode swap still fires (the staged flags survived).
	v2Bin := sfc.GetContractBin()
	admin := idx.ValidatorID(1)
	other := idx.ValidatorID(2)
	_, err = env.ApplyTxs(nextEpoch, env.Transfer(admin, other, utils.ToVC(1)))
	require.NoError(t, err)
	_, err = env.ApplyTxs(nextEpoch, env.Transfer(admin, other, utils.ToVC(1)))
	require.NoError(t, err)

	gotAfterSeal, err := env.CodeAt(context.TODO(), sfc.ContractAddress, nil)
	require.NoError(t, err)
	require.True(t, bytes.Equal(gotAfterSeal, v2Bin),
		"SFC V2 bytecode swap must still fire after governance interleaving")

	postSealEs := env.store.GetEpochState()
	require.True(t, postSealEs.Rules.Upgrades.SfcV2,
		"stored epoch state SfcV2 must be true after the seal")
}

// TestRuntimeActivationIdempotentAcrossRestart pins the multi-restart
// idempotency of the staging logic. The first newService call stages
// SfcV2/Podgorica/Elemont into bs.DirtyRules. A second newService call on
// the same store must observe the already-staged flags and produce no
// duplicate "Staged ..." log lines and no further mutation of DirtyRules.
// Without this guarantee a regression that re-clears or re-stages on every
// restart would either churn the persisted block state or leak log noise.
func TestRuntimeActivationIdempotentAcrossRestart(t *testing.T) {
	logger.SetTestMode(t)

	var stagedLogCount int64
	prevHandler := log.Root().GetHandler()
	defer log.Root().SetHandler(prevHandler)
	log.Root().SetHandler(log.FuncHandler(func(r *log.Record) error {
		switch r.Msg {
		case "Staged SfcV2 upgrade from binary rules; will activate at next epoch seal",
			"Staged Podgorica upgrade from binary rules; will activate at next epoch seal",
			"Staged Elemont upgrade from binary rules; will activate at next epoch seal":
			atomic.AddInt64(&stagedLogCount, 1)
		}
		return prevHandler.Log(r)
	}))

	// First boot: staging fires for all three flags, producing 3 log lines.
	env := newPreElemontTestEnv(2, 3)
	defer env.Close()

	require.Equal(t, int64(3), atomic.LoadInt64(&stagedLogCount),
		"first boot must stage SfcV2, Podgorica, and Elemont exactly once each")

	bs1 := env.store.GetBlockState()
	require.NotNil(t, bs1.DirtyRules)
	require.True(t, bs1.DirtyRules.Upgrades.SfcV2)
	require.True(t, bs1.DirtyRules.Upgrades.Podgorica)
	require.True(t, bs1.DirtyRules.Upgrades.Elemont)

	// Second boot: re-create the service from the same store. The staging
	// loop must observe DirtyRules already at the target values and produce
	// no additional log lines. The second service is intentionally not
	// started or Close()d — only the constructor path (staging) is under
	// test here; background goroutines never run, so there is nothing to
	// tear down. The first env's defer env.Close() owns the store cleanup.
	atomic.StoreInt64(&stagedLogCount, 0)
	store := env.store
	blockProc := DefaultBlockProc()
	blockProc.EventsModule = testConfirmedEventsModule{blockProc.EventsModule, env}
	engine, vecClock := makeTestEngine(store)
	txPool := &dummyTxPool{}
	_, err := newService(DefaultConfig(cachescale.Identity), store, blockProc, engine, vecClock, func(_ evmcore.StateReader) TxPool {
		return txPool
	})
	require.NoError(t, err)

	require.Equal(t, int64(0), atomic.LoadInt64(&stagedLogCount),
		"second boot must not re-stage already-staged upgrade flags")

	bs2 := env.store.GetBlockState()
	require.NotNil(t, bs2.DirtyRules)
	require.True(t, bs2.DirtyRules.Upgrades.SfcV2)
	require.True(t, bs2.DirtyRules.Upgrades.Podgorica)
	require.True(t, bs2.DirtyRules.Upgrades.Elemont)
}
