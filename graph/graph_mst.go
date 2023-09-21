package graph

import "go-algs/basic"

// Minimum Spanning Tree 最小生成树，应用场景包括：
//・Dithering. see https://en.wikipedia.org/wiki/Dither
// 例如Floyd–Steinberg dithering 抖动算法
// 抖动是一种有意应用的噪声形式，用于随机化量化误差，防止出现大规模图案，例如图像中的色带。抖动通常用于数字音频和视频数据的处理。
// 抖动的常见用途是将灰度图像转换为黑白图像，使得新图像中黑点的密度接近原始图像中的平均灰度级。
//・Cluster analysis. 聚簇分析
//・Max bottleneck paths. 最大瓶颈路径
//・Real-time face verification. 实时人脸验证
//・LDPC codes for error correction. LDPC 纠错码, see: https://zhuanlan.zhihu.com/p/514670102
//・Image registration with Renyi entropy. 基于Renyi熵的图像配准（registration）
//・Find road networks in satellite and aerial imagery.寻找卫星航空图像中的道路
//・Reducing data storage in sequencing amino acids in a protein.减少蛋白质序列的数据存储
//・Model locality of particle interactions in turbulent fluid flows.模拟湍流中粒子相互作用的局部性
//・Autoconfig protocol for Ethernet bridging to avoid cycles in a network. 自动配置以太网桥避免环
//・Approximation algorithms for NP-hard problems (e.g., TSP, Steiner tree). NP问题估计(Traveling Salesman Problem旅行商问题, Steiner树-组合优化）
//・Network design (communication, electrical, hydraulic, computer, road).网络设计（通信、电力、液压系统、计算机网络、道路）

// 定义“切分”: 一个图的切分指将定点分割成两个非空的集合
// 定义“跨连接边”:一个“跨连接边”连接两个集合中的两点
// 切分的性质：给定一个切分，跨连接边中的最小边属于MST（最小生成树）
//Def. A cut in a graph is a partition of its vertices into two (nonempty) sets.
//Def. A crossing edge connects a vertex in one set with a vertex in the other.
//Cut property. Given any cut, the crossing edge of min weight is in the MST.

// Greedy MST algorithm
// Start with all edges colored gray.
// Find cut with no black crossing edges; color its min-weighted edge black.
// Repeat until V-1 edges are colored black.

// 命题：贪心算法会计算出MST
// Proposition. The greedy algorithm computes the MST.
// Any edge colored black is in the MST (via cut property).
//・Fewer than V - 1 black edges ⇒ cut with no black crossing edges.
//(consider cut whose vertices are one connected component)

//高效实现：选择cut? 找到最小权重的边?
//Efficient implementations. Choose cut? Find min-weight edge?
//Ex 1. Kruskal's algorithm. [stay tuned]
//Ex 2. Prim's algorithm. [stay tuned]
//Ex 3. Borüvka's algorithm.

type MST interface {
	Edges() []*Edge  // all edges of MST
	Weight() float32 // total weight of MST
}

type KruskalMST struct {
	mst *basic.LinkedQueue[*Edge]
}

// NewKruskalMST 最小生成树算法：
// 把所有的边按权重排序，每次取出最小边构建到树中（除非生成环）
// Sort edges in ascending order of weight.
// Add next edge to tree T unless doing so would create a cycle
// 如何高效判断环？
// 使用Union-Find算法
// 技术总结：优先级队列MinPQ + 并查集Union-Find
func NewKruskalMST(g *EdgeWeightedGraph) *KruskalMST {
	q := basic.NewMinPQ[float32, *Edge]()
	uf := basic.NewQuickUnionUF(g.V())
	result := basic.NewLinkedQueue[*Edge]()
	for _, edge := range g.Edges() {
		q.Insert(edge)
	}

	for !q.IsEmpty() {
		nextMinEdge := q.DelMin()
		if !uf.Connected(nextMinEdge.v, nextMinEdge.w) {
			uf.Union(nextMinEdge.v, nextMinEdge.w)
			result.Enqueue(nextMinEdge)
		}
	}

	return &KruskalMST{mst: result}
}
