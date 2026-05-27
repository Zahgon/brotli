package brotli

func utf8Position(last uint, c uint, clamp uint) uint {
	_ = "STUB: not implemented"
	return

	/* Next one is the 'Byte 1' again. */
	0
}

/* Next one is the 'Byte 2' of utf-8 encoding. */

/* Let's decide over the last byte if this ends the sequence. */

/* Completed two or three byte coding. */ /* Next one is the 'Byte 3' of utf-8 encoding. */

func decideMultiByteStatsLevel(pos uint, len uint, mask uint, data []byte) uint {
	_ = "STUB: not implemented"
	return
	/* should be 2, but 1 compresses better. */ 0
}

func estimateBitCostsForLiteralsUTF8(pos uint, len uint, mask uint, data []byte, cost []float32) {
	_ = "STUB: not implemented"
	return
}

/* Bootstrap histograms. */

/* max_utf8 is 0 (normal ASCII single byte modeling),
   1 (for 2-byte UTF-8 modeling), or 2 (for 3-byte UTF-8 modeling). */

/* Compute bit costs with sliding window. */

/* Remove a byte in the past. */

/* Add a byte in the future. */

/* Make the first bytes more expensive -- seems to help, not sure why.
   Perhaps because the entropy source is changing its properties
   rapidly in the beginning of the file, perhaps because the beginning
   of the data is a statistical "anomaly". */

func estimateBitCostsForLiterals(pos uint, len uint, mask uint, data []byte, cost []float32) {
	_ = "STUB: not implemented"
	return
}

/* Bootstrap histogram. */

/* Compute bit costs with sliding window. */

/* Remove a byte in the past. */

/* Add a byte in the future. */
