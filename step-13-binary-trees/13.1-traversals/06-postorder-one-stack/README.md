# Iterative Postorder with One Stack

Harder than the other iterative versions, because a node must be visited **after both** subtrees.

Idea: walk left like inorder. When you can't go further left, look at the top node:
- if it has a **right child that you haven't visited yet**, go there,
- otherwise **visit** the top and pop it, remembering it as `lastVisited`.

```
cur = root, last = nil
while cur != nil or stack not empty:
    if cur != nil: push cur; cur = cur.Left
    else:
        top = peek
        if top.Right != nil and top.Right != last: cur = top.Right
        else: visit top; last = pop
```
`last` tells you whether you're coming **back up** from the right subtree (so it's time to visit) or arriving for the first time.
- **Time:** O(n), **Space:** O(h)
