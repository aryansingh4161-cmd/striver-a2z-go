# Aggressive Cows

**Problem:** `stalls` are positions on a line. Place `k` cows in stalls so that the **minimum distance between any two cows is as large as possible**. Return that distance.

`stalls = [0 3 4 7 10 9]`, k = 4 → 3 (cows at 0, 3, 7, 10)

This is a **max of min** problem → binary search on the answer, looking for the **largest** distance that works.

## Checker: `canPlace(dist)`
Sort the stalls. Put the first cow in the first stall. Walk right; place the next cow in the first stall that is at least `dist` away from the last cow. Can we place `k` cows?
(Greedy: placing each cow as early as possible leaves the most room for the rest.)

## Search space
`1 .. (max − min)`. Smaller distance is easier → ✅ ✅ ✅ ❌ ❌ → find the **last** ✅ (on success, save and go **right**).

- **Time:** O(n log n) to sort + O(n × log(range)), **Space:** O(1)

## Dry run (sorted `[0 3 4 7 9 10]`), k = 4
| dist | placements | cows | ok? |
|---|---|---|---|
| 5 | 0, 7 | 2 | ❌ |
| 2 | 0, 3, 7, 9 | 4 | ✅ |
| 3 | 0, 3, 7, 10 | 4 | ✅ |
| 4 | 0, 4, 9 | 3 | ❌ → answer **3** |
