# Recursive Insertion Sort

**Idea:** To sort the first `i+1` elements: insert `a[i]` into the already sorted `a[0..i-1]`, then do the same for `i+1`.

```
insert(i):
    if i == n: return
    j = i
    while j > 0 and a[j-1] > a[j]: swap, j--
    insert(i + 1)
```

- **Time:** O(n²), O(n) best
- **Space:** O(n) recursion stack

Same as [insertion sort](../03-insertion-sort/) with the outer loop replaced by recursion.
