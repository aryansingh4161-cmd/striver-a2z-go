# Reverse Nodes in K-Group

**Problem (LeetCode 25):** Reverse the list in groups of `k`. A last group with fewer than `k` nodes stays as it is.

`1 2 3 4 5`, k = 2 → `2 1 4 3 5`
`1 2 3 4 5`, k = 3 → `3 2 1 4 5`

## Approach (optimal, iterative)
Use a dummy node and a pointer `groupPrev` = the node just before the current group.
1. From `groupPrev`, walk `k` steps to find `kth`, the last node of the group. If there are fewer than `k` nodes left, stop.
2. Remember `groupNext = kth.Next`.
3. Reverse the group: standard reversal, but start `prev` at `groupNext` so the reversed group's tail already points onward.
4. Reconnect: the old first node of the group is now its last. `groupPrev.Next = kth` (new first), and `groupPrev` = old first node.

- **Time:** O(n) (each node is visited about twice), **Space:** O(1)

## Picture for k = 3
```
dummy → [1 → 2 → 3] → 4 → 5
dummy → [3 → 2 → 1] → 4 → 5     groupPrev moves to 1
         only 2 nodes left (< 3) → stop
```
