package brotli

import "github.com/andybalholm/brotli/matchfinder"

// An Encoder implements the matchfinder.Encoder interface, writing in Brotli format.
type Encoder struct {
	wroteHeader bool
	bw          bitWriter
	distCache   []distanceCode
}

func (e *Encoder) Reset() { _ = "STUB: not implemented"; return }

func (e *Encoder) Encode(dst []byte, src []byte, matches []matchfinder.Match, lastBlock bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// islast + isempty

// first pass: build the histograms

// d is the ring buffer of the last 4 distances.

// If the stream ends with unmatched bytes, we need a dummy copy length.

// In my testing, codes 10–15 actually reduced the compression ratio.

// If the stream ends with unmatched bytes, we need a dummy copy length.

// islast + isempty

type distanceCode struct {
	code      int
	nExtra    uint
	extraBits uint64
}

func getDistanceCode(distance int) distanceCode {
	_ = "STUB: not implemented"
	return *new(distanceCode)
}
