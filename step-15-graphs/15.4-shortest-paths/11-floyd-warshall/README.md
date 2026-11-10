# Floyd-Warshall Algorithm

**Problem:** Find the shortest distance between **every pair** of nodes. Input is an `n × n` matrix where `m[i][j]` is the edge weight, `-1` means no edge. Replace each entry with the shortest distance (`-1` if unreachable). Also detect negative cycles.

## Idea: allow one more "via" node at a time
`dist[i][j]` starts as the direct edge. Then for each node `k`, ask: *is going `i → k → j` shorter than what I have?*
```
for k in 0..n-1:            // k MUST be the outer loop
    for i in 0..n-1:
        for j in 0..n-1:
            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
```
After processing `k`, `dist[i][j]` is the shortest path whose middle nodes are all from `{0..k}`. After the last `k`, any middle node is allowed → final answer.

**Why must `k` be outermost?** Because each "via k" step needs all paths using via-nodes `{0..k-1}` to be finished first.

## Negative cycle
After the algorithm, if `dist[i][i] < 0` for some `i`, then `i` sits on a negative cycle (going around it from `i` back to `i` costs less than 0).

## Dry run
```
0→1 (2), 1→0 (1), 1→2 (3), 3→0 (3), 3→1 (5), 3→2 (4)
```
- k=0: 3→1 via 0 = 3+2 = 5 (same), 1→... nothing new
- k=1: 0→2 via 1 = 2+3 = 5 (new!)
- k=2, k=3: nothing new (no edges leave 2; nothing enters 3)
```
 0  2  5 -1
 1  0  3 -1
-1 -1  0 -1
 3  5  4  0
```
✅

- **Time:** O(n³), **Space:** O(n²) (in place)

## When to use what
- All pairs, small `n` (≤ ~400), maybe negative weights → **Floyd-Warshall** (also the easiest to code).
- All pairs, big sparse graph, no negative weights → Dijkstra from every node: O(V · E log V).
