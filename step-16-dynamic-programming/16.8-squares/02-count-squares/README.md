# Count Square Submatrices with All 1s

**Problem:** A binary matrix. How many **square** submatrices contain only 1s?

```
0 1 1 1
1 1 1 1
0 1 1 1   → 15 (10 of size 1, 4 of size 2, 1 of size 3)
```

## Idea: `dp[i][j]` = side of the largest all-ones square whose **bottom-right corner** is `(i, j)`
```
dp[i][j] = 0                                                   if mat[i][j] == 0
dp[i][j] = mat[i][j]                                           in the first row / column
dp[i][j] = 1 + min(dp[i−1][j], dp[i][j−1], dp[i−1][j−1])       otherwise
```
Why the `min` of three? A square of side `s` ending at `(i, j)` needs squares of side `s − 1` ending **above**, **left**, and **diagonally up-left**. The smallest of the three limits how big it can grow.

## Why the sum of `dp` is the answer
If the largest square ending at `(i, j)` has side `s`, then squares of side `1, 2, ..., s` all end there: exactly `s` squares. So:
```
answer = Σ dp[i][j]
```

## Dry run (example)
```
dp:
0 1 1 1
1 1 2 2
0 1 2 3   → sum = 0+1+1+1 + 1+1+2+2 + 0+1+2+3 = 15 ✅
```

- **Time:** O(n · m), **Space:** O(n · m) (or O(m) with one row)

The same table also gives the **largest** square (LeetCode "Maximal Square"): `max(dp)²`.
