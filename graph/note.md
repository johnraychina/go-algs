

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