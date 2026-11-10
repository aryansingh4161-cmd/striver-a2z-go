# Hand of Straights

**Problem (LeetCode 846):** Can the cards be split into groups of `groupSize` **consecutive** values?

`[1 2 3 6 2 3 4 7 8]`, groupSize 3 → true (`123 234 678`)

## Approach: Smallest card must start a group (greedy)
The smallest remaining card can't be in the middle of a group (nothing smaller is left), so it **must start** one.
1. Count each value; get the distinct values in sorted order (or use a min-heap).
2. For each value `v` in increasing order with `count[v] > 0`: it starts `c = count[v]` groups. Each of `v, v+1, …, v+groupSize−1` must have at least `c` cards; subtract `c` from each. If any is short → false.

- **Time:** O(n log n) for sorting (+ O(n × groupSize) worst case for the checks), **Space:** O(n)

Quick check: `len(hand) % groupSize != 0` → false immediately.
