package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{1, 2, 3, 4, 4, -1, 5, 6, 7, 7}
	//a := []int{3, 5, 4}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	//a := []int{5, 7, 6}
	fmt.Println(a)

	// 靠回忆自己写的，用递归方式做，不但stack深，而且每层stack都要消耗辅助数组，代码丑陋，效率低下
	NaiveRecursiveMergeSort(a)

	// 《算法4》实现:
	// 1.top-down递归方式
	// 2.bottom-up，无递归，采用for循环+一个辅助数组搞定，左右部分对比，分别插入，主数组和辅助数组互换，非常精妙
	BottomUpMergeSort(a)

	fmt.Println(a)

	//println(medianOf3([]int{3, 2, 1}, 0, 1, 2))
}

func BottomUpMergeSort(a []int) {

	aux := make([]int, len(a))
	size := 1
	lo, hi := 0, len(a)-1

	for ; size < len(a); size = size * 2 { // size: 1,2,4,6,8
		for i := lo; i < hi; i = i + size + size { // segment
			mergeSort(a, i, i+size, min(hi, i+size+size), aux)
		}
		a, aux = aux, a // switch array a with auxiliary array aux
	}
}

func mergeSort(a []int, lo, mid, hi int, aux []int) {
	i, j := lo, mid
	for k := lo; k <= hi; k++ {
		if i >= mid {
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
