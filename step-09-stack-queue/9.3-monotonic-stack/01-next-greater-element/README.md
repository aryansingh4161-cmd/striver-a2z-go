# Next Greater Element

**Problem:** For each element, find the first element **to its right** that is greater. −1 if none.

`[4 5 2 10 8]` → `[5 10 10 −1 −1]`

(LeetCode 496 asks this for a subset of values; same core idea.)

## Approach 1: Brute force
For each `i`, scan right. **O(n²)**.

## Approach 2: Monotonic stack (optimal)
Scan **from the right**, keeping a stack of candidates that are **decreasing from bottom to top**.
For `a[i]`:
1. **Pop** every element `<= a[i]`. They're useless now: `a[i]` is bigger *and* closer for everyone to the left.
2. The top (if any) is the answer for `i`.
3. **Push** `a[i]`.

- **Time:** O(n): each element is pushed once and popped at most once. **Space:** O(n)

## Dry run `[4 5 2 10 8]` (from the right)
| i | a[i] | stack before | pop | answer | stack after |
|---|---|---|---|---|---|
| 4 | 8 | [] | – | −1 | [8] |
| 3 | 10 | [8] | 8 | −1 | [10] |
| 2 | 2 | [10] | – | 10 | [10 2] |
| 1 | 5 | [10 2] | 2 | 10 | [10 5] |
| 0 | 4 | [10 5] | – | 5 | [10 5 4] |
