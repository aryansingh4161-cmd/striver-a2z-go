# Count Occurrences in a Sorted Array

**Problem:** Count how many times `x` appears.

`[2 2 3 3 3 3 4]`, x = 3 → 4

## Approach 1: Linear count
**O(n)**.

## Approach 2: Binary search (optimal)
`count = upperBound(x) − lowerBound(x)` (or `last − first + 1` from the [previous problem](../06-first-last-occurrence/)).
- **Time:** O(log n), **Space:** O(1)
