# Zigzag Level Order Traversal

**Problem (LeetCode 103):** Level order, but alternate direction each level: left→right, then right→left, …

`[3 9 20 null null 15 7]` → `[[3] [20 9] [15 7]]`

## Approach
Normal [level order](../../13.1-traversals/02-level-order/) with a flag. For right→left levels, **fill the level slice from the back** (`level[size−1−i] = val`) instead of reversing afterwards.
- **Time:** O(n), **Space:** O(n)
