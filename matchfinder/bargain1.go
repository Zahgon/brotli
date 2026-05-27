package matchfinder

const (
	bargain1TableBits = 18
	bargain1TableSize = 1 << bargain1TableBits
)

// Bargain1 is a MatchFinder that attempts to find the encoding with the lowest
// "bit cost", using 1 hash length (6).
type Bargain1 struct {
	MaxDistance int

	// Skip is whether to look for matches at every other byte instead of every
	// byte (to increase speed but decrease compression).
	Skip bool

	history []byte
	table6  [bargain1TableSize]tableEntry

	// holding onto buffers to reduce allocations:

	arrivals []arrival
	matches  []Match
}

func (z *Bargain1) Reset() { _ = "STUB: not implemented"; return }

func (z *Bargain1) FindMatches(dst []Match, src []byte) []Match {
	_ = "STUB: not implemented"
	return nil
}

// Each element in arrivals corresponds to the position just after
// the corresponding byte in src.

// There's no room to check hashes.

// Look for a repeat match, unless there is no previous distance, or a match at
// that distance has already been found.

// We have a repeat of the previous match distance.

// The match was extended backwards. Add it with and without the extra.

// We've found the shortest path; now walk it backward and store the matches.

func (z *Bargain1) hash6(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }
