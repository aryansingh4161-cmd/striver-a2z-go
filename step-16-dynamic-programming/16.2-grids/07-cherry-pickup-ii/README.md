# Cherry Pickup II (Ninja and his Friends): 3D DP

**Problem:** A grid of cherries, `n` rows × `m` columns. Robot A starts at `(0, 0)`, robot B at `(0, m−1)`. Each step **both** go down one row, to column `j−1`, `j` or `j+1`. They collect the cherries on cells they visit; if both stand on the same cell, it counts **once**. Both must reach the last row. Maximise the total.

```
3 1 1
2 5 1
1 5 5
2 1 1    → 24  (A: 3,2,5,2  B: 1,5,5,1 → 12 + 12)
```

## Why move them together
Both robots are always on the **same row** (they move down together). So the state is just `(row, colA, colB)`: a 3D DP.

## Recursion
```
f(i, a, b) = −∞                                 if a or b is off the grid
value      = grid[i][a] + (a != b ? grid[i][b] : 0)
f(n−1, a, b) = value                            (last row)
f(i, a, b) = value + max over da, db in {−1, 0, +1} of f(i+1, a+da, b+db)
answer = f(0, 0, m−1)
```
Each step tries 3 × 3 = **9** combinations.

## Stages
- Recursion: O(9ⁿ)
- Memo: `n · m · m` states × 9 → **O(9 · n · m²)**, Space O(n · m²) + stack
- Tabulation: fill from the **last row upward**: `dp[n−1][a][b]` is the base, then `dp[i]` from `dp[i+1]`. Answer `dp[0][0][m−1]`.
- Space optimised: only the row below is needed → two `m × m` layers. **O(m²)** space.

## Dry run (start)
At row 0, A = col 0 (3), B = col 2 (1) → value 4. Best continuation from row 1 is A → col 0 (2), B → col 1 (5)... building up the table gives `dp[0][0][2] = 24` ✅.
