# Find Median from Data Stream

**Problem (LeetCode 295):** `AddNum(x)` adds a number; `FindMedian()` returns the median of everything so far.

add 1, add 2 → 1.5; add 3 → 2

## Approach 1: Keep a sorted slice
Insert in the right place (binary search + shift): **AddNum O(n)**, FindMedian O(1).

## Approach 2: Two heaps (optimal)
Split the numbers into two halves:
- `low`: a **max-heap** with the smaller half (its top is the largest of the small numbers)
- `high`: a **min-heap** with the bigger half (its top is the smallest of the big numbers)

Rules: every number in `low` ≤ every number in `high`, and `len(low)` is equal to `len(high)` or one more.
```
AddNum(x):
    push x into low
    move low's top into high        // keeps the order rule
    if len(high) > len(low): move high's top back into low   // keeps the size rule
FindMedian:
    odd total  → low.top
    even total → (low.top + high.top) / 2
```
- **AddNum:** O(log n), **FindMedian:** O(1), **Space:** O(n)
