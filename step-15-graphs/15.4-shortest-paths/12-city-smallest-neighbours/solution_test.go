package citysmallestneighbours

import "testing"

func TestFindCity(t *testing.T) {
	if got := FindCity(4, [][3]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4); got != 3 {
		t.Errorf("got %d want 3", got)
	}
	edges := [][3]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	if got := FindCity(5, edges, 2); got != 0 {
		t.Errorf("got %d want 0", got)
	}
	if got := FindCity(3, nil, 5); got != 2 {
		t.Errorf("no roads: all counts 0, largest wins; got %d", got)
	}
}
