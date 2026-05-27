package brotli

/* Copyright 2016 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/*
Dynamically grows array capacity to at least the requested size
T: data type
A: array
C: capacity
R: requested size
*/
func brotli_ensure_capacity_uint8_t(a *[]byte, c *uint, r uint) { _ = "STUB: not implemented"; return }

func brotli_ensure_capacity_uint32_t(a *[]uint32, c *uint, r uint) {
	_ = "STUB: not implemented"
	return
}
