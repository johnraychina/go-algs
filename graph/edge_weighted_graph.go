package graph

import "go-algs/basic"

type EdgeWeightedGraph struct {
	v   int             // number of vertices
	adj []map[int]*Edge // adjacent vertices
}

func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	adj := make([]map[int]*Edge, v)
	for i, _ := range adj {
		adj[i] = make(map[int]*Edge)
	}
	return &EdgeWeightedGraph{v: v, adj: adj}
}

func (g *EdgeWeightedGraph) AddEdge(edge *Edge) {
	// 无向边，本质是两个方向都可以
	g.adj[edge.v][edge.w] = edge
	g.adj[edge.w][edge.v] = edge
}

func (g *EdgeWeightedGraph) AdjOf(v int) map[int]*Edge {
	return g.adj[v]
}

func (g *EdgeWeightedGraph) Edges() (edges []*Edge) {
	for _, m := range g.adj {
		for _, e := range m {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *EdgeWeightedGraph) V() int {
	return g.v
}

func (g *EdgeWeightedGraph) E() int {
	edgeCount := 0
	for _, m := range g.adj {
		edgeCount += len(m)
	}
	return edgeCount
}

type Edge struct {
	basic.ComparableKey[float32]
	v      int
	w      int
	weight float32
}

// NewEdge create a weighted edge v-w
func NewEdge(v, w int, weight float32) *Edge {
	return &Edge{v: v, w: w, weight: weight}
}
func (e *Edge) Key() float32 {
	return e.weight
}
func (e *Edge) Either() int {
	return e.v
}
func (e *Edge) Other(v int) int {
	if e.v == v {
		return e.w
	}
	return v
}

func (e *Edge) CompareTo(that Edge) int {
	if e.weight < that.weight {
		return -1
	} else if e.weight > that.weight {
		return 1
	}
	return 0
}

func (e *Edge) Weight() float32 {
	return e.weight
}
