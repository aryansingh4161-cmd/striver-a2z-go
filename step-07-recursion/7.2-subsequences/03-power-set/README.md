# Power Set (All Subsequences)

**Problem (LeetCode 78):** Return all subsets of an array of distinct numbers.

`[1 2 3]` → `[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]` (8 = 2³)

A **subsequence** keeps the original order but may skip elements. For distinct elements, subsets and subsequences are the same thing.

## Approach 1: Pick / not pick recursion
At index `i`, either **take** `a[i]` or **skip** it. At the end, save the current path (a **copy**).
```
f(i):
    if i == n: save copy of path; return
    path.push(a[i]); f(i+1); path.pop()   // pick
    f(i+1)                                // not pick
```
- **Time:** O(2ⁿ × n), **Space:** O(n) depth

## Approach 2: Bitmasks
Each number `mask` from 0 to 2ⁿ−1 is one subset: bit `i` set → include `a[i]`.
- **Time:** O(2ⁿ × n), **Space:** O(1) besides the answer

## Recursion tree for [1 2]
```
              f(0) []
        pick 1 /       \ skip
       f(1) [1]        f(1) []
      /      \         /     \
  [1 2]     [1]      [2]     []
```

## Go gotcha
`res = append(res, path)` would save a **reference** that changes later. Always save `append([]int{}, path...)`.
