# Reverse Pairs

**Problem (LeetCode 493):** Count pairs `(i, j)` with `i < j` and `a[i] > 2 × a[j]`.

`[1 3 2 3 1]` → 2 (`(3,1)` at indices 1,4 and 3,4)

## Brute
All pairs. **O(n²)**.

## Optimal: Merge sort, but count in a separate pass
It looks like [count inversions](../10-count-inversions/), but the condition `a[i] > 2·a[j]` doesn't line up with the merge comparison `a[i] > a[j]`. So before merging, do a separate **counting** pass on the two sorted halves:
```
r = mid + 1
for l = lo .. mid:
    while r <= hi and a[l] > 2 * a[r]: r++
    count += r - (mid + 1)
```
Because both halves are sorted, `r` only moves forward: the pass is O(n). Then merge normally.
- **Time:** O(n log n), **Space:** O(n)

(In C++/Java, `2 * a[j]` can overflow 32-bit ints; Go's `int` is 64-bit.)
