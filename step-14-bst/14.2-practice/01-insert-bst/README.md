# Insert into a BST

**Problem (LeetCode 701):** Insert `x` (not already present) and return the root.

## Approach
Walk down as in search until you'd fall off the tree; attach the new node there. A new value always becomes a **leaf**.
```
if root == nil: return new node
cur = root
loop:
    if x < cur.val: if cur.left == nil: cur.left = new; break; else cur = cur.left
    else:           if cur.right == nil: cur.right = new; break; else cur = cur.right
```
- **Time:** O(h), **Space:** O(1) iterative
