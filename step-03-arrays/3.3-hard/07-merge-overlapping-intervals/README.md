# Merge Overlapping Intervals

**Problem (LeetCode 56):** Merge all overlapping intervals.

`[[1 3] [2 6] [8 10] [15 18]]` → `[[1 6] [8 10] [15 18]]`

## Approach: Sort + one pass (optimal)
1. **Sort by start.** Now overlapping intervals sit next to each other.
2. Walk through. Compare each interval with the **last merged** one:
   - `start <= last.end` → overlap: `last.end = max(last.end, end)`
   - otherwise → start a new merged interval.

- **Time:** O(n log n) for the sort + O(n), **Space:** O(n) for the answer

Use `max` for the end: `[1 10]` and `[2 3]` overlap, but `[2 3]` ends earlier.

The brute force (for each interval, scan all later ones) is O(n²) and gives the same result. The sort makes it unnecessary.
