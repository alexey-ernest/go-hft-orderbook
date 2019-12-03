package types

import "fmt"

// A self-balancing Binary Search Tree with 2*lgN worst case garantees for
// search, put, delete, min, max, select, rank, floor, ceiling operations.
// Average runtine for search-based operations estimated as 1*lgN

type nodeRedBlack struct {
	key float64
	value *LimitOrder
	left *nodeRedBlack
	right *nodeRedBlack
	size int
	isRed bool
}

type redBlackBST struct {
	root *nodeRedBlack
	minC *nodeRedBlack // cached min/max keys for O(1) access
	maxC *nodeRedBlack
}

func NewRedBlackBST() redBlackBST {
	return redBlackBST{}
}

func (t *redBlackBST) Size() int {
	return t.size(t.root)
}

func (t *redBlackBST) size(n *nodeRedBlack) int {
	if n == nil {
		return 0
	}

	return n.size
}

func (t *redBlackBST) IsEmpty() bool {
	return t.size(t.root) == 0
}

func (t *redBlackBST) panicIfEmpty() {
	if t.IsEmpty() {
		panic("Red Black BST is empty")
	}
}

func (t *redBlackBST) Contains(key float64) bool {
	return t.get(t.root, key) != nil
}

func (t *redBlackBST) Get(key float64) *LimitOrder {
	t.panicIfEmpty()

	x := t.get(t.root, key)
	if x == nil {
		panic(fmt.Sprintf("key %0.8f does not exist", key))
	}

	return x.value
}

func (t *redBlackBST) get(n *nodeRedBlack, key float64) *nodeRedBlack {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	}

	if n.key > key {
		return t.get(n.left, key)
	} else {
		return t.get(n.right, key)
	}
}

func (t *redBlackBST) isRed(n *nodeRedBlack) bool {
	if n == nil {
		// nil nodes are black by default
		return false
	}

	return n.isRed
}

func (t *redBlackBST) flipColors(n *nodeRedBlack) {
	if n == nil {
		return
	}

	// inverse children colors
	if n.left != nil {
		n.left.isRed = !n.left.isRed
	}
	if n.right != nil {
		n.right.isRed = !n.right.isRed
	}

	// inverse node color
	n.isRed = !n.isRed
}

func (t *redBlackBST) rotateLeft(n *nodeRedBlack) *nodeRedBlack {
	x := n.right
	n.right = x.left
	x.left = n

	x.isRed = n.isRed
	n.isRed = true

	// re-calculate sizes
	n.size = t.size(n.left) + 1 + t.size(n.right)
	x.size = t.size(x.left) + 1 + t.size(x.right)

	return x
}

func (t *redBlackBST) rotateRight(n *nodeRedBlack) *nodeRedBlack {
	x := n.left
	n.left = x.right
	x.right = n

	x.isRed = n.isRed
	n.isRed = true

	// re-calculate sizes
	n.size = t.size(n.left) + 1 + t.size(n.right)
	x.size = t.size(x.left) + 1 + t.size(x.right)

	return x
}

func (t *redBlackBST) Put(key float64, value *LimitOrder) {
	t.root = t.put(t.root, key, value)

	// keeping root black
	t.root.isRed = false
}

func (t *redBlackBST) put(n *nodeRedBlack, key float64, value *LimitOrder) *nodeRedBlack {
	if n == nil {
		// search miss, creating a new node with a red link as a part of 3- or 4-node
		n := &nodeRedBlack{
			key: key,
			value: value,
			size: 1,
			isRed: true,
		}

		if t.minC == nil || key < t.minC.key {
			// new min
			t.minC = n
		}
		if t.maxC == nil || key > t.maxC.key {
			// new max
			t.maxC = n
		}

		return n
	}

	if n.key == key {
		// search hit, updating the value
		n.value = value
		return n
	}

	if n.key > key {
		n.left = t.put(n.left, key, value)
	} else {
		n.right = t.put(n.right, key, value)
	}

	// balancing the tree
	if t.isRed(n.right) && !t.isRed(n.left) {
		// fixing right leaning red link case
		// this can lead to the next case in upper level
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		// making 4-node
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		// convert 4-node into 3 2-nodes
		t.flipColors(n)
	}

	// re-calc size
	n.size = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *redBlackBST) Height() int {
	if t.IsEmpty() {
		return 0
	}

	return t.height(t.root)
}

func (t *redBlackBST) height(n *nodeRedBlack) int {
	if n == nil {
		return 0
	}

	lheight := t.height(n.left)
	rheight := t.height(n.right)
	
	height := lheight
	if rheight > lheight {
		height = rheight
	}

	return height + 1
}

