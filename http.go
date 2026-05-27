package brotli

import (
	"io"
	"net/http"
	"strings"
)

// HTTPCompressor chooses a compression method (brotli, gzip, or none) based on
// the Accept-Encoding header, sets the Content-Encoding header, and returns a
// WriteCloser that implements that compression. The Close method must be called
// before the current HTTP handler returns.
func HTTPCompressor(w http.ResponseWriter, r *http.Request) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

func HTTPCompressorWithLevel(w http.ResponseWriter, r *http.Request, level int) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// negotiateContentEncoding returns the best offered content encoding for the
// request's Accept-Encoding header. If two offers match with equal weight and
// then the offer earlier in the list is preferred. If no offers are
// acceptable, then "" is returned.
func negotiateContentEncoding(r *http.Request, offers []string) string {
	_ = "STUB: not implemented"
	return ""
}

// acceptSpec describes an Accept* header.
type acceptSpec struct {
	Value string
	Q     float64
}

// parseAccept parses Accept* headers.
func parseAccept(header http.Header, key string) (specs []acceptSpec) {
	_ = "STUB: not implemented"
	return nil
}

func skipSpace(s string) (rest string) { _ = "STUB: not implemented"; return "" }

func expectTokenSlash(s string) (token, rest string) { _ = "STUB: not implemented"; return "", "" }

func expectQuality(s string) (q float64, rest string) { _ = "STUB: not implemented"; return 0, "" }

// Octet types from RFC 2616.
var octetTypes [256]octetType

type octetType byte

const (
	isToken octetType = 1 << iota
	isSpace
)

func init() {
	// OCTET      = <any 8-bit sequence of data>
	// CHAR       = <any US-ASCII character (octets 0 - 127)>
	// CTL        = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
	// CR         = <US-ASCII CR, carriage return (13)>
	// LF         = <US-ASCII LF, linefeed (10)>
	// SP         = <US-ASCII SP, space (32)>
	// HT         = <US-ASCII HT, horizontal-tab (9)>
	// <">        = <US-ASCII double-quote mark (34)>
	// CRLF       = CR LF
	// LWS        = [CRLF] 1*( SP | HT )
	// TEXT       = <any OCTET except CTLs, but including LWS>
	// separators = "(" | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\" | <">
	//              | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP | HT
	// token      = 1*<any CHAR except CTLs or separators>
	// qdtext     = <any TEXT except <">>

	for c := 0; c < 256; c++ {
		var t octetType
		isCtl := c <= 31 || c == 127
		isChar := 0 <= c && c <= 127
		isSeparator := strings.ContainsRune(" \t\"(),/:;<=>?@[]\\{}", rune(c))
		if strings.ContainsRune(" \t\r\n", rune(c)) {
			t |= isSpace
		}
		if isChar && !isCtl && !isSeparator {
			t |= isToken
		}
		octetTypes[c] = t
	}
}
