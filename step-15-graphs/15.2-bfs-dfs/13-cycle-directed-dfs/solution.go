package cycledirecteddfs

// HasCycle detects a cycle in a directed graph with DFS path tracking. O(n + m).
func HasCycle(adj [][]int) bool {
	n := len(adj)
	visited, onPath := make([]bool, n), make([]bool, n)
	var dfs func(int) bool
	dfs = func(u int) bool {
		visited[u], onPath[u] = true, true
		for _, v := range adj[u] {
			if onPath[v] {
				return true // back edge
			}
			if !visited[v] && dfs(v) {
				return true
			}
		}
		onPath[u] = false
		return false
	}
	for s := 0; s < n; s++ {
		if !visited[s] && dfs(s) {
			return true
		}
	}
	return false
}
