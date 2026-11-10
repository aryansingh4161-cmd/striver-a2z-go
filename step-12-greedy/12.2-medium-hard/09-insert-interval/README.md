# Insert Interval

**Problem (LeetCode 57):** Non-overlapping intervals sorted by start. Insert `newInterval`, merging where needed.

`[[1 3] [6 9]]`, new `[2 5]` → `[[1 5] [6 9]]`
`[[1 2] [3 5] [6 7] [8 10] [12 16]]`, new `[4 8]` → `[[1 2] [3 10] [12 16]]`

## Approach: Three phases in one pass (optimal)
1. Copy every interval that ends **before** the new one starts (`end < new.start`).
2. **Merge** every interval that overlaps (`start <= new.end`): `new.start = min(...)`, `new.end = max(...)`. Then add the merged interval.
3. Copy the rest.
- **Time:** O(n), **Space:** O(n) for the answer
