package basic

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(a)
	shuffle(a)
	fmt.Println(a)
}
