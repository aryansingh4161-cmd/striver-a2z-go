# Number of Islands

**Problem (LeetCode 200):** Grid of `'1'` (land) and `'0'` (water). An island is land connected **horizontally or vertically**. Count islands.

```
1 1 0 0 0
1 1 0 0 0     → 3
0 0 1 0 0
0 0 0 1 1
```
(Striver's GFG version also counts **diagonal** connections: just use 8 directions.)

## Approach: Flood each island once
Scan the grid. At each unvisited land cell, `count++` and run DFS/BFS to mark the whole island as visited.
```
dirs = up, down, left, right
dfs(r, c): if out of grid or water or visited → return
           mark; dfs on the 4 neighbours
```
- **Time:** O(n × m), **Space:** O(n × m) (visited / recursion)
