package matchfinder

// A TextEncoder is an Encoder that produces a human-readable representation of
// the LZ77 compression. Matches are replaced with <Length,Distance> symbols.
type TextEncoder struct{}

func (t TextEncoder) Reset() { _ = "STUB: not implemented"; return }

func (t TextEncoder) Encode(dst []byte, src []byte, matches []Match, lastBlock bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// A NoMatchFinder implements MatchFinder, but doesn't find any matches.
// It can be used to implement the equivalent of the standard library flate package's
// HuffmanOnly setting.
type NoMatchFinder struct{}

func (n NoMatchFinder) Reset() { _ = "STUB: not implemented"; return }

func (n NoMatchFinder) FindMatches(dst []Match, src []byte) []Match {
	_ = "STUB: not implemented"
	return nil
}

// AutoReset wraps a MatchFinder that can return references to data in previous
// blocks, and calls Reset before each block. It is useful for (e.g.) using a
// snappy Encoder with a MatchFinder designed for flate. (Snappy doesn't
// support references between blocks.)
type AutoReset struct {
	MatchFinder
}

func (a AutoReset) FindMatches(dst []Match, src []byte) []Match {
	_ = "STUB: not implemented"
	return nil
}
