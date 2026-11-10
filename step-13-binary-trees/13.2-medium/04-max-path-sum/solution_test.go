package maxpathsum

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestMaxPathSum(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want int
	}{{[]int{1, 2, 3}, 6}, {[]int{-10, 9, 20, N, N, 15, 7}, 42}, {[]int{-3}, -3}, {[]int{2, -1}, 2}, {[]int{-2, -1}, -1}}
	for _, tc := range tests {
		if got := MaxPathSum(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("MaxPathSum(%v) = %d", tc.lvl, got)
		}
	}
}
