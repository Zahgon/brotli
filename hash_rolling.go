package brotli

/* Copyright 2018 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* NOTE: this hasher does not search in the dictionary. It is used as
   backup-hasher, the main hasher already searches in it. */

const kRollingHashMul32 uint32 = 69069

const kInvalidPosHashRolling uint32 = 0xffffffff

/*
This hasher uses a longer forward length, but returning a higher value here

	will hurt compression by the main hasher when combined with a composite
	hasher. The hasher tests for forward itself instead.
*/
func (*hashRolling) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (*hashRolling) StoreLookahead() uint {
	_ = "STUB: not implemented"

	/*
	   Computes a code from a single byte. A lookup table of 256 values could be

	   	used, but simply adding 1 works about as good.
	*/return 0
}

func (*hashRolling) HashByte(b byte) uint32 { _ = "STUB: not implemented"; return 0 }

func (h *hashRolling) HashRollingFunctionInitial(state uint32, add byte, factor uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (h *hashRolling) HashRollingFunction(state uint32, add byte, rem byte, factor uint32, factor_remove uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

/*
Rolling hash for long distance long string matches. Stores one position

	per bucket, bucket key is computed over a long region.
*/
type hashRolling struct {
	hasherCommon

	jump int

	state         uint32
	table         []uint32
	next_ix       uint
	factor        uint32
	factor_remove uint32
}

func (h *hashRolling) Initialize(params *encoderParams) { _ = "STUB: not implemented"; return }

/* Compute the factor of the oldest byte to remove: factor**steps modulo
   0xffffffff (the multiplications rely on 32-bit overflow) */

func (h *hashRolling) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	/* Too small size, cannot use this hasher. */ return
}

func (*hashRolling) Store(data []byte, mask uint, ix uint) { _ = "STUB: not implemented"; return }

func (*hashRolling) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashRolling) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ring_buffer_mask uint) {
	_ = "STUB: not implemented"
	return

	/* In this case we must re-initialize the hasher from scratch from the
	   current position. */
}

/* wrapping around ringbuffer not handled. */

func (*hashRolling) PrepareDistanceCache(distance_cache []int) { _ = "STUB: not implemented"; return }

func (h *hashRolling) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}

/* Not enough lookahead */

/* The cast to 32-bit makes backward distances up to 4GB work even
   if cur_ix is above 4GB, despite using 32-bit values in the table. */
