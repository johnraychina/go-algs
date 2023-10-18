package string

import (
	"fmt"
	"testing"
)

func TestMSD(t *testing.T) {
	a := []string{"Zack", "Eddie", "Tommy", "Clark", "Peter", "McLan"}
	MSDSort(a)
	fmt.Println(a)
}
