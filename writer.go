package brotli

import (
	"errors"
	"io"

	"github.com/andybalholm/brotli/matchfinder"
)

const (
	BestSpeed          = 0
	BestCompression    = 11
	DefaultCompression = 6
)

// WriterOptions configures Writer.
type WriterOptions struct {
	// Quality controls the compression-speed vs compression-density trade-offs.
	// The higher the quality, the slower the compression. Range is 0 to 11.
	Quality int
	// LGWin is the base 2 logarithm of the sliding window size.
	// Range is 10 to 24. 0 indicates automatic configuration based on Quality.
	LGWin int
}

var (
	errEncode       = errors.New("brotli: encode error")
	errWriterClosed = errors.New("brotli: Writer is closed")
)

// Writes to the returned writer are compressed and written to dst.
// It is the caller's responsibility to call Close on the Writer when done.
// Writes may be buffered and not flushed until Close.
func NewWriter(dst io.Writer) *Writer { _ = "STUB: not implemented"; return nil }

// NewWriterLevel is like NewWriter but specifies the compression level instead
// of assuming DefaultCompression.
// The compression level can be DefaultCompression or any integer value between
// BestSpeed and BestCompression inclusive.
func NewWriterLevel(dst io.Writer, level int) *Writer { _ = "STUB: not implemented"; return nil }

// NewWriterOptions is like NewWriter but specifies WriterOptions
func NewWriterOptions(dst io.Writer, options WriterOptions) *Writer {
	_ = "STUB: not implemented"
	return nil
}

// Reset discards the Writer's state and makes it equivalent to the result of
// its original state from NewWriter or NewWriterLevel, but writing to dst
// instead. This permits reusing a Writer rather than allocating a new one.
func (w *Writer) Reset(dst io.Writer) { _ = "STUB: not implemented"; return }

func (w *Writer) writeChunk(p []byte, op int) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Flush outputs encoded data for all input provided to Write. The resulting
// output can be decoded to match all input before Flush, but the stream is
// not yet complete until after Close.
// Flush has a negative impact on compression.
func (w *Writer) Flush() error { _ = "STUB: not implemented"; return nil }

// Close flushes remaining data to the decorated writer.
func (w *Writer) Close() error {
	_ = "STUB: not implemented"
	// If stream is already closed, it is reported by `writeChunk`.
	return nil
}

// Write implements io.Writer. Flush or Close must be called to ensure that the
// encoded bytes are actually flushed to the underlying Writer.
func (w *Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error {
	_ = "STUB: not implemented"

	// NewWriterV2 is like NewWriterLevel, but it uses the new implementation
	// based on the matchfinder package. It currently supports up to level 9;
	// if a higher level is specified, level 9 will be used.
	return nil
}

func NewWriterV2(dst io.Writer, level int) *matchfinder.Writer {
	_ = "STUB: not implemented"
	return nil
}
