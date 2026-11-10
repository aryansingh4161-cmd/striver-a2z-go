# Edit Distance

**Problem:** Minimum operations to convert `a` into `b`. Allowed operations on `a`: **insert** a character, **delete** a character, **replace** a character.

```
"horse" → "ros"        : 3 (horse → rorse → rose → ros)
"intention" → "execution": 5
```

## Recursion (compare last characters)
`f(i, j)` = min ops to turn `a[0..i]` into `b[0..j]`.
- `i < 0`: `a` is empty → insert all remaining `j + 1` chars of `b`.
- `j < 0`: `b` is empty → delete all remaining `i + 1` chars of `a`.
- `a[i] == b[j]`: nothing to do → `f(i−1, j−1)`.
- Otherwise, `1 +` the best of:
  - **insert** `b[j]` at the end of `a` → now `b[j]` is matched: `f(i, j−1)`
  - **delete** `a[i]` → `f(i−1, j)`
  - **replace** `a[i]` with `b[j]` → both matched: `f(i−1, j−1)`

```
f(i, j) = f(i−1, j−1)                                    if a[i] == b[j]
        = 1 + min(f(i, j−1), f(i−1, j), f(i−1, j−1))     otherwise
```

## Shifted table
`dp[i][0] = i` (delete everything), `dp[0][j] = j` (insert everything).

## Stages
- Recursion: O(3^(n+m))
- Memo / table: O(n · m)
- Two rows: **O(m)** space (remember `cur[0] = i` at the start of each row!)

## Dry run (`horse` → `ros`)
```
      ""  r  o  s
""     0  1  2  3
h      1  1  2  3
o      2  2  1  2
r      3  2  2  2
s      4  3  3  2
e      5  4  4  3   → 3 ✅
```
