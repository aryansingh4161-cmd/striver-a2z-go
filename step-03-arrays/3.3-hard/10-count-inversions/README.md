# Count Inversions

**Problem:** An inversion is a pair `(i, j)` with `i < j` and `a[i] > a[j]`. Count them. (It measures how far the array is from sorted.)

`[5 3 2 4 1]` → 8

## Approach 1: All pairs (brute)
**O(n²)**.

## Approach 2: Merge sort (optimal)
During the **merge** step, both halves are sorted. When we take `right[j]` because `left[i] > right[j]`, then **every remaining element of the left half** (`left[i..end]`) is also bigger than `right[j]`. That's `mid − i + 1` inversions in one go.

```
if a[left] <= a[right]: take left
else: count += (mid - left + 1); take right
```
Total = inversions inside left half + inside right half + across (counted in merge).
- **Time:** O(n log n), **Space:** O(n)

Note: this sorts the array as a side effect. The solution works on a copy so the caller's input is unchanged.

## Dry run of one merge
left = `[2 3 5]`, right = `[1 4]`:
- 2 > 1 → take 1, add 3 (2,3,5 all > 1)
- 2 < 4 → take 2; 3 < 4 → take 3
- 5 > 4 → take 4, add 1 (only 5)
Across-count = **4**.
