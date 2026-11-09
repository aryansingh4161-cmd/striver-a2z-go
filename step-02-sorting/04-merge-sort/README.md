# Merge Sort

**Idea (divide and conquer):**
1. Split the array into two halves.
2. Sort each half (recursively).
3. **Merge** the two sorted halves into one sorted array.

A single element is already sorted, which is the base case.

## Merge step (the heart of it)
Two pointers, one on each sorted half. Take the smaller one each time into a temporary slice. When one half runs out, copy the rest of the other. Copy the temp back.

```
mergeSort(lo, hi):
    if lo >= hi: return
    mid = (lo + hi) / 2
    mergeSort(lo, mid)
    mergeSort(mid+1, hi)
    merge(lo, mid, hi)
```

## Dry run `[3 1 2 4 1 5 2 6 4]`
```
split:  [3 1 2 4 1] [5 2 6 4]
        [3 1 2] [4 1]   [5 2] [6 4]
        [3 1] [2] [4] [1] [5] [2] [6] [4]
merge:  [1 3] [2] → [1 2 3]   [1 4]   [2 5]   [4 6]
        [1 1 2 3 4]           [2 4 5 6]
        [1 1 2 2 3 4 4 5 6]
```

## Complexity
- **Time:** O(n log n) in every case. There are log n levels and each level merges n elements.
- **Space:** O(n) for the temp slice + O(log n) recursion stack.
- **Stable:** yes (take from the left half when equal: use `<=`).

## Why it matters later
The merge step lets you count things across the two halves, e.g. **Count Inversions** and **Reverse Pairs** (Step 3 hard).
