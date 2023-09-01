package basic

import (
	"fmt"
	"testing"
)


func TestRedBlackTreeGet(t *testing.T) {
	st := NewRedBlackTree[string, string]()
	st.Put("alice", "New York")
	st.Put("bob", "Chicago")
	st.Put("cathy", "Washington")
	st.Put("david", "Beijing")

	
	fmt.Println(st.Get("alice"))
	fmt.Println(st.Get("david"))
}

