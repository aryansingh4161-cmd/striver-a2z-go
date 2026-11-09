# Step 16: Dynamic Programming

> Status: theory ready, problems coming ⏳

## What DP is
Recursion where the **same subproblems repeat**, so you **store answers and reuse them**.
Remember the Fibonacci tree in [Step 1.5](../step-01-basics/1.5-basic-recursion/): `fib(2)` computed again and again. DP computes each one once.

## The 3-step recipe
1. **Write the recursion.**
   - Express the problem with **indexes** (`f(i)`, `f(i, j)`, `f(i, target)`).
   - Do **all the choices** at that index.
   - Return the best (min/max) or the total (count).
   - Write the base cases.
2. **Memoisation (top-down):** add a `dp` table; before computing `f(i)`, check if `dp[i]` is already filled. Time drops from exponential to **(number of states) × (work per state)**.
3. **Tabulation (bottom-up):** fill the table with loops from the base cases upward. No recursion stack.
4. **(Bonus) Space optimisation:** if `dp[i]` only needs `dp[i−1]` (and `dp[i−2]`), keep just those rows/variables.

```go
// Climbing stairs: ways to reach step n taking 1 or 2 steps
// 1. recursion:  f(n) = f(n-1) + f(n-2), f(0) = f(1) = 1
// 2. memo:
func ways(n int, dp []int) int {
    if n <= 1 { return 1 }
    if dp[n] != -1 { return dp[n] }
    dp[n] = ways(n-1, dp) + ways(n-2, dp)
    return dp[n]
}
// 3. tabulation + 4. space optimised:
func waysOpt(n int) int {
    prev2, prev1 := 1, 1
    for i := 2; i <= n; i++ {
        prev2, prev1 = prev1, prev1+prev2
    }
    return prev1
}
```

## How to spot a DP problem
- "Count the number of ways…"
- "Minimum / maximum cost / length / profit…"
- Choices at each step, and greedy fails on a small example.

## The families (the sheet is organised by these)
| Family | State | Example |
|---|---|---|
| 1D | `f(i)` | climbing stairs, frog jump, house robber |
| Grids | `f(row, col)` | unique paths, min path sum, triangle |
| Subsequences / knapsack | `f(i, target)` | subset sum, 0/1 knapsack, coin change |
| Strings | `f(i, j)` over two strings | LCS, edit distance, wildcard matching |
| Stocks | `f(i, canBuy, [transactions])` | buy/sell I–VI |
| LIS | `f(i, prevIndex)` or binary search | longest increasing subsequence |
| Partition (MCM) | `f(i, j)` over a range, try every split `k` | matrix chain, burst balloons |
| Squares / rectangles | histogram per row | maximal rectangle, count squares |

## Problems

### 16.1 1D DP
| # | Problem | Status |
|---|---|---|
| 1 | Climbing stairs | ⏳ |
| 2 | Frog jump | ⏳ |
| 3 | Frog jump with K distance | ⏳ |
| 4 | Maximum sum of non-adjacent elements | ⏳ |
| 5 | House robber II | ⏳ |

### 16.2 2D / 3D DP and grids
| # | Problem | Status |
|---|---|---|
| 1 | Ninja's training | ⏳ |
| 2 | Grid unique paths | ⏳ |
| 3 | Unique paths II (obstacles) | ⏳ |
| 4 | Minimum path sum | ⏳ |
| 5 | Triangle | ⏳ |
| 6 | Minimum / maximum falling path sum | ⏳ |
| 7 | Cherry pickup II (3D DP) | ⏳ |

### 16.3 Subsequences
| # | Problem | Status |
|---|---|---|
| 1 | Subset sum equal to target | ⏳ |
| 2 | Partition equal subset sum | ⏳ |
| 3 | Partition with minimum absolute difference | ⏳ |
| 4 | Count subsets with sum K | ⏳ |
| 5 | Count partitions with given difference | ⏳ |
| 6 | 0/1 knapsack | ⏳ |
| 7 | Minimum coins | ⏳ |
| 8 | Target sum | ⏳ |
| 9 | Coin change II | ⏳ |
| 10 | Unbounded knapsack | ⏳ |
| 11 | Rod cutting | ⏳ |

### 16.4 Strings
| # | Problem | Status |
|---|---|---|
| 1 | Longest common subsequence | ⏳ |
| 2 | Print LCS | ⏳ |
| 3 | Longest common substring | ⏳ |
| 4 | Longest palindromic subsequence | ⏳ |
| 5 | Minimum insertions to make a palindrome | ⏳ |
| 6 | Minimum insertions/deletions to convert A to B | ⏳ |
| 7 | Shortest common supersequence | ⏳ |
| 8 | Distinct subsequences | ⏳ |
| 9 | Edit distance | ⏳ |
| 10 | Wildcard matching | ⏳ |

### 16.5 Stocks
| # | Problem | Status |
|---|---|---|
| 1 | Buy and sell stock I | ⏳ |
| 2 | Buy and sell stock II (unlimited) | ⏳ |
| 3 | Buy and sell stock III (at most 2) | ⏳ |
| 4 | Buy and sell stock IV (at most K) | ⏳ |
| 5 | With cooldown | ⏳ |
| 6 | With transaction fee | ⏳ |

### 16.6 Longest increasing subsequence
| # | Problem | Status |
|---|---|---|
| 1 | LIS | ⏳ |
| 2 | Print LIS | ⏳ |
| 3 | LIS with binary search | ⏳ |
| 4 | Largest divisible subset | ⏳ |
| 5 | Longest string chain | ⏳ |
| 6 | Longest bitonic subsequence | ⏳ |
| 7 | Number of LIS | ⏳ |

### 16.7 Partition DP (MCM)
| # | Problem | Status |
|---|---|---|
| 1 | Matrix chain multiplication | ⏳ |
| 2 | Minimum cost to cut a stick | ⏳ |
| 3 | Burst balloons | ⏳ |
| 4 | Evaluate boolean expression to true | ⏳ |
| 5 | Palindrome partitioning II | ⏳ |
| 6 | Partition array for maximum sum | ⏳ |

### 16.8 DP on squares
| # | Problem | Status |
|---|---|---|
| 1 | Maximal rectangle of 1s | ⏳ |
| 2 | Count square submatrices of 1s | ⏳ |
