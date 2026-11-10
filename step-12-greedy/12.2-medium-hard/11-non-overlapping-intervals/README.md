# Non-overlapping Intervals

**Problem (LeetCode 435):** Remove the minimum number of intervals so the rest don't overlap. (Touching like `[1 2]` and `[2 3]` is fine.)

`[[1 2] [2 3] [3 4] [1 3]]` → 1 (remove `[1 3]`)

## Key insight
Minimum removed = n − **maximum number of non-overlapping intervals you can keep**. Keeping the maximum is exactly [N meetings in one room](../01-n-meetings/) (activity selection).

## Greedy rule
Sort by **end**. Keep an interval if its start ≥ the end of the last kept one; otherwise remove it.
- **Time:** O(n log n), **Space:** O(1) besides sorting

Note the comparison: here touching is allowed (`>=`); in N meetings it isn't (`>`). Always check the problem's rule.
