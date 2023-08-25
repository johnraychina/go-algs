package basic

import (
	"fmt"
	"testing"
)

func TestSymbolTableDelete(t *testing.T) {
	st := NewBST[string, int]()
	st.Put("bob", 2)
	st.Put("alice", 1)
	st.Put("cathy", 3)

	for _, k := range st.Keys() {
		fmt.Println(k, ":", st.Get(k))
	}

	st.Delete("bob")
	st.Delete("alice")
	st.Delete("cathy")
}

func TestSymbolTablePutGet(t *testing.T) {

	st := NewBST[string, int]()
	st.Put("bob", 2)
	st.Put("alice", 1)
	st.Put("cathy", 3)

	////func TestSymbolTablePutGet(t *testing.T) {
	//scanner := bufio.NewScanner(os.Stdin)
	//
	//// scan line by line from standard input
	//// put text line to the symbol table
	//for scanner.Scan() {
	//	text := scanner.Text()
	//	if len(text) == 0 {
	//		fmt.Println("end")
	//		break
	//	}
	//	columns := strings.Split(text, " ")
	//	st.Put(columns[0], columns[1])
	//}

	for _, k := range st.Keys() {
		fmt.Println(k, ":", st.Get(k))
	}
}
