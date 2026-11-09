# Search a 2D Matrix

**Problem (LeetCode 74):** Each row is sorted, and each row's first value is greater than the previous row's last. Find `target`.

```
1  3  5  7
10 11 16 20     target 3 → true
23 30 34 60
```

## Approach 1: Scan everything
**O(n × m)**.

## Approach 2: Pick the row, then binary search it
For each row, if `row[0] <= target <= row[m−1]`, binary search that row. **O(n + log m)**.

## Approach 3: Treat it as one sorted array (optimal)
Reading row by row gives one sorted list of `n × m` values. Binary search over index `0 .. n·m − 1` and convert:
`row = idx / m`, `col = idx % m`.
- **Time:** O(log(n × m)), **Space:** O(1)
