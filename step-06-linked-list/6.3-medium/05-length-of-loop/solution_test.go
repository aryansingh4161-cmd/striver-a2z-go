package lengthofloop

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestLoopLength(t *testing.T) {
	tests := []struct {
		a         []int
		pos, want int
	}{{[]int{1, 2, 3, 4, 5}, 1, 4}, {[]int{1, 2, 3}, -1, 0}, {[]int{1}, 0, 1}, {[]int{1, 2, 3, 4, 5, 6}, 5, 1}}
	for _, tc := range tests {
		if got := LoopLength(list.WithCycle(tc.a, tc.pos)); got != tc.want {
			t.Errorf("LoopLength(%v,%d) = %d", tc.a, tc.pos, got)
		}
		if got := LoopLengthMap(list.WithCycle(tc.a, tc.pos)); got != tc.want {
			t.Errorf("LoopLengthMap(%v,%d) = %d", tc.a, tc.pos, got)
		}
	}
}
