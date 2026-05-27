package brotli

/* Copyright 2010 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Entropy encoding (Huffman) utilities. */

/* A node of a Huffman tree. */
type huffmanTree struct {
	total_count_          uint32
	index_left_           int16
	index_right_or_value_ int16
}

func initHuffmanTree(self *huffmanTree, count uint32, left int16, right int16) {
	_ = "STUB: not implemented"
	return
}

/* Input size optimized Shell sort. */
type huffmanTreeComparator func(huffmanTree, huffmanTree) bool

var sortHuffmanTreeItems_gaps = []uint{132, 57, 23, 10, 4, 1}

func sortHuffmanTreeItems(items []huffmanTree, n uint, comparator huffmanTreeComparator) {
	_ = "STUB: not implemented"

	/* Insertion sort. */
	return
}

/* Returns 1 if assignment of depths succeeded, otherwise 0. */
func setDepth(p0 int, pool []huffmanTree, depth []byte, max_depth int) bool {
	_ = "STUB: not implemented"
	return false
}

/* Sort the root nodes, least popular first. */
func sortHuffmanTree(v0 huffmanTree, v1 huffmanTree) bool { _ = "STUB: not implemented"; return false }

/*
This function will create a Huffman tree.

	The catch here is that the tree cannot be arbitrarily deep.
	Brotli specifies a maximum depth of 15 bits for "code trees"
	and 7 bits for "code length code trees."

	count_limit is the value that is to be faked as the minimum value
	and this minimum value is raised until the tree matches the
	maximum length requirement.

	This algorithm is not of excellent performance for very long data blocks,
	especially when population counts are longer than 2**tree_limit, but
	we are not planning to use this with extremely long blocks.

	See http://en.wikipedia.org/wiki/Huffman_coding
*/
func createHuffmanTree(data []uint32, length uint, tree_limit int, tree []huffmanTree, depth []byte) {
	_ = "STUB: not implemented"
	return
}

/* For block sizes below 64 kB, we never need to do a second iteration
   of this loop. Probably all of our block sizes will be smaller than
   that, so this loop is mostly of academic interest. If we actually
   would need this, we would be better off with the Katajainen algorithm. */

/* Only one element. */

/* The nodes are:
   [0, n): the sorted leaf nodes that we start with.
   [n]: we add a sentinel here.
   [n + 1, 2n): new parent nodes are added here, starting from
                (n+1). These are naturally in ascending order.
   [2n]: we add a sentinel at the end as well.
   There will be (2n+1) elements at the end. */

/* Points to the next leaf node. */
/* Points to the next non-leaf node. */

/* The sentinel node becomes the parent node. */

/* Add back the last sentinel node. */

/* We need to pack the Huffman tree in tree_limit bits. If this was not
   successful, add fake entities to the lowest values and retry. */

func reverse(v []byte, start uint, end uint) { _ = "STUB: not implemented"; return }

func writeHuffmanTreeRepetitions(previous_value byte, value byte, repetitions uint, tree_size *uint, tree []byte, extra_bits_data []byte) {
	_ = "STUB: not implemented"
	return
}

func writeHuffmanTreeRepetitionsZeros(repetitions uint, tree_size *uint, tree []byte, extra_bits_data []byte) {
	_ = "STUB: not implemented"
	return
}

/*
Change the population counts in a way that the consequent

	Huffman tree compression, especially its RLE-part will be more
	likely to compress this data more efficiently.

	length contains the size of the histogram.
	counts contains the population counts.
	good_for_rle is a buffer of at least length size
*/
func optimizeHuffmanCountsForRLE(length uint, counts []uint32, good_for_rle []byte) {
	_ = "STUB: not implemented"
	return
}

/* Let's make the Huffman code more compatible with RLE encoding. */

/* All zeros. */

/* Now counts[0..length - 1] does not have trailing zeros. */

/* Small histogram will model it well. */

/* 2) Let's mark all population counts that already can be encoded
   with an RLE code. */

/* Let's not spoil any of the existing good RLE codes.
   Mark any seq of 0's that is longer as 5 as a good_for_rle.
   Mark any seq of non-0's that is longer as 7 as a good_for_rle. */

/* 3) Let's replace those population counts that lead to more RLE codes.
   Math here is in 24.8 fixed point representation. */

/* The stride must end, collapse what we have, if we have enough (4). */

/* Don't make an all zeros stride to be upgraded to ones. */

/* We don't want to change value at counts[i],
   that is already belonging to the next stride. Thus - 1. */

/* All interesting strides have a count of at least 4, */
/* at least when non-zeros. */

func decideOverRLEUse(depth []byte, length uint, use_rle_for_non_zero *bool, use_rle_for_zero *bool) {
	_ = "STUB: not implemented"
	return
}

/*
Write a Huffman tree from bit depths into the bit-stream representation

	of a Huffman tree. The generated Huffman tree is to be compressed once
	more using a Huffman tree
*/
func writeHuffmanTree(depth []byte, length uint, tree_size *uint, tree []byte, extra_bits_data []byte) {
	_ = "STUB: not implemented"
	return
}

/* Throw away trailing zeros. */

/* First gather statistics on if it is a good idea to do RLE. */

/* Find RLE coding for longer codes.
   Shorter codes seem not to benefit from RLE. */

/* Actual RLE coding. */

var reverseBits_kLut = [16]uint{
	0x00,
	0x08,
	0x04,
	0x0C,
	0x02,
	0x0A,
	0x06,
	0x0E,
	0x01,
	0x09,
	0x05,
	0x0D,
	0x03,
	0x0B,
	0x07,
	0x0F,
}

func reverseBits(num_bits uint, bits uint16) uint16 { _ = "STUB: not implemented"; return 0 }

/* 0..15 are values for bits */
const maxHuffmanBits = 16

/* Get the actual bit values for a tree of bit depths. */
func convertBitDepthsToSymbols(depth []byte, len uint, bits []uint16) {
	_ = "STUB: not implemented"
	return
}

/* In Brotli, all bit depths are [1..15]
   0 bit depth means that the symbol does not exist. */
