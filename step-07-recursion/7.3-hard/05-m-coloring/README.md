# M-Coloring Problem

**Problem:** Undirected graph with `n` nodes. Can you colour every node with one of `m` colours so that **no edge connects two nodes of the same colour**?

4 nodes, edges `0-1 1-2 2-3 3-0 0-2`, m = 3 → true

## Approach: Backtracking node by node
```
f(node):
    if node == n: return true
    for color = 1 .. m:
        if no neighbour of node already has this color:
            colour[node] = color
            if f(node + 1): return true
            colour[node] = 0      // backtrack
    return false
```
- **Time:** O(mⁿ) worst case, **Space:** O(n) for colours + depth

Store the graph as an **adjacency list** so checking neighbours costs O(degree). (More on graph storage in Step 15.)

m = 2 is the special case "**is the graph bipartite?**", which BFS solves in O(V + E) (Step 15).
