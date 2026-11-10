package courseschedule

// FindOrder returns a valid order to take all courses, or nil if impossible.
// pre[i] = [a, b] means b must come before a (edge b → a). O(V + E).
func FindOrder(n int, pre [][2]int) []int {
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, p := range pre {
		a, b := p[0], p[1]
		adj[b] = append(adj[b], a)
		indeg[a]++
	}
	queue := []int{}
	for v := 0; v < n; v++ {
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}
	order := make([]int, 0, n)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(order) < n {
		return nil // a cycle blocks some courses
	}
	return order
}

// CanFinish is Course Schedule I.
func CanFinish(n int, pre [][2]int) bool {
	return FindOrder(n, pre) != nil
}
