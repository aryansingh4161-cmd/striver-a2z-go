# Search Insert Position

**Problem (LeetCode 35):** Sorted array of distinct values. Return the index of `target` if present; otherwise the index where it should be inserted to keep the order.

`[1 3 5 6]`: target 5 → 2, target 2 → 1, target 7 → 4

## Key insight
"Where it should be inserted" = the first index with `a[i] >= target`. That's just **lower bound**.
- **Time:** O(log n), **Space:** O(1)
