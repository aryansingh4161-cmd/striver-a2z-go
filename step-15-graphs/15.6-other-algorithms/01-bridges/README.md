# Bridges in a Graph (Tarjan's Algorithm)

**Problem:** In an undirected connected graph, a **bridge** (critical connection) is an edge whose removal **disconnects** the graph. Find all bridges.

```
0 - 1 - 2          bridges: 1-3   (0,1,2 form a triangle; 3 hangs off 1)
 \     /
   ---             also: 3-4 if 4 hangs off 3
    1 - 3 - 4
```
LeetCode example: `n = 4`, edges `0-1, 1-2, 2-0, 1-3` → bridge `[1, 3]`.

## Brute force
Remove each edge, check if the graph is still connected (BFS). O(E · (V + E)).

## Tarjan's idea: two timestamps per node
Do one DFS. For every node keep:
- `tin[u]`: the time (step number) when DFS first reached `u`.
- `low[u]`: the **smallest `tin`** reachable from `u`'s DFS subtree using tree edges down **plus at most one back edge** (an edge to an already-visited node), not using the edge back to the parent.

**Bridge rule:** tree edge `u → v` (v is a child) is a bridge iff
```
low[v] > tin[u]
```
Meaning: nothing in `v`'s subtree can climb back to `u` or above without that edge. Cut it and `v`'s subtree falls off.

## The DFS
```
dfs(u, parent):
    tin[u] = low[u] = timer++
    for v in adj[u]:
        if v == parent: continue
        if v not visited:
            dfs(v, u)
            low[u] = min(low[u], low[v])       // child's best climb is mine too
            if low[v] > tin[u]: bridge (u, v)
        else:
            low[u] = min(low[u], tin[v])       // back edge
```

## Dry run (`0-1, 1-2, 2-0, 1-3`)
- dfs(0): tin0=0
  - dfs(1): tin1=1
    - dfs(2): tin2=2; neighbour 0 visited (not parent) → low2 = 0
    - back at 1: low1 = min(1, 0) = 0; low2 (0) > tin1 (1)? no
    - dfs(3): tin3=3, low3=3; back at 1: low3 (3) > tin1 (1) → **bridge 1-3** ✅
  - back at 0: low1 (0) > tin0 (0)? no

- **Time:** O(V + E), **Space:** O(V)

## Careful
- Skip the parent **once by edge**, not by node, if the graph can have parallel edges (two edges between the same pair are never bridges). The code skips the parent edge by its index.
