package main

import (
	"fmt"
)

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 4, 7, 7}
	//rand.Shuffle(len(a), func(i, j int) {
	//	a[i], a[j] = a[j], a[i]
	//})
	a := []int{5, 7, 6}
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

	// sort left, right
	QuickSort(a, lo, midIdx-1) // 这里必须+1和-1，否则会只剩最后2个元素时，导致无穷递归
	QuickSort(a, midIdx+1, hi)
}

func partition(a []int, lo, hi int) int {
	fmt.Printf("before partition: %d\t %d\t %v\n", lo, hi, a)

	left, right := lo+1, hi // 边界条件极易出错，考虑：只有两个元素的场景；重复元素的场景
	for left < right {
		// 注意：为什么是到达最左和最右呢？ 为什么不能用left < right?
		// 例如[5 7 6]，right=(1)，以7为为中心，分为[5], 7, [6]两个partition，没做swap.
		// 如果用left < hi, right > lo作为界限，right=(0)，分为5, [7, 6]，后续会做swap.

		// 向右扫描，直到有一个值比pivot大，或者left 到达最右。
		for ; a[left] < a[lo] && left < hi; left++ {
			println("left = ", left, " right=", right)
		}

		// 向左扫描，直到有一个值比pivot小，或者right 到达最左。
		for ; a[right] > a[lo] && right > lo; right-- {
			println("right1 = ", right, " left=", left)
		}

		println("left: ", left, " right: ", right)

		// 交换逆序的二者
		if left < right {
			swap(a, left, right)
			fmt.Println("l-r swap", a)
		}
	}

	// 将pivot交换到中间位置
	if a[lo] > a[right] {
		swap(a, lo, right)
		fmt.Println("pivot swap", a)
	}

	fmt.Printf("after  partition: %d\t %d\t %v\n", lo, hi, a)
	return right
}

func swap(a []int, x, y int) {
	a[x], a[y] = a[y], a[x]
}
