# Step 15: Graphs

> Status: theory, 15.1 – 15.3 done (23 problems); shortest paths, MST and SCC coming ⏳

## Words to know
- **Vertex (node)** and **edge** (connection).
- **Undirected** (two-way) vs **directed** (one-way).
- **Weighted:** edges have a cost.
- **Path, cycle, connected component** (a group of nodes reachable from each other).
- **Degree:** number of edges at a node (in-degree / out-degree for directed graphs).
- **DAG:** directed acyclic graph (directed, no cycles).

## Storing a graph in Go
```go
// adjacency list: n nodes, edges given as pairs
adj := make([][]int, n)
for _, e := range edges {
    u, v := e[0], e[1]
    adj[u] = append(adj[u], v)
    adj[v] = append(adj[v], u) // skip this line for directed graphs
}

// weighted: store (neighbour, weight)
type Edge struct{ To, W int }
wadj := make([][]Edge, n)
```
Adjacency list uses O(V + E) memory. A matrix uses O(V²); only use it for dense or small graphs.

**Grids are graphs too:** each cell is a node; neighbours are the 4 (or 8) directions.
```go
dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
```

## BFS and DFS
| | BFS | DFS |
|---|---|---|
| Uses | queue | recursion / stack |
| Visits | level by level | goes deep first |
| Good for | shortest path in **unweighted** graphs, levels/distances | components, cycles, topo sort, backtracking |
| Time | O(V + E) | O(V + E) |

Always keep a `visited` array, or you'll loop forever on cycles. For disconnected graphs, start a BFS/DFS from every unvisited node.

## Big algorithms and when to use them
| Need | Algorithm | Time |
|---|---|---|
| Order tasks with dependencies (DAG) | **Topological sort** (DFS or Kahn's BFS with in-degrees) | O(V + E) |
| Shortest path, unit weights | BFS | O(V + E) |
| Shortest path, non-negative weights | **Dijkstra** (min-heap) | O(E log V) |
| Shortest path, negative weights / detect negative cycle | **Bellman-Ford** (relax all edges V−1 times) | O(V·E) |
| All-pairs shortest paths | **Floyd-Warshall** | O(V³) |
| Minimum spanning tree | **Prim's** (heap) or **Kruskal's** (sort edges + DSU) | O(E log E) |
| Are two nodes connected? (merging groups) | **Disjoint Set Union** with path compression + union by size | ~O(1) per op |
| Critical edges / nodes | **Tarjan's** bridges / articulation points | O(V + E) |
| Strongly connected components | **Kosaraju's** | O(V + E) |

## Cycle detection
- **Undirected:** during BFS/DFS, reaching a visited node that isn't your parent → cycle.
- **Directed:** DFS with a "path visited" (on the current recursion stack) array; or Kahn's algorithm: if topo order has fewer than V nodes → cycle.

## Problems

### 15.1 Learning
| # | Problem | Key idea |
|---|---|---|
| 1 | [Graph representation](15.1-learning/01-graph-representation/) | adjacency list vs matrix |
| 2 | [Connected components](15.1-learning/02-connected-components/) | start a traversal from each unvisited node |
| 3 | [BFS traversal](15.1-learning/03-bfs/) | queue, mark visited when pushing |
| 4 | [DFS traversal](15.1-learning/04-dfs/) | recursion, go deep first |

### 15.2 BFS / DFS problems
| # | Problem | Key idea |
|---|---|---|
| 1 | [Number of provinces](15.2-bfs-dfs/01-number-of-provinces/) | count DFS starts |
| 2 | [Number of islands](15.2-bfs-dfs/02-number-of-islands/) | grid DFS, 8 directions |
| 3 | [Flood fill](15.2-bfs-dfs/03-flood-fill/) | DFS from start, recolour same-colour cells |
| 4 | [Rotting oranges](15.2-bfs-dfs/04-rotting-oranges/) | multi-source BFS, count levels |
| 5 | [Detect cycle in undirected graph (BFS / DFS)](15.2-bfs-dfs/05-cycle-undirected/) | visited neighbour that isn't the parent |
| 6 | [0/1 matrix (distance to nearest 0)](15.2-bfs-dfs/06-distance-nearest-zero/) | multi-source BFS from all zeros |
| 7 | [Surrounded regions](15.2-bfs-dfs/07-surrounded-regions/) | DFS from border O's, flip the rest |
| 8 | [Number of enclaves](15.2-bfs-dfs/08-number-of-enclaves/) | remove land touching the border |
| 9 | [Word ladder I](15.2-bfs-dfs/09-word-ladder/) | BFS, change one letter at a time |
| 10 | [Word ladder II](15.2-bfs-dfs/10-word-ladder-ii/) | BFS levels + backtrack parents |
| 11 | [Number of distinct islands](15.2-bfs-dfs/11-distinct-islands/) | shape = offsets from the first cell |
| 12 | [Bipartite graph](15.2-bfs-dfs/12-bipartite-graph/) | 2-colour, neighbour same colour = fail |
| 13 | [Detect cycle in directed graph (DFS)](15.2-bfs-dfs/13-cycle-directed-dfs/) | onPath array (back edge) |

### 15.3 Topological sort
| # | Problem | Key idea |
|---|---|---|
| 1 | [Topological sort (DFS)](15.3-topo-sort/01-topo-sort-dfs/) | push after children, reverse |
| 2 | [Kahn's algorithm (BFS)](15.3-topo-sort/02-kahns-algorithm/) | pop in-degree 0 nodes |
| 3 | [Cycle detection in directed graph (BFS)](15.3-topo-sort/03-cycle-directed-bfs/) | Kahn's processed < V |
| 4 | [Course schedule I and II](15.3-topo-sort/04-course-schedule/) | edge b → a, then Kahn's |
| 5 | [Find eventual safe states](15.3-topo-sort/05-eventual-safe-states/) | reverse graph + Kahn's |
| 6 | [Alien dictionary](15.3-topo-sort/06-alien-dictionary/) | first differing letter = edge, topo sort |

### 15.4 Shortest paths
| # | Problem | Status |
|---|---|---|
| 1 | Shortest path, unit weights | ⏳ |
| 2 | Shortest path in a DAG | ⏳ |
| 3 | Dijkstra's algorithm | ⏳ |
| 4 | Shortest path in a binary maze | ⏳ |
| 5 | Path with minimum effort | ⏳ |
| 6 | Cheapest flights within K stops | ⏳ |
| 7 | Network delay time | ⏳ |
| 8 | Number of ways to arrive at destination | ⏳ |
| 9 | Minimum multiplications to reach end | ⏳ |
| 10 | Bellman-Ford | ⏳ |
| 11 | Floyd-Warshall | ⏳ |
| 12 | City with the smallest number of neighbours | ⏳ |

### 15.5 MST and Disjoint Set
| # | Problem | Status |
|---|---|---|
| 1 | Prim's algorithm | ⏳ |
| 2 | Disjoint Set (union by rank/size, path compression) | ⏳ |
| 3 | Kruskal's algorithm | ⏳ |
| 4 | Operations to make network connected | ⏳ |
| 5 | Most stones removed | ⏳ |
| 6 | Accounts merge | ⏳ |
| 7 | Number of islands II | ⏳ |
| 8 | Making a large island | ⏳ |
| 9 | Swim in rising water | ⏳ |

### 15.6 Other algorithms
| # | Problem | Status |
|---|---|---|
| 1 | Bridges in a graph (Tarjan) | ⏳ |
| 2 | Articulation points | ⏳ |
| 3 | Kosaraju's algorithm (SCC) | ⏳ |
