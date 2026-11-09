# Rotate Matrix by 90° (clockwise)

**Problem (LeetCode 48):** Rotate an n×n matrix 90° clockwise, in place.

```
1 2 3      7 4 1
4 5 6  →   8 5 2
7 8 9      9 6 3
```

## Approach 1: New matrix (brute)
Element at `(i, j)` goes to `(j, n−1−i)`. Fill a new matrix. **Time O(n²), Space O(n²)**.

## Approach 2: Transpose + reverse each row (optimal)
Look at the answer: its first **row** is the original first **column**, read bottom-up.
1. **Transpose** (swap `(i,j)` with `(j,i)` for `j > i`): rows become columns.
2. **Reverse each row**.

```
1 2 3    transpose   1 4 7    reverse rows   7 4 1
4 5 6   ---------->  2 5 8   ------------->  8 5 2
7 8 9                3 6 9                   9 6 3
```
- **Time:** O(n²), **Space:** O(1)

Swap only above the diagonal (`j > i`), or you swap everything twice and end up where you started.

**Anticlockwise:** transpose, then reverse each **column** (or reverse rows first, then transpose).
