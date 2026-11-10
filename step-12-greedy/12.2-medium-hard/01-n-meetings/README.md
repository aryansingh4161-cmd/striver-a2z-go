# N Meetings in One Room

**Problem:** Meetings with `start[i]` and `end[i]`. One room. A meeting can start only **after** the previous one ends (strictly: `start > lastEnd`). Return the maximum number of meetings (and which ones).

start `[1 3 0 5 8 5]`, end `[2 4 6 7 9 9]` → 4 (meetings 1, 2, 4, 5 in 1-based order)

## Greedy rule: earliest finish first
Sort meetings by **end time**. Take a meeting if it starts after the last taken one ends.

**Why end time?** Finishing as early as possible leaves the most room for everything after. (Sorting by start time or by duration fails; try `[0, 10]`, `[1, 2]`, `[3, 4]`.)

- **Time:** O(n log n), **Space:** O(n)

This is the classic **activity selection** problem. [Non-overlapping intervals](../11-non-overlapping-intervals/) is the same idea.
