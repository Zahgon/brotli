package brotli

var kInsBase = []uint32{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	8,
	10,
	14,
	18,
	26,
	34,
	50,
	66,
	98,
	130,
	194,
	322,
	578,
	1090,
	2114,
	6210,
	22594,
}

var kInsExtra = []uint32{
	0,
	0,
	0,
	0,
	0,
	0,
	1,
	1,
	2,
	2,
	3,
	3,
	4,
	4,
	5,
	5,
	6,
	7,
	8,
	9,
	10,
	12,
	14,
	24,
}

var kCopyBase = []uint32{
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	12,
	14,
	18,
	22,
	30,
	38,
	54,
	70,
	102,
	134,
	198,
	326,
	582,
	1094,
	2118,
}

var kCopyExtra = []uint32{
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	1,
	1,
	2,
	2,
	3,
	3,
	4,
	4,
	5,
	5,
	6,
	7,
	8,
	9,
	10,
	24,
}

func getInsertLengthCode(insertlen uint) uint16 { _ = "STUB: not implemented"; return 0 }

func getCopyLengthCode(copylen uint) uint16 { _ = "STUB: not implemented"; return 0 }

func combineLengthCodes(inscode uint16, copycode uint16, use_last_distance bool) uint16 {
	_ = "STUB: not implemented"
	return 0
}

/* Specification: 5 Encoding of ... (last table) */
/* offset = 2 * index, where index is in range [0..8] */

/* All values in specification are K * 64,
   where   K = [2, 3, 6, 4, 5, 8, 7, 9, 10],
       i + 1 = [1, 2, 3, 4, 5, 6, 7, 8,  9],
   K - i - 1 = [1, 1, 3, 0, 0, 2, 0, 1,  2] = D.
   All values in D require only 2 bits to encode.
   Magic constant is shifted 6 bits left, to avoid final multiplication. */

func getLengthCode(insertlen uint, copylen uint, use_last_distance bool, code *uint16) {
	_ = "STUB: not implemented"
	return
}

func getInsertBase(inscode uint16) uint32 { _ = "STUB: not implemented"; return 0 }

func getInsertExtra(inscode uint16) uint32 { _ = "STUB: not implemented"; return 0 }

func getCopyBase(copycode uint16) uint32 { _ = "STUB: not implemented"; return 0 }

func getCopyExtra(copycode uint16) uint32 { _ = "STUB: not implemented"; return 0 }

type command struct {
	insert_len_  uint32
	copy_len_    uint32
	dist_extra_  uint32
	cmd_prefix_  uint16
	dist_prefix_ uint16
}

/* distance_code is e.g. 0 for same-as-last short code, or 16 for offset 1. */
func makeCommand(dist *distanceParams, insertlen uint, copylen uint, copylen_code_delta int, distance_code uint) (cmd command) {
	_ = "STUB: not implemented"
	/* Don't rely on signed int representation, use honest casts. */ return *new(command)
}

/* The distance prefix and extra bits are stored in this Command as if
   npostfix and ndirect were 0, they are only recomputed later after the
   clustering if needed. */

func makeInsertCommand(insertlen uint) (cmd command) {
	_ = "STUB: not implemented"
	return *new(command)
}

func commandRestoreDistanceCode(self *command, dist *distanceParams) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func commandDistanceContext(self *command) uint32 { _ = "STUB: not implemented"; return 0 }

func commandCopyLen(self *command) uint32 { _ = "STUB: not implemented"; return 0 }

func commandCopyLenCode(self *command) uint32 { _ = "STUB: not implemented"; return 0 }
