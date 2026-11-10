package cycleundirected

// HasCycleBFS detects a cycle in an undirected graph using BFS. O(n + m).
func HasCycleBFS(adj [][]int) bool {
	n := len(adj)
	visited := make([]bool, n)
	for s := 0; s < n; s++ {
		if visited[s] {
			continue
		}
		visited[s] = true
		q := [][2]int{{s, -1}}
		for len(q) > 0 {
			u, parent := q[0][0], q[0][1]
			q = q[1:]
			for _, v := range adj[u] {
				if !visited[v] {
					visited[v] = true
					q = append(q, [2]int{v, u})
				} else if v != parent {
					return true
				}
			}
		}
	}
	return false
}

// HasCycleDFS detects a cycle in an undirected graph using DFS. O(n + m).
func HasCycleDFS(adj [][]int) bool {
	visited := make([]bool, len(adj))
	var dfs func(u, parent int) bool
	dfs = func(u, parent int) bool {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				if dfs(v, u) {
					return true
				}
			} else if v != parent {
				return true
			}
		}
		return false
	}
	for s := range adj {
		if !visited[s] && dfs(s, -1) {
			return true
		}
	}
	return false
}
