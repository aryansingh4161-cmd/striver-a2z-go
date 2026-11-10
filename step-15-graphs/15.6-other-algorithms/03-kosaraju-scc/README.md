# Kosaraju's Algorithm (Strongly Connected Components)

**Problem:** In a **directed** graph, a **strongly connected component (SCC)** is a largest group of nodes where every node can reach every other node in the group. Count (or list) the SCCs.

```
0 → 1 → 2 → 0      {0,1,2} is one SCC (a cycle)
2 → 3 → 4 → 5 → 3  {3,4,5} is another
5 → 6 → 7 → 6      {6,7}, and so on
```

## The trick in three steps
1. **DFS on the original graph; record nodes by finish time** (like [topo sort](../../15.3-topo-sort/01-topo-sort-dfs/)). The node that finishes **last** is in a "source" SCC (one nothing else points into).
2. **Reverse every edge.** SCCs stay the same (a cycle reversed is still a cycle), but the edges *between* SCCs flip direction.
3. **DFS on the reversed graph, in decreasing finish time.** Each DFS started in step 3 stays inside exactly one SCC, because the edges that used to leave this SCC now point *into* it, and the SCCs they come from were already visited. Number of DFS starts = number of SCCs.

```
order = finish order of DFS on G
R = reverse(G)
for u in reverse(order):
    if !visited[u]: count++; dfs on R from u (collect the SCC)
```

## Dry run
Edges: 0→1, 1→2, 2→0, 2→3, 3→4, 4→5, 5→3, 5→6, 6→7, 7→6
- Step 1 from 0: finish order `7 6 5 4 3 2 1 0` (deepest finishes first)
- Step 3, take nodes from the end of that list: 0, 1, 2, 3 ...
  - start 0 on reversed graph: reaches 2 (edge 2→0 reversed), 1 → **{0, 1, 2}**
  - next unvisited: 3 → reaches 5, 4 → **{3, 4, 5}**
  - next unvisited: 6 → reaches 7 → **{6, 7}**
- **3 SCCs** ✅

- **Time:** O(V + E) (two DFS + reversing), **Space:** O(V + E)

## Uses
Condensing a graph into a DAG of SCCs, 2-SAT, "which nodes can all reach each other". (Tarjan's SCC algorithm does it in one DFS with `low` values; Kosaraju is easier to remember.)
