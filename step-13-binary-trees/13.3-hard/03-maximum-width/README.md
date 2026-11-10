# Maximum Width of a Binary Tree

**Problem (LeetCode 662):** Width of a level = distance between its leftmost and rightmost non-null nodes, **counting the null gaps between them** (as if the tree were complete). Return the max width.

`[1 3 2 5 3 null 9]` → 4 (last level: 5 3 _ 9)

## Approach: BFS with heap-style indexes
Number positions like a heap: root = 0, left child = `2i + 1`, right child = `2i + 2`. Width of a level = `last index − first index + 1`.

**Overflow fix:** indexes double every level, so in a deep, skinny tree they overflow even 64 bits. At each level, subtract the level's first index from every index before computing children (only differences matter).
- **Time:** O(n), **Space:** O(width)
