# Shortest Path in a DAG (weighted)

**Problem:** A directed acyclic graph with weighted edges `[u, v, w]`. Find the shortest distance from `src` (node 0) to every node; `-1` if unreachable. Weights may even be negative (no cycles, so that's fine).

## Idea: relax edges in topological order
"Relax edge `u → v`" means: `if dist[u] + w < dist[v] { dist[v] = dist[u] + w }`.

In a topological order, **all** edges into `v` come from nodes placed before `v`. So if we process nodes in that order, by the time we reach `u`, `dist[u]` is already final. Relax its outgoing edges once and move on. Each edge is relaxed exactly once.

```
order = topoSort(graph)
dist = [INF...], dist[src] = 0
for u in order:
    if dist[u] == INF: continue     // not reachable yet, nothing to push
    for (v, w) in adj[u]:
        dist[v] = min(dist[v], dist[u] + w)
```

## Dry run
Edges: 0→1 (2), 0→4 (1), 4→5 (4), 4→2 (2), 1→2 (3), 2→3 (6), 5→3 (1); topo order e.g. `0 1 4 5 2 3`
- u=0: dist1=2, dist4=1
- u=1: dist2=5
- u=4: dist5=5, dist2=min(5,3)=3
- u=5: dist3=6
- u=2: dist3=min(6, 9)=6
→ `[0 2 3 6 1 5]` ✅

- **Time:** O(V + E) (faster than Dijkstra, and allows negative weights)
- **Space:** O(V + E)
