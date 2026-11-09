# Merge Two Sorted Arrays Without Extra Space

**Problem:** `a` (size n) and `b` (size m) are sorted. Rearrange so that the smallest n numbers are in `a` and the rest in `b`, both sorted, using O(1) extra space.

`a = [1 4 8 10]`, `b = [2 3 9]` → `a = [1 2 3 4]`, `b = [8 9 10]`

(LeetCode 88 is a different version: `a` has m empty slots at the end. See `MergeIntoFirst`.)

## Brute
Merge into a third array with two pointers, copy back. **Time O(n + m), Space O(n + m)**.

## Optimal 1: Swap ends, then sort
The biggest elements of `a` belong in `b`. Compare `a`'s **last** with `b`'s **first**:
```
i = n-1, j = 0
while i >= 0 and j < m and a[i] > b[j]:
    swap a[i], b[j]; i--; j++
sort a; sort b
```
- **Time:** O(min(n,m) + n log n + m log m), **Space:** O(1)

## Optimal 2: Gap method (Shell sort idea)
Treat `a` and `b` as one array of length `n + m`. Start with `gap = ceil((n+m)/2)`. Compare elements `gap` apart and swap if out of order. Then `gap = ceil(gap/2)`; stop after the pass with `gap = 1`.
- **Time:** O((n+m) log(n+m)), **Space:** O(1)

## LeetCode 88 (space at the end of `a`)
Fill from the **back**: compare the largest remaining of each, write it at `a[i+j+1]`. No shifting needed. **Time O(n + m), Space O(1)**.
