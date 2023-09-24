package basic

import (
	"fmt"
	"testing"
)

type Student struct {
	name string
	age  int
}

func (s *Student) Key() int {
	return s.age
}

func TestInsert(t *testing.T) {

	q := NewIndexMinPQ[int, *Student](5)
	q.Insert(1, &Student{"Tom", 15})
	q.Insert(2, &Student{"Read", 14})
	q.Insert(3, &Student{"Jenny", 13})
	q.Insert(4, &Student{"Kent", 16})
	q.Insert(5, &Student{"Joey", 14})

	// q.get
	//for i := 1; i < 6; i++ {
	//	fmt.Println(q.Get(i))
	//}

	// get min
	for !q.IsEmpty() {
		minIdx := q.DelMin()
		fmt.Println(minIdx)
	}
}
