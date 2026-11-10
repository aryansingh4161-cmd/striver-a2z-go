package roottonodepath

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestPath(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{1, 2, 3, 4, 5, N, N, N, N, 6, 7})
	if got := PathTo(r, 7); !slices.Equal(got, []int{1, 2, 5, 7}) {
		t.Errorf("got %v", got)
	}
	if got := PathTo(r, 9); got != nil {
		t.Errorf("missing: %v", got)
	}
	if got := AllRootToLeafPaths(tree.FromLevel([]int{1, 2, 3, N, 5})); !slices.Equal(got, []string{"1->2->5", "1->3"}) {
		t.Errorf("got %v", got)
	}
}
