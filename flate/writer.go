package flate

import (
	"io"

	"github.com/andybalholm/brotli/matchfinder"
)

// NewWriter returns a new matchfinder.Writer that compresses data at the given level,
// in flate encoding. Levels 1–9 are available; levels outside this range will
// be replaced with the closest level available.
func NewWriter(w io.Writer, level int) *matchfinder.Writer { _ = "STUB: not implemented"; return nil }

// NewGZIPWriter returns a new matchfinder.Writer that compresses data at the given
// level, in gzip encoding. Levels 1–9 are available; levels outside this range
// will be replaced by the closest level available.
func NewGZIPWriter(w io.Writer, level int) *matchfinder.Writer {
	_ = "STUB: not implemented"
	return nil
}

func newWriter(w io.Writer, level int, e matchfinder.Encoder) *matchfinder.Writer {
	_ = "STUB: not implemented"
	return nil
}
