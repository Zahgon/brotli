package brotli

const fastOnePassCompressionQuality = 0

const fastTwoPassCompressionQuality = 1

const zopflificationQuality = 10

const hqZopflificationQuality = 11

const maxQualityForStaticEntropyCodes = 2

const minQualityForBlockSplit = 4

const minQualityForNonzeroDistanceParams = 4

const minQualityForOptimizeHistograms = 4

const minQualityForExtensiveReferenceSearch = 5

const minQualityForContextModeling = 5

const minQualityForHqContextModeling = 7

const minQualityForHqBlockSplitting = 10

/*
For quality below MIN_QUALITY_FOR_BLOCK_SPLIT there is no block splitting,

	so we buffer at most this much literals and commands.
*/
const maxNumDelayedSymbols = 0x2FFF

/* Returns hash-table size for quality levels 0 and 1. */
func maxHashTableSize(quality int) uint { _ = "STUB: not implemented"; return 0 }

/* The maximum length for which the zopflification uses distinct distances. */
const maxZopfliLenQuality10 = 150

const maxZopfliLenQuality11 = 325

/* Do not thoroughly search when a long copy is found. */
const longCopyQuickStep = 16384

func maxZopfliLen(params *encoderParams) uint { _ = "STUB: not implemented"; return 0 }

/* Number of best candidates to evaluate to expand Zopfli chain. */
func maxZopfliCandidates(params *encoderParams) uint { _ = "STUB: not implemented"; return 0 }

func sanitizeParams(params *encoderParams) { _ = "STUB: not implemented"; return }

/* Returns optimized lg_block value. */
func computeLgBlock(params *encoderParams) int { _ = "STUB: not implemented"; return 0 }

/*
Returns log2 of the size of main ring buffer area.

	Allocate at least lgwin + 1 bits for the ring buffer so that the newly
	added block fits there completely and we still get lgwin bits and at least
	read_block_size_bits + 1 bits because the copy tail length needs to be
	smaller than ring-buffer size.
*/
func computeRbBits(params *encoderParams) int { _ = "STUB: not implemented"; return 0 }

func maxMetablockSize(params *encoderParams) uint { _ = "STUB: not implemented"; return 0 }

/*
When searching for backward references and have not seen matches for a long

	time, we can skip some match lookups. Unsuccessful match lookups are very
	expensive and this kind of a heuristic speeds up compression quite a lot.
	At first 8 byte strides are taken and every second byte is put to hasher.
	After 4x more literals stride by 16 bytes, every put 4-th byte to hasher.
	Applied only to qualities 2 to 9.
*/
func literalSpreeLengthForSparseSearch(params *encoderParams) uint {
	_ = "STUB: not implemented"
	return 0
}

func chooseHasher(params *encoderParams, hparams *hasherParams) { _ = "STUB: not implemented"; return }

/* Different hashers for large window brotli: not for qualities <= 2,
   these are too fast for large window. Not for qualities >= 10: their
   hasher already works well with large window. So the changes are:
   H3 --> H35: for quality 3.
   H54 --> H55: for quality 4 with size hint > 1MB
   H6 --> H65: for qualities 5, 6, 7, 8, 9. */
