package numberofprovinces

// FindCircleNum counts connected components in an adjacency matrix. O(n^2).
func FindCircleNum(m [][]int) int {
	n := len(m)
	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for v := 0; v < n; v++ {
			if m[u][v] == 1 && !visited[v] {
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
