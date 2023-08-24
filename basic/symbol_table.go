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
	for x != nil { // 非递归写法，还有一种是递归传TreeNode的写法
		if k < x.key {
			if x.left == nil {
				x.left = &TreeNode[K, V]{key: k, val: v}
				return
			}
			x = x.left
		} else if k > x.key {
			if x.right == nil {
				x.right = &TreeNode[K, V]{key: k, val: v}
				return
			}
			x = x.right
		} else {
			x.val = v // equal key, override with the value
			return
		}
	}
}

func (b *BinarySearchTree[K, V]) Delete(k K) {
	//TODO implement me
	panic("implement me")
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

	// bfs todo
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
