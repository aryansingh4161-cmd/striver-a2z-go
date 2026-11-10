package connectedcomponents

// Count returns the number of connected components in an undirected graph.
// Time O(n + m), space O(n).
func Count(n int, edges [][2]int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	count := 0
	for u := 0; u < n; u++ {
		if !visited[u] {
			count++
			dfs(u)
		}
	}
	return count
}
