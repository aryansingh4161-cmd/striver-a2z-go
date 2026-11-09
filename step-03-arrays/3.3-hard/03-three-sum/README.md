# 3Sum

**Problem (LeetCode 15):** Return all **unique** triplets `[a, b, c]` with `a + b + c = 0` (different indices).

`[-1 0 1 2 -1 -4]` → `[[-1 -1 2] [-1 0 1]]`

## Approach 1: Three loops (brute)
Try every triplet, sort each one and put it in a set to remove duplicates. **Time O(n³ log)**.

## Approach 2: Hash set for the third (better)
Fix `i` and `j`; the third must be `−(a[i] + a[j])`. Check a set of the elements between `i` and `j`. **Time O(n² log m), Space O(n)** + the set of answers.

## Approach 3: Sort + two pointers (optimal)
1. Sort the array.
2. Fix `i`. **Skip** it if `a[i] == a[i−1]` (same first number → same triplets).
3. Two pointers `j = i+1`, `k = n−1`:
   - `sum < 0` → `j++`
   - `sum > 0` → `k--`
   - `sum == 0` → save, then move **both** and skip repeats: `while a[j] == a[j−1]: j++`.

- **Time:** O(n log n + n²) = O(n²), **Space:** O(1) besides the answer

Sorting lets us skip duplicates without a set and makes the two-pointer moves meaningful.
