# Combination Sum II

**Problem (LeetCode 40):** Numbers may repeat. Each **position** can be used **at most once**. Return unique combinations summing to `target` (no duplicate combinations).

`[10 1 2 7 6 1 5]`, target 8 → `[1 1 6] [1 2 5] [1 7] [2 6]`

## Approach 1: Pick / not pick + a set to remove duplicates
Works, but generates duplicates first and dedupes later: O(2ⁿ × k log …). Slow.

## Approach 2: Sort + loop over the next choice, skipping equal values (optimal)
Think of each level as "**which number goes next** in the combination?":
```
sort(a)
f(start, rem):
    if rem == 0: save path; return
    for j = start .. n-1:
        if j > start and a[j] == a[j-1]: continue   // same value already tried at this level
        if a[j] > rem: break                         // sorted: everything after is bigger
        pick a[j]; f(j+1, rem - a[j]); unpick
```
**Why the skip works:** at one level, choosing the first `1` or the second `1` as "the next number" leads to identical combinations. Trying only the first one at each level removes duplicates without a set. (Deeper levels can still use the second `1`, which is how `[1 1 6]` is found.)

- **Time:** O(2ⁿ × k) worst case, **Space:** O(n) depth
