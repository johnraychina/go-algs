package basic

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		fmt.Println(v)
	}

	fmt.Println("size:", stack.Size())
}
