# Bellman-Ford Algorithm

**Problem:** Directed graph, edges `[u, v, w]`, weights **can be negative**. Find shortest distances from `src`. If there's a **negative cycle** reachable from `src`, report it (shortest paths are then undefined: you could loop forever and keep getting cheaper).

## Idea: relax every edge, V − 1 times
A shortest path (with no negative cycle) visits each node at most once, so it has at most `V − 1` edges.
- After round 1, every shortest path that uses **1** edge is correct.
- After round 2, every one with **≤ 2** edges is correct.
- ...after round `V − 1`, all of them are.

```
dist = [INF...], dist[src] = 0
repeat V-1 times:
    for (u, v, w) in edges:
        if dist[u] != INF and dist[u] + w < dist[v]:
            dist[v] = dist[u] + w
// round V: if anything still improves → negative cycle
for (u, v, w) in edges:
    if dist[u] != INF and dist[u] + w < dist[v]: return NEGATIVE CYCLE
```

## Dry run
Edges: 3→2 (6), 5→3 (1), 0→1 (5), 1→5 (−3), 1→2 (−2), 3→4 (−2), 2→4 (3); src 0, V = 6
- round 1 (edge order above): 1 = 5, 5 = 2, 2 = 3, 4 = 6
- round 2: 3 = 3 (via 5), 4 = min(6, 3−2=1) = 1
- rounds 3–5: nothing changes
→ `[0 5 3 3 1 2]` ✅

## Negative cycle example
`0→1 (1), 1→2 (−1), 2→0 (−1)`: going around once costs −1, so it never stops improving. Round V still improves → report.

## Dijkstra vs Bellman-Ford
| | Dijkstra | Bellman-Ford |
|---|---|---|
| Negative weights | ✗ | ✓ |
| Finds negative cycles | ✗ | ✓ |
| Time | O(E log V) | O(V · E) |

Undirected graph with a negative edge? Each undirected edge is two directed ones, so `u ↔ v` with −1 is already a negative cycle.

- **Time:** O(V · E), **Space:** O(V)
