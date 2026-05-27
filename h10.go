package brotli

/* Copyright 2016 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

func (*h10) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (*h10) StoreLookahead() uint { _ = "STUB: not implemented"; return 0 }

func hashBytesH10(data []byte) uint32 { _ = "STUB: not implemented"; return 0 }

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

/*
A (forgetful) hash table where each hash bucket contains a binary tree of

	sequences whose first 4 bytes share the same hash code.
	Each sequence is 128 long and is identified by its starting
	position in the input data. The binary tree is sorted by the lexicographic
	order of the sequences, and it is also a max-heap with respect to the
	starting positions.
*/
type h10 struct {
	hasherCommon
	window_mask_ uint
	buckets_     [1 << 17]uint32
	invalid_pos_ uint32
	forest       []uint32
}

func (h *h10) Initialize(params *encoderParams) { _ = "STUB: not implemented"; return }

func (h *h10) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	return
}

func leftChildIndexH10(self *h10, pos uint) uint { _ = "STUB: not implemented"; return 0 }

func rightChildIndexH10(self *h10, pos uint) uint { _ = "STUB: not implemented"; return 0 }

/*
Stores the hash of the next 4 bytes and in a single tree-traversal, the

	hash bucket's binary tree is searched for matches and is re-rooted at the
	current position.

	If less than 128 data is available, the hash bucket of the
	current position is searched for matches, but the state of the hash table
	is not changed, since we can not know the final sorting order of the
	current (incomplete) sequence.

	This function must be called with increasing cur_ix positions.
*/
func storeAndFindMatchesH10(self *h10, data []byte, cur_ix uint, ring_buffer_mask uint, max_length uint, max_backward uint, best_len *uint, matches []backwardMatch) []backwardMatch {
	_ = "STUB: not implemented"
	return nil
}

/* The forest index of the rightmost node of the left subtree of the new
   root, updated as we traverse and re-root the tree of the hash bucket. */

/* The forest index of the leftmost node of the right subtree of the new
   root, updated as we traverse and re-root the tree of the hash bucket. */

/* The match length of the rightmost node of the left subtree of the new
   root, updated as we traverse and re-root the tree of the hash bucket. */

/* The match length of the leftmost node of the right subtree of the new
   root, updated as we traverse and re-root the tree of the hash bucket. */

/*
Finds all backward matches of &data[cur_ix & ring_buffer_mask] up to the

	length of max_length and stores the position cur_ix in the hash table.

	Sets *num_matches to the number of matches found, and stores the found
	matches in matches[0] to matches[*num_matches - 1]. The matches will be
	sorted by strictly increasing length and (non-strictly) increasing
	distance.
*/
func findAllMatchesH10(handle *h10, dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, cur_ix uint, max_length uint, max_backward uint, gap uint, params *encoderParams, matches []backwardMatch) uint {
	_ = "STUB: not implemented"
	return 0
}

/*
Stores the hash of the next 4 bytes and re-roots the binary tree at the

	current sequence, without returning any matches.
	REQUIRES: ix + 128 <= end-of-current-block
*/
func (h *h10) Store(data []byte, mask uint, ix uint) { _ = "STUB: not implemented"; return }

/* Maximum distance is window size - 16, see section 9.1. of the spec. */

func (h *h10) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *h10) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

/* Store the last `128 - 1` positions in the hasher.
   These could not be calculated before, since they require knowledge
   of both the previous and the current block. */

/* Maximum distance is window size - 16, see section 9.1. of the spec.
   Furthermore, we have to make sure that we don't look further back
   from the start of the next block than the window size, otherwise we
   could access already overwritten areas of the ring-buffer. */

/* We know that i + 128 <= position + num_bytes, i.e. the
   end of the current block and that we have at least
   128 tail in the ring-buffer. */

/* MAX_NUM_MATCHES == 64 + MAX_TREE_SEARCH_DEPTH */
const maxNumMatchesH10 = 128

func (*h10) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}

func (*h10) PrepareDistanceCache(distance_cache []int) { _ = "STUB: not implemented"; return }
