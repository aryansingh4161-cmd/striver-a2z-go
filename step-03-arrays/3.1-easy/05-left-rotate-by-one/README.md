# Left Rotate Array by One

**Problem:** Shift every element one place left; the first element moves to the end.

`[1 2 3 4 5]` → `[2 3 4 5 1]`

## Approach (optimal)
```
temp = a[0]
for i = 1 .. n-1: a[i-1] = a[i]
a[n-1] = temp
```
- **Time:** O(n), **Space:** O(1)

Save `a[0]` **before** shifting, or it gets overwritten.
