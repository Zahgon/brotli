package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Bit reading helpers */

const shortFillBitWindowRead = (8 >> 1)

var kBitMask = [33]uint32{
	0x00000000,
	0x00000001,
	0x00000003,
	0x00000007,
	0x0000000F,
	0x0000001F,
	0x0000003F,
	0x0000007F,
	0x000000FF,
	0x000001FF,
	0x000003FF,
	0x000007FF,
	0x00000FFF,
	0x00001FFF,
	0x00003FFF,
	0x00007FFF,
	0x0000FFFF,
	0x0001FFFF,
	0x0003FFFF,
	0x0007FFFF,
	0x000FFFFF,
	0x001FFFFF,
	0x003FFFFF,
	0x007FFFFF,
	0x00FFFFFF,
	0x01FFFFFF,
	0x03FFFFFF,
	0x07FFFFFF,
	0x0FFFFFFF,
	0x1FFFFFFF,
	0x3FFFFFFF,
	0x7FFFFFFF,
	0xFFFFFFFF,
}

func bitMask(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

type bitReader struct {
	val_      uint64
	bit_pos_  uint32
	input     []byte
	input_len uint
	byte_pos  uint
}

type bitReaderState struct {
	val_      uint64
	bit_pos_  uint32
	input     []byte
	input_len uint
	byte_pos  uint
}

/* Initializes the BrotliBitReader fields. */

/*
Ensures that accumulator is not empty.

	May consume up to sizeof(brotli_reg_t) - 1 bytes of input.
	Returns false if data is required but there is no input available.
	For BROTLI_ALIGNED_READ this function also prepares bit reader for aligned
	reading.
*/
func bitReaderSaveState(from *bitReader, to *bitReaderState) { _ = "STUB: not implemented"; return }

func bitReaderRestoreState(to *bitReader, from *bitReaderState) { _ = "STUB: not implemented"; return }

func getAvailableBits(br *bitReader) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Returns amount of unread bytes the bit reader still has buffered from the

	BrotliInput, including whole bytes in br->val_.
*/
func getRemainingBytes(br *bitReader) uint { _ = "STUB: not implemented"; return 0 }

/*
Checks if there is at least |num| bytes left in the input ring-buffer

	(excluding the bits remaining in br->val_).
*/
func checkInputAmount(br *bitReader, num uint) bool { _ = "STUB: not implemented"; return false }

/*
Guarantees that there are at least |n_bits| + 1 bits in accumulator.

	Precondition: accumulator contains at least 1 bit.
	|n_bits| should be in the range [1..24] for regular build. For portable
	non-64-bit little-endian build only 16 bits are safe to request.
*/
func fillBitWindow(br *bitReader, n_bits uint32) { _ = "STUB: not implemented"; return }

/* here same as -= 32 because of the if condition */

/*
Mostly like BrotliFillBitWindow, but guarantees only 16 bits and reads no

	more than BROTLI_SHORT_FILL_BIT_WINDOW_READ bytes of input.
*/
func fillBitWindow16(br *bitReader) { _ = "STUB: not implemented"; return }

/*
Tries to pull one byte of input to accumulator.

	Returns false if there is no input available.
*/
func pullByte(br *bitReader) bool { _ = "STUB: not implemented"; return false }

/*
Returns currently available bits.

	The number of valid bits could be calculated by BrotliGetAvailableBits.
*/
func getBitsUnmasked(br *bitReader) uint64 { _ = "STUB: not implemented"; return 0 }

/*
Like BrotliGetBits, but does not mask the result.

	The result contains at least 16 valid bits.
*/
func get16BitsUnmasked(br *bitReader) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Returns the specified number of bits from |br| without advancing bit

	position.
*/
func getBits(br *bitReader, n_bits uint32) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Tries to peek the specified amount of bits. Returns false, if there

	is not enough input.
*/
func safeGetBits(br *bitReader, n_bits uint32, val *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/* Advances the bit pos by |n_bits|. */
func dropBits(br *bitReader, n_bits uint32) { _ = "STUB: not implemented"; return }

func bitReaderUnload(br *bitReader) { _ = "STUB: not implemented"; return }

/*
Reads the specified number of bits from |br| and advances the bit pos.

	Precondition: accumulator MUST contain at least |n_bits|.
*/
func takeBits(br *bitReader, n_bits uint32, val *uint32) { _ = "STUB: not implemented"; return }

/*
Reads the specified number of bits from |br| and advances the bit pos.

	Assumes that there is enough input to perform BrotliFillBitWindow.
*/
func readBits(br *bitReader, n_bits uint32) uint32 { _ = "STUB: not implemented"; return 0 }

/*
Tries to read the specified amount of bits. Returns false, if there

	is not enough input. |n_bits| MUST be positive.
*/
func safeReadBits(br *bitReader, n_bits uint32, val *uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/*
Advances the bit reader position to the next byte boundary and verifies

	that any skipped bits are set to zero.
*/
func bitReaderJumpToByteBoundary(br *bitReader) bool { _ = "STUB: not implemented"; return false }

/*
Copies remaining input bytes stored in the bit reader to the output. Value

	|num| may not be larger than BrotliGetRemainingBytes. The bit reader must be
	warmed up again after this.
*/
func copyBytes(dest []byte, br *bitReader, num uint) { _ = "STUB: not implemented"; return }

func initBitReader(br *bitReader) { _ = "STUB: not implemented"; return }

func warmupBitReader(br *bitReader) bool {
	_ = "STUB: not implemented"
	/* Fixing alignment after unaligned BrotliFillWindow would result accumulator
	   overflow. If unalignment is caused by BrotliSafeReadBits, then there is
	   enough space in accumulator to fix alignment. */return false
}
