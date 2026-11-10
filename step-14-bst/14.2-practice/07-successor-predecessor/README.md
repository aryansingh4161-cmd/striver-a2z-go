# Inorder Successor and Predecessor in a BST

**Problem:** For value `x`, the **successor** is the smallest value `> x`; the **predecessor** is the largest value `< x`. (−1 if none.)

BST of `1..9`, x = 5 → successor 6, predecessor 4

## Approach 1: Inorder list
Find `x` in the sorted inorder list; look at neighbours. **O(n)**.

## Approach 2: Walk down once (optimal)
**Successor:** at each node, if `node.val > x`, it's a candidate: record it, go **left** (look for a smaller one). Otherwise go **right**.
**Predecessor:** mirror: if `node.val < x`, record and go **right**; else go **left**.

Like [ceil](../../14.1-concepts/03-ceil-bst/) and [floor](../../14.1-concepts/04-floor-bst/), but **strictly** greater / smaller.
- **Time:** O(h), **Space:** O(1)
