# Preorder, Inorder, Postorder (recursive)

The three depth-first orders differ only in **when you visit the root**:

| Order | Visit | For the tree below |
|---|---|---|
| Preorder | **Root**, Left, Right | 1 2 4 5 3 |
| Inorder | Left, **Root**, Right | 4 2 5 1 3 |
| Postorder | Left, Right, **Root** | 4 5 2 3 1 |

```
      1
     / \
    2   3
   / \
  4   5
```

```go
func preorder(n *Node, out *[]int) {
    if n == nil { return }
    *out = append(*out, n.Val)   // move this line to change the order
    preorder(n.Left, out)
    preorder(n.Right, out)
}
```
- **Time:** O(n), **Space:** O(h) recursion stack (h = height; O(n) for a skewed tree)

**Memory trick:** "pre" = root first, "post" = root last, "in" = root in the middle.
