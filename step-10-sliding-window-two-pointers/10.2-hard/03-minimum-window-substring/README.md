# Minimum Window Substring

**Problem (LeetCode 76):** Return the shortest substring of `s` that contains every character of `t` (with multiplicity). `""` if none.

`s = "ADOBECODEBANC"`, `t = "ABC"` → `"BANC"`

## Approach 1: All substrings
For each start, extend until all needed characters are covered. **O(n² )** (with O(1) coverage checks using counts).

## Approach 2: Sliding window, "shortest valid" pattern (optimal)
1. `need[c]` = how many of `c` we still need (from `t`). `missing` = `len(t)`.
2. Expand `r`: if `need[s[r]] > 0`, it's a useful character → `missing--`. Then `need[s[r]]--` (can go negative = extra copies).
3. **While `missing == 0`** (window is valid):
   - record it if shorter than the best,
   - remove `s[l]`: `need[s[l]]++`; if it becomes `> 0`, we just lost a needed character → `missing++`. `l++`.

- **Time:** O(n + m), **Space:** O(256)

## Pattern reminder
| Want | Shrink when | Update answer |
|---|---|---|
| longest valid | window becomes **invalid** | after shrinking |
| shortest valid | window is **valid** | inside the shrink loop |
