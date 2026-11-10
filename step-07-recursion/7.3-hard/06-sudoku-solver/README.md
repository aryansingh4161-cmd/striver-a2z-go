# Sudoku Solver

**Problem (LeetCode 37):** Fill the empty cells (`.`) so every row, column and 3×3 box contains digits 1–9 exactly once. The puzzle has one solution.

## Approach: Backtracking cell by cell
```
solve():
    find the next empty cell (r, c)
    if none: return true            // solved
    for d = '1' .. '9':
        if d can go at (r, c):
            board[r][c] = d
            if solve(): return true
            board[r][c] = '.'       // backtrack
    return false                    // no digit fits → an earlier choice was wrong
```

### Checking a digit
Brute force: scan the row, column and box (27 cells). Faster: keep three tables of booleans:
`rowHas[r][d]`, `colHas[c][d]`, `boxHas[b][d]` with box index `b = (r/3)*3 + c/3`. Then the check is O(1).

- **Time:** O(9^(empty cells)) in the worst case; with pruning real puzzles solve in milliseconds
- **Space:** O(81) recursion depth at most

## Box index picture
```
b = (r/3)*3 + c/3
0 | 1 | 2
3 | 4 | 5
6 | 7 | 8
```
