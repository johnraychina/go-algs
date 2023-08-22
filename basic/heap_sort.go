package basic

// InplaceBottomUpHeapSort 不借助额外的空间，对原数组进行堆排序。
// 要点：
// 1. sink，自底向上构建.
// 2. 最大值交换，sink重排序。
func InplaceBottomUpHeapSort(a []int) {
	N := len(a) - 1
	// first pass: build heap using bottom-up method
	for k := N / 2; k >= 1; k-- {
		sink(a, k, N)
	}

	// second pass:
	// - swap the maximum to the end, one at a time.
	// - Leave in array, instead of nulling out.
	for k := N; k > 1; {
		swap(a, 1, k)
		k--           // 从后往前移动
		sink(a, 1, k) // 重新排序，为下一次swap做准备
	}
}

// MaxPQ(heap)只保证层与层之间的相对顺序，并不保证同一层之间的相对顺序。
// heap sort 利用了MaxPQ(heap)，进一步加工，对数组排序（从小到大）。
func naiveHeapSort(a []int) {

	// 全部塞到 maxPQ中，才能取到最大值
	heap := NewMaxPQ()
	for _, v := range a {
		heap.insert(v)
	}

	// 每次del max 从后往前put到数组，最终数组就是有序的。
	for i := len(a) - 1; i > 0; i-- {
		a[i] = heap.delMax()
	}
}

func sink(a []int, k int, N int) {
	for 2*k <= N {
		i := 2 * k
		if i+1 <= N && less(a, i, i+1) {
			i++
		}
		if !less(a, k, i) {
			break
		}
		swap(a, k, i)
		k = i
	}
}

//* Helper functions for comparisons and swaps.
//* Indices are "off-by-one" to support 1-based indexing.
func swap(a []int, l int, r int) {
	a[l], a[r] = a[r], a[l]
}

func less(a []int, l int, r int) bool {
	return a[l] < a[r]
}
