package regex

import (
	"fmt"
	"testing"
)

func TestNFA(t *testing.T) {
	//nfa := NewNFA("")
	//nfa := NewNFA("abc")
	//nfa := NewNFA("abc*")
	//nfa := NewNFA(".*(a|b)cd")
	nfa := NewNFA(".*((a|b)*cd)ef.*")

	match := nfa.Recognize("---ab--abbbcdeffffff---")
	fmt.Println(match)
}
