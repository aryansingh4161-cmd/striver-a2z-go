# Minimum Coins (Coin Change)

**Problem:** Coins of given denominations, **unlimited supply** of each. Fewest coins that make exactly `amount`? `-1` if impossible.

```
coins = [1 2 5], amount = 11 → 3 (5 + 5 + 1)
coins = [2], amount = 3     → -1
```

## Why greedy fails
"Biggest coin first" works for real money, not in general: `coins = [1 3 4]`, amount 6 → greedy 4+1+1 = 3 coins, best 3+3 = **2**.

## Recursion: take and **stay**
Unlimited supply changes one thing: after taking coin `i`, you may take it **again**, so you stay on index `i`.
```
f(i, t) = fewest coins from coins[0..i] making t
f(0, t) = t / coins[0] if t % coins[0] == 0 else ∞
f(i, t) = min( f(i−1, t),                      // never use coin i again
               1 + f(i, t − coins[i])  if coins[i] <= t )   // use it, stay at i
```

## Stages
- Recursion: exponential (the "stay" makes it worse than 2ⁿ)
- Memo / table: O(n · amount)
- **One row, t going upward:** O(amount) space.
  Going **upward** is exactly what allows reuse: `dp[t − c]` already includes coin `c` from this round. (Compare with the downward loop in [0/1 knapsack](../06-knapsack-01/).)

```
dp[0] = 0, dp[t > 0] = ∞
for c in coins:
    for t = c .. amount:
        dp[t] = min(dp[t], 1 + dp[t − c])
```

## Dry run (`[1 2 5]`, 11)
| after coin | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 1 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 |
| 2 | 0 | 1 | 1 | 2 | 2 | 3 | 3 | 4 | 4 | 5 | 5 | 6 |
| 5 | 0 | 1 | 1 | 2 | 2 | 1 | 2 | 2 | 3 | 3 | 2 | **3** |
