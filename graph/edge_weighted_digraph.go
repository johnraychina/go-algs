package graph

// EdgeWeightedDiGraph 有权重有向图
type EdgeWeightedDiGraph struct {
	v   int                     // number of vertices
	adj []map[int]*DirectedEdge // adjacent vertices
}

func NewEdgeWeightedDiGraph(v int) *EdgeWeightedDiGraph {
	adj := make([]map[int]*DirectedEdge, v)
	for i := range adj {
		adj[i] = make(map[int]*DirectedEdge)
	}
	return &EdgeWeightedDiGraph{v: v, adj: adj}
}

func (g *EdgeWeightedDiGraph) AddDirectedEdge(e *DirectedEdge) {
	g.adj[e.From()][e.To()] = e
}

func (g *EdgeWeightedDiGraph) AdjOf(v int) map[int]*DirectedEdge {
	return g.adj[v]
}

func (g *EdgeWeightedDiGraph) DirectedEdges() (DirectedEdges []*DirectedEdge) {
	for _, m := range g.adj {
		for _, e := range m {
			DirectedEdges = append(DirectedEdges, e)
		}
	}
	return DirectedEdges
}

func (g *EdgeWeightedDiGraph) V() int {
	return g.v
}

func (g *EdgeWeightedDiGraph) E() int {
	DirectedEdgeCount := 0
	for _, m := range g.adj {
		DirectedEdgeCount += len(m)
	}
	return DirectedEdgeCount
}

type DirectedEdge struct {
	v, w   int
	weight float64
}

func (e *DirectedEdge) From() int {
	return e.v
}
func (e *DirectedEdge) To() int {
	return e.w
}
func (e *DirectedEdge) Weight() float64 {
	return e.weight
}
func (e *DirectedEdge) Key() float64 {
	return e.weight
}
