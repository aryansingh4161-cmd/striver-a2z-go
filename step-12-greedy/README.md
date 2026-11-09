# Step 12: Greedy Algorithms

> Status: theory ready, problems coming ⏳

## The idea
At each step, take the choice that looks **best right now**, and never undo it.
Greedy is fast and simple, but it only works when a local best choice is guaranteed to lead to a global best answer.

## How to know if greedy works
1. Find a rule ("always take the meeting that ends earliest").
2. Try to break it with small examples. If you find one, greedy fails → think DP.
3. Classic failure: coin change with coins `{1, 3, 4}` and amount 6. Greedy takes 4+1+1 (3 coins), but 3+3 (2 coins) is better. Works for Indian currency denominations, fails in general.

## Common greedy rules
| Problem type | Rule |
|---|---|
| Interval scheduling (max meetings) | sort by **end** time, pick if it starts after the last picked ends |
| Remove fewest intervals to avoid overlap | same as above; answer = n − picked |
| Minimum platforms | sort arrivals and departures separately, sweep with two pointers |
| Fractional knapsack | sort by value/weight, take the best ratio first (can split items) |
| Job sequencing | sort by profit, place each job in the latest free slot before its deadline |
| Jump game | track the farthest index reachable |
| Assign cookies | sort both, match the smallest cookie that satisfies each child |

Most greedy solutions start with a **sort**: O(n log n).

## Problems

### 12.1 Easy
| # | Problem | Status |
|---|---|---|
| 1 | Assign cookies | ⏳ |
| 2 | Fractional knapsack | ⏳ |
| 3 | Minimum coins (Indian currency) | ⏳ |
| 4 | Lemonade change | ⏳ |
| 5 | Valid parenthesis string (with `*`) | ⏳ |

### 12.2 Medium / hard
| # | Problem | Status |
|---|---|---|
| 1 | N meetings in one room | ⏳ |
| 2 | Jump game | ⏳ |
| 3 | Jump game II | ⏳ |
| 4 | Minimum number of platforms | ⏳ |
| 5 | Job sequencing | ⏳ |
| 6 | Candy | ⏳ |
| 7 | Shortest job first (SJF) scheduling | ⏳ |
| 8 | LRU page faults | ⏳ |
| 9 | Insert interval | ⏳ |
| 10 | Merge intervals | ⏳ |
| 11 | Non-overlapping intervals | ⏳ |
