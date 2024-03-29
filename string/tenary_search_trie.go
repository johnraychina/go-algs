package string

// ・Store characters and values in nodes (not keys).
// ・Each node has 3 children: smaller (left), equal (middle), larger (right).
// 和常规的Trie树不同，为了减少null link所占空间，使用了left, middle, right 3个子节点.

type TNode[V comparable] struct {
	value            V
	c                uint8
	left, mid, right *TNode[V]
}

type TrieTST[V comparable] struct {
	root *TNode[V]
}

func NewTrieTST[V comparable]() *TrieTST[V] {
	trie := &TrieTST[V]{}
	return trie
}

func (t *TrieTST[V]) Put(key string, v V) {
	t.root = t.put(t.root, key, v, 0)
}

func (t *TrieTST[V]) put(node *TNode[V], key string, v V, d int) *TNode[V] {
	if node == nil {
		node = &TNode[V]{c: key[d]}
	}

	c := key[d]

	if c < node.c {
		node.left = t.put(node.left, key, v, d)
	} else if c > node.c {
		node.right = t.put(node.right, key, v, d)
	} else if d < len(key)-1 {
		node.mid = t.put(node.mid, key, v, d+1)
	} else {
		node.value = v // match c, down to the next
	}

	return node
}

func (t *TrieTST[V]) Get(key string) (v V) {
	node := t.get(t.root, key, 0)
	if node != nil {
		return node.value
	}
	return v
}

func (t *TrieTST[V]) get(node *TNode[V], key string, d int) *TNode[V] {
	if node == nil {
		return nil
	}

	c := key[d]
	if c < node.c {
		return t.get(node.left, key, d)
	} else if c > node.c {
		return t.get(node.right, key, d)
	} else if d < len(key)-1 {
		return t.get(node.mid, key, d+1)
	} else {
		return node
	}
}
