# Left Rotate Array by D Places

**Problem:** Rotate the array left by `d` places. (LeetCode 189 rotates **right**: the same idea in mirror.)

`[1 2 3 4 5 6 7]`, d = 3 → `[4 5 6 7 1 2 3]`

First: `d = d % n`. Rotating by n gives the same array back.

## Approach 1: Temp array (brute)
1. Copy the first `d` elements into `temp`.
2. Shift the rest left by `d`.
3. Put `temp` at the end.
- **Time:** O(n), **Space:** O(d)

## Approach 2: Reversal algorithm (optimal)
1. Reverse the first `d` elements.
2. Reverse the remaining `n − d`.
3. Reverse the whole array.

```
[1 2 3 | 4 5 6 7]
[3 2 1 | 4 5 6 7]   reverse first 3
[3 2 1 | 7 6 5 4]   reverse the rest
[4 5 6 7 1 2 3]     reverse all  ✓
```
- **Time:** O(n) (each element is moved about twice), **Space:** O(1)

**Right rotate by k:** reverse all, reverse first k, reverse the rest.
