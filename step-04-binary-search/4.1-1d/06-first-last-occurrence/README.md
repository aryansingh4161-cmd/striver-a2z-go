# First and Last Occurrence

**Problem (LeetCode 34):** Sorted array with duplicates. Return `[first, last]` index of `target`, or `[-1, -1]`.

`[5 7 7 8 8 10]`, target 8 → `[3, 4]`

## Approach 1: Linear scan
**O(n)**.

## Approach 2: Lower + upper bound
- first = lower bound (check that `a[first] == target`!)
- last = upper bound − 1

## Approach 3: Two direct binary searches (clearest)
- **first:** on a match, save `mid`, go **left** (`hi = mid − 1`).
- **last:** on a match, save `mid`, go **right** (`lo = mid + 1`).

- **Time:** O(log n), **Space:** O(1)

## Dry run (first), target 8
| lo | hi | mid | a[mid] | action | first |
|---|---|---|---|---|---|
| 0 | 5 | 2 | 7 | < 8 → lo = 3 | −1 |
| 3 | 5 | 4 | 8 | match → save, hi = 3 | 4 |
| 3 | 3 | 3 | 8 | match → save, hi = 2 | **3** |
