package string

//Standard quicksort.
//・Uses ~ 2 N ln N string compares on average.
//・Costly for keys with long common prefixes (and this is a common case!)

//3-way string (radix) quicksort.
//・Uses ~ 2 N ln N character compares on average for random strings. ・Avoids re-comparing long common prefixes.

//MSD string sort.
//・Is cache-inefficient.
//・Too much memory storing count[].
//・Too much overhead reinitializing count[] and aux[].

//3-way string quicksort.
//・Has a short inner loop. ・Is cache-friendly.
//・Is in-place.

func ThreeWaySort(a []string) {
	threeWaySort(a, 0, len(a)-1, 0)
}

func threeWaySort(a []string, lo, hi, d int) {
	// 3-way partitioning (using dth character)
	//[...[lo...lt...gt...hi]....]
	if lo >= hi {
		return
	}
	lt, gt := lo, hi
	i := lo + 1
	v := charAt(a[lo], d) // o handle variable-length strings
	for i <= gt {
		t := charAt(a[i], d)
		if t < v {
			swap(a, i, lt)
			i++
			lt++
		} else if t > v {
			swap(a, i, gt)
			gt-- // 注意：这里没有i++，因为a[i]与a[gt]交换后，需要继续与v对比（下个循环）
		} else {
			i++
		}
	}

	// sort 3 part
	threeWaySort(a, lo, lt-1, d)
	if v >= 0 {
		threeWaySort(a, lt, gt, d+1)
	}
	threeWaySort(a, gt+1, hi, d)
}

func swap(a []string, x, y int) {
	a[x], a[y] = a[y], a[x]
}

// LRS longest repeated substring
func LRS(s string) {

}
