# Longest Subarray with Sum K (positives and negatives)

**Problem:** Elements can be negative. Return the length of the longest subarray with sum exactly `k`.

`[1 -1 5 -2 3]`, k = 3 → 4 (`[1 -1 5 -2]`)

## Key idea: prefix sums
`prefix[i]` = sum of `a[0..i]`.
Sum of subarray `a[j+1..i]` = `prefix[i] − prefix[j]`.
We want that to be `k` → we need an earlier prefix equal to **`prefix[i] − k`**.

## Approach (optimal for mixed signs)
Store each prefix sum's **first** index in a map. For each `i`:
1. `sum += a[i]`
2. If `sum == k` → the whole `a[0..i]` works: length `i + 1`.
3. If `sum − k` is in the map at index `j` → length `i − j`.
4. Put `sum` in the map **only if it's not already there** (keep the earliest index → longest subarray).

- **Time:** O(n) average, **Space:** O(n)

## Dry run `[1 -1 5 -2 3]`, k = 3
| i | a[i] | sum | sum−k | found at | length | map adds |
|---|---|---|---|---|---|---|
| 0 | 1 | 1 | −2 | — | | 1→0 |
| 1 | −1 | 0 | −3 | — | | 0→1 |
| 2 | 5 | 5 | 2 | — | | 5→2 |
| 3 | −2 | 3 | 0 | sum==k | **4** | 3→3 |
| 4 | 3 | 6 | 3 | 3 | 1 | 6→4 |

## Why "only first index" matters (zeros!)
`[2 0 0 3]`, k = 3: prefix 2 appears at index 0, 1 and 2. Keeping index 0 gives the longest answer `[0 0 3]` = 3.
