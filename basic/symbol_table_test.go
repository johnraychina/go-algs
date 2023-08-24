package basic

import (
	"fmt"
	"testing"
)

func TestSymbolTable(t *testing.T) {

	st := NewBST[string, int]()
	st.Put("alice", 1)
	st.Put("bob", 2)
	st.Put("cathy", 3)

	////func TestSymbolTable(t *testing.T) {
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
