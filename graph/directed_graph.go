package graph

type DiGraph struct {
	Graph
	v   int            // number of vertices
	adj []map[int]bool // adjacent bag
}

func NewDiGraph(v int) *DiGraph {
	g := &DiGraph{v: v, adj: make([]map[int]bool, v)}
	for i := range g.adj {
		g.adj[i] = make(map[int]bool)
	}
	return g
}

// AddEdge add an edge v-w
func (g *DiGraph) AddEdge(v, w int) {
	g.adj[v][w] = true
}

// AdjOf vertices adjacent to v
func (g *DiGraph) AdjOf(v int) map[int]bool {
	return g.adj[v]
}

// V number of vertices
func (g *DiGraph) V() int {
	return g.v
}

// E number of edges
func (g *DiGraph) E() int {
	cnt := 0
	for _, m := range g.adj {
		cnt += len(m)
	}
	return cnt
}

// Degree compute the degree of v
func (g *DiGraph) Degree(v int) int {
	panic("implement me")
}
