# Minimum in Rotated Sorted Array

**Problem (LeetCode 153):** Distinct values, sorted then rotated. Return the minimum in O(log n).

`[4 5 6 7 0 1 2]` → 0

## Approach (optimal)
Again, one half is always sorted. In a sorted half, its **first** element is that half's minimum.
- If the left half `[lo..mid]` is sorted: `a[lo]` is a candidate; the true minimum might still be on the right → `ans = min(ans, a[lo])`, `lo = mid + 1`.
- Else the right half is sorted: `a[mid]` is a candidate → `ans = min(ans, a[mid])`, `hi = mid − 1`.

**Shortcut:** if `a[lo] <= a[hi]`, the whole range is sorted → `a[lo]` is the min of it; stop.

- **Time:** O(log n), **Space:** O(1)

(With duplicates, add the same `a[lo]==a[mid]==a[hi]` shrink trick as [search II](../09-search-rotated-sorted-array-ii/).)
