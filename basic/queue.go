package basic

type ArrayQueue[V any] struct {
	a []V
}

func NewArrayQueue[V any]() *ArrayQueue[V] {
	return &ArrayQueue[V]{}
}

func (q *ArrayQueue[V]) IsEmpty() bool {
	return len(q.a) == 0
}

func (q *ArrayQueue[V]) Enqueue(val V) {
	q.a = append(q.a, val)
}

func (q *ArrayQueue[V]) Dequeue() (val V) {
	first := q.a[0]
	q.a = q.a[1:]
	return first
}

type LinkedQueue[V any] struct {
	head *Node[V] // pseudo head
	tail *Node[V] // tail
}

func NewLinkedQueue[V any]() *LinkedQueue[V] {
	h := &Node[V]{}
	t := h
	h.next = t
	t.prev = h

	return &LinkedQueue[V]{head: h, tail: t}
}

func (q *LinkedQueue[V]) IsEmpty() bool {
	return q.head == q.tail
}

func (q *LinkedQueue[V]) Enqueue(val V) {
	newNode := &Node[V]{val: val}

	q.tail.next = newNode
	newNode.prev = q.tail

	q.tail = newNode
}

func (q *LinkedQueue[V]) Dequeue() (val V) {
	if q.IsEmpty() {
		return val
	}

	// 摘除节点前，先暂存一下
	val = q.head.next.val
	// 重组关系

	q.head.next = q.head.next.next
	if q.head.next != nil { // end of queue
		q.head.next.prev = q.head
	} else {
		q.tail = q.head
	}

	return val
}
