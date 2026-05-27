package brotli

/* Copyright 2015 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Greedy block splitter for one block category (literal, command or distance).
 */
type blockSplitterDistance struct {
	alphabet_size_     uint
	min_block_size_    uint
	split_threshold_   float64
	num_blocks_        uint
	split_             *blockSplit
	histograms_        []histogramDistance
	histograms_size_   *uint
	target_block_size_ uint
	block_size_        uint
	curr_histogram_ix_ uint
	last_histogram_ix_ [2]uint
	last_entropy_      [2]float64
	merge_last_count_  uint
}

func initBlockSplitterDistance(self *blockSplitterDistance, alphabet_size uint, min_block_size uint, split_threshold float64, num_symbols uint, split *blockSplit, histograms *[]histogramDistance, histograms_size *uint) {
	_ = "STUB: not implemented"
	return
}

/* We have to allocate one more histogram than the maximum number of block
   types for the current histogram when the meta-block is too big. */

/* Clear only current histogram. */

/*
Does either of three things:

	(1) emits the current block with a new block type;
	(2) emits the current block with the type of the second last block;
	(3) merges the current block with the last block.
*/
func blockSplitterFinishBlockDistance(self *blockSplitterDistance, is_final bool) {
	_ = "STUB: not implemented"
	return
}

/* Create first block. */

/* Create new block. */

/* Combine this block with second last block. */

/* Combine this block with last block. */

/*
Adds the next symbol to the current histogram. When the current histogram

	reaches the target size, decides on merging the block.
*/
func blockSplitterAddSymbolDistance(self *blockSplitterDistance, symbol uint) {
	_ = "STUB: not implemented"
	return
}

/* is_final = */
