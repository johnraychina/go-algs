package ch

import (
	"context"
	"strconv"
	"testing"
)

func TestChannelStream(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	str := []string{"do", "rei", "mi", "fa", "so", "la", "ti"}
	log := func(r string) { t.Log(r) }
	even := func(i int) bool { return i%2 == 0 }
	toString := func(i int) string {
		return strconv.Itoa(i) + ": " + str[i-1]
	}

	ctx := context.Background()
	ForEach(ctx, log, Map(ctx, toString, Filter(ctx, even, Stream(ctx, data))))
}
