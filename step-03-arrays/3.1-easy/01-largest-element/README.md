# Largest Element in an Array

**Problem:** Return the largest element.

`[3 2 1 5 2]` → 5

## Approach 1: Sort (brute)
Sort, return the last element. **Time O(n log n)**. Overkill.

## Approach 2: Current champion (optimal)
```
largest = a[0]
for each x in a:
    if x > largest: largest = x
```
- **Time:** O(n), **Space:** O(1)

Start with `a[0]`, **not 0 or -1**, so all-negative arrays still work: `[-5 -2 -9]` → −2.
