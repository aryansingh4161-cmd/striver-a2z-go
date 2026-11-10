package checkbalanced

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestIsBalanced(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want bool
	}{{[]int{3, 9, 20, N, N, 15, 7}, true}, {[]int{1, 2, 2, 3, 3, N, N, 4, 4}, false}, {[]int{}, true}, {[]int{1, 2, N, 3}, false}}
	for _, tc := range tests {
		if got := IsBalanced(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("IsBalanced(%v) = %v", tc.lvl, got)
		}
	}
}
