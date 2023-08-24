package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {

	//a := []int{5, 7, 6}
	a := []int{1, 2, 3, 4, 4, 5, 6, 7, 7, -1}
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	fmt.Println(a)

	implType := 2
	// 1. 靠回忆自己写的，用递归方式做，不但stack深，而且每层stack都要消耗辅助数组，代码丑陋，效率低下
	// 2.《算法4》top-down递归方式：
	//      三部曲：算mid, sort left&right, merge
	//		一个要点：主数组和辅助数组互换
	// 3. 《算法4》bottom-up，无递归，采用for循环+一个辅助数组搞定，左右部分对比，分别插入，主数组和辅助数组互换，非常精妙
	switch implType {
	case 1:
		NaiveRecursiveMergeSort(a)
	case 2:
		TopDownMergeSort(a)
	case 3:
		BottomUpMergeSort(a)
	}

	fmt.Println(a)

	//println(medianOf3([]int{3, 2, 1}, 0, 1, 2))
}
