package brotli

type zopfliNode struct {
	length              uint32
	distance            uint32
	dcode_insert_length uint32
	u                   struct {
		cost     float32
		next     uint32
		shortcut uint32
	}
}

const maxEffectiveDistanceAlphabetSize = 544

const kInfinity float32 = 1.7e38 /* ~= 2 ^ 127 */

var kDistanceCacheIndex = []uint32{0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1}

var kDistanceCacheOffset = []int{0, 0, 0, 0, -1, 1, -2, 2, -3, 3, -1, 1, -2, 2, -3, 3}

func initZopfliNodes(array []zopfliNode, length uint) { _ = "STUB: not implemented"; return }

func zopfliNodeCopyLength(self *zopfliNode) uint32 { _ = "STUB: not implemented"; return 0 }

func zopfliNodeLengthCode(self *zopfliNode) uint32 { _ = "STUB: not implemented"; return 0 }

func zopfliNodeCopyDistance(self *zopfliNode) uint32 { _ = "STUB: not implemented"; return 0 }

func zopfliNodeDistanceCode(self *zopfliNode) uint32 { _ = "STUB: not implemented"; return 0 }

func zopfliNodeCommandLength(self *zopfliNode) uint32 { _ = "STUB: not implemented"; return 0 }

/* Histogram based cost model for zopflification. */
type zopfliCostModel struct {
	cost_cmd_               [numCommandSymbols]float32
	cost_dist_              []float32
	distance_histogram_size uint32
	literal_costs_          []float32
	min_cost_cmd_           float32
	num_bytes_              uint
}

func initZopfliCostModel(self *zopfliCostModel, dist *distanceParams, num_bytes uint) {
	_ = "STUB: not implemented"
	return
}

func cleanupZopfliCostModel(self *zopfliCostModel) { _ = "STUB: not implemented"; return }

func setCost(histogram []uint32, histogram_size uint, literal_histogram bool, cost []float32) {
	_ = "STUB: not implemented"
	return
}

/* Shannon bits for this symbol. */

/* Cannot be coded with less than 1 bit */

func zopfliCostModelSetFromCommands(self *zopfliCostModel, position uint, ringbuffer []byte, ringbuffer_mask uint, commands []command, last_insert_len uint) {
	_ = "STUB: not implemented"
	return
}

func zopfliCostModelSetFromLiteralCosts(self *zopfliCostModel, position uint, ringbuffer []byte, ringbuffer_mask uint) {
	_ = "STUB: not implemented"
	return
}

func zopfliCostModelGetCommandCost(self *zopfliCostModel, cmdcode uint16) float32 {
	_ = "STUB: not implemented"
	return 0
}

func zopfliCostModelGetDistanceCost(self *zopfliCostModel, distcode uint) float32 {
	_ = "STUB: not implemented"
	return 0
}

func zopfliCostModelGetLiteralCosts(self *zopfliCostModel, from uint, to uint) float32 {
	_ = "STUB: not implemented"
	return 0
}

func zopfliCostModelGetMinCostCmd(self *zopfliCostModel) float32 {
	_ = "STUB: not implemented"
	return 0
}

/* REQUIRES: len >= 2, start_pos <= pos */
/* REQUIRES: cost < kInfinity, nodes[start_pos].cost < kInfinity */
/* Maintains the "ZopfliNode array invariant". */
func updateZopfliNode(nodes []zopfliNode, pos uint, start_pos uint, len uint, len_code uint, dist uint, short_code uint, cost float32) {
	_ = "STUB: not implemented"
	return
}

type posData struct {
	pos            uint
	distance_cache [4]int
	costdiff       float32
	cost           float32
}

/* Maintains the smallest 8 cost difference together with their positions */
type startPosQueue struct {
	q_   [8]posData
	idx_ uint
}

func initStartPosQueue(self *startPosQueue) { _ = "STUB: not implemented"; return }

func startPosQueueSize(self *startPosQueue) uint { _ = "STUB: not implemented"; return 0 }

func startPosQueuePush(self *startPosQueue, posdata *posData) { _ = "STUB: not implemented"; return }

/* Restore the sorted order. In the list of |len| items at most |len - 1|
   adjacent element comparisons / swaps are required. */

func startPosQueueAt(self *startPosQueue, k uint) *posData { _ = "STUB: not implemented"; return nil }

/* Returns the minimum possible copy length that can improve the cost of any */
/* future position. */
func computeMinimumCopyLength(start_cost float32, nodes []zopfliNode, num_bytes uint, pos uint) uint {
	_ = "STUB: not implemented"
	return 0
}

/* Compute the minimum possible cost of reaching any future position. */

/* We already reached (pos + len) with no more cost than the minimum
   possible cost of reaching anything from this pos, so there is no point in
   looking for lengths <= len. */

/* We reached the next copy length code bucket, so we add one more
   extra bit to the minimum cost. */

