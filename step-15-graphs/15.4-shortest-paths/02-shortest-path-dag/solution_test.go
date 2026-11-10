package shortestpathdag

import (
	"slices"
	"testing"
)

func TestShortestPath(t *testing.T) {
	edges := [][3]int{{0, 1, 2}, {0, 4, 1}, {4, 5, 4}, {4, 2, 2}, {1, 2, 3}, {2, 3, 6}, {5, 3, 1}}
	if got := ShortestPath(6, edges); !slices.Equal(got, []int{0, 2, 3, 6, 1, 5}) {
		t.Errorf("got %v", got)
	}
	// node 2 is unreachable from 0; negative weight is allowed
	if got := ShortestPath(3, [][3]int{{0, 1, -4}, {2, 1, 1}}); !slices.Equal(got, []int{0, -4, -1}) {
		t.Errorf("got %v", got)
	}
}
