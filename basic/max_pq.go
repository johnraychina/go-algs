package basic

type MaxPQ struct {
	a []int
}

func NewMaxPQ() *MaxPQ {
	return &MaxPQ{a: []int{0}} // a[0]=0 is useless, a[1] is the max item
}

func (q *MaxPQ) insert(v int) {
	// append to the last -> swim
	q.a = append(q.a, v)
	q.swim(len(q.a) - 1)
}

func (q *MaxPQ) swim(k int) {
	for k > 1 {
		i := k / 2
		if !q.less(i, k) { // stop swimming
			break
		}
		q.swap(i, k)
		k = i
	}
}

func (q *MaxPQ) sink(k int) {
	for 2*k < len(q.a) { // 避免越界
		i := 2 * k
		if i+1 < len(q.a) && q.less(i, i+1) {
			i++
		}
		if !q.less(k, i) {
			break
		}
		q.swap(k, i)
		k = i
	}
}

func (q *MaxPQ) swap(l int, r int) {
	q.a[l], q.a[r] = q.a[r], q.a[l]
}

func (q *MaxPQ) less(l int, r int) bool {
	return q.a[l] < q.a[r]
}

func (q *MaxPQ) isEmpty() bool {
	return len(q.a) <= 1 // a[0] is useless
}

func (q *MaxPQ) delMax() int {
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
