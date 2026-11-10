# Longest Common Substring

**Problem:** Length of the longest string that appears **contiguously** in both strings.

```
a = "abcjklp", b = "acjkp" → 3 ("cjk")
```

## How it differs from LCS
A **substring** must be continuous. So when characters don't match, the run is **broken**: the value resets to 0 instead of taking `max(up, left)`.

`dp[i][j]` = length of the longest common substring **ending exactly** at `a[i−1]` and `b[j−1]`.
```
dp[i][j] = 1 + dp[i−1][j−1]   if a[i−1] == b[j−1]
         = 0                  otherwise
answer = max over all cells (not dp[n][m]!)
```

## Dry run (`abcjklp` / `acjkp`, only the interesting cells)
- c = c → dp = 1
- j = j → dp = 1 + 1 = 2
- k = k → dp = 2 + 1 = **3**
- l vs p → 0, the run is broken

## Stages
- Brute force: try every pair of starting points and extend: O(n · m · min(n, m)).
- Table: O(n · m) time, O(n · m) space.
- Two rows: **O(m)** space.

(There's no natural "recursion first" for this one: the table is the clean way to write it.)
