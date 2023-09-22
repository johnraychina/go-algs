package graph

import (
	"go-algs/basic"
	"math"
)

// 问题抽象：给定一个有权重的有向图，请找到s点到t点的最短路径。
// Given an edge-weighted digraph, find the shortest path from s to t.

//・PERT/CPM. 项目管理，关键路径。
//・Map routing. 地图寻路
//・Seam carving. 图像伸缩
//・Robot navigation. 机器人导航
//・Texture mapping. 着色 https://zhuanlan.zhihu.com/p/453457104
//・Typesetting in TeX. TeX排版
//・Urban traffic planning. 城市交通规划
//・Optimal pipelining of VLSI chip. VLSI(Very Large Scale Integration)超大规模集成芯片最优流水线设计
// see https://yearn.xyz/docs/vlsi-%E5%90%8E%E7%AB%AF%E8%AE%BE%E8%AE%A1/3.1-vlsi-%E5%90%8E%E7%AB%AF%E8%AE%BE%E8%AE%A1-%E7%AE%80%E4%BB%8B/
//・Telemarketer operator scheduling. 电话接线员调度
//・Routing of telecommunications messages. 电信消息路由
//・Network routing protocols (OSPF, BGP, RIP). 网络路由协议(OSPF, BGP骨干网, RIP）
// see https://orhanergun.net/ospf-vs-rip#:~:text=OSPF%20is%20an%20open%20standard,metric%20for%20determining%20best%20path.
//・Exploiting arbitrage opportunities in currency exchange. 在汇率交易中发现套利机会
//・Optimal truck routing through given traffic congestion pattern. 给定拥堵模式下，求卡车运输最优路径。

// 最短路径问题的变体
//Which vertices?
//・Single source: from one vertex s to every other vertex.
//・Source-sink: from one vertex s to another t.
//・All pairs: between all pairs of vertices.

//Restrictions on edge weights?
//・Non-negative weights.
//・Euclidean weights.
//・Arbitrary weights.

//Cycles?
//・No directed cycles.
//・No "negative cycles."

// ShortestPaths Single-source shortest paths API
type ShortestPaths struct {
	s      int
	edgeTo []*DirectedEdge // parent link representation. from source to any vertices forms a spanning tree.
	distTo []float32
}

// todo iterate
//Goal. Find the shortest path from s to every other vertex.
//Observation. A shortest-paths tree (SPT) solution exists. Why?
//Consequence. Can represent the SPT with two vertex-indexed arrays:
//・ distTo[v] is length of shortest path from s to v.
//・ edgeTo[v] is last edge on shortest path from s to v.

// 松弛算法，本质是动态规划的变体（会更新权重）
// 如果有别的更短路径s→v→w 比直接 s→w更短，则更新（relax松弛）到w的路径和长度。
//Relax edge e = v→w.
//・ distTo[v] is length of shortest known path from s to v.
//・ distTo[w] is length of shortest known path from s to w.
//・ edgeTo[w] is last edge on shortest known path from s to w.
//・If e = v→w gives shorter path to w through v,
//	update both distTo[w] and edgeTo[w].

func NewShortestPaths(g EdgeWeightedDiGraph, source int) *ShortestPaths {
	// todo 检测环

	sp := &ShortestPaths{s: source}
	edgeTo := make([]*DirectedEdge, g.V()) // parent link representation
	distTo := make([]float32, g.V())       // dp状态

	// Initialize distTo[s] = 0 and distTo[v] = ∞ for all other vertices.
	//Repeat until optimality conditions are satisfied:
	// - Relax any edge.
	// 初始条件
	distTo[source] = 0
	for i := range edgeTo {
		distTo[i] = math.MaxFloat32
	}

	// candidate edges, s->any distance as key.
	minPQ := basic.NewIndexMinPQ[float32, *DirectedEdge](g.V())

	// 初始：从source开始的临边开始向外遍历
	// 和Prim算法非常像，只不过Prim只关心最短路径是最外层的边（外侧贪心），而这里是关心从source到当前顶点的边（整体动态规划）
	for v, edge := range g.AdjOf(source) {
		minPQ.Insert(v, edge) // take vertex as the index of priority queue
	}

	for !minPQ.IsEmpty() {
		// 最小路径的点: min(s->any) as v
		v := minPQ.DelMin()

		// （动态规划状态转移）relax最小边的邻边, v == e.From
		relaxed := false
		for _, e := range g.AdjOf(v) {
			relaxed = relaxed && sp.relax(e, minPQ)
		}
		// 停止条件：没有更短的路径可以做relax了
		// todo 考虑相同distance的情况
		if !relaxed {
			break
		}
	}

	return sp
}

func (sp *ShortestPaths) relax(e *DirectedEdge, minPQ *basic.IndexMinPQ[float32, *DirectedEdge]) bool {
	v := e.From()
	w := e.To()
	if sp.distTo[v]+e.Weight() < sp.distTo[w] {
		sp.edgeTo[w] = e
		sp.distTo[w] = sp.distTo[v] + e.Weight()

		if minPQ.Contains(w) {
			// s->w已有路径，distance更新后需要更新优先级队列
			minPQ.DecreaseKey(w, sp.distTo[w]) // 注意：新松弛（更新）后，更备选边的排序
		} else {
			// w首次纳入，span the tree: 把v的邻接点，加入s->any外侧顶点队列
			minPQ.Insert(w, e)
		}
		return true
	}
	return false
}

// PathTo the shortest path from s to v
func (sp *ShortestPaths) PathTo(v int) []*DirectedEdge {
	//todo
}

// DistTo length of shortest path from s to v
func (sp *ShortestPaths) DistTo(v int) []*DirectedEdge {
	//todo
}

// HasPathTo is there any path from s to v?
func (sp *ShortestPaths) HasPathTo(v int) bool {
	//todo
}
