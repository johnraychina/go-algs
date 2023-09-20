package basic

import (
	"fmt"
	"testing"
)

type Int struct {
	val int
}

func (i *Int) Key() int {
	return i.val
}

func TestMinPQ(t *testing.T) {
	maxQ := NewMinPQ[*Int]()
	maxQ.Insert(&Int{val: 2})
	maxQ.Insert(&Int{val: 3})
	maxQ.Insert(&Int{val: 4})
	maxQ.Insert(&Int{val: 3})
	maxQ.Insert(&Int{val: 1})
	fmt.Println(maxQ.a)

	fmt.Println("show min")
	for !maxQ.isEmpty() {
		fmt.Println(maxQ.DelMin())
	}

}
