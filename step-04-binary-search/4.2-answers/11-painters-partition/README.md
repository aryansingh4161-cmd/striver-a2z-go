# Painter's Partition

**Problem:** `boards[i]` = length of board `i`. `k` painters each paint a **contiguous** set of boards; painting 1 unit takes 1 unit of time, and painters work at the same time. Minimise the time to paint everything.

`[10 20 30 40]`, k = 2 → 60 (`[10 20 30] [40]`)

## Same pattern again
Total time = the busiest painter's total → minimise the maximum contiguous sum with at most `k` parts.
Identical to [book allocation](../09-book-allocation/) and [split array](../10-split-array-largest-sum/).

- **Time:** O(n log sum), **Space:** O(1)
