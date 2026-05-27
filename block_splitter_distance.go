package brotli

import "math"

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

func initialEntropyCodesDistance(data []uint16, length uint, stride uint, num_histograms uint, histograms []histogramDistance) {
	_ = "STUB: not implemented"
	return
}

func randomSampleDistance(seed *uint32, data []uint16, length uint, stride uint, sample *histogramDistance) {
	_ = "STUB: not implemented"
	return
}

func refineEntropyCodesDistance(data []uint16, length uint, stride uint, num_histograms uint, histograms []histogramDistance) {
	_ = "STUB: not implemented"
	return
}

/*
Assigns a block id from the range [0, num_histograms) to each data element

	in data[0..length) and fills in block_id[0..length) with the assigned values.
	Returns the number of blocks, i.e. one plus the number of block switches.
*/
func findBlocksDistance(data []uint16, length uint, block_switch_bitcost float64, num_histograms uint, histograms []histogramDistance, insert_cost []float64, cost []float64, switch_signal []byte, block_id []byte) uint {
	_ = "STUB: not implemented"
	return 0
}

/* After each iteration of this loop, cost[k] will contain the difference
   between the minimum cost of arriving at the current byte position using
   entropy code k, and the minimum cost of arriving at the current byte
   position. This difference is capped at the block switch cost, and if it
   reaches block switch cost, it means that when we trace back from the last
   position, we need to switch here. */

/* We are coding the symbol in data[byte_ix] with entropy code k. */

/* More blocks for the beginning. */

/* Trace back from the last position and switch at the marked places. */

var remapBlockIdsDistance_kInvalidId uint16 = 256

func remapBlockIdsDistance(block_ids []byte, length uint, new_id []uint16, num_histograms uint) uint {
	_ = "STUB: not implemented"
	return 0
}

func buildBlockHistogramsDistance(data []uint16, length uint, block_ids []byte, num_histograms uint, histograms []histogramDistance) {
	_ = "STUB: not implemented"
	return
}

var clusterBlocksDistance_kInvalidIndex uint32 = math.MaxUint32

func clusterBlocksDistance(data []uint16, length uint, num_blocks uint, block_ids []byte, split *blockSplit) {
	_ = "STUB: not implemented"
	return
}

func splitByteVectorDistance(data []uint16, length uint, literals_per_histogram uint, max_histograms uint, sampling_stride_length uint, block_switch_cost float64, params *encoderParams, split *blockSplit) {
	_ = "STUB: not implemented"
	return
}

/* Find good entropy codes. */

/* Find a good path through literals with the good entropy codes. */
