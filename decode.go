package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

const (
	decoderResultError           = 0
	decoderResultSuccess         = 1
	decoderResultNeedsMoreInput  = 2
	decoderResultNeedsMoreOutput = 3
)

/**
 * Error code for detailed logging / production debugging.
 *
 * See ::BrotliDecoderGetErrorCode and ::BROTLI_LAST_ERROR_CODE.
 */
const (
	decoderNoError                          = 0
	decoderSuccess                          = 1
	decoderNeedsMoreInput                   = 2
	decoderNeedsMoreOutput                  = 3
	decoderErrorFormatExuberantNibble       = -1
	decoderErrorFormatReserved              = -2
	decoderErrorFormatExuberantMetaNibble   = -3
	decoderErrorFormatSimpleHuffmanAlphabet = -4
	decoderErrorFormatSimpleHuffmanSame     = -5
	decoderErrorFormatClSpace               = -6
	decoderErrorFormatHuffmanSpace          = -7
	decoderErrorFormatContextMapRepeat      = -8
	decoderErrorFormatBlockLength1          = -9
	decoderErrorFormatBlockLength2          = -10
	decoderErrorFormatTransform             = -11
	decoderErrorFormatDictionary            = -12
	decoderErrorFormatWindowBits            = -13
	decoderErrorFormatPadding1              = -14
	decoderErrorFormatPadding2              = -15
	decoderErrorFormatDistance              = -16
	decoderErrorDictionaryNotSet            = -19
	decoderErrorInvalidArguments            = -20
	decoderErrorAllocContextModes           = -21
	decoderErrorAllocTreeGroups             = -22
	decoderErrorAllocContextMap             = -25
	decoderErrorAllocRingBuffer1            = -26
	decoderErrorAllocRingBuffer2            = -27
	decoderErrorAllocBlockTypeTrees         = -30
	decoderErrorUnreachable                 = -31
)

const huffmanTableBits = 8

const huffmanTableMask = 0xFF

/*
We need the slack region for the following reasons:
  - doing up to two 16-byte copies for fast backward copying
  - inserting transformed dictionary word (5 prefix + 24 base + 8 suffix)
*/
const kRingBufferWriteAheadSlack uint32 = 42

var kCodeLengthCodeOrder = [codeLengthCodes]byte{1, 2, 3, 4, 0, 5, 17, 6, 16, 7, 8, 9, 10, 11, 12, 13, 14, 15}

/* Static prefix code for the complex code length code lengths. */
var kCodeLengthPrefixLength = [16]byte{2, 2, 2, 3, 2, 2, 2, 4, 2, 2, 2, 3, 2, 2, 2, 4}

var kCodeLengthPrefixValue = [16]byte{0, 4, 3, 2, 0, 4, 3, 1, 0, 4, 3, 2, 0, 4, 3, 5}

/* Saves error code and converts it to BrotliDecoderResult. */
func saveErrorCode(s *Reader, e int) int { _ = "STUB: not implemented"; return 0 }

/*
Decodes WBITS by reading 1 - 7 bits, or 0x11 for "Large Window Brotli".

	Precondition: bit-reader accumulator has at least 8 bits.
*/
func decodeWindowBits(s *Reader, br *bitReader) int { _ = "STUB: not implemented"; return 0 }

/* Decodes a number in the range [0..255], by reading 1 - 11 bits. */
func decodeVarLenUint8(s *Reader, br *bitReader, value *uint32) int {
	_ = "STUB: not implemented"
	return 0
}

/* Fall through. */

/* Use output value as a temporary storage. It MUST be persisted. */

/* Fall through. */

/* Decodes a metablock length and flags by reading 2 - 31 bits. */
func decodeMetaBlockLength(s *Reader, br *bitReader) int { _ = "STUB: not implemented"; return 0 }

/* Fall through. */

/* Fall through. */

/* Fall through. */

/* Fall through. */

/* Fall through. */

/* Fall through. */

/*
Decodes the Huffman code.

	This method doesn't read data from the bit reader, BUT drops the amount of
	bits that correspond to the decoded symbol.
	bits MUST contain at least 15 (BROTLI_HUFFMAN_MAX_CODE_LENGTH) valid bits.
*/
func decodeSymbol(bits uint32, table []huffmanCode, br *bitReader) uint32 {
	_ = "STUB: not implemented"
	return 0
}

