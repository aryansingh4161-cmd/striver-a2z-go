# Best Time to Buy and Sell Stock with Transaction Fee

**Problem:** Unlimited transactions, but every completed transaction costs a `fee`. Max profit.

```
prices = [1 3 2 8 4 9], fee = 2 → 8
buy 1 sell 8: 7 − 2 = 5;  buy 4 sell 9: 5 − 2 = 3;  total 8
```

## One change to Stock II
Pay the fee once per transaction. The simplest place is **when selling**:
```
canBuy:  max( −p[i] + f(i+1, hold),        f(i+1, canBuy) )
holding: max( +p[i] − fee + f(i+1, canBuy), f(i+1, hold) )
```
(Charging it on buy works too, just don't charge both.)

Why greedy (sum of rises) fails now: many small rises each pay the fee. In the example, `1→3` (+2) minus fee is 0, and it's better to hold through `3→2→8`.

## Stages
- Recursion: O(2ⁿ) → table O(n) → **two variables O(1)**

## Dry run (space optimised, from the end)
| p | buy | hold |
|---|---|---|
| 9 | 0 | 7 |
| 4 | 3 | 7 |
| 8 | 3 | 9 |
| 2 | 7 | 9 |
| 3 | 7 | 9 |
| 1 | **8** | |
