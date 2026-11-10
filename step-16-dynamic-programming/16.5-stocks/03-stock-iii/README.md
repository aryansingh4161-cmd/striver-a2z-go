# Best Time to Buy and Sell Stock III (At Most 2 Transactions)

**Problem:** Like [Stock II](../02-stock-ii/), but at most **2** complete transactions (buy + sell = 1 transaction).

```
[3 3 5 0 0 3 1 4] → 6 (buy 0 sell 3: +3, buy 1 sell 4: +3)
```

## Add one more state: transactions left
`f(i, canBuy, cap)` where `cap` = how many transactions you may still **complete**.
- A transaction is completed on **sell**, so `cap` drops by 1 when selling.
- `cap == 0` → profit 0 (you're done).
```
f(n, *, *) = 0;   f(*, *, 0) = 0
canBuy:  max( −p[i] + f(i+1, 0, cap),   f(i+1, 1, cap) )
holding: max( +p[i] + f(i+1, 1, cap−1), f(i+1, 0, cap) )
answer = f(0, 1, 2)
```

## Stages
- Recursion: O(2ⁿ)
- Memo / table `dp[n+1][2][3]`: O(n · 2 · 3) = **O(n)**
- Space optimised: `ahead[2][3]` → **O(1)**

## A compact way to see it (four running values)
Scanning left to right, keep the best balance after each action:
```
buy1  = max(buy1,  −p)          // first buy
sell1 = max(sell1, buy1 + p)    // first sell
buy2  = max(buy2,  sell1 − p)   // second buy uses first profit
sell2 = max(sell2, buy2 + p)    // second sell
answer = sell2
```
This is the same DP, written forwards. The test checks it too.

## Dry run (four values on `[3 3 5 0 0 3 1 4]`)
| p | buy1 | sell1 | buy2 | sell2 |
|---|---|---|---|---|
| 3 | −3 | 0 | −3 | 0 |
| 5 | −3 | 2 | −3 | 2 |
| 0 | 0 | 2 | 2 | 2 |
| 3 | 0 | 3 | 2 | 5 |
| 1 | 0 | 3 | 2 | 5 |
| 4 | 0 | 4 | 2 | **6** |
