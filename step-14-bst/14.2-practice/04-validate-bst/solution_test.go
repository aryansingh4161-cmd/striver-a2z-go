package validatebst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestIsValid(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want bool
	}{{[]int{2, 1, 3}, true}, {[]int{5, 1, 4, N, N, 3, 6}, false}, {[]int{5, 4, 6, N, N, 3, 7}, false}, {[]int{2, 2, 2}, false}, {[]int{}, true}}
	for _, tc := range tests {
		if got := IsValid(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("IsValid(%v) = %v", tc.lvl, got)
		}
	}
}
