# BST Iterator

**Problem (LeetCode 173):** An iterator over a BST in ascending order: `Next()` returns the next smallest value, `HasNext()` says whether one exists. Average O(1) time, O(h) memory.

## Approach 1: Flatten to a slice first
Inorder into a slice, walk an index. O(1) per call but **O(n)** memory.

## Approach 2: Controlled iterative inorder (optimal)
Keep a stack with the **left spine** of what's left to visit.
- **Constructor:** push root and all its left descendants.
- **Next():** pop the top (that's the answer); then push its right child and that child's whole left spine.
- **HasNext():** stack not empty.

It's [iterative inorder](../../../step-13-binary-trees/13.1-traversals/04-iterative-inorder/) paused between steps.
- Each node is pushed and popped once → **amortised O(1)** per `Next`, **O(h)** memory.
