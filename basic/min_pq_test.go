package basic

import (
	"fmt"
	"testing"
)

func TestMinPQ(t *testing.T) {
	maxQ := NewMinPQ[int, *Int]()
	maxQ.Insert(&Int{k: 2})
	maxQ.Insert(&Int{k: 3})
	maxQ.Insert(&Int{k: 4})
	maxQ.Insert(&Int{k: 3})
	maxQ.Insert(&Int{k: 1})
	fmt.Println(maxQ.a)

	fmt.Println("show min")
	for !maxQ.IsEmpty() {
		fmt.Println(maxQ.DelMin())
	}

}
