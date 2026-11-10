# Candy

**Problem (LeetCode 135):** Children stand in a line with `ratings`. Each child gets at least 1 candy; a child with a **higher rating than a neighbour** must get more than that neighbour. Minimum total candies?

`[1 0 2]` → 5 (`2 1 2`), `[1 2 2]` → 4 (`1 2 1`)

## Approach 1: Two passes with two arrays
- **Left pass:** `left[i] = left[i−1] + 1` if `r[i] > r[i−1]`, else 1. Satisfies the left-neighbour rule.
- **Right pass:** `right[i] = right[i+1] + 1` if `r[i] > r[i+1]`, else 1.
- Each child needs `max(left[i], right[i])`.
- **Time:** O(n), **Space:** O(n)

## Approach 2: One array (same idea, less memory)
Do the left pass into `candy`, then in the right pass update `candy[i] = max(candy[i], candy[i+1] + 1)` when `r[i] > r[i+1]`.
- **Time:** O(n), **Space:** O(n)

(An O(1)-space "slope counting" version exists: count lengths of increasing and decreasing runs. It's tricky; the two-pass version is what interviewers expect first.)

## Dry run `[1 3 4 5 2]`
left  `1 2 3 4 1`
right `1 1 1 2 1`
max   `1 2 3 4 1` → **11**
