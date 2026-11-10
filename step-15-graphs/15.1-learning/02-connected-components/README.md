# Connected Components

**Problem:** Count the groups of nodes that are connected to each other (undirected graph).

n = 5, edges `0-1, 1-2, 3-4` → 2 components

## Approach: Start a traversal from every unvisited node
```
count = 0
for u in 0..n-1:
    if not visited[u]:
        count++
        dfs(u)            // marks the whole component
```
Every traversal problem on a possibly disconnected graph needs this outer loop, otherwise nodes in other components are never seen.
- **Time:** O(n + m), **Space:** O(n)
