# Network Delay Time

**Problem:** `times[i] = [u, v, w]`: a signal takes `w` time to go from node `u` to `v` (directed). Nodes are `1..n`. A signal starts at `k`. How long until **all** nodes receive it? `-1` if some node never does.

## Idea
The signal reaches node `v` at time = shortest distance from `k` to `v`. Everyone has it when the **last** (farthest) node gets it:
```
answer = max over v of dist[v]   (or -1 if any dist is ∞)
```
Distances with non-negative weights → [Dijkstra](../03-dijkstra/).

## Dry run
`times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2`
- dist from 2: node1 = 1, node3 = 1, node4 = 2
- max = **2** ✅

- **Time:** O((V + E) log V), **Space:** O(V + E)

## Edge cases
- Node labels start at 1: shift by -1 (or make arrays of size n+1).
- An unreachable node → -1.
