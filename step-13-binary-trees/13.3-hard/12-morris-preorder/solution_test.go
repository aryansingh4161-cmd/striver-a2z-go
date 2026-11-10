package morrispreorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestPreorder(t *testing.T) {
	N := tree.Null
	lvl := []int{1, 2, 3, 4, 5, N, 6, N, N, 7}
	r := tree.FromLevel(lvl)
	if got := Preorder(r); !slices.Equal(got, []int{1, 2, 4, 5, 7, 3, 6}) {
		t.Errorf("got %v", got)
	}
	if got := tree.ToLevel(r); !slices.Equal(got, lvl) {
		t.Errorf("tree not restored: %v", got)
	}
}
