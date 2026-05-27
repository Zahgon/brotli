package brotli

/* The distance symbols effectively used by "Large Window Brotli" (32-bit). */
const numHistogramDistanceSymbols = 544

type histogramLiteral struct {
	data_        [numLiteralSymbols]uint32
	total_count_ uint
	bit_cost_    float64
}

func histogramClearLiteral(self *histogramLiteral) { _ = "STUB: not implemented"; return }

func clearHistogramsLiteral(array []histogramLiteral, length uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddLiteral(self *histogramLiteral, val uint) { _ = "STUB: not implemented"; return }

func histogramAddVectorLiteral(self *histogramLiteral, p []byte, n uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddHistogramLiteral(self *histogramLiteral, v *histogramLiteral) {
	_ = "STUB: not implemented"
	return
}

func histogramDataSizeLiteral() uint { _ = "STUB: not implemented"; return 0 }

type histogramCommand struct {
	data_        [numCommandSymbols]uint32
	total_count_ uint
	bit_cost_    float64
}

func histogramClearCommand(self *histogramCommand) { _ = "STUB: not implemented"; return }

func clearHistogramsCommand(array []histogramCommand, length uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddCommand(self *histogramCommand, val uint) { _ = "STUB: not implemented"; return }

func histogramAddVectorCommand(self *histogramCommand, p []uint16, n uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddHistogramCommand(self *histogramCommand, v *histogramCommand) {
	_ = "STUB: not implemented"
	return
}

func histogramDataSizeCommand() uint { _ = "STUB: not implemented"; return 0 }

type histogramDistance struct {
	data_        [numDistanceSymbols]uint32
	total_count_ uint
	bit_cost_    float64
}

func histogramClearDistance(self *histogramDistance) { _ = "STUB: not implemented"; return }

func clearHistogramsDistance(array []histogramDistance, length uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddDistance(self *histogramDistance, val uint) { _ = "STUB: not implemented"; return }

func histogramAddVectorDistance(self *histogramDistance, p []uint16, n uint) {
	_ = "STUB: not implemented"
	return
}

func histogramAddHistogramDistance(self *histogramDistance, v *histogramDistance) {
	_ = "STUB: not implemented"
	return
}

func histogramDataSizeDistance() uint { _ = "STUB: not implemented"; return 0 }

type blockSplitIterator struct {
	split_  *blockSplit
	idx_    uint
	type_   uint
	length_ uint
}

func initBlockSplitIterator(self *blockSplitIterator, split *blockSplit) {
	_ = "STUB: not implemented"
	return
}

func blockSplitIteratorNext(self *blockSplitIterator) { _ = "STUB: not implemented"; return }

func buildHistogramsWithContext(cmds []command, literal_split *blockSplit, insert_and_copy_split *blockSplit, dist_split *blockSplit, ringbuffer []byte, start_pos uint, mask uint, prev_byte byte, prev_byte2 byte, context_modes []int, literal_histograms []histogramLiteral, insert_and_copy_histograms []histogramCommand, copy_dist_histograms []histogramDistance) {
	_ = "STUB: not implemented"
	return
}

/* TODO: unwrap iterator blocks. */
