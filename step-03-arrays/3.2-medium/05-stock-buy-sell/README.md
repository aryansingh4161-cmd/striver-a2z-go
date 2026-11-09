# Best Time to Buy and Sell Stock

**Problem (LeetCode 121):** `prices[i]` is the price on day `i`. Buy once, sell once **later**. Return the maximum profit (0 if you can't profit).

`[7 1 5 3 6 4]` → 5 (buy at 1, sell at 6)

## Approach 1: All pairs (brute)
**Time O(n²)**.

## Approach 2: Track the minimum so far (optimal)
If you sell on day `i`, the best buy day is the **cheapest day before it**. So keep `minPrice` while walking:
```
minPrice = prices[0], best = 0
for p in prices:
    best = max(best, p - minPrice)
    minPrice = min(minPrice, p)
```
- **Time:** O(n), **Space:** O(1)

This "remember the best of the past" idea is a tiny form of **dynamic programming**.
