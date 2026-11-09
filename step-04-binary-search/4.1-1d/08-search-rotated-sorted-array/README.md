# Search in Rotated Sorted Array I

**Problem (LeetCode 33):** A sorted array of **distinct** values was rotated at some unknown point, e.g. `[0 1 2 4 5 6 7]` → `[4 5 6 7 0 1 2]`. Find `target`'s index or −1, in O(log n).

`[4 5 6 7 0 1 2]`, target 0 → 4

## Key insight
At any `mid`, **at least one half is sorted**.
- If `a[lo] <= a[mid]` → the **left** half `[lo..mid]` is sorted.
- Otherwise → the **right** half `[mid..hi]` is sorted.

In a sorted half you can easily check "is the target inside this range?". If yes, search there; if not, search the other half.

```
if a[mid] == target: return mid
if a[lo] <= a[mid]:                     // left sorted
    if a[lo] <= target < a[mid]: hi = mid - 1
    else: lo = mid + 1
else:                                   // right sorted
    if a[mid] < target <= a[hi]: lo = mid + 1
    else: hi = mid - 1
```
- **Time:** O(log n), **Space:** O(1)

## Dry run `[4 5 6 7 0 1 2]`, target 0
| lo | hi | mid | a[mid] | sorted half | target inside? | action |
|---|---|---|---|---|---|---|
| 0 | 6 | 3 | 7 | left [4..7] | no | lo = 4 |
| 4 | 6 | 5 | 1 | left [0..1] | yes | hi = 4 |
| 4 | 4 | 4 | 0 | found → **4** | | |
