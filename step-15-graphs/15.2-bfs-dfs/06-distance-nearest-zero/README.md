# Distance of Nearest 0 (01 Matrix)

**Problem (LeetCode 542):** For each cell, the distance (in 4-direction steps) to the nearest `0`.

```
0 0 0        0 0 0
0 1 0   →    0 1 0
1 1 1        1 2 1
```

## Approach 1: BFS from every 1
O((n·m)²). Too slow.

## Approach 2: Multi-source BFS from all 0s (optimal)
Every `0` has distance 0: push **all** of them first. BFS outward; the first time a cell is reached is its shortest distance.
- **Time:** O(n × m), **Space:** O(n × m)

Same technique as [rotting oranges](../04-rotting-oranges/). (The GFG version asks for the nearest **1** instead: just swap the roles.)
