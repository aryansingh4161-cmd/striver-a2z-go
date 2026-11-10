# Fractional Knapsack

**Problem:** Items with `value` and `weight`; a bag of capacity `W`. You may take **fractions** of items. Maximise total value.

values `[60 100 120]`, weights `[10 20 30]`, W = 50 → 240
(all of item 1 and 2 = 160, then 20/30 of item 3 = 80)

## Greedy rule
Sort by **value per unit weight** (`value / weight`), highest first. Take whole items while they fit; take a fraction of the next one to fill the bag exactly.
- **Time:** O(n log n), **Space:** O(n)

## Why greedy works here, but not for 0/1 knapsack
With fractions, every unit of capacity should hold the most valuable "stuff" available, so the best ratio always wins.
In **0/1 knapsack** (all or nothing) the best-ratio item might waste space: values `[60 100 120]`, weights `[10 20 30]`, W = 50 → greedy takes items 1, 2 (160) but the best is items 2, 3 (220). That one needs **DP** (Step 16).
