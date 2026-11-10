# Check if Two Trees are Identical

**Problem (LeetCode 100):** Same shape **and** same values.

## Approach: Compare recursively
```
same(a, b):
    if a == nil or b == nil: return a == b     // both nil → true
    return a.val == b.val and same(a.left, b.left) and same(a.right, b.right)
```
- **Time:** O(n), **Space:** O(h)

(Comparing only one traversal isn't enough: different shapes can have the same preorder.)
