# Find a Peak Element II (2D)

**Problem (LeetCode 1901):** A peak is strictly greater than its 4 neighbours (outside the grid counts as −1). Adjacent cells are never equal. Return the position of any peak, in O(n log m) or O(m log n).

```
1  4
3  2      → [0, 1] (4) or [1, 0] (3)
```

## Approach 1: Check every cell
**O(n × m × 4)**.

## Approach 2: Binary search on columns (optimal)
Pick the middle column `mid`. Find the **row with the maximum value** in that column: call it `r`. That cell is already bigger than its up/down neighbours (it's the column max). Now compare with left and right:
- bigger than both → it's a peak ✅
- left neighbour bigger → a peak exists on the **left** → `hi = mid − 1`
- right neighbour bigger → go **right** → `lo = mid + 1`

Same "walk uphill" logic as the [1D peak](../../4.1-1d/13-find-peak-element/), applied to column maxima.
- **Time:** O(n × log m), **Space:** O(1)
