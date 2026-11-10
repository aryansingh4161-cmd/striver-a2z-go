# Bottom View of a Binary Tree

**Problem:** For each column, the **bottom-most** node (if two share the bottom level, the one that comes **later** in level order, i.e. the right one).

```
        20
       /  \
      8    22
     / \     \
    5   3     25
       / \
      10  14
```
Bottom view: `5 10 3 14 25`

## Approach: BFS with columns, keep overwriting
Same as the [top view](../09-top-view/), but **overwrite** the column's value every time: the last node BFS sees in a column is the lowest one.
- **Time:** O(n), **Space:** O(n)
