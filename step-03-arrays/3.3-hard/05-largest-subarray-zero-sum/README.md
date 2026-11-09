# Largest Subarray with Sum 0

**Problem:** Return the length of the longest subarray whose sum is 0.

`[15 -2 2 -8 1 7 10 23]` → 5 (`[-2 2 -8 1 7]`)

## Brute
All subarrays with a running sum. **O(n²)**.

## Optimal: prefix sum + first index
This is [longest subarray with sum K](../../3.1-easy/14-longest-subarray-sum-k/) with `k = 0`:
if the same prefix sum appears at indices `j` and `i`, the part between them sums to 0 → length `i − j`.
```
first = map{}
sum = 0
for i, x in a:
    sum += x
    if sum == 0: best = i + 1
    else if sum in first: best = max(best, i - first[sum])
    else: first[sum] = i
```
- **Time:** O(n), **Space:** O(n)
