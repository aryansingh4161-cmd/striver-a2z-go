package lrupagefaults

import "testing"

func TestPageFaults(t *testing.T) {
	if PageFaults([]int{5, 0, 1, 3, 2, 4, 1, 0, 5}, 4) != 8 {
		t.Error("example 1")
	}
	if PageFaults([]int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2}, 4) != 6 {
		t.Error("example 2")
	}
	if PageFaults([]int{1, 1, 1}, 1) != 1 {
		t.Error("repeats")
	}
}
