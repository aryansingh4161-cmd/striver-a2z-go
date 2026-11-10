# Partition Equal Subset Sum

**Problem:** Can `nums` be split into **two subsets with equal sums**?

```
[1 5 11 5] → true ({11} and {1, 5, 5})
[1 2 3 5]  → false
```

## Reduce to subset sum
If the two halves are equal, each is `total / 2`.
- If `total` is **odd**, impossible → false.
- Otherwise: is there a subset with sum `total / 2`? The rest automatically sums to the other half.

That's exactly [Subset Sum](../01-subset-sum/) with `target = total / 2`.

## Dry run
`[1 5 11 5]`: total 22 → target 11 → `{11}` works → **true** ✅
`[1 2 3 5]`: total 11 is odd → **false** ✅

- **Time:** O(n · total/2), **Space:** O(total/2) with the one-row version
