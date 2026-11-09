# Capacity to Ship Packages Within D Days

**Problem (LeetCode 1011):** Packages must be shipped **in the given order**. Each day the ship loads packages until the next one would exceed its capacity. Find the minimum capacity to ship everything in `days` days.

`[1 2 3 4 5 6 7 8 9 10]`, days = 5 → 15

## Checker: `daysNeeded(cap)`
Greedy: keep adding packages to today's load; when the next doesn't fit, start a new day.

## Search space
- **Low** = `max(weights)` (the heaviest package must fit).
- **High** = `sum(weights)` (ship everything in one day).
Bigger capacity → fewer days → monotonic ✅.

- **Time:** O(n × log(sum)), **Space:** O(1)

## Check capacity 15
Day 1: 1+2+3+4+5 = 15 · Day 2: 6+7 = 13 · Day 3: 8 · Day 4: 9 · Day 5: 10 → 5 days ✅
Capacity 14: 1..4 = 10 (+5 = 15 > 14), 5+6 = 11, 7, 8, 9, 10 → 6 days ❌. Answer **15**.

This "split into ≤ k contiguous groups, minimise the largest group" shape comes back in book allocation, split array and painter's partition.
