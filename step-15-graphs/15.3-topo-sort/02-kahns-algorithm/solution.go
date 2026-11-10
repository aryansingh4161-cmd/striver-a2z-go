package kahnsalgorithm

// TopoSort returns a topological order using Kahn's algorithm.
// If the graph has a cycle, the returned slice is shorter than len(adj).
// O(V + E).
func TopoSort(adj [][]int) []int {
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
	return order
}
