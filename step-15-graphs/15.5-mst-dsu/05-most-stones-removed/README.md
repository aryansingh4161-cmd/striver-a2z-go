# Most Stones Removed with Same Row or Column

**Problem:** Stones sit on grid points `[row, col]`. You can remove a stone if **another remaining stone** shares its row or column. What's the maximum number you can remove?

```
stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]  → 5
```

## Key insight
Link two stones if they share a row or a column. That makes groups (components).
- In a group of size `k` you can always remove `k − 1` stones (remove leaves of a spanning tree one by one; the last one stays).
- You can never remove the last stone of a group (nothing shares its row/column anymore).

So: **answer = total stones − number of groups.**

## Building groups fast: rows and columns are the nodes
Comparing every pair of stones is O(n²). Better: treat **each row** and **each column** as a DSU node. A stone at `(r, c)` just does `Union(row r, col c)`.
- Rows get ids `0 .. maxRow`, columns get ids `maxRow + 1 + c`.
- Two stones in the same row both touch that row node, so they end up together. Same for columns.
- Count groups only among row/column nodes that actually have a stone.

## Dry run
Stones (0,0) (0,1) (1,0) (1,2) (2,1) (2,2); maxRow = 2 → col c is node 3 + c
- (0,0): R0–C0 · (0,1): R0–C1 · (1,0): R1–C0 · (1,2): R1–C2 · (2,1): R2–C1 · (2,2): R2–C2
- All six used nodes end up in one group → groups = 1
- **6 − 1 = 5** ✅

- **Time:** O(n · α), **Space:** O(maxRow + maxCol)
