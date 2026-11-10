# Flood Fill

**Problem (LeetCode 733):** Starting at `(sr, sc)`, change that pixel and every 4-directionally connected pixel of the **same original colour** to `color`.

```
1 1 1        2 2 2
1 1 0   →    2 2 0      (start (1,1), color 2)
1 0 1        2 0 1
```

## Approach: DFS from the start cell
Remember `orig = image[sr][sc]`. If `orig == color`, return immediately (otherwise the DFS would loop forever, since painted cells still look "original").
DFS: if inside the grid and the cell equals `orig`, paint it and recurse on 4 neighbours. The new colour itself acts as the "visited" mark.
- **Time:** O(n × m), **Space:** O(n × m) recursion worst case
