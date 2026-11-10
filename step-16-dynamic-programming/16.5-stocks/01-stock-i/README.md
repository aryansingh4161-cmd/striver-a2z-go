# Best Time to Buy and Sell Stock I

**Problem:** `prices[i]` is the price on day `i`. Buy **once** and sell **once** later. Max profit (0 if you can't profit).

```
[7 1 5 3 6 4] → 5 (buy at 1, sell at 6)
[7 6 4 3 1]   → 0
```

## Brute force
Try every (buy day, sell day) pair: O(n²).

## Optimal: remember the cheapest day so far
If you sell on day `i`, the best buy day is the **cheapest day before `i`**. So walk once, keeping `minSoFar`:
```
minSoFar = prices[0], best = 0
for i = 1..n−1:
    best = max(best, prices[i] − minSoFar)
    minSoFar = min(minSoFar, prices[i])
```
This is DP in its simplest form: `minSoFar` is the stored answer to "cheapest price up to here".

## Dry run
| day | price | minSoFar | profit if sold | best |
|---|---|---|---|---|
| 0 | 7 | 7 | – | 0 |
| 1 | 1 | 1 | −6 | 0 |
| 2 | 5 | 1 | 4 | 4 |
| 3 | 3 | 1 | 2 | 4 |
| 4 | 6 | 1 | 5 | **5** |
| 5 | 4 | 1 | 3 | 5 |

- **Time:** O(n), **Space:** O(1)
