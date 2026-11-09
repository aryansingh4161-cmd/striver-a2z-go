# Floor and Ceil in a Sorted Array

**Problem:** For value `x`:
- **floor** = largest element `≤ x`
- **ceil** = smallest element `≥ x`

Return −1 if it doesn't exist.

`[3 4 4 7 8 10]`, x = 5 → floor 4, ceil 7
x = 8 → floor 8, ceil 8

## Approach: two binary searches (optimal)
- **Ceil** = value at the lower bound.
- **Floor:** when `a[mid] <= x`, save `a[mid]` and go **right** (look for a bigger one that still fits). Otherwise go left.

- **Time:** O(log n) each, **Space:** O(1)
