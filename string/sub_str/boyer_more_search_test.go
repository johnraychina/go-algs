package sub_str

import (
	"fmt"
	"testing"
)

func TestBoyerMoreSearch(t *testing.T) {

	idx := BoyerMooreSearch("FIND AN NEEDLE IN HEY", "NEEDLE")
	fmt.Println(idx)
}
