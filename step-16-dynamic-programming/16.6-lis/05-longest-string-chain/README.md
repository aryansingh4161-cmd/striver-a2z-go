# Longest String Chain

**Problem:** Word `a` is a **predecessor** of `b` if you can insert **exactly one** letter anywhere in `a` to get `b`. A chain is `w1 → w2 → ...` where each is a predecessor of the next. Return the longest chain length.

```
["a" "b" "ba" "bca" "bda" "bdca"] → 4 (a → ba → bda → bdca)
```

## Same LIS shape
1. **Sort words by length** (a predecessor is always exactly 1 shorter, so it comes first).
2. `dp[i]` = longest chain ending at word `i`:
   `dp[i] = 1 + max(dp[j])` over `j < i` where `words[j]` is a predecessor of `words[i]`.

## Checking "predecessor" with two pointers
`a` is a predecessor of `b` iff `len(b) == len(a) + 1` and `a` is a subsequence of `b` where `b` has **one** extra letter:
```
i, j = 0, 0           // i in a, j in b
while j < len(b):
    if i < len(a) and a[i] == b[j]: i++
    j++
return i == len(a) and len(b) == len(a) + 1
```

## Dry run
Sorted: a, b, ba, bca, bda, bdca
| word | dp | from |
|---|---|---|
| a | 1 | |
| b | 1 | |
| ba | 2 | a or b |
| bca | 3 | ba |
| bda | 3 | ba |
| bdca | 4 | bca or bda |
→ **4** ✅

- **Time:** O(n² · L) where L = word length, **Space:** O(n)

(Faster alternative: a map from word → best chain, trying to delete each letter: O(n · L²).)
