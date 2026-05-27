package matchfinder

// Trio is a MatchFinder that uses 3 different hash lengths, and
// overlap parsing.
type Trio struct {
	MaxDistance int
	history     []byte
	table5      [1 << 16]tableEntry
	table8      [1 << 17]tableEntry
	table12     [1 << 18]tableEntry
}

func (z *Trio) Reset() { _ = "STUB: not implemented"; return }

func (z *Trio) FindMatches(dst []Match, src []byte) []Match { _ = "STUB: not implemented"; return nil }

// matches stores the matches that have been found but not emitted,
// in reverse order. (matches[0] is the most recent one.)

// Search for a match, starting after the last match emitted.

// t will contain the match offset when we find one.

// Look for a repeat match one byte after the current position.

// There is a repeated match at s+2.

// There is a 12-byte match at s.

// Look for a "lazy" match with a longer hash at s+1.

// We have a match in matches[0].
// Now look for overlapping matches.

// Store some entries that haven't been indexed yet.

// There is a 12-byte match at s.

// There is a long match at s.

// No overlapping match was found.

// The new match isn't longer than the old one, so we break out of the loop
// of looking for overlapping matches.

// We have three matches, so it's time to emit one and/or eliminate one.

// The first and third matches overlap; discard the one in between.

// The first and third matches don't overlap, but there's no room for
// another match between them. Emit the first match and discard the second.

// Emit the first match, shortening it if necessary to avoid overlap with the second.

// Store some entries up to the end of the last match.

// We're done looking for overlapping matches; emit the ones we have.

func (z *Trio) hash5(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (z *Trio) hash8(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (z *Trio) hash12(u uint64, e uint32) uint32 { _ = "STUB: not implemented"; return 0 }
