# LIS in O(n log n) with Binary Search

**Problem:** Same [LIS length](../01-lis/), but for big `n` (like 10⁵), O(n²) is too slow.

## The `tails` idea
Keep an array `tails` where `tails[k]` = the **smallest possible last element** of an increasing subsequence of length `k + 1` seen so far.

`tails` is always **sorted** (a longer increasing subsequence must end with a bigger value). For each `x`:
- If `x` is bigger than everything in `tails`, it **extends** the longest one → append.
- Otherwise, find the first `tails[k] >= x` and **replace** it with `x`. Same length, smaller ending → better for the future.

The answer is `len(tails)`.

```
for x in nums:
    k = lowerBound(tails, x)    // first index with tails[k] >= x
    if k == len(tails): append x
    else: tails[k] = x
```

## Careful: `tails` is not the LIS itself
After replacements, `tails` can be a mix of different subsequences. Only its **length** is guaranteed right. (To print, use [Print LIS](../02-print-lis/).)

## Why lower bound (`>=`)?
"Strictly increasing" means equal values can't both be used. Replacing the equal value (instead of appending) makes sure duplicates don't extend the length. For "non-decreasing" you'd use upper bound (`>`).

## Dry run (`[1 7 8 4 5 6 −1 9]`)
| x | tails |
|---|---|
| 1 | [1] |
| 7 | [1 7] |
| 8 | [1 7 8] |
| 4 | [1 4 8] |
| 5 | [1 4 5] |
| 6 | [1 4 5 6] |
| −1 | [−1 4 5 6] |
| 9 | [−1 4 5 6 9] → **5** ✅ |

- **Time:** O(n log n), **Space:** O(n)
