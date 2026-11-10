# Task Scheduler

**Problem (LeetCode 621):** Tasks labelled A–Z, each takes 1 unit. Two runs of the **same** task need at least `n` units between them (CPU can idle). Minimum total time?

`[A A A B B B]`, n = 2 → 8 (`A B idle A B idle A B`)

## Approach 1: Simulation with a max-heap
Each round of `n + 1` slots: pop up to `n + 1` tasks with the highest remaining counts, run each once, push back those still remaining. Add `n + 1` per round, except the last round, which only counts the tasks actually run.
- **Time:** O(total × log 26), **Space:** O(26)

## Approach 2: Formula (optimal)
Let `maxFreq` = the highest task count, and `countMax` = how many tasks have that count.
The most frequent task forms `maxFreq − 1` full blocks of length `n + 1`, plus a final block containing the `countMax` most frequent tasks:
```
answer = max(len(tasks), (maxFreq − 1) × (n + 1) + countMax)
```
The `max` handles the case where there are so many different tasks that no idling is needed.
- **Time:** O(len(tasks)), **Space:** O(26)

Check: `[A A A B B B]`, n = 2 → (3−1)×3 + 2 = 8 ✓
