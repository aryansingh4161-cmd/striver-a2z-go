# Detect a Cycle in a Directed Graph (BFS / Kahn's)

**Problem:** Does the directed graph contain a cycle? Use BFS.

## Idea
Run [Kahn's algorithm](../02-kahns-algorithm/). Count how many nodes get popped.
- Every node popped had all its prerequisites done.
- Nodes on a cycle (and anything that depends on a cycle) never reach in-degree 0.

So: **cycle ⇔ processed count < V**.

## Dry run
`0 → 1 → 2 → 0`, `2 → 3`
- indeg: 0:1, 1:1, 2:1, 3:1 → queue empty from the start
- processed = 0 < 4 → **cycle** ✅

`0 → 1, 0 → 2, 2 → 1` (diamond, no cycle)
- indeg 0:0 → pop 0, then 2, then 1 → processed 3 = V → no cycle ✅

- **Time:** O(V + E), **Space:** O(V)

Compare with the [DFS version](../../15.2-bfs-dfs/13-cycle-directed-dfs/), which tracks the current recursion path.
