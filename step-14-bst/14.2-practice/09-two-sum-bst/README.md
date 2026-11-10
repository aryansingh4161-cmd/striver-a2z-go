# Two Sum in a BST

**Problem (LeetCode 653):** Are there two different nodes whose values add up to `k`?

## Approach 1: Inorder + two pointers
Inorder gives a sorted slice; use the classic two-pointer 2Sum. **O(n)** time, O(n) space.

## Approach 2: Hash set during any traversal
**O(n)** time and space.

## Approach 3: Two BST iterators (optimal space)
A **forward** iterator (next smallest) and a **backward** iterator (next largest), each using a stack of size O(h). Then run two pointers on them:
- `sum < k` → advance the forward one
- `sum > k` → advance the backward one
- stop when they meet

- **Time:** O(n), **Space:** O(h)
