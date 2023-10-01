package graph

import (
	"go-algs/basic"
	"math"
)

// todo 探索检测正向环

//Bellman-Ford shortest paths算法：
//初始化：distTo[s]=0, distTo[v]=Max
//迭代V次：
//- 更新顶点v对应的边（relax each edge）

type BellmanFordSP struct {
	distTo []float64
	edgeTo []*DirectedEdge
	onQ    []bool
	queue  *basic.LinkedQueue[int] // queue of vertices
	cost   int                     // relax cost
	cycle  *basic.Stack[*DirectedEdge]
}

func NewBellmanFordSP(g *EdgeWeightedDiGraph, source int) *BellmanFordSP {
	// initialize
	sp := &BellmanFordSP{}
	sp.distTo = make([]float64, g.V())
	for i, _ := range sp.distTo {
		sp.distTo[i] = math.MaxFloat64
	}
	sp.distTo[source] = 0

	sp.edgeTo = make([]*DirectedEdge, g.V())

	sp.onQ = make([]bool, g.V())
	for i, _ := range sp.onQ {
		sp.onQ[i] = false
	}

	sp.queue = basic.NewLinkedQueue[int]()
	sp.queue.Enqueue(source)
	sp.onQ[source] = true

	// traverse and relax
	for !sp.queue.IsEmpty() && !sp.hasNegativeCycle() {
		v := sp.queue.Dequeue()
		sp.onQ[v] = false
		sp.relax(g, v)
	}

	return sp
}

func (sp *BellmanFordSP) relax(g *EdgeWeightedDiGraph, v int) {
	for w, e := range g.AdjOf(v) {
		if sp.distTo[w] > sp.distTo[v]+e.weight {
			sp.distTo[w] = sp.distTo[v] + e.weight
			sp.edgeTo[w] = e
			if !sp.onQ[w] {
				sp.queue.Enqueue(w)
				sp.onQ[w] = true
			}
		}

		sp.cost++
		if sp.cost%g.V() == 0 { // 每次relax 所有个顶点后，都判断一次是否存在环
			sp.findNegativeCycle()
			if sp.hasNegativeCycle() {
				return
			}
		}
	}
}

func (sp *BellmanFordSP) findNegativeCycle() {
	V := len(sp.edgeTo)
	//用遍历过的edge来构造一个graph
	g := NewEdgeWeightedDiGraph(V)
	for _, edge := range sp.edgeTo {
		if edge != nil {
			g.AddDirectedEdge(edge)
		}
	}

	//对graph做dfs遍历，用一个marked数组标记跟踪，一个onStack数组标记在dfs遍历的栈
	c := NewEdgeWeightedDirectedCycle(g)
	sp.cycle = c.Get()
}

func (sp *BellmanFordSP) hasNegativeCycle() bool {
	return sp.cycle != nil && !sp.cycle.IsEmpty()
}

type EdgeWeightedDirectedCycle struct {
	edgeTo  []*DirectedEdge
	marked  []bool
	onStack []bool
	cycle   *basic.Stack[*DirectedEdge]
}

func NewEdgeWeightedDirectedCycle(g *EdgeWeightedDiGraph) *EdgeWeightedDirectedCycle {
	edgeTo := make([]*DirectedEdge, g.V())
	marked := make([]bool, g.V())
	onStack := make([]bool, g.V())
	//cycle := basic.NewStack[*DirectedEdge]()
	c := &EdgeWeightedDirectedCycle{edgeTo: edgeTo, marked: marked, onStack: onStack}
	for v := 0; v < g.V(); v++ {
		if !c.marked[v] {
			c.dfs(g, v)
		}
	}

	return c
}

func (c *EdgeWeightedDirectedCycle) dfs(g *EdgeWeightedDiGraph, v int) {
	c.onStack[v] = true
	c.marked[v] = true

	for w, e := range g.AdjOf(v) {
		// short circuit if directed cycle found
		if c.cycle != nil {
			return
		}
		if !c.marked[w] {
			c.edgeTo[w] = e
			c.dfs(g, w)
		} else if c.onStack[w] {
			// trace back directed cycle
			c.cycle = basic.NewStack[*DirectedEdge]()
			f := e
			for f.From() != w { // todo: 思考，如果有多个cycle会怎么样？
				c.cycle.Push(f)
				f = c.edgeTo[f.From()]
			}
			c.cycle.Push(f)

			return
		}
	}

	c.onStack[v] = false
}

func (c *EdgeWeightedDirectedCycle) Get() *basic.Stack[*DirectedEdge] {
	return c.cycle
}
