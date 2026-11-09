# Maximum Subarray Sum (Kadane's Algorithm) + Print the Subarray

**Problem (LeetCode 53):** Return the largest sum of a non-empty contiguous subarray. Follow-up: also return which subarray.

`[-2 1 -3 4 -1 2 1 -5 4]` → 6 (`[4 -1 2 1]`)

## Approach 1: All subarrays (brute / better)
Fix start `i`, extend `j` with a running sum. **Time O(n²)**.

## Approach 2: Kadane's algorithm (optimal)
Walk left to right with a running `sum`:
- add the current element,
- update `best`,
- **if `sum` drops below 0, reset it to 0.** A negative prefix can only make any future subarray smaller, so drop it.

```
best = -∞, sum = 0
for x in a:
    sum += x
    best = max(best, sum)
    if sum < 0: sum = 0
```
- **Time:** O(n), **Space:** O(1)

Update `best` **before** resetting, so an all-negative array returns its largest element (e.g. `[-3 -1 -2]` → −1).

## Printing the subarray
Remember where the current sum started: when `sum == 0` (just reset / beginning), the next element is a possible `start`. When `best` improves, save `(start, i)`.

## Dry run `[-2 1 -3 4 -1 2 1 -5 4]`
| x | sum | best | note |
|---|---|---|---|
| −2 | −2 → 0 | −2 | reset |
| 1 | 1 | 1 | |
| −3 | −2 → 0 | 1 | reset |
| 4 | 4 | 4 | start = 3 |
| −1 | 3 | 4 | |
| 2 | 5 | 5 | |
| 1 | 6 | **6** | end = 6 |
| −5 | 1 | 6 | |
| 4 | 5 | 6 | |
