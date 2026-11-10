package dfs

import (
	"slices"
	"testing"
)

func TestDFS(t *testing.T) {
	adj := [][]int{{1, 2}, {0, 3, 4}, {0, 5}, {1}, {1}, {2}}
	want := []int{0, 1, 3, 4, 2, 5}
	if got := DFS(adj, 0); !slices.Equal(got, want) {
		t.Errorf("rec %v", got)
	}
	if got := DFSIterative(adj, 0); !slices.Equal(got, want) {
		t.Errorf("iter %v", got)
	}
}
