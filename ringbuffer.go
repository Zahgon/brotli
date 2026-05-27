package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/*
A ringBuffer(window_bits, tail_bits) contains `1 << window_bits' bytes of

	data in a circular manner: writing a byte writes it to:
	  `position() % (1 << window_bits)'.
	For convenience, the ringBuffer array contains another copy of the
	first `1 << tail_bits' bytes:
	  buffer_[i] == buffer_[i + (1 << window_bits)], if i < (1 << tail_bits),
	and another copy of the last two bytes:
	  buffer_[-1] == buffer_[(1 << window_bits) - 1] and
	  buffer_[-2] == buffer_[(1 << window_bits) - 2].
*/
type ringBuffer struct {
	size_       uint32
	mask_       uint32
	tail_size_  uint32
	total_size_ uint32
	cur_size_   uint32
	pos_        uint32
	data_       []byte
	buffer_     []byte
}

func ringBufferInit(rb *ringBuffer) { _ = "STUB: not implemented"; return }

func ringBufferSetup(params *encoderParams, rb *ringBuffer) { _ = "STUB: not implemented"; return }

const kSlackForEightByteHashingEverywhere uint = 7

/*
Allocates or re-allocates data_ to the given length + plus some slack

	region before and after. Fills the slack regions with zeros.
*/
func ringBufferInitBuffer(buflen uint32, rb *ringBuffer) { _ = "STUB: not implemented"; return }

func ringBufferWriteTail(bytes []byte, n uint, rb *ringBuffer) { _ = "STUB: not implemented"; return }

/* Just fill the tail buffer with the beginning data. */

/* Push bytes into the ring buffer. */
func ringBufferWrite(bytes []byte, n uint, rb *ringBuffer) { _ = "STUB: not implemented"; return }

/* Special case for the first write: to process the first block, we don't
   need to allocate the whole ring-buffer and we don't need the tail
   either. However, we do this memory usage optimization only if the
   first write is less than the tail size, which is also the input block
   size, otherwise it is likely that other blocks will follow and we
   will need to reallocate to the full size anyway. */

/* Lazily allocate the full buffer. */

/* Initialize the last two bytes to zero, so that we don't have to worry
   later when we copy the last two bytes to the first two positions. */

/* The length of the writes is limited so that we do not need to worry
   about a write */

/* A single write fits. */

/* Split into two writes.
   Copy into the end of the buffer, including the tail buffer. */

/* Copy into the beginning of the buffer */

/* Wrap, but preserve not-a-first-lap feature. */
