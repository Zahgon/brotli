package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Class to model the static dictionary. */

const maxStaticDictionaryMatchLen = 37

const kInvalidMatch uint32 = 0xFFFFFFF

/*
Copyright 2013 Google Inc. All Rights Reserved.

	Distributed under MIT license.
	See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/
func hash(data []byte) uint32 { _ = "STUB: not implemented"; return 0 }

/* The higher bits contain more mixture from the multiplication,
   so we take our results from there. */

func addMatch(distance uint, len uint, len_code uint, matches []uint32) {
	_ = "STUB: not implemented"
	return
}

func dictMatchLength(dict *dictionary, data []byte, id uint, len uint, maxlen uint) uint {
	_ = "STUB: not implemented"
	return 0
}

func isMatch(d *dictionary, w dictWord, data []byte, max_length uint) bool {
	_ = "STUB: not implemented"
	return false
}

/* Match against base dictionary word. */

/* Match against uppercase first transform.
   Note that there are only ASCII uppercase words in the lookup table. */

/* Match against uppercase all transform.
   Note that there are only ASCII uppercase words in the lookup table. */

func findAllStaticDictionaryMatches(dict *encoderDictionary, data []byte, min_length uint, max_length uint, matches []uint32) bool {
	_ = "STUB: not implemented"
	return false
}

/* Transform "" + BROTLI_TRANSFORM_IDENTITY + "" */

/* Transforms "" + BROTLI_TRANSFORM_OMIT_LAST_1 + "" and
   "" + BROTLI_TRANSFORM_OMIT_LAST_1 + "ing " */

/* Transform "" + BROTLI_TRANSFORM_OMIT_LAST_# + "" (# = 2 .. 9) */

/* Transforms "" + BROTLI_TRANSFORM_IDENTITY + <suffix> */

/* Set is_all_caps=0 for BROTLI_TRANSFORM_UPPERCASE_FIRST and
    is_all_caps=1 otherwise (BROTLI_TRANSFORM_UPPERCASE_ALL)
transform. */

/* Transform "" + kUppercase{First,All} + "" */

/* Transforms "" + kUppercase{First,All} + <suffix> */

/* Transforms with prefixes " " and "." */

/* Transforms " " + BROTLI_TRANSFORM_IDENTITY + "" and
   "." + BROTLI_TRANSFORM_IDENTITY + "" */

/* Transforms " " + BROTLI_TRANSFORM_IDENTITY + <suffix> and
   "." + BROTLI_TRANSFORM_IDENTITY + <suffix>
*/

/* Set is_all_caps=0 for BROTLI_TRANSFORM_UPPERCASE_FIRST and
    is_all_caps=1 otherwise (BROTLI_TRANSFORM_UPPERCASE_ALL)
transform. */

/* Transforms " " + kUppercase{First,All} + "" */

/* Transforms " " + kUppercase{First,All} + <suffix> */

/* Transforms with prefixes "e ", "s ", ", " and "\xC2\xA0" */

/* Transforms with prefixes " the " and ".com/" */
