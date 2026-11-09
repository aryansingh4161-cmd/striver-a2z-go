# Single Element in a Sorted Array

**Problem (LeetCode 540):** Sorted array where every element appears **twice** except one. Find it in O(log n).

`[1 1 2 3 3 4 4 8 8]` → 2

## Approach 1: XOR or linear scan
**O(n)**. Works, but doesn't use the sorted order.

## Approach 2: Binary search on index parity (optimal)
Look at the pairs **before** the single element: each pair starts at an **even** index → `(even, odd)`.
**After** the single element, the pattern shifts: pairs start at odd indexes → `(odd, even)`.

So at `mid`:
- If `mid` is even and `a[mid] == a[mid+1]`, **or** `mid` is odd and `a[mid] == a[mid−1]` → we're still in the "before" zone → the single element is on the **right**: `lo = mid + 1`.
- Otherwise → it's at `mid` or on the **left**: `hi = mid` … (or check `mid` directly first).

Handle the edges first (`n == 1`, first element, last element) so `mid−1` and `mid+1` are always valid.
- **Time:** O(log n), **Space:** O(1)

## Dry run `[1 1 2 3 3 4 4 8 8]`
| lo | hi | mid | a[mid] | check | action |
|---|---|---|---|---|---|
| 1 | 7 | 4 | 3 | mid is even, but a[5] = 4 ≠ 3 → pairing is broken → single is on the left | hi = 3 |
| 1 | 3 | 2 | 2 | a[1] = 1 and a[3] = 3 both differ | **found 2** |
