# Next Permutation

**Problem (LeetCode 31):** Rearrange the numbers into the next bigger arrangement in dictionary order. If it's already the biggest, return the smallest (sorted ascending). In place.

| Input | Output |
|---|---|
| `[1 2 3]` | `[1 3 2]` |
| `[3 2 1]` | `[1 2 3]` |
| `[2 1 5 4 3 0 0]` | `[2 3 0 0 1 4 5]` |

## Brute force
Generate all permutations, sort them, find the current one, return the next. **O(n! × n)**, impossible beyond tiny n.

## Optimal: 3 steps
Think about how a dictionary orders words: keep the longest possible **prefix** the same, and change as late as possible.

1. **Find the break point:** from the right, find the first index `i` with `a[i] < a[i+1]`. Everything after `i` is decreasing (already the largest arrangement of those numbers).
   - No such `i`? The whole array is decreasing → reverse it and stop.
2. **Swap:** from the right, find the first `a[j] > a[i]` (the smallest number bigger than `a[i]` in the suffix). Swap `a[i]` and `a[j]`.
3. **Reverse** the part after `i` (it's decreasing; reversing makes it the smallest).

- **Time:** O(n), **Space:** O(1)

## Dry run `[2 1 5 4 3 0 0]`
1. From the right: 0≥0, 3≥0, 4≥3, 5≥4, **1 < 5** → i = 1 (value 1).
2. From the right, first value > 1: **3** (index 4). Swap → `[2 3 5 4 1 0 0]`.
3. Reverse after index 1 → `[2 3 0 0 1 4 5]`. ✓
