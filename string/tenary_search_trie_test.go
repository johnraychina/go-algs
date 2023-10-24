package string

import (
	"fmt"
	"testing"
)

func TestTST(t *testing.T) {
	trie := NewTrieTST[int]()

	trie.Put("She", 1)
	trie.Put("Shell", 2)
	trie.Put("Should", 3)

	fmt.Println(trie.Get("She"))
	fmt.Println(trie.Get("Shell"))
	fmt.Println(trie.Get("Should"))
	fmt.Println(trie.Get("Shall"))

	//trie.Delete("Shell")
	//trie.Delete("Should")

}
