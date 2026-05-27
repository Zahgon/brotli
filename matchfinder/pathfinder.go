package matchfinder

// Pathfinder is a MatchFinder that uses hash chains to find matches, and a
// shortest-path optimizer to choose which matches to use.
type Pathfinder struct {
	// MaxDistance is the maximum distance (in bytes) to look back for
	// a match. The default is 65535.
	MaxDistance int

	// MinLength is the length of the shortest match to return.
	// The default is 4.
	MinLength int

	// HashLen is the number of bytes to use to calculate the hashes.
	// The maximum is 8 and the default is 6.
	HashLen int

	// TableBits is the number of bits in the hash table indexes.
	// The default is 17 (128K entries).
	TableBits int

	// ChainLength is how many entries to search on the "match chain" of older
	// locations with the same hash as the current location.
	ChainLength int

	table []uint32
	chain []uint32

	history []byte

	// holding onto buffers to reduce allocations:

	arrivals     []arrival
	foundMatches []absoluteMatch
	matches      []Match
}

func (q *Pathfinder) Reset() { _ = "STUB: not implemented"; return }

// An arrival represents how we got to a certain byte position.
// The cost is the total cost to get there from the beginning of the block.
// If distance > 0, the arrival is with a match.
// If distance == 0, the arrival is with a run of literals.
type arrival struct {
	length   uint32
	distance uint32
	cost     float32
}

const (
	baseMatchCost float32 = 4
)

func (q *Pathfinder) FindMatches(dst []Match, src []byte) []Match {
	_ = "STUB: not implemented"
	return nil
}

// Each element in arrivals corresponds to the position just after
// the corresponding byte in src.

// Trim down the history buffer.

// Append src to the history buffer.

// Calculate hashes and build the chain.

// Look for matches, and collect them in foundMatches. Later we'll figure out
// which ones to actually use.

// Look for a repeat match at i+1.

// We were looking for an overlapping match, but we didn't find one longer
// than the previous match. So we'll go back to sequential search,
// starting right after the previous match.

// No match found. Continue with sequential search.

// We've found a match; now look for matches overlapping the end of it.

// Matches shorter than 6 are comparatively rare, and therefore
// have longer codes.

// If a match from an earlier position extends far enough past the current
// position, try using the tail of it, starting from here.

// Matches shorter than 6 are comparatively rare, and therefore
// have longer codes.

// We've found the shortest path; now walk it backward and store the matches.
