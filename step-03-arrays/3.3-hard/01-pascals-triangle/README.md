# Pascal's Triangle

```
row 1:        1
row 2:       1 1
row 3:      1 2 1
row 4:     1 3 3 1
row 5:    1 4 6 4 1
```
Each number = sum of the two numbers above it. Edges are 1.

There are **three** variants (rows and columns counted from 1):

## Variant 1: Value at (row r, column c)
It equals the combination **C(r−1, c−1)**.
Computing factorials overflows fast, so build it step by step:
```
C(n, k) = (n/1) × ((n−1)/2) × … × ((n−k+1)/k)
res = 1
for i = 0 .. k-1:
    res = res * (n - i) / (i + 1)    // always divides exactly
```
- **Time:** O(c), **Space:** O(1)

## Variant 2: Print the whole r-th row
Brute: compute C(r−1, c−1) for every column → O(r²).
**Optimal:** each element comes from the previous one in the same row:
`next = prev × (r − c) / c`
Row 5: 1 → 1×4/1 = 4 → 4×3/2 = 6 → 6×2/3 = 4 → 4×1/4 = 1.
- **Time:** O(r)

## Variant 3: The first N rows (LeetCode 118)
Generate each row with Variant 2. **Time O(N²)** (that's the size of the output, can't do better).
