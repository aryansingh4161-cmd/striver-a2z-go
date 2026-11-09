# Median of a Row-wise Sorted Matrix

**Problem:** n × m matrix, each row sorted, `n × m` is **odd**. Return the median of all elements.

```
1 4 9
2 5 6     → all values sorted: 1 2 3 4 5 6 7 8 9 → median 5
3 7 8
```

## Approach 1: Flatten and sort
**O(nm log(nm)), Space O(nm)**.

## Approach 2: Binary search on the value (optimal)
The median is the smallest value `x` such that **more than half** of all elements are `≤ x`.
- Search `x` between the smallest first-column value and the largest last-column value.
- `countLE(x)` = Σ over rows of **upper bound** of `x` in that row (O(log m) per row).
- Need `countLE(x) > (n·m) / 2`.

- **Time:** O(log(range) × n × log m), **Space:** O(1)

## Dry run (example above, need count > 4)
| x | count ≤ x | > 4? |
|---|---|---|
| 5 | row1: 2, row2: 2, row3: 1 = 5 | ✅ try smaller |
| 4 | 2 + 1 + 1 = 4 | ❌ → answer **5** |
