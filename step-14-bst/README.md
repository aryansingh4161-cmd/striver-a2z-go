# Step 14: Binary Search Trees (BST)

> Status: theory ready, problems coming ⏳

## The rule
For **every** node: all values in the left subtree are **smaller**, all values in the right subtree are **larger**.

Consequences:
- **Inorder traversal gives sorted order.** Many BST problems are just "do inorder and use the sorted sequence".
- Search, insert and delete take **O(height)**: O(log n) if balanced, O(n) if it degenerates into a line (e.g. inserting 1, 2, 3, 4…).

## Search (iterative, O(h))
```go
for root != nil && root.Val != target {
    if target < root.Val {
        root = root.Left
    } else {
        root = root.Right
    }
}
return root
```

## Key ideas
- **Floor / ceil:** walk down like search, remembering the best candidate seen.
- **Insert:** walk down to a `nil` spot and attach.
- **Delete:** if the node has two children, replace it with its inorder successor (smallest in the right subtree), or attach the left subtree under the right subtree's minimum.
- **Validate BST:** pass a valid range `(low, high)` down; each node must sit strictly inside it. Checking only parent-child is **wrong**.
- **LCA in BST:** if both values are smaller → go left; both bigger → go right; otherwise this node is the LCA.
- **Kth smallest:** inorder, count to K.
- **BST iterator:** a stack holding the "left spine" gives next-smallest in O(1) average.

## Problems

### 14.1 Concepts
| # | Problem | Status |
|---|---|---|
| 1 | Search in a BST | ⏳ |
| 2 | Minimum / maximum in a BST | ⏳ |
| 3 | Ceil in a BST | ⏳ |
| 4 | Floor in a BST | ⏳ |

### 14.2 Practice
| # | Problem | Status |
|---|---|---|
| 1 | Insert a node | ⏳ |
| 2 | Delete a node | ⏳ |
| 3 | Kth smallest / largest | ⏳ |
| 4 | Validate a BST | ⏳ |
| 5 | LCA in a BST | ⏳ |
| 6 | Build BST from preorder | ⏳ |
| 7 | Inorder successor / predecessor | ⏳ |
| 8 | BST iterator | ⏳ |
| 9 | Two sum in a BST | ⏳ |
| 10 | Recover BST (two swapped nodes) | ⏳ |
| 11 | Largest BST in a binary tree | ⏳ |
