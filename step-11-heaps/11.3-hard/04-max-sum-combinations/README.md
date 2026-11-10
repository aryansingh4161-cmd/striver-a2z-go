# Maximum Sum Combinations

**Problem:** Arrays `a` and `b` of size n. Pick one element from each; return the `k` largest sums `a[i] + b[j]` (over all n² pairs), in decreasing order.

`a = [3 2]`, `b = [1 4]`, k = 2 → `[7 6]`

## Approach 1: All pairs
Compute n² sums, sort, take k. **O(n² log n)**.

## Approach 2: Sort + max-heap + visited set (optimal)
1. Sort both arrays **descending**. The biggest sum is `a[0] + b[0]`.
2. Put `(0, 0)` in a max-heap (by sum) and in a `visited` set.
3. Repeat k times: pop the best `(i, j)`, record its sum, then push its two "next best" neighbours `(i+1, j)` and `(i, j+1)` if not visited.

Like walking a sorted 2D grid from the top-left corner, always expanding the best frontier cell.
- **Time:** O(n log n + k log k), **Space:** O(k)