func (t *redBlackBST) IsRedBlack() bool {
	balanced, _ := t.isBalanced(t.root)
	return balanced && t.is23(t.root)
}

func (t *redBlackBST) isBalanced(n *nodeRedBlack) (bool, int) {
	if n == nil {
		// nil node is black by default
		return true, 1
	}

	lb, l := t.isBalanced(n.left)
	rb, r := t.isBalanced(n.right)

	b := l
	if r > l {
		b = r
	}

	if !t.isRed(n) {
		b += 1
	}

	return lb && rb && l == r, b
}

func (t *redBlackBST) is23(n *nodeRedBlack) bool {
	if n == nil {
		return true
	}

	if t.isRed(n.right) {
		// it should has only left leaning red links
		return false
	}

	if t.isRed(n) && t.isRed(n.left) {
		// no node should be connected by two red links
		return false
	}

	return t.is23(n.left) && t.is23(n.right)
}

func (t *redBlackBST) Min() float64 {
	t.panicIfEmpty()
	return t.minC.key
}

func (t *redBlackBST) MinValue() *LimitOrder {
	t.panicIfEmpty()
	return t.minC.value
}

func (t *redBlackBST) min(n *nodeRedBlack) *nodeRedBlack {
	if n.left == nil {
		return n
	}

	return t.min(n.left)
}

func (t *redBlackBST) Max() float64 {
	t.panicIfEmpty()
	return t.maxC.key
}

func (t *redBlackBST) MaxValue() *LimitOrder {
	t.panicIfEmpty()
	return t.maxC.value
}

func (t *redBlackBST) max(n *nodeRedBlack) *nodeRedBlack {
	if n.right == nil {
		return n
	}

	return t.max(n.right)
}

func (t *redBlackBST) Floor(key float64) float64 {
	t.panicIfEmpty()

	floor := t.floor(t.root, key)
	if floor == nil {
		panic(fmt.Sprintf("there are no keys <= %0.8f", key))
	}

	return floor.key
}

func (t *redBlackBST) floor(n *nodeRedBlack, key float64) *nodeRedBlack {
	if n == nil {
		// search miss
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key > key {
		// floor must be in the left sub-tree
		return t.floor(n.left, key)
	}

	// key could be in the right sub-tree, if not, using current root
	floor := t.floor(n.right, key)
	if floor != nil {
		return floor
	}

	return n
}

func (t *redBlackBST) Ceiling(key float64) float64 {
	t.panicIfEmpty()

	ceiling := t.ceiling(t.root, key)
	if ceiling == nil {
		panic(fmt.Sprintf("there are no keys >= %0.8f", key))
	}

	return ceiling.key
}

func (t *redBlackBST) ceiling(n *nodeRedBlack, key float64) *nodeRedBlack {
	if n == nil {
		// search miss
		return nil
	}

	if n.key == key {
		// search hit
		return n
	}

	if n.key < key {
		// ceiling must be in the right sub-tree
		return t.ceiling(n.right, key)
	}

	// the key could be in the left sub-tree, if not, using current root
	ceiling := t.ceiling(n.left, key)
	if ceiling != nil {
		return ceiling
	}

	return n
}

func (t *redBlackBST) Select(k int) float64 {
	if k < 0 || k >= t.Size() {
		panic("index out of range")
	}

	return t.selectNode(t.root, k).key
}

func (t *redBlackBST) selectNode(n *nodeRedBlack, k int) *nodeRedBlack {
	if t.size(n.left) == k {
		return n
	}

	if t.size(n.left) > k {
		return t.selectNode(n.left, k)
	}

	k = k - t.size(n.left) - 1
	return t.selectNode(n.right, k)
}

func (t *redBlackBST) Rank(key float64) int {
	t.panicIfEmpty()
	return t.rank(t.root, key)
}

func (t *redBlackBST) rank(n *nodeRedBlack, key float64) int {
	if n == nil {
		return 0
	}

	if n.key == key {
		return t.size(n.left)
	}

	if n.key > key {
		return t.rank(n.left, key)
	}

	return t.size(n.left) + 1 + t.rank(n.right, key)
}

func (t *redBlackBST) makeLeftRed(n *nodeRedBlack) *nodeRedBlack {
	// assuming that n.left and n.left.left are black and n is red
	t.flipColors(n)
	// now n is black and both left and right are red

	// fixing red black invariat that no node can be connected with two red links
	if t.isRed(n.right.left) {
		n.right = t.rotateRight(n.right)
		// now n.right and n.right.right are red, fixing that by rotating n
		n = t.rotateLeft(n) 
		// now n.right, n.left and n.left.left are red
	}

	return n
}

func (t *redBlackBST) DeleteMin() {
	t.panicIfEmpty()

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		// making root red temporarily to fit invariant required for makeLeftRed method
		t.root.isRed = true
	}
	t.root = t.deleteMin(t.root)
	if !t.IsEmpty() {
		t.root.isRed = false
	}
}

