# Longest Substring Without Repeating Characters

**Problem (LeetCode 3):** Length of the longest substring with all distinct characters.

`"abcabcbb"` → 3 (`"abc"`), `"bbbbb"` → 1, `"pwwkew"` → 3 (`"wke"`)

## Approach 1: All substrings
For each start, extend while characters are new (use a `[256]bool`). **O(n²)**.

## Approach 2: Sliding window + last-seen index (optimal)
Window `[l, r]` always has distinct characters. Store `last[c]` = the last index where `c` appeared.
When `s[r]` was already seen **inside the window** (`last[s[r]] >= l`), jump `l` to `last[s[r]] + 1`. No step-by-step shrinking needed.
```
for r in 0..n-1:
    if last[s[r]] >= l: l = last[s[r]] + 1
    last[s[r]] = r
    best = max(best, r - l + 1)
```
- **Time:** O(n), **Space:** O(256)

## Dry run `"abcabcbb"`
| r | char | last seen | l | window | best |
|---|---|---|---|---|---|
| 0 | a | – | 0 | a | 1 |
| 1 | b | – | 0 | ab | 2 |
| 2 | c | – | 0 | abc | 3 |
| 3 | a | 0 | 1 | bca | 3 |
| 4 | b | 1 | 2 | cab | 3 |
| 5 | c | 2 | 3 | abc | 3 |
| 6 | b | 4 | 5 | cb | 3 |
| 7 | b | 6 | 7 | b | 3 |

Initialise `last` to −1 so "never seen" is never inside the window.
