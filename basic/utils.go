package basic

// Helper functions for comparisons and swaps.
func swap(a []int, l int, r int) {
	a[l], a[r] = a[r], a[l]
}

func less(a []int, l int, r int) bool {
	return a[l] < a[r]
}
