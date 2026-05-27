package brotli

/* Copyright 2010 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/*
A (forgetful) hash table to the data seen by the compressor, to

	help create backward references to previous data.

	This is a hash map of fixed size (bucket_size_) to a ring buffer of
	fixed size (block_size_). The ring buffer contains the last block_size_
	index positions of the given hash key in the compressed data.
*/
func (*h6) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (*h6) StoreLookahead() uint {
	_ = "STUB: not implemented"

	/* HashBytes is the function that chooses the bucket to place the address in. */
	return 0
}

func hashBytesH6(data []byte, mask uint64, shift int) uint32 { _ = "STUB: not implemented"; return 0 }

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

type h6 struct {
	hasherCommon
	bucket_size_ uint
	block_size_  uint
	hash_shift_  int
	hash_mask_   uint64
	block_mask_  uint32
	num          []uint16
	buckets      []uint32
}

func (h *h6) Initialize(params *encoderParams) { _ = "STUB: not implemented"; return }

func (h *h6) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	return
}

/* Partial preparation is 100 times slower (per socket). */

/*
Look at 4 bytes at &data[ix & mask].

	Compute a hash from these, and store the value of ix at that position.
*/
func (h *h6) Store(data []byte, mask uint, ix uint) { _ = "STUB: not implemented"; return }

func (h *h6) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *h6) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

/* Prepare the hashes for three last bytes of the last write.
   These could not be calculated before, since they require knowledge
   of both the previous and the current block. */

func (h *h6) PrepareDistanceCache(distance_cache []int) { _ = "STUB: not implemented"; return }

/*
Find a longest backward match of &data[cur_ix] up to the length of

	max_length and stores the position cur_ix in the hash table.

	REQUIRES: PrepareDistanceCacheH6 must be invoked for current distance cache
	          values; if this method is invoked repeatedly with the same distance
	          cache values, it is enough to invoke PrepareDistanceCacheH6 once.

	Does not look for matches longer than max_length.
	Does not look for matches further away than max_backward.
	Writes the best match into |out|.
	|out|->score is updated only if a better match is found.
*/
func (h *h6) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}

/* Don't accept a short copy from far away. */

/* Try last distance first. */

/* Comparing for >= 2 does not change the semantics, but just saves for
   a few unnecessary binary logarithms in backward reference score,
   since we are not interested in such short matches. */

/* Comparing for >= 3 does not change the semantics, but just saves
   for a few unnecessary binary logarithms in backward reference
   score, since we are not interested in such short matches. */
