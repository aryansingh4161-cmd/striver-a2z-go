# Maximum Sum of Non-Adjacent Elements (House Robber)

**Problem:** Pick some elements of `nums` so that **no two picked elements are next to each other**. Maximise their sum. (Same as LeetCode "House Robber": you can't rob two neighbouring houses.)

```
[2 1 4 9] → 11 (2 + 9)
[1 2 4]   → 5  (1 + 4)
```

## Pick / not pick
At index `i` you have exactly two choices:
- **Pick** `nums[i]`: then you can't take `i−1`, so add `f(i−2)`.
- **Don't pick**: the best up to `i−1`, i.e. `f(i−1)`.

```
f(i) = max( nums[i] + f(i−2),  f(i−1) )
f(0) = nums[0]
f(negative) = 0
```
This "pick / not pick" pattern is the core of the subsequence problems in [16.3](../../16.3-subsequences/).

## Stages
| | Time | Space |
|---|---|---|
| Recursion | O(2ⁿ) | O(n) stack |
| Memo | O(n) | O(n) + stack |
| Tabulation | O(n) | O(n) |
| Two variables | O(n) | **O(1)** |

## Dry run (`[2 1 4 9]`)
| i | pick | not pick | dp |
|---|---|---|---|
| 0 | 2 | 0 | 2 |
| 1 | 1 + 0 | 2 | 2 |
| 2 | 4 + 2 | 2 | 6 |
| 3 | 9 + 2 | 6 | **11** |

## Edge cases
- One element → that element.
- All elements are non-negative in the standard problem, so picking nothing is never better than picking something.
