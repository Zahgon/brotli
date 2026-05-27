package brotli

import (
	"errors"
	"io"
)

type decodeError int

func (err decodeError) Error() string { _ = "STUB: not implemented"; return "" }

var errExcessiveInput = errors.New("brotli: excessive input")
var errInvalidState = errors.New("brotli: invalid state")

// readBufSize is a "good" buffer size that avoids excessive round-trips
// between C and Go but doesn't waste too much memory on buffering.
// It is arbitrarily chosen to be equal to the constant used in io.Copy.
const readBufSize = 32 * 1024

// NewReader creates a new Reader reading the given reader.
func NewReader(src io.Reader) *Reader { _ = "STUB: not implemented"; return nil }

// Reset discards the Reader's state and makes it equivalent to the result of
// its original state from NewReader, but reading from src instead.
// This permits reusing a Reader rather than allocating a new one.
// Error is always nil
func (r *Reader) Reset(src io.Reader) error { _ = "STUB: not implemented"; return nil }

// There was an unrecoverable error, leaving the Reader's state
// undefined. Clear out everything but the buffers.

func (r *Reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// If readErr is `nil`, we just proxy underlying stream behavior.

// Calling r.src.Read may block. Don't block if we have data to return.

// Top off the buffer.

// Not enough data to complete decoding.
