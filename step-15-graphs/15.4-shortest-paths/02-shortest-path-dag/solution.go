package shortestpathdag

import "math"

// ShortestPath returns distances from node 0 in a weighted DAG
// (edges are [u, v, w]); -1 marks unreachable nodes. O(V + E).
func ShortestPath(n int, edges [][3]int) []int {
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
	}

	// 1. topological order with DFS finish times
	visited := make([]bool, n)
	order := make([]int, 0, n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for _, e := range adj[u] {
			if !visited[e.to] {
				dfs(e.to)
			}
		}
		order = append(order, u)
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs(s)
		}
	}

	// 2. relax edges in topological order (order is reversed finish order)
	const inf = math.MaxInt
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0
	for i := n - 1; i >= 0; i-- {
		u := order[i]
		if dist[u] == inf {
			continue
		}
		for _, e := range adj[u] {
			dist[e.to] = min(dist[e.to], dist[u]+e.w)
		}
	}
	for i := range dist {
		if dist[i] == inf {
			dist[i] = -1
		}
	}
	return dist
}
