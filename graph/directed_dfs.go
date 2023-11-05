package graph

import "strconv"

type DirectedDFS struct {
	marked []bool // marked[v]=true if v is reachable from source(s)
	count  int    // number of vertices reachable from source(s)
}

func NewDirectedDFS(G *DiGraph, v []int) *DirectedDFS {
	d := &DirectedDFS{
		marked: make([]bool, G.V(), G.V()),
		count:  0,
	}
	d.validateVertices(v)

	for i := range v {
		if !d.marked[v[i]] {
			d.dfs(G, v[i])
		}
	}

	return d
}

func (d *DirectedDFS) validateVertices(vertices []int) {
	V := len(d.marked)
	for s := range vertices {
		if s >= V || s < 0 {
			panic("illegal vertex:" + strconv.Itoa(s))
		}
	}
}

func (d *DirectedDFS) dfs(G *DiGraph, v int) {
	d.marked[v] = true
	d.count++

	for w := range G.AdjOf(v) {
		if !d.marked[w] {
			d.dfs(G, w)
		}
	}
}

func (d *DirectedDFS) Marked(v int) bool {
	return d.marked[v]
}

func (d *DirectedDFS) Count() int {
	return d.count
}
