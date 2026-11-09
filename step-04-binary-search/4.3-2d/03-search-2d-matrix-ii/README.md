# Search in a Row-wise and Column-wise Sorted Matrix

**Problem (LeetCode 240):** Each row is sorted left → right and each column is sorted top → bottom (but rows don't continue each other). Find `target`.

```
1  4  7  11 15
2  5  8  12 19
3  6  9  16 22      target 5 → true
10 13 14 17 24
18 21 23 26 30
```

## Approach 1: Binary search every row
**O(n log m)**.

## Approach 2: Start at the top-right corner (optimal)
At the top-right, moving **left** makes values smaller and moving **down** makes them bigger, so each comparison eliminates a whole row or column:
- value == target → found
- value > target → this whole **column** below is even bigger → move **left**
- value < target → this whole **row** to the left is even smaller → move **down**

- **Time:** O(n + m), **Space:** O(1)

(Bottom-left also works. Top-left and bottom-right don't, because both directions go the same way there.)
