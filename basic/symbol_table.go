package basic

import (
	"cmp"
)

type SymbolTable[K cmp.Ordered, V any] interface {
	Put(K, V)
	Get(K) V
	Delete(K)
	Contains(K) bool
	IsEmpty() bool
	Size() int
	Keys() []K

	Min() K
	Max() K
	Floor(K) K                  // largest key less than or equal to the key
	Ceiling(K) K                // smallest key greater than or equal to the key
	Rank(K) int                 // number of keys less than the key
	Select(int) K               // key of rank
	KeysOfRange(lo, hi int) int // keys in [lo...hi] in sorted order
}

func NewBST[K cmp.Ordered, V any]() SymbolTable[K, V] {
	return &BinarySearchTree[K, V]{}
}

type BinarySearchTree[K cmp.Ordered, V any] struct {
	root *TreeNode[K, V]
}

func (b *BinarySearchTree[K, V]) Get(k K) (v V) {
	x := b.root
	for x != nil {
		if k < x.key {
			x = x.left
		} else if k > x.key {
			x = x.right
		} else {
			return x.val
		}
	}
	return v // no equal key, return zero value
}

func (b *BinarySearchTree[K, V]) Put(k K, v V) {
	// traverse through the tree
	x := b.root

	// 递归写法
	b.root = PutRecursive(x, k, v)

	// 非递归写法
	//for x != nil {
	//	if k < x.key {
	//		if x.left == nil {
	//			x.left = &TreeNode[K, V]{key: k, val: v}
	//			return
	//		}
	//		x = x.left
	//	} else if k > x.key {
	//		if x.right == nil {
	//			x.right = &TreeNode[K, V]{key: k, val: v}
	//			return
	//		}
	//		x = x.right
	//	} else {
	//		x.val = v // equal key, override with the value
	//		return
	//	}
	//}
}

func PutRecursive[K cmp.Ordered, V any](x *TreeNode[K, V], k K, v V) *TreeNode[K, V] {
	if x == nil {
		return &TreeNode[K, V]{key: k, val: v, count: 1}
	}

	if k < x.key {
		x.left = PutRecursive(x.left, k, v)
	} else if k > x.key {
		x.right = PutRecursive(x.right, k, v)
	} else {
		x.val = v
	}
	x.count = 1 + SizeOf(x.left) + SizeOf(x.right)
	return x // 避免在本层调用内部对参数重新赋值无效问题，通过返回值在上层重新赋值。
}

// Delete Hibbard deletion todo
func (b *BinarySearchTree[K, V]) Delete(k K) {

}

func (b *BinarySearchTree[K, V]) DeleteRecursive(node *TreeNode[K, V], k K) {

}

func (b *BinarySearchTree[K, V]) DeleteMin(node *TreeNode[K, V]) *TreeNode[K, V] {

	// 左子树为空，则当前节点为最小节点，删除他，然后返回右子树作为替代。
	if node.left == nil {
		return node.right
	}

	// 左子树不为空，最小节点一定在他下面，删除之，更新指向左子树的链接。
	node.left = b.DeleteMin[K, V](node.left)
	return node
}

func (b *BinarySearchTree[K, V]) Contains(k K) bool {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) IsEmpty() bool {
	return b.Size() > 0
}

func (b *BinarySearchTree[K, V]) Size() int {
	return SizeOf(b.root)
}

func SizeOf[K cmp.Ordered, V any](x *TreeNode[K, V]) int {
	if x == nil {
		return 0
	}
	return x.count
}

func (b *BinarySearchTree[K, V]) Keys() (keys []K) {
	return b.collectKey(b.root)
}

func (b *BinarySearchTree[K, V]) collectKey(root *TreeNode[K, V]) []K {

	var keys []K

	// bfs
	queue := NewLinkedQueue[*TreeNode[K, V]]()
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		node := queue.Dequeue()

		if node.left != nil {
			queue.Enqueue(node.left)
		}
		keys = append(keys, node.key)
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

	// dfs recursive
	//if x == nil {
	//	return
	//}
	//keys = append(keys, b.collectKey(x.left)...)
	//keys = append(keys, x.key)
	//keys = append(keys, b.collectKey(x.right)...)

	return keys
}

func (b *BinarySearchTree[K, V]) Min() K {
	x := b.root

	// the most left one
	for x.right != nil {
		x = x.right
	}
	return x.key
}

func (b *BinarySearchTree[K, V]) Max() K {
	x := b.root

	// the largest key is the most left one
	for x.left != nil {
		x = x.left
	}
	return x.key
}

func (b *BinarySearchTree[K, V]) Floor(k K) (result K) {

	x := b.root
	if x == nil {
		return result
	}

	x = b.FloorRecursive(x, k)
	if x == nil {
		return result
	}
	return x.key
}

// FloorRecursive find the largest key <= the given key
// 1. node key > the given key, the floor must in the left subtree.
// 2. node key < the given key, the floor could be in the right subtree(or the node key)
// 想要找到左边界：
// 1. 当前节点值较小，一定是向左子树找。
// 2. 相等，直接return
// 3. 当前节点较小，可能就是当前节点或者右子树。
func (b *BinarySearchTree[K, V]) FloorRecursive(node *TreeNode[K, V], k K) *TreeNode[K, V] {
	// largest key <= the given key
	// 递归实现
	if k < node.key {
		return b.FloorRecursive(node.left, k)
	} else if k == node.key {
		return node
	} else {
		t := b.FloorRecursive(node.right, k)
		if t != nil {
			return t
		}
		return node
	}
}

// Ceiling Smallest key ≥ a given key.
func (b *BinarySearchTree[K, V]) Ceiling(k K) (result K) {
	x := b.root
	if x == nil {
		return result
	}

	x = b.CeilingRecursive(x, k)
	if x == nil {
		return result
	}
	return x.key
}

// CeilingRecursive
// 想要找右边界（比给定key值大的最小值）
// 1. 当前node节点较小，一定找右子树。
// 2. 相等，直接return
// 3. 当前node节点较大，可能就是当前节点，或者左子树。
func (b *BinarySearchTree[K, V]) CeilingRecursive(node *TreeNode[K, V], k K) *TreeNode[K, V] {

	if k > node.key {
		return b.CeilingRecursive(node.right, k)
	} else if k == node.key {
		return node
	} else {
		t := b.CeilingRecursive(node.left, k)
		if t != nil {
			return t
		}
		return node
	}
}

func (b *BinarySearchTree[K, V]) Rank(k K) int {
	return RankOf(b.root, k)
}

func RankOf[K cmp.Ordered, V any](node *TreeNode[K, V], k K) int {
	if node == nil {
		return 0
	}

	if k < node.key {
		return RankOf(node.left, k)
	} else if k == node.key {
		return SizeOf(node)
	} else {
		return SizeOf(node.left) + 1 + RankOf(node.right, k)
	}
}

func (b *BinarySearchTree[K, V]) Select(i int) K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) KeysOfRange(lo, hi int) int {
	//TODO implement me
	panic("implement me")
}
