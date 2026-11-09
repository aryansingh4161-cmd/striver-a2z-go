# Check if Array is Sorted

**Problem:** Return true if the array is sorted in non-decreasing order.

| Input | Output |
|---|---|
| `[1 2 2 3]` | true |
| `[1 3 2]` | false |

## Approach (optimal)
Every element must be ≥ the one before it.
```
for i = 1 .. n-1:
    if a[i] < a[i-1]: return false
return true
```
- **Time:** O(n), **Space:** O(1)

## Variant: LeetCode 1752 (sorted and rotated)
A sorted array rotated once has **at most one** "drop" (`a[i] > a[i+1]`), counting the wrap-around from last to first. `[3 4 5 1 2]` has one drop (5→1) → true.
