package string

// Trie performance
// Search hit: Need to examine all L characters for equality
// Search miss:
//   - Could have mismatched on the first character.
//   - Typical case: only examine a few characters(sublinear).
//Space: R null links at each leaf. (but sublinear space possible if many short strings share common prefixes)

// 精髓：
// 1. 将搜索转换为查表（c -> next[c] 数组索引）.
// 2. 共享前缀，节省空间.
// 这两者都是基于合理的假设：ascii字符数较少(256), 英文字母很多共享前缀.
// Make good assumptions about the problems.

type TrieST[V comparable] struct {
	R    int // radix
	root *Node[V]
}

type Node[V comparable] struct {
	value V
	next  []*Node[V]
}

func NewTrieST[V comparable](radix int) *TrieST[V] {
	trie := &TrieST[V]{}
	trie.R = radix

	return trie
}

func (t *TrieST[V]) Put(key string, val V) {
	t.root = t.put(t.root, key, val, 0)
}

// put value into trie with key, starting from depth d.
func (t *TrieST[V]) put(node *Node[V], key string, val V, d int) *Node[V] {

	if node == nil {
		node = &Node[V]{next: make([]*Node[V], t.R)}
	}

	// end condition
	if len(key) == d {
		node.value = val
		return node
	}

	// recursive call
	// NOTICE: node.next[c] maybe is nil
	c := key[d]
	node.next[c] = t.put(node.next[c], key, val, d+1)
	return node
}

func (t *TrieST[V]) Get(key string) (v V) {
	node := t.get(t.root, key, 0)
	if node == nil {
		return v
	}
	return node.value
}

func (t *TrieST[V]) get(node *Node[V], key string, d int) *Node[V] {
	// end condition
	if node == nil {
		return nil
	}

	// recursive call
	if len(key) != d {
		return t.get(node.next[key[d]], key, d+1)
	}
	// hit key
	return node
}

func (t *TrieST[V]) Delete(key string) {
	t.root = t.delete(t.root, key, 0)
}

// To delete a key-value pair:
// ・Find the node corresponding to key and set value to null.
// ・If node has null value and all null links, remove that node (and recur).

func (t *TrieST[V]) delete(node *Node[V], key string, d int) *Node[V] {
	var zeroV V

	if node == nil {
		return nil // end condition
	}
	if len(key) == d {
		// hit the key
		node.value = zeroV // set value to null
	} else {
		c := key[d]
		node.next[c] = t.delete(node.next[c], key, d+1) // recursive call
	}

	// remove subtrie rooted at x if it is completely empty
	if node.value == zeroV {
		for _, n := range node.next {
			if n != nil {
				return node // has non-null links
			}
		}
		return nil // all null links, remove the node
	}
	return node
}
