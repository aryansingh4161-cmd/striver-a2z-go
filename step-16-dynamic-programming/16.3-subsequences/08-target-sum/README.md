# Target Sum

**Problem:** Put a `+` or `−` in front of every number in `nums`. How many ways does the expression equal `target`?

```
nums = [1 1 1 1 1], target = 3 → 5
(−1+1+1+1+1, +1−1+1+1+1, ...)
```

## Recursion (direct)
```
f(i, sum) = f(i−1, sum + nums[i]) + f(i−1, sum − nums[i])
```
Works, but sums can be negative, which makes memoising with an array awkward (you'd shift by `total`).

## The neat reduction
Numbers with `+` form group `P`, numbers with `−` form group `N`:
```
P − N = target
P + N = total
→ N = (total − target) / 2
```
So it's **exactly** [Count partitions with given difference](../05-count-partitions-difference/) with `d = target`: count subsets with sum `(total − target) / 2`.
- If `total − target` is negative or odd → 0.
- Zeros: each zero can be `+0` or `−0`: 2 ways. The counting code already doubles for zeros.

## Dry run
`[1 1 1 1 1]`, total 5, target 3 → N = 1 → subsets with sum 1: choose which single 1 → **5** ✅

- **Time:** O(n · N), **Space:** O(N)

The test also runs the direct `±` recursion to double-check the reduction.
