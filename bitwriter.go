package brotli

/* Copyright 2010 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Write bits into a byte array. */

type bitWriter struct {
	dst []byte

	// Data waiting to be written is the low nbits of bits.
	bits  uint64
	nbits uint
}

func (w *bitWriter) writeBits(nb uint, b uint64) { _ = "STUB: not implemented"; return }

func (w *bitWriter) writeSingleBit(bit bool) { _ = "STUB: not implemented"; return }

func (w *bitWriter) jumpToByteBoundary() { _ = "STUB: not implemented"; return }

// Avoid underflow
