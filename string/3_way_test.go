package string

import (
	"fmt"
	"testing"
)

func Test3Way(t *testing.T) {

	a := []string{"Eddie", "Tommy", "Clark", "Peter", "McLan"}

	ThreeWaySort(a)
	fmt.Println(a)
}
