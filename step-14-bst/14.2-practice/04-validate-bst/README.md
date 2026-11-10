# Validate a BST

**Problem (LeetCode 98):** Is it a valid BST (strictly smaller on the left, strictly greater on the right, for **every** ancestor, not just the parent)?

```
    5
   / \
  1   4      → false (3 is in 5's right subtree but 3 < 5)
     / \
    3   6
```

## Wrong approach (common mistake)
Checking only `left.val < node.val < right.val` at each node misses the example above.

## Approach 1: Inorder must be strictly increasing
**O(n)**; compare each value with the previous one.

## Approach 2: Pass down a valid range (optimal)
```
valid(n, low, high):
    if n == nil: return true
    if not (low < n.val < high): return false
    return valid(n.left, low, n.val) and valid(n.right, n.val, high)
```
Start with `(−∞, +∞)`. Use `math.MinInt` / `math.MaxInt`, or pointers for "no bound" if values can reach those limits.
- **Time:** O(n), **Space:** O(h)
