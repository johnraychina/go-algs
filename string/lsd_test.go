package string

import (
	"fmt"
	"testing"
)

func TestLSD(t *testing.T) {

	// radix value
	s := "eknry"
	for _, c := range s {
		println(c)
	}

	a := []string{"Eddie", "Tommy", "Clark", "Peter", "McLan"}

	LSDSort(a, 5)
	fmt.Println(a)
}
