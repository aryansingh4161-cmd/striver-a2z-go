# Kth Element of Two Sorted Arrays

**Problem:** Return the k-th smallest element (1-based) of the two sorted arrays combined.

`[2 3 6 7 9]`, `[1 4 8 10]`, k = 5 → 6

## Approach 1: Merge until k
Two pointers, count up to k. **Time O(k), Space O(1)**.

## Approach 2: Partition binary search (optimal)
Same as the [median](../13-median-two-sorted-arrays/), but the left side must hold exactly **k** elements instead of half.

The range for `cut1` changes:
- `cut1 >= max(0, k − n2)`: array 2 can provide at most n2 elements.
- `cut1 <= min(k, n1)`: can't take more than k, or more than n1.

When the cut is correct, the answer is `max(l1, l2)`.
- **Time:** O(log(min(n1, n2))), **Space:** O(1)
