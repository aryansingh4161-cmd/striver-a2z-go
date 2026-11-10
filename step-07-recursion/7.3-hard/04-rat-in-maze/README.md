# Rat in a Maze

**Problem:** n×n grid, `1` = open, `0` = blocked. A rat starts at `(0,0)` and must reach `(n−1, n−1)` moving **D**own, **L**eft, **R**ight, **U**p, never revisiting a cell. Return all paths as direction strings, in lexicographic order.

```
1 0 0 0
1 1 0 1       → "DDRDRR", "DRDDRR"
1 1 0 0
0 1 1 1
```

## Approach: DFS with backtracking
From `(r, c)`, try moves in alphabetical order **D, L, R, U** (so the output is already sorted). A move is allowed if the new cell is inside the grid, open, and not visited on the current path.
```
f(r, c, path):
    if (r, c) is the target: save path; return
    mark (r, c)
    for each move in D, L, R, U: if allowed → f(next, path + move)
    unmark (r, c)
```
If the start or the target is blocked → no paths.
- **Time:** O(4^(n²)) worst case (in practice far less), **Space:** O(n²) for visited + depth

**Tip:** store moves as parallel arrays `dr = {1, 0, 0, −1}`, `dc = {0, −1, 1, 0}`, `name = "DLRU"` instead of writing four `if` blocks.
