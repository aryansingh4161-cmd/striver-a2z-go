# Asteroid Collision

**Problem (LeetCode 735):** Each asteroid's sign is its direction (+ right, − left), absolute value its size. When a right-mover meets a left-mover, the smaller explodes; equal sizes both explode. Return the final state.

`[5 10 −5]` → `[5 10]`
`[8 −8]` → `[]`
`[10 2 −5]` → `[10]`

## Approach: Stack of survivors (optimal)
Only a **right-moving asteroid on the stack** can collide with a **new left-moving** asteroid.
For each asteroid `x`:
- if `x > 0`: push.
- if `x < 0`: while the top is positive and smaller than `|x|`, pop it (it explodes).
  - top positive and **equal** → pop it, and `x` explodes too.
  - top positive and **bigger** → `x` explodes.
  - stack empty or top negative → push `x` (it flies off safely).

- **Time:** O(n) (each asteroid pushed/popped once), **Space:** O(n)
