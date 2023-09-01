package basic

import (
	"cmp"
)

type SymbolTable[K cmp.Ordered, V any] interface {
	Put(K, V)
	Get(K) V
	Delete(K)
	DeleteMin()
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
	KeysOfRange(lo, hi int) []K // keys in [lo...hi] in sorted order
}

func NewBST[K cmp.Ordered, V comparable]() SymbolTable[K, V] {
	return &BinarySearchTree[K, V]{}
}

type BinarySearchTree[K cmp.Ordered, V comparable] struct {
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

// Delete Hibbard deletion
func (b *BinarySearchTree[K, V]) Delete(k K) {
	if !b.Contains(k) {
		panic("No such a key")
	}
	b.root = DeleteRecursive(b.root, k)
}

// DeleteRecursive 删除树下对应键值的节点，可能当前树的根节点会发生变化，返回变化后的node。
func DeleteRecursive[K cmp.Ordered, V any](node *TreeNode[K, V], k K) *TreeNode[K, V] {
	if node == nil {
		return nil
	}

	if k > node.key { // 搜索key的过程本质上和Get是一样的
		node.right = DeleteRecursive(node.right, k)
	} else if k < node.key {
		node.left = DeleteRecursive(node.left, k)
	} else { // 找到对应键值的节点
		// case 0: 无子树
		if node.left == nil && node.right == nil {
			return nil
		}
		// case 1: 有一个子树
		if node.left == nil && node.right != nil {
			return node.right
		}
		if node.left != nil && node.right == nil {
			return node.left
		}

		// case 2: 有两个子树，选取min(node.right) 或者max(node.left)作为当前node的替代节点。
		// 巧妙的是，由于树是中序inorder，这样选取的继任节点是不可能有子树的，直接使用即可。
		successor := MinNode(node.right)
		// 将 successor 提升为继任者
		successor.right = DeleteMinOf(node.right)
		successor.left = node.left
		return successor
	}

	return node
}

func MinNode[K cmp.Ordered, V any](node *TreeNode[K, V]) *TreeNode[K, V] {
	if node.left != nil {
		return MinNode(node.left)
	}
	return node
}

func (b *BinarySearchTree[K, V]) DeleteMin() {
	if b.root == nil {
		panic("empty tree")
	}
	b.root = DeleteMinOf(b.root)
}

func DeleteMinOf[K cmp.Ordered, V any](node *TreeNode[K, V]) *TreeNode[K, V] {

	// 左子树为空，则当前节点为最小节点，删除他，然后返回右子树作为替代。
	if node.left == nil {
		return node.right
	}

	// 左子树不为空，最小节点一定在他下面，删除之，更新指向左子树的链接。
	node.left = DeleteMinOf(node.left)
	node.count = 1 + SizeOf(node.left) + SizeOf(node.right)
	return node
}

func (b *BinarySearchTree[K, V]) Contains(k K) bool {
	var zeroValue V
	return b.Get(k) != zeroValue
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

func (b *BinarySearchTree[K, V]) KeysOfRange(lo, hi int) (keys []K) {
	// in-order traversal the tree, and collect the keys in the range [lo, hi]
	
	x := b.root
	if x == nil {
		return keys
	}

	lowKey := b.Select(lo)
	highKey := b.Select(hi)
	
	keys = append(keys, b.KeysOfRangeRecursive(x, lowKey, highKey)...)
	return keys
}

func (b *BinarySearchTree[K,V]) KeysOfRangeRecursive(node *TreeNode[K,V], lowKey, hiKey K) []K {
	var keys []K
	if node == nil {
		return keys
	}

	// dfs
	if lowKey < node.key {
		keys = append(keys, b.KeysOfRangeRecursive(node.left, lowKey, hiKey)...)
	}
	if lowKey <= node.key && node.key <= hiKey {
		keys = append(keys, node.key)
	}
	if hiKey > node.key {
		keys = append(keys, b.KeysOfRangeRecursive(node.right, lowKey, hiKey)...)
	}

	return keys
}

func (b *BinarySearchTree[K, V]) Keys() (keys []K) {
	return b.collectKey(b.root)
}

func (b *BinarySearchTree[K, V]) collectKey(node *TreeNode[K, V]) []K {

	var keys []K
	if node == nil {
		return keys
	}

	// bfs
	queue := NewLinkedQueue[*TreeNode[K, V]]()
	queue.Enqueue(node)
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
	if x == nil {
		panic("empty tree")
	}

	// the most left one
	for x.left != nil {
		x = x.left
	}
	return x.key
}

func (b *BinarySearchTree[K, V]) Max() K {
	x := b.root
	if x == nil {
		panic("empty tree")
	}

	// the largest key is the most left one
	for x.right != nil {
		x = x.right
	}
	return x.key
}

func (b *BinarySearchTree[K, V]) Floor(k K) (result K) {

	x := b.root
	if x == nil {
		panic("empty tree")
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
	if node == nil {
		return nil
	}
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
		panic("empty tree")
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
	if node == nil {
		return nil
	}

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
	if b.root == nil {
		panic("empty tree")
	}
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
	
	// select the i-th smallest key
	if b.root == nil {
		panic("empty tree")
	}

	if i > b.Size() || i < 0 {
		panic("rank out of range")
	}
	
	return selectOf(b.root, i).key
}

func selectOf[K cmp.Ordered, V any](node *TreeNode[K, V], i int) *TreeNode[K, V] {
	if node == nil {
		return nil
	}

	t := SizeOf(node.left)
	if t > i {
		return selectOf(node.left, i)
	} else if t < i {
		return selectOf(node.right, i-t-1)
	} else {
		return node
	}
}
