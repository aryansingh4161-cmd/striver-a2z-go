# Largest BST in a Binary Tree

**Problem:** Return the size (node count) of the largest subtree that is a valid BST.

```
      10
     /  \
    5    15
   / \     \
  1   8     7
```
→ 3 (the subtree `5, 1, 8`)

## Approach 1: Validate every subtree
For each node, check "is this a BST?" (O(n)) and count nodes. **O(n²)**.

## Approach 2: Postorder returning (min, max, size) (optimal)
Each subtree reports up: its **minimum**, **maximum**, and **size of the largest BST inside it**.
A node forms a BST with its subtrees if: `left.max < node.val < right.min` (and both children are BSTs).
- If yes: return `(min(left.min, val), max(right.max, val), left.size + right.size + 1)`.
- If no: return `(−∞, +∞, max(left.size, right.size))`. The `(−∞, +∞)` range makes every ancestor fail the check automatically.
- An empty subtree returns `(+∞, −∞, 0)` so any value passes.

- **Time:** O(n), **Space:** O(h)
