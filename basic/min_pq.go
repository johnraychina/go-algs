package basic

import "cmp"

// MinPQ
// 注意：
// golang泛型无法像java做这种constraints嵌套 type MinPQ[V ComparableObj[K cmp.Ordered]]
// 但是可以分别做constraint声明
type MinPQ[K cmp.Ordered, V ComparableObj[K]] struct {
	a []V
}

func NewMinPQ[K cmp.Ordered, V ComparableObj[K]]() *MinPQ[K, V] {
	return &MinPQ[K, V]{a: make([]V, 1)} // a[0]=0 is useless, a[1] is the max item
}

func (q *MinPQ[K, V]) Insert(v V) {
	// append to the last -> swim
	q.a = append(q.a, v)
	q.swim(len(q.a) - 1)
}

func (q *MinPQ[K, V]) swim(k int) {
	for k > 1 {
		i := k / 2
		if !q.greater(i, k) { // stop swimming
			break
		}
		q.swap(i, k)
		k = i
	}
}

func (q *MinPQ[K, V]) sink(k int) {
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

func (q *MinPQ[K, V]) swap(l int, r int) {
	q.a[l], q.a[r] = q.a[r], q.a[l]
}

func (q *MinPQ[K, V]) greater(l int, r int) bool {
	return q.a[l].Key() > q.a[r].Key()
}

func (q *MinPQ[K, V]) IsEmpty() bool {
	return len(q.a) <= 1 // a[0] is useless
}

func (q *MinPQ[K, V]) DelMin() V {
	// hold the max value
	v := q.a[1]

	// copy last to first
	lastIdx := len(q.a) - 1
	q.a[1] = q.a[lastIdx]
	// slice, ignore the last
	q.a = q.a[:lastIdx]

	// sink down
	q.sink(1)

	return v
}
