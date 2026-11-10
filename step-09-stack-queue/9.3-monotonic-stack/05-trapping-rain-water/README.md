# Trapping Rain Water

**Problem (LeetCode 42):** Bars of width 1 with heights `h`. How much water is trapped after rain?

`[0 1 0 2 1 0 1 3 2 1 2 1]` → 6

## Key formula
Water above bar `i` = `min(maxLeft[i], maxRight[i]) − h[i]`, where `maxLeft` / `maxRight` are the tallest bars on each side (including i). The lower wall decides the water level.

## Approach 1: For each bar, scan both sides
**O(n²)**.

## Approach 2: Prefix max + suffix max arrays
Precompute `maxLeft` (left to right) and `maxRight` (right to left), then apply the formula.
**Time O(n), Space O(n)**.

## Approach 3: Two pointers (optimal)
`l = 0`, `r = n−1`, with `leftMax` and `rightMax` seen so far.
- If `h[l] <= h[r]`: the right side has a wall at least as tall as `h[l]`, so the water at `l` depends only on `leftMax`. Add `leftMax − h[l]` (after updating `leftMax`), `l++`.
- Otherwise do the same on the right side.
- **Time:** O(n), **Space:** O(1)

## Approach 4: Monotonic stack (fills water layer by layer)
Keep indexes of bars with decreasing heights. When a taller bar arrives, pop the bottom `mid`; the water above it is bounded by the new top (left wall) and the current bar (right wall):
`width = i − left − 1`, `height = min(h[left], h[i]) − h[mid]`.
- **Time:** O(n), **Space:** O(n)
