# Swim in Rising Water

**Problem:** `n × n` grid of **distinct** heights `0 .. n²−1`. At time `t` the water level is `t`, and you can swim between 4-neighbouring cells if **both** heights are `≤ t`. Starting at `(0,0)`, what's the earliest time you can reach `(n−1, n−1)`?

```
0 2
1 3   → 3 (you can't even stand on (1,1) before t = 3)
```

The answer is the smallest possible value of **the highest cell on a path** from start to end.

## Approach 1: DSU, adding cells in rising order (Step 15.5 style)
As `t` goes `0, 1, 2, ...`, the cell with height `t` becomes usable. When it does, union it with its already-usable neighbours. The first `t` where start and end are in the same group is the answer.
- Because heights are `0 .. n²−1` and distinct, `pos[t]` tells us which cell opens at time `t`.
- **Time:** O(n² · α), **Space:** O(n²)

## Approach 2: Dijkstra with max (like [Path with minimum effort](../../15.4-shortest-paths/05-path-minimum-effort/))
Cost of a path = max height on it. Min-heap by that cost; the first time the end is popped is the answer.
- **Time:** O(n² log n)

## Dry run (approach 1)
```
0  1  2  3  4
24 23 22 21 5
12 13 14 15 16
11 17 18 19 20
10 9  8  7  6
```
Cells open in order 0, 1, 2, ... The top row and the right column through 5 connect early, but the bottom-right `6` can only be reached through the snake: `5 → 16 → 15 → ... → 12 → 11 → 10 → ... → 6`. The highest cell needed is **16**, so the answer is 16 ✅.

## Edge case
`n = 1` → answer is `grid[0][0]`.
