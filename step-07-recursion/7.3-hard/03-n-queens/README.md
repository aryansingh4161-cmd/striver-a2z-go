# N-Queens

**Problem (LeetCode 51):** Place `n` queens on an n×n board so that no two attack each other (same row, column or diagonal). Return all boards.

n = 4 →
```
. Q . .      . . Q .
. . . Q      Q . . .
Q . . .      . . . Q
. . Q .      . Q . .
```

## Approach: One queen per column (or row), backtracking
Place queens **row by row**. In row `r`, try every column `c` that is safe; place, recurse to `r+1`, remove.

### Checking "safe" in O(1)
Keep three boolean arrays:
| Array | Index | Why |
|---|---|---|
| `cols[c]` | column | same column |
| `diag1[r+c]` | `r + c` (0 .. 2n−2) | same "/" diagonal has the same r+c |
| `diag2[r−c+n−1]` | `r − c` shifted to be ≥ 0 | same "\" diagonal has the same r−c |

The brute-force check walks up the column and both diagonals: O(n) per check. The arrays make it O(1).

- **Time:** O(n!) (row 1 has n choices, row 2 at most n−1, …), **Space:** O(n)

Number of solutions: n=4 → 2, n=8 → 92.
