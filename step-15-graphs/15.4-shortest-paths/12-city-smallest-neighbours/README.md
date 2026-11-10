# City With the Smallest Number of Neighbours at a Threshold Distance

**Problem:** `n` cities, undirected weighted roads. For each city count how many **other** cities are reachable with total distance `≤ threshold`. Return the city with the **smallest** count; on a tie, return the one with the **largest** number.

## Idea
We need distances between **all pairs** → [Floyd-Warshall](../11-floyd-warshall/) (n ≤ 100 in the problem, so O(n³) is fine). Then count per city.

```
run Floyd-Warshall
best, bestCity = INF, -1
for city in 0..n-1:
    cnt = number of j != city with dist[city][j] <= threshold
    if cnt <= best:          // "<=" so a later (larger) city wins ties
        best, bestCity = cnt, city
```

## Dry run
`n = 4`, roads `0-1 (3), 1-2 (1), 1-3 (4), 2-3 (1)`, threshold 4
| city | reachable (≤ 4) | count |
|---|---|---|
| 0 | 1 (3), 2 (4) | 2 |
| 1 | 0, 2, 3 (2) | 3 |
| 2 | 0 (4), 1, 3 | 3 |
| 3 | 1 (2), 2 (1) | 2 |
Tie between 0 and 3 → larger → **3** ✅

- **Time:** O(n³), **Space:** O(n²)

Alternative: Dijkstra from every city: O(n · E log n), better for big sparse graphs.
