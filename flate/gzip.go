package flate

import (
	"github.com/andybalholm/brotli/matchfinder"
)

func NewGZIPEncoder() matchfinder.Encoder {
	_ = "STUB: not implemented"
	return *new(matchfinder.Encoder)
}

type gzipEncoder struct {
	f           matchfinder.Encoder
	length      uint32
	crc         uint32
	wroteHeader bool
}

func (g *gzipEncoder) Reset() { _ = "STUB: not implemented"; return }

func appendUint32(dst []byte, n uint32) []byte { _ = "STUB: not implemented"; return nil }

func (g *gzipEncoder) Encode(dst []byte, src []byte, matches []matchfinder.Match, lastBlock bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// magic number
// CM = flate
// FLG

// XFL
// OS (unspecified)
