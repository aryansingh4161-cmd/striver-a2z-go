# Number of Islands II (Online Queries)

**Problem:** An `n × m` grid starts as all water. Each operation `[r, c]` turns that cell into land. After **each** operation, report the number of islands (4-directional).

```
n = 4, m = 5, ops = [[1,1],[0,1],[3,3],[3,4]] → [1 1 2 2]
```

## Why not BFS every time?
Running a full island count after each of `k` operations costs O(k · n · m). Too slow when `k` is large.

## Idea: DSU, one cell at a time
Give cell `(r, c)` the id `r·m + c`.
```
for each op (r, c):
    if already land: answer stays the same
    mark land; islands++                       // a brand new island
    for each 4-neighbour that is land:
        if Union(cell, neighbour) succeeded: islands--   // two islands became one
    record islands
```
Each successful union glues two different islands into one, so the count drops by exactly 1.

## Dry run
- (1,1): new island → 1
- (0,1): new → 2; neighbour (1,1) is land, different group → union → **1**
- (3,3): new → **2**
- (3,4): new → 3; neighbour (3,3) → union → **2**
→ `[1 1 2 2]` ✅

## Edge case
The same cell can be added twice: don't count it again.

- **Time:** O(k · α), **Space:** O(n · m)
