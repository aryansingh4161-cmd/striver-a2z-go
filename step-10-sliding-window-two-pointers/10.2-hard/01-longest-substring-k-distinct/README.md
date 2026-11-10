# Longest Substring with At Most K Distinct Characters

**Problem:** Return the length of the longest substring with at most `k` distinct characters.

`"aaabbccd"`, k = 2 → 5 (`"aaabb"`)

## Approach 1: All substrings
Extend from each start with a count map; stop when distinct > k. **O(n²)**.

## Approach 2: Sliding window (optimal)
Expand `r`, `count[s[r]]++`. While `len(count) > k`: decrease `count[s[l]]` (delete at 0), `l++`. Track `r − l + 1`.
- **Time:** O(n), **Space:** O(k)

This is the general form of [fruit into baskets](../../10.1-medium/03-fruit-into-baskets/) (k = 2).
