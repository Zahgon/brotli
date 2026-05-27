package brotli

/* Copyright 2013 Google Inc. All Rights Reserved.

   Distributed under MIT license.
   See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
*/

/* Utilities for building Huffman decoding tables. */

type symbolList struct {
	storage []uint16
	offset  int
}

func symbolListGet(sl symbolList, i int) uint16 { _ = "STUB: not implemented"; return 0 }

func symbolListPut(sl symbolList, i int, val uint16) { _ = "STUB: not implemented"; return }
