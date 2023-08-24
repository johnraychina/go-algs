package basic

import (
	"fmt"
)

func TopDownMergeSort(a []int) {
	aux := make([]int, len(a))
	topDownMergeSort(a, 0, len(a)-1, aux)
}

func topDownMergeSort(a []int, lo, hi int, aux []int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	//fmt.Println(lo, hi, mid, a, aux)

	topDownMergeSort(a, lo, mid, aux)
	topDownMergeSort(a, mid+1, hi, aux)
	merge(a, lo, mid, hi, aux)
}

func BottomUpMergeSort(a []int) {

	aux := make([]int, len(a))
	size := 1
	lo, hi := 0, len(a)-1

	for ; size < len(a); size = size * 2 { // size: 1,2,4,6,8
		for i := lo; i < hi; i = i + size + size { // segment
			merge(a, i, i+size, min(hi, i+size+size), aux)
		}
		a, aux = aux, a // switch array a with auxiliary array aux
	}
}

func merge(a []int, lo, mid, hi int, aux []int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			aux[k] = a[j]
			j++
		} else if j > hi {
			aux[k] = a[i]
			i++
		} else if a[i] < a[j] {
			aux[k] = a[i]
			i++
		} else {
			aux[k] = a[j]
			j++
		}
	}

	for i := lo; i <= hi; i++ {
		a[i] = aux[i]
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func NaiveRecursiveMergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	// mid: (lo + hi) / 2
	mid := len(a) / 2
	x := 0

	//left part and right part
	left := make([]int, mid)
	for i := 0; i < len(left); i++ {
		left[i] = a[x]
		x++
	}

	right := make([]int, len(a)-mid)
	for j := 0; j < len(right); j++ {
		right[j] = a[x]
		x++
	}

	NaiveRecursiveMergeSort(left)
	NaiveRecursiveMergeSort(right)

	// merge two sorted partition into one
	i, l, r := 0, 0, 0
	for i < len(a) {
		if l < len(left) && r < len(right) {
			if left[l] < right[r] {
				a[i] = left[l]
				l++
				i++
			} else {
				a[i] = right[r]
				r++
				i++
			}
		} else if l == len(left) {
			for r < len(right) {
				a[i] = right[r]
				r++
				i++
			}
		} else if r == len(right) {
			for l < len(left) {
				a[i] = left[l]
				l++
				i++
			}
		}

	}
	fmt.Println("a:", a)
}
