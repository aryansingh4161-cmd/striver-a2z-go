# Partition a Set into Two Subsets with Minimum Absolute Sum Difference

**Problem:** Split `arr` (positive integers) into two groups so that `|sum1 − sum2|` is as small as possible. Return that difference.

```
[1 6 11 5] → 1 ({1, 5, 6} = 12, {11} = 11)
```

## Idea: which sums can one group have?
If group 1 has sum `s1`, group 2 has `total − s1`, and the difference is `|total − 2·s1|`.

The [Subset Sum](../01-subset-sum/) table answers "is sum `t` possible?" for **every** `t` from `0` to `total` at once: it's the last row `dp[n−1][t]`. So:
1. Build the subset-sum row for target = `total`.
2. For every reachable `s1` (only need `s1 ≤ total / 2`), compute `total − 2·s1`. Take the minimum.

## Dry run (`[1 2 7]`, total 10)
Reachable sums: 0, 1, 2, 3, 7, 8, 9, 10
| s1 ≤ 5 | diff = 10 − 2·s1 |
|---|---|
| 0 | 10 |
| 1 | 8 |
| 2 | 6 |
| 3 | **4** |
Answer **4** ({1, 2} vs {7}) ✅

- **Time:** O(n · total), **Space:** O(total)

## Note
This only works for **non-negative** numbers (the table index is the sum). With negatives, use "meet in the middle" (split into halves, sort, binary search), which is a different technique.
