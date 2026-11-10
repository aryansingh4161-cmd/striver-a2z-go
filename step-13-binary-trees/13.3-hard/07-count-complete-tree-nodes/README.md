# Count Nodes in a Complete Binary Tree

**Problem (LeetCode 222):** Count nodes in a **complete** tree faster than O(n).

## Approach 1: Any traversal
**O(n)**.

## Approach 2: Use the "perfect subtree" shortcut (optimal)
At a node, measure the height going **only left** (`lh`) and **only right** (`rh`).
- If `lh == rh`, the subtree is **perfect**: it has `2^lh − 1` nodes. Done, no recursion.
- Otherwise, `1 + count(left) + count(right)`.

In a complete tree, at every level at least one of the two children is perfect, so only one branch keeps recursing.
- **Time:** O(log² n) (log n levels, O(log n) height measurement at each), **Space:** O(log n)
