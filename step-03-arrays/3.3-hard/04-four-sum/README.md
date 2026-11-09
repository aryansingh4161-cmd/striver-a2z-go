# 4Sum

**Problem (LeetCode 18):** Return all unique quadruplets that add up to `target`.

`[1 0 -1 0 -2 2]`, target 0 → `[[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]`

## Brute → better
- 4 loops: O(n⁴).
- 3 loops + hash set for the 4th: O(n³).

## Optimal: Sort + two fixed + two pointers
Same as [3Sum](../03-three-sum/) with one more fixed index:
1. Sort.
2. Loop `i`; skip duplicates.
3. Loop `j` from `i+1`; skip duplicates (`j > i+1 && a[j] == a[j−1]`).
4. Two pointers `k = j+1`, `l = n−1` for the remaining two.

- **Time:** O(n³), **Space:** O(1) besides the answer

**Overflow note:** in C++/Java the sum of four ints can overflow 32 bits, so videos use `long long`. Go's `int` is 64-bit, so plain `int` is fine.
