package basic

import (
	"cmp"
	"fmt"
)

// IndexMinPQ 支持按索引访问（注意，index不重复，key可以重复）的最小堆。
type IndexMinPQ[K cmp.Ordered] struct {
	n    int   // number of elements
	maxN int   // max number of elements
	pq   []int // position->index, binary heap using 1-based indexing
	a    []K   // a[index] = priority of index
	qp   []int // index->position, inverse of pq,  qp[pq[i]] = pq[qp[i]] = i
}

func NewIndexMinPQ[K cmp.Ordered](maxN int) *IndexMinPQ[K] {
	q := &IndexMinPQ[K]{
		n:    0,
		maxN: maxN,
		pq:   make([]int, maxN+1), // 1-based indexing
		a:    make([]K, maxN+1),
		qp:   make([]int, maxN+1),
	}
	for i, _ := range q.qp {
		q.qp[i] = -1
	}
	return q
}

func (q *IndexMinPQ[K]) Insert(index int, k K) {
	q.validateIndex(index)
	q.n++ // 1-based indexing
	q.pq[q.n] = index
	q.qp[index] = q.n
	q.a[index] = k
	q.swim(q.n)
}

func (q *IndexMinPQ[K]) Get(idx int) K {
	q.validateIndex(idx)
	if !q.Contains(idx) {
		panic(fmt.Sprintf("index is not in the priority queue"))
	} else {
		return q.a[idx]
	}
}

// MinIndex Returns an index associated with a minimum key.
func (q *IndexMinPQ[K]) MinIndex() int {
	if q.n == 0 {
		panic("Priority queue underflow")
	}
	return q.pq[1]
}

func (q *IndexMinPQ[K]) swim(k int) {
	for k > 1 {
		i := k / 2
		if !q.greater(i, k) { // stop swimming
			break
		}
		q.swap(i, k)
		k = i
	}
}

func (q *IndexMinPQ[K]) sink(k int) {
	for 2*k <= q.n { // 避免越界
		i := 2 * k
		if i < q.n && q.greater(i, i+1) {
			i++
		}
		if !q.greater(k, i) {
			break
		}
		q.swap(k, i)
		k = i
	}
}

func (q *IndexMinPQ[K]) swap(a int, b int) {
	// 交换索引位置
	q.pq[a], q.pq[b] = q.pq[b], q.pq[a]
	// 更新qp(index->position) 与 pq(position->index)保持一致
	q.qp[q.pq[a]] = a
	q.qp[q.pq[b]] = b
}

func (q *IndexMinPQ[K]) greater(l int, r int) bool {
	// pos->index->value
	return q.a[q.pq[l]] > q.a[q.pq[r]]
}

func (q *IndexMinPQ[K]) IsEmpty() bool {
	return q.n == 0 // a[0] is useless
}

// DelMin Removes a minimum key and returns its associated index.
// return an index associated with a minimum key
// throws NoSuchElementException if this priority queue is empty
func (q *IndexMinPQ[K]) DelMin() int {
	if q.IsEmpty() {
		panic("already empty")
	}

	// hold the index of min value
	indexOfMinKey := q.pq[1]

	// swap min with last
	q.swap(1, q.n)
	q.n--
	// sink down
	q.sink(1)

	var zeroVal K
	q.qp[indexOfMinKey] = -1     // delete
	q.a[indexOfMinKey] = zeroVal // to help with garbage collection
	q.pq[q.n+1] = -1             // not needed, as pq is guarded by q.n

	return indexOfMinKey
}

func (q *IndexMinPQ[K]) Contains(idx int) bool {
	q.validateIndex(idx)
	return q.qp[idx] != -1
}

func (q *IndexMinPQ[K]) validateIndex(idx int) {
	if idx < 0 {
		panic(fmt.Sprintf("index is negative:%d", idx))
	}
	if idx >= q.maxN {
		panic(fmt.Sprintf("index >= capacity: %d", idx))
	}
}

func (q *IndexMinPQ[K]) DecreaseKey(idx int, key K) {
	q.validateIndex(idx)
	if !q.Contains(idx) {
		panic("index does not exist")
	}
	if q.a[idx] == key {
		panic("DecreaseKey with an equal key")
	}
	if q.a[idx] < key {
		panic("DecreaseKey with a larger key")
	}

	//update weight
	q.a[idx] = q.a[idx]
	// swim up
	q.swim(q.qp[idx])
}
