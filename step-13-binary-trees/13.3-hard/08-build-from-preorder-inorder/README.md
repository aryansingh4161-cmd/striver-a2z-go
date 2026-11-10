# Construct a Binary Tree from Preorder and Inorder

**Problem (LeetCode 105):** Unique values. Rebuild the tree.

preorder `[3 9 20 15 7]`, inorder `[9 3 15 20 7]` →
```
    3
   / \
  9  20
     / \
    15  7
```

## Key facts
- Preorder's **first** element is the root.
- In inorder, everything **left of the root** is the left subtree, everything right is the right subtree.
- The left subtree's size tells you how many preorder elements belong to it.

## Approach: Recursion with a value → inorder-index map
```
build(preStart, preEnd, inStart, inEnd):
    if preStart > preEnd: return nil
    root = pre[preStart]
    k = indexIn[root]             // O(1) with the map
    leftSize = k - inStart
    root.left  = build(preStart+1, preStart+leftSize, inStart, k-1)
    root.right = build(preStart+leftSize+1, preEnd, k+1, inEnd)
```
- **Time:** O(n), **Space:** O(n) for the map + O(h) stack

(Preorder + postorder **can't** uniquely define a general binary tree: you need inorder to know where left ends and right begins.)
