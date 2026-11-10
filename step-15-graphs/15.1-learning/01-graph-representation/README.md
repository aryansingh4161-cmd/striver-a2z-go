# Graph Representation

Input usually comes as `n` nodes (numbered `0..n−1` or `1..n`) and a list of edges `(u, v)`.

## 1. Adjacency matrix
`adj[u][v] = 1` if there's an edge.
- **Space:** O(n²): fine for n ≤ ~1000, too big beyond that.
- Check "is there an edge u→v?" in O(1).

## 2. Adjacency list (use this by default)
`adj[u]` = slice of neighbours of `u`.
- **Space:** O(n + m) (m = edges; each undirected edge is stored twice).
- Visit all neighbours of `u` in O(degree(u)).

```go
adj := make([][]int, n)
for _, e := range edges {
    adj[e[0]] = append(adj[e[0]], e[1])
    adj[e[1]] = append(adj[e[1]], e[0]) // undirected only
}
```

## Weighted graphs
Store pairs: `adj[u] = append(adj[u], Edge{To: v, W: w})`.

## Degree
- Undirected: `degree(u) = len(adj[u])`, and the sum of all degrees = `2m`.
- Directed: **in-degree** (edges coming in) and **out-degree** (going out). In-degrees drive Kahn's topological sort.
