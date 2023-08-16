package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{1, 2, 3, 4, 4, 5, 6, 7, 8}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	fmt.Println(a)

	QuickSort(a, 0, len(a)-1)

	fmt.Println(a)
}
func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	// partition
	midIdx := partition(a, lo, hi)
	fmt.Println("partition:", a[lo], "lo", lo, "hi", hi, "midIdx", midIdx)

	// sort left, right
	QuickSort(a, lo, midIdx-1) // 这里必须+1和-1，否则会只剩最后2个元素时，导致无穷递归
	QuickSort(a, midIdx+1, hi)
}

func partition(a []int, lo, hi int) int {
	left, right := lo, hi // 边界条件极易出错，考虑：只有两个元素的场景；重复元素的场景
	for left < right {

		// 向右扫描，直到有一个值比pivot大，或者left 到达最右。
		// todo 为什么这里要用等于？ 如果只有两个元素，要保证左右元素都纳入比较
		for ; a[left] <= a[lo] && left < hi; left++ {
		}

		// todo 想想为什么这里不能用left < right? 这会导致过早停止，对比不充分，导致指针位置不对
		// 向左扫描，直到有一个值比pivot小，或者right 到达最左。
		for ; a[right] > a[lo] && right > lo; right-- {
		}

		// 交换逆序的二者
		if left < right {
			swap(a, left, right)
			fmt.Println(a)
		}
	}

	// 将pivot交换到中间位置
	if a[lo] > a[right] {
		swap(a, lo, right)
		fmt.Println(a)
	}
	return right
}

func swap(a []int, x, y int) {
	a[x], a[y] = a[y], a[x]
}
