# Kth Largest Element in a Stream

**Problem (LeetCode 703):** Numbers arrive one by one. After each `Add(x)`, return the current k-th largest.

k = 3, start `[4 5 8 2]`; add 3 → 4, add 5 → 5, add 10 → 5, add 9 → 8, add 4 → 8

## Approach: Min-heap of size k (optimal)
Keep only the k largest numbers seen so far in a min-heap. On `Add(x)`: push x; if size > k, pop. The top is the answer.
- **Add:** O(log k), **Space:** O(k)

Same idea as [kth largest in an array](../../11.2-medium/01-kth-largest/), but the heap lives between calls.
