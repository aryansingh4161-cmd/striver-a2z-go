# Sum of Beauty of All Substrings

**Problem (LeetCode 1781):** The **beauty** of a string = (highest character frequency) − (lowest frequency among characters that appear). Return the sum of beauty over all substrings.

`"aabcb"` → 5

## Approach: All substrings with a running count (optimal for this problem)
Fix the start `i`; extend `j` one character at a time, updating a `[26]` frequency array. For each `(i, j)`, compute max and min over the 26 counts (ignoring zeros).
- **Time:** O(n² × 26), **Space:** O(1)

Recomputing the counts from scratch for every substring would be O(n³); reusing the running count as `j` grows is what makes it O(n²).
