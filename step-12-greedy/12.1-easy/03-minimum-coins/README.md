# Minimum Coins (Indian currency)

**Problem:** Using denominations `1 2 5 10 20 50 100 200 500 2000` (unlimited supply), pay amount `V` with the fewest coins/notes. Return the coins used.

V = 49 → `20 20 5 2 2` (5 coins)

## Greedy rule
Always take the **largest denomination that fits**, repeat.
- **Time:** O(V / smallest) worst case in coins taken, or O(number of denominations) with division. **Space:** O(1)

## When does greedy fail?
It works for "canonical" coin systems like Indian or US currency, but **not in general**:
coins `{1 3 4}`, V = 6 → greedy `4 1 1` (3 coins), best `3 3` (2 coins).
For arbitrary coins use DP: **Coin Change** in Step 16.
