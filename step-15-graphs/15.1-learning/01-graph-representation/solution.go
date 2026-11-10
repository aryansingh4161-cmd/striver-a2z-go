package graphrepresentation

// Edge is a weighted directed edge (to, weight).
type Edge struct{ To, W int }

// AdjMatrix builds an n x n matrix for an undirected graph.
func AdjMatrix(n int, edges [][2]int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	for _, e := range edges {
		m[e[0]][e[1]], m[e[1]][e[0]] = 1, 1
	}
	return m
}

// AdjList builds an adjacency list. directed=false adds both directions.
func AdjList(n int, edges [][2]int, directed bool) [][]int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		if !directed {
			adj[e[1]] = append(adj[e[1]], e[0])
		}
	}
	return adj
}

// WeightedAdj builds a weighted adjacency list from (u, v, w) triples.
func WeightedAdj(n int, edges [][3]int, directed bool) [][]Edge {
	adj := make([][]Edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], Edge{e[1], e[2]})
		if !directed {
			adj[e[1]] = append(adj[e[1]], Edge{e[0], e[2]})
		}
	}
	return adj
}

// InDegrees returns the in-degree of every node in a directed adjacency list.
func InDegrees(adj [][]int) []int {
	in := make([]int, len(adj))
	for _, nbrs := range adj {
		for _, v := range nbrs {
			in[v]++
		}
	}
	return in
}
