# Climbing Stairs

**Problem:** You're at step 0 and want to reach step `n`. Each move climbs **1 or 2** steps. How many different ways are there?

```
n = 3 → 3 ways: 1+1+1, 1+2, 2+1
```

This is the "hello world" of DP. We'll go through all four stages of the recipe from the [Step 16 README](../../README.md).

## 1. Recursion
To land on step `n`, the last move came from `n−1` (a 1-step) or `n−2` (a 2-step). Those are separate ways, so add them:
```
f(n) = f(n−1) + f(n−2)
f(0) = 1   (standing still at the start is one way)
f(1) = 1
```
- **Time:** O(2ⁿ): the call tree doubles. `f(n−2)` gets computed from both `f(n)` and `f(n−1)`, and so on.

## 2. Memoisation (top-down)
Store `dp[i]` the first time it's computed; return it on later calls.
- **Time:** O(n) (each `i` once), **Space:** O(n) table + O(n) recursion stack

## 3. Tabulation (bottom-up)
Fill `dp[0], dp[1]`, then loop `i = 2..n`: `dp[i] = dp[i−1] + dp[i−2]`.
- **Time:** O(n), **Space:** O(n), no recursion

## 4. Space optimisation
`dp[i]` only looks back two cells. Keep two variables `prev2, prev1`.
- **Time:** O(n), **Space:** O(1)

## Dry run (n = 5)
| i | 0 | 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|---|
| ways | 1 | 1 | 2 | 3 | 5 | **8** |

Yes, it's Fibonacci shifted by one.
