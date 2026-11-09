# Set Matrix Zeroes

**Problem (LeetCode 73):** If a cell is 0, set its whole row and column to 0. In place.

```
1 1 1        1 0 1
1 0 1   →    0 0 0
1 1 1        1 0 1
```

**Trap:** if you zero things while scanning, the new zeros trigger more zeroing. You must first **remember** which rows/columns to clear, then clear them.

## Approach 1: Mark with a special value (brute)
For each 0, mark its row/column cells (that aren't 0) with a placeholder, then turn placeholders into 0. Placeholders are hacky in Go (no spare value for ints). **Time O(n·m·(n+m))**. Skip this one.

## Approach 2: Row and column arrays (better)
`row[i] = true` if row `i` has a 0; `col[j] = true` if column `j` has one. Second pass: zero cell `(i,j)` if `row[i] || col[j]`.
- **Time:** O(2·n·m), **Space:** O(n + m)

## Approach 3: Use the first row and column as markers (optimal, O(1) space)
Store `row[i]` in `matrix[i][0]` and `col[j]` in `matrix[0][j]`.
Problem: `matrix[0][0]` would be shared by row 0 and column 0, so keep column 0's flag in a separate variable `col0`.

1. Scan every cell. If `matrix[i][j] == 0`: set `matrix[i][0] = 0`, and `matrix[0][j] = 0` (if `j == 0`, set `col0 = 0` instead).
2. Zero the inner cells `(i ≥ 1, j ≥ 1)` using the markers.
3. If `matrix[0][0] == 0`, zero the first **row**.
4. If `col0 == 0`, zero the first **column**.

Steps 3 and 4 must come last, otherwise you destroy markers you still need.
- **Time:** O(2·n·m), **Space:** O(1)
