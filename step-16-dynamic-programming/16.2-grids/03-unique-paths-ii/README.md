# Unique Paths II (with Obstacles)

**Problem:** Same as [Unique Paths](../02-unique-paths/), but some cells are blocked (`1` = obstacle, `0` = free). Count paths from top-left to bottom-right moving right/down, never stepping on an obstacle.

```
0 0 0
0 1 0
0 0 0   → 2
```

## Only one change
A blocked cell has **0 ways** to reach it, so it contributes nothing to its neighbours:
```
f(i, j) = 0                       if grid[i][j] is blocked
f(0, 0) = 1                       (if not blocked)
f(i, j) = f(i−1, j) + f(i, j−1)   otherwise
```
The maths formula from part I no longer works; the DP handles obstacles naturally.

## Stages
Same as part I: recursion O(2^(m+n)) → memo O(m·n) → table O(m·n) → one row O(n) space.

## Dry run
```
1 1 1
1 0 1
1 1 2   → 2 ✅
```
(the middle is 0 because it's blocked; the first row/column stay 1 because nothing blocks them)

## Edge cases
- Start or end blocked → 0.
- An obstacle in the first row blocks every cell to its right in that row (they can only be reached from the left). The one-row loop handles this automatically because blocked cells reset to 0.
