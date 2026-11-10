package diameter

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestDiameter(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want int
	}{{[]int{1, 2, 3, 4, 5}, 3}, {[]int{1, 2}, 1}, {[]int{1}, 0}, {[]int{1, 2, N, 3, 4, 5, N, N, 6}, 4}}
	for _, tc := range tests {
		if got := Diameter(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("Diameter(%v) = %d", tc.lvl, got)
		}
	}
}
