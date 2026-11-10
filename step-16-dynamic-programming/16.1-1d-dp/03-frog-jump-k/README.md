# Frog Jump with K Distance

**Problem:** Same as [Frog Jump](../02-frog-jump/), but from stone `i` the frog may jump to any of `i+1, i+2, ..., i+k`. Minimum total cost to reach the last stone.

## Recursion
Try all `k` previous stones:
```
f(0) = 0
f(i) = min over j = 1..k (with i−j ≥ 0) of  f(i−j) + |h[i] − h[i−j]|
```
- **Time:** O(kⁿ)

## Memoisation / Tabulation
Each of the `n` states looks at up to `k` earlier states.
- **Time:** O(n · k), **Space:** O(n)

```
dp[0] = 0
for i in 1..n-1:
    dp[i] = min over j in 1..k, i-j >= 0: dp[i-j] + |h[i] - h[i-j]|
```

## Space optimisation?
You'd need the last `k` values, so a ring buffer of size `k` gives O(k) space. When `k` can be as big as `n` this doesn't help, so the O(n) table is the usual answer.

## Dry run (`h = [10 30 40 50 20]`, k = 3)
| i | candidates | dp |
|---|---|---|
| 0 | | 0 |
| 1 | 0+20 | 20 |
| 2 | 20+10, 0+30 | 30 |
| 3 | 30+10, 20+20, 0+40 | 40 |
| 4 | 40+30, 30+20, 20+10 | **30** |
