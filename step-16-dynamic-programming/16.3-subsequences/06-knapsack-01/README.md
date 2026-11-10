# 0/1 Knapsack

**Problem:** `n` items, item `i` has weight `wt[i]` and value `val[i]`. A bag holds at most weight `W`. Each item can be taken **at most once** (that's the "0/1"). Maximise total value.

```
wt  = [1 2 4 5]
val = [5 4 8 6], W = 5 → 13 (items 0 and 2: weight 5, value 13)
```

## Why greedy fails
"Best value per weight first" picks item 0 (5/1), then item 1 (4/2) → weight 3, can't fit 4 → value 9. Best is 13. Need DP.

## Recursion: take / not take
```
f(i, w) = best value using items 0..i with capacity w
f(0, w) = val[0] if wt[0] <= w else 0
f(i, w) = max( f(i−1, w),                              // skip item i
               val[i] + f(i−1, w − wt[i])  if wt[i] <= w ) // take item i
answer = f(n−1, W)
```

## Stages
| | Time | Space |
|---|---|---|
| Recursion | O(2ⁿ) | O(n) |
| Memo | O(n · W) | O(n · W) + stack |
| Table | O(n · W) | O(n · W) |
| Two rows | O(n · W) | O(W) |
| **One row, w from high to low** | O(n · W) | **O(W)** |

Why high to low with one row? `dp[w − wt[i]]` must still be the value **without** item `i`. Going downward, the smaller indices haven't been updated yet in this round. (Going upward would let item `i` be used twice: that's the [unbounded](../10-unbounded-knapsack/) version!)

## Dry run (one row, W = 5)
| after item | w=0 | 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|---|
| start | 0 | 0 | 0 | 0 | 0 | 0 |
| (1, 5) | 0 | 5 | 5 | 5 | 5 | 5 |
| (2, 4) | 0 | 5 | 5 | 9 | 9 | 9 |
| (4, 8) | 0 | 5 | 5 | 9 | 9 | 13 |
| (5, 6) | 0 | 5 | 5 | 9 | 9 | **13** |
