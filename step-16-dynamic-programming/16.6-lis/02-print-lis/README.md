# Print the LIS

**Problem:** Return one longest increasing subsequence itself, not just its length.

## Idea: remember where each `dp[i]` came from
Use the "ending at `i`" DP from [LIS](../01-lis/) and add a `parent` array:
```
dp[i] = 1, parent[i] = i
for j < i with nums[j] < nums[i] and dp[j] + 1 > dp[i]:
    dp[i] = dp[j] + 1
    parent[i] = j            // the element before nums[i] in its best LIS
```
Then:
1. Find `last` = index with the largest `dp`.
2. Follow `parent` back until `parent[x] == x`.
3. Reverse.

## Dry run (`[5 4 11 1 16 8]`)
| i | nums | dp | parent |
|---|---|---|---|
| 0 | 5 | 1 | 0 |
| 1 | 4 | 1 | 1 |
| 2 | 11 | 2 | 0 (5) |
| 3 | 1 | 1 | 3 |
| 4 | 16 | 3 | 2 (11) |
| 5 | 8 | 2 | 0 (5) |
Best ends at 4 → 16 → 11 → 5 → reversed **[5 11 16]** ✅

- **Time:** O(n²), **Space:** O(n)
