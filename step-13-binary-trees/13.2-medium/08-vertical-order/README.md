# Vertical Order Traversal

**Problem (LeetCode 987):** Give each node a position: root at (row 0, col 0); left child (row+1, col−1); right child (row+1, col+1). Return nodes column by column from left to right. Inside a column, sort by **row**; nodes at the same row and column are sorted by **value**.

`[3 9 20 null null 15 7]` → `[[9] [3 15] [20] [7]]`

## Approach
1. Traverse (BFS or DFS) and record `(col, row, val)` for every node.
2. Sort by `col`, then `row`, then `val`.
3. Group consecutive entries with the same `col`.
- **Time:** O(n log n), **Space:** O(n)

(In C++, Striver uses `map<col, map<row, multiset>>`. Go has no sorted map, so "collect, then sort" is the clean Go version.)
