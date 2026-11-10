# House Robber II (Circular)

**Problem:** Same as [non-adjacent max sum](../04-max-sum-non-adjacent/), but the houses are in a **circle**: the first and last houses are neighbours too.

```
[2 3 2]   → 3 (can't take both 2s, they touch)
[1 2 3 1] → 4 (1 + 3)
```

## Idea: break the circle in two ways
The first and last house can't **both** be taken. So the answer is one of:
1. Best of `nums[0 .. n−2]` (**last** house excluded), or
2. Best of `nums[1 .. n−1]` (**first** house excluded).

Each is the ordinary straight-line problem. Take the max.

```
if n == 1: return nums[0]
return max( linear(nums[:n-1]), linear(nums[1:]) )
```

## Dry run (`[1 2 3 1]`)
- without last: `[1 2 3]` → 1 + 3 = 4
- without first: `[2 3 1]` → 3 (or 2 + 1 = 3)
- answer **4** ✅

## Edge case
`n == 1`: both slices would be empty, so return the single house directly.

- **Time:** O(n), **Space:** O(1)
