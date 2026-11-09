# Longest Subarray with Sum K (positives and zeros only)

**Problem:** All elements are ≥ 0. Return the length of the longest **subarray** (contiguous part) whose sum is exactly `k`.

`[1 2 3 1 1 1 1 4 2 3]`, k = 3 → 3 (`[1 1 1]`)

## Approach 1: All subarrays (brute)
Fix the start `i`, extend the end `j`, keep a running sum.
**Time O(n²), Space O(1)**.

## Approach 2: Prefix sum + hash map (better)
Works even with negatives. See the [next problem](../14-longest-subarray-sum-k/).
**Time O(n), Space O(n)**.

## Approach 3: Sliding window / two pointers (optimal for non-negatives)
Because all numbers are ≥ 0, **growing the window never decreases the sum, and shrinking never increases it**. So:
- add `a[right]` to `sum`,
- while `sum > k`: remove `a[left]`, `left++`,
- if `sum == k`: update the answer with `right − left + 1`.

```
left = 0, sum = 0
for right = 0 .. n-1:
    sum += a[right]
    while sum > k and left <= right:
        sum -= a[left]; left++
    if sum == k: best = max(best, right-left+1)
```
- **Time:** O(2n) = O(n) (each element enters and leaves once), **Space:** O(1)

## Why it fails with negatives
With negatives, shrinking can **increase** the sum, so "shrink while too big" is no longer correct.
