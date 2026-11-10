# Shortest Job First (SJF) Scheduling

**Problem:** All jobs arrive at time 0 with burst times `bt`. Run them one at a time (no interruption). Minimise the **average waiting time** (time a job waits before it starts). Return the average (floored, as on GFG).

`[4 3 7 1 2]` → 4
order 1, 2, 3, 4, 7 → waits 0, 1, 3, 6, 10 → total 20 → average 4

## Greedy rule
Run the **shortest job first**. Every job's length is added to the waiting time of **all jobs after it**, so short jobs should go early.
1. Sort burst times ascending.
2. Waiting time of job `i` = sum of all earlier burst times.
- **Time:** O(n log n), **Space:** O(1) besides sorting
