# Sort Characters by Frequency

**Problem (LeetCode 451):** Rearrange the string so characters appear in decreasing order of frequency (ties in any order).

`"tree"` → `"eert"` or `"eetr"`

## Approach 1: Count + sort (simple)
1. Count each character (`[256]int`).
2. Make a list of characters that appear, sort by count descending.
3. Build the answer by repeating each character `count` times.
- **Time:** O(n + k log k) where k ≤ 256 distinct characters, **Space:** O(n)

## Approach 2: Bucket sort
`buckets[f]` = characters with frequency `f`. Walk `f` from n down to 1. **O(n)**.

The solution breaks ties by character so the output is deterministic (easier to test).
