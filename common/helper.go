package common

import (
	"errors"
	"io"
)

var errWhence = errors.New("Seek: invalid whence")
var errOffset = errors.New("Seek: invalid offset")

// OffsetReader implements Read, and ReadAt on a section
// of an underlying io.ReaderAt.
// The main difference between io.SectionReader and OffsetReader is that
// NewOffsetReadSeeker does not require the user to know the number of readable bytes.
//
// It also partially implements Seek, where the implementation panics if io.SeekEnd is passed.
// This is because, OffsetReader does not know the end of the file therefore cannot seek relative
// to it.
type OffsetReader struct {
	r    io.ReaderAt
	base int64
	off  int64
}

// NewOffsetReadSeeker returns an OffsetReader that reads from r
// starting offset offset off and stops with io.EOF when r reaches its end.
// The Seek function will panic if whence io.SeekEnd is passed.
func NewOffsetReadSeeker(r io.ReaderAt, off int64) *OffsetReader {
	return &OffsetReader{r, off, off}
}

func (o *OffsetReader) Read(p []byte) (n int, err error) {
	n, err = o.r.ReadAt(p, o.off)
	o.off += int64(n)
	if n == 0 {
		return n, io.EOF
	}
	return
}

func (o *OffsetReader) ReadAt(p []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, io.EOF
	}
	off += o.base
	return o.r.ReadAt(p, off)
}

func (o *OffsetReader) ReadByte() (byte, error) {
	b := []byte{0}
	_, err := o.Read(b)
	return b[0], err
}

func (o *OffsetReader) Offset() int64 {
	return o.off
}

func (o *OffsetReader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		o.off = offset + o.base
	case io.SeekCurrent:
		o.off += offset
	case io.SeekEnd:
		panic("unsupported whence: SeekEnd")
	}
	return o.Position(), nil
}

// Position returns the current position of this reader relative to the initial offset.
func (o *OffsetReader) Position() int64 {
	return o.off - o.base
}

// An OffsetWriter maps writes at offset base to offset base+off in the underlying writer.
type OffsetWriter struct {
	w    io.WriterAt
	base int64 // the original offset
	off  int64 // the current offset
}

// NewOffsetWriter returns an OffsetWriter that writes to w
// starting at offset off.
func NewOffsetWriter(w io.WriterAt, off int64) *OffsetWriter {
	return &OffsetWriter{w, off, off}
}

func (o *OffsetWriter) Write(p []byte) (n int, err error) {
	n, err = o.w.WriteAt(p, o.off)
	o.off += int64(n)
	return
}

func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error) {
	off += o.base
	return o.w.WriteAt(p, off)
}

func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errWhence
	case io.SeekStart:
		offset += o.base
	case io.SeekCurrent:
		offset += o.off
	}
	if offset < o.base {
		return 0, errOffset
	}
	o.off = offset
	return offset - o.base, nil
}
