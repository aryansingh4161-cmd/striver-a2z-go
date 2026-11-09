# Find Peak Element

**Problem (LeetCode 162):** A peak is an element strictly greater than its neighbours. Imagine `a[−1] = a[n] = −∞`. Return the index of **any** peak. Neighbours are never equal.

`[1 2 3 1]` → 2
`[1 2 1 3 5 6 4]` → 1 or 5

## Approach 1: Linear scan
Return the first `i` where `a[i] > a[i+1]` (or the last index). **O(n)**.

## Approach 2: Binary search on the slope (optimal)
The array isn't sorted, but look at `mid` and its right neighbour:
- `a[mid] < a[mid+1]` → we're on an **uphill** slope. Keep walking right and you must hit a peak (the end is −∞) → `lo = mid + 1`.
- otherwise → downhill (or at a peak) → a peak is at `mid` or to the left → `hi = mid`.

Loop while `lo < hi`; when they meet, that's a peak.
- **Time:** O(log n), **Space:** O(1)

This shows binary search works whenever you can **throw away half** for sure, not only on sorted arrays.
