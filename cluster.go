package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Functions for clustering similar histograms together. */

type histogramPair struct {
	idx1       uint32
	idx2       uint32
	cost_combo float64
	cost_diff  float64
}

func histogramPairIsLess(p1 *histogramPair, p2 *histogramPair) bool {
	_ = "STUB: not implemented"
	return false
}

/* Returns entropy reduction of the context map when we combine two clusters. */
func clusterCostDiff(size_a uint, size_b uint) float64 { _ = "STUB: not implemented"; return 0 }
