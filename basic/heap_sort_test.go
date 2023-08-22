package basic

import (
	"fmt"
	"testing"
)

func TestInplaceHeapSort(t *testing.T) {
	a := []int{0, 6, 4, 7, 7, 1, 2, 3, 4, 5} // 为了计算父子节点关系方便直接/2, a[0]是占位符。
	//a := []int{0, 1, 2, 3} // 为了计算父子节点关系方便直接/2, a[0]是占位符。

	fmt.Println(a)
	InplaceBottomUpHeapSort(a)
	fmt.Println(a)
}

func TestNaiveHeapSort(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 4, 7, 7}

	fmt.Println(a)
	naiveHeapSort(a)
	fmt.Println(a)
}
