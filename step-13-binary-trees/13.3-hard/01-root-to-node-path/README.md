# Root to Node Path

**Problem:** Return the values on the path from the root to the node with value `x` (values are unique).

`[1 2 3 4 5 null null null null 6 7]`, x = 7 → `[1 2 5 7]`

## Approach: DFS with backtracking
```
path(n):
    if n == nil: return false
    add n.val to path
    if n.val == x: return true
    if path(n.left) or path(n.right): return true
    remove n.val (backtrack)
    return false
```
- **Time:** O(n), **Space:** O(h)

**Bonus (LeetCode 257):** all root-to-leaf paths: same DFS, save a copy of the path at each leaf. Both are in `solution.go`.
