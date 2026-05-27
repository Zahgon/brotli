package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Heuristics for deciding about the UTF8-ness of strings. */

const kMinUTF8Ratio float64 = 0.75

/*
Returns 1 if at least min_fraction of the bytes between pos and

	pos + length in the (data, mask) ring-buffer is UTF8-encoded, otherwise
	returns 0.
*/
func parseAsUTF8(symbol *int, input []byte, size uint) uint {
	_ = "STUB: not implemented"
	/* ASCII */ return 0
}

/* 2-byte UTF8 */

/* 3-byte UFT8 */

/* 4-byte UFT8 */

/* Not UTF8, emit a special symbol above the UTF8-code space */

/* Returns 1 if at least min_fraction of the data is UTF8-encoded.*/
func isMostlyUTF8(data []byte, pos uint, mask uint, length uint, min_fraction float64) bool {
	_ = "STUB: not implemented"
	return false
}
