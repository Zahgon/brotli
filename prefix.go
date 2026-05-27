package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Functions for encoding of integers into prefix codes the amount of extra
   bits, and the actual values of the extra bits. */

/*
Here distance_code is an intermediate code, i.e. one of the special codes or

	the actual distance increased by BROTLI_NUM_DISTANCE_SHORT_CODES - 1.
*/
func prefixEncodeCopyDistance(distance_code uint, num_direct_codes uint, postfix_bits uint, code *uint16, extra_bits *uint32) {
	_ = "STUB: not implemented"
	return
}
