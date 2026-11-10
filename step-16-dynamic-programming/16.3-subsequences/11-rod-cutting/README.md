# Rod Cutting

**Problem:** A rod of length `n`. `price[i]` is the price of a piece of length `i + 1`. Cut the rod into pieces (any number, any lengths summing to `n`) to maximise total price.

```
n = 5, price = [2 5 7 8 10] → 12 (lengths 2 + 2 + 1: 5 + 5 + 2)
```

## It's unbounded knapsack in disguise
- "Items" = piece lengths `1..n` with value `price[len − 1]`.
- "Capacity" = rod length `n`.
- You can cut the same length many times → **unbounded**.
- The pieces must use the **whole** rod, but length-1 pieces always exist, so leftovers can always be filled. Maximising "≤ n" equals "exactly n" here.

```
f(0, L) = L * price[0]                  // only length-1 pieces
f(i, L) = max( f(i−1, L),
               price[i] + f(i, L − (i+1))   if i+1 <= L )
```

## One row
```
dp[0..n] = 0
for len = 1..n:
    for L = len..n:
        dp[L] = max(dp[L], price[len−1] + dp[L − len])
```

## Dry run (n = 5)
| after length | L=1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|
| 1 (2) | 2 | 4 | 6 | 8 | 10 |
| 2 (5) | 2 | 5 | 7 | 10 | 12 |
| 3 (7) | 2 | 5 | 7 | 10 | 12 |
| 4 (8) | 2 | 5 | 7 | 10 | 12 |
| 5 (10) | 2 | 5 | 7 | 10 | **12** |

- **Time:** O(n²), **Space:** O(n)
