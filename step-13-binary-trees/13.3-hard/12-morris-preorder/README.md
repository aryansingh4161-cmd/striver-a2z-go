# Morris Preorder Traversal

Same threading as [Morris inorder](../11-morris-inorder/); only **when you visit** changes.

Preorder visits a node **before** its left subtree, so visit `cur` at the moment you **create** the thread (first arrival), not when you remove it.
```
if cur.Left == nil: visit; cur = cur.Right
else:
    pred = rightmost of cur.Left
    if pred.Right == nil: visit cur; pred.Right = cur; cur = cur.Left
    else: pred.Right = nil; cur = cur.Right
```
- **Time:** O(n), **Space:** O(1)
