# Distinct Subsequences

**Problem:** In how many ways can `t` be formed as a **subsequence** of `s`? (Count by positions used.)

```
s = "babgbag", t = "bag" → 5
s = "rabbbit", t = "rabbit" → 3
```

## Recursion
`f(i, j)` = ways to form `t[0..j]` from `s[0..i]`.
- If `j < 0`: all of `t` is matched → **1** way.
- If `i < 0` (and `j ≥ 0`): `s` ran out → **0**.
- If `s[i] == t[j]`: two choices:
  - **use** `s[i]` for `t[j]` → `f(i−1, j−1)`
  - **don't use** it (maybe an earlier `s` char matches `t[j]`) → `f(i−1, j)`
- If they differ: `s[i]` is useless for `t[j]` → `f(i−1, j)`.

```
f(i, j) = f(i−1, j−1) + f(i−1, j)   if s[i] == t[j]
        = f(i−1, j)                 otherwise
```

## Stages
- Recursion: exponential
- Memo / table (shifted: `dp[i][0] = 1`, `dp[0][j>0] = 0`): O(n · m)
- **One row:** loop `j` from **high to low** so `dp[j−1]` is still the previous row. **O(m)** space.

## Dry run (`babgbag` / `bag`, one row `dp[0..3]`, start `[1 0 0 0]`)
| s char | dp after |
|---|---|
| b | 1 1 0 0 |
| a | 1 1 1 0 |
| b | 1 2 1 0 |
| g | 1 2 1 1 |
| b | 1 3 1 1 |
| a | 1 3 4 1 |
| g | 1 3 4 **5** |

Counts can get very large; LeetCode guarantees the answer fits in 32 bits, the test sizes are small.
