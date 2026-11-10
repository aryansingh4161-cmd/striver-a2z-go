# Kth Smallest / Kth Largest in a BST

**Problem (LeetCode 230):** k-th smallest value (1-based).

## Approach 1: Inorder into a slice
Inorder gives sorted order; return `a[k−1]`. **O(n)** time and space.

## Approach 2: Inorder with a counter, stop early (optimal)
Count nodes as you visit them in inorder; stop at the k-th. **Time O(h + k), Space O(h)**.

**Kth largest:** reverse inorder (Right, Root, Left), or the `(n − k + 1)`-th smallest.

(Morris inorder makes it O(1) extra space.)
