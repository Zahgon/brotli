// The matchfinder package defines reusable components for data compression.
//
// Many compression libraries have two main parts:
//   - Something that looks for repeated sequences of bytes
//   - An encoder for the compressed data format (often an entropy coder)
//
// Although these are logically two separate steps, the implementations are
// usually closely tied together. You can't use flate's matcher with snappy's
// encoder, for example. This package defines interfaces and an intermediate
// representation to allow mixing and matching compression components.
package matchfinder

import "io"

// A Match is the basic unit of LZ77 compression.
type Match struct {
	Unmatched int // the number of unmatched bytes since the previous match
	Length    int // the number of bytes in the matched string; it may be 0 at the end of the input
	Distance  int // how far back in the stream to copy from
}

// A MatchFinder performs the LZ77 stage of compression, looking for matches.
type MatchFinder interface {
	// FindMatches looks for matches in src, appends them to dst, and returns dst.
	FindMatches(dst []Match, src []byte) []Match

	// Reset clears any internal state, preparing the MatchFinder to be used with
	// a new stream.
	Reset()
}

// An Encoder encodes the data in its final format.
type Encoder interface {
	// Encode appends the encoded format of src to dst, using the match
	// information from matches.
	Encode(dst []byte, src []byte, matches []Match, lastBlock bool) []byte

	// Reset clears any internal state, preparing the Encoder to be used with
	// a new stream.
	Reset()
}

// A Writer uses MatchFinder and Encoder to write compressed data to Dest.
type Writer struct {
	Dest        io.Writer
	MatchFinder MatchFinder
	Encoder     Encoder

	// BlockSize is the number of bytes to compress at a time. If it is zero,
	// each Write operation will be treated as one block.
	BlockSize int

	err     error
	inBuf   []byte
	outBuf  []byte
	matches []Match
}

func (w *Writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) writeBlock(p []byte, lastBlock bool) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Reset(newDest io.Writer) { _ = "STUB: not implemented"; return }
