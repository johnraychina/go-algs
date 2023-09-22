package basic

import (
	"cmp"
	"fmt"
)

// IndexMinPQ 支持按索引访问（注意，index不重复，key可以重复）的最小堆。
type IndexMinPQ[K cmp.Ordered, V ComparableObj[K]] struct {
	n  int   // number of elements
	pq []int // position->index, binary heap using 1-based indexing
	a  []V   // a[index] = priority of index
	qp []int // index->position, inverse of pq,  qp[pq[i]] = pq[qp[i]] = i
}

func NewIndexMinPQ[K cmp.Ordered, V ComparableObj[K]](maxN int) *IndexMinPQ[K, V] {
	return &IndexMinPQ[K, V]{
		n:  0,
		pq: make([]int, maxN+1), // 1-based indexing
		a:  make([]V, maxN+1),
		qp: make([]int, maxN+1),
	}
}

func (q *IndexMinPQ[K, V]) Insert(index int, v V) {
	q.n++ // 1-based indexing
	q.pq[q.n] = index
	q.qp[index] = q.n
	q.a[index] = v
	q.swim(len(q.pq) - 1)
}

func (q *IndexMinPQ[K, V]) Get(idx int) V {
	q.validateIndex(idx)
	if !q.Contains(idx) {
		panic(fmt.Sprintf("index is not in the priority queue"))
	} else {
		return q.a[idx]
	}
}

// MinIndex Returns an index associated with a minimum key.
func (q *IndexMinPQ[K, V]) MinIndex() int {
	if q.n == 0 {
		panic("Priority queue underflow")
	}
	return q.pq[1]
}

func (q *IndexMinPQ[K, V]) swim(k int) {
	for k > 1 {
		i := k / 2
		if !q.greater(i, k) { // stop swimming
			break
		}
		q.swap(i, k)
		k = i
	}
}

func (q *IndexMinPQ[K, V]) sink(k int) {
	for 2*k < len(q.a) { // 避免越界
		i := 2 * k
		if i+1 < len(q.a) && q.greater(i, i+1) {
			i++
		}
		if !q.greater(k, i) {
			break
		}
		q.swap(k, i)
		k = i
	}
}

func (q *IndexMinPQ[K, V]) swap(l int, r int) {
	// 交换索引位置
	q.pq[l], q.pq[l] = q.pq[r], q.pq[l]
	// 更新qp(index->position) 与 pq(position->index)保持一致
	q.qp[q.pq[l]] = l
	q.qp[q.pq[r]] = r
}

func (q *IndexMinPQ[K, V]) greater(l int, r int) bool {
	return q.a[q.pq[l]].Key() > q.a[q.pq[r]].Key()
}

func (q *IndexMinPQ[K, V]) IsEmpty() bool {
	return len(q.pq) <= 1 // a[0] is useless
}

// DelMin Removes a minimum key and returns its associated index.
// return an index associated with a minimum key
// throws NoSuchElementException if this priority queue is empty
func (q *IndexMinPQ[K, V]) DelMin() int {
	// hold the index of min value
	indexOfMinKey := q.pq[1]

	// swap min with last
	q.swap(1, q.n)
	q.n--
	// sink down
	q.sink(1)

	var zeroVal V
	q.qp[indexOfMinKey] = -1     // delete
	q.a[indexOfMinKey] = zeroVal // to help with garbage collection
	q.pq[q.n+1] = -1             // not needed, as pq is guarded by q.n

	return indexOfMinKey
}

func (q *IndexMinPQ[K, V]) Contains(idx int) bool {
	q.validateIndex(idx)
	return q.qp[idx] != -1
}

func (q *IndexMinPQ[K, V]) validateIndex(idx int) {
	if idx < 0 {
		panic(fmt.Sprintf("index is negative:%d", idx))
	}
	if idx >= len(q.a) {
		panic(fmt.Sprintf("index >= capacity: %d", idx))
	}
}

// todo
func (q *IndexMinPQ[K, V]) DecreaseKey(w int, weight float32) {

}
