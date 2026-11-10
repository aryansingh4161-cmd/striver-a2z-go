# Minimum Time to Burn a Tree

**Problem:** A fire starts at node `start`. Every second it spreads to all neighbours (children and parent). How many seconds until the whole tree burns?

```
        1
      /   \
     2     3
    /     / \
   4     5   6
```
start 1 → 2 seconds; start 4 → 4 seconds (4 → 2 → 1 → 3 → 5/6)

## Approach: Parent map + BFS by levels
Same as [nodes at distance K](../05-nodes-at-distance-k/): treat the tree as an undirected graph and BFS from `start`. The answer is the number of BFS levels after the first (the distance to the farthest node).
- **Time:** O(n), **Space:** O(n)
