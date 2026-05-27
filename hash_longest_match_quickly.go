package brotli

/* Copyright 2010 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* For BUCKET_SWEEP == 1, enabling the dictionary lookup makes compression
   a little faster (0.5% - 1%) and it compresses 0.15% better on small text
   and HTML inputs. */

func (*hashLongestMatchQuickly) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (*hashLongestMatchQuickly) StoreLookahead() uint {
	_ = "STUB: not implemented"

	/*
	   HashBytes is the function that chooses the bucket to place

	   	the address in. The HashLongestMatch and hashLongestMatchQuickly
	   	classes have separate, different implementations of hashing.
	*/return 0
}

func (h *hashLongestMatchQuickly) HashBytes(data []byte) uint32 {
	_ = "STUB: not implemented"
	return 0
}

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

/*
A (forgetful) hash table to the data seen by the compressor, to

	help create backward references to previous data.

	This is a hash map of fixed size (1 << 16). Starting from the
	given index, 1 buckets are used to store values of a key.
*/
type hashLongestMatchQuickly struct {
	hasherCommon

	bucketBits    uint
	bucketSweep   int
	hashLen       uint
	useDictionary bool

	buckets []uint32
}

func (h *hashLongestMatchQuickly) Initialize(params *encoderParams) {
	_ = "STUB: not implemented"
	return
}

func (h *hashLongestMatchQuickly) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	return
}

/* Partial preparation is 100 times slower (per socket). */

/* It is not strictly necessary to fill this buffer here, but
   not filling will make the results of the compression stochastic
   (but correct). This is because random data would cause the
   system to find accidentally good backward references here and there. */

/*
Look at 5 bytes at &data[ix & mask].

	Compute a hash from these, and store the value somewhere within
	[ix .. ix+3].
*/
func (h *hashLongestMatchQuickly) Store(data []byte, mask uint, ix uint) {
	_ = "STUB: not implemented"
	return
}

/* Wiggle the value with the bucket sweep range. */

func (h *hashLongestMatchQuickly) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashLongestMatchQuickly) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

/* Prepare the hashes for three last bytes of the last write.
   These could not be calculated before, since they require knowledge
   of both the previous and the current block. */

func (*hashLongestMatchQuickly) PrepareDistanceCache(distance_cache []int) {
	_ = "STUB: not implemented"

	/*
	   Find a longest backward match of &data[cur_ix & ring_buffer_mask]

	   	up to the length of max_length and stores the position cur_ix in the
	   	hash table.

	   	Does not look for matches longer than max_length.
	   	Does not look for matches further away than max_backward.
	   	Writes the best match into |out|.
	   	|out|->score is updated only if a better match is found.
	*/return
}

func (h *hashLongestMatchQuickly) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}

/* Only one to look for, don't bother to prepare for a loop. */
