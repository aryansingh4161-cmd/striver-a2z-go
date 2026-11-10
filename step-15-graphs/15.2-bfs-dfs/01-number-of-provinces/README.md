# Number of Provinces

**Problem (LeetCode 547):** `isConnected[i][j] = 1` if city `i` and `j` are directly connected. A province is a group of cities connected directly or indirectly. Count provinces.

```
1 1 0
1 1 0     → 2
0 0 1
```

## Approach
This is [connected components](../../15.1-learning/02-connected-components/) where the graph is given as an **adjacency matrix**. DFS from every unvisited city; count how many times you start.
- **Time:** O(n²) (reading the matrix), **Space:** O(n)

Disjoint Set Union (Step 15.5) also solves it.
