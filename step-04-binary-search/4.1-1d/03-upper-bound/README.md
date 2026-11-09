# Upper Bound

**Problem:** In a sorted array, return the **first index `i` with `a[i] > x`** (strictly greater). If none, return `n`.

`[2 3 6 7 8 8 11 11 11 12]`, x = 6 → 3 (7)
x = 11 → 9, x = 12 → 10

## Approach: Binary search (optimal)
Exactly like [lower bound](../02-lower-bound/), but the condition is `a[mid] > x`.
- **Time:** O(log n), **Space:** O(1)

## Lower vs upper bound
| | condition | for x = 8 in `[2 3 6 7 8 8 11]` |
|---|---|---|
| lower bound | first `≥ x` | 4 |
| upper bound | first `> x` | 6 |

`upper − lower` = number of times `x` appears (here 2). See [count occurrences](../07-count-occurrences/).
