package maxdepth

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		lvl  []int
		want int
	}{{[]int{3, 9, 20, tree.Null, tree.Null, 15, 7}, 3}, {[]int{1, tree.Null, 2}, 2}, {[]int{}, 0}, {[]int{1}, 1}}
	for _, tc := range tests {
		r := tree.FromLevel(tc.lvl)
		if MaxDepth(r) != tc.want || MaxDepthBFS(r) != tc.want {
			t.Errorf("%v", tc.lvl)
		}
	}
}
