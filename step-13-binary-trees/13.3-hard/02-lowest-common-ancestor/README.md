# Lowest Common Ancestor (LCA)

**Problem (LeetCode 236):** The LCA of `p` and `q` is the deepest node that has both as descendants (a node counts as its own descendant).

```
        3
      /   \
     5     1
    / \   / \
   6   2 0   8
      / \
     7   4
```
LCA(5, 1) = 3, LCA(5, 4) = 5, LCA(7, 4) = 2

## Approach 1: Compare root-to-node paths
Find both [paths](../01-root-to-node-path/); the last common node is the LCA. **O(n)** time, O(h) extra space, two passes.

## Approach 2: One recursive pass (optimal)
```
lca(n):
    if n == nil or n == p or n == q: return n
    left = lca(n.left); right = lca(n.right)
    if left != nil and right != nil: return n    // p and q split here
    return left if left != nil else right        // pass up what was found
```
- **Time:** O(n), **Space:** O(h)

Why does returning `n` when `n == p` work even if `q` is below `p`? Because then `p` **is** the LCA, and the other branch returns nil.
