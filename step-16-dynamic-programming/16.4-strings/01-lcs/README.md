# Longest Common Subsequence (LCS)

**Problem:** Length of the longest sequence that is a **subsequence** of both strings (characters in order, not necessarily next to each other).

```
a = "abcde", b = "ace" → 3 ("ace")
a = "abc",   b = "def" → 0
```

This is the most important string DP. Many problems below are LCS with a twist.

## Recursion: compare the last characters
`f(i, j)` = LCS of `a[0..i]` and `b[0..j]`.
- If `a[i] == b[j]`: that character can be in the LCS → `1 + f(i−1, j−1)`.
- Otherwise: drop one of them → `max(f(i−1, j), f(i, j−1))`.
- Base: if either index is `< 0` (an empty string), LCS = 0.

```
f(i, j) = 0                               if i < 0 or j < 0
        = 1 + f(i−1, j−1)                 if a[i] == b[j]
        = max(f(i−1, j), f(i, j−1))       otherwise
```

## Shifting indexes by one (table trick)
Base case `i < 0` can't be an array index. So the table uses `dp[i][j]` = LCS of the first `i` chars of `a` and first `j` chars of `b`. Then row 0 and column 0 are the empty-string base cases (all 0), and we compare `a[i−1]` with `b[j−1]`.

## Stages
| | Time | Space |
|---|---|---|
| Recursion | O(2^(n+m)) | O(n + m) |
| Memo | O(n · m) | O(n · m) + stack |
| Table | O(n · m) | O(n · m) |
| Two rows | O(n · m) | **O(m)** |

## Dry run (`abcde` vs `ace`)
```
      ""  a  c  e
""     0  0  0  0
a      0  1  1  1
b      0  1  1  1
c      0  1  2  2
d      0  1  2  2
e      0  1  2  3   → 3 ✅
```