func (t *redBlackBST) deleteMin(n *nodeRedBlack) *nodeRedBlack {
	if n.left == nil {
		// we've reached the least leave of the tree
		return nil
	}

	// making current node a part of 3 or 4 node by moving red link to the left
	if !t.isRed(n.left) && !t.isRed(n.left.left) {
		n = t.makeLeftRed(n)
	}

	left := n.left
	n.left = t.deleteMin(n.left)

	// updating global min
	if t.minC == left && n.left == nil {
		t.minC = n
	}

	// we have to restore balance of the tree moving from bottom to top now
	if t.isRed(n.right) {
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *redBlackBST) makeRightRed(n *nodeRedBlack) *nodeRedBlack {
	// assuming n is red, n.right and n.right.left are black
	t.flipColors(n)
	// now n is black and n.right is red

	if !t.isRed(n.left.left) {
		// meaning n.left should be red now
		n = t.rotateRight(n)
		// now n.left is black, n.right and n.right.right are red
	}
	return n
}

func (t *redBlackBST) DeleteMax() {
	t.panicIfEmpty()

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.isRed = true
	}
	t.root = t.deleteMax(t.root)
	if !t.IsEmpty() {
		t.root.isRed = false;
	}
}

func (t *redBlackBST) deleteMax(n *nodeRedBlack) *nodeRedBlack {
	if t.isRed(n.left) {
		// making right red by rotating
		n = t.rotateRight(n)
	}
	if n.right == nil {
		return nil
	}

	// making right left on the way from top to bottom
	if !t.isRed(n.right) && !t.isRed(n.right.left) {
		n = t.makeRightRed(n)
	}

	right := n.right
	n.right = t.deleteMax(n.right)

	// updating global max
	if t.maxC == right && n.right == nil {
		t.maxC = n
	}
	
	// balancing back on the way from bottom to top
	if t.isRed(n.right) {
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *redBlackBST) Delete(key float64) {
	t.panicIfEmpty()

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.isRed = true
	}
	t.root = t.delete(t.root, key)
	if !t.IsEmpty() {
		t.root.isRed = false;
	}
}

func (t *redBlackBST) delete(n *nodeRedBlack, key float64) *nodeRedBlack {
	if n.key > key {
		// looking into the left sub-tree
		if !t.isRed(n.left) && !t.isRed(n.left.left) {
			n = t.makeLeftRed(n)
		}

		left := n.left
		n.left = t.delete(n.left, key)
		if t.minC == left && n.left == nil {
			// min was deleted, updating global min
			t.minC = n
		}

	} else {
		// checking current node and right sub-tree if required
		if t.isRed(n.left) {
			n = t.rotateRight(n)
		}
		if n.key == key && n.right == nil {
			return nil
		}

		if !t.isRed(n.right) && !t.isRed(n.right.left) {
			n = t.makeRightRed(n)
		}

		if n.key == key {
			// removing a node from the middle of the tree
			rightMin := t.min(n.right)
			n.key = rightMin.key
			n.value = rightMin.value
			n.right = t.deleteMin(n.right)
		} else {
			right := n.right
			n.right = t.delete(n.right, key)
			if t.maxC == right && n.right == nil {
				// max was deleted, updating global max
				t.maxC = n
			}
		}
	}

	// balance
	if t.isRed(n.right) {
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = t.size(n.left) + 1 + t.size(n.right)
	return n
}

func (t *redBlackBST) Keys(lo, hi float64) []float64 {
	if lo < t.Min() || hi > t.Max() {
		panic("keys out of range")
	}

	return t.keys(t.root, lo, hi)
}

func (t *redBlackBST) keys(n *nodeRedBlack, lo, hi float64) []float64 {
	if n == nil {
		return nil
	}

	if n.key < lo {
		return t.keys(n.right, lo, hi)
	} else if n.key > hi {
		return t.keys(n.left, lo, hi)
	}

	l := t.keys(n.left, lo, hi)
	r := t.keys(n.right, lo, hi)
	
	keys := make([]float64, 0)
	if l != nil {
		keys = append(keys, l...)
	}
	keys = append(keys, n.key)
	if r != nil {
		keys = append(keys, r...)
	}

	return keys
}
