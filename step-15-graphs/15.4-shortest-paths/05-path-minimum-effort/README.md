# Path With Minimum Effort

**Problem:** `heights` is a grid. Go from top-left to bottom-right (4 directions). The **effort** of a path is the **largest** absolute height difference between two consecutive cells on it. Find the minimum possible effort.

```
1 2 2
3 8 2
5 3 5
```
Path 1→3→5→3→5 has differences 2,2,2,2 → effort **2** (best).

## Approach 1: Binary search on the answer + BFS
"Can I reach the end using only steps with difference ≤ k?" is a yes/no question that's monotonic in k. Binary search k in `[0, maxHeight]`, check each with BFS.
- **Time:** O(R·C · log(maxH))

## Approach 2: Dijkstra with a different "distance" (optimal, Striver's way)
Let `effort[cell]` = smallest possible max-difference to reach that cell. Moving from `a` to `b`:
```
newEffort = max(effort[a], |h[a] - h[b]|)
```
This "cost" never goes down as the path grows (max can only stay or rise), which is exactly the property Dijkstra needs. So run Dijkstra with `max` instead of `+`.

The first time we **pop** the target, its effort is final → return it.

## Dry run (start)
- pop (0, (0,0)); neighbours (0,1): max(0,1)=1, (1,0): max(0,2)=2
- pop (1,(0,1)); (0,2): 1, (1,1): 6
- pop (1,(0,2)); (1,2): 1
- pop (1,(1,2)); (2,2): max(1,3)=3
- pop (2,(1,0)); (2,0): 2, (1,1): 5
- pop (2,(2,0)); (2,1): 2
- pop (2,(2,1)); (2,2): max(2,2)=2 < 3 → update
- pop (2,(2,2)) → **2** ✅

- **Time:** O(R·C · log(R·C)), **Space:** O(R·C)
