# Sort a K-Sorted (Nearly Sorted) Array

**Problem:** Each element is at most `k` positions away from its sorted position. Sort it.

`[6 5 3 2 8 10 9]`, k = 3 → `[2 3 5 6 8 9 10]`

## Approach 1: Normal sort
**O(n log n)**: ignores the extra information.

## Approach 2: Min-heap of size k+1 (optimal)
The smallest remaining element must be within the next `k+1` elements. So keep a window of `k+1` elements in a min-heap; pop the smallest into the output, push the next element.
- **Time:** O(n log k), **Space:** O(k)

When k is small, this beats a full sort.
