# Count Partitions with Given Difference

**Problem:** Split `arr` into two subsets `S1` and `S2` with `sum(S1) − sum(S2) = d` (`S1` is the bigger one). How many ways? (mod `1e9 + 7`)

```
arr = [5 2 6 4], d = 3 → 1 ({6, 4} = 10 and {5, 2} = 7)
```

## Turn it into "count subsets with sum K"
We know:
```
S1 + S2 = total
S1 − S2 = d
```
Subtract: `2 · S2 = total − d`, so **`S2 = (total − d) / 2`**.

Every subset that sums to `S2` gives exactly one valid partition (the rest is `S1`). So:
- If `total − d < 0` or `total − d` is **odd** → 0 ways.
- Otherwise → [count subsets with sum](../04-count-subsets-sum-k/) `(total − d) / 2`.

## Dry run
`[5 2 6 4]`, total 17, d 3 → S2 = 7 → subsets with sum 7: `{5, 2}` only → **1** ✅

Zeros matter here too (each zero can go to either side), and the counting code already handles them.

- **Time:** O(n · S2), **Space:** O(S2)

This exact reduction also solves [Target Sum](../08-target-sum/).
