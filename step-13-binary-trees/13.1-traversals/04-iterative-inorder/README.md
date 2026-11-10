# Iterative Inorder (with a stack)

Inorder = go **as far left as possible**, visit, then turn right.
```
cur = root
loop:
    while cur != nil: push cur; cur = cur.Left    // dive left
    if stack empty: break
    cur = pop; visit cur                           // leftmost unvisited
    cur = cur.Right                                // then its right subtree
```
- **Time:** O(n), **Space:** O(h)

This exact loop is the basis of the **BST iterator** in Step 14.
