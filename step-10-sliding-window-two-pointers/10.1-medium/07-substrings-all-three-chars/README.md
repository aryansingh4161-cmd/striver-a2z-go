# Number of Substrings Containing All Three Characters

**Problem (LeetCode 1358):** String of `a`, `b`, `c`. Count substrings that contain at least one of each.

`"abcabc"` → 10

## Approach 1: All substrings
For each start, extend until all three appear; then **every longer substring** also works, so add `n − j` and stop. **O(n²)** worst case.

## Approach 2: Last-seen positions (optimal)
For each end `r`, track `last[a]`, `last[b]`, `last[c]`. A substring ending at `r` contains all three iff it starts at or before `min(last)`. So there are `min(last) + 1` valid starts.
```
last = [-1, -1, -1]
for r: last[s[r]] = r
       count += min(last) + 1     // adds 0 while some letter hasn't appeared
```
- **Time:** O(n), **Space:** O(1)

## Dry run `"abcabc"`
| r | last a,b,c | min + 1 | total |
|---|---|---|---|
| 0 | 0,−1,−1 | 0 | 0 |
| 1 | 0,1,−1 | 0 | 0 |
| 2 | 0,1,2 | 1 | 1 |
| 3 | 3,1,2 | 2 | 3 |
| 4 | 3,4,2 | 3 | 6 |
| 5 | 3,4,5 | 4 | **10** |
