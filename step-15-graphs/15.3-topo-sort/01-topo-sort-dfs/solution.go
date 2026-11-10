package toposortdfs

// TopoSort returns a topological order of a DAG using DFS finish times.
// O(V + E) time, O(V) space.
func TopoSort(adj [][]int) []int {
	n := len(adj)
	visited := make([]bool, n)
	stack := make([]int, 0, n) // nodes in the order they finish
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs(v)
			}
		}
		stack = append(stack, u) // all of u's descendants are already in
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs(s)
		}
	}
	// reverse the finish order = popping the stack
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

// IsTopoOrder checks that order is a valid topological order of adj.
// Shared by the tests of this step.
func IsTopoOrder(adj [][]int, order []int) bool {
	n := len(adj)
	if len(order) != n {
		return false
	}
	pos := make([]int, n)
	for i := range pos {
		pos[i] = -1
	}
	for i, u := range order {
		if u < 0 || u >= n || pos[u] != -1 {
			return false
		}
		pos[u] = i
	}
	for u := range adj {
		for _, v := range adj[u] {
			if pos[u] > pos[v] {
				return false
			}
		}
	}
	return true
}
