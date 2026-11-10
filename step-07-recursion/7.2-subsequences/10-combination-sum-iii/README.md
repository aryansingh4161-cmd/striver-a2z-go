# Combination Sum III

**Problem (LeetCode 216):** Find all combinations of **k distinct numbers from 1 to 9** that add up to `n`.

k = 3, n = 9 → `[1 2 6] [1 3 5] [2 3 4]`

## Approach: Loop-based backtracking over 1..9
```
f(start, rem):
    if len(path) == k: if rem == 0 save; return
    for d = start .. 9:
        if d > rem: break          // numbers only get bigger
        pick d; f(d+1, rem - d); unpick
```
Moving to `d+1` makes each number used at most once and keeps combinations increasing (so no duplicates).
- **Time:** at most C(9, k) leaves, tiny. **Space:** O(k)
