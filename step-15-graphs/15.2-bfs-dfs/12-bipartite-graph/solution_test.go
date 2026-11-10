package bipartitegraph

import "testing"

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		adj  [][]int
		want bool
	}{
		{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
		{[][]int{{}, {2}, {1}}, true},
		{[][]int{{1}, {0}, {3, 4}, {2, 4}, {2, 3}}, false}, // odd cycle in 2nd component
	}
	for _, tc := range tests {
		if IsBipartiteBFS(tc.adj) != tc.want || IsBipartiteDFS(tc.adj) != tc.want {
			t.Errorf("%v", tc.adj)
		}
	}
}
