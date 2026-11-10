# Subsets II (input has duplicates)

**Problem (LeetCode 90):** Return all **unique** subsets.

`[1 2 2]` → `[] [1] [1 2] [1 2 2] [2] [2 2]`

## Approach 1: Power set + set of sorted subsets
Generate all 2ⁿ, dedupe with a set. Wasteful.

## Approach 2: Sort + loop with skip (optimal)
Same idea as [Combination Sum II](../07-combination-sum-ii/), but **every node of the recursion is an answer** (not only when the sum hits a target):
```
sort(a)
f(start):
    save path                       // every prefix is a subset
    for j = start .. n-1:
        if j > start and a[j] == a[j-1]: continue
        pick a[j]; f(j+1); unpick
```
- **Time:** O(2ⁿ × n), **Space:** O(n) depth

## Recursion tree for [1 2 2]
```
[]
├─ [1]
│  └─ [1 2]
│     └─ [1 2 2]
│  (second 2 skipped at this level)
└─ [2]
   └─ [2 2]
(second 2 skipped at the top level)
```
