
通用的最短路径SPT算法，本质上是一个动态规划算法，distTo保存状态。
对于起始点s, 初始化distTo[s]=0, 对于其他的节点v, distTo[v] = ∞
重复以下处理，直到达到最优条件:
- 松弛（更新）任意的边

Generic algorithm (to compute SPT from s)
Initialize distTo[s] = 0 and distTo[v] = ∞ for all other vertices. 
Repeat until optimality conditions are satisfied:
- Relax any edge.

我们如何选择松弛的边？
例子1: Dijkstra最短路径算法（权重不能为负）
例子2: 拓扑排序(无环)
例子3: Bellman-Ford算法 (无负数的环)

Efficient implementations. How to choose which edge to relax? 
Ex 1. Dijkstra's algorithm (nonnegative weights).
Ex 2. Topological sort algorithm (no directed cycles).
Ex 3. Bellman-Ford algorithm (no negative cycles).


Dijkstra 最短路径算法和Prim非常相似，都会生成树，区别在于如何选择下一个节点：
Prim: （通过无向图）选择离树最近的节点
Dijkstra: (通过有向图)选择离source起点最近的节点
・Prim’s: Closest vertex to the tree (via an undirected edge). 
・Dijkstra’s: Closest vertex to the source (via a directed path).

## Dijkstra’s algorithm.
・Nearly linear-time when weights are non-negative. 
・Generalization encompasses DFS, BFS, and Prim.
## Acyclic edge-weighted digraphs.
・Arise in applications.
・Faster than Dijkstra’s algorithm. 
・Negative weights are no problem.
## Negative weights and negative cycles.
・Arise in applications.
・If no negative cycles, can find shortest paths via Bellman-Ford. 
・If negative cycles, can find one via Bellman-Ford.



Remark 1. Directed cycles make the problem harder.
Remark 2. Negative weights make the problem harder.
Remark 3. Negative cycles makes the problem intractable.


有向环难
负数权重难
负数环让问题不可解


# 为什么负数会让问题变难？
负数破坏了状态转换的依赖假设，从而破坏最优条件的判断。

# 为什么有环会变难，无环更简单？
无环，可以得到拓扑排序，可以做更强依赖假设，更少的状态转换。
because of topological order, no edge pointing to v
will be relaxed after v is relaxed
