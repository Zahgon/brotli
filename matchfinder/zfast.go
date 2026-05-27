package matchfinder

type tableEntry struct {
	val    uint32
	offset int32
}

const (
	zfastTableBits = 15
	zfastTableSize = 1 << zfastTableBits
	zfastHashLen   = 6

	prime6Bytes = 227718039650203
)

// ZFast is a MatchFinder based on the "Fastest" setting in
// github.com/klauspost/compress/zstd.
type ZFast struct {
	MaxDistance int

	history []byte
	// current is the offset at the start of history
	current int32
	table   [zfastTableSize]tableEntry
}

func (z *ZFast) Reset() { _ = "STUB: not implemented"; return }

func (z *ZFast) FindMatches(dst []Match, src []byte) []Match { _ = "STUB: not implemented"; return nil }

// Protect against overflow of current.

// history doesn't have enough capacity to hold the new block.

// Move down

// t will contain the match offset when we find one.
// When exiting the search loop, we have already checked 4 bytes.

// There is a repeated match at s+2.

// found a regular match

// A 4-byte match has been found. We'll later see if more than
// 4 bytes.

// Check offset 2

// Store the hash, since we have it.

func (z *ZFast) hash(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }
