package stream

import (
	"strconv"
	"strings"
	"testing"
)

func TestSliceStream(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7}
	str := []string{"do", "rei", "mi", "fa", "so", "la", "ti"}
	log := func(r string) { t.Log(r) }
	even := func(i int) bool { return i%2 == 0 }
	toString := func(i int) string {
		return strconv.Itoa(i) + ": " + str[i-1]
	}

	// anonymous function
	// Java : (i) -> { ... }, target type inference do the dirty work.
	// Go: Go could do target-type inference,
	//   but it intentionally doesn’t—because it would conflict with Go’s design goals:
	//   simplicity, predictability, and fast, local compilation.
	Map(MakeIntStream(data).Filter(even), toString).ForEach(log)

	t.Log(strings.Repeat("-", 10))
	// Go native, composition: ForEach( Map( Filter() ) )
	for _, v := range data {
		if even(v) {
			s := toString(v)
			log(s)
		}
	}
}
