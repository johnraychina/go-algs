package string

// MSDSort Most-significant-digit-first string sort
// - Partition array into R pieces according to first character (use key-indexed counting).
// - Recursively sort all strings that start with each character (key-indexed counts delineate subarrays to sort).
// For variable length strings:
// treat strings as if they had an extra char at end (smaller than any char).
// 技术要点：
// 1. 按基数分片为R个子数组
// 2. 递归对子数组排序
func MSDSort(a []string) {
	aux := make([]string, len(a))
	msdSort(a, aux, 0, len(a)-1, 0)
}

const R = 256 // radix

func msdSort(a []string, aux []string, lo, hi, d int) {
	if hi <= lo {
		return
	}
	count := make([]int, R+2) // 注意这里R+2, count[0]对应r=-1， count[1]对应r=0(预留)
	for i := lo; i <= hi; i++ {
		s := a[i]
		count[charAt(s, d)+2]++ // 由于charAt可能返回-1，所以需要+2，从保证从count[1]开始。
	}

	// 累计count
	for r := 0; r < R; r++ {
		count[r+1] += count[r]
	}

	// copy a to aux based on count
	for i := lo; i <= hi; i++ {
		radix := charAt(a[i], d) + 1
		aux[count[radix]] = a[i]
		count[radix]++
	}
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	for r := 0; r < R; r++ {
		msdSort(a, aux, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

// Treat strings as if they had an extra char at end (smaller than any char).
func charAt(s string, i int) int {
	if i < len(s) {
		return int(s[i])
	}
	return -1
}
