package bfs

// BFS returns nodes in BFS order from start. Time O(n + m).
func BFS(adj [][]int, start int) []int {
	visited := make([]bool, len(adj))
	visited[start] = true
	q := []int{start}
	order := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			if !visited[v] {
				visited[v] = true // mark when pushing
				q = append(q, v)
			}
		}
	}
	return order
}

// Distances returns the number of edges from start to every node (-1 if unreachable).
func Distances(adj [][]int, start int) []int {
	dist := make([]int, len(adj))
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	q := []int{start}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				q = append(q, v)
			}
		}
	}
	return dist
}
