# Maximum Consecutive Ones

**Problem (LeetCode 485):** Binary array. Return the longest run of 1s.

`[1 1 0 1 1 1]` → 3

## Approach (optimal)
Keep a running `count`. On 1: `count++` and update `best`. On 0: reset `count = 0`.
- **Time:** O(n), **Space:** O(1)
