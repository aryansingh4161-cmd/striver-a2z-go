package largestbstinbt

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestLargestBST(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want int
	}{{[]int{10, 5, 15, 1, 8, N, 7}, 3}, {[]int{5, 3, 8, 1, 4, 7, 9}, 7}, {[]int{1}, 1}, {[]int{}, 0}, {[]int{2, 3, 1}, 1}}
	for _, tc := range tests {
		if got := LargestBST(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("LargestBST(%v) = %d", tc.lvl, got)
		}
	}
}
