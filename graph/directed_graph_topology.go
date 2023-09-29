package graph

import "go-algs/basic"

// DepthFirstOrder reverse DFS post order
type DepthFirstOrder struct {
	marked      []bool
	reversePost *basic.Stack[int]
}

func NewDepthFirstOrder(g Graph) *DepthFirstOrder {

	dfsOrder := &DepthFirstOrder{}
	dfsOrder.marked = make([]bool, g.V())
	dfsOrder.reversePost = basic.NewStack[int]()

	//dfs
	dfsOrder.dfs(g, 0)

	return dfsOrder
}

func (d *DepthFirstOrder) dfs(g Graph, v int) {
	d.marked[v] = true
	for w := range g.AdjOf(v) {
		if !d.marked[w] {
			d.marked[w] = true
			d.dfs(g, w)
		}
	}
	// dfs post order 再reverse一下
	d.reversePost.Push(v)
}
