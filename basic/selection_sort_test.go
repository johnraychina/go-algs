package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 4, 7, 7}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	fmt.Println(a)
	selectionSort(a)
	fmt.Println(a)
}
