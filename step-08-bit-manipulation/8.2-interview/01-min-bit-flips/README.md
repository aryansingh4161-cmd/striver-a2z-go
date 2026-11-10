# Minimum Bit Flips to Convert A to B

**Problem (LeetCode 2220):** How many bits must you flip in `a` to get `b`?

a = 10 (`1010`), b = 7 (`0111`) → 3

## Approach (optimal)
`a ^ b` has a 1 **exactly where `a` and `b` differ**. Count its set bits.
`1010 ^ 0111 = 1101` → 3 ones.
- **Time:** O(number of set bits), **Space:** O(1)
