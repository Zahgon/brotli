package matchfinder

// M0 is an implementation of the MatchFinder interface based
// on the algorithm used by snappy, but modified to be more like the algorithm
// used by compression level 0 of the brotli reference implementation.
//
// It has a maximum block size of 65536 bytes.
type M0 struct {
	// Lazy turns on "lazy matching," for higher compression but less speed.
	Lazy bool

	MaxDistance int
	MaxLength   int
}

func (M0) Reset() { _ = "STUB: not implemented"; return }

const (
	m0HashLen = 5

	m0TableBits = 16
	m0TableSize = 1 << m0TableBits
	m0Shift     = 32 - m0TableBits
	// m0TableMask is redundant, but helps the compiler eliminate bounds
	// checks.
	m0TableMask = m0TableSize - 1
)

func (m M0) hash(data uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// FindMatches looks for matches in src, appends them to dst, and returns dst.
// src must not be longer than 65536 bytes.
func (m M0) FindMatches(dst []Match, src []byte) []Match { _ = "STUB: not implemented"; return nil }

// sLimit is when to stop looking for offset/length copies. The inputMargin
// lets us use a fast path for emitLiteral in the main loop, while we are
// looking for copies.

// nextEmit is where in src the next emitLiteral should start from.

// The encoded form must start with a literal, as there are no previous
// bytes to copy, so we start looking for hash matches at s == 1.

// Copied from the C++ snappy implementation:
//
// Heuristic match skipping: If 32 bytes are scanned with no matches
// found, start looking only at every other byte. If 32 more bytes are
// scanned (or skipped), look at every third byte, etc.. When a match
// is found, immediately go back to looking at every byte. This is a
// small loss (~5% performance, ~0.1% density) for compressible data
// due to more bookkeeping, but for non-compressible data (such as
// JPEG) it's a huge win since the compressor quickly "realizes" the
// data is incompressible and doesn't bother looking for matches
// everywhere.
//
// The "skip" variable keeps track of how many bytes there are since
// the last match; dividing it by 32 (ie. right-shifting by five) gives
// the number of bytes to move ahead for each iteration.

// Invariant: we have a 4-byte match at s.

// If lazy matching is enabled, we update the hash table for
// every byte in the match.

// We update the hash table only at base+1

// We could immediately start working at s now, but to improve
// compression we first update the hash table.
