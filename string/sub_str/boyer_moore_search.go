package sub_str

//Intuition.
//・Scan characters in pattern from right to left.
//・Can skip as many as M text chars when finding one not in the pattern.
// 排除法思想，如果pattern中某个字符在text一段范围内都没有，无需一个个滑动匹配，直接跳过M个字符。

//Property.
//Substring search with the Boyer-Moore mismatched character heuristic takes about ~ N / M character compares to search for a pattern of
//length M in a text of length N.
//Worst-case. Can be as bad as ~ M N

//Boyer-Moore variant.
//Can improve worst case to ~ 3 N character compares by adding a KMP-like rule
//to guard against repetitive patterns.

// 适用于重复模式不多的情况（容易触发skip）
func BoyerMooreSearch(txt string, pat string) int {

	M := len(pat)
	// build skip table
	right := make([]int, Radix)
	for c := 0; c < Radix; c++ {
		right[c] = -1
	}
	for j := 0; j < M; j++ {
		right[pat[j]] = j
	}

	//todo
	return -1
}
