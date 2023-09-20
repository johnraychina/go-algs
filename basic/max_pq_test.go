package basic

import (
	"fmt"
	"testing"
)

func TestMaxPQ(t *testing.T) {
	maxQ := NewMaxPQ()
	maxQ.Insert(2)
	maxQ.Insert(3)
	maxQ.Insert(4)
	maxQ.Insert(3)
	maxQ.Insert(1)

	fmt.Println(maxQ.a)

	fmt.Println("show max")
	for !maxQ.isEmpty() {
		fmt.Println(maxQ.DelMax())
	}

}
