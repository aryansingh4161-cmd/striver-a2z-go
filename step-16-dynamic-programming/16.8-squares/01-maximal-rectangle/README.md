# Maximal Rectangle of 1s

**Problem:** A binary matrix. Find the area of the largest rectangle containing **only 1s**.

```
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0   → 6
```

## Idea: one histogram per row
Treat each row as the **floor** of a histogram. `height[j]` = number of consecutive 1s going **up** from this row in column `j`:
```
height[j] = height[j] + 1   if mat[row][j] == 1
height[j] = 0               otherwise (the bar breaks)
```
Any all-ones rectangle has a bottom row. With that row as the floor, the rectangle is a rectangle inside that row's histogram. So:
```
answer = max over rows of largestRectangleInHistogram(height)
```
"Largest rectangle in histogram" is the stack problem from [Step 9.3](../../../step-09-stack-queue/9.3-monotonic-stack/10-largest-rectangle-histogram/) (this exact problem also appears there as [9.3 #11](../../../step-09-stack-queue/9.3-monotonic-stack/11-maximal-rectangle/); here we see it as "DP on rows"): for every bar, find the nearest smaller bar on each side; the bar's height times that width is a candidate.

## Histograms for the example
| row | heights | largest rectangle |
|---|---|---|
| 0 | 1 0 1 0 0 | 1 |
| 1 | 2 0 2 1 1 | 3 |
| 2 | 3 1 3 2 2 | **6** (height 2 × width 3, columns 2–4) |
| 3 | 4 0 0 3 0 | 4 |
→ **6** ✅

## Complexity
- n rows × O(m) stack pass → **O(n · m)** time, **O(m)** space.
