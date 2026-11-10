# Check if a Number is a Power of 2

**Problem (LeetCode 231):** 1, 2, 4, 8, … → true.

## Approach 1: Keep dividing by 2
While even, divide by 2; check you end at 1. **O(log n)**.

## Approach 2: `n & (n−1)` (optimal)
A power of 2 has **exactly one** set bit: `8 = 1000`.
Subtracting 1 flips that bit and turns on everything below it: `7 = 0111`.
So `n & (n−1) == 0` exactly for powers of 2 (when `n > 0`).
- **Time:** O(1)

`n & (n−1)` = "**clear the lowest set bit**". It's one of the most useful tricks in this step.
