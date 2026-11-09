# Sort a Linked List

**Problem (LeetCode 148):** Sort the list in O(n log n).

`4 → 2 → 1 → 3` → `1 → 2 → 3 → 4`

## Approach 1: Copy to slice, sort, write back
**Time O(n log n), Space O(n)**.

## Approach 2: Merge sort on the list (optimal)
Merge sort suits linked lists: splitting and merging only re-link pointers, no shifting.
1. Find the **first** middle (slow/fast), cut the list into two halves.
2. Sort each half recursively.
3. **Merge** two sorted lists with a dummy node (LeetCode 21).

- **Time:** O(n log n), **Space:** O(log n) recursion stack
