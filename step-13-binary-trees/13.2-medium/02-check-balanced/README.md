# Check if a Binary Tree is Balanced

**Problem (LeetCode 110):** Balanced = at **every** node, the heights of the left and right subtrees differ by at most 1.

`[3 9 20 null null 15 7]` → true
`[1 2 2 3 3 null null 4 4]` → false

## Approach 1: Check every node with a separate height call
At each node compute both heights (O(n)) and recurse. **O(n²)** worst case.

## Approach 2: Height that returns −1 for "unbalanced" (optimal)
Compute heights bottom-up (postorder). If a subtree is unbalanced, or the two heights differ by more than 1, return **−1** and pass it straight up.
```
h(n):
    if n == nil: return 0
    l = h(left);  if l == -1: return -1
    r = h(right); if r == -1: return -1
    if |l - r| > 1: return -1
    return 1 + max(l, r)
balanced = h(root) != -1
```
- **Time:** O(n), **Space:** O(h)
