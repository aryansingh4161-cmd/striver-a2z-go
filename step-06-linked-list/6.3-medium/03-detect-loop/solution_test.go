package detectloop

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		a    []int
		pos  int
		want bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
		{[]int{1}, 0, true},
		{[]int{1, 2, 3}, -1, false},
		{[]int{}, -1, false},
	}
	for _, tc := range tests {
		if got := HasCycle(list.WithCycle(tc.a, tc.pos)); got != tc.want {
			t.Errorf("HasCycle(%v,%d) = %v", tc.a, tc.pos, got)
		}
		if got := HasCycleSet(list.WithCycle(tc.a, tc.pos)); got != tc.want {
			t.Errorf("HasCycleSet(%v,%d) = %v", tc.a, tc.pos, got)
		}
	}
}
