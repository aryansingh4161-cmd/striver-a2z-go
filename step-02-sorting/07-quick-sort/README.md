# Quick Sort

**Idea (divide and conquer):**
1. Pick a **pivot** (here: the first element).
2. **Partition**: move everything ≤ pivot to its left and everything > pivot to its right. The pivot is now in its final sorted position.
3. Recursively sort the left part and the right part.

Unlike merge sort, the hard work happens **before** the recursive calls, and there is no merge step.

## Partition (Hoare-style, as taught by Striver)
```
pivot = a[lo]; i = lo; j = hi
while i < j:
    while a[i] <= pivot and i <= hi-1: i++   // find first element > pivot
    while a[j] > pivot and j >= lo+1:  j--   // find first element <= pivot from the right
    if i < j: swap a[i], a[j]
swap a[lo], a[j]   // put pivot in place
return j
```

## Dry run `[4 6 2 5 7 9 1 3]`, pivot = 4
| step | i stops at | j stops at | action |
|---|---|---|---|
| 1 | 6 (idx 1) | 3 (idx 7) | swap → 4 3 2 5 7 9 1 6 |
| 2 | 5 (idx 3) | 1 (idx 6) | swap → 4 3 2 1 7 9 5 6 |
| 3 | 7 (idx 4) | 1 (idx 3) | i > j, stop |
| | | | swap pivot with a[3] → **1 3 2 4** 7 9 5 6 |

4 is now in place. Recurse on `[1 3 2]` and `[7 9 5 6]`.

## Complexity
- **Average:** O(n log n)
- **Worst:** O(n²): when the pivot is always the smallest/largest (e.g. sorted input with first-element pivot). A **random pivot** makes this very unlikely.
- **Space:** O(log n) stack on average, no extra array
- **Not stable**
