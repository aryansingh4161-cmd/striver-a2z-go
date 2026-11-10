# Maximal Rectangle (of 1s in a binary matrix)

**Problem (LeetCode 85):** Find the largest rectangle containing only `1`s.

```
1 0 1 0 0
1 0 1 1 1     → 6
1 1 1 1 1
1 0 0 1 0
```

## Approach: Turn each row into a histogram (optimal)
For each row, `height[j]` = number of consecutive `1`s ending at this row in column `j` (reset to 0 on a `0`).
Each row is now a histogram. Run [largest rectangle in histogram](../10-largest-rectangle-histogram/) on it; the answer is the best over all rows.

| row | heights | best rectangle |
|---|---|---|
| 0 | 1 0 1 0 0 | 1 |
| 1 | 2 0 2 1 1 | 3 |
| 2 | 3 1 3 2 2 | **6** |
| 3 | 4 0 0 3 0 | 4 |

- **Time:** O(n × m), **Space:** O(m)
