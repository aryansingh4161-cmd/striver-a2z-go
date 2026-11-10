# Depth-First Search (DFS)

Go as **deep** as possible along one path before backing up.
```
dfs(u):
    visited[u] = true; visit u
    for v in adj[u]:
        if not visited[v]: dfs(v)
```
- **Time:** O(n + m), **Space:** O(n) (visited + recursion depth)

## BFS vs DFS
| | BFS | DFS |
|---|---|---|
| Data structure | queue | recursion / stack |
| Finds shortest path (unweighted)? | **yes** | no |
| Good for | distances, levels, "nearest" | components, cycles, topo sort, backtracking |

Very deep graphs (e.g. a path of 10⁶ nodes) can make recursion slow or hit limits. Go's stack grows dynamically so it's usually fine; otherwise use an explicit stack (`DFSIterative` in `solution.go`).
