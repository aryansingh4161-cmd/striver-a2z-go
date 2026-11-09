# Binary Search to Find X

**Problem (LeetCode 704):** Sorted array. Return the index of `target`, or −1.

`[-1 0 3 5 9 12]`, target 9 → 4

## Approach 1: Linear search
**O(n)**. Ignores the fact that the array is sorted.

## Approach 2: Binary search, iterative (optimal)
Look at the middle:
- equal → found,
- middle is smaller → the target can only be on the **right**: `lo = mid + 1`,
- middle is bigger → only on the **left**: `hi = mid − 1`.

Stop when `lo > hi` (empty range).
- **Time:** O(log n), **Space:** O(1)

## Approach 3: Binary search, recursive
Same logic, `f(lo, hi)` calls itself on one half. **Space O(log n)** for the stack.

## Dry run `[-1 0 3 5 9 12]`, target 9
| lo | hi | mid | a[mid] | action |
|---|---|---|---|---|
| 0 | 5 | 2 | 3 | 3 < 9 → lo = 3 |
| 3 | 5 | 4 | 9 | found → **4** |

## Overflow note
`mid := (lo + hi) / 2` can overflow when `lo + hi` exceeds the int range (a real bug in 32-bit languages). `lo + (hi−lo)/2` never overflows. With Go's 64-bit ints it doesn't matter in practice, but the safe form is a good habit.
