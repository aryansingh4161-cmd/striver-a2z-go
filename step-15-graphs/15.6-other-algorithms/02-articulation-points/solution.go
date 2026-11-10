package articulationpoints

// Points returns the articulation points in increasing order. O(V + E).
func Points(n int, edges [][2]int) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	tin, low := make([]int, n), make([]int, n)
	visited, mark := make([]bool, n), make([]bool, n)
	timer := 0
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		visited[u] = true
		tin[u], low[u] = timer, timer
		timer++
		children := 0
		for _, v := range adj[u] {
			if v == parent {
				continue // parallel edges don't matter for cut vertices
			}
			if !visited[v] {
				dfs(v, u)
				children++
				low[u] = min(low[u], low[v])
				if parent != -1 && low[v] >= tin[u] {
					mark[u] = true
				}
			} else {
				low[u] = min(low[u], tin[v])
			}
		}
		if parent == -1 && children > 1 {
			mark[u] = true
		}
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs(s, -1)
		}
	}
	out := []int{}
	for u := 0; u < n; u++ {
		if mark[u] {
			out = append(out, u)
		}
	}
	return out
}
