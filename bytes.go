package gtl

import (
	"bytes"
	"io"
	"unsafe"
)

// Bytes defines a `Vec` of the type byte.
type Bytes Vec[byte]

var (
	_ io.WriterTo   = Bytes{}
	_ io.ReaderFrom = Bytes{}
)

// NewBytes creates a new `Bytes` with a len and cap
func NewBytes(size, capacity int) Bytes {
	return BytesFrom(make([]byte, size, capacity))
}

// BytesFrom creates a `Bytes` from a byte slice.
func BytesFrom(bts []byte) Bytes {
	return Bytes(bts)
}

// Slice returns a Bytes object slice in parts.
// Calling Slice is equivalent to `b[low:max]`.
// But `max` can be <= 0. If so, calling Slice is
// equivalent to `b[low:]`.
func (b Bytes) Slice(low, max int) Bytes {
	if max <= 0 {
		return b[low:]
	}

	return b[low:max]
}

// CopyFrom copies the bytes from
func (b *Bytes) CopyFrom(b2 Bytes) {
	b.Reserve(b2.Len())
	copy(*b, b2)
}

// WriteTo writes the bytes from `b` to `w`.
func (b Bytes) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b)

	return int64(n), err
}

// ReadFrom reads the contents from `r` to `b`.
func (b Bytes) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(b)

	return int64(n), err
}

// LimitReadFrom reads a limited amount from `r` to `b`.
//
// If `n` is greater than `b`'s len, b will be resized.
func (b *Bytes) LimitReadFrom(r io.Reader, n int) (int64, error) {
	b.Reserve(n)
	return b.Slice(0, n).ReadFrom(r)
}

// LimitWriteTo writes a limited amount from `b` to `w`.
func (b Bytes) LimitWriteTo(w io.Writer, n int) (int64, error) {
	return b.Slice(0, Min(b.Len(), n)).WriteTo(w)
}

// Len returns the length of Bytes.
func (b Bytes) Len() int {
	return len(b)
}

// Cap returns the capacity of Bytes.
func (b Bytes) Cap() int {
	return cap(b)
}

// String returns the `Bytes`' string representation.
func (b Bytes) String() string {
	return string(b)
}

// UnsafeString returns the `Bytes`' string representation avoiding allocations.
// This method makes use of the `unsafe` package, and any modification to the string will affect `b`.
func (b Bytes) UnsafeString() string {
	return *(*string)(unsafe.Pointer(&b))
}

// Index returns the index of `c`.
func (b Bytes) Index(c byte) int {
	return bytes.IndexByte(b, c)
}

// Contains returns whether `b` contains `c` or not.
func (b Bytes) Contains(c byte) bool {
	return bytes.IndexByte(b, c) != -1
}

// Resize increases (if needed) the length of Bytes to match `n`.
func (b *Bytes) Resize(n int) {
	vc := (*Vec[byte])(b)
	vc.Resize(n)
}

// Reserve ensures that Bytes has at least len = n.
func (b *Bytes) Reserve(n int) {
	vc := (*Vec[byte])(b)
	vc.Reserve(n)
}

// Append appends the bytes to the end of Bytes.
func (b *Bytes) Append(bts ...byte) {
	vc := (*Vec[byte])(b)
	vc.Append(bts...)
}

// Push pushes the `bts` into the first positions of Bytes.
func (b *Bytes) Push(bts ...byte) {
	vc := (*Vec[byte])(b)
	vc.Push(bts...)
}
