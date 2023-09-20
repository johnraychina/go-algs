package basic

// ComparableKey golang的泛型是用typeset实现的，得预先指定实现范围，没有java那样灵活，
// 为了利用comparable，这里用接口包一层，获取key，只要key是comparable即可。
type ComparableKey[K comparable] interface {
	Key() K
}

type Int struct {
	ComparableKey[int]
	k int
}

func (i Int) Key() int {
	return i.k
}

type Float32 struct {
	ComparableKey[float32]
	k float32
}

func (f *Float32) Key() float32 {
	return f.k
}
