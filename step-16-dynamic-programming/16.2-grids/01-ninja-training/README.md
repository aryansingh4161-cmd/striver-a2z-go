# Ninja's Training

**Problem:** For `n` days, each day the ninja does one of 3 activities. `points[day][act]` is the merit for doing `act` on `day`. He can't do the **same activity two days in a row**. Maximise total points.

```
points = [[1 2 5],
          [3 1 1],
          [3 3 3]]   → 11  (5, then 3, then 3)
```

## Why greedy fails
Taking the best activity each day can block a much better choice tomorrow: `[[10 50 1],[5 100 11]]` → greedy takes 50 then can't take 100 (gets 61); best is 10 + 100 = 110.

## State: (day, last activity)
What matters for today's choice is **which activity was done yesterday**. So the state is `f(day, last)` = best total for days `0..day` if the activity **after** `day` is `last` (so `day` can't use `last`). `last = 3` means "no restriction".

```
f(day, last) = max over act != last of  points[day][act] + f(day−1, act)
f(0, last)   = max over act != last of  points[0][act]
answer = f(n−1, 3)
```

## Stages
- **Recursion:** O(3ⁿ)
- **Memo:** states `n × 4`, each tries 3 → **O(n · 4 · 3)**, Space O(n · 4) + stack
- **Tabulation:** `dp[day][last]`, fill day 0, then each next day from the previous.
- **Space optimised:** only the previous day's 4 values are needed → `prev[4]`. **Space O(1)**.

## Dry run (tabulation, example)
`dp[day][last]` for last = 0, 1, 2, 3:
| day | last=0 | last=1 | last=2 | last=3 |
|---|---|---|---|---|
| 0 | max(2,5)=5 | max(1,5)=5 | max(1,2)=2 | 5 |
| 1 | max(1+5, 1+2)=6 | max(3+5, 1+2)=8 | max(3+5, 1+5)=8 | 8 |
| 2 | max(3+8, 3+8)=11 | 11 | 11 | **11** |
