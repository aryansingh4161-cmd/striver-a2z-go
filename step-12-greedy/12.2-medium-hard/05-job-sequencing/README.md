# Job Sequencing Problem

**Problem:** Each job has a `deadline` and a `profit`, and takes 1 unit of time. Only one job at a time. A job must finish by its deadline. Maximise profit; return (number of jobs, total profit).

jobs (id, deadline, profit): `(1,4,20) (2,1,10) (3,1,40) (4,1,30)` → 2 jobs, profit 60 (jobs 3 and 1)

## Greedy rule
1. Sort jobs by **profit, highest first**.
2. For each job, place it in the **latest free slot ≤ its deadline** (doing it as late as possible keeps early slots free for other jobs).
3. If no slot is free, skip it.

- **Time:** O(n log n + n × maxDeadline), **Space:** O(maxDeadline)
- (A Disjoint Set Union can find the latest free slot in nearly O(1), giving O(n log n) total. DSU is taught in Step 15.)

## Dry run
Sorted: (3,1,40) (4,1,30) (1,4,20) (2,1,10). Slots 1..4 empty.
- job 3 → slot 1
- job 4 → slot 1 taken, nothing earlier → skip
- job 1 → slot 4
- job 2 → slot 1 taken → skip
Total **2 jobs, 60**.
