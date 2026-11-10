# Jump Game II

**Problem (LeetCode 45):** Same jumps as [Jump Game](../02-jump-game/); the end is always reachable. Return the **minimum number of jumps**.

`[2 3 1 1 4]` → 2 (0 → 1 → 4)

## Approach 1: Recursion / DP
Try every jump length; O(n²) with memoisation.

## Approach 2: BFS by levels (greedy, optimal)
Think of it as levels: with `j` jumps you can reach a **range** of indexes `[l, r]`. With `j+1` jumps you reach up to `max(i + a[i])` for `i` in that range.
```
jumps = 0, l = r = 0
while r < n-1:
    farthest = max(i + a[i]) for i in l..r
    l = r + 1
    r = farthest
    jumps++
```
- **Time:** O(n) (each index is scanned once), **Space:** O(1)

## Dry run `[2 3 1 1 4]`
| jumps | range | farthest |
|---|---|---|
| 0 | [0,0] | 2 |
| 1 | [1,2] | 4 |
| 2 | [3,4] → reached the end |
