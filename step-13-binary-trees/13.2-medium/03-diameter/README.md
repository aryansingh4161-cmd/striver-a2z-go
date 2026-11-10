# Diameter of a Binary Tree

**Problem (LeetCode 543):** The diameter is the number of **edges** on the longest path between any two nodes. The path doesn't have to pass through the root.

`[1 2 3 4 5]` → 3 (path 4 → 2 → 1 → 3, or 5 → 2 → 1 → 3)

## Key insight
The longest path "bends" at some node: it goes down the left side and down the right side. Through node `n` its length is `height(left) + height(right)`.

## Approach 1: For every node, compute both heights
**O(n²)**.

## Approach 2: One height pass, update a global answer (optimal)
```
h(n):
    if n == nil: return 0
    l = h(left); r = h(right)
    best = max(best, l + r)     // the one extra line
    return 1 + max(l, r)
```
- **Time:** O(n), **Space:** O(h)
