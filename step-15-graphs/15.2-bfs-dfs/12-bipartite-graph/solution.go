package bipartitegraph

// IsBipartiteBFS 2-colours the graph with BFS. O(n + m).
func IsBipartiteBFS(adj [][]int) bool {
	color := make([]int, len(adj))
	for i := range color {
		color[i] = -1
	}
	for s := range adj {
		if color[s] != -1 {
			continue
		}
		color[s] = 0
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range adj[u] {
				if color[v] == -1 {
					color[v] = 1 - color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					return false
				}
			}
		}
	}
	return true
}

// IsBipartiteDFS 2-colours the graph with DFS. O(n + m).
func IsBipartiteDFS(adj [][]int) bool {
	color := make([]int, len(adj))
	for i := range color {
		color[i] = -1
	}
	var dfs func(u, c int) bool
	dfs = func(u, c int) bool {
		color[u] = c
		for _, v := range adj[u] {
			if color[v] == -1 {
				if !dfs(v, 1-c) {
					return false
				}
			} else if color[v] == c {
				return false
			}
		}
		return true
	}
	for s := range adj {
		if color[s] == -1 && !dfs(s, 0) {
			return false
		}
	}
	return true
}
