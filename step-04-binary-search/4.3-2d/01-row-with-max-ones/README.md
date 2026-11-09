# Row with Maximum 1s

**Problem:** Binary matrix, every row sorted (0s then 1s). Return the index of the row with the most 1s (the first such row on ties), or −1 if there are no 1s.

```
0 0 1 1 1
0 0 0 0 0
0 1 1 1 1   → row 2 (four 1s)
0 0 1 1 1
```

## Approach 1: Count every row
**O(n × m)**.

## Approach 2: Lower bound per row (optimal-ish)
In a sorted row, ones = `m − (index of first 1)`. The first 1 is the lower bound of 1.
- **Time:** O(n log m), **Space:** O(1)

## Bonus: Staircase walk, O(n + m)
Start at the top-right. Move **left** while you see 1 (this row has more ones), otherwise move **down**. The last row where you moved left is the answer.
