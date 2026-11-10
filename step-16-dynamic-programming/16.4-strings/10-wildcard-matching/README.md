# Wildcard Matching

**Problem:** Does pattern `p` match the **whole** string `s`?
- `?` matches exactly one character.
- `*` matches **any sequence** of characters (including empty).

```
s = "aa",    p = "a"    → false
s = "aa",    p = "*"    → true
s = "adceb", p = "*a*b" → true
s = "acdcb", p = "a*c?b" → false
```

## Recursion (from the ends)
`f(i, j)` = does `p[0..i]` match `s[0..j]`?
- Both empty (`i < 0`, `j < 0`) → true.
- Pattern empty, string not → false.
- String empty, pattern not → true only if the remaining pattern is **all `*`** (each matches empty).
- `p[i] == s[j]` or `p[i] == '?'` → `f(i−1, j−1)`.
- `p[i] == '*'` → two choices:
  - `*` matches **nothing** → `f(i−1, j)`
  - `*` **eats** `s[j]` and stays available for more → `f(i, j−1)`
- Otherwise → false.

## Shifted table
- `dp[0][0] = true`
- `dp[0][j > 0] = false` (empty pattern can't match a non-empty string)
- `dp[i][0] = true` only if `p[0..i−1]` is all `*`

## Stages
- Recursion: exponential
- Memo / table: O(n · m)
- Two rows: **O(m)** space (where `m = len(s)`)

## Dry run (`p = "*a*b"`, `s = "adceb"`)
The first `*` eats nothing, `a` matches `a`, the second `*` eats `dce`, `b` matches `b` → **true** ✅. The table finds this automatically by trying both `*` options everywhere.
