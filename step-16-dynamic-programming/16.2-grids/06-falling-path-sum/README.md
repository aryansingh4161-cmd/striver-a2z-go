# Minimum / Maximum Falling Path Sum

**Problem:** `n × m` matrix. Start at **any** cell of the first row, end at **any** cell of the last row. From `(i, j)` move to `(i+1, j−1)`, `(i+1, j)` or `(i+1, j+1)` (down-left, down, down-right). Find the minimum (LeetCode 931) or maximum (Striver's version) path sum.

```
2 1 3
6 5 4
7 8 9   → min 13 (1 → 5 → 7 or 1 → 4 → 8), max 17 (3 → 5 → 9)
```

## Recursion
Both ends are free, so pick a direction and try every start. Let `f(i, j)` = best sum of a path from row 0 **ending** at `(i, j)`:
```
f(0, j) = mat[0][j]
f(i, j) = ±∞ if j is off the grid
f(i, j) = mat[i][j] + best( f(i−1, j−1), f(i−1, j), f(i−1, j+1) )
answer  = best over j of f(n−1, j)
```
"best" = min for the minimum version, max for the maximum version. The code takes a flag.

## Stages
- Recursion: O(3ⁿ) per start → O(m · 3ⁿ)
- Memo: O(n · m) states × 3 → **O(n · m)**
- Tabulation: row 0 = matrix row 0; each next row from the previous one.
- Space optimised: only the previous row → **O(m)** space.

## Dry run (min, example)
```
row 0: 2 1 3
row 1: 6+1=7  5+1=6  4+1=5
row 2: 7+6=13 8+5=13 9+5=14   → min 13 ✅
```
