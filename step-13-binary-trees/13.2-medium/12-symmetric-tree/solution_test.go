package symmetrictree

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestIsSymmetric(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want bool
	}{{[]int{1, 2, 2, 3, 4, 4, 3}, true}, {[]int{1, 2, 2, N, 3, N, 3}, false}, {[]int{}, true}, {[]int{1}, true}, {[]int{1, 2, 3}, false}}
	for _, tc := range tests {
		if got := IsSymmetric(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("IsSymmetric(%v) = %v", tc.lvl, got)
		}
	}
}
