# Find the Missing Number

**Problem (LeetCode 268):** The array has `n` distinct numbers from the range `0..n`. One is missing; find it.

`[3 0 1]` → 2 (n = 3, range 0..3)

## Approach 1: Check each number (brute)
For each value `0..n`, linear search for it. **Time O(n²)**.

## Approach 2: Hash array (better)
Mark seen values in a boolean array of size n+1, then find the unmarked one. **Time O(n), Space O(n)**.

## Approach 3a: Sum formula (optimal)
Expected sum of `0..n` is `n(n+1)/2`. Missing = expected − actual sum.
- **Time:** O(n), **Space:** O(1)
- Risk: the sum can overflow for huge n in languages with 32-bit ints. Go's 64-bit `int` is fine here.

## Approach 3b: XOR (optimal, no overflow)
XOR facts: `x ^ x = 0`, `x ^ 0 = x`, order doesn't matter.
XOR all numbers `0..n` **and** all array elements. Every present number appears twice and cancels; only the missing one remains.
- **Time:** O(n), **Space:** O(1)

## Dry run XOR `[3 0 1]`
`(0^1^2^3) ^ (3^0^1)` → pairs cancel → **2**
