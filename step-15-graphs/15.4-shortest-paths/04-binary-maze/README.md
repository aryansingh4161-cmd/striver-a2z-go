# Shortest Path in a Binary Maze

**Problem:** A grid of `1` (open) and `0` (wall). Move up/down/left/right through open cells. Return the fewest steps from `src` to `dst`, or `-1`.

## Why BFS (and not Dijkstra)
Every move costs 1, so it's the [unit-weight](../01-shortest-path-unit-weights/) case: BFS level = distance. Striver also shows it with Dijkstra's idea (a queue of `(dist, cell)`); with all weights equal, the queue is already sorted by distance, so the heap isn't needed.

```
dist[src] = 0; queue = [src]
pop cell; for each of 4 neighbours:
    in grid, open, not seen → dist = dist + 1, push
    if neighbour == dst → return its dist (early exit)
```

## Dry run
```
1 1 1 1
1 1 0 1
1 1 1 1
1 1 0 0
1 0 0 1
```
src (0,1), dst (2,2): (0,1) → (1,1) → (2,1) → (2,2) = **3** ✅

## Edge cases
- `src == dst` → 0
- src or dst is a wall → -1
- dst walled off → -1

- **Time:** O(rows · cols), **Space:** O(rows · cols)
