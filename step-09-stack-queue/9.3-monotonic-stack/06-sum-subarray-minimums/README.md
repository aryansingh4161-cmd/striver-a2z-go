# Sum of Subarray Minimums

**Problem (LeetCode 907):** Sum `min(subarray)` over **all** subarrays, modulo 10⁹+7.

`[3 1 2 4]` → 17
(mins: 3,1,1,1, 1,1,1, 2,2, 4 → 17)

## Approach 1: All subarrays with a running min
**O(n²)**.

## Approach 2: Contribution of each element (optimal)
Instead of looping over subarrays, ask: **in how many subarrays is `a[i]` the minimum?**
- `left` = how far you can extend to the left while everything is **greater** = `i − PSE index`
- `right` = how far to the right while everything is **greater or equal** = `NSE index − i`

Then `a[i]` is the minimum of `left × right` subarrays, contributing `a[i] × left × right`.

Find PSE / NSE **indexes** with monotonic stacks.

**Ties:** use "strictly smaller" on one side and "smaller or equal" on the other. Otherwise a subarray with two equal minimums is counted twice (or zero times).
- **Time:** O(n), **Space:** O(n)

## Dry run `[3 1 2 4]`
| i | a[i] | left | right | contribution |
|---|---|---|---|---|
| 0 | 3 | 1 | 1 | 3 |
| 1 | 1 | 2 | 3 | 6 |
| 2 | 2 | 1 | 2 | 4 |
| 3 | 4 | 1 | 1 | 4 |
Total **17** ✓
