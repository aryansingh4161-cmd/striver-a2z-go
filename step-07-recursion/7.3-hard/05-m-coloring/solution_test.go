package mcoloring

import "testing"

func TestCanColor(t *testing.T) {
	square := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}}
	if !CanColor(4, square, 3) {
		t.Error("should be 3-colourable")
	}
	if CanColor(4, square, 2) {
		t.Error("triangle 0-1-2 needs 3 colours")
	}
	triangle := [][2]int{{0, 1}, {1, 2}, {0, 2}}
	if CanColor(3, triangle, 2) || !CanColor(3, triangle, 3) {
		t.Error("triangle")
	}
	k4 := [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}
	if CanColor(4, k4, 3) || !CanColor(4, k4, 4) {
		t.Error("K4")
	}
	if !CanColor(3, nil, 1) {
		t.Error("no edges, 1 colour")
	}
}
