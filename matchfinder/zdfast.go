package matchfinder

const (
	zdfastLongTableBits = 17
	zdfastLongTableSize = 1 << zdfastLongTableBits
)

// ZDFast is a MatchFinder based on the "Default" setting in
// github.com/klauspost/compress/zstd.
type ZDFast struct {
	MaxDistance int

	history []byte
	// current is the offset at the start of history
	current   int32
	table     [zfastTableSize]tableEntry
	longTable [zdfastLongTableSize]tableEntry
}

func (z *ZDFast) Reset() { _ = "STUB: not implemented"; return }

func (z *ZDFast) FindMatches(dst []Match, src []byte) []Match {
	_ = "STUB: not implemented"
	return nil
}

// Protect against overflow of current.

// history doesn't have enough capacity to hold the new block.

// Move down

// t will contain the match offset when we find one.
// When exiting the search loop, we have already checked 4 bytes.

// There is a repeated match at s+1.

// found a long match (likely at least 8 bytes)

// Found a regular match.
// See if we can find a long match at s+1

// We found a long match at s+1, so we'll use that instead
// of the regular match at s.

// A 4-byte match has been found. We'll later see if more than
// 4 bytes.

// Store some table entries near the start and end of the match.

// Check offset 2

// Store the hashes, since we have them.

func (z *ZDFast) hashShort(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (z *ZDFast) hashLong(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }
