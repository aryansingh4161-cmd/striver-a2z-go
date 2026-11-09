# Smallest Divisor Given a Threshold

**Problem (LeetCode 1283):** Pick a positive divisor `d`. Divide every element by `d`, **rounding up**, and add the results. Return the smallest `d` with that sum ≤ `threshold`.

`[1 2 5 9]`, threshold 6 → 5 (1+1+1+2 = 5)

## Checker
`sum(d) = Σ ceil(a[i] / d)`. Bigger `d` → smaller sum → monotonic ✅.

## Search space
`1 .. max(a)` (dividing by `max` makes every term 1, the smallest possible sum).

Same shape as [Koko](../03-koko-eating-bananas/).
- **Time:** O(n log max), **Space:** O(1)
