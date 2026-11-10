# Step 12: Greedy Algorithms

> Status: ✅ theory and all 16 problems done

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
| # | Problem | Key idea |
|---|---|---|
| 1 | [Assign cookies](12.1-easy/01-assign-cookies/) | smallest cookie that fits the least greedy child |
| 2 | [Fractional knapsack](12.1-easy/02-fractional-knapsack/) | best value/weight first |
| 3 | [Minimum coins (Indian currency)](12.1-easy/03-minimum-coins/) | largest note first (fails for general coins) |
| 4 | [Lemonade change](12.1-easy/04-lemonade-change/) | give 10+5 before 5+5+5 |
| 5 | [Valid parenthesis string with *](12.1-easy/05-valid-parenthesis-string/) | track [lo, hi] open count |

### 12.2 Medium / hard
| # | Problem | Key idea |
|---|---|---|
| 1 | [N meetings in one room](12.2-medium-hard/01-n-meetings/) | earliest end first |
| 2 | [Jump game](12.2-medium-hard/02-jump-game/) | farthest reachable |
| 3 | [Jump game II](12.2-medium-hard/03-jump-game-ii/) | BFS levels of reach |
| 4 | [Minimum platforms](12.2-medium-hard/04-minimum-platforms/) | sort arrivals and departures, sweep |
| 5 | [Job sequencing](12.2-medium-hard/05-job-sequencing/) | highest profit, latest free slot |
| 6 | [Candy](12.2-medium-hard/06-candy/) | left pass + right pass |
| 7 | [Shortest job first](12.2-medium-hard/07-shortest-job-first/) | sort burst times |
| 8 | [LRU page faults](12.2-medium-hard/08-lru-page-faults/) | evict least recently used |
| 9 | [Insert interval](12.2-medium-hard/09-insert-interval/) | before / merge / after |
| 10 | [Merge intervals](12.2-medium-hard/10-merge-intervals/) | sort by start, extend |
| 11 | [Non-overlapping intervals](12.2-medium-hard/11-non-overlapping-intervals/) | n − max kept (sort by end) |
