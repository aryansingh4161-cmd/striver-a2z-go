# Recursive Bubble Sort

**Idea:** One bubble pass puts the largest element at the end. Then the problem is the same, just one element smaller.

```
bubble(n):
    if n == 1: return
    swapped = false
    for j = 0 .. n-2:
        if a[j] > a[j+1]: swap, swapped = true
    if not swapped: return
    bubble(n - 1)
```

- **Time:** O(n²), O(n) best
- **Space:** O(n) recursion stack (the loop version uses O(1))

This is the same as [bubble sort](../02-bubble-sort/) with the outer loop replaced by recursion. Good recursion practice.
