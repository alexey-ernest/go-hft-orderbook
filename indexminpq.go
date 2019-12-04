package hftorderbook

// Indexed mininum oriented Priority Queue
type indexMinPQ struct {
	keys []float64
	index2offset []int
	offset2index []int
	n int
}

func NewIndexMinPQ(size int) indexMinPQ {
	return indexMinPQ {
		keys: make([]float64, size + 1),
		index2offset: make([]int, size + 1),
		offset2index: make([]int, size + 1),
	}
}

func (pq *indexMinPQ) Size() int {
	return pq.n
}

func (pq *indexMinPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *indexMinPQ) Insert(i int, key float64) {
	pq.checkIndex(i)

	if pq.index2offset[i] > 0 {
		panic("index already used")
	}
	if pq.n + 1 == cap(pq.keys) {
		panic("pq is full")
	}

	pq.n++
	pq.keys[pq.n] = key
	pq.index2offset[i] = pq.n
	pq.offset2index[pq.n] = i

	// restore order
	pq.swim(i)
}

func (pq *indexMinPQ) Change(i int, key float64) {
	pq.checkIndex(i)

	offset := pq.index2offset[i]
	if offset == 0 {
		panic("a key does not exist")
	}

	// updating the key
	k := pq.keys[offset]
	pq.keys[offset] = key

	// restore order
	if key > k {
		pq.sink(i)
	} else if key < k {
		pq.swim(i)
	}
}

func (pq *indexMinPQ) Contains(i int) bool {
	pq.checkIndex(i)

	return pq.index2offset[i] > 0
}

func (pq *indexMinPQ) Delete(i int) {
	pq.checkIndex(i)

	offset := pq.index2offset[i]
	if offset == 0 {
		panic("invalid index")
	}

	// replace key with the lask key
	pq.keys[offset] = pq.keys[pq.n]

	// update indexes
	lastkeyindex := pq.offset2index[pq.n]
	pq.index2offset[lastkeyindex] = offset
	pq.offset2index[offset] = lastkeyindex
	
	// nullify removed data
	pq.offset2index[pq.n] = 0
	pq.index2offset[i] = 0

	pq.n--

	// restore order
	pq.sink(lastkeyindex)
}

func (pq *indexMinPQ) Top() float64 {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	return pq.keys[1]
}

func (pq *indexMinPQ) TopIndex() int {
	if pq.IsEmpty() {
		panic("pq is empty")
	}

	return pq.offset2index[1]
}

// removes minimal element and returns it's index
func (pq *indexMinPQ) DelTop() int {
	minindex := pq.TopIndex()
	pq.Delete(minindex)
	return minindex
}


// helpers

func (pq *indexMinPQ) checkIndex(i int) {
	if i < 0 || i + 1 >= cap(pq.keys) {
		panic("invalid index")
	}
}

func (pq *indexMinPQ) swim(i int) {
	k := pq.index2offset[i]
	for k > 1 && pq.keys[k] < pq.keys[k/2] {
		// swap keys
		pq.keys[k], pq.keys[k/2] = pq.keys[k/2], pq.keys[k]

		// swap indexes
		kid := pq.offset2index[k]
		k2id := pq.offset2index[k/2]
		pq.index2offset[kid], pq.index2offset[k2id] = pq.index2offset[k2id], pq.index2offset[kid]
		pq.offset2index[k], pq.offset2index[k/2] = pq.offset2index[k/2], pq.offset2index[k]

		k = k/2
	}
}

func (pq *indexMinPQ) sink(i int) {
	k := pq.index2offset[i]
	for 2*k <= pq.n {
		c := 2*k

		// select minimum of two children
		if c < pq.n && pq.keys[c+1] < pq.keys[c] {
			c++
		}

		if pq.keys[k] > pq.keys[c] {
			// swap keys
			pq.keys[k], pq.keys[c] = pq.keys[c], pq.keys[k]

			// swap indexes
			kid := pq.offset2index[k]
			cid := pq.offset2index[c]
			pq.index2offset[kid], pq.index2offset[cid] = pq.index2offset[cid], pq.index2offset[kid]
			pq.offset2index[k], pq.offset2index[c] = pq.offset2index[c], pq.offset2index[k]

			k = c
		} else {
			break
		}
	}
}
