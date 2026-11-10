# Symmetric Tree

**Problem (LeetCode 101):** Is the tree a mirror of itself around its centre?

`[1 2 2 3 4 4 3]` → true, `[1 2 2 null 3 null 3]` → false

## Approach: Compare two subtrees as mirrors
```
mirror(a, b):
    if a == nil or b == nil: return a == b
    return a.val == b.val
       and mirror(a.left,  b.right)    // outside pair
       and mirror(a.right, b.left)     // inside pair
answer = mirror(root.left, root.right)
```
Compare with [identical trees](../05-identical-trees/): the only change is that the children are crossed.
- **Time:** O(n), **Space:** O(h)
