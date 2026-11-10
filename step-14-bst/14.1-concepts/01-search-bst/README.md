# Search in a BST

**Problem (LeetCode 700):** Return the node with value `x`, or nil.

## Approach
At each node, the BST rule tells you which side `x` must be on:
- `x == node` → found
- `x < node` → go left
- `x > node` → go right

```go
for n != nil && n.Val != x {
    if x < n.Val { n = n.Left } else { n = n.Right }
}
```
- **Time:** O(h): O(log n) for a balanced tree, O(n) for a skewed one. **Space:** O(1)

This is binary search on a tree.
