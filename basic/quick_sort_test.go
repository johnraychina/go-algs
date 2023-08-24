package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6, 4, 7, 7}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	//a := []int{5, 7, 6}
	fmt.Println(a)
	//QuickSort(a, 0, len(a)-1)
	//Quick3Way(a, 0, len(a)-1)
	//fmt.Println("top 3", TopK(a, 3))
	fmt.Println(a)

	//println(medianOf3([]int{3, 2, 1}, 0, 1, 2))
}
