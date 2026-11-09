# Step 4: Binary Search

> Status: ✅ theory and all 32 problems done

## The idea
In a **sorted** space, look at the middle. Its value tells you which half **cannot** contain the answer, so throw that half away. Each step halves the search space → **O(log n)**.
n = 10⁹ needs only ~30 steps.

## Template (find target in a sorted slice)
```go
func search(a []int, target int) int {
    lo, hi := 0, len(a)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2 // same as (lo+hi)/2 but can never overflow
        switch {
        case a[mid] == target:
            return mid
        case a[mid] < target:
            lo = mid + 1      // answer is on the right
        default:
            hi = mid - 1      // answer is on the left
        }
    }
    return -1
}
```

## Lower bound / upper bound
- **Lower bound:** first index with `a[i] >= x`.
- **Upper bound:** first index with `a[i] > x`.

Pattern: keep a possible answer `ans = n`, and when `a[mid]` satisfies the condition, record `ans = mid` and keep searching **left** (`hi = mid − 1`) for an earlier one.

Go's standard library: `sort.SearchInts(a, x)` = lower bound; `sort.Search(n, func(i int) bool {...})` finds the first index where the condition becomes true.

## Three families of problems
1. **On a sorted array (1D):** find, first/last occurrence, rotated arrays, peak.
   - Rotated sorted array: at every `mid`, **one half is always sorted**. Check if the target lies in the sorted half.
2. **On the answer ("binary search on answers"):** the array isn't sorted, but the **answer range** is monotonic.
   - Ask: "If the answer were `x`, is it possible?" If possible for `x`, it's also possible for every bigger `x` (or smaller). Then binary search on `x`.
   - Examples: Koko eating bananas, ship packages in D days, aggressive cows, book allocation.
   - Recipe: find `low` and `high` of the possible answers → write `possible(x)` → binary search for the smallest (or largest) `x` that works.
3. **On 2D matrices:** treat a fully sorted matrix as a 1D array (`row = idx / cols`, `col = idx % cols`), or walk from a corner.

## Problems

The [4.2 README](4.2-answers/) explains the "binary search on answers" recipe in detail.

### 4.1 Binary search on 1D arrays
| # | Problem | Key idea |
|---|---|---|
| 1 | [Binary search to find X](4.1-1d/01-binary-search/) | halve the range |
| 2 | [Lower bound](4.1-1d/02-lower-bound/) | first ≥ x: save, go left |
| 3 | [Upper bound](4.1-1d/03-upper-bound/) | first > x |
| 4 | [Search insert position](4.1-1d/04-search-insert-position/) | = lower bound |
| 5 | [Floor and ceil](4.1-1d/05-floor-and-ceil/) | save candidate, go right / left |
| 6 | [First and last occurrence](4.1-1d/06-first-last-occurrence/) | on match keep searching left / right |
| 7 | [Count occurrences](4.1-1d/07-count-occurrences/) | upper − lower |
| 8 | [Search in rotated sorted array I](4.1-1d/08-search-rotated-sorted-array/) | one half is always sorted |
| 9 | [Search in rotated sorted array II](4.1-1d/09-search-rotated-sorted-array-ii/) | shrink ends when lo = mid = hi |
| 10 | [Minimum in rotated sorted array](4.1-1d/10-min-in-rotated-sorted-array/) | min of the sorted half |
| 11 | [How many times rotated](4.1-1d/11-times-array-rotated/) | index of the minimum |
| 12 | [Single element in sorted array](4.1-1d/12-single-element-sorted-array/) | pair parity (even, odd) |
| 13 | [Find peak element](4.1-1d/13-find-peak-element/) | walk uphill |

### 4.2 Binary search on answers
| # | Problem | Key idea |
|---|---|---|
| 1 | [Square root](4.2-answers/01-square-root/) | last x with x² ≤ n |
| 2 | [Nth root](4.2-answers/02-nth-root/) | early-stop power |
| 3 | [Koko eating bananas](4.2-answers/03-koko-eating-bananas/) | min speed, sum of ceilings |
| 4 | [Minimum days for M bouquets](4.2-answers/04-min-days-bouquets/) | count adjacent bloomed |
| 5 | [Smallest divisor](4.2-answers/05-smallest-divisor/) | sum of ceilings ≤ threshold |
| 6 | [Ship packages within D days](4.2-answers/06-ship-packages/) | greedy days count |
| 7 | [Kth missing positive](4.2-answers/07-kth-missing-positive/) | missing = a[i] − (i+1) |
| 8 | [Aggressive cows](4.2-answers/08-aggressive-cows/) | max of min, greedy placement |
| 9 | [Book allocation](4.2-answers/09-book-allocation/) | min of max, greedy split |
| 10 | [Split array largest sum](4.2-answers/10-split-array-largest-sum/) | same as book allocation |
| 11 | [Painter's partition](4.2-answers/11-painters-partition/) | same as book allocation |
| 12 | [Minimise max distance to gas station](4.2-answers/12-minimise-max-distance-gas-station/) | binary search on reals |
| 13 | [Median of two sorted arrays](4.2-answers/13-median-two-sorted-arrays/) | partition the smaller array |
| 14 | [Kth element of two sorted arrays](4.2-answers/14-kth-element-two-sorted-arrays/) | partition with k on the left |

### 4.3 Binary search on 2D arrays
| # | Problem | Key idea |
|---|---|---|
| 1 | [Row with maximum 1s](4.3-2d/01-row-with-max-ones/) | lower bound per row / staircase |
| 2 | [Search a 2D matrix](4.3-2d/02-search-2d-matrix/) | flatten: idx/m, idx%m |
| 3 | [Search row & column sorted matrix](4.3-2d/03-search-2d-matrix-ii/) | start top-right |
| 4 | [Find peak element II](4.3-2d/04-find-peak-element-ii/) | binary search on column maxima |
| 5 | [Matrix median](4.3-2d/05-matrix-median/) | binary search on value, count ≤ x |
