# Shortest Common Supersequence

**Problem:** Return the shortest string that has **both** `a` and `b` as subsequences.

```
a = "brute", b = "groot" → "bgruoote" (length 8; other answers of length 8 exist)
```

## Length first
Characters in the [LCS](../01-lcs/) can be **shared**; everything else must appear separately:
```
length = n + m − LCS
```
`brute` / `groot`: LCS = "rt" (2) → 5 + 5 − 2 = 8.

## Building the string from the LCS table
Walk back from `(n, m)` like [Print LCS](../02-print-lcs/), but **don't throw away** the characters you skip:
```
while i > 0 and j > 0:
    if a[i−1] == b[j−1]: add it once; i--; j--        // shared char
    else if dp[i−1][j] > dp[i][j−1]: add a[i−1]; i--  // a's char, only once
    else: add b[j−1]; j--                             // b's char
add whatever is left of a (i > 0) and of b (j > 0)
reverse
```

## Dry run (shortened)
From the end: `e` (from a), `t` shared, `o`, `o` (from b), `u` (from a), `r` shared, `g` (from b), `b` (from a) → reversed: **"bgruoote"** ✅ (length 8, contains b-r-u-t-e and g-r-o-o-t)

- **Time:** O(n · m), **Space:** O(n · m)
