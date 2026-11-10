# Number of Distinct Islands

**Problem:** Count islands with **different shapes**. Two islands are the same if one can be **moved (translated)** onto the other (no rotation/reflection).

```
1 1 0 1 1
1 0 0 0 0
0 0 0 1 1
1 1 0 1 0
```
(two L-shapes and two 1×2 bars → **2** distinct shapes)

## Approach: Record each island's shape relative to its first cell
DFS each island. For every cell, store `(r − r0, c − c0)` where `(r0, c0)` is the cell where the DFS started (the top-left-most first cell found by the scan). Because the scan and DFS order are fixed, identical shapes produce identical lists. Put each list (as a string key) into a set; the answer is the set size.
- **Time:** O(n × m), **Space:** O(n × m)
