# Lower Bound

**Problem:** In a sorted array, return the **first index `i` with `a[i] >= x`**. If none, return `n` (length of the array).

`[3 5 8 15 19]`, x = 9 → 3 (15 is the first ≥ 9)
x = 8 → 2, x = 20 → 5

## Approach 1: Linear scan
First `i` with `a[i] >= x`. **O(n)**.

## Approach 2: Binary search (optimal)
`ans = n`. When `a[mid] >= x`, `mid` **might** be the answer: save it and keep looking **left** for an earlier one. Otherwise go right.
```
ans = n
while lo <= hi:
    mid = (lo+hi)/2
    if a[mid] >= x: ans = mid; hi = mid - 1
    else: lo = mid + 1
```
- **Time:** O(log n), **Space:** O(1)

## Dry run `[3 5 8 15 19]`, x = 9
| lo | hi | mid | a[mid] | ≥ 9? | ans |
|---|---|---|---|---|---|
| 0 | 4 | 2 | 8 | no → lo = 3 | 5 |
| 3 | 4 | 3 | 15 | yes → hi = 2 | **3** |

## In Go's standard library
`sort.SearchInts(a, x)` is exactly lower bound.
This "save and keep searching" pattern is used in almost every problem in this step.
