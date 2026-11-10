# LCA in a BST

**Problem (LeetCode 235):** Lowest common ancestor of `p` and `q` in a BST.

## Approach: Use the ordering (optimal)
Starting at the root:
- both values **smaller** than the node → LCA is on the left
- both **greater** → on the right
- otherwise (they split, or one equals the node) → **this node is the LCA**

- **Time:** O(h), **Space:** O(1)

Much simpler than the [general binary tree LCA](../../../step-13-binary-trees/13.3-hard/02-lowest-common-ancestor/), because the values tell you where to go.
