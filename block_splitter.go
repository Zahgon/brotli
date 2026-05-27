package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Block split point selection utilities. */

type blockSplit struct {
	num_types          uint
	num_blocks         uint
	types              []byte
	lengths            []uint32
	types_alloc_size   uint
	lengths_alloc_size uint
}

const (
	kMaxLiteralHistograms        uint    = 100
	kMaxCommandHistograms        uint    = 50
	kLiteralBlockSwitchCost      float64 = 28.1
	kCommandBlockSwitchCost      float64 = 13.5
	kDistanceBlockSwitchCost     float64 = 14.6
	kLiteralStrideLength         uint    = 70
	kCommandStrideLength         uint    = 40
	kSymbolsPerLiteralHistogram  uint    = 544
	kSymbolsPerCommandHistogram  uint    = 530
	kSymbolsPerDistanceHistogram uint    = 544
	kMinLengthForBlockSplitting  uint    = 128
	kIterMulForRefining          uint    = 2
	kMinItersForRefining         uint    = 100
)

func countLiterals(cmds []command) uint {
	_ = "STUB: not implemented"
	return

	/* Count how many we have. */
	0
}

func copyLiteralsToByteArray(cmds []command, data []byte, offset uint, mask uint, literals []byte) {
	_ = "STUB: not implemented"
	return
}

func myRand(seed *uint32) uint32 {
	_ = "STUB: not implemented"
	/* Initial seed should be 7. In this case, loop length is (1 << 29). */ return 0
}

func bitCost(count uint) float64 { _ = "STUB: not implemented"; return 0 }

const histogramsPerBatch = 64

const clustersPerBatch = 16

func initBlockSplit(self *blockSplit) { _ = "STUB: not implemented"; return }

func splitBlock(cmds []command, data []byte, pos uint, mask uint, params *encoderParams, literal_split *blockSplit, insert_and_copy_split *blockSplit, dist_split *blockSplit) {
	_ = "STUB: not implemented"
	return
}

/* Create a continuous array of literals. */

/* Create the block split on the array of literals.
   Literal histograms have alphabet size 256. */

/* Compute prefix codes for commands. */

/* Create the block split on the array of command prefixes. */

/* TODO: reuse for distances? */

/* Create a continuous array of distance prefixes. */

/* Create the block split on the array of distance prefixes. */
