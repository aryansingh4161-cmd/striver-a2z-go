# Check if Any Subsequence Has Sum K

**Problem:** Return true if some subsequence sums to `k` (and show one).

`[1 2 3]`, k = 5 → true (`[2 3]`)

## Approach: Pick / not pick that **returns a bool** and stops early
```
f(i, sum):
    if i == n: return sum == k
    pick a[i]; if f(i+1, sum + a[i]): return true
    unpick
    return f(i+1, sum)
```
The moment one branch returns true, every caller returns true immediately, so we stop at the **first** answer instead of exploring all 2ⁿ paths.

The three variants of this template:
| Want | Return type | Combine |
|---|---|---|
| all answers | nothing (save at leaf) | explore everything |
| any one answer | `bool` | `if f(...) { return true }` |
| count | `int` | `return f(pick) + f(notPick)` |

- **Time:** O(2ⁿ) worst case, **Space:** O(n)
