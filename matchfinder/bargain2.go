package matchfinder

const (
	bargain2TableBits = 18
	bargain2TableSize = 1 << bargain2TableBits
)

// Bargain2 is a MatchFinder that attempts to find the encoding with the lowest
// "bit cost", using 2 hash lengths (5 and 8).
type Bargain2 struct {
	MaxDistance int

	// Skip is whether to look for matches at every other byte instead of every
	// byte (to increase speed but decrease compression).
	Skip bool

	history []byte
	table5  [bargain2TableSize]tableEntry
	table8  [bargain2TableSize]tableEntry

	// holding onto buffers to reduce allocations:

	arrivals []arrival
	matches  []Match
}

func (z *Bargain2) Reset() { _ = "STUB: not implemented"; return }

func (z *Bargain2) FindMatches(dst []Match, src []byte) []Match {
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

// The match was extended backwards. Add it with and without the extra.

// We've found the shortest path; now walk it backward and store the matches.

func (z *Bargain2) hash5(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func (z *Bargain2) hash8(u uint64) uint32 { _ = "STUB: not implemented"; return 0 }
