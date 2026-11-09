# Sort an Array of 0s, 1s and 2s

**Problem (LeetCode 75, "Sort Colors"):** Sort in place without a library sort.

`[2 0 2 1 1 0]` → `[0 0 1 1 2 2]`

## Approach 1: Sort (brute)
Any sort. **O(n log n)**.

## Approach 2: Count (better)
Count 0s, 1s and 2s, then overwrite. **Time O(2n), Space O(1)**. Two passes.

## Approach 3: Dutch National Flag (optimal, one pass)
Three pointers split the array into four zones:
```
[0 .. low-1]   all 0s
[low .. mid-1] all 1s
[mid .. high]  unknown (still to check)
[high+1 .. ]   all 2s
```
Look at `a[mid]`:
- `0` → swap with `a[low]`, `low++`, `mid++`
- `1` → `mid++`
- `2` → swap with `a[high]`, `high--` (**don't** move `mid`: the swapped-in value is still unknown)

Stop when `mid > high`.
- **Time:** O(n), one pass, **Space:** O(1)

## Dry run `[2 0 1]`
| low | mid | high | a[mid] | action | array |
|---|---|---|---|---|---|
| 0 | 0 | 2 | 2 | swap mid,high; high-- | 1 0 2 |
| 0 | 0 | 1 | 1 | mid++ | 1 0 2 |
| 0 | 1 | 1 | 0 | swap low,mid; low++, mid++ | 0 1 2 |
| 1 | 2 | 1 | stop | | **0 1 2** |
