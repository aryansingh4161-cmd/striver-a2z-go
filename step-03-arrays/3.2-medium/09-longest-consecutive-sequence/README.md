# Longest Consecutive Sequence

**Problem (LeetCode 128):** Unsorted array. Return the length of the longest run of consecutive integers (order in the array doesn't matter).

`[100 4 200 1 3 2]` → 4 (`1 2 3 4`)

## Approach 1: Linear search for the next (brute)
For each `x`, keep searching for `x+1`, `x+2`, … with linear search. **Time O(n³)** worst case.

## Approach 2: Sort (better)
Sort, then walk: if `a[i] == last + 1` → `count++`; if `a[i] == last` → skip duplicate; else → `count = 1`.
**Time O(n log n), Space O(1)** (if you may sort the input).

## Approach 3: Hash set (optimal)
Put everything in a set. A number `x` **starts** a sequence only if `x − 1` is **not** in the set. Only from those starts, count upward `x+1, x+2, …`.
- **Time:** O(n) average: each number is visited by at most one counting walk.
- **Space:** O(n)

## Dry run `[100 4 200 1 3 2]`
| x | x−1 in set? | start? | length |
|---|---|---|---|
| 100 | 99 no | yes | 1 |
| 4 | 3 yes | no | |
| 200 | no | yes | 1 |
| 1 | 0 no | yes | 1,2,3,4 → **4** |
| 3, 2 | yes | no | |
