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
// (service.go:412-414). This is deliberate: the now-removed legacy gofuzz
// harness (its makeFuzzedHandler) left those readers zero-valued, which
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
		// handleMsg lazily creates per-peer state keyed on peer.id -- a rate-limit
		// bucket (peerRateLimit.Allow) and event/stream quota entries
		// (peerEventQuota/peerStreamQuota.Acquire). Production reclaims these in
		// unregisterPeer, but that short-circuits unless the peer is registered in
		// h.peers (which this harness never does), so mirror its per-peer RemovePeer
		// calls directly. Each iteration uses a fresh random id (so the 500-msg/10s
		// rate limit never trips and the decoders stay reachable); without this
		// cleanup those maps would grow one entry per iteration over the CI run.
		defer func() {
			h.peerRateLimit.RemovePeer(peer.id)
			h.peerEventQuota.RemovePeer(peer.id)
			h.peerStreamQuota.RemovePeer(peer.id)
		}()
		// handleMsg returning an error on malformed input is EXPECTED and fine;
		// a panic/hang is the bug the fuzzer hunts.
		_ = h.handleMsg(peer)
	})
}

// FuzzHandleMsgDeep is the nightly counterpart to FuzzHandleMsg: it fuzzes ALL
// 16 protocol message codes (0..15), INCLUDING the four Request*Stream codes
// (RequestEventsStream=8, RequestBVsStream=10, RequestBRsStream=12,
// RequestEPsStream=14) that FuzzHandleMsg deliberately excludes.
//
// FuzzHandleMsg omits 8/10/12/14 because each dispatches to
// seeder.NotifyRequestReceived, which blocks on a 16-buffered notifyReceivedRequest
// channel (lachesis-base basestreamseeder/seeder.go) drained ONLY by the seeder's
// readerLoop -- and FuzzHandleMsg never starts the seeders, so >16 synthesized
// same-type stream requests would deadlock the harness. FuzzHandleMsgDeep closes
// that gap by additionally starting the four stream seeders (and the matching
// processors), so the reader loops drain those channels and the codes become
// fuzzable without hanging. This wider component set has a longer startup/teardown
// cost per run, so it is reserved for the nightly (longer-runtime) workflow while
// FuzzHandleMsg stays lean for the per-PR job.
//
// As in FuzzHandleMsg, handleMsg returning an error on malformed input is EXPECTED
// and not a finding; a panic or hang is what the fuzzer hunts.
func FuzzHandleMsgDeep(f *testing.F) {
	// Same global-log detachment rationale as FuzzHandleMsg: other gossip tests
	// bind log.Root() to their own *testing.T at debug level and never restore it,
	// so newTestEnv's debug logging would route into a completed test's t.Log and
	// panic. Install a discard handler for the duration and restore on exit.
	prevHandler := log.Root().GetHandler()
	log.Root().SetHandler(log.DiscardHandler())
	defer log.Root().SetHandler(prevHandler)

	// firstEpoch=1, validatorsNum=3 -- identical wired construction to FuzzHandleMsg.
	env := newTestEnv(1, 3)
	defer env.Close()
	h := env.handler

	// Open the AcceptTxs()/AcceptEvents() gates so the fuzzed payloads reach the
	// real tx/event decoders (see FuzzHandleMsg for the full sync-state rationale).
	h.syncStatus.MarkMaybeSynced()
	h.syncStatus.Set(ssEvents)

	// Start the worker components in the SAME order handler.Start() uses
	// (handler.go ~:488-506), MINUS the four leechers. FuzzHandleMsg starts only
	// the first four (dagFetcher/txFetcher/Heavycheck/dagProcessor); Deep adds the
	// remaining processor+seeder pairs so the Request*Stream codes can drain:
	//   RequestEventsStream(8)  -> h.dagSeeder.NotifyRequestReceived (handler_sync.go:595)
	//   RequestBVsStream(10)    -> h.bvSeeder.NotifyRequestReceived  (handler_sync.go:659)
	//   RequestBRsStream(12)    -> h.brSeeder.NotifyRequestReceived  (handler_sync.go:721)
	//   RequestEPsStream(14)    -> h.epSeeder.NotifyRequestReceived  (handler_sync.go:780)
	// Each seeder's started readerLoop drains its 16-buffered notifyReceivedRequest
	// channel (basestreamseeder/seeder.go), so synthesized stream requests no longer
	// deadlock. The matching processors (ep/dag/bv/br) are started too, mirroring
	// Start()'s processor-before-seeder pairing and keeping the response-decode paths
	// reachable. The seeders are self-contained -- their reader loops call only the
	// store-backed iterate/forEach callbacks wired in newHandler, so they do not need
	// a fully-started Service.
	//
	// We do NOT start the four leechers (dagLeecher/bvLeecher/brLeecher/epLeecher):
	// their loops initiate OUTBOUND stream requests to peers and would spin/hang
	// against the no-op fuzzRW. Leaving them unstarted ALSO keeps the four
	// *StreamResponse codes (9/11/13/15) breaking early at leecher.IsValidSession
	// (mutex-only check, returns false while session.agent is nil) before any channel
	// op -- exactly as in FuzzHandleMsg.
	//
	// Stops run in defer/LIFO (reverse of start order) so no goroutine leaks across
	// the run.
	h.dagFetcher.Start()
	defer h.dagFetcher.Stop()
	h.txFetcher.Start()
	defer h.txFetcher.Stop()
	h.checkers.Heavycheck.Start()
	defer h.checkers.Heavycheck.Stop()

	h.epProcessor.Start()
	defer h.epProcessor.Stop()
	h.epSeeder.Start()
	defer h.epSeeder.Stop()

	h.dagProcessor.Start()
	defer h.dagProcessor.Stop()
	h.dagSeeder.Start()
	defer h.dagSeeder.Stop()

	h.bvProcessor.Start()
	defer h.bvProcessor.Stop()
	h.bvSeeder.Start()
	defer h.bvSeeder.Stop()

	h.brProcessor.Start()
	defer h.brProcessor.Stop()
	h.brSeeder.Start()
	defer h.brSeeder.Stop()

	// Seed corpus: byte[0] selects a message code (mod 16), the rest is the payload.
	// Include seeds that land on the four Request*Stream codes now under test.
	f.Add([]byte{0x00})
	f.Add([]byte{0x04, 0x01, 0x02, 0x03})
	f.Add([]byte{0x02})
	f.Add([]byte{0x08})             // RequestEventsStream
	f.Add([]byte{0x0a, 0x01, 0x02}) // RequestBVsStream
	f.Add([]byte{0x0c, 0x03, 0x04}) // RequestBRsStream
	f.Add([]byte{0x0e, 0x05, 0x06}) // RequestEPsStream

	f.Fuzz(func(t *testing.T, data []byte) {
		msg, err := decodeFuzzMsgAll(data)
		if err != nil {
			return // not interesting
		}
		// Same synthetic-peer construction as FuzzHandleMsg: newPeer initializes the
		// mapset/queue/datasemaphore/term state that the reachable handleMsg paths
		// touch (a bare literal would panic/hang), starts the broadcaster against the
		// no-op fuzzRW, and a fresh random id per iteration keeps the rate limiter and
		// quota maps from collapsing or tripping across the corpus.
		peer := newPeer(
			ProtocolVersion,
			p2p.NewPeer(randFuzzID(), "fuzz-peer", []p2p.Cap{}),
			&fuzzRW{msg},
			DefaultPeerCacheConfig(cachescale.Identity),
		)
		defer peer.Close()
		// Mirror unregisterPeer's per-peer RemovePeer calls (the harness never
		// registers the peer in h.peers, so production cleanup short-circuits).
		defer func() {
			h.peerRateLimit.RemovePeer(peer.id)
			h.peerEventQuota.RemovePeer(peer.id)
			h.peerStreamQuota.RemovePeer(peer.id)
		}()
		// An error on malformed input is EXPECTED; a panic/hang is the bug.
		_ = h.handleMsg(peer)
	})
}

