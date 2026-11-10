package numberofprovinces

import "testing"

func TestFindCircleNum(t *testing.T) {
	if FindCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}) != 2 || FindCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}) != 3 {
		t.Error("examples")
	}
}
