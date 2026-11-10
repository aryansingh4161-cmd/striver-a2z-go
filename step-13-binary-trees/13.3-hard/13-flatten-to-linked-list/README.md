# Flatten a Binary Tree to a Linked List

**Problem (LeetCode 114):** Rearrange the tree in place into a "linked list" using `Right` pointers, in **preorder**, with every `Left` set to nil.

```
    1                1
   / \                \
  2   5      →         2
 / \   \                \
3   4   6                3 → 4 → 5 → 6 (all via Right)
```

## Approach 1: Reverse preorder with a `prev` pointer
Visit nodes in **Right, Left, Root** order (reverse preorder). Each node's `Right` becomes `prev` (the node visited just before), `Left` becomes nil.
- **Time:** O(n), **Space:** O(h) recursion

## Approach 2: Morris-style, O(1) space (optimal)
For each `cur` with a left child: find the rightmost node of the left subtree, attach `cur.Right` there, move the whole left subtree to the right, and set `cur.Left = nil`. Then move to `cur.Right`.
- **Time:** O(n), **Space:** O(1)
