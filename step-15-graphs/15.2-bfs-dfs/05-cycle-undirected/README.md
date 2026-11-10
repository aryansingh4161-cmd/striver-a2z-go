# Detect a Cycle in an Undirected Graph (BFS and DFS)

**Problem:** Does the undirected graph contain a cycle?

## Key idea
In an undirected graph every edge goes both ways, so the node you **came from** (the parent) will always look "already visited". A cycle exists only if you reach a visited node that is **not your parent**.

## BFS version
Queue of `(node, parent)`. For each neighbour `v` of `u`:
- not visited → mark, push `(v, u)`
- visited and `v != parent` → **cycle**

## DFS version
`dfs(u, parent)`: for each neighbour `v`: if not visited, recurse (return true if it finds a cycle); else if `v != parent`, cycle.

Run from every unvisited node (the graph may be disconnected).
- **Time:** O(n + m), **Space:** O(n)
