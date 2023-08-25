package basic

import (
	"cmp"
)

type SymbolTable[K cmp.Ordered, V any] interface {
	Put(K, V)
	Get(K) V
	Delete(K)
	Contains(K) bool
	IsEmpty()
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
	if b.root == nil {
		b.root = &TreeNode[K, V]{key: k, val: v, left: nil, right: nil}
		return
	}

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
		return &TreeNode[K, V]{key: k, val: v}
	}

	if k < x.key {
		x.left = PutRecursive(x.left, k, v)
	} else if k > x.key {
		x.right = PutRecursive(x.right, k, v)
	} else {
		x.val = v
	}
	return x // 避免在本层调用内部对参数重新赋值无效问题，通过返回值在上层重新赋值。
}

// Delete Hibbard deletion todo
func (b *BinarySearchTree[K, V]) Delete(k K) {

	// 遍历到对应的node，需要冗余记录父节点p.
	var p *TreeNode[K, V]
	x := b.root
	for x != nil {
		if k < x.key {
			p = x
			x = x.left
		} else if k > x.key {
			p = x
			x = x.right
		} else {
			// this is it!
		}
	}

	// 有子节点，需要重组：长兄为父。
	// 对应key的node无子节点，比较好删除
	if x.left != nil && x.right != nil {
		//todo
	}

}

func (b *BinarySearchTree[K, V]) Contains(k K) bool {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) IsEmpty() {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Size() int {
	//TODO implement me
	panic("implement me")
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
		keys = append(keys, node.key)

		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

	// dfs recursive
	//if x == nil {
	//	return
	//}
	//keys = append(keys, x.key)
	//keys = append(keys, b.collectKey(x.left)...)
	//keys = append(keys, b.collectKey(x.right)...)

	return keys
}

func (b *BinarySearchTree[K, V]) Min() K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Max() K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Floor(k K) K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Ceiling(k K) K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Rank(k K) int {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) Select(i int) K {
	//TODO implement me
	panic("implement me")
}

func (b *BinarySearchTree[K, V]) KeysOfRange(lo, hi int) int {
	//TODO implement me
	panic("implement me")
}
