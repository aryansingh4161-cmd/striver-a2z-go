# Best Time to Buy and Sell Stock IV (At Most K Transactions)

**Problem:** Same as [Stock III](../03-stock-iii/), but at most `k` transactions.

## Approach 1: replace 2 with k
`f(i, canBuy, cap)` with `cap` from `k` down to 0. Exactly the Stock III code with `k` in place of 2.
- **Time:** O(n · 2 · k), **Space:** O(2 · k) with the "ahead" arrays

## Approach 2: one "transaction number" instead of (canBuy, cap)
Number the actions: `0 = buy #1, 1 = sell #1, 2 = buy #2, ... , 2k − 1 = sell #k`. Even numbers are buys, odd numbers are sells.
```
f(i, t) = 0                                          if i == n or t == 2k
even t: max( −p[i] + f(i+1, t+1), f(i+1, t) )
odd t:  max( +p[i] + f(i+1, t+1), f(i+1, t) )
answer = f(0, 0)
```
One dimension less, same complexity. Both are in the code.

## Edge case
If `k ≥ n / 2`, the limit never matters: it becomes [Stock II](../02-stock-ii/) (unlimited). Not needed for correctness, but a nice speed-up for huge `k`.

## Dry run (`k = 2`, `[3 2 6 5 0 3]`)
Buy 2 sell 6 (+4), buy 0 sell 3 (+3) → **7** ✅
