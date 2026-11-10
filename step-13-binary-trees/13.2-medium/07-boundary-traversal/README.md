# Boundary Traversal

**Problem:** Return the boundary anticlockwise starting from the root:
1. the root,
2. the **left boundary** (top to bottom, excluding leaves),
3. **all leaves** (left to right),
4. the **right boundary** (bottom to top, excluding leaves).

```
        1
       / \
      2   7
     /     \
    3       8
     \     /
      4   9
     / \  / \
    5  6 10 11
```
→ `1 2 3 4 5 6 10 11 9 8 7`

## Approach
1. Add the root (if it isn't a leaf).
2. **Left boundary:** start at `root.Left`; at each non-leaf node add it, then go left if possible, else right.
3. **Leaves:** any traversal (preorder) that adds nodes with no children.
4. **Right boundary:** start at `root.Right`; go right if possible, else left; collect non-leaves, then add them **reversed**.

- **Time:** O(n), **Space:** O(h) (plus the answer)
