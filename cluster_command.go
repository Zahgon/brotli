package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/*
Computes the bit cost reduction by combining out[idx1] and out[idx2] and if

	it is below a threshold, stores the pair (idx1, idx2) in the *pairs queue.
*/
func compareAndPushToQueueCommand(out []histogramCommand, cluster_size []uint32, idx1 uint32, idx2 uint32, max_num_pairs uint, pairs []histogramPair, num_pairs *uint) {
	_ = "STUB: not implemented"
	return
}

/* Replace the top of the queue if needed. */

func histogramCombineCommand(out []histogramCommand, cluster_size []uint32, symbols []uint32, clusters []uint32, pairs []histogramPair, num_clusters uint, symbols_size uint, max_clusters uint, max_num_pairs uint) uint {
	_ = "STUB: not implemented"
	return 0
}

/* We maintain a vector of histogram pairs, with the property that the pair
   with the maximum bit cost reduction is the first. */

/* Take the best pair from the top of heap. */

/* Remove pairs intersecting the just combined best pair. */

/* Remove invalid pair from the queue. */

/* Replace the top of the queue if needed. */

/* Push new pairs formed with the combined histogram to the heap. */

/* What is the bit cost of moving histogram from cur_symbol to candidate. */
func histogramBitCostDistanceCommand(histogram *histogramCommand, candidate *histogramCommand) float64 {
	_ = "STUB: not implemented"
	return 0
}
