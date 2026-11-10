package mcoloring

// CanColor reports whether the graph (n nodes, undirected edges) can be
// coloured with m colours. Time O(m^n) worst case.
func CanColor(n int, edges [][2]int, m int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	color := make([]int, n) // 0 = not coloured yet
	var f func(node int) bool
	f = func(node int) bool {
		if node == n {
			return true
		}
		for c := 1; c <= m; c++ {
			ok := true
			for _, nb := range adj[node] {
				if color[nb] == c {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			color[node] = c
			if f(node + 1) {
				return true
			}
			color[node] = 0
		}
		return false
	}
	return f(0)
}
