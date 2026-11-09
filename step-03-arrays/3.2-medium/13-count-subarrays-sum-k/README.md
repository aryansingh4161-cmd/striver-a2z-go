# Count Subarrays with Sum K

**Problem (LeetCode 560):** Count the subarrays whose sum is exactly `k`. Elements may be negative.

`[3 1 2 4]`, k = 6 → 2 (`[3 1 2]` and `[2 4]`)

## Approach 1: All subarrays (brute)
Fix start, extend end with a running sum. **Time O(n²)**.

## Approach 2: Prefix sum + hash map (optimal)
Same idea as [longest subarray with sum K](../../3.1-easy/14-longest-subarray-sum-k/), but now we **count**.
A subarray ending at `i` has sum `k` when an earlier prefix equals `sum − k`. How many such subarrays? As many times as `sum − k` appeared before.

```
count = map{0: 1}   // empty prefix: lets subarrays starting at index 0 count
sum = 0, ans = 0
for x in a:
    sum += x
    ans += count[sum - k]
    count[sum]++
```
- **Time:** O(n), **Space:** O(n)

## Dry run `[3 1 2 4]`, k = 6
| x | sum | sum−k | count[sum−k] | ans | map after |
|---|---|---|---|---|---|
| | | | | | {0:1} |
| 3 | 3 | −3 | 0 | 0 | {0:1, 3:1} |
| 1 | 4 | −2 | 0 | 0 | +4 |
| 2 | 6 | 0 | 1 | 1 | +6 |
| 4 | 10 | 4 | 1 | **2** | +10 |

Why `{0: 1}` at the start? `[3 1 2]` starts at index 0: we need "prefix before index 0" = 0.
