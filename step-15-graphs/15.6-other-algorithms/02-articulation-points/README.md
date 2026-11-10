# Articulation Points (Cut Vertices)

**Problem:** In an undirected graph, an **articulation point** is a node whose removal (together with its edges) increases the number of connected components. Find all of them.

```
0 - 1 - 2        1 is an articulation point (removing it cuts 0 off)
    |   |        2, 3, 4 are not: they sit on the loop 1-2-3-4-1,
    4 - 3        so removing any one of them leaves the rest connected
```

## Same tools as [bridges](../01-bridges/), one small change
Keep `tin[u]` and `low[u]` from a DFS. Two differences:

1. **For back edges and children, `low` still uses `tin[v]`** (same as before).
2. **The rule changes from `>` to `>=`, and it's about the node `u`:**
   ```
   u (not the DFS root) is an articulation point if some child v has low[v] >= tin[u]
   ```
   `low[v] >= tin[u]` means `v`'s subtree can climb back **at most to u itself**, never above it. Remove `u` and that subtree is cut off. (For bridges we needed `>`: climbing back to `u` itself is enough to keep the edge non-critical. For nodes it isn't, because `u` is the thing being removed.)

3. **The root is special:** it has no "above". The root is an articulation point iff it has **2 or more DFS children** (those subtrees can only meet through the root).

```
dfs(u, parent):
    tin[u] = low[u] = timer++; children = 0
    for v in adj[u]:
        if v == parent: continue
        if !visited[v]:
            dfs(v, u); children++
            low[u] = min(low[u], low[v])
            if parent != -1 and low[v] >= tin[u]: mark u
        else:
            low[u] = min(low[u], tin[v])
    if parent == -1 and children > 1: mark u
```

## Dry run
Edges `0-1, 1-2, 2-3, 3-4, 4-1`
- dfs(0) t0 = 0 → dfs(1) t1 = 1 → dfs(2) t2 = 2 → dfs(3) t3 = 3 → dfs(4) t4 = 4; 4 sees 1 (back edge) → low4 = 1
- low3 = 1, low2 = 1; back at 1: low2 (1) ≥ tin1 (1) → **1 is a cut vertex** ✅
- back at 0 (root): only 1 child → not a cut vertex
- Answer: `[1]`

- **Time:** O(V + E), **Space:** O(V)

## Bridges vs articulation points
| | Bridge (edge u–v) | Articulation point (node u) |
|---|---|---|
| Rule | `low[v] > tin[u]` | `low[v] >= tin[u]` (non-root) |
| Root | no special case | cut iff ≥ 2 DFS children |
