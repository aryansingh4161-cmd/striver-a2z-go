# Unbounded Knapsack

**Problem:** Like [0/1 knapsack](../06-knapsack-01/), but you may take **any number of copies** of each item. Maximise value within capacity `W`.

```
wt = [2 4 6], val = [5 11 13], W = 10 → 27 (4 + 4 + 2: 11 + 11 + 5)
```

## The only change from 0/1: stay on the same item after taking it
```
f(0, w) = (w / wt[0]) * val[0]          // fill with as many copies of item 0 as fit
f(i, w) = max( f(i−1, w),
               val[i] + f(i, w − wt[i])   if wt[i] <= w )   // note: f(i, ...) not f(i−1, ...)
```

## 0/1 vs unbounded, side by side
| | take item i | one-row loop over w |
|---|---|---|
| 0/1 | `f(i−1, w − wt)` | **downward** |
| unbounded | `f(i, w − wt)` | **upward** |

## Stages
- Recursion: exponential
- Memo / table: O(n · W)
- One row (w upward): **O(W)** space

## Dry run (one row, W = 10)
- item (2, 5): dp[w] = 5·⌊w/2⌋ → `0 0 5 5 10 10 15 15 20 20 25`
- item (4, 11): dp[4] = 11, dp[6] = 16, dp[8] = 22, dp[10] = 27
- item (6, 13): dp[10] = max(27, 13 + dp[4] = 24) = **27** ✅
