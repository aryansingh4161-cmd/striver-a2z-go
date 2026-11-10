# Count Subsequences with Sum K

**Problem:** Count subsequences (any subset of positions) whose sum is `k`.

`[1 2 1]`, k = 2 → 2 (`[1 1]` and `[2]`)

## Approach: Pick / not pick that **returns a count**
```
count(i, sum):
    if i == n: return 1 if sum == k else 0
    return count(i+1, sum + a[i])    // pick
         + count(i+1, sum)           // not pick
```
No path to store: we only need the number.

**Pruning:** if all numbers are positive and `sum > k`, return 0 early.
- **Time:** O(2ⁿ), **Space:** O(n) depth

This exact recursion becomes the **"count subsets with sum K"** DP problem in Step 16, where memoisation turns O(2ⁿ) into O(n × k).
