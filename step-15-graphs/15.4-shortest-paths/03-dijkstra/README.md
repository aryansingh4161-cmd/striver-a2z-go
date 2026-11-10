# Dijkstra's Algorithm

**Problem:** Weighted graph with **non-negative** weights. Find the shortest distance from `src` to every node.

## The idea in one sentence
Always expand the **closest not-yet-finished** node; with no negative edges, nothing found later can make it any closer.

## How it works
- `dist[src] = 0`, everything else = ∞
- A **min-heap** of `(distance, node)` gives the closest candidate fast.
```
push (0, src)
while heap not empty:
    (d, u) = pop smallest
    if d > dist[u]: continue      // stale entry, we already found better
    for (v, w) in adj[u]:
        if d + w < dist[v]:
            dist[v] = d + w
            push (dist[v], v)
```
The "stale entry" check is how Go/most languages handle it: we can't easily decrease a key inside the heap, so we push a new pair and ignore the old one when it pops.

## Dry run
Edges (undirected): 0-1 (4), 0-2 (4), 1-2 (2), 2-3 (3), 2-4 (1), 2-5 (6), 3-5 (2), 4-5 (3); src 0
| pop | updates | dist |
|---|---|---|
| (0,0) | 1→4, 2→4 | 0 4 4 ∞ ∞ ∞ |
| (4,1) | 2: 6 not better | |
| (4,2) | 3→7, 4→5, 5→10 | 0 4 4 7 5 10 |
| (5,4) | 5→8 | 0 4 4 7 5 8 |
| (7,3) | 5: 9 not better | |
| (8,5) | — | |
| (10,5) | stale, skip | |
→ `[0 4 4 7 5 8]` ✅

- **Time:** O((V + E) log V) with a heap
- **Space:** O(V + E)

## Why negative weights break it
Once a node is popped we treat it as final. A negative edge found later could still lower it, and Dijkstra would never go back. Use [Bellman-Ford](../10-bellman-ford/) then.

## Variants you'll see
- **Set version (C++ `set`)**: lets you erase the old pair instead of skipping stale ones. Same complexity; the heap version is simpler in Go.
- **Path reconstruction**: keep `parent[v] = u` whenever you improve `dist[v]`; walk back from the target. `PathTo` below does this.
- **Why the heap and not a queue?** A normal queue gives the right answer too, but it may revisit nodes many times: much slower.
