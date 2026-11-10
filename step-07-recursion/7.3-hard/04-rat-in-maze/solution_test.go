package ratinmaze

import (
	"slices"
	"testing"
)

func TestFindPaths(t *testing.T) {
	maze := [][]int{{1, 0, 0, 0}, {1, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 1, 1}}
	if got := FindPaths(maze); !slices.Equal(got, []string{"DDRDRR", "DRDDRR"}) {
		t.Errorf("got %v", got)
	}
	if got := FindPaths([][]int{{1, 0}, {1, 0}}); len(got) != 0 {
		t.Errorf("blocked target: %v", got)
	}
	if got := FindPaths([][]int{{1}}); !slices.Equal(got, []string{""}) {
		t.Errorf("1x1: %v", got)
	}
	if got := FindPaths([][]int{{1, 1}, {1, 1}}); !slices.Equal(got, []string{"DR", "RD"}) {
		t.Errorf("2x2: %v", got)
	}
}
