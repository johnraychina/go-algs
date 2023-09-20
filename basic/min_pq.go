package basic

type MinPQ[V ComparableKey[K comparable]] struct {
	a []V
}

func NewMinPQ[V ComparableKey[int]]() *MinPQ[V] {
	return &MinPQ[V]{a: make([]V, 1)} // a[0]=0 is useless, a[1] is the max item
}

func (q *MinPQ[V]) Insert(v V) {
	// append to the last -> swim
	q.a = append(q.a, v)
	q.swim(len(q.a) - 1)
}

func (q *MinPQ[V]) swim(k int) {
	for k > 1 {
		i := k / 2
		if !q.greater(i, k) { // stop swimming
			break
		}
		q.swap(i, k)
		k = i
	}
}

func (q *MinPQ[V]) sink(k int) {
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

func (q *MinPQ[V]) swap(l int, r int) {
	q.a[l], q.a[r] = q.a[r], q.a[l]
}

func (q *MinPQ[V]) greater(l int, r int) bool {
	return q.a[l].Key() > q.a[r].Key()
}

func (q *MinPQ[V]) isEmpty() bool {
	return len(q.a) <= 1 // a[0] is useless
}

func (q *MinPQ[V]) DelMin() V {
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
