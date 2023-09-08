package basic

import "hash/fnv"

type Hash interface {
	comparable
	HashCode() uint
}

type HashTable[K Hash, V comparable] struct {
	size uint
	tab  []*HashNode[K, V] // table of nodes indexed by hash
}

func NewHashTable[K Hash, V comparable](capacity uint) *HashTable[K, V] {
	if capacity <= 0 {
		panic("capacity <= 0")
	}

	c := NearestPowerOfTwo(capacity)

	return &HashTable[K, V]{
		size: 0,
		tab:  make([]*HashNode[K, V], c),
	}
}

func (t *HashTable[K, V]) Get(k K) (v V) {
	hc := k.HashCode()
	m := uint(len(t.tab))
	idx := hc % m

	node := t.tab[idx]
	for node != nil {
		if node.key == k {
			return node.val
		}
		node = node.next
	}
	return v
}

func (t *HashTable[K, V]) Put(k K, v V) {
	hc := k.HashCode()
	m := uint(len(t.tab))
	idx := hc % m

	node := t.tab[idx]
	if node == nil {
		t.tab[idx] = &HashNode[K, V]{key: k, val: v}
		t.size++
		return
	}

	for node != nil {
		if node.key == k { // found
			node.val = v
			return
		} else if node.next != nil { // keep searching
			node = node.next
		} else { // not found
			node.next = &HashNode[K, V]{key: k, val: v}
			t.size++
			return
		}
	}
}

func (t *HashTable[K, V]) Delete(k K) {
	hc := k.HashCode()
	m := uint(len(t.tab))
	idx := hc % m

	node := t.tab[idx]
	if node == nil {
		return
	}
	var zeroV V
	if t.Get(k) != zeroV {
		t.tab[idx] = listDelete(node, k)
		t.size--
	}
}

func listDelete[K Hash, V any](node *HashNode[K, V], k K) *HashNode[K, V] {
	if node == nil {
		return nil
	}
	if node.key == k {
		return node.next
	}

	node.next = listDelete(node.next, k)
	return node
}

func (t *HashTable[K, V]) Size() uint {
	return t.size
}

func (t *HashTable[K, V]) ensureCapacity() {
	panic("todo")
}

func (t *HashTable[K, V]) Keys() (keys []K) {

	for _, node := range t.tab {
		for node != nil {
			keys = append(keys, node.key)
			node = node.next
		}
	}
	return keys
}

type String string

func (s String) HashCode() uint {
	f := fnv.New32()
	f.Write([]byte(s))
	return uint(f.Sum32()) // fixme
}

type Integer struct {
	v int
}

func (i Integer) HashCode() uint {
	//return uint(i.v)
	return uint(i.v & 0x7fffffff)
}