/*
Reads and decodes the next Huffman code from bit-stream.

	This method peeks 16 bits of input and drops 0 - 15 of them.
*/
func readSymbol(table []huffmanCode, br *bitReader) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Same as DecodeSymbol, but it is known that there is less than 15 bits of

	input are currently available.
*/
func safeDecodeSymbol(table []huffmanCode, br *bitReader, result *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/* No valid bits at all. */

/* Not enough bits for the first level. */

/* Not enough bits to move to the second level. */

/* Speculatively drop HUFFMAN_TABLE_BITS. */

/* Not enough bits for the second level. */

func safeReadSymbol(table []huffmanCode, br *bitReader, result *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/* Makes a look-up in first level Huffman table. Peeks 8 bits. */
func preloadSymbol(safe int, table []huffmanCode, br *bitReader, bits *uint32, value *uint32) {
	_ = "STUB: not implemented"
	return
}

/*
Decodes the next Huffman code using data prepared by PreloadSymbol.

	Reads 0 - 15 bits. Also peeks 8 following bits.
*/
func readPreloadedSymbol(table []huffmanCode, br *bitReader, bits *uint32, value *uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func log2Floor(x uint32) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Reads (s->symbol + 1) symbols.

	Totally 1..4 symbols are read, 1..11 bits each.
	The list of symbols MUST NOT contain duplicates.
*/
func readSimpleHuffmanSymbols(alphabet_size uint32, max_symbol uint32, s *Reader) int {
	_ = "STUB: not implemented"
	return 0
}

/* max_bits == 1..11; symbol == 0..3; 1..44 bits will be read. */

/*
Process single decoded symbol code length:

	A) reset the repeat variable
	B) remember code length (if it is not 0)
	C) extend corresponding index-chain
	D) reduce the Huffman space
	E) update the histogram
*/
func processSingleCodeLength(code_len uint32, symbol *uint32, repeat *uint32, space *uint32, prev_code_len *uint32, symbol_lists symbolList, code_length_histo []uint16, next_symbol []int) {
	_ = "STUB: not implemented"
	return
}

/* code_len == 1..15 */

/*
Process repeated symbol code length.

	 A) Check if it is the extension of previous repeat sequence; if the decoded
	    value is not BROTLI_REPEAT_PREVIOUS_CODE_LENGTH, then it is a new
	    symbol-skip
	 B) Update repeat variable
	 C) Check if operation is feasible (fits alphabet)
	 D) For each symbol do the same operations as in ProcessSingleCodeLength

	PRECONDITION: code_len == BROTLI_REPEAT_PREVIOUS_CODE_LENGTH or
	              code_len == BROTLI_REPEAT_ZERO_CODE_LENGTH
*/
func processRepeatedCodeLength(code_len uint32, repeat_delta uint32, alphabet_size uint32, symbol *uint32, repeat *uint32, space *uint32, prev_code_len *uint32, repeat_code_len *uint32, symbol_lists symbolList, code_length_histo []uint16, next_symbol []int) {
	_ = "STUB: not implemented"
	return
	/* for BROTLI_REPEAT_ZERO_CODE_LENGTH */ /* for BROTLI_REPEAT_ZERO_CODE_LENGTH */
}

/* Reads and decodes symbol codelengths. */
func readSymbolCodeLengths(alphabet_size uint32, s *Reader) int {
	_ = "STUB: not implemented"
	return 0
}

/* Use 1..5 bits. */
/* code_len == 0..17 */

/* code_len == 16..17, extra_bits == 2..3 */

func safeReadSymbolCodeLengths(alphabet_size uint32, s *Reader) int {
	_ = "STUB: not implemented"
	return 0
}

/* code_len == 0..17 */

/* code_len == 16..17, extra_bits == 2..3 */

/*
Reads and decodes 15..18 codes using static prefix code.

	Each code is 2..4 bits long. In total 30..72 bits are used.
*/
func readCodeLengthCodeLengths(s *Reader) int { _ = "STUB: not implemented"; return 0 }

/* space is 0 or wrapped around. */

/*
Decodes the Huffman tables.

	There are 2 scenarios:
	 A) Huffman code contains only few symbols (1..4). Those symbols are read
	    directly; their code lengths are defined by the number of symbols.
	    For this scenario 4 - 49 bits will be read.

	 B) 2-phase decoding:
	 B.1) Small Huffman table is decoded; it is specified with code lengths
	      encoded with predefined entropy code. 32 - 74 bits are used.
	 B.2) Decoded table is used to decode code lengths of symbols in resulting
	      Huffman table. In worst case 3520 bits are read.
*/
func readHuffmanCode(alphabet_size uint32, max_symbol uint32, table []huffmanCode, opt_table_size *uint32, s *Reader) int {
	_ = "STUB: not implemented"
	return 0

	/* Unnecessary masking, but might be good for safety. */
}

/* State machine. */

/* The value is used as follows:
   1 for simple code;
   0 for no skipping, 2 skips 2 code lengths, 3 skips 3 code lengths */

/* num_codes */

/* Read symbols, codes & code lengths directly. */

/* num_symbols */

/* Decode Huffman-coded code lengths. */

/* Decodes a block length by reading 3..39 bits. */
func readBlockLength(table []huffmanCode, br *bitReader) uint32 {
	_ = "STUB: not implemented"
	return 0
}

/* nbits == 2..24 */

/*
WARNING: if state is not BROTLI_STATE_READ_BLOCK_LENGTH_NONE, then

	reading can't be continued with ReadBlockLength.
*/
func safeReadBlockLength(s *Reader, result *uint32, table []huffmanCode, br *bitReader) bool {
	_ = "STUB: not implemented"
	return false
}

/* nbits == 2..24 */

/*
Transform:

 1. initialize list L with values 0, 1,... 255

 2. For each input element X:
    2.1) let Y = L[X]
    2.2) remove X-th element from L
    2.3) prepend Y to L
    2.4) append Y to output

    In most cases max(Y) <= 7, so most of L remains intact.
    To reduce the cost of initialization, we reuse L, remember the upper bound
    of Y values, and reinitialize only first elements in L.

    Most of input values are 0 and 1. To reduce number of branches, we replace
    inner for loop with do-while.
*/
func inverseMoveToFrontTransform(v []byte, v_len uint32, state *Reader) {
	_ = "STUB: not implemented"
	return
}

/* Transform the input. */

/* Decodes a series of Huffman table using ReadHuffmanCode function. */
func huffmanTreeGroupDecode(group *huffmanTreeGroup, s *Reader) int {
	_ = "STUB: not implemented"
	return 0
}

/*
Decodes a context map.

	Decoding is done in 4 phases:
	 1) Read auxiliary information (6..16 bits) and allocate memory.
	    In case of trivial context map, decoding is finished at this phase.
	 2) Decode Huffman table using ReadHuffmanCode function.
	    This table will be used for reading context map items.
	 3) Read context map items; "0" values could be run-length encoded.
	 4) Optionally, apply InverseMoveToFront transform to the resulting map.
*/
func decodeContextMap(context_map_size uint32, num_htrees *uint32, context_map_arg *[]byte, s *Reader) int {
	_ = "STUB: not implemented"
	return 0
}

/* Fall through. */

/* In next stage ReadHuffmanCode uses at least 4 bits, so it is safe
   to peek 4 bits ahead. */

/* Use RLE for zeros. */

/* Fall through. */

/* Fall through. */

/* RLE sub-stage. */

/*
Decodes a command or literal and updates block type ring-buffer.

	Reads 3..54 bits.
*/
func decodeBlockTypeAndLength(safe int, s *Reader, tree_type int) bool {
	_ = "STUB: not implemented"
	return false
}

/* Read 0..15 + 3..39 bits. */

func detectTrivialLiteralBlockTypes(s *Reader) { _ = "STUB: not implemented"; return }

func prepareLiteralDecoding(s *Reader) { _ = "STUB: not implemented"; return }

/*
Decodes the block type and updates the state for literal context.

	Reads 3..54 bits.
*/
func decodeLiteralBlockSwitchInternal(safe int, s *Reader) bool {
	_ = "STUB: not implemented"
	return false
}

func decodeLiteralBlockSwitch(s *Reader) { _ = "STUB: not implemented"; return }

func safeDecodeLiteralBlockSwitch(s *Reader) bool { _ = "STUB: not implemented"; return false }

/*
Block switch for insert/copy length.

	Reads 3..54 bits.
*/
func decodeCommandBlockSwitchInternal(safe int, s *Reader) bool {
	_ = "STUB: not implemented"
	return false
}

func decodeCommandBlockSwitch(s *Reader) { _ = "STUB: not implemented"; return }

func safeDecodeCommandBlockSwitch(s *Reader) bool { _ = "STUB: not implemented"; return false }

/*
Block switch for distance codes.

	Reads 3..54 bits.
*/
func decodeDistanceBlockSwitchInternal(safe int, s *Reader) bool {
	_ = "STUB: not implemented"
	return false
}

func decodeDistanceBlockSwitch(s *Reader) { _ = "STUB: not implemented"; return }

func safeDecodeDistanceBlockSwitch(s *Reader) bool { _ = "STUB: not implemented"; return false }

func unwrittenBytes(s *Reader, wrap bool) uint { _ = "STUB: not implemented"; return 0 }

/*
Dumps output.

	Returns BROTLI_DECODER_NEEDS_MORE_OUTPUT only if there is more output to push
	and either ring-buffer is as big as window size, or |force| is true.
*/
func writeRingBuffer(s *Reader, available_out *uint, next_out *[]byte, total_out *uint, force bool) int {
	_ = "STUB: not implemented"
	return 0
}

/* Wrap ring buffer only if it has reached its maximal size. */

func wrapRingBuffer(s *Reader) { _ = "STUB: not implemented"; return }

/*
Allocates ring-buffer.

	s->ringbuffer_size MUST be updated by BrotliCalculateRingBufferSize before
	this function is called.

	Last two bytes of ring-buffer are initialized to 0, so context calculation
	could be done uniformly for the first two and all other positions.
*/
func ensureRingBuffer(s *Reader) bool { _ = "STUB: not implemented"; return false }

func copyUncompressedBlockToOutput(available_out *uint, next_out *[]byte, total_out *uint, s *Reader) int {
	_ = "STUB: not implemented"
	/* TODO: avoid allocation for single uncompressed block. */ return 0
}

/* State machine */

/* Copy remaining bytes from s->br.buf_ to ring-buffer. */

/*
Calculates the smallest feasible ring buffer.

	If we know the data size is small, do not allocate more ring buffer
	size than needed to reduce memory usage.

	When this method is called, metablock size and flags MUST be decoded.
*/
func calculateRingBufferSize(s *Reader) { _ = "STUB: not implemented"; return }

/* We need at least 2 bytes of ring buffer size to get the last two
   bytes for context from there */

/* If maximum is already reached, no further extension is retired. */

/* Metadata blocks does not touch ring buffer. */

/* Reduce ring buffer size to save memory when server is unscrupulous.
   In worst case memory usage might be 1.5x bigger for a short period of
   ring buffer reallocation. */

/* Reads 1..256 2-bit context modes. */
func readContextModes(s *Reader) int { _ = "STUB: not implemented"; return 0 }

func takeDistanceFromRingBuffer(s *Reader) { _ = "STUB: not implemented"; return }

/* Compensate double distance-ring-buffer roll for dictionary items. */

/* kDistanceShortCodeIndexOffset has 2-bit values from LSB:
   3, 2, 1, 0, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2 */

/* kDistanceShortCodeValueOffset has 2-bit values from LSB:
   -0, 0,-0, 0,-1, 1,-2, 2,-3, 3,-1, 1,-2, 2,-3, 3 */

/* A huge distance will cause a () soon.
   This is a little faster than failing here. */

func safeReadBitsMaybeZero(br *bitReader, n_bits uint32, val *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/* Precondition: s->distance_code < 0. */
func readDistanceInternal(safe int, s *Reader, br *bitReader) bool {
	_ = "STUB: not implemented"
	return false
}

/* Convert the distance code to the actual distance by possibly
   looking up past distances from the s->ringbuffer. */

/* This branch also works well when s->distance_postfix_bits == 0. */

/* Restore precondition. */

func readDistance(s *Reader, br *bitReader) { _ = "STUB: not implemented"; return }

func safeReadDistance(s *Reader, br *bitReader) bool { _ = "STUB: not implemented"; return false }

func readCommandInternal(safe int, s *Reader, br *bitReader, insert_length *int) bool {
	_ = "STUB: not implemented"
	return false
}

func readCommand(s *Reader, br *bitReader, insert_length *int) { _ = "STUB: not implemented"; return }

func safeReadCommand(s *Reader, br *bitReader, insert_length *int) bool {
	_ = "STUB: not implemented"
	return false
}

func checkInputAmountMaybeSafe(safe int, br *bitReader, num uint) bool {
	_ = "STUB: not implemented"
	return false
}

func processCommandsInternal(safe int, s *Reader) int { _ = "STUB: not implemented"; return 0 }

/* Jump into state machine. */

/* 156 bits + 7 bytes */

/* Read the insert/copy length in the command. */

/* Read the literals in the command. */

/* 162 bits + 7 bytes */

/* 162 bits + 7 bytes */

/* Implicit distance case. */

/* Read distance code in the command, unless it was implicitly zero. */

/* Apply copy of LZ77 back-reference, or static dictionary reference if
   the distance is larger than the max LZ77 distance */

/* The maximum allowed distance is BROTLI_MAX_ALLOWED_DISTANCE = 0x7FFFFFFC.
   With this choice, no signed overflow can occur after decoding
   a special distance code (e.g., after adding 3 to the last distance). */

/* Compensate double distance-ring-buffer roll. */

/* Update the recent distances cache. */

/* There are 32+ bytes of slack in the ring-buffer allocation.
   Also, we have 16 short codes, that make these 16 bytes irrelevant
   in the ring-buffer. Let's copy over them as a first guess. */

/* Regions intersect. */

/* At least one region wraps. */

/* This branch covers about 45% cases.
   Fixed size short copy allows more compiler optimizations. */

/* Next metablock, if any. */

/* Next metablock, if any. */

func processCommands(s *Reader) int { _ = "STUB: not implemented"; return 0 }

func safeProcessCommands(s *Reader) int { _ = "STUB: not implemented"; return 0 }

/* Returns the maximum number of distance symbols which can only represent
   distances not exceeding BROTLI_MAX_ALLOWED_DISTANCE. */

var maxDistanceSymbol_bound = [maxNpostfix + 1]uint32{0, 4, 12, 28}
var maxDistanceSymbol_diff = [maxNpostfix + 1]uint32{73, 126, 228, 424}

func maxDistanceSymbol(ndirect uint32, npostfix uint32) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Invariant: input stream is never overconsumed:
  - invalid input implies that the whole stream is invalid -> any amount of
    input could be read and discarded
  - when result is "needs more input", then at least one more byte is REQUIRED
    to complete decoding; all input data MUST be consumed by decoder, so
    client could swap the input buffer
  - when result is "needs more output" decoder MUST ensure that it doesn't
    hold more than 7 bits in bit reader; this saves client from swapping input
    buffer ahead of time
  - when result is "success" decoder MUST return all unused data back to input
    buffer; this is possible because the invariant is held on enter
*/
func decoderDecompressStream(s *Reader, available_in *uint, next_in *[]byte, available_out *uint, next_out *[]byte) int {
	_ = "STUB: not implemented"
	return 0
}

/* Do not try to process further in a case of unrecoverable error. */

/* Just connect bit reader to input stream. */

/* At least one byte of input is required. More than one byte of input may
   be required to complete the transaction -> reading more data must be
   done in a loop -> do it in a main loop. */

/* State machine */

/* Error, needs more input/output. */

/* Pro-actively push output. */

/* WriteRingBuffer checks s->meta_block_remaining_len validity. */

/* Used with internal buffer. */

/* Successfully finished read transaction.
   Accumulator contains less than 8 bits, because internal buffer
   is expanded byte-by-byte until it is enough to complete read. */

/* Switch to input stream and restart. */

/* Not enough data in buffer, but can take one more byte from
   input stream. */

/* Retry with more data in buffer. */

/* Can't finish reading and no more input. */

/* Input stream doesn't contain enough input. */

/* Copy tail to internal buffer and return. */

/* Unreachable. */

/* Fail or needs more output. */

/* Just consumed the buffered input and produced some output. Otherwise
   it would result in "needs more input". Reset internal buffer. */

/* Using input stream in last iteration. When decoder switches to input
   stream it has less than 8 bits in accumulator, so it is safe to
   return unused accumulator bits there. */

/* Prepare to the first read. */

/* Decode window size. */
/* Reads 1..8 bits. */

/* Maximum distance, see section 9.1. of the spec. */
/* Fall through. */

/* Allocate memory for both block_type_trees and block_len_trees. */

/* Fall through. */

/* Fall through. */

/* Reads 2 - 31 bits. */

/* Read one byte and ignore it. */

/* Reads 1..11 bits. */

/* Next metablock, if any. */

/* BROTLI_STATE_COMMAND_INNER_WRITE */

func decoderHasMoreOutput(s *Reader) bool {
	_ = "STUB: not implemented"
	/* After unrecoverable error remaining output is considered nonsensical. */ return false
}

func decoderGetErrorCode(s *Reader) int { _ = "STUB: not implemented"; return 0 }

func decoderErrorString(c int) string { _ = "STUB: not implemented"; return "" }
