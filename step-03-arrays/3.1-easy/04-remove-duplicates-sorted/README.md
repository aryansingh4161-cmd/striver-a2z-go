# Remove Duplicates from Sorted Array

**Problem (LeetCode 26):** The array is sorted. Remove duplicates **in place** so each unique value appears once at the front. Return `k`, the number of unique values. Anything after index `k−1` doesn't matter.

`[1 1 2 2 2 3 3]` → k = 3, array starts with `[1 2 3 …]`

## Approach 1: Set (brute)
Put everything in a set, sort the set, write it back. **Time O(n log n), Space O(n)**.

## Approach 2: Two pointers (optimal)
- `i` = last position of the unique part.
- `j` scans ahead. When `a[j] != a[i]`, it's a new value: move `i` forward and copy `a[j]` there.

```
i = 0
for j = 1 .. n-1:
    if a[j] != a[i]:
        i++
        a[i] = a[j]
return i + 1
```
- **Time:** O(n), **Space:** O(1)

## Dry run `[1 1 2 2 2 3 3]`
| j | a[j] | new? | i | front |
|---|---|---|---|---|
| 1 | 1 | no | 0 | 1 |
| 2 | 2 | yes | 1 | 1 2 |
| 3,4 | 2 | no | 1 | 1 2 |
| 5 | 3 | yes | 2 | 1 2 3 |
| 6 | 3 | no | 2 | → k = **3** |

Works only because the array is **sorted** (duplicates sit next to each other).
