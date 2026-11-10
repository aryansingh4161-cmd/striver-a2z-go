# Step 16: Dynamic Programming

> Status: theory, 16.1 – 16.5 done (39 problems); more coming ⏳

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Climbing stairs](16.1-1d-dp/01-climbing-stairs/) | f(n) = f(n−1) + f(n−2) |
| 2 | [Frog jump](16.1-1d-dp/02-frog-jump/) | min of the two previous stones |
| 3 | [Frog jump with K distance](16.1-1d-dp/03-frog-jump-k/) | min over the k previous stones |
| 4 | [Maximum sum of non-adjacent elements](16.1-1d-dp/04-max-sum-non-adjacent/) | pick (i−2) or skip (i−1) |
| 5 | [House robber II](16.1-1d-dp/05-house-robber-ii/) | skip first or skip last |

### 16.2 2D / 3D DP and grids
| # | Problem | Key idea |
|---|---|---|
| 1 | [Ninja's training](16.2-grids/01-ninja-training/) | state (day, last activity) |
| 2 | [Grid unique paths](16.2-grids/02-unique-paths/) | up + left; or C(m+n−2, m−1) |
| 3 | [Unique paths II (obstacles)](16.2-grids/03-unique-paths-ii/) | blocked cell = 0 ways |
| 4 | [Minimum path sum](16.2-grids/04-minimum-path-sum/) | cell + min(up, left) |
| 5 | [Triangle](16.2-grids/05-triangle/) | recurse from the top, fill from the bottom |
| 6 | [Minimum / maximum falling path sum](16.2-grids/06-falling-path-sum/) | 3 parents per cell, free start and end |
| 7 | [Cherry pickup II (3D DP)](16.2-grids/07-cherry-pickup-ii/) | state (row, colA, colB), 9 moves |

### 16.3 Subsequences
| # | Problem | Key idea |
|---|---|---|
| 1 | [Subset sum equal to target](16.3-subsequences/01-subset-sum/) | take / not take, boolean table |
| 2 | [Partition equal subset sum](16.3-subsequences/02-partition-equal-subset-sum/) | subset sum of total / 2 |
| 3 | [Partition with minimum absolute difference](16.3-subsequences/03-partition-min-difference/) | last row of subset sum table |
| 4 | [Count subsets with sum K](16.3-subsequences/04-count-subsets-sum-k/) | take + not take counts (zeros!) |
| 5 | [Count partitions with given difference](16.3-subsequences/05-count-partitions-difference/) | count subsets with (total − d) / 2 |
| 6 | [0/1 knapsack](16.3-subsequences/06-knapsack-01/) | f(i, W) = max(skip, val + f(i−1, W − wt)) |
| 7 | [Minimum coins](16.3-subsequences/07-minimum-coins/) | unbounded: stay on index after taking |
| 8 | [Target sum](16.3-subsequences/08-target-sum/) | same as count partitions with difference |
| 9 | [Coin change II](16.3-subsequences/09-coin-change-ii/) | count ways, unbounded |
| 10 | [Unbounded knapsack](16.3-subsequences/10-unbounded-knapsack/) | take and stay |
| 11 | [Rod cutting](16.3-subsequences/11-rod-cutting/) | unbounded knapsack on lengths |

### 16.4 Strings
| # | Problem | Key idea |
|---|---|---|
| 1 | [Longest common subsequence](16.4-strings/01-lcs/) | match → 1 + diag, else max(up, left) |
| 2 | [Print LCS](16.4-strings/02-print-lcs/) | walk the table back from (n, m) |
| 3 | [Longest common substring](16.4-strings/03-longest-common-substring/) | match → 1 + diag, else 0 |
| 4 | [Longest palindromic subsequence](16.4-strings/04-longest-palindromic-subsequence/) | LCS(s, reverse(s)) |
| 5 | [Minimum insertions to make a palindrome](16.4-strings/05-min-insertions-palindrome/) | n − LPS |
| 6 | [Minimum insertions/deletions to convert A to B](16.4-strings/06-min-insert-delete/) | n + m − 2·LCS |
| 7 | [Shortest common supersequence](16.4-strings/07-shortest-common-supersequence/) | n + m − LCS, build from the table |
| 8 | [Distinct subsequences](16.4-strings/08-distinct-subsequences/) | match → use + skip |
| 9 | [Edit distance](16.4-strings/09-edit-distance/) | insert / delete / replace |
| 10 | [Wildcard matching](16.4-strings/10-wildcard-matching/) | * = empty or eat one char |

### 16.5 Stocks
| # | Problem | Key idea |
|---|---|---|
| 1 | [Buy and sell stock I](16.5-stocks/01-stock-i/) | track the minimum so far |
| 2 | [Buy and sell stock II (unlimited)](16.5-stocks/02-stock-ii/) | state (i, canBuy) |
| 3 | [Buy and sell stock III (at most 2)](16.5-stocks/03-stock-iii/) | state (i, canBuy, transactions left) |
| 4 | [Buy and sell stock IV (at most K)](16.5-stocks/04-stock-iv/) | same with K |
| 5 | [With cooldown](16.5-stocks/05-stock-cooldown/) | after selling jump to i + 2 |
| 6 | [With transaction fee](16.5-stocks/06-stock-fee/) | pay the fee when selling |

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
