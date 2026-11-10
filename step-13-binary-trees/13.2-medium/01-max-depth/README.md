# Maximum Depth (Height) of a Binary Tree

**Problem (LeetCode 104):** Number of nodes on the longest root-to-leaf path.

`[3 9 20 null null 15 7]` → 3

## Approach 1: Recursion (postorder)
`height(n) = 0` if nil, else `1 + max(height(left), height(right))`.
- **Time:** O(n), **Space:** O(h)

## Approach 2: Level order
Count the number of levels in a BFS. **Time O(n), Space O(width)**.

This tiny function is the base of many "hard" problems: balanced tree, diameter, max path sum all add **one line** to it.
