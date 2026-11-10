# Merge Intervals

Full notes are in [Step 3 hard](../../../step-03-arrays/3.3-hard/07-merge-overlapping-intervals/). It appears again here because it's a greedy algorithm:

**Sort by start**, then for each interval either **extend** the last merged one (`start <= last.end` → `last.end = max(last.end, end)`) or **start a new** one.
- **Time:** O(n log n), **Space:** O(n)

`[[1 3] [2 6] [8 10] [15 18]]` → `[[1 6] [8 10] [15 18]]`
