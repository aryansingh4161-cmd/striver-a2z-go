package bfs

import (
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {
	adj := [][]int{{1, 2}, {0, 3, 4}, {0, 5}, {1}, {1}, {2}, {}}
	if got := BFS(adj, 0); !slices.Equal(got, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("got %v", got)
	}
	if got := Distances(adj, 0); !slices.Equal(got, []int{0, 1, 1, 2, 2, 2, -1}) {
		t.Errorf("dist %v", got)
	}
}
