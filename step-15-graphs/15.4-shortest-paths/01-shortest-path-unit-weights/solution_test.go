package shortestpathunitweights

import (
	"slices"
	"testing"
)

func TestShortestPath(t *testing.T) {
	adj := FromEdges(9, [][2]int{{0, 1}, {0, 3}, {1, 2}, {3, 4}, {4, 5}, {5, 6}, {2, 6}, {6, 7}, {6, 8}, {7, 8}})
	if got := ShortestPath(adj, 0); !slices.Equal(got, []int{0, 1, 2, 1, 2, 3, 3, 4, 4}) {
		t.Errorf("got %v", got)
	}
	adj = FromEdges(3, [][2]int{{0, 1}})
	if got := ShortestPath(adj, 0); !slices.Equal(got, []int{0, 1, -1}) {
		t.Errorf("unreachable: got %v", got)
	}
}
