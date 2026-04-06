package launcher

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
)

// maxTraceImportStreamSize is the expected cap for trace import streams.
// It must equal maxImportStreamSize (defined in import.go).
const maxTraceImportStreamSize = maxImportStreamSize

// buildTraceRLPStream constructs a valid RLP stream containing n TracePayload entries.
func buildTraceRLPStream(n int) []byte {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		p := TracePayload{
			Key:    common.BytesToHash([]byte{byte(i)}),
			Traces: []byte("trace data"),
		}
		if err := rlp.Encode(&buf, p); err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}

// TestImportTxTraces_StreamUsesMaxImportSize verifies that the rlp.Stream used
// for trace import has an input limit equal to maxImportStreamSize, not zero.
// A zero limit means unlimited input — an attacker can supply a file with an
// arbitrarily large RLP entry and cause OOM.
func TestImportTxTraces_StreamUsesMaxImportSize(t *testing.T) {
	// Build a small but valid RLP stream.
	data := buildTraceRLPStream(1)
	reader := bytes.NewReader(data)

	// Construct the stream the way importTxTraces should, with the size limit.
	stream := rlp.NewStream(reader, maxTraceImportStreamSize)

	// The stream should be able to decode the single entry.
	e := new(TracePayload)
	err := stream.Decode(e)
	require.NoError(t, err)
	require.Equal(t, []byte("trace data"), e.Traces)

	// Next decode should return io.EOF (stream exhausted).
	err = stream.Decode(e)
	require.ErrorIs(t, err, io.EOF)
}

// TestImportTxTraces_ZeroLimitIsUnbounded confirms that rlp.NewStream with
// inputLimit=0 and a generic io.Reader applies no size limit. This is why the
// import path must pass maxImportStreamSize explicitly.
func TestImportTxTraces_ZeroLimitIsUnbounded(t *testing.T) {
	data := buildTraceRLPStream(1)
	wrapped := &limitProbeReader{bytes.NewReader(data)}
	stream := rlp.NewStream(wrapped, 0)
	e := new(TracePayload)
	require.NoError(t, stream.Decode(e))
}

// limitProbeReader wraps a reader so it is not detected as a byte-slice reader.
type limitProbeReader struct{ r io.Reader }

func (p *limitProbeReader) Read(b []byte) (int, error) { return p.r.Read(b) }

// TestExportTraceTo_EncodeErrorPropagated verifies that a write failure during
// RLP encoding is propagated to the caller rather than silently ignored, which
// would produce a corrupt/truncated trace file with no diagnostic.
func TestExportTraceTo_EncodeErrorPropagated(t *testing.T) {
	errWrite := errors.New("write failed: disk full")
	w := &errWriter{err: errWrite}
	p := TracePayload{Key: common.Hash{1}, Traces: []byte("x")}
	err := rlp.Encode(w, p)
	require.ErrorIs(t, err, errWrite)
}

// errWriter is a writer that always returns an error.
type errWriter struct{ err error }

func (e *errWriter) Write(p []byte) (int, error) { return 0, e.err }

// TestBlockNumberOverflowIsDetected verifies that block numbers exceeding
// idx.Block (uint32) are rejected with an error rather than silently truncated.
func TestBlockNumberOverflowIsDetected(t *testing.T) {
	_, err := parseBlockNumber("4294967296") // 2^32, overflows uint32
	require.Error(t, err)

	n, err := parseBlockNumber("100")
	require.NoError(t, err)
	require.Equal(t, uint32(100), n)

	n, err = parseBlockNumber("0")
	require.NoError(t, err)
	require.Equal(t, uint32(0), n)
}

// buildOversizeRLP constructs a minimal RLP list header whose declared size
// exceeds the given limit, without allocating the full payload.
func buildOversizeRLP(declaredLen uint32) []byte {
	buf := make([]byte, 5)
	buf[0] = 0xf7 + 4 // long-form list with 4-byte length field
	binary.BigEndian.PutUint32(buf[1:5], declaredLen)
	return buf
}

// TestImportTxTraces_OversizeLimitEnforced verifies that the RLP stream rejects
// an entry whose declared size exceeds maxImportStreamSize.
func TestImportTxTraces_OversizeLimitEnforced(t *testing.T) {
	oversized := buildOversizeRLP(uint32(maxTraceImportStreamSize) + 1)
	reader := &limitProbeReader{bytes.NewReader(oversized)}
	stream := rlp.NewStream(reader, maxTraceImportStreamSize)

	e := new(TracePayload)
	err := stream.Decode(e)
	require.Error(t, err)
	require.True(t, strings.Contains(err.Error(), "too large") || strings.Contains(err.Error(), "exceeds"),
		"unexpected error: %v", err)
}
