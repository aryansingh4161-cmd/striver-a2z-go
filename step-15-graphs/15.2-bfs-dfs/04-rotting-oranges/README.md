# Rotting Oranges

**Problem (LeetCode 994):** `0` empty, `1` fresh, `2` rotten. Every minute, fresh oranges next to a rotten one (4 directions) become rotten. Minimum minutes until no fresh orange is left, or −1 if impossible.

```
2 1 1
1 1 0     → 4
0 1 1
```

## Approach: Multi-source BFS
All rotten oranges start rotting **at the same time**, so push **all of them** into the queue at minute 0. Then BFS level by level; each level is one minute. Count fresh oranges; decrease as they rot.
At the end: if fresh > 0 → −1, else the number of minutes.
- **Time:** O(n × m), **Space:** O(n × m)

"Multi-source BFS" (start from many cells at once) also solves [distance to nearest 0](../06-distance-nearest-zero/).
