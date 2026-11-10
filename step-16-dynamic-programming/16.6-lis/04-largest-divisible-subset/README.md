# Largest Divisible Subset

**Problem:** Distinct positive integers. Return the largest subset where for **every pair** `(a, b)`, `a % b == 0` or `b % a == 0`.

```
[1 2 3]     → [1 2] (or [1 3])
[1 2 4 8]   → [1 2 4 8]
```

## Key observation: sort first, then it's LIS
After sorting, a valid subset looks like a **chain** `a1 | a2 | a3 | ...` where each divides the next. Divisibility is **transitive** (if `a | b` and `b | c` then `a | c`), so checking **neighbours** in the chain is enough.

That's [Print LIS](../02-print-lis/) with the condition changed:
```
"nums[j] < nums[i]"   →   "nums[i] % nums[j] == 0"
```
(Sorting is what makes "subset" behave like "subsequence".)

## Dry run (sorted `[1 2 4 7 8]`)
| i | x | dp | parent |
|---|---|---|---|
| 0 | 1 | 1 | – |
| 1 | 2 | 2 | 1 |
| 2 | 4 | 3 | 2 |
| 3 | 7 | 2 | 1 |
| 4 | 8 | 4 | 4 |
Back from 8: 8 → 4 → 2 → 1 → **[1 2 4 8]** ✅

- **Time:** O(n²) (+ sort), **Space:** O(n)
