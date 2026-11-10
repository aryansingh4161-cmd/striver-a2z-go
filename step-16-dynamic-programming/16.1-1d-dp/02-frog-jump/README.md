# Frog Jump

**Problem:** A frog is on stone `0` and wants stone `n−1`. From stone `i` it can jump to `i+1` or `i+2`. A jump from `i` to `j` costs `|height[i] − height[j]|`. Find the minimum total cost.

```
heights = [10 20 30 10] → 20
0 → 1 (10) → 3 (10) = 20
```

## Why greedy fails
"Always take the cheaper next jump" can trap you. Heights `[30 10 60 10 60 50]`:
- Greedy: 0→1 (20), 1→3 (0), 3→5 (40) = **60**
- Best: 0→2 (30), 2→4 (0), 4→5 (10) = **40**

The expensive first jump pays off later. We need to try **both** options everywhere → DP.

## 1. Recursion
`f(i)` = min cost to reach stone `i` from stone 0.
```
f(0) = 0
f(i) = min( f(i−1) + |h[i] − h[i−1]|,
            f(i−2) + |h[i] − h[i−2]|   if i ≥ 2 )
```
- **Time:** O(2ⁿ)

## 2. Memoisation
Cache `dp[i]`. **Time** O(n), **Space** O(n) + stack.

## 3. Tabulation
`dp[0] = 0`; for `i = 1..n−1` take the min of the two options.

## 4. Space optimisation
Only `dp[i−1]` and `dp[i−2]` are used → two variables. **Space** O(1).

## Dry run (`[10 20 30 10]`)
| i | from i−1 | from i−2 | dp |
|---|---|---|---|
| 0 | | | 0 |
| 1 | 0 + 10 | – | 10 |
| 2 | 10 + 10 | 0 + 20 | 20 |
| 3 | 20 + 20 | 10 + 10 | **20** |
