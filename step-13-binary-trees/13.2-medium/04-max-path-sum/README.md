# Binary Tree Maximum Path Sum

**Problem (LeetCode 124):** A path is any sequence of connected nodes (each used once, at least one node). Return the maximum sum. Values can be negative.

`[−10 9 20 null null 15 7]` → 42 (15 → 20 → 7)

## Approach: Same shape as the diameter (optimal)
`gain(n)` = the best sum of a path that **starts at `n` and goes down one side**.
```
gain(n):
    if n == nil: return 0
    l = max(0, gain(left))      // a negative branch is better skipped
    r = max(0, gain(right))
    best = max(best, n.val + l + r)   // path bending at n
    return n.val + max(l, r)          // a parent can only use one side
```
Start `best` at the root's value (or −∞) so all-negative trees work: `[−3]` → −3.
- **Time:** O(n), **Space:** O(h)
