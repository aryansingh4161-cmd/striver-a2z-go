package cycledirecteddfs

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		adj  [][]int
		want bool
	}{
		{[][]int{{1, 2}, {}, {1}}, false}, // diamond-ish, no cycle
		{[][]int{{1}, {2}, {0}}, true},
		{[][]int{{1}, {2}, {3}, {4, 5}, {}, {3}}, true},
		{[][]int{{}}, false},
		{[][]int{{0}}, true}, // self loop
	}
	for _, tc := range tests {
		if got := HasCycle(tc.adj); got != tc.want {
			t.Errorf("HasCycle(%v) = %v", tc.adj, got)
		}
	}
}
