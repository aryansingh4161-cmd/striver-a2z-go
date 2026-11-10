# Step 14: Binary Search Trees (BST)

> Status: ✅ theory and all 15 problems done

Nodes are the same type as Step 13. [`bst/`](bst/) has `bst.Build(values)` (insert one by one) and `bst.Inorder(root)` for tests.

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Search in a BST](14.1-concepts/01-search-bst/) | go left / right by value |
| 2 | [Minimum / maximum](14.1-concepts/02-min-max-bst/) | leftmost / rightmost |
| 3 | [Ceil](14.1-concepts/03-ceil-bst/) | record and go left |
| 4 | [Floor](14.1-concepts/04-floor-bst/) | record and go right |

### 14.2 Practice
| # | Problem | Key idea |
|---|---|---|
| 1 | [Insert a node](14.2-practice/01-insert-bst/) | walk down, attach a leaf |
| 2 | [Delete a node](14.2-practice/02-delete-bst/) | replace with inorder successor |
| 3 | [Kth smallest / largest](14.2-practice/03-kth-smallest-largest/) | inorder with a counter |
| 4 | [Validate a BST](14.2-practice/04-validate-bst/) | pass (low, high) down |
| 5 | [LCA in a BST](14.2-practice/05-lca-bst/) | split point by value |
| 6 | [BST from preorder](14.2-practice/06-bst-from-preorder/) | upper-bound recursion, O(n) |
| 7 | [Inorder successor / predecessor](14.2-practice/07-successor-predecessor/) | strict ceil / floor walk |
| 8 | [BST iterator](14.2-practice/08-bst-iterator/) | stack of the left spine |
| 9 | [Two sum in a BST](14.2-practice/09-two-sum-bst/) | forward + backward iterators |
| 10 | [Recover BST](14.2-practice/10-recover-bst/) | find drops in inorder |
| 11 | [Largest BST in a binary tree](14.2-practice/11-largest-bst-in-bt/) | postorder (min, max, size) |
