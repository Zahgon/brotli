package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Functions to estimate the bit cost of Huffman trees. */
func shannonEntropy(population []uint32, size uint, total *uint) float64 {
	_ = "STUB: not implemented"
	return 0
}

func bitsEntropy(population []uint32, size uint) float64 { _ = "STUB: not implemented"; return 0 }

/* At least one bit per literal is needed. */

const kOneSymbolHistogramCost float64 = 12
const kTwoSymbolHistogramCost float64 = 20
const kThreeSymbolHistogramCost float64 = 28
const kFourSymbolHistogramCost float64 = 37

func populationCostLiteral(histogram *histogramLiteral) float64 {
	_ = "STUB: not implemented"
	return 0
}

/* Sort */

/* In this loop we compute the entropy of the histogram and simultaneously
   build a simplified histogram of the code length codes where we use the
   zero repeat code 17, but we don't use the non-zero repeat code 16. */

/* Compute -log2(P(symbol)) = -log2(count(symbol)/total_count) =
   = log2(total_count) - log2(count(symbol)) */

/* Approximate the bit depth by round(-log2(P(symbol))) */

/* Compute the run length of zeros and add the appropriate number of 0
   and 17 code length codes to the code length code histogram. */

/* Don't add any cost for the last zero run, since these are encoded
   only implicitly. */

/* Add the 3 extra bits for the 17 code length code. */

/* Add the estimated encoding cost of the code length code histogram. */

/* Add the entropy of the code length code histogram. */

func populationCostCommand(histogram *histogramCommand) float64 {
	_ = "STUB: not implemented"
	return 0
}

/* Sort */

/* In this loop we compute the entropy of the histogram and simultaneously
   build a simplified histogram of the code length codes where we use the
   zero repeat code 17, but we don't use the non-zero repeat code 16. */

/* Compute -log2(P(symbol)) = -log2(count(symbol)/total_count) =
   = log2(total_count) - log2(count(symbol)) */

/* Approximate the bit depth by round(-log2(P(symbol))) */

/* Compute the run length of zeros and add the appropriate number of 0
   and 17 code length codes to the code length code histogram. */

/* Don't add any cost for the last zero run, since these are encoded
   only implicitly. */

/* Add the 3 extra bits for the 17 code length code. */

/* Add the estimated encoding cost of the code length code histogram. */

/* Add the entropy of the code length code histogram. */

func populationCostDistance(histogram *histogramDistance) float64 {
	_ = "STUB: not implemented"
	return 0
}

/* Sort */

/* In this loop we compute the entropy of the histogram and simultaneously
   build a simplified histogram of the code length codes where we use the
   zero repeat code 17, but we don't use the non-zero repeat code 16. */

/* Compute -log2(P(symbol)) = -log2(count(symbol)/total_count) =
   = log2(total_count) - log2(count(symbol)) */

/* Approximate the bit depth by round(-log2(P(symbol))) */

/* Compute the run length of zeros and add the appropriate number of 0
   and 17 code length codes to the code length code histogram. */

/* Don't add any cost for the last zero run, since these are encoded
   only implicitly. */

/* Add the 3 extra bits for the 17 code length code. */

/* Add the estimated encoding cost of the code length code histogram. */

/* Add the entropy of the code length code histogram. */
