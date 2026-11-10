# Number of Ways to Arrive at Destination

**Problem:** Undirected weighted graph, nodes `0..n-1`. Count how many **different shortest paths** go from `0` to `n-1`. Answer modulo `1e9 + 7`.

## Idea: Dijkstra + a `ways` array
Next to `dist[v]`, keep `ways[v]` = number of shortest paths to `v` found so far.

When relaxing edge `u → v` with `nd = dist[u] + w`:
- `nd < dist[v]`: found a **strictly shorter** route. Old routes no longer count → `dist[v] = nd`, `ways[v] = ways[u]`, push.
- `nd == dist[v]`: another route of the **same** best length → `ways[v] += ways[u]`.
- `nd > dist[v]`: ignore.

Start: `dist[0] = 0`, `ways[0] = 1`.

Why is this correct? Dijkstra pops `u` only when `dist[u]` is final, and every node that could be the second-to-last stop on a shortest path to `u` is popped before `u` (non-negative weights; with **zero** weights this can fail, but the problem uses positive weights). So `ways[u]` is complete before `u` spreads it.

## Dry run
Edges: 0-1 (2), 0-2 (1), 2-1 (1), 1-3 (1), 2-3 (2)
- pop 0: 1 = (2, 1 way), 2 = (1, 1 way)
- pop 2: to 1: 1+1 = 2 == 2 → ways1 = 2; to 3: 3 → (3, 1)
- pop 1: to 3: 2+1 = 3 == 3 → ways3 = 1 + 2 = **3** ✅ (0-1-3, 0-2-1-3, 0-2-3)

- **Time:** O((V + E) log V), **Space:** O(V + E)

## Gotchas
- Path lengths can be big: use `int` (64-bit in Go) for distances.
- Take `% mod` every time you add ways.
