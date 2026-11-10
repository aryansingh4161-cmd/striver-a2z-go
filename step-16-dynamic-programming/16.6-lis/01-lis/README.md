# Longest Increasing Subsequence (LIS)

**Problem:** Length of the longest **strictly increasing** subsequence.

```
[10 9 2 5 3 7 101 18] → 4 ([2 3 7 18] or [2 5 7 101])
```

## Approach 1: recursion with "previous index"
To decide whether `nums[i]` can be added, you need to know the **last element taken**. State: `f(i, prev)`.
```
f(n, prev) = 0
f(i, prev) = max( f(i+1, prev),                          // skip nums[i]
                  1 + f(i+1, i)  if prev == −1 or nums[i] > nums[prev] )   // take it
answer = f(0, −1)
```
`prev = −1` means "nothing taken yet". In a table, shift `prev` by +1 so it fits index 0.
- Recursion: O(2ⁿ)
- Memo / table: `n × (n+1)` states → **O(n²)** time, O(n²) space
- Space optimised (two rows of size n+1): O(n)

## Approach 2: `dp[i]` = LIS **ending at** i (simplest O(n²), most useful)
```
dp[i] = 1 + max( dp[j] for j < i with nums[j] < nums[i] ),  or 1 if none
answer = max(dp)
```
This form is the one we extend in the next problems (print LIS, divisible subset, string chain, bitonic, count).
- **Time:** O(n²), **Space:** O(n)

## Dry run (approach 2)
| i | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|---|---|---|---|---|---|---|---|---|
| nums | 10 | 9 | 2 | 5 | 3 | 7 | 101 | 18 |
| dp | 1 | 1 | 1 | 2 | 2 | 3 | 4 | 4 |
max = **4** ✅

Faster: O(n log n) with [binary search](../03-lis-binary-search/).
