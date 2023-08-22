package basic

import "math/rand"

// 要点：rand(i + 1)
// 证明： ...* [1/i] * [i/i+1] * [i+1/i+2]....
func shuffle(a []int) {
	for i := 0; i < len(a); i++ {
		r := rand.Intn(i + 1)
		swap(a, i, r)
	}
}
