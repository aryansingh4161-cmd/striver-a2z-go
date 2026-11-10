# Detect a Cycle in a Directed Graph (DFS)

**Problem:** Does the directed graph have a cycle?

## Why the undirected method fails
In a directed graph, reaching an already-visited node does **not** mean a cycle: `0→1, 0→2, 2→1` visits 1 twice, but there's no cycle.

## Approach: Track the current DFS path
Keep two arrays:
- `visited[u]`: ever visited
- `onPath[u]`: currently on the recursion stack (the path we're exploring)

A cycle exists iff DFS reaches a node that is **on the current path** (a back edge). When `dfs(u)` finishes, set `onPath[u] = false`.

(Equivalently, use 3 colours: white = unvisited, grey = on path, black = finished.)
- **Time:** O(n + m), **Space:** O(n)

The BFS version uses Kahn's algorithm: see [Step 15.3](../../15.3-topo-sort/03-cycle-directed-bfs/).
