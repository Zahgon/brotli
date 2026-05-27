package brotli

import "math"

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/*
Computes the bit cost reduction by combining out[idx1] and out[idx2] and if

	it is below a threshold, stores the pair (idx1, idx2) in the *pairs queue.
*/
func compareAndPushToQueueDistance(out []histogramDistance, cluster_size []uint32, idx1 uint32, idx2 uint32, max_num_pairs uint, pairs []histogramPair, num_pairs *uint) {
	_ = "STUB: not implemented"
	return
}

/* Replace the top of the queue if needed. */

func histogramCombineDistance(out []histogramDistance, cluster_size []uint32, symbols []uint32, clusters []uint32, pairs []histogramPair, num_clusters uint, symbols_size uint, max_clusters uint, max_num_pairs uint) uint {
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
func histogramBitCostDistanceDistance(histogram *histogramDistance, candidate *histogramDistance) float64 {
	_ = "STUB: not implemented"
	return 0
}

/*
Find the best 'out' histogram for each of the 'in' histograms.

	When called, clusters[0..num_clusters) contains the unique values from
	symbols[0..in_size), but this property is not preserved in this function.
	Note: we assume that out[]->bit_cost_ is already up-to-date.
*/
func histogramRemapDistance(in []histogramDistance, in_size uint, clusters []uint32, num_clusters uint, out []histogramDistance, symbols []uint32) {
	_ = "STUB: not implemented"
	return
}

/* Recompute each out based on raw and symbols. */

/* Reorders elements of the out[0..length) array and changes values in
   symbols[0..length) array in the following way:
     * when called, symbols[] contains indexes into out[], and has N unique
       values (possibly N < length)
     * on return, symbols'[i] = f(symbols[i]) and
                  out'[symbols'[i]] = out[symbols[i]], for each 0 <= i < length,
       where f is a bijection between the range of symbols[] and [0..N), and
       the first occurrences of values in symbols'[i] come in consecutive
       increasing order.
   Returns N, the number of unique values in symbols[]. */

var histogramReindexDistance_kInvalidIndex uint32 = math.MaxUint32

func histogramReindexDistance(out []histogramDistance, symbols []uint32, length uint) uint {
	_ = "STUB: not implemented"
	return 0
}

/* TODO: by using idea of "cycle-sort" we can avoid allocation of
   tmp and reduce the number of copying by the factor of 2. */

func clusterHistogramsDistance(in []histogramDistance, in_size uint, max_histograms uint, out []histogramDistance, out_size *uint, histogram_symbols []uint32) {
	_ = "STUB: not implemented"
	return
}

/* For the first pass of clustering, we allow all pairs. */

/* For the second pass, we limit the total number of histogram pairs.
   After this limit is reached, we only keep searching for the best pair. */

/* Collapse similar histograms. */

/* Find the optimal map from original histograms to the final ones. */

/* Convert the context map to a canonical form. */
