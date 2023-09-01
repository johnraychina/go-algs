package basic

import "cmp"

// 用BST表示2-3树
// 利用朝左侧倾斜的”内部“链接作为3分支（2节点）的”胶水“
// 1. represent 2-3 tree as a BST(binary search tree)
// 2. use left-leaning "internal" links as "glue" for 3-nodes
// 个人理解，红黑树本质：加一个颜色字段，将2-3树的复杂操作简化（空间换时间）。

// 为了保证平衡，红黑树做了如下约束：
// 1. 一个节点不能有两个红色链接.
// 2. 每个从root到null的节点的路径上，黑色链接数目相等.
// 3. 红色链接是朝左边倾斜的.
// 1. No node has two red links connected to it.
// 2. Every path from root to null link has the same number of black links.
// 3. Red links lean left.

type RedBlackTree[K cmp.Ordered, V comparable] struct {
	root *RBTreeNode[K, V]
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

// 右旋转 todo
// Left rotation: orient a (temporarily) right-leaning link to lean left.
//           parent
//              |
//             left - right
//             /\      /\
//
