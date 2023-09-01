package basic

import (
	"fmt"
	"testing"
)

func TestSymbolTableSize(t *testing.T) {

}

func TestKeysOfRange(t *testing.T) {
	st := NewBST[string, int]()
	st.Put("alice", 0)
	st.Put("bob", 1)
	st.Put("cathy", 3)
	st.Put("david", 2)

	fmt.Println(st.KeysOfRange(0, 0))
	fmt.Println(st.KeysOfRange(1, 2))
	fmt.Println(st.KeysOfRange(3, 3))
}


func TestSelect(t *testing.T) {
	st := NewBST[string, int]()
	st.Put("cathy", 3)
	st.Put("david", 4)
	st.Put("alice", 1)
	st.Put("bob", 2)

	for i := 0; i < st.Size(); i++ {
		fmt.Println(st.Select(i))
	}
}

func TestDeleteKey(t *testing.T) {
	st := NewBST[string, int]()
	st.Put("bob", 2)
	st.Put("alice", 1)
	st.Put("cathy", 3)
	st.Put("david", 4)

	fmt.Println(st.Keys())
	st.Delete("cathy")
	fmt.Println(st.Keys())
	st.Delete("david")
	fmt.Println(st.Keys())
	st.Delete("john")
	fmt.Println(st.Keys())
}

func TestDeleteMin(t *testing.T) {
	st := NewBST[string, int]()
	st.Put("bob", 2)
	st.Put("alice", 1)
	st.Put("cathy", 3)
	st.Put("david", 4)

	fmt.Println(st.Min(), ":", st.Get(st.Min()))
	st.DeleteMin()
	fmt.Println(st.Min(), ":", st.Get(st.Min()))
	st.DeleteMin()
	fmt.Println(st.Min(), ":", st.Get(st.Min()))
	st.DeleteMin()
	fmt.Println(st.Min(), ":", st.Get(st.Min()))
	st.DeleteMin()
	fmt.Println(st.Min(), ":", st.Get(st.Min()))
}

func TestSymbolTableFloorCeilingRank(t *testing.T) {

	st := NewBST[string, int]()
	st.Put("bob", 2)
	st.Put("alice", 1)
	st.Put("cathy", 3)
	st.Put("david", 4)

	key := "aa"
	fmt.Printf("key:%s, floor: %s, ceiling: %s, rank: %d \n", key, st.Floor(key), st.Ceiling(key), st.Rank(key))
	key = "bb"
	fmt.Printf("key:%s, floor: %s, ceiling: %s, rank: %d \n", key, st.Floor(key), st.Ceiling(key), st.Rank(key))
	key = "cc"
	fmt.Printf("key:%s, floor: %s, ceiling: %s, rank: %d \n", key, st.Floor(key), st.Ceiling(key), st.Rank(key))
	key = "dd"
	fmt.Printf("key:%s, floor: %s, ceiling: %s, rank: %d \n", key, st.Floor(key), st.Ceiling(key), st.Rank(key))

}

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
