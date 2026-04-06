package emitter

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	notify "github.com/ethereum/go-ethereum/event"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/gossip/emitter/mock"
	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/vecmt"
)

// newTestEmitter builds a fresh Emitter backed entirely by mocks and returns
// both the emitter and the mocks so callers can set additional EXPECT()s.
func newTestEmitter(t *testing.T) (
	em *Emitter,
	external *mock.MockExternal,
	txPool *mock.MockTxPool,
	signer *mock.MockSigner,
	txSigner *mock.MockTxSigner,
) {
	t.Helper()

	cfg := DefaultConfig()
	gValidators := makefakegenesis.GetFakeValidators(3)
	vv := pos.NewBuilder()
	for _, v := range gValidators {
		vv.Set(v.ID, pos.Weight(1))
	}
	validators := vv.Build()
	cfg.Validator.ID = gValidators[0].ID

	ctrl := gomock.NewController(t)
	external = mock.NewMockExternal(ctrl)
	txPool = mock.NewMockTxPool(ctrl)
	signer = mock.NewMockSigner(ctrl)
	txSigner = mock.NewMockTxSigner(ctrl)

	// Expectations shared by init() and the emitter background goroutine.
	external.EXPECT().Lock().AnyTimes()
	external.EXPECT().Unlock().AnyTimes()
	external.EXPECT().DagIndex().Return((*vecmt.Index)(nil)).AnyTimes()
	external.EXPECT().IsSynced().Return(true).AnyTimes()
	external.EXPECT().PeersNum().Return(3).AnyTimes()
	external.EXPECT().GetRules().Return(opera.FakeNetRules()).AnyTimes()
	external.EXPECT().GetEpochValidators().Return(validators, idx.Epoch(1)).AnyTimes()
	external.EXPECT().GetLastEvent(idx.Epoch(1), cfg.Validator.ID).Return((*hash.Event)(nil)).AnyTimes()
	external.EXPECT().GetGenesisTime().Return(inter.Timestamp(uint64(time.Now().UnixNano()))).AnyTimes()

	em = NewEmitter(cfg, World{
		External: external,
		TxPool:   txPool,
		Signer:   signer,
		TxSigner: txSigner,
	})
	return em, external, txPool, signer, txSigner
}

// nopSubscription returns an event.Subscription whose goroutine exits cleanly
// when Unsubscribe is called.
func nopSubscription() notify.Subscription {
	return notify.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

// TestEmitter_StartStop verifies that Start() launches a background goroutine
// and Stop() shuts it down cleanly. A second Stop() must not panic.
func TestEmitter_StartStop(t *testing.T) {
	em, _, txPool, _, _ := newTestEmitter(t)
	txPool.EXPECT().SubscribeNewTxsNotify(gomock.Any()).Return(nopSubscription()).Times(1)

	em.Start()
	require.NotNil(t, em.done, "done channel should be non-nil after Start")

	em.Stop()
	require.Nil(t, em.done, "done channel should be nil after Stop")

	// A second Stop() must be a no-op — not a panic.
	require.NotPanics(t, func() { em.Stop() })
}

// TestEmitter_Tick_BusyPreventsEmission verifies that IsBusy()=true prevents
// event emission: prevEmittedAtTime must stay unchanged.
func TestEmitter_Tick_BusyPreventsEmission(t *testing.T) {
	em, external, _, _, _ := newTestEmitter(t)
	external.EXPECT().IsBusy().Return(true).AnyTimes()

	em.init()
	before := em.prevEmittedAtTime
	em.tick()

	require.Equal(t, before, em.prevEmittedAtTime,
		"tick() must not update prevEmittedAtTime when IsBusy returns true")
}

// TestEmitter_Tick_NotBusyCallsBuild verifies that when the emitter is not
// busy and the emit interval has elapsed, tick() reaches External.Build —
// proving it actually attempts event creation.
func TestEmitter_Tick_NotBusyCallsBuild(t *testing.T) {
	em, external, txPool, _, _ := newTestEmitter(t)

	external.EXPECT().IsBusy().Return(false).AnyTimes()
	txPool.EXPECT().Count().Return(0).AnyTimes()
	txPool.EXPECT().Pending(gomock.Any()).Return(nil, nil).AnyTimes()
	external.EXPECT().GetHeads(idx.Epoch(1)).Return(hash.Events{}).AnyTimes()
	external.EXPECT().GetLatestBlockIndex().Return(idx.Block(0)).AnyTimes()

	// FakeNetRules has Llr=true, so the LLR vote helpers run. Nil returns
	// short-circuit them without producing vote payloads.
	external.EXPECT().GetLowestEpochToDecide().Return(idx.Epoch(0)).AnyTimes()
	external.EXPECT().GetLastEV(gomock.Any()).Return((*idx.Epoch)(nil)).AnyTimes()
	external.EXPECT().GetEpochRecordHash(gomock.Any()).Return((*hash.Hash)(nil)).AnyTimes()
	external.EXPECT().GetLowestBlockToDecide().Return(idx.Block(0)).AnyTimes()
	external.EXPECT().GetLastBV(gomock.Any()).Return((*idx.Block)(nil)).AnyTimes()
	external.EXPECT().GetBlockRecordHash(gomock.Any()).Return((*hash.Hash)(nil)).AnyTimes()

	// ErrNotEnoughGasPower short-circuits createEvent cleanly. We must not
	// invoke onIndexed: it calls quorumIndexer.GetMetricOf which dereferences
	// the nil vecmt.Index returned by DagIndex().
	var buildCalled atomic.Bool
	external.EXPECT().Build(gomock.Any(), gomock.Any()).
		DoAndReturn(func(e *inter.MutableEventPayload, onIndexed func()) error {
			buildCalled.Store(true)
			return ErrNotEnoughGasPower
		}).AnyTimes()

	em.init()
	em.intervals.DoublesignProtection = 0 // disable; see isSyncedToEmit
	em.prevEmittedAtTime = time.Time{}

	em.tick()

	require.True(t, buildCalled.Load(),
		"tick() should call External.Build when the emitter is not busy and the emit interval has elapsed")
}
