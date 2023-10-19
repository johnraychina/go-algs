package string

import (
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	trie := NewTrieST[int](256)

	trie.Put("She", 1)
	trie.Put("Shell", 2)
	trie.Put("Should", 3)

	fmt.Println(trie.Get("She"))
	fmt.Println(trie.Get("Shell"))
	fmt.Println(trie.Get("Should"))
	fmt.Println(trie.Get("Shall"))

	fmt.Println("------delete-----")
	trie.Delete("She")
	trie.Delete("Shell")
	fmt.Println(trie.Get("She"))
	fmt.Println(trie.Get("Shell"))
	fmt.Println(trie.Get("Should"))
	fmt.Println(trie.Get("Shall"))

	//trie.Delete("Shell")
	//trie.Delete("Should")

}
