package graph

import (
	"fmt"
	"math/rand"
)

// skip list
//Level 3:        1 ----------- 9
//Level 2:        1 ----- 5 ----- 9
//Level 1:        1 - 3 - 5 - 7 - 9
//Level 0:        1 - 2 - 3 - 4 - 5 - 6 - 7 - 8 - 9

// n: all node count
// k: like tree branches, how many child nodes under each parent node in average(statistically).
// P=1/k: node promotion possibility (1/2, 1/4 in general)
// MaxLevel: about logk(n), grow up as nodes getting more

type Node[T any] struct {
	key   int
	value T

	//next[i]: forward pointer of the i-th level
	//len(next): the height of the current node, random(P).
	next []*Node[T]
}

type SkipList[T any] struct {
	head  *Node[T] // 哨兵节点，减少空指针判断，不存实际数据。
	p     float64  // level升级概率, 1/2 相当于二叉树，1/4 相当于四叉树
	level int      // 当前最大层数，预估: logk(max node), k=1/p
}

const MaxLevel = 31 // 最大层, zero based

type SkipListOption struct {
	P float64
}
type Option func(*SkipListOption)

func NewSkipList[T any](opts ...Option) *SkipList[T] {
	option := &SkipListOption{
		P: 0.5,
	}

	for _, opt := range opts {
		opt(option)
	}

	if !(option.P > 0 && option.P < 1) {
		panic("P must be between 0 and 1")
	}

	node := &Node[T]{next: make([]*Node[T], MaxLevel+1)} // zero based
	return &SkipList[T]{
		level: 0,
		p:     option.P,
		head:  node,
	}
}

// skip list
//Level 3:        1 ------------------------------ 9
//Level 2:        1 ------------- 5 -------------- 9
//Level 1:        1   -   3   -   5   -   7   -   9
//Level 0:        1 - 2 - 3 - 4 - 5 - 6 - 7 - 8 - 9

func (sl *SkipList[T]) Search(target int) (ok bool, value T) {
	// 永远通过“前驱”来操作结构，而不是直接命中节点
	cur := sl.head
	level := sl.level

	// to exclude out of range nodes,
	// move the cur pointer to the largest key that is less or equal to the target
	// top -> down
	for l := level; l >= 0; l-- {
		// left -> right
		for cur.next[l] != nil && cur.next[l].key < target {
			cur = cur.next[l]
		}
	}

	// cur.key < target
	// cur.next.key >= target
	cur = cur.next[0]
	if cur != nil && cur.key == target {
		return true, cur.value
	}

	return false, value // not found ,return zero value
}

func (sl *SkipList[T]) Insert(key int, value T) {

	// 找到每层前驱节点
	// find each prev node of the key in each l
	cur := sl.head
	curMaxLevel := sl.level
	update := make([]*Node[T], MaxLevel+1)

	for l := curMaxLevel; l >= 0; l-- {
		for cur.next[l] != nil && cur.next[l].key < key {
			cur = cur.next[l]
		}

		update[l] = cur
	}

	// 已存在，则覆盖 fixed
	// override the value if the key already exists
	if cur.next[0] != nil && cur.next[0].key == key {
		cur.next[0].value = value
	}

	// 确定新节点的层数: 随机数<P，则升一层
	// generate random number for each level of the new node
	// promote level if less than p
	level := 0
	for rand.Float64() < sl.p && level < MaxLevel {
		level++
	}

	// 新节点层数超过当前最大层，则扩层，用head作为前驱节点
	if level > sl.level {
		for i := sl.level + 1; i <= level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 新节点
	newNode := &Node[T]{
		key:   key,
		value: value,
		next:  make([]*Node[T], level+1),
	}

	// 对应层插入节点
	// prev ------------ newNode ------------ prev.next
	//            prev - newNode - prev.next
	for i := 0; i <= level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

// Print show skip list nodes in a horizontal way
func (sl *SkipList[T]) Print() {

	// iterate level 0 node
	for cur := sl.head; cur != nil; cur = cur.next[0] {

		// iterate each level
		fmt.Printf("[%d]", cur.key)
		for i := range cur.next {
			if cur.next[i] == nil {
				fmt.Printf("%s ", "-")
			} else {
				fmt.Printf("%d ", cur.next[i].key)
			}
		}
		fmt.Println()
	}
}

func (sl *SkipList[T]) Range(from, to int) []T {
	//todo
	var r []T
	return r
}
