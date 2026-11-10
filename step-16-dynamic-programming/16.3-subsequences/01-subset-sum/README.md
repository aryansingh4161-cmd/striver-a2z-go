# Subset Sum Equal to Target

**Problem:** Given positive integers `arr` and a `target`, is there a **subset** whose sum is exactly `target`?

```
arr = [1 2 3 4], target = 4 → true ({4} or {1, 3})
arr = [4 3 2 1], target = 11 → false (total is only 10)
```

## The pattern: take / not take
This is the base of every "subsequence DP". Look at elements from the last index down. At index `i`, element `arr[i]` is either **in** the subset or **not**.
```
f(i, t) = "can some subset of arr[0..i] sum to t?"
f(i, 0) = true                      (empty subset)
f(0, t) = (arr[0] == t)
f(i, t) = f(i−1, t)                 // not take
       || (arr[i] <= t && f(i−1, t − arr[i]))   // take
```

## Stages
| | Time | Space |
|---|---|---|
| Recursion | O(2ⁿ) | O(n) stack |
| Memo | O(n · target) | O(n · target) + stack |
| Tabulation `dp[i][t]` | O(n · target) | O(n · target) |
| One row | O(n · target) | **O(target)** |

## Tabulation details
- `dp[i][0] = true` for every `i`.
- `dp[0][arr[0]] = true` (only if `arr[0] ≤ target`).
- For `i = 1..n−1`, `t = 1..target`: `dp[i][t] = dp[i−1][t] || (arr[i] ≤ t && dp[i−1][t−arr[i]])`.

## One row trick
Row `i` only reads row `i−1`. Keep `prev` and build `cur`, then swap. (Or use **one** array and loop `t` from **high to low**, so `dp[t − arr[i]]` is still the old value.)

## Dry run (`[1 2 3]`, target 5)
| i \ t | 0 | 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|---|---|
| 0 (1) | T | T | F | F | F | F |
| 1 (2) | T | T | T | T | F | F |
| 2 (3) | T | T | T | T | T | **T** |
`{2, 3}` → true ✅
