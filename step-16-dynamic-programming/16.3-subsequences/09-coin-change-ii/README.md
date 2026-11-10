# Coin Change II (Number of Ways)

**Problem:** Unlimited coins of each denomination. How many **combinations** make `amount`? (Order doesn't matter: `1+2` and `2+1` are the same.)

```
coins = [1 2 5], amount = 5 → 4
5 · 2+2+1 · 2+1+1+1 · 1+1+1+1+1
```

## Recursion
Same take-and-stay as [Minimum Coins](../07-minimum-coins/), but count instead of minimise:
```
f(0, t) = 1 if t % coins[0] == 0 else 0
f(i, t) = f(i−1, t) + (coins[i] <= t ? f(i, t − coins[i]) : 0)
```

## Why this counts combinations (not orderings)
Coins are decided **in index order**: we finish deciding how many of coin `i` before moving to `i−1`. So each combination is built in exactly one way.

In the 1D table, that's why the **coin loop is outside** and the amount loop inside:
```
ways[0] = 1
for c in coins:            // outer: coins
    for t = c .. amount:   // inner: amounts, upward (reuse allowed)
        ways[t] += ways[t − c]
```
Swapping the loops (amount outside) would count **orderings** instead (that's "Combination Sum IV").

## Dry run (`[1 2 5]`, 5)
| after coin | 0 | 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|---|
| 1 | 1 | 1 | 1 | 1 | 1 | 1 |
| 2 | 1 | 1 | 2 | 2 | 3 | 3 |
| 5 | 1 | 1 | 2 | 2 | 3 | **4** |

- **Time:** O(n · amount), **Space:** O(amount)
