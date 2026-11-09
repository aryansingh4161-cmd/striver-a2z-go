# Single Number (appears once, others twice)

**Problem (LeetCode 136):** Every element appears twice except one. Find it.

`[4 1 2 1 2]` → 4

## Approach 1: Count each (brute)
For each element, count its occurrences with a nested loop. **Time O(n²)**.

## Approach 2: Hash map (better)
Count with a map, return the key with count 1. **Time O(n), Space O(n)**.

## Approach 3: XOR (optimal)
`x ^ x = 0` and `x ^ 0 = x`. XOR of everything: pairs cancel, the single one stays.
- **Time:** O(n), **Space:** O(1)

`4 ^ 1 ^ 2 ^ 1 ^ 2 = 4 ^ (1^1) ^ (2^2) = 4 ^ 0 ^ 0 = 4`
