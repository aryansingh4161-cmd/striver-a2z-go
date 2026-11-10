# Sum of Subarray Ranges

**Problem (LeetCode 2104):** Range of a subarray = max − min. Sum the ranges of all subarrays.

`[1 2 3]` → 4 (ranges: 0,0,0,1,1,2)

## Approach 1: All subarrays with running max and min
**O(n²)**. Accepted on LeetCode.

## Approach 2: Contributions (optimal)
`Σ(max − min) = Σ max − Σ min`.
- `Σ min` over all subarrays: exactly [sum of subarray minimums](../06-sum-subarray-minimums/) (without the modulo).
- `Σ max`: the same thing with greater instead of smaller.

- **Time:** O(n), **Space:** O(n)