/*
REQUIRES: nodes[pos].cost < kInfinity

	REQUIRES: nodes[0..pos] satisfies that "ZopfliNode array invariant".
*/
func computeDistanceShortcut(block_start uint, pos uint, max_backward_limit uint, gap uint, nodes []zopfliNode) uint32 {
	_ = "STUB: not implemented"
	return 0
}

/* Since |block_start + pos| is the end position of the command, the copy part
   starts from |block_start + pos - clen|. Distances that are greater than
   this or greater than |max_backward_limit| + |gap| are static dictionary
   references, and do not update the last distances.
   Also distance code 0 (last distance) does not update the last distances. */

/*
Fills in dist_cache[0..3] with the last four distances (as defined by

	Section 4. of the Spec) that would be used at (block_start + pos) if we
	used the shortest path of commands from block_start, computed from
	nodes[0..pos]. The last four distances at block_start are in
	starting_dist_cache[0..3].
	REQUIRES: nodes[pos].cost < kInfinity
	REQUIRES: nodes[0..pos] satisfies that "ZopfliNode array invariant".
*/
func computeDistanceCache(pos uint, starting_dist_cache []int, nodes []zopfliNode, dist_cache []int) {
	_ = "STUB: not implemented"
	return
}

/* Because of prerequisite, p >= clen + ilen >= 2. */

/*
Maintains "ZopfliNode array invariant" and pushes node to the queue, if it

	is eligible.
*/
func evaluateNode(block_start uint, pos uint, max_backward_limit uint, gap uint, starting_dist_cache []int, model *zopfliCostModel, queue *startPosQueue, nodes []zopfliNode) {
	_ = "STUB: not implemented"
	/* Save cost, because ComputeDistanceCache invalidates it. */ return
}

/* Returns longest copy length. */
func updateNodes(num_bytes uint, block_start uint, pos uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, max_backward_limit uint, starting_dist_cache []int, num_matches uint, matches []backwardMatch, model *zopfliCostModel, queue *startPosQueue, nodes []zopfliNode) uint {
	_ = "STUB: not implemented"
	return 0
}

/* Go over the command starting positions in order of increasing cost
   difference. */

/* Look for last distance matches using the distance cache from this
   starting position. */

/* Word dictionary -> ignore. */

/* Regular backward reference. */

/* At higher iterations look only for new last distance matches, since
   looking only for new command start positions with the same distances
   does not help much. */

/* Loop through all possible copy lengths at this position. */

/* We already tried all possible last distance matches, so we can use
   normal distance code here. */

/* Try all copy lengths up until the maximum copy length corresponding
   to this distance. If the distance refers to the static dictionary, or
   the maximum length is long enough, try only one maximum length. */

func computeShortestPathFromNodes(num_bytes uint, nodes []zopfliNode) uint {
	_ = "STUB: not implemented"
	return 0
}

/* REQUIRES: nodes != NULL and len(nodes) >= num_bytes + 1 */
func zopfliCreateCommands(num_bytes uint, block_start uint, nodes []zopfliNode, dist_cache []int, last_insert_len *uint, params *encoderParams, commands *[]command, num_literals *uint) {
	_ = "STUB: not implemented"
	return
}

func zopfliIterate(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, gap uint, dist_cache []int, model *zopfliCostModel, num_matches []uint32, matches []backwardMatch, nodes []zopfliNode) uint {
	_ = "STUB: not implemented"
	return 0
}

/*
Computes the shortest path of commands from position to at most

	position + num_bytes.

	On return, path->size() is the number of commands found and path[i] is the
	length of the i-th command (copy length plus insert length).
	Note that the sum of the lengths of all commands can be less than num_bytes.

	On return, the nodes[0..num_bytes] array will have the following
	"ZopfliNode array invariant":
	For each i in [1..num_bytes], if nodes[i].cost < kInfinity, then
	  (1) nodes[i].copy_length() >= 2
	  (2) nodes[i].command_length() <= i and
	  (3) nodes[i - nodes[i].command_length()].cost < kInfinity

REQUIRES: nodes != nil and len(nodes) >= num_bytes + 1
*/
func zopfliComputeShortestPath(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, dist_cache []int, hasher *h10, nodes []zopfliNode) uint {
	_ = "STUB: not implemented"
	return 0
}

/* Add the tail of the copy to the hasher. */

func createZopfliBackwardReferences(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, hasher *h10, dist_cache []int, last_insert_len *uint, commands *[]command, num_literals *uint) {
	_ = "STUB: not implemented"
	return
}

func createHqZopfliBackwardReferences(num_bytes uint, position uint, ringbuffer []byte, ringbuffer_mask uint, params *encoderParams, hasher hasherHandle, dist_cache []int, last_insert_len *uint, commands *[]command, num_literals *uint) {
	_ = "STUB: not implemented"
	return
}

/* Ensure that we have enough free slots. */

/* Add the tail of the copy to the hasher. */
