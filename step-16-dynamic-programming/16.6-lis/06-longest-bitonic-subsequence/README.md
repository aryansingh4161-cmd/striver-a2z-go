# Longest Bitonic Subsequence

**Problem:** A subsequence is **bitonic** if it first strictly increases and then strictly decreases. Either part may be empty, so a purely increasing or purely decreasing subsequence also counts. Return the longest length.

```
[1 11 2 10 4 5 2 1] → 6 ([1 2 10 4 2 1])
```

## Idea: LIS from both sides
For every index `i` as the **peak**:
- `inc[i]` = longest increasing subsequence **ending** at `i` (scan left → right, normal [LIS](../01-lis/)).
- `dec[i]` = longest decreasing subsequence **starting** at `i` = LIS ending at `i` when scanning right → left.
- Bitonic with peak `i`: `inc[i] + dec[i] − 1` (the peak is counted twice).

```
answer = max over i of inc[i] + dec[i] − 1
```

## Dry run
| i | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|---|---|---|---|---|---|---|---|---|
| nums | 1 | 11 | 2 | 10 | 4 | 5 | 2 | 1 |
| inc | 1 | 2 | 2 | 3 | 3 | 4 | 2 | 1 |
| dec | 1 | 5 | 2 | 4 | 3 | 3 | 2 | 1 |
| sum − 1 | 1 | 6 | 3 | **6** | 5 | **6** | 3 | 1 |
→ **6** ✅

- **Time:** O(n²), **Space:** O(n)
