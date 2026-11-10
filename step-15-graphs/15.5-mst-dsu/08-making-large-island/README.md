# Making a Large Island

**Problem:** `n × n` binary grid. You may change **at most one** `0` into `1`. What's the size of the largest island you can get?

```
1 0
0 1   → flip (0,1): island of size 3
```

## Brute force
Try every 0, flip it, run a full BFS for the biggest island. O(n⁴). Too slow.

## Idea: precompute islands once with DSU
1. **Union** every pair of neighbouring 1-cells. Now each island is one DSU group and `Size(root)` is its area.
2. For each `0` cell: look at its up to 4 neighbours that are 1. Collect their **distinct** roots (two neighbours can be the same island!). New size = `1 + sum of those islands' sizes`.
3. Answer = max over all 0s. If there are **no** 0s, the whole grid is one island: `n · n`.

## Dry run
```
1 1 0
0 1 0
1 0 1
```
- Islands: A = {(0,0),(0,1),(1,1)} size 3; B = {(2,0)} 1; C = {(2,2)} 1
- Flip (0,2): neighbours (0,1)=A, (1,2)=0 → 1 + 3 = 4
- Flip (1,0): neighbours (0,0)=A, (1,1)=A (same!), (2,0)=B → 1 + 3 + 1 = 5
- Flip (2,1): neighbours (1,1)=A, (2,0)=B, (2,2)=C → 1 + 3 + 1 + 1 = **6** ✅

- **Time:** O(n² · α), **Space:** O(n²)
