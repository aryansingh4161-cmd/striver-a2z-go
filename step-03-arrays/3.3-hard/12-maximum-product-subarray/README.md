# Maximum Product Subarray

**Problem (LeetCode 152):** Return the largest product of a contiguous subarray.

`[2 3 -2 4]` → 6 (`[2 3]`)
`[-2 3 -4]` → 24 (whole array)

## Brute
All subarrays with a running product. **O(n²)**.

## Optimal 1: Prefix and suffix products (easiest to remember)
Observations:
- With an **even** number of negatives (and no zeros), the whole array is the answer.
- With an **odd** number, you must drop one negative, so the answer is a prefix or a suffix that excludes it.
- A **zero** splits the array; restart the product after it.

So: walk from the left (prefix product) and from the right (suffix product) at the same time, resetting to 1 after a zero, and take the max of everything seen.
- **Time:** O(n), **Space:** O(1)

## Optimal 2: Track max and min (Kadane-style)
Keep both the max and the min product ending here. A negative number flips them: the smallest (most negative) becomes the largest after multiplying.
```
for x in a[1:]:
    if x < 0: swap(maxHere, minHere)
    maxHere = max(x, maxHere * x)
    minHere = min(x, minHere * x)
    best = max(best, maxHere)
```
- **Time:** O(n), **Space:** O(1)
