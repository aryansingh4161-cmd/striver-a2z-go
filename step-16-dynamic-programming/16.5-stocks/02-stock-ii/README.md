# Best Time to Buy and Sell Stock II (Unlimited Transactions)

**Problem:** You may buy and sell **as many times** as you like, but hold at most **one share** at a time (sell before buying again). Max profit.

```
[7 1 5 3 6 4] → 7 (buy 1 sell 5: +4, buy 3 sell 6: +3)
```

## The DP framework for all stock problems
On each day you're in one of two states:
- **can buy** (you hold nothing)
- **can sell** (you hold one share)

`f(i, canBuy)` = best profit from day `i` to the end, given the state.
```
f(n, *) = 0                                           // no days left
if canBuy:  f(i, 1) = max( −prices[i] + f(i+1, 0),    // buy today
                            f(i+1, 1) )               // skip
else:       f(i, 0) = max( +prices[i] + f(i+1, 1),    // sell today
                            f(i+1, 0) )               // skip
answer = f(0, 1)
```
Paying when buying (`−price`) and receiving when selling (`+price`) keeps profit = sum of everything.

Stocks III, IV, cooldown and fee are all **small changes to this one recurrence**.

## Stages
- Recursion: O(2ⁿ)
- Memo / table `dp[i][2]`: O(n)
- Space optimised: only day `i+1` is needed → 2 variables (`ahead buy / ahead sell`). **O(1)**.

## Greedy (works only for this version)
Collect every upward step: `profit = Σ max(0, p[i] − p[i−1])`. Any rising stretch can be split into daily rises, so this equals the DP. The test checks both.

## Dry run (space optimised, from the last day back)
Start `aheadBuy = aheadSell = 0`.
| day (price) | buy = max(−p + aheadSell, aheadBuy) | sell = max(p + aheadBuy, aheadSell) |
|---|---|---|
| 5 (4) | max(−4, 0) = 0 | max(4, 0) = 4 |
| 4 (6) | max(−6 + 4, 0) = 0 | max(6, 4) = 6 |
| 3 (3) | max(3, 0) = 3 | max(3, 6) = 6 |
| 2 (5) | max(1, 3) = 3 | max(8, 6) = 8 |
| 1 (1) | max(7, 3) = 7 | max(4, 8) = 8 |
| 0 (7) | max(1, 7) = **7** | |
