package basic

import "cmp"

// 用BST表示2-3树
// 利用朝左侧倾斜的”内部“链接作为3分支（2节点）的”胶水“
// 1. represent 2-3 tree as a BST(binary search tree)
// 2. use left-leaning "internal" links as "glue" for 3-nodes
// 个人理解，红黑树本质：加一个颜色字段，将2-3树的复杂操作简化（空间换时间）
// 通过颜色变化，实现4-node分解为2个2-node。


// 为了保证平衡，红黑树做了如下约束：
// 1. 一个节点不能有两个红色链接.
// 2. 每个从root到null的节点的路径上，黑色链接数目相等.
// 3. 红色链接是朝左边倾斜的.
// 1. No node has two red links connected to it.
// 2. Every path from root to null link has the same number of black links.
// 3. Red links lean left.


// Q: 如何实现红黑树的约束？
// A: rotate and flip colors 旋转 + 颜色变化

type RedBlackTree[K cmp.Ordered, V comparable] struct {
	root *RBTreeNode[K, V]
}


func NewRedBlackTree[K cmp.Ordered, V comparable]() *RedBlackTree[K, V]{
	return &RedBlackTree[K,V]{}
}

func (t *RedBlackTree[K, V]) Put(k K, v V) {
	t.root = t.PutRecursive(t.root, k, v)
	// 重要的细节，如果没有这个细节，连续插入3个kv就会报错。
	t.root.red = false
}

func (t *RedBlackTree[K, V]) PutRecursive(node *RBTreeNode[K, V], k K, v V) *RBTreeNode[K, V] {
	
	if node == nil { 
		// new node with red link
		return &RBTreeNode[K, V]{key: k, val: v, red: true} 
	}
	
	if k == node.key {
		node.val = v
	} else if k < node.key {
		node.left = t.PutRecursive(node.left, k, v)
	} else if k > node.key {
		node.right = t.PutRecursive(node.right, k, v)
	}

	// 添加到子节点后，再从父节点校验是否需要旋转或翻转颜色（分裂）.
	// need rotate or flip color?
	 // case1: right leaning -> rotate left, to left leaning
	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	// case2: double red left leaning(4-nodes) --> rotate right, to balance 4-nodes
	if isRed(node.left) && isRed(node.left.left) { 
		node = rotateRight(node)
	}
	// case3: left red + right red --> flip colors, to split 4-nodes
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}
	return node
}


// Get 同 BinarySearchTree
func (t *RedBlackTree[K, V]) Get(k K) (v V) {
	x := t.root
	for x != nil {
		if x.key < k {
			x = x.right
		} else if x.key > k {
			x = x.left
		} else {
			return x.val
		}
	}
	return v
}

func isRed[K cmp.Ordered, V comparable](node *RBTreeNode[K, V]) bool {
	if node == nil {
		return false
	}
	return node.red
}



// 左旋转: S的左子树 成为 E的右子树
// Left rotation: orient a (temporarily) right-leaning link to lean left.
//           parent
//              |
//              E --> S
//             /     /\
//
//                 parent
//                   |
//             E <-- S
//             /\     \
//
//  
func rotateLeft[K cmp.Ordered, V any](h *RBTreeNode[K, V]) *RBTreeNode[K, V] {
	
	x := h.right
	
	h.right = x.left
	x.left = h
	
	x.red = h.red
	h.red = true

	return x
}


// 右旋转: E的右子树 变成S的左子树
// Right rotation: orient a left-leaning link to (temporarily) lean right.
//
//                 parent
//                   |
//             E <-- S
//             /\     \
//
//           parent
//              |
//              E --> S
//             /     /\
//
func rotateRight[K cmp.Ordered,V any](h *RBTreeNode[K,V]) *RBTreeNode[K,V] {
	
	x := h.left
	h.left = x.right
	x.right = h
	
	x.red = h.red
	h.red = true

	return x
}

// 翻转颜色：左右link都是红色，说明是一个3node，需要拆分
// Color flip. Recolor to split a (temporary) 4-node.
func flipColors[K cmp.Ordered,V any](h *RBTreeNode[K,V]) {

	if !h.red && h.left.red && h.right.red {
		h.left.red = false
		h.right.red = false
		h.red = true		
	} else {
		panic("illegal color status, can't flip color!")
	}
}
