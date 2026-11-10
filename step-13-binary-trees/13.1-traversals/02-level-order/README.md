# Level Order Traversal (BFS)

**Problem (LeetCode 102):** Return values level by level.

`[3 9 20 null null 15 7]` → `[[3] [9 20] [15 7]]`

## Approach: Queue, one level at a time
```
queue = [root]
while queue not empty:
    size = len(queue)          // nodes on this level
    level = []
    repeat size times:
        node = pop front
        level.append(node.val)
        push node.Left, node.Right (if not nil)
    result.append(level)
```
Reading `size` **before** the inner loop is what separates levels: children pushed during the loop belong to the next level.
- **Time:** O(n), **Space:** O(width) ≤ O(n)

Many problems are this template plus one change: zigzag (reverse alternate levels), right view (last of each level), max width, etc.
