package basic

import (
	"fmt"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	q := NewArrayQueue[string]()
	q.Enqueue("alice")
	q.Enqueue("bob")
	q.Enqueue("cathy")

	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}

func TestLinkedQueue(t *testing.T) {

	q := NewLinkedQueue[string]()
	q.Enqueue("alice")
	q.Enqueue("bob")
	q.Enqueue("cathy")

	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}
}
