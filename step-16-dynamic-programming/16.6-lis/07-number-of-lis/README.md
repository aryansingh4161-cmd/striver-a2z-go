# Number of Longest Increasing Subsequences

**Problem:** How many different (by index) longest strictly increasing subsequences are there?

```
[1 3 5 4 7] → 2 ([1 3 4 7] and [1 3 5 7])
[2 2 2 2 2] → 5 (each single element)
```

## Two arrays
- `dp[i]` = length of the LIS **ending** at `i` (as in [LIS](../01-lis/)).
- `cnt[i]` = how many LIS of that length end at `i`.

For each `j < i` with `nums[j] < nums[i]`:
- `dp[j] + 1 > dp[i]`: found a **longer** way → `dp[i] = dp[j] + 1`, `cnt[i] = cnt[j]` (start counting fresh)
- `dp[j] + 1 == dp[i]`: another way of the **same** length → `cnt[i] += cnt[j]`

Answer: sum of `cnt[i]` over all `i` with `dp[i] == max(dp)`.

## Dry run (`[1 3 5 4 7]`)
| i | nums | dp | cnt |
|---|---|---|---|
| 0 | 1 | 1 | 1 |
| 1 | 3 | 2 | 1 |
| 2 | 5 | 3 | 1 |
| 3 | 4 | 3 | 1 |
| 4 | 7 | 4 | 1 (from 5) + 1 (from 4) = 2 |
max dp = 4 → total count **2** ✅

- **Time:** O(n²), **Space:** O(n)
