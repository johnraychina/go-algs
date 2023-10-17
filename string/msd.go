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
	msdSort(a, aux, 0, len(a), 0)
}

func msdSort(a []string, aux []string, lo, hi, d int) {
	panic("implement me")
}
