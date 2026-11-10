# Delete a Node from a BST

**Problem (LeetCode 450):** Delete value `key` and return the root.

## Cases for the node to delete
1. **Leaf** → just remove it.
2. **One child** → replace the node with that child.
3. **Two children** → the tricky one.

## Two children, method A: inorder successor
Copy the **smallest value of the right subtree** (the successor) into this node, then delete that value from the right subtree (it has at most one child).

## Two children, method B (Striver's): attach left under right
Find the **largest node of the left subtree** (rightmost), attach the node's right subtree as its right child, and replace the node with its left child. No values are copied.

Both keep every BST rule. Method A is implemented in `solution.go`.
- **Time:** O(h), **Space:** O(h) recursion
