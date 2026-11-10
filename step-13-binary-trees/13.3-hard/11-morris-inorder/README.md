# Morris Inorder Traversal (O(1) extra space)

Normal inorder needs O(h) space for the stack (or recursion). **Morris traversal** needs O(1) by **temporarily threading** the tree.

## Idea
Before going left from `cur`, find its **inorder predecessor** (the rightmost node of its left subtree). Make that node's `Right` point back to `cur`: a thread that tells you how to come back up.

```
cur = root
while cur != nil:
    if cur.Left == nil:
        visit cur; cur = cur.Right
    else:
        pred = rightmost node in cur.Left (stop if pred.Right == cur)
        if pred.Right == nil:          // first time here: make the thread
            pred.Right = cur; cur = cur.Left
        else:                          // came back via the thread: remove it
            pred.Right = nil; visit cur; cur = cur.Right
```
- **Time:** O(n) (each edge is walked at most a few times), **Space:** O(1)
- The tree is **restored** by the end (every thread is removed).
