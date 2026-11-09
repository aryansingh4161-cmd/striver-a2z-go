# Union (and Intersection) of Two Sorted Arrays

**Union:** all **distinct** values from both arrays, sorted.
`a = [1 1 2 3 4 5]`, `b = [2 3 4 4 5 6]` → `[1 2 3 4 5 6]`

**Intersection:** values present in both (here counted with repeats, like LeetCode 350 on sorted input).
`a = [1 2 2 3 3 4 5 6]`, `b = [2 3 3 5 6 6 7]` → `[2 3 3 5 6]`

## Union, approach 1: Set (brute)
Insert everything into a set, then sort. **Time O((n+m) log(n+m)), Space O(n+m)**.

## Union, approach 2: Two pointers (optimal)
Like the merge step of merge sort. Take the smaller value; add it only if it differs from the **last value added** (that skips duplicates).
```
i = j = 0
while i < n and j < m:
    if a[i] <= b[j]: add(a[i]); i++
    else:            add(b[j]); j++
add the leftovers of a, then of b
add(x): if result is empty or last != x: append x
```
- **Time:** O(n + m), **Space:** O(n + m) for the answer only

## Intersection, two pointers
```
if a[i] < b[j]: i++
else if a[i] > b[j]: j++
else: add a[i]; i++; j++
```
- **Time:** O(n + m)
