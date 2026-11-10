# Is the Graph Bipartite?

**Problem (LeetCode 785):** Can the nodes be split into two groups so that every edge connects the two groups? (Equivalently: can you colour the graph with **2 colours** so no edge joins same-coloured nodes?)

## Key fact
A graph is bipartite **iff it has no odd-length cycle**.

## Approach: Try to 2-colour it with BFS (or DFS)
Colour the start node 0. Every neighbour must get the **opposite** colour. If a neighbour already has the **same** colour → not bipartite. Repeat from every uncoloured node (disconnected graphs).
- **Time:** O(n + m), **Space:** O(n)

This is the [M-colouring problem](../../../step-07-recursion/7.3-hard/05-m-coloring/) with m = 2, but no backtracking is needed: each node's colour is forced.
