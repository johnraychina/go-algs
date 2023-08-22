package basic

import (
	"fmt"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	maxQ := NewMaxPQ()
	maxQ.insert(2)
	maxQ.insert(3)
	maxQ.insert(4)
	maxQ.insert(3)
	maxQ.insert(1)

	fmt.Println(maxQ.a)

	fmt.Println("show max")
	for !maxQ.isEmpty() {
		fmt.Println(maxQ.delMax())
	}

}
