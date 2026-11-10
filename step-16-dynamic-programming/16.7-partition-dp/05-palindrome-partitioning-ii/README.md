# Palindrome Partitioning II (Minimum Cuts)

**Problem:** Cut string `s` into pieces that are **all palindromes**. Minimum number of cuts?

```
"aab"     → 1 ("aa" | "b")
"ababbbabbababa" → 3 ("a" | "babbbab" | "b" | "ababa")
"a"       → 0
```

## Front partition (a 1D variant of partition DP)
Instead of splitting a range in two, **cut off the first piece** and recurse on the rest:
```
f(i) = fewest palindrome pieces for s[i..]
f(n) = 0
f(i) = min over j = i..n−1 where s[i..j] is a palindrome of  1 + f(j+1)
answer = f(0) − 1          (pieces − 1 = cuts)
```
Counting pieces avoids an off-by-one at the end: the last piece doesn't need a cut after it.

## Make the palindrome check O(1)
Checking `s[i..j]` each time costs O(n), giving O(n³) total. Precompute `pal[i][j]`:
```
pal[i][j] = s[i] == s[j] && (j − i < 2 || pal[i+1][j−1])
```
(fill `i` from the end so `pal[i+1][...]` is ready)

## Stages
- Recursion: exponential
- Memo / table on `f(i)`: O(n²) with the `pal` table, Space O(n²)

## Dry run (`aab`)
- f(3) = 0
- f(2): "b" → 1 + f(3) = 1
- f(1): "a" → 1 + f(2) = 2; "ab" no → 2
- f(0): "a" → 1 + f(1) = 3; "aa" → 1 + f(2) = 2; "aab" no → 2
- cuts = 2 − 1 = **1** ✅
