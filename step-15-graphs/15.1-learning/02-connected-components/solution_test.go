package connectedcomponents

import "testing"

func TestCount(t *testing.T) {
	if Count(5, [][2]int{{0, 1}, {1, 2}, {3, 4}}) != 2 || Count(4, nil) != 4 || Count(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}) != 1 {
		t.Error("examples")
	}
}
