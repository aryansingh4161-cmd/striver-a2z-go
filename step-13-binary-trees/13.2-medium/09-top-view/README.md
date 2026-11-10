# Top View of a Binary Tree

**Problem:** The nodes you'd see looking down from above: for each **column**, the **topmost** node (left to right).

```
      1
     / \
    2   3
     \
      4
       \
        5
```
Top view: `2 1 3` (4 and 5 are hidden under 1 and 3).

## Approach: BFS with columns (optimal)
BFS visits nodes **top to bottom**, so the **first** node seen in each column is the top one. Store it only if the column is new. Track `minCol..maxCol` to output in order.
- **Time:** O(n), **Space:** O(n)

DFS also works if you store `(row, val)` and keep the smaller row, but BFS makes "first seen = topmost" automatic.
