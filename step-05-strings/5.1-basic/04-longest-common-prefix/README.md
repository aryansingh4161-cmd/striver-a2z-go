# Longest Common Prefix

**Problem (LeetCode 14):** Return the longest prefix shared by all strings.

`["flower" "flow" "flight"]` → `"fl"`

## Approach 1: Vertical scan (optimal)
Compare column by column: character `i` of the first word against character `i` of every other word. Stop at the first mismatch or when any word ends.
- **Time:** O(total characters), **Space:** O(1)

## Approach 2: Sort, compare first and last
After sorting, the first and last strings are the **most different**. Their common prefix is shared by everything in between.
- **Time:** O(n log n × L) for the sort
