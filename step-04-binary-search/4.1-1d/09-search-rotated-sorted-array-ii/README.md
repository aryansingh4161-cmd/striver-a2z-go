# Search in Rotated Sorted Array II (with duplicates)

**Problem (LeetCode 81):** Same as [part I](../08-search-rotated-sorted-array/) but values can repeat. Return true/false.

`[2 5 6 0 0 1 2]`, target 0 → true

## What breaks
With duplicates you can get `a[lo] == a[mid] == a[hi]`, e.g. `[3 1 2 3 3 3 3]`. Now you **can't tell** which half is sorted.

## Fix
In exactly that case, shrink both ends: `lo++`, `hi--`, and continue. Those values equal `a[mid]`, which we already know isn't the target, so nothing is lost.

- **Time:** O(log n) on average, **O(n) worst case** (e.g. all elements equal except one)
- **Space:** O(1)
