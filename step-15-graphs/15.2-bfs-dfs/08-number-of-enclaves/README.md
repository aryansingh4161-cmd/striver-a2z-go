# Number of Enclaves

**Problem (LeetCode 1020):** Grid of `0` (sea) and `1` (land). Count land cells from which you **can't walk off the border** (4 directions).

```
0 0 0 0
1 0 1 0     → 3
0 1 1 0
0 0 0 0
```

## Approach: Same reverse trick as surrounded regions
BFS/DFS from every land cell on the **border**, marking everything reachable. Count the land cells **not** marked.
- **Time:** O(n × m), **Space:** O(n × m)
