# Binary Subarrays With Sum

**Problem (LeetCode 930):** Binary array. Count subarrays with sum exactly `goal`.

`[1 0 1 0 1]`, goal 2 → 4

## Approach 1: Prefix sum + hash map
Same as [count subarrays with sum K](../../../step-03-arrays/3.2-medium/13-count-subarrays-sum-k/). **Time O(n), Space O(n)**.

## Approach 2: atMost(goal) − atMost(goal − 1) (optimal, O(1) space)
"Exactly" is hard for a window, "at most" is easy:
```
atMost(g):
    if g < 0: return 0
    for r: sum += a[r]
           while sum > g: sum -= a[l]; l++
           count += r - l + 1        // every subarray ending at r and starting in [l, r]
```
Why `r − l + 1`? All those subarrays end at `r` and have sum ≤ g (smaller windows inside a valid window are valid too, since elements are ≥ 0).

`exactly(goal) = atMost(goal) − atMost(goal − 1)`
- **Time:** O(n), **Space:** O(1)
