package startofloop

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		a   []int
		pos int
	}{{[]int{3, 2, 0, -4}, 1}, {[]int{1, 2}, 0}, {[]int{1}, -1}, {[]int{1, 2, 3, 4, 5, 6}, 3}, {[]int{1}, 0}}
	for _, tc := range tests {
		for name, fn := range map[string]func(*Node) *Node{"floyd": DetectCycle, "set": DetectCycleSet} {
			head := list.WithCycle(tc.a, tc.pos)
			var want *Node
			if tc.pos >= 0 {
				want = list.NodeAt(head, tc.pos)
			}
			if got := fn(head); got != want {
				t.Errorf("%s(%v,%d) wrong node", name, tc.a, tc.pos)
			}
		}
	}
}
