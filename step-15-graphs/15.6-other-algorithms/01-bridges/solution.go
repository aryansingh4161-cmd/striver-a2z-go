package bridges

import "sort"

// Bridges returns every bridge [u, v] (u < v), sorted. Tarjan, O(V + E).
func Bridges(n int, edges [][2]int) [][2]int {
	type half struct{ to, id int }
	adj := make([][]half, n)
	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], half{e[1], i})
		adj[e[1]] = append(adj[e[1]], half{e[0], i})
	}
	tin, low := make([]int, n), make([]int, n)
	visited := make([]bool, n)
	timer := 0
	var out [][2]int
	var dfs func(u, parentEdge int)
	dfs = func(u, parentEdge int) {
		visited[u] = true
		tin[u], low[u] = timer, timer
		timer++
		for _, h := range adj[u] {
			if h.id == parentEdge {
				continue // don't go back along the same edge
			}
			if !visited[h.to] {
				dfs(h.to, h.id)
				low[u] = min(low[u], low[h.to])
				if low[h.to] > tin[u] {
					a, b := u, h.to
					if a > b {
						a, b = b, a
					}
					out = append(out, [2]int{a, b})
				}
			} else {
				low[u] = min(low[u], tin[h.to]) // back edge
			}
		}
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs(s, -1)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i][0] != out[j][0] {
			return out[i][0] < out[j][0]
		}
		return out[i][1] < out[j][1]
	})
	return out
}
