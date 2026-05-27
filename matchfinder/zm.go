package matchfinder

const (
	zmTableBits     = 15
	zmTableSize     = 1 << zmTableBits
	zmLongTableBits = 17
	zmLongTableSize = 1 << zmLongTableBits
)

// ZM is a MatchFinder that combines the cache tables of ZDFast with the
// overlap-based parsing of M4.
type ZM struct {
	MaxDistance int
	history     []byte
	table       [zmTableSize]tableEntry
	longTable   [zmLongTableSize]tableEntry
}

func (z *ZM) Reset() { _ = "STUB: not implemented"; return }

func (z *ZM) FindMatches(dst []Match, src []byte) []Match { _ = "STUB: not implemented"; return nil }

// matches stores the matches that have been found but not emitted,
// in reverse order. (matches[0] is the most recent one.)

// Search for a match, starting after the last match emitted.

// t will contain the match offset when we find one.

// Look for a repeat match one byte after the current position.

// There is a repeated match at s+2.

// There is a long match at s.

// There is a regular match at s.
// See if we can find a long match at s+1.

// We found a long match at s+1, so we'll use that instead
// of the regular match at s.

// Store some table entries after s.

// We have a match in matches[0].
// Now look for overlapping matches.

// There is a long match at s.

// There is a regular match at s.

// See if we can find a long match at s+1.

// We found a long match at s+1, so we'll use that instead
// of the regular match at s.

// No overlapping match was found.

// The new match isn't longer than the old one, so we break out of the loop
// of looking for overlapping matches.

// We have three matches, so it's time to emit one and/or eliminate one.

// The first and third matches overlap; discard the one in between.

// The first and third matches don't overlap, but there's no room for
// another match between them. Emit the first match and discard the second.

// Emit the first match, shortening it if necessary to avoid overlap with the second.

// Store some table entries at the end of the last match.

// We're done looking for overlapping matches; emit the ones we have.

func (z *ZM) hashShort(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (z *ZM) hashLong(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }
