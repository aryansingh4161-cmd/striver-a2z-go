# Jump Game

**Problem (LeetCode 55):** `a[i]` = the maximum jump length from index `i`. Starting at index 0, can you reach the last index?

`[2 3 1 1 4]` → true, `[3 2 1 0 4]` → false (stuck at index 3)

## Approach 1: Recursion / DP
Try every jump. Exponential without memoisation, O(n²) with it.

## Approach 2: Track the farthest reachable index (greedy, optimal)
```
farthest = 0
for i in 0..n-1:
    if i > farthest: return false      // can't even reach i
    farthest = max(farthest, i + a[i])
return true
```
- **Time:** O(n), **Space:** O(1)

## Dry run `[3 2 1 0 4]`
farthest: 3, 3, 3, 3 → at i = 4, 4 > 3 → **false**.
