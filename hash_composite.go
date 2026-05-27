package brotli

/* Copyright 2018 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

func (h *hashComposite) HashTypeLength() uint { _ = "STUB: not implemented"; return 0 }

func (h *hashComposite) StoreLookahead() uint { _ = "STUB: not implemented"; return 0 }

/*
Composite hasher: This hasher allows to combine two other hashers, HASHER_A

	and HASHER_B.
*/
type hashComposite struct {
	hasherCommon
	ha     hasherHandle
	hb     hasherHandle
	params *encoderParams
}

func (h *hashComposite) Initialize(params *encoderParams) {
	_ = "STUB: not implemented"

	/*
	   TODO: Initialize of the hashers is defered to Prepare (and params

	   	remembered here) because we don't get the one_shot and input_size params
	   	here that are needed to know the memory size of them. Instead provide
	   	those params to all hashers InitializehashComposite
	*/return
}

func (h *hashComposite) Prepare(one_shot bool, input_size uint, data []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *hashComposite) Store(data []byte, mask uint, ix uint) { _ = "STUB: not implemented"; return }

func (h *hashComposite) StoreRange(data []byte, mask uint, ix_start uint, ix_end uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashComposite) StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ring_buffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

func (h *hashComposite) PrepareDistanceCache(distance_cache []int) {
	_ = "STUB: not implemented"
	return
}

func (h *hashComposite) FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult) {
	_ = "STUB: not implemented"
	return
}
