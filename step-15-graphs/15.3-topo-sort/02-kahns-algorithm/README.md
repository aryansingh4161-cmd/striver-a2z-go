# Kahn's Algorithm (Topological Sort with BFS)

**Problem:** Same as before: order the nodes of a DAG so every edge `u → v` has `u` first. This time use BFS.

## Idea: start with nodes nobody depends on
**In-degree** of a node = number of edges coming **into** it = number of things that must happen before it.

- A node with in-degree 0 has no pending prerequisites, so it can go now.
- When we "do" a node, remove its outgoing edges: every neighbour's in-degree drops by 1. Any neighbour that hits 0 is now free.

```
indeg[v] = count of edges into v
queue = all nodes with indeg 0
while queue not empty:
    u = pop
    order.append(u)
    for v in adj[u]:
        indeg[v]--
        if indeg[v] == 0: push v
```

## Dry run
Edges: `5→0, 5→2, 4→0, 4→1, 2→3, 3→1`
- indeg: 0:2, 1:2, 2:1, 3:1, 4:0, 5:0 → queue [4 5]
- pop 4 → order [4]; indeg 0→1, 1→1
- pop 5 → order [4 5]; indeg 0→0 (push), 2→0 (push)
- pop 0 → [4 5 0]
- pop 2 → [4 5 0 2]; 3→0 push
- pop 3 → [4 5 0 2 3]; 1→0 push
- pop 1 → **[4 5 0 2 3 1]** ✅

## Bonus: cycle detection for free
Nodes on a cycle never reach in-degree 0 (they wait on each other forever). So if `len(order) < V`, the graph has a cycle. This is the next problem.

- **Time:** O(V + E), **Space:** O(V)

## DFS vs Kahn
| | DFS | Kahn |
|---|---|---|
| Detects cycles | needs extra onPath array | yes, `len(order) < V` |
| Recursion depth | can be deep | none (queue) |
| Easy to get "smallest label first" order | no | yes, use a min-heap instead of a queue |
