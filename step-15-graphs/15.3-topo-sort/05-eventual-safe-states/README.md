# Find Eventual Safe States

**Problem:** In a directed graph, a node is **terminal** if it has no outgoing edges. A node is **safe** if *every* path starting from it ends at a terminal node (it can never get stuck in a cycle). Return all safe nodes in sorted order.

```
0 → 1 → 3 (terminal? no: 3 → 0)
0 → 2 → 5 (terminal)
2 → 5, 4 → 5, 6 (terminal)
```
In LeetCode's example `graph = [[1,2],[2,3],[5],[0],[5],[],[]]` the answer is `[2 4 5 6]`.

## Approach 1: DFS with path tracking
A node is **unsafe** if it's on a cycle or can reach one. Run the [directed-cycle DFS](../../15.2-bfs-dfs/13-cycle-directed-dfs/): when a DFS from `u` finishes **without** finding a cycle, `u` is safe. If a cycle is found, we return early and leave `onPath` set, so every node on that path stays "unsafe".

- **Time:** O(V + E), **Space:** O(V)

## Approach 2: Reverse the graph + Kahn's (recommended, easy to reason about)
Terminal nodes are safe. A node becomes safe once **all** its outgoing edges lead to safe nodes. That's Kahn's algorithm on the **reversed** graph:
1. Reverse every edge: `u → v` becomes `v → u`.
2. In the reversed graph, in-degree of `u` = out-degree of `u` in the original.
3. Start the queue with terminal nodes (out-degree 0). Pop `v`; for each `u` that pointed to `v`, decrease its count; when it hits 0, `u` is safe.
4. Everything popped is safe. Sort it.

Nodes on cycles never reach 0, so they are left out automatically.

## Dry run (approach 2)
Out-degrees: 0:2, 1:2, 2:1, 3:1, 4:1, 5:0, 6:0 → queue [5 6]
- pop 5: who points to 5? 2 and 4 → both drop to 0 → push 2, 4
- pop 6: nobody
- pop 2: 0 → 1, 1 → 1
- pop 4: nobody
- Safe: {5, 6, 2, 4} → sorted **[2 4 5 6]** ✅ (0, 1, 3 are on the cycle 0→1→3→0)

- **Time:** O(V + E) (+ O(V) to collect in order), **Space:** O(V + E)
