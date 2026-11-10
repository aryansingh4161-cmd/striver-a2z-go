# Grid Unique Paths

**Problem:** An `m × n` grid. Start top-left, finish bottom-right, moving only **right** or **down**. How many different paths?

```
m = 3, n = 2 → 3 paths: RDD, DRD, DDR
```

## 1. Recursion (go backwards from the target)
`f(i, j)` = number of paths from `(0,0)` to `(i, j)`. You arrive at `(i, j)` either from above `(i−1, j)` or from the left `(i, j−1)`:
```
f(0, 0) = 1
f(i, j) = 0 if i < 0 or j < 0   (off the grid)
f(i, j) = f(i−1, j) + f(i, j−1)
answer = f(m−1, n−1)
```
- **Time:** O(2^(m+n))

## 2. Memo → O(m · n) time, O(m · n) space + stack
## 3. Tabulation → `dp[i][j] = dp[i−1][j] + dp[i][j−1]`, first row/column are all 1
## 4. Space optimisation: one row
`dp[i][j]` needs the cell above (previous row, same column) and the cell to the left (current row). Keep a single row `row[j]` and update left to right: before the update `row[j]` is "above", `row[j−1]` is already "left".
- **Space:** O(n)

## Bonus: the maths answer
Every path is `m−1` downs and `n−1` rights in some order. Choose which of the `m+n−2` moves are downs:
```
answer = C(m + n − 2, m − 1)
```
Compute it with a running product to avoid overflow: `res = res · (n − 1 + i) / i` for `i = 1..m−1`.
- **Time:** O(min(m, n)), **Space:** O(1)

## Dry run (3 × 3 table)
```
1 1 1
1 2 3
1 3 6   → 6 = C(4, 2) ✅
```
