package graph

import (
	"fmt"
	"testing"
)

func TestInsertSearch(t *testing.T) {
	tree := NewSkipList[string]()
	tree.Print()

	t.Run("insert", func(t *testing.T) {
		for i := 1; i <= 1024; i++ {
			tree.Insert(i, "a")
		}
		tree.Print()
		fmt.Println("----------------------------")
	})

	//t.Run("search", func(t *testing.T) {
	//for i := 1; i < 10; i++ {
	//	if ok, _ := tree.Search(i); !ok {
	//		t.Errorf("not found: %d", i)
	//	}
	//}

	//for i := 100; i < 1000; i++ {
	//	if ok, _ := tree.Search(i); ok {
	//		t.Errorf("found non-existent key: %d", i)
	//	}
	//}
	//})

	//tree.Print()

}
