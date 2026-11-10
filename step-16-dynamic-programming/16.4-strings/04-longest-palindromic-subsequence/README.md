# Longest Palindromic Subsequence

**Problem:** Length of the longest subsequence of `s` that reads the same forwards and backwards.

```
"bbbab"    → 4 ("bbbb")
"bbabcbcab" → 7 ("babcbab")
```

## The trick: LCS with the reverse
A palindrome reads the same backwards. So a palindromic subsequence of `s` is also a subsequence of `reverse(s)`, and the longest common one is the longest palindrome:
```
LPS(s) = LCS(s, reverse(s))
```
Reuse [LCS](../01-lcs/). Done.

## Direct interval DP (alternative)
`f(i, j)` = LPS inside `s[i..j]`:
```
f(i, i) = 1
f(i, j) = 2 + f(i+1, j−1)          if s[i] == s[j]
        = max(f(i+1, j), f(i, j−1)) otherwise
```
Same O(n²). The test checks both agree.

## Dry run (`bbbab` vs reverse `babbb`)
LCS = "bbbb"... (b,b,b,b) → **4** ✅

- **Time:** O(n²), **Space:** O(n) with the two-row LCS
