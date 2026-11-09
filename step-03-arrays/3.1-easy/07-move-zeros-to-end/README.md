# Move Zeros to End

**Problem (LeetCode 283):** Move all zeros to the end, keeping the order of the non-zero elements. Do it in place.

`[1 0 2 3 0 4 0 1]` → `[1 2 3 4 1 0 0 0]`

## Approach 1: Temp array (brute)
Copy non-zeros into `temp`, write them back, fill the rest with 0.
**Time O(n), Space O(n)**.

## Approach 2: Two pointers (optimal)
`j` = where the next non-zero should go. Scan with `i`; whenever `a[i] != 0`, swap it to position `j` and move `j` on.

```
j = 0
for i = 0 .. n-1:
    if a[i] != 0:
        swap a[i], a[j]
        j++
```
- **Time:** O(n), **Space:** O(1)

## Dry run `[1 0 2 0 3]`
| i | a[i] | action | array | j |
|---|---|---|---|---|
| 0 | 1 | swap(0,0) | 1 0 2 0 3 | 1 |
| 1 | 0 | skip | | 1 |
| 2 | 2 | swap(2,1) | 1 2 0 0 3 | 2 |
| 3 | 0 | skip | | 2 |
| 4 | 3 | swap(4,2) | 1 2 3 0 0 | 3 |
