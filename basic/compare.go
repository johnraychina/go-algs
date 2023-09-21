package basic

import "cmp"

// ComparableObj 由于golang的泛型是用typeset实现的，得预先指定实现范围，没有java那样灵活，
// 为了利用comparable，这里用接口包一层，获取key，只要key是 Ordered 即可。
type ComparableObj[K cmp.Ordered] interface {
	Key() K
}

type Int struct {
	ComparableObj[int]
	k int
}

func (i *Int) Key() int {
	return i.k
}

type Float32 struct {
	ComparableObj[float32]
	k float32
}

func (f *Float32) Key() float32 {
	return f.k
}
