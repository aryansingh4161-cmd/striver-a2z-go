package numberofenclaves

import "testing"

func TestNumEnclaves(t *testing.T) {
	if NumEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}) != 3 {
		t.Error("example 1")
	}
	if NumEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}) != 0 {
		t.Error("example 2")
	}
}
