# Kruskal's Algorithm (Minimum Spanning Tree)

**Problem:** Same as [Prim's](../01-prims-algorithm/): total weight of a minimum spanning tree.

## Idea: cheapest edges first, skip the ones that make a cycle
1. Sort all edges by weight.
2. Go through them in order. For edge `u – v`:
   - If `u` and `v` are already connected → adding it makes a cycle → **skip**.
   - Otherwise → **take it** and merge their groups.
3. Stop after `V − 1` edges.

"Are they already connected?" and "merge" are exactly what a [Disjoint Set](../02-disjoint-set/) does in ~O(1).

```
sort edges by w
dsu = New(n)
for (w, u, v) in edges:
    if dsu.Union(u, v): sum += w      // Union is false when already connected
```

## Dry run
Edges sorted: 0-1 (2), 1-2 (3), 1-4 (5), 0-3 (6), 2-4 (7), 1-3 (8)
| edge | connected already? | action | sum |
|---|---|---|---|
| 0-1 (2) | no | take | 2 |
| 1-2 (3) | no | take | 5 |
| 1-4 (5) | no | take | 10 |
| 0-3 (6) | no | take | 16 |
| 2-4 (7) | yes (2-1-4) | skip | |
| 1-3 (8) | yes | skip | |
**16** ✅ (same as Prim)

- **Time:** O(E log E) for sorting, **Space:** O(V + E)

## Prim or Kruskal?
| | Prim | Kruskal |
|---|---|---|
| Needs | adjacency list + heap | edge list + DSU |
| Feels like | Dijkstra | sort + greedy |
| Disconnected graph | gives MST of one component | gives a minimum spanning **forest** |
Both give the same total.
