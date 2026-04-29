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
	head     *Node[T] // 哨兵节点，减少空指针判断，不存实际数据。
	p        float64  // level升级概率, 1/2 相当于二叉树，1/4 相当于四叉树
	maxLevel int      // 最大层数，预估: logk(max node), k=1/p
}

type SkipListOption struct {
	MaxLevel int
	P        float64
}
type Option func(*SkipListOption)

func NewSkipList[T any](opts ...Option) *SkipList[T] {
	option := &SkipListOption{
		MaxLevel: 32,
		P:        2,
	}

	for _, opt := range opts {
		opt(option)
	}

	if !(option.MaxLevel > 0 && option.MaxLevel < 1024) {
		panic("max level must be greater than 1024")
	}

	if !(option.P > 0 && option.P < 1) {
		panic("P must be between 0 and 1")
	}

	node := &Node[T]{next: make([]*Node[T], option.MaxLevel+1)}
	return &SkipList[T]{
		maxLevel: option.MaxLevel,
		p:        option.P,
		head:     node,
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
	maxLevel := sl.maxLevel

	// to exclude out of range nodes,
	// move the cur pointer to the largest key that is less or equal to the target
	// top -> down
	for level := maxLevel; level >= 0; level-- {
		// left -> right
		for cur.next[level] != nil && cur.next[level].key < target {
			cur = cur.next[level]
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
	// find each prev node of the key in each level
	cur := sl.head
	maxLevel := sl.maxLevel
	update := make([]*Node[T], maxLevel+1)

	for level := maxLevel; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].key < key {
			fmt.Printf("level: %d, cur: %d, next:%d \n", level, cur.key, cur.next[level].key)
			cur = cur.next[level]
		}

		update[level] = cur
	}

	// 确定层数: 随机数<P，则升一层
	// generate random number for each level
	level := 0
	for rand.Float64() < sl.p && level < maxLevel {
		level++
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
