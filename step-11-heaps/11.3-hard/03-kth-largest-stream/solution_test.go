package kthlargeststream

import "testing"

func TestKthLargest(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	for _, step := range [][2]int{{3, 4}, {5, 5}, {10, 5}, {9, 8}, {4, 8}} {
		if got := kl.Add(step[0]); got != step[1] {
			t.Errorf("Add(%d) = %d, want %d", step[0], got, step[1])
		}
	}
}
