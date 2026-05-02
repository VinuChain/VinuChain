package fast

type Reader struct {
	buf    []byte
	offset int
}

type Writer struct {
	buf []byte
}

// NewReader wraps bytes with reading buffer.
func NewReader(bb []byte) *Reader {
	return &Reader{
		buf:    bb,
		offset: 0,
	}
}

// NewWriter wraps bytes with writing buffer.
func NewWriter(bb []byte) *Writer {
	return &Writer{
		buf: bb,
	}
}

// WriteByteFast appends a single byte to the buffer.
//
// Renamed from WriteByte to avoid go vet flagging it as a non-conforming
// implementation of io.ByteWriter (which requires `WriteByte(byte) error`).
// This buffer's append-to-slice cannot fail and the caller does not need
// the error contract; preserving the panic-free semantic and matching the
// fast.Reader.ReadByteFast naming keeps the package's intent visible.
func (b *Writer) WriteByteFast(v byte) {
	b.buf = append(b.buf, v)
}

// Write the byte to the buffer.
func (b *Writer) Write(v []byte) {
	b.buf = append(b.buf, v...)
}

// Read n bytes.
func (b *Reader) Read(n int) []byte {
	var res []byte
	res = b.buf[b.offset : b.offset+n]
	b.offset += n

	return res
}

// ReadByteFast reads one byte and advances the cursor. Out-of-bounds
// reads panic via the underlying slice index, matching the fast-path
// semantic of every other Reader method on this type (callers treat any
// short-buffer condition as a malformed-encoding panic, not a recoverable
// error).
//
// Renamed from ReadByte to avoid go vet flagging it as a non-conforming
// implementation of io.ByteReader (which requires `ReadByte() (byte, error)`).
// Returning the (byte, error) tuple here would force every hot-path caller
// in utils/cser to add an error-discarding shim that obscures the
// panic-on-malformed contract this package is built around.
func (b *Reader) ReadByteFast() byte {
	res := b.buf[b.offset]
	b.offset++
	return res
}

// Position of internal cursor.
func (b *Reader) Position() int {
	return b.offset
}

// Bytes of internal buffer
func (b *Reader) Bytes() []byte {
	return b.buf
}

// Bytes of internal buffer
func (b *Writer) Bytes() []byte {
	return b.buf
}

// Empty returns true if the whole buffer is consumed
func (b *Reader) Empty() bool {
	return len(b.buf) == b.offset
}
