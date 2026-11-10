# Best Time to Buy and Sell Stock with Cooldown

**Problem:** Unlimited transactions like [Stock II](../02-stock-ii/), but after you **sell**, you can't buy on the **next day** (one day of cooldown).

```
[1 2 3 0 2] → 3 (buy 1, sell 2, cooldown, buy 0, sell 2)
```

## One change to Stock II
When you sell on day `i`, the next day you're allowed to buy is `i + 2`:
```
canBuy:  max( −p[i] + f(i+1, hold),  f(i+1, canBuy) )
holding: max( +p[i] + f(i+2, canBuy), f(i+1, hold) )    // ← i+2 instead of i+1
f(i >= n, *) = 0
```

## Stages
- Recursion: exponential
- Table `dp[n+2][2]` (two extra rows of zeros so `i + 2` is always valid): O(n)
- Space optimised: keep day `i+1` **and** day `i+2` → O(1)

## Dry run (table from the end, `[1 2 3 0 2]`)
| i (p) | buy | hold |
|---|---|---|
| 5, 6 | 0 | 0 |
| 4 (2) | max(−2 + 0, 0) = 0 | max(2 + 0, 0) = 2 |
| 3 (0) | max(0 + 2, 0) = 2 | max(0 + 0, 2) = 2 |
| 2 (3) | max(−3 + 2, 2) = 2 | max(3 + 0, 2) = 3 |
| 1 (2) | max(−2 + 3, 2) = 2 | max(2 + 2, 3) = 4 |
| 0 (1) | max(−1 + 4, 2) = **3** | |
