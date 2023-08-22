package basic

// selectionSort 要点：每次选择最小放左边，逐步向右处理。
// In each iteration,find the index of the smallest remaining entry.
// swap a[i] with a[minIdx]
//[...sorted | unsorted....]
func selectionSort(a []int) {
	if len(a) <= 1 {
		return
	}
	for i := 0; i < len(a); i++ {
		minIdx := i
		for j := i + 1; j < len(a); j++ {
			if less(a, j, minIdx) {
				minIdx = j
			}
		}
		swap(a, i, minIdx)
	}

	return
}
