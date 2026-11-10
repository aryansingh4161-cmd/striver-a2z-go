# Replace Each Element by Its Rank

**Problem (LeetCode 1331):** Replace every value by its rank: the smallest value gets 1, equal values get the same rank.

`[20 15 26 2 98 6]` → `[4 3 5 1 6 2]`
`[40 10 20 40]` → `[3 1 2 3]`

## Approach 1: For each element, count smaller distinct values
**O(n²)**.

## Approach 2: Sort + map (optimal)
1. Copy and sort the values.
2. Walk the sorted copy, giving rank 1, 2, 3… to each **new** value (skip duplicates) in a map.
3. Replace each original value with `rank[value]`.

A min-heap can replace the sort (pop in order, assign ranks): same O(n log n).
- **Time:** O(n log n), **Space:** O(n)
