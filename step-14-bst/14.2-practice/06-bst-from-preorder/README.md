# Construct a BST from Preorder

**Problem (LeetCode 1008):** Build the BST whose preorder is the given array.

`[8 5 1 7 10 12]` →
```
      8
     / \
    5   10
   / \    \
  1   7    12
```

## Approach 1: Insert one by one
**O(n²)** worst case (skewed).

## Approach 2: Sort to get inorder, then use preorder + inorder
**O(n log n)**.

## Approach 3: Upper bound recursion (optimal)
Read preorder left to right with a shared index. A node may be placed only if its value is below the **upper bound** allowed at that position.
```
build(bound):
    if i == n or pre[i] > bound: return nil
    root = pre[i]; i++
    root.left  = build(root.val)    // left values must be < root
    root.right = build(bound)       // right values keep the parent's bound
```
Start with `bound = +∞`. Each value is used exactly once → **O(n)**.
- **Time:** O(n), **Space:** O(h)
