package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 4, 7, 7}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	//a := []int{5, 7, 6}
	fmt.Println(a)

	QuickSort(a, 0, len(a)-1)

	fmt.Println(a)
}

const CUTOFF = 10

func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	// 优化：insertion sort
	if hi <= lo+CUTOFF-1 {
		InsertionSort(a, lo, hi)
		return
	}

	// 优化: 使用中间元素
	// ~ 12/7 n ln n compares (14% fewer)
	// ~ 12/35 n ln n exchanges (3% more)
	median := medianOf3(a, lo, lo+(hi-lo)/2, hi)
	swap(a, lo, median)

	// partition
	midIdx := partition(a, lo, hi)

	// sort left, right
	QuickSort(a, lo, midIdx-1) // 这里必须+1和-1，否则会只剩最后2个元素时，导致无穷递归
	QuickSort(a, midIdx+1, hi)
}

func InsertionSort(a []int, lo int, hi int) {
	if lo == hi {
		return
	}

	for i := lo; i <= hi; i++ {
		// find smallest of the rest
		minIdx := i
		for j := i + 1; j <= hi; j++ {
			if less(a[j], a[minIdx]) {
				minIdx = j
			}
		}

		// swap the min to i position
		swap(a, i, minIdx)
	}
}

func less(a, b int) bool {
	return a < b
}
func medianOf3(a []int, lo int, i int, hi int) int {
	if a[lo] < a[hi] {
		if a[i] < a[lo] {
			return lo
		} else if a[i] > a[hi] {
			return hi
		} else {
			return i
		}
	} else {
		if a[i] > a[lo] {
			return lo
		} else if a[i] < a[hi] {
			return hi
		} else {
			return i
		}
	}
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
