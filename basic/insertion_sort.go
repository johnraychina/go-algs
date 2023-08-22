package basic

// insertionSort 要点：2层循环，对于一个元素a[i]，把大于它的其他元素交换到右边。
//・In iteration i, swap a[i] with each larger entry to its right.
func insertionSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := i; j >= 1; j-- {
			if less(a, j, j-1) {
				swap(a, j, j-1)
			}
		}
	}
}
