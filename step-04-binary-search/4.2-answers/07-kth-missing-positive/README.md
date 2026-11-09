# Kth Missing Positive Number

**Problem (LeetCode 1539):** Strictly increasing positive integers. Return the k-th positive integer missing from the array.

`[2 3 4 7 11]`, k = 5 → 9 (missing: 1 5 6 8 **9** 10 …)

## Approach 1: Walk and shift
For each `x` in order: if `x <= k`, then `k++` (one of the first k numbers is present, so the answer shifts up). Else stop. Answer = `k`.
**O(n)**.

## Approach 2: Binary search (optimal)
Missing numbers before index `i`: `missing(i) = a[i] − (i + 1)` (if nothing were missing, `a[i]` would be `i+1`).
`[2 3 4 7 11]` → missing = `[1 1 1 3 6]` (increasing ✅).

Binary search for the last index with `missing < k`. Call it `hi` after the loop (it ends with `lo = hi + 1`).
Answer = `a[hi] + (k − missing(hi))` = … which simplifies to **`lo + k`**.

- **Time:** O(log n), **Space:** O(1)

## Dry run `[2 3 4 7 11]`, k = 5
| lo | hi | mid | missing | < 5? |
|---|---|---|---|---|
| 0 | 4 | 2 | 1 | yes → lo = 3 |
| 3 | 4 | 3 | 3 | yes → lo = 4 |
| 4 | 4 | 4 | 6 | no → hi = 3 |
lo = 4 → answer = 4 + 5 = **9** ✓
