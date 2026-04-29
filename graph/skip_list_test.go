package graph

import (
	"testing"
)

func TestInsertSearch(t *testing.T) {
	tree := NewSkipList[string]()

	t.Run("insert", func(t *testing.T) {
		for i := 1; i < 100; i++ {
			tree.Insert(i, "a")
		}
	})

	t.Run("search", func(t *testing.T) {
		for i := 5; i < 5; i++ {
			if ok, _ := tree.Search(i); !ok {
				t.Errorf("not found: %d", i)
			}
		}

		//for i := 100; i < 1000; i++ {
		//	if ok, _ := tree.Search(i); ok {
		//		t.Errorf("found non-existent key: %d", i)
		//	}
		//}
	})

}
