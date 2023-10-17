package string

//Least-significant-digit-first string (radix) sort

//・Consider characters from right to left.
//・Stably sort using dth character as the key (using key-indexed counting).

// LSDSort
// 基数排序技术要点
// 1. 固定列数W，有限基数范围[0 ~ 256]
// 2. count基数计数表 + aux 辅助
// 3. 从右到左逐列处理
func LSDSort(a []string, W int) {

	R := 256 // ascii radix
	N := len(a)
	aux := make([]string, N)

	// sort by the d-th character
	// do key-indexed counting
	// for each digit from right to left
	for d := W - 1; d >= 0; d-- {
		// 注意：R radix如果基数太大，内存分配释放比较消耗资源
		count := make([]int, R+1)

		for _, s := range a {
			count[s[d]+1]++ // 这里要注意
		}

		// key-indexed counting
		// 这里叠罗汉的方式设定边界: 已经在做基数r的相对排序了。
		// 4 +----+
		// 6 +------+
		// 6 +------+
		// 变为：4, 10, 16 +----+---------+---------+
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		// sort by copying a[...] to aux[...] based on count
		for _, s := range a {
			radix := s[d]
			aux[count[radix]] = s
			count[radix]++
		}

		// copy back
		for i := 0; i < N; i++ {
			a[i] = aux[i]
		}
	}
}
