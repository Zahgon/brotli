package brotli

/* Copyright 2016 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

func (*hashForgetfulChain) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (*hashForgetfulChain) StoreLookahead() uint {
	_ = "STUB: not implemented"

	/* HashBytes is the function that chooses the bucket to place the address in.*/
	return 0
}

func (h *hashForgetfulChain) HashBytes(data []byte) uint { _ = "STUB: not implemented"; return 0 }

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

type slot struct {
	delta uint16
	next  uint16
}

/*
A (forgetful) hash table to the data seen by the compressor, to

	help create backward references to previous data.

	Hashes are stored in chains which are bucketed to groups. Group of chains
	share a storage "bank". When more than "bank size" chain nodes are added,
	oldest nodes are replaced; this way several chains may share a tail.
*/
type hashForgetfulChain struct {
	hasherCommon

	bucketBits              uint
	numBanks                uint
	bankBits                uint
	numLastDistancesToCheck int

	addr          []uint32
	head          []uint16
	tiny_hash     [65536]byte
	banks         [][]slot
	free_slot_idx []uint16
	max_hops      uint
}

func (h *hashForgetfulChain) Initialize(params *encoderParams) { _ = "STUB: not implemented"; return }

func (h *hashForgetfulChain) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	return
}

/* Partial preparation is 100 times slower (per socket). */

/* See InitEmpty comment. */

/* Fill |addr| array with 0xCCCCCCCC value. Because of wrapping, position
   processed by hasher never reaches 3GB + 64M; this makes all new chains
   to be terminated after the first node. */

/*
Look at 4 bytes at &data[ix & mask]. Compute a hash from these, and prepend

	node to corresponding chain; also update tiny_hash for current position.
*/
func (h *hashForgetfulChain) Store(data []byte, mask uint, ix uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashForgetfulChain) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashForgetfulChain) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ring_buffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

/* Prepare the hashes for three last bytes of the last write.
   These could not be calculated before, since they require knowledge
   of both the previous and the current block. */

func (h *hashForgetfulChain) PrepareDistanceCache(distance_cache []int) {
	_ = "STUB: not implemented"
	return
}

/*
Find a longest backward match of &data[cur_ix] up to the length of

	max_length and stores the position cur_ix in the hash table.

	REQUIRES: PrepareDistanceCachehashForgetfulChain must be invoked for current distance cache
	          values; if this method is invoked repeatedly with the same distance
	          cache values, it is enough to invoke PrepareDistanceCachehashForgetfulChain once.

	Does not look for matches longer than max_length.
	Does not look for matches further away than max_backward.
	Writes the best match into |out|.
	|out|->score is updated only if a better match is found.
*/
func (h *hashForgetfulChain) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}

/* Don't accept a short copy from far away. */

/* Try last distance first. */

/* For distance code 0 we want to consider 2-byte matches. */

/* Comparing for >= 3 does not change the semantics, but just saves
   for a few unnecessary binary logarithms in backward reference
   score, since we are not interested in such short matches. */
