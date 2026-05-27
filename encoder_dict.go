package brotli

/* Dictionary data (words and transforms) for 1 possible context */
type encoderDictionary struct {
	words                 *dictionary
	cutoffTransformsCount uint32
	cutoffTransforms      uint64
	hash_table            []uint16
	buckets               []uint16
	dict_words            []dictWord
}

func initEncoderDictionary(dict *encoderDictionary) { _ = "STUB: not implemented"; return }
