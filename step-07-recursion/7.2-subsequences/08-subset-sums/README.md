# Subset Sums

**Problem:** Return the sums of **all** 2ⁿ subsets, sorted.

`[2 3]` → `0 2 3 5`

## Approach: Pick / not pick carrying the sum
```
f(i, sum):
    if i == n: save sum; return
    f(i+1, sum + a[i])   // pick
    f(i+1, sum)          // not pick
sort the results
```
- **Time:** O(2ⁿ + 2ⁿ log 2ⁿ) = O(2ⁿ × n) with the sort, **Space:** O(2ⁿ) for the answer
