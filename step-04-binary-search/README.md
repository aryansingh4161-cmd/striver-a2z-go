# Step 4: Binary Search

> Status: theory ready, problems coming in the next batch ⏳

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

### 4.1 Binary search on 1D arrays
| # | Problem | Status |
|---|---|---|
| 1 | Binary search to find X | ⏳ |
| 2 | Lower bound | ⏳ |
| 3 | Upper bound | ⏳ |
| 4 | Search insert position | ⏳ |
| 5 | Floor and ceil in sorted array | ⏳ |
| 6 | First and last occurrence | ⏳ |
| 7 | Count occurrences | ⏳ |
| 8 | Search in rotated sorted array I | ⏳ |
| 9 | Search in rotated sorted array II (duplicates) | ⏳ |
| 10 | Minimum in rotated sorted array | ⏳ |
| 11 | How many times is the array rotated | ⏳ |
| 12 | Single element in a sorted array | ⏳ |
| 13 | Find peak element | ⏳ |

### 4.2 Binary search on answers
| # | Problem | Status |
|---|---|---|
| 1 | Square root of a number | ⏳ |
| 2 | Nth root of a number | ⏳ |
| 3 | Koko eating bananas | ⏳ |
| 4 | Minimum days to make M bouquets | ⏳ |
| 5 | Smallest divisor given a threshold | ⏳ |
| 6 | Capacity to ship packages within D days | ⏳ |
| 7 | Kth missing positive number | ⏳ |
| 8 | Aggressive cows | ⏳ |
| 9 | Book allocation | ⏳ |
| 10 | Split array: largest sum | ⏳ |
| 11 | Painter's partition | ⏳ |
| 12 | Minimise max distance to gas station | ⏳ |
| 13 | Median of two sorted arrays | ⏳ |
| 14 | Kth element of two sorted arrays | ⏳ |

### 4.3 Binary search on 2D arrays
| # | Problem | Status |
|---|---|---|
| 1 | Row with maximum 1s | ⏳ |
| 2 | Search in a 2D matrix | ⏳ |
| 3 | Search in a row and column wise sorted matrix | ⏳ |
| 4 | Find peak element II | ⏳ |
| 5 | Matrix median | ⏳ |
