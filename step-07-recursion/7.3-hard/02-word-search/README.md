# Word Search

**Problem (LeetCode 79):** Can `word` be traced in the grid by moving up/down/left/right, without using the same cell twice?

```
A B C E
S F C S      "ABCCED" → true,  "ABCB" → false
A D E E
```

## Approach: DFS from every cell with backtracking
1. For each cell matching `word[0]`, start a DFS.
2. `dfs(r, c, k)`: does `word[k:]` start at `(r, c)`?
   - out of the grid, or `grid[r][c] != word[k]` → false
   - `k == len(word) − 1` → true (last letter matched)
   - **mark** the cell as used (e.g. set it to `'#'`)
   - try the 4 neighbours with `k + 1`
   - **unmark** it (backtrack) and return the result

Marking in the grid itself avoids a separate `visited` array.
- **Time:** O(n × m × 3^L) (after the first step each cell has at most 3 new directions), **Space:** O(L) recursion depth
