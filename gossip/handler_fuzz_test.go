package gossip

import (
	"bytes"
	"errors"
	"math/rand"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// FuzzHandleMsg fuzzes the P2P message-handling path (handler.handleMsg) with
// arbitrary attacker-controlled bytes. The first input byte selects one of the
// protocol message codes; the remainder is fed as the message payload.
//
// The handler is built from a fully-wired testEnv (newTestEnv -> newService),
// which initializes BOTH the heavy-check and gas-power-check readers
// (service.go:412-414). This is deliberate: the legacy gofuzz harness
// (handler_fuzz.go's makeFuzzedHandler) leaves those readers zero-valued, which
// panics in a background goroutine once the handler starts -- a setup bug, not
// an attacker-input finding. By going through newService here, the deep
// validation path is reachable and any panic/hang surfaced is rooted in fed
// input.
//
// handleMsg returning an error on malformed input is EXPECTED and not a finding;
// a panic or hang is what the fuzzer hunts.
func FuzzHandleMsg(f *testing.F) {
	// Detach the global go-ethereum log root before constructing the env. Other
	// gossip tests (e.g. TestSFC via logger.SetTestMode + SetLevel("debug"))
	// bind the shared log.Root() handler to their own *testing.T and raise the
	// level to debug, but never restore it. The debug-level genesis/snapshot
	// logging triggered by newTestEnv below would then route into that already
	// completed test's t.Log, which the testing framework turns into a
	// "Log in goroutine after <Test> has completed" panic. That is a cross-test
	// harness artifact, not a handleMsg finding, so we install a discard handler
	// for the duration and restore the previous one on exit.
	prevHandler := log.Root().GetHandler()
	log.Root().SetHandler(log.DiscardHandler())
	defer log.Root().SetHandler(prevHandler)

	// firstEpoch=1, validatorsNum=3 -- mirrors the wired construction exemplar
	// in heavycheck_test.go (newTestEnv(startEpoch, validatorsNum)).
	env := newTestEnv(1, 3)
	defer env.Close()

	// *handler from the embedded *Service (service.go:146, set at :453). All
	// fields handleMsg relies on (msgSemaphore, peerRateLimit, txFetcher,
	// dagFetcher, dagProcessor, checkers, store) are initialized by newHandler.
	h := env.handler

	// Mark the handler synced so the AcceptTxs()/AcceptEvents() gates in
	// handleMsg are OPEN. newTestEnv builds the handler but never starts it, so
	// syncStatus sits at the zero value (stage=ssUnknown, maybeSynced=0); under
	// that state AcceptEvents()=Is(ssEvents) and AcceptTxs()=MaybeSynced() &&
	// Is(ssEvents) are BOTH false (sync.go), so the EvmTxs/NewEvmTxHashes/
	// GetEvmTxs/NewEventIDs/GetEvents/Events/EventsStreamResponse branches
	// early-return at the gate BEFORE decoding the fuzzed payload -- the harness
	// would pass without ever exercising its headline decoders. MarkMaybeSynced()
	// + Set(ssEvents) opens both gates so the fuzzed bytes reach the real tx/event
	// decoders and downstream processing.
	h.syncStatus.MarkMaybeSynced()
	h.syncStatus.Set(ssEvents)

	// Start ONLY the worker components the now-reachable gated decoders feed,
	// once, before fuzzing (NOT per-iteration). Each pushes to a bounded channel
	// drained solely by its own started loop; left unstarted they would DEADLOCK
	// the harness once the buffer fills (the same failure class as the seeder
	// request channel). Downstream wiring per code:
	//   EvmTxsMsg(2)        -> h.txFetcher.NotifyReceived + txpool.AddRemotes
	//   NewEvmTxHashesMsg(3)-> h.txFetcher.NotifyAnnounces
	//   NewEventIDsMsg(5)   -> h.dagFetcher.NotifyAnnounces
	//   EventsMsg(7)        -> h.dagFetcher.NotifyAnnounces + h.dagProcessor.Enqueue
	// dagProcessor's parentless check runs checkers.Heavycheck, so Heavycheck is
	// started too. GetEvmTxs(4)/GetEvents(6) are self-contained (txpool + the
	// peer's own queue, drained by newPeer's broadcaster). We do NOT start the
	// dag/bv/br/ep seeders or leechers, so the Request*Stream codes stay excluded
	// from decodeFuzzMsg and the *StreamResponse codes break at IsValidSession
	// (session.agent nil while the leecher is unstarted) before any channel op.
	// Stops run in defer/LIFO so no goroutine leaks across the run.
	h.checkers.Heavycheck.Start()
	defer h.checkers.Heavycheck.Stop()
	h.dagFetcher.Start()
	defer h.dagFetcher.Stop()
	h.txFetcher.Start()
	defer h.txFetcher.Stop()
	h.dagProcessor.Start()
	defer h.dagProcessor.Stop()

	// Seed corpus: byte[0] selects a message code, the rest is the payload.
	f.Add([]byte{0x00})
	f.Add([]byte{0x04, 0x01, 0x02, 0x03})
	f.Add([]byte{0x02})

	f.Fuzz(func(t *testing.T, data []byte) {
		msg, err := decodeFuzzMsg(data)
		if err != nil {
			return // not interesting
		}
		// Construct the synthetic peer through newPeer rather than a bare struct
		// literal. The reachable handleMsg paths touch peer state that a literal
		// leaves nil and which would then PANIC or HANG on otherwise-valid input,
		// turning a harness setup gap into a spurious fail-closed CI failure:
		//   - knownTxs / knownEvents are mapset.Set interfaces; handleTxs /
		//     handleEvents / handleEventHashes call MarkTransaction / MarkEvent
		//     (peer.go) which invoke .Cardinality()/.Add() -> nil-interface panic.
		//   - queue / queuedDataSemaphore are read by EnqueueSendTransactions /
		//     EnqueueSendEventsRLP on the GetEvmTxsMsg / GetEventsMsg paths;
		//     a nil channel send hangs and a nil *datasemaphore panics.
		//   - term is closed by Close to reap the broadcaster.
		// newPeer initializes all of them and starts peer.broadcast, whose
		// p2p.Send target is the no-op fuzzRW.WriteMsg (no real I/O). A fresh
		// peer per iteration avoids cross-iteration state and concurrency issues;
		// defer p.Close() reaps that iteration's broadcaster goroutine so the
		// fuzzer does not leak goroutines across the corpus. The non-empty id
		// "fuzz-peer" keeps the per-peer rate-limit key from collapsing to "".
		peer := newPeer(
			ProtocolVersion,
			p2p.NewPeer(randFuzzID(), "fuzz-peer", []p2p.Cap{}),
			&fuzzRW{msg},
			DefaultPeerCacheConfig(cachescale.Identity),
		)
		defer peer.Close()
		// handleMsg returning an error on malformed input is EXPECTED and fine;
		// a panic/hang is the bug the fuzzer hunts.
		_ = h.handleMsg(peer)
	})
}

// decodeFuzzMsg maps data[0] to one of the protocol message codes and wraps the
// remaining bytes as the payload. Distinct name from the gofuzz-tagged
// newFuzzMsg so the two files can coexist.
//
// The code set is the handleMsg-dispatched range (protocol.go: HandshakeMsg=0 ..
// EPsStreamResponse=15) MINUS the four Request*Stream codes (8/10/12/14). Those
// dispatch to seeder.NotifyRequestReceived, which sends to a 16-buffered
// notifyReceivedRequest channel (lachesis-base basestreamseeder/seeder.go) that
// is drained ONLY by the seeder reader loop started in handler.Start(). This
// lightweight harness intentionally does not start the handler, so once the
// fuzzer synthesizes >16 valid same-type stream requests the buffer fills and
// handleMsg blocks forever -- a CI hang, not a finding. A future nightly harness
// that starts the seeders could cover them.
//
// The four *StreamResponse codes (9/11/13/15) ARE kept: each response branch
// gates on leecher.IsValidSession before touching any processor/quota channel,
// and IsValidSession is a mutex-only check returning false while session.agent
// is nil -- which it always is here because the leechers are never started -- so
// those branches break out before they could block. Codes 0-7 are likewise safe
// (GetEvmTxs/GetEvents enqueue to the peer's OWN queue, drained by newPeer's
// broadcaster goroutine).
func decodeFuzzMsg(data []byte) (*p2p.Msg, error) {
	if len(data) < 1 {
		return nil, errors.New("empty data")
	}

	var (
		// Request*Stream codes 8/10/12/14 deliberately excluded -- see the doc
		// comment above (they would deadlock the unstarted seeder).
		codes = []uint64{
			HandshakeMsg,         // 0
			ProgressMsg,          // 1
			EvmTxsMsg,            // 2
			NewEvmTxHashesMsg,    // 3
			GetEvmTxsMsg,         // 4
			NewEventIDsMsg,       // 5
			GetEventsMsg,         // 6
			EventsMsg,            // 7
			EventsStreamResponse, // 9
			BVsStreamResponse,    // 11
			BRsStreamResponse,    // 13
			EPsStreamResponse,    // 15
		}
		code = codes[int(data[0])%len(codes)]
	)
	data = data[1:]

	return &p2p.Msg{
		Code:    code,
		Size:    uint32(len(data)),
		Payload: bytes.NewReader(data),
	}, nil
}

// fuzzRW is a p2p.MsgReadWriter that replays a single canned message. Distinct
// name from the gofuzz-tagged fuzzMsgReadWriter.
type fuzzRW struct {
	msg *p2p.Msg
}

func (rw *fuzzRW) ReadMsg() (p2p.Msg, error) {
	return *rw.msg, nil
}

func (rw *fuzzRW) WriteMsg(p2p.Msg) error {
	return nil
}

// randFuzzID returns a random enode ID for the synthetic peer. Distinct name
// from the gofuzz-tagged randomID.
func randFuzzID() (id enode.ID) {
	for i := range id {
		id[i] = byte(rand.Intn(255))
	}
	return id
}
