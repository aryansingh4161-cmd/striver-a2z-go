package cycleundirected

import "testing"

func build(n int, edges [][2]int) [][]int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	return adj
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}, {3, 4}}, true},
		{4, [][2]int{{0, 1}, {1, 2}, {2, 3}}, false},
		{6, [][2]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true}, // cycle in second component
		{3, nil, false},
	}
	for _, tc := range tests {
		adj := build(tc.n, tc.edges)
		if HasCycleBFS(adj) != tc.want || HasCycleDFS(adj) != tc.want {
			t.Errorf("%v", tc.edges)
		}
	}
}
