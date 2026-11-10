# Triangle

**Problem:** A triangle of numbers: row `i` has `i + 1` entries. Start at the top; from `(i, j)` you may move to `(i+1, j)` or `(i+1, j+1)`. Find the minimum path sum from top to **any** cell of the bottom row.

```
   2
  3 4
 6 5 7
4 1 8 3     → 11 (2 + 3 + 5 + 1)
```

## The trick: recurse from the top, so the end is free
The start is fixed (top) but the end is **any** bottom cell. If we define `f(i, j)` = min sum from `(i, j)` **down to the bottom**, the answer is just `f(0, 0)` and the many possible endings become the base case:
```
f(n−1, j) = tri[n−1][j]                       (bottom row)
f(i, j)   = tri[i][j] + min( f(i+1, j), f(i+1, j+1) )
answer    = f(0, 0)
```
No bounds trouble either: from row `i` (length `i+1`), both `j` and `j+1` exist in row `i+1`.

## Stages
- Recursion: O(2ⁿ)
- Memo: O(n²) time, O(n²) space + stack
- Tabulation: **start from the bottom row** and fill upward. Answer `dp[0][0]`.
- Space optimised: keep one row (the row below). O(n) space.

## Dry run (tabulation, bottom-up)
```
row 3:  4  1  8  3
row 2:  6+1=7  5+1=6  7+3=10
row 1:  3+6=9  4+6=10
row 0:  2+9 = 11 ✅
```
