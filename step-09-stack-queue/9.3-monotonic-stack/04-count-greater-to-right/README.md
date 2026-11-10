# Number of Greater Elements to the Right

**Problem:** For each query index `i`, count how many elements to the right of `i` are greater than `a[i]` (all of them, not just the next one).

`[3 4 2 7 5 8 10 6]`, queries `[0 5]` → `[6 1]`

## Approach 1: Scan per query (what Striver shows)
For each query, scan to the right and count. **O(n) per query, O(n·q) total**.

This one is in the stack section because it looks like "next greater", but a monotonic stack only finds the **nearest** one, not the count of all of them.

## Approach 2: For many queries, count all at once
Count "greater elements to the right" for **every** index, then answer each query in O(1). This is the [count inversions](../../../step-03-arrays/3.3-hard/10-count-inversions/) idea from the other direction: with merge sort, when an element of the left half is smaller than an element of the right half, it gains one "greater to the right" for every remaining right-half element.
- **Time:** O(n log n + q), **Space:** O(n)

Both versions are in `solution.go`.
