package brotli

type hasherCommon struct {
	params           hasherParams
	is_prepared_     bool
	dict_num_lookups uint
	dict_num_matches uint
}

func (h *hasherCommon) Common() *hasherCommon { _ = "STUB: not implemented"; return nil }

type hasherHandle interface {
	Common() *hasherCommon
	Initialize(params *encoderParams)
	Prepare(one_shot bool, input_size uint, data []byte)
	StitchToPreviousBlock(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint)
	HashTypeLength() uint
	StoreLookahead() uint
	PrepareDistanceCache(distance_cache []int)
	FindLongestMatch(dictionary *encoderDictionary, data []byte, ring_buffer_mask uint, distance_cache []int, cur_ix uint, max_length uint, max_backward uint, gap uint, max_distance uint, out *hasherSearchResult)
	StoreRange(data []byte, mask uint, ix_start uint, ix_end uint)
	Store(data []byte, mask uint, ix uint)
}

const kCutoffTransformsCount uint32 = 10

/*   0,  12,   27,    23,    42,    63,    56,    48,    59,    64 */
/* 0+0, 4+8, 8+19, 12+11, 16+26, 20+43, 24+32, 28+20, 32+27, 36+28 */
const kCutoffTransforms uint64 = 0x071B520ADA2D3200

type hasherSearchResult struct {
	len            uint
	distance       uint
	score          uint
	len_code_delta int
}

/*
kHashMul32 multiplier has these properties:
  - The multiplier must be odd. Otherwise we may lose the highest bit.
  - No long streaks of ones or zeros.
  - There is no effort to ensure that it is a prime, the oddity is enough
    for this use.
  - The number has been tuned heuristically against compression benchmarks.
*/
const kHashMul32 uint32 = 0x1E35A7BD

const kHashMul64 uint64 = 0x1E35A7BD1E35A7BD

const kHashMul64Long uint64 = 0x1FE35A7BD3579BD3

func hash14(data []byte) uint32 { _ = "STUB: not implemented"; return 0 }

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

func prepareDistanceCache(distance_cache []int, num_distances int) {
	_ = "STUB: not implemented"
	return
}

const literalByteScore = 135

const distanceBitPenalty = 30

/* Score must be positive after applying maximal penalty. */
const scoreBase = (distanceBitPenalty * 8 * 8)

/*
Usually, we always choose the longest backward reference. This function

	allows for the exception of that rule.

	If we choose a backward reference that is further away, it will
	usually be coded with more bits. We approximate this by assuming
	log2(distance). If the distance can be expressed in terms of the
	last four distances, we use some heuristic constants to estimate
	the bits cost. For the first up to four literals we use the bit
	cost of the literals from the literal cost model, after that we
	use the average bit cost of the cost model.

	This function is used to sometimes discard a longer backward reference
	when it is not much longer and the bit cost for encoding it is more
	than the saved literals.

	backward_reference_offset MUST be positive.
*/
func backwardReferenceScore(copy_length uint, backward_reference_offset uint) uint {
	_ = "STUB: not implemented"
	return 0
}

func backwardReferenceScoreUsingLastDistance(copy_length uint) uint {
	_ = "STUB: not implemented"
	return 0
}

func backwardReferencePenaltyUsingLastDistance(distance_short_code uint) uint {
	_ = "STUB: not implemented"
	return 0
}

func testStaticDictionaryItem(dictionary *encoderDictionary, item uint, data []byte, max_length uint, max_backward uint, max_distance uint, out *hasherSearchResult) bool {
	_ = "STUB: not implemented"
	return false
}

func searchInStaticDictionary(dictionary *encoderDictionary, handle hasherHandle, data []byte, max_length uint, max_backward uint, max_distance uint, out *hasherSearchResult, shallow bool) {
	_ = "STUB: not implemented"
	return
}

type backwardMatch struct {
	distance        uint32
	length_and_code uint32
}

func initBackwardMatch(self *backwardMatch, dist uint, len uint) { _ = "STUB: not implemented"; return }

func initDictionaryBackwardMatch(self *backwardMatch, dist uint, len uint, len_code uint) {
	_ = "STUB: not implemented"
	return
}

func backwardMatchLength(self *backwardMatch) uint { _ = "STUB: not implemented"; return 0 }

func backwardMatchLengthCode(self *backwardMatch) uint { _ = "STUB: not implemented"; return 0 }

func hasherReset(handle hasherHandle) { _ = "STUB: not implemented"; return }

func newHasher(typ int) hasherHandle { _ = "STUB: not implemented"; return *new(hasherHandle) }

func hasherSetup(handle *hasherHandle, params *encoderParams, data []byte, position uint, input_size uint, is_last bool) {
	_ = "STUB: not implemented"
	return
}

func initOrStitchToPreviousBlock(handle *hasherHandle, data []byte, mask uint, params *encoderParams, position uint, input_size uint, is_last bool) {
	_ = "STUB: not implemented"
	return
}
