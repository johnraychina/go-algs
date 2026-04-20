package basic

import "cmp"

// DWayHeap is a d-way min heap (Johnson 1975).
// Each node has up to d children instead of 2, reducing tree height to O(log_d n).
// For 1-based indexing:
//   - children of node k: d*(k-1)+2  ...  d*(k-1)+d+1
//   - parent of node k:   (k-2)/d + 1
type DWayHeap[K cmp.Ordered] struct {
	d int  // branching factor
	n int  // number of elements
	a []K  // heap array, 1-based indexing
}

func NewDWayHeap[K cmp.Ordered](d int) *DWayHeap[K] {
	if d < 2 {
		panic("branching factor must be >= 2")
	}
	return &DWayHeap[K]{
		d: d,
		n: 0,
		a: make([]K, 1), // a[0] is unused (1-based)
	}
}

func (h *DWayHeap[K]) Insert(key K) {
	h.a = append(h.a, key)
	h.n++
	h.swim(h.n)
}

func (h *DWayHeap[K]) Min() K {
	if h.n == 0 {
		panic("priority queue underflow")
	}
	return h.a[1]
}

func (h *DWayHeap[K]) DelMin() K {
	if h.n == 0 {
		panic("priority queue underflow")
	}
	min := h.a[1]
	h.swap(1, h.n)
	h.n--
	h.a = h.a[:h.n+1]
	h.sink(1)
	return min
}

func (h *DWayHeap[K]) IsEmpty() bool {
	return h.n == 0
}

func (h *DWayHeap[K]) Size() int {
	return h.n
}

// parent of node k: (k-2)/d + 1
func (h *DWayHeap[K]) swim(k int) {
	for k > 1 {
		p := (k-2)/h.d + 1
		if h.a[p] <= h.a[k] {
			break
		}
		h.swap(p, k)
		k = p
	}
}

// children of node k: d*(k-1)+2 .. d*(k-1)+d+1
func (h *DWayHeap[K]) sink(k int) {
	for {
		firstChild := h.d*(k-1) + 2
		if firstChild > h.n {
			break
		}
		// find the smallest child
		minChild := firstChild
		lastChild := firstChild + h.d - 1
		if lastChild > h.n {
			lastChild = h.n
		}
		for c := firstChild + 1; c <= lastChild; c++ {
			if h.a[c] < h.a[minChild] {
				minChild = c
			}
		}
		if h.a[k] <= h.a[minChild] {
			break
		}
		h.swap(k, minChild)
		k = minChild
	}
}

func (h *DWayHeap[K]) swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}
