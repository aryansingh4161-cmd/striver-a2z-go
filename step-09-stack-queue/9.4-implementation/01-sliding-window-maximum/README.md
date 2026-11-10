# Sliding Window Maximum

**Problem (LeetCode 239):** Return the maximum of every window of size `k`.

`[1 3 −1 −3 5 3 6 7]`, k = 3 → `[3 3 5 5 6 7]`

## Approach 1: Scan each window
**O(n × k)**.

## Approach 2: Monotonic deque (optimal)
Keep a **deque of indexes** whose values are **decreasing** from front to back. The front is always the window's maximum.
For each `i`:
1. Remove the front if it has **left the window** (`front <= i − k`).
2. Remove from the **back** every index with value `<= a[i]`: they can never be a maximum again while `a[i]` is in the window (it's bigger and leaves later).
3. Push `i` at the back.
4. Once `i >= k − 1`, the front is the answer for this window.

- **Time:** O(n) (each index enters and leaves the deque once), **Space:** O(k)

## Dry run (first few steps)
| i | a[i] | deque (values) | output |
|---|---|---|---|
| 0 | 1 | [1] | |
| 1 | 3 | [3] (1 removed) | |
| 2 | −1 | [3 −1] | 3 |
| 3 | −3 | [3 −1 −3] | 3 |
| 4 | 5 | [5] (3 left window; −1, −3 removed) | 5 |
