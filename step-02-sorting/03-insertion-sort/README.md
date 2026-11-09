# Insertion Sort

**Idea:** Like sorting playing cards in your hand. The left part is always sorted; take the next element and slide it left until it is in the right place.

## Steps
```
for i = 1 .. n-1:
    j = i
    while j > 0 and a[j-1] > a[j]:
        swap a[j-1], a[j]
        j--
```

## Dry run `[14 9 15 12 6]`
| i | element | after inserting |
|---|---|---|
| 1 | 9 | **9 14** 15 12 6 |
| 2 | 15 | **9 14 15** 12 6 |
| 3 | 12 | **9 12 14 15** 6 |
| 4 | 6 | **6 9 12 14 15** |

## Complexity
- **Worst / average:** O(n²)
- **Best:** O(n) when already sorted (inner loop never runs)
- **Space:** O(1), **Stable:** yes
- Great for small or **nearly sorted** arrays. Real sorting libraries use it for small pieces.
