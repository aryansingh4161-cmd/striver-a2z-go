package kosarajuscc

import "sort"

// SCC returns the strongly connected components (each sorted, list sorted
// by first node). Kosaraju, O(V + E).
func SCC(adj [][]int) [][]int {
	n := len(adj)
	// 1. finish order on the original graph
	visited := make([]bool, n)
	order := make([]int, 0, n)
	var dfs1 func(int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		order = append(order, u)
	}
	for s := 0; s < n; s++ {
		if !visited[s] {
			dfs1(s)
		}
	}
	// 2. reverse the graph
	rev := make([][]int, n)
	for u := range adj {
		for _, v := range adj[u] {
			rev[v] = append(rev[v], u)
		}
	}
	// 3. DFS on the reversed graph by decreasing finish time
	for i := range visited {
		visited[i] = false
	}
	var comps [][]int
	var cur []int
	var dfs2 func(int)
	dfs2 = func(u int) {
		visited[u] = true
		cur = append(cur, u)
		for _, v := range rev[u] {
			if !visited[v] {
				dfs2(v)
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if u := order[i]; !visited[u] {
			cur = nil
			dfs2(u)
			sort.Ints(cur)
			comps = append(comps, cur)
		}
	}
	sort.Slice(comps, func(i, j int) bool { return comps[i][0] < comps[j][0] })
	return comps
}

// Count returns the number of SCCs.
func Count(adj [][]int) int { return len(SCC(adj)) }
