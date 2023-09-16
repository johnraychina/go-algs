package basic

import (
	"fmt"
	"os"
)

// 图问题分类
// 1. 路径问题(Path)，从一点到另外一点是否可达？最短距离是多少？
// 2. 环问题(Cycle)，Euler tour， Hamilton tour。
// 3. 连通性问题(Connectivity: 遍历所有节点的最优方案是什么？是否有关键节点，移除后会分为2个或以上的图？
// 4. 平面问题(Planarity), 是否能在一个平面内将图划出而没有边交叉？图的等效变形？
// 想一想，这里面哪些问题是简单的？哪些是困难一些的但是可解？哪一些是不可解的(intractable)？

type Graph interface {
	// AddEdge add an edge v-w
	AddEdge(v, w int)

	// AdjOf vertices adjacent to v
	AdjOf(v int) map[int]bool

	// V number of vertices
	V() int

	// E number of edges
	E() int

	// Degree compute the degree of v
	Degree(v int) int
}

type UndirectedGraph struct {
	Graph
	v   int            // number of vertices
	adj []map[int]bool // adjacent bag
}

// NewGraphFromFile create a graph from input file
func NewGraphFromFile(f os.File) *Graph {
	//todo
	return nil
}

// NewUndirectedGraph create an empty graph with V vertices
func NewUndirectedGraph(v int) Graph {
	g := &UndirectedGraph{v: v, adj: make([]map[int]bool, v)}
	for i, _ := range g.adj {
		g.adj[i] = make(map[int]bool)
	}
	return g
}

func (g *UndirectedGraph) AddEdge(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)

	g.adj[v][w] = true
	g.adj[w][v] = true
}

func (g *UndirectedGraph) validateVertex(v int) {
	if !(v >= 0 && v < g.v) {
		panic(fmt.Sprintf("vertex:%d out of bound:%d", v, g.v))
	}
}

func (g *UndirectedGraph) AdjOf(v int) map[int]bool {
	g.validateVertex(v)
	return g.adj[v]
}

func (g *UndirectedGraph) V() int {
	return g.v
}

func (g *UndirectedGraph) E() int {
	cnt := 0
	for _, m := range g.adj {
		cnt += len(m)
	}
	return cnt
}

func (g *UndirectedGraph) Degree(v int) int {
	g.validateVertex(v)
	return len(g.AdjOf(v))
}

type Paths interface {
	HasPathTo(v int) bool

	// PathTo path from s to v; empty if no such path
	PathTo(v int) []int
}

type DepthFirstPaths struct {
	Paths
	s int // source vertex
	//g *Graph // graph

	// to keep tree of paths. (edgeTo[w] == v) means that edge v-w taken to visit w for first time
	// 这个设计看似平平无奇， 但是意图非常巧妙，它保存从source遍历到的所有节点的前序节点。
	// 因为从一个点出发做dfs遍历，形成的是一棵树，
	// 想要或者完整的路径，可以设计以s为出发点的top-down的tree-node
	// 还有一种简单的办法，就是bottom-up反着来！以目标点为索引下标，存储它的前序节点。
	edgeTo []int

	marked []bool //to mark visited vertices
}

func DepthFirstPathsOf(g Graph, source int) *DepthFirstPaths {
	edgeTo := make([]int, g.V())
	marked := make([]bool, g.V())
	paths := &DepthFirstPaths{s: source, edgeTo: edgeTo, marked: marked}
	paths.dfs(g, source)
	return paths
}

func (p *DepthFirstPaths) dfs(g Graph, v int) {
	p.marked[v] = true
	for w := range g.AdjOf(v) {
		if !p.marked[w] {
			p.edgeTo[w] = v // 保存前序（父节点）
			p.dfs(g, w)
		}
	}
}

func (p *DepthFirstPaths) HasPathTo(w int) bool {
	return p.marked[w]
}

func (p *DepthFirstPaths) PathTo(w int) (path []int) {
	if !p.HasPathTo(w) {
		return path
	}

	for w != p.s {
		path = append(path, w)
		w = p.edgeTo[w]
	}
	path = append(path, p.s)

	// 由于edgeTo保存的的是父节点，会从w遍历到s, 我们想要返回s->w，需要调换顺序。
	l := len(path)
	for i := 0; i < l/2; i++ {
		path[i], path[l-1-i] = path[l-1-i], path[i]
	}
	return path
}

type BreadFirstPaths struct {
	Paths
	s      int // source vertex
	marked []bool
	edgeTo []int // parent links, keep the tree of the paths.
}

func (p *BreadFirstPaths) HasPathTo(w int) bool {
	return p.marked[w]
}

func (p *BreadFirstPaths) PathTo(w int) (path []int) {
	if !p.HasPathTo(w) {
		return path
	}

	// 自底向上遍历edgeTo得到出发点s的路径
	for p.edgeTo[w] != p.s {
		path = append(path, w)
		w = p.edgeTo[w]
	}
	path = append(path, p.s)

	return path
}

func BreadthFirstPathsOf(g Graph, s int) *BreadFirstPaths {
	return bfs(g, s)
}

func bfs(g Graph, s int) *BreadFirstPaths {
	// 广度优先搜索，以queue作为存储，保存一层（一圈）的节点
	// 这种搜索方式可以理解为把图组织成一个圆， 看成从圆心s开始，向外遍历，每次遍历一圈（一层）节点
	// 或者把图想象为一张渔网，你把渔网的一个结拎起来，从上到下逐层遍历节点。
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	paths := &BreadFirstPaths{marked: marked, edgeTo: edgeTo}

	queue := NewLinkedQueue[int]()
	queue.Enqueue(s) // s 起始节点
	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for w := range g.AdjOf(v) {
			if !paths.marked[w] {
				queue.Enqueue(w)
				marked[w] = true
				edgeTo[w] = v
			}
		}
	}

	return paths
}
