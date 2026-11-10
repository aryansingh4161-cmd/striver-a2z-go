# Construct a Binary Tree from Postorder and Inorder

**Problem (LeetCode 106):** Same idea as [preorder + inorder](../08-build-from-preorder-inorder/), except the root is the **last** element of postorder.

postorder `[9 15 7 20 3]`, inorder `[9 3 15 20 7]` → same tree as before.

```
build(postStart, postEnd, inStart, inEnd):
    root = post[postEnd]
    k = indexIn[root]; leftSize = k - inStart
    root.left  = build(postStart, postStart+leftSize-1, inStart, k-1)
    root.right = build(postStart+leftSize, postEnd-1, k+1, inEnd)
```
- **Time:** O(n), **Space:** O(n)
