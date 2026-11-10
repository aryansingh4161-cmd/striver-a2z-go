# Number of Operations to Make Network Connected

**Problem:** `n` computers, `connections[i] = [a, b]` are cables. In one operation you can unplug any cable and plug it between two other computers. Minimum operations to connect all computers? `-1` if impossible.

## Two observations
1. To connect `n` computers you need at least `n − 1` cables. Fewer than that → **-1**, no matter what.
2. If there are `c` separate groups (components), you need exactly `c − 1` new links to join them. Each link can come from an **extra** cable: one that closes a cycle inside a group (removing it disconnects nothing).

With `len(connections) ≥ n − 1`, there are always enough extra cables. So:
```
if len(connections) < n-1: return -1
return components - 1
```

## Counting components with DSU
Start with `n` groups; each successful Union merges two → `groups--`.

## Dry run
`n = 6`, `[[0,1],[0,2],[0,3],[1,2],[1,3]]`
- 5 cables ≥ 5 ✅
- Groups: {0,1,2,3}, {4}, {5} → 3 groups → **2** operations ✅ (cables 1-2 and 1-3 are extra)

- **Time:** O(E · α(n)), **Space:** O(n)

(DFS counting components works just as well; DSU is the Step 15.5 way.)
