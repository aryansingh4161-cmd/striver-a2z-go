# Selection Sort

**Idea:** Find the minimum of the unsorted part and swap it to the front. Repeat.

## Steps
```
for i = 0 .. n-2:
    minIdx = i
    for j = i+1 .. n-1:
        if a[j] < a[minIdx]: minIdx = j
    swap a[i], a[minIdx]
```

## Dry run `[13 46 24 52 20 9]`
| i | minimum in a[i:] | after swap |
|---|---|---|
| 0 | 9 | **9** 46 24 52 20 13 |
| 1 | 13 | 9 **13** 24 52 20 46 |
| 2 | 20 | 9 13 **20** 52 24 46 |
| 3 | 24 | 9 13 20 **24** 52 46 |
| 4 | 46 | 9 13 20 24 **46** 52 |

## Complexity
- **Time:** O(n²) always: it scans the rest even if the array is already sorted.
- **Space:** O(1)
- **Swaps:** at most n−1 (fewest of the simple sorts).
- **Not stable:** a long-distance swap can jump over an equal element.
