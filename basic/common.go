package basic

// Helper functions for comparisons and swaps.
func swap(a []int, l int, r int) {
	a[l], a[r] = a[r], a[l]
}

func less(a []int, l int, r int) bool {
	return a[l] < a[r]
}

type RBTreeNode[K comparable, V any] struct {
	key   K
	val   V
	left  *RBTreeNode[K, V]
	right *RBTreeNode[K, V]
	red   bool // color of parent link
}

type TreeNode[K comparable, V any] struct {
	key   K
	val   V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
	count int
}

type Node[V any] struct {
	val  V
	prev *Node[V]
	next *Node[V]
}
