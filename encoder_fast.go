package brotli

import (
	"github.com/andybalholm/brotli/matchfinder"
)

func gaussianProbability(x, mean, stdDev float64) float64 { _ = "STUB: not implemented"; return 0 }

// A FastEncoder implements the matchfinder.Encoder interface, writing in Brotli
// format. It uses a simplified encoding (like level 0 in the reference
// implementation) to save time.
type FastEncoder struct {
	wroteHeader   bool
	bw            bitWriter
	commandHisto  [704]uint32
	distanceHisto [64]uint32
}

func (e *FastEncoder) Reset() { _ = "STUB: not implemented"; return }

func (e *FastEncoder) Encode(dst []byte, src []byte, matches []matchfinder.Match, lastBlock bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Fill the histograms with default statistics.

// For the command codes we're using for insert lengths (insert + 2-byte copy),
// fill the histogram with a Zipf-squared distribution.

// For the command codes we're using for copy lengths (0 insert + copy
// (length - 2), with repeat distance),
// fill the histogram with Zipf distribution starting at code 1 (match length 5),
// but a smaller frequency for code 0.

// Fill in the combined codes for short insert and copy lengths.

// Fill e.distanceHisto with a normal distribution.

// islast + isempty

// Reset the statistics, starting with a count of 1 for each symbol we might use.

// Write a command with the appropriate insert length, and a copy length of 2.

// We can use a combined insert/copy code with no extra bits.

// Write the literals, if any.

// Write the distance code.

// Write a command for the remainder of the match (after the first two bytes
// from before), using the previous distance.

// We don't need to finish the length.

// islast + isempty
