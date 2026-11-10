# Right View and Left View

**Right view (LeetCode 199):** the last node of every level (what you see from the right side).
**Left view:** the first node of every level.

`[1 2 3 null 5 null 4]` → right view `[1 3 4]`, left view `[1 2 5]`

## Approach 1: Level order
Take the last (or first) node of each level. **O(n)** time, O(width) space.

## Approach 2: DFS with the level number (shorter code)
For the **right view**, visit **right before left**. The first time you reach a new depth (`depth == len(result)`), that node is the rightmost of that level.
For the left view, visit left before right.
- **Time:** O(n), **Space:** O(h)
