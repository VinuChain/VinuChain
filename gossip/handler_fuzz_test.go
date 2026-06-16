package gossip

import (
	"bytes"
	"errors"
	"math/rand"
	"testing"

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
	// dagFetcher, dagProcessor, dagSeeder, checkers, store) are initialized by
	// newHandler. handleMsg does not depend on handler.Start (only invoked from
	// the full Service.Start, which newTestEnv does not call), so we do NOT
	// start the handler -- that would launch background loops that assume a
	// fully-started service and add nothing handleMsg needs.
	h := env.handler

	// Seed corpus: byte[0] selects a message code, the rest is the payload.
	f.Add([]byte{0x00})
	f.Add([]byte{0x04, 0x01, 0x02, 0x03})
	f.Add([]byte{0x02})

	f.Fuzz(func(t *testing.T, data []byte) {
		msg, err := decodeFuzzMsg(data)
		if err != nil {
			return // not interesting
		}
		peer := &peer{
			version: ProtocolVersion,
			Peer:    p2p.NewPeer(randFuzzID(), "fuzz-peer", []p2p.Cap{}),
			rw:      &fuzzRW{msg},
		}
		// handleMsg returning an error on malformed input is EXPECTED and fine;
		// a panic/hang is the bug the fuzzer hunts.
		_ = h.handleMsg(peer)
	})
}

// decodeFuzzMsg maps data[0] to one of the protocol message codes and wraps the
// remaining bytes as the payload. Distinct name from the gofuzz-tagged
// newFuzzMsg so the two files can coexist.
func decodeFuzzMsg(data []byte) (*p2p.Msg, error) {
	if len(data) < 1 {
		return nil, errors.New("empty data")
	}

	var (
		codes = []uint64{
			HandshakeMsg,
			EvmTxsMsg,
			ProgressMsg,
			NewEventIDsMsg,
			GetEventsMsg,
			EventsMsg,
			RequestEventsStream,
			EventsStreamResponse,
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
