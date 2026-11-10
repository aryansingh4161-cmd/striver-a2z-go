# Shortest Path in an Undirected Graph with Unit Weights

**Problem:** Every edge has weight 1. Find the shortest distance from `src` to every node (`-1` if unreachable).

## Why plain BFS is enough
BFS visits nodes **level by level**: first everything 1 edge away, then 2 edges away, and so on. So the first time BFS reaches a node is through the fewest edges. With all weights = 1, "fewest edges" = "shortest distance".

```
dist[src] = 0, others = -1
queue = [src]
while queue:
    u = pop
    for v in adj[u]:
        if dist[v] == -1:          // first time we see v
            dist[v] = dist[u] + 1
            push v
```
The `dist` array doubles as the visited array.

## Dry run
Edges: 0-1, 0-3, 1-2, 3-4, 4-5, 5-6, 2-6, 6-7, 6-8, 7-8; src = 0
- level 0: {0}
- level 1: {1, 3}
- level 2: {2, 4}
- level 3: {6, 5}
- level 4: {7, 8}
→ `dist = [0 1 2 1 2 3 3 4 4]` ✅

- **Time:** O(V + E), **Space:** O(V)

## When this breaks
As soon as weights differ, the path with the fewest edges might not be the cheapest. Then use [Dijkstra](../03-dijkstra/) (or a [DAG](../02-shortest-path-dag/) trick, or 0-1 BFS for weights 0/1).
