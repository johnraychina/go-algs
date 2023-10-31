package sub_str

import (
	"fmt"
	"testing"
)

func TestKMPSearch(t *testing.T) {
	dfa := BuildDFA("AAAAB")
	index := dfa.KMPSearch("AAAACAAABAAAAB")

	fmt.Println(index)
}
