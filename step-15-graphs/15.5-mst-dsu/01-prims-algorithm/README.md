# Prim's Algorithm (Minimum Spanning Tree)

## What is a spanning tree?
For a connected undirected graph with `V` nodes, a **spanning tree** keeps **all V nodes** and exactly **V − 1 edges**, with no cycle, and everything still connected.

A **minimum spanning tree (MST)** is the spanning tree with the **smallest total weight**. Think: connect all houses with cable using the least cable.

```
    2       3
 0 --- 1 ------ 2        MST: 0-1 (2), 1-2 (3), 1-4 (5), 0-3 (6)   total 16
 |   / |      /          (8 and 7 are left out)
6| 8/  |5   /7
 | /   |  /
 3     4
```

## Prim's idea: grow one tree, always add the cheapest edge leaving it
1. Start from any node (0). It's "in the tree".
2. Among all edges with one end in the tree and the other outside, take the **cheapest**. Add that node.
3. Repeat until all nodes are in.

A min-heap of `(weight, node, parent)` gives the cheapest outgoing edge fast.
```
push (0, 0, -1)
while heap:
    (w, u, p) = pop
    if inTree[u]: continue        // already added by a cheaper edge
    inTree[u] = true; sum += w; if p != -1: mstEdges.add(p-u)
    for (v, wt) in adj[u]: if !inTree[v]: push (wt, v, u)
```

## Why the cheapest edge is always safe (the "cut" idea)
Split the nodes into "in tree" and "not yet". Some edge must cross that split in any spanning tree. If the MST used a pricier crossing edge, swap it for the cheapest one: still a spanning tree, and not more expensive. So taking the cheapest crossing edge never hurts.

## Dry run (graph above, start 0)
| pop | add | heap gets |
|---|---|---|
| (0,0) | 0 | (2,1) (6,3) |
| (2,1) | 1, sum 2 | (3,2) (8,3) (5,4) |
| (3,2) | 2, sum 5 | (7,4) |
| (5,4) | 4, sum 10 | |
| (6,3) | 3, sum 16 | |
| rest | already in tree | |
**16** ✅

- **Time:** O(E log E), **Space:** O(V + E)

Compare with [Kruskal's](../03-kruskals-algorithm/), which sorts all edges instead of growing one tree.