// decodeFuzzMsgAll is the FuzzHandleMsgDeep counterpart to decodeFuzzMsg: it maps
// data[0] to one of ALL 16 protocol message codes (HandshakeMsg=0 ..
// EPsStreamResponse=15, protocol.go) and wraps the remaining bytes as the payload.
// Unlike decodeFuzzMsg it INCLUDES the four Request*Stream codes (8/10/12/14),
// which are safe to fuzz here only because FuzzHandleMsgDeep starts the seeders
// that drain their request channels (see that target's doc comment).
func decodeFuzzMsgAll(data []byte) (*p2p.Msg, error) {
	if len(data) < 1 {
		return nil, errors.New("empty data")
	}

	var (
		// The full handleMsg-dispatched range 0..15, including the four
		// Request*Stream codes excluded by decodeFuzzMsg.
		codes = []uint64{
			HandshakeMsg,         // 0
			ProgressMsg,          // 1
			EvmTxsMsg,            // 2
			NewEvmTxHashesMsg,    // 3
			GetEvmTxsMsg,         // 4
			NewEventIDsMsg,       // 5
			GetEventsMsg,         // 6
			EventsMsg,            // 7
			RequestEventsStream,  // 8
			EventsStreamResponse, // 9
			RequestBVsStream,     // 10
			BVsStreamResponse,    // 11
			RequestBRsStream,     // 12
			BRsStreamResponse,    // 13
			RequestEPsStream,     // 14
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

// FuzzHandleMsgDeepLLR is the block-record-stage counterpart to FuzzHandleMsgDeep.
// The two targets are mutually exclusive by design: the sync stage gates the
// untrusted decoders into two non-overlapping phases (sync.go), and a single run
// can only sit in one phase.
//
//   - FuzzHandleMsgDeep fixes the stage at ssEvents, so AcceptEvents()==true /
//     AcceptBlockRecords()==(!Is(ssEvents))==false. That reaches the event/tx
//     decoders but the three block-record response branches BVsStreamResponse(11),
//     BRsStreamResponse(13), EPsStreamResponse(15) early-`break` at their
//     `if !AcceptBlockRecords()` guard (handler_sync.go:671, 733, 792) BEFORE
//     msg.Decode.
//   - FuzzHandleMsgDeepLLR leaves the stage at its newTestEnv default (ssUnknown),
//     so AcceptBlockRecords()==(!Is(ssEvents))==true and those three branches now
//     PROCEED to msg.Decode (handler_sync.go:675, 739, 797), exercising the
//     bvsChunk/brsChunk/epsChunk decoders on attacker bytes. They then break at the
//     leecher IsValidSession check (leechers unstarted) -- but the decode has run,
//     which is the whole point. RequestLLR()==(!Is(ssEvents) || MaybeSynced())==true
//     in this stage as well. AcceptEvents()/AcceptTxs() are false here (both require
//     Is(ssEvents)); the event/tx codes are deliberately left to FuzzHandleMsgDeep.
//
// As in the other targets, an error on malformed input is EXPECTED; a panic or hang
// is what the fuzzer hunts. Reserved for the nightly (longer-runtime) workflow.
func FuzzHandleMsgDeepLLR(f *testing.F) {
	// Same global-log detachment rationale as the other targets (see FuzzHandleMsg).
	prevHandler := log.Root().GetHandler()
	log.Root().SetHandler(log.DiscardHandler())
	defer log.Root().SetHandler(prevHandler)

	// firstEpoch=1, validatorsNum=3 -- identical wired construction to FuzzHandleMsgDeep.
	env := newTestEnv(1, 3)
	defer env.Close()
	h := env.handler

	// CRITICAL: do NOT call h.syncStatus.Set(ssEvents). The stage stays at the
	// newTestEnv default (ssUnknown), so AcceptBlockRecords() == !Is(ssEvents) == true
	// and the BVs/BRs/EPs StreamResponse decoders become reachable. MarkMaybeSynced()
	// is set so RequestLLR() == (!Is(ssEvents) || MaybeSynced()) == true (it is already
	// true via the first disjunct in this stage; this makes the LLR-request intent
	// explicit and robust if the stage ever changes). AcceptEvents()/AcceptTxs() are
	// (correctly) false in this stage -- the event/tx decoders are FuzzHandleMsgDeep's
	// responsibility.
	h.syncStatus.MarkMaybeSynced()

	// Start the SAME components as FuzzHandleMsgDeep, in handler.Start() order minus
	// the leechers, with deferred Stop in LIFO order. The seeders drain the four
	// Request*Stream codes' notifyReceivedRequest channels (so 10/12/14 cannot
	// deadlock); the leechers stay unstarted so the *StreamResponse branches break at
	// IsValidSession AFTER decode. None of these components panics on the seed corpus
	// in this stage.
	h.dagFetcher.Start()
	defer h.dagFetcher.Stop()
	h.txFetcher.Start()
	defer h.txFetcher.Stop()
	h.checkers.Heavycheck.Start()
	defer h.checkers.Heavycheck.Stop()

	h.epProcessor.Start()
	defer h.epProcessor.Stop()
	h.epSeeder.Start()
	defer h.epSeeder.Stop()

	h.dagProcessor.Start()
	defer h.dagProcessor.Stop()
	h.dagSeeder.Start()
	defer h.dagSeeder.Stop()

	h.bvProcessor.Start()
	defer h.bvProcessor.Stop()
	h.bvSeeder.Start()
	defer h.bvSeeder.Stop()

	h.brProcessor.Start()
	defer h.brProcessor.Stop()
	h.brSeeder.Start()
	defer h.brSeeder.Stop()

	// Seed corpus: byte[0] selects a code (mod len). Include seeds landing on the
	// three block-record response codes whose decoders this stage newly reaches.
	f.Add([]byte{0x00})             // HandshakeMsg
	f.Add([]byte{0x03, 0x01})       // BVsStreamResponse (index 3 in decodeFuzzMsgLLR)
	f.Add([]byte{0x05, 0x02})       // BRsStreamResponse (index 5)
	f.Add([]byte{0x07, 0x03})       // EPsStreamResponse (index 7)
	f.Add([]byte{0x02, 0x04, 0x05}) // RequestBVsStream (index 2)

	f.Fuzz(func(t *testing.T, data []byte) {
		msg, err := decodeFuzzMsgLLR(data)
		if err != nil {
			return // not interesting
		}
		// Same synthetic-peer construction + cleanup as the other Deep target.
		peer := newPeer(
			ProtocolVersion,
			p2p.NewPeer(randFuzzID(), "fuzz-peer", []p2p.Cap{}),
			&fuzzRW{msg},
			DefaultPeerCacheConfig(cachescale.Identity),
		)
		defer peer.Close()
		defer func() {
			h.peerRateLimit.RemovePeer(peer.id)
			h.peerEventQuota.RemovePeer(peer.id)
			h.peerStreamQuota.RemovePeer(peer.id)
		}()
		_ = h.handleMsg(peer)
	})
}

// decodeFuzzMsgLLR maps data[0] to one of the block-record / LLR-stage message
// codes whose decoders FuzzHandleMsgDeepLLR's ssUnknown stage actually reaches.
// The three *StreamResponse codes 11/13/15 only get past their
// `if !AcceptBlockRecords()` guard in this stage; the three Request*Stream codes
// 10/12/14 decode unconditionally and drain through the started seeders. The two
// stage-agnostic codes Handshake(0)/Progress(1) are included as cheap filler.
// EventsStreamResponse(9) is deliberately omitted -- it gates on AcceptEvents()
// (handler_sync.go:607), which is false here, so it belongs to FuzzHandleMsgDeep.
func decodeFuzzMsgLLR(data []byte) (*p2p.Msg, error) {
	if len(data) < 1 {
		return nil, errors.New("empty data")
	}

	var (
		codes = []uint64{
			HandshakeMsg,      // 0  (stage-agnostic filler)
			ProgressMsg,       // 1  (stage-agnostic filler)
			RequestBVsStream,  // 10 (decodes unconditionally; drains via bvSeeder)
			BVsStreamResponse, // 11 (reaches msg.Decode only when AcceptBlockRecords())
			RequestBRsStream,  // 12 (decodes unconditionally; drains via brSeeder)
			BRsStreamResponse, // 13 (reaches msg.Decode only when AcceptBlockRecords())
			RequestEPsStream,  // 14 (decodes unconditionally; drains via epSeeder)
			EPsStreamResponse, // 15 (reaches msg.Decode only when AcceptBlockRecords())
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

// decodeFuzzMsg maps data[0] to one of the protocol message codes and wraps the
// remaining bytes as the payload. (The distinct name dates from coexisting
// with the now-removed legacy gofuzz harness.)
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

// fuzzRW is a p2p.MsgReadWriter that replays a single canned message. (Named
// distinctly from the now-removed legacy gofuzz harness's fuzzMsgReadWriter.)
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
// (replaces the now-removed legacy gofuzz harness's randomID).
func randFuzzID() (id enode.ID) {
	for i := range id {
		id[i] = byte(rand.Intn(255))
	}
	return id
}
