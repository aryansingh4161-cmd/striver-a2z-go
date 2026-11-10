# Longest Repeating Character Replacement

**Problem (LeetCode 424):** Uppercase string. You may change at most `k` characters. Return the longest substring of one repeated letter you can get.

`"AABABBA"`, k = 1 → 4 (`"AABA"` → `"AAAA"`)

## Key insight
A window can be made uniform if
**changes needed = window length − (count of its most frequent letter) ≤ k**.

## Approach 1: All substrings with a frequency array
**O(26 · n²)** or O(n²).

## Approach 2: Sliding window (optimal)
Keep `count[26]` and `maxFreq` = the highest count seen in any window so far.
```
for r:
    count[s[r]]++
    maxFreq = max(maxFreq, count[s[r]])
    if (r - l + 1) - maxFreq > k:      // invalid
        count[s[l]]--; l++
    best = max(best, r - l + 1)
```
**Why don't we lower `maxFreq` when shrinking?** The answer only grows when some window has a **bigger** `maxFreq`. A stale (too high) `maxFreq` can't produce a wrong bigger answer; it just stops the window from growing until a real better frequency appears.
- **Time:** O(n), **Space:** O(26)
