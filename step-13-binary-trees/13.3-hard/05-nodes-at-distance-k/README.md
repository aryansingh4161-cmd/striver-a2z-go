# All Nodes at Distance K

**Problem (LeetCode 863):** Return values of all nodes exactly `k` edges away from `target` (any direction: down, or up through the parent).

```
        3
      /   \
     5     1
    / \   / \
   6   2 0   8
      / \
     7   4
```
target 5, k = 2 → `7 4 1`

## Approach: Turn the tree into an undirected graph, then BFS
1. One traversal to record each node's **parent** (map node → parent).
2. BFS from `target`, moving to left, right **and parent**, with a `visited` set. Stop after `k` levels; the queue holds the answer.
- **Time:** O(n), **Space:** O(n)

The same "parent map + BFS" trick solves [burning tree](../06-burn-tree/).
