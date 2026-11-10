# Minimum Path Sum

**Problem:** A grid of non-negative numbers. Go from top-left to bottom-right moving only right or down. Minimise the **sum of numbers on the path** (both ends included).

```
1 3 1
1 5 1
4 2 1   → 7 (1 → 3 → 1 → 1 → 1)
```

## Recursion
Same shape as unique paths, but "count" becomes "min":
```
f(0, 0) = grid[0][0]
f(i, j) = ∞ if off the grid
f(i, j) = grid[i][j] + min( f(i−1, j), f(i, j−1) )
```
Using ∞ for off-grid makes the first row/column work without special cases: there's only one real option there.

## Stages
- Recursion O(2^(m+n)) → Memo O(m·n) → Table O(m·n) → **one row O(n)** space.

## Dry run (table)
```
1  4  5
2  7  6
6  8  7   → 7 ✅
```
e.g. `dp[1][1] = 5 + min(4, 2) = 7`, `dp[2][2] = 1 + min(6, 8) = 7`.

## Careful
In Go use a big number (like `math.MaxInt / 2`) for ∞ so `∞ + grid` can't overflow, or just skip the missing neighbour (which the table code does).
