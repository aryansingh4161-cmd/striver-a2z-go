# Highest and Lowest Frequency Element

**Problem:** Return the element that appears the most and the one that appears the least.

`[10 5 10 15 10 5]` → most: 10 (3 times), least: 15 (1 time)

## Approach
1. Count with a map (see [count frequency](../01-count-frequency/)).
2. Walk the counts, keeping a "current champion" for max and for min.

- **Time:** O(n), **Space:** O(k)

## Ties
The problem usually doesn't care which one you return on a tie. Go maps iterate in random order, so to be **deterministic** this solution walks the elements in first-seen order and only replaces the champion on a strictly better count. So ties go to the element seen first.
