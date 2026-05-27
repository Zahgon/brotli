package brotli

import (
	"sync"
)

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Function to find backward reference copies. */

func computeDistanceCode(distance uint, max_distance uint, dist_cache []int) uint {
	_ = "STUB: not implemented"
	return 0
}

var hasherSearchResultPool sync.Pool

func createBackwardReferences(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, hasher hasherHandle, dist_cache []int, last_insert_len *uint, commands *[]command, num_literals *uint) {
	_ = "STUB: not implemented"
	return
}

/* Set maximum distance, see section 9.1. of the spec. */

/* For speed up heuristics for random data. */

/* Minimum score to accept a backward reference. */

/* Found a match. Let's look for something even better ahead. */

/* Ok, let's just write one byte for now and start a match from the
   next byte. */

/* The first 16 codes are special short-codes,
   and the minimum offset is 1. */

/* Put the hash keys into the table, if there are enough bytes left.
   Depending on the hasher implementation, it can push all positions
   in the given range or only a subset of them.
   Avoid hash poisoning with RLE data. */

/* If we have not seen matches for a long time, we can skip some
   match lookups. Unsuccessful match lookups are very very expensive
   and this kind of a heuristic speeds up compression quite
   a lot. */

/* Going through uncompressible data, jump. */

/* It is quite a long time since we saw a copy, so we assume
   that this data is not compressible, and store hashes less
   often. Hashes of non compressible data are less likely to
   turn out to be useful in the future, too, so we store less of
   them to not to flood out the hash table of good compressible
   data. */
