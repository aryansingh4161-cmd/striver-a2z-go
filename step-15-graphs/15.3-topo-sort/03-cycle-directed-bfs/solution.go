package cycledirectedbfs

// HasCycle reports whether the directed graph has a cycle, using Kahn's
// algorithm: a cycle exists iff not every node can be removed. O(V + E).
func HasCycle(adj [][]int) bool {
	n := len(adj)
	indeg := make([]int, n)
	for u := range adj {
		for _, v := range adj[u] {
			indeg[v]++
		}
	}
	queue := []int{}
	for v := 0; v < n; v++ {
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}
	processed := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		processed++
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return processed < n
}
