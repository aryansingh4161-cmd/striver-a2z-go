package dfs

// DFS returns nodes in recursive DFS order from start. Time O(n + m).
func DFS(adj [][]int, start int) []int {
	visited := make([]bool, len(adj))
	order := []int{}
	var f func(int)
	f = func(u int) {
		visited[u] = true
		order = append(order, u)
		for _, v := range adj[u] {
			if !visited[v] {
				f(v)
			}
		}
	}
	f(start)
	return order
}

// DFSIterative gives the same order as DFS using an explicit stack.
func DFSIterative(adj [][]int, start int) []int {
	visited := make([]bool, len(adj))
	order := []int{}
	st := []int{start}
	for len(st) > 0 {
		u := st[len(st)-1]
		st = st[:len(st)-1]
		if visited[u] {
			continue
		}
		visited[u] = true
		order = append(order, u)
		for i := len(adj[u]) - 1; i >= 0; i-- { // reverse so the first neighbour is popped first
			if !visited[adj[u][i]] {
				st = append(st, adj[u][i])
			}
		}
	}
	return order
}
