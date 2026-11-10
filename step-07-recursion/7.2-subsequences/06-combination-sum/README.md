# Combination Sum

**Problem (LeetCode 39):** Distinct positive numbers. Return all unique combinations that sum to `target`. **Each number can be used unlimited times.**

`[2 3 6 7]`, target 7 → `[2 2 3] [7]`

## Approach: Pick / not pick, but "pick" stays at the same index
```
f(i, remaining):
    if i == n: if remaining == 0 save path; return
    if a[i] <= remaining:
        pick a[i]; f(i, remaining - a[i]); unpick      // stay at i: can reuse
    f(i+1, remaining)                                   // move on
```
Staying at `i` after picking allows reuse; moving to `i+1` means "done with this number". Because we only move forward, `[2 2 3]` and `[3 2 2]` can't both appear.

- **Time:** exponential: about O(2^(target/min) × k) where k is the average combination length
- **Space:** O(target/min) depth

## Recursion sketch, target 7
`2 → 2 → 2` (rem 1, stuck) · `2 → 2 → 3` ✓ · `2 → 3 → …` too big · `7` ✓
