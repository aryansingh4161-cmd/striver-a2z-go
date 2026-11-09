# Bubble Sort

**Idea:** Compare neighbours and swap if they are in the wrong order. After each pass, the biggest remaining element has "bubbled" to the end.

## Steps
```
for end = n-1 down to 1:
    swapped = false
    for j = 0 .. end-1:
        if a[j] > a[j+1]: swap, swapped = true
    if not swapped: stop   // already sorted
```

## Dry run `[5 1 4 2]`
| pass | after the pass |
|---|---|
| 1 | 1 4 2 **5** |
| 2 | 1 2 **4 5** |
| 3 | no swaps → stop |

## Complexity
- **Worst / average:** O(n²)
- **Best:** O(n) on an already sorted array, thanks to the `swapped` check
- **Space:** O(1), **Stable:** yes
