package morrisinorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestInorder(t *testing.T) {
	N := tree.Null
	lvl := []int{1, 2, 3, 4, 5, N, 6, N, N, 7}
	r := tree.FromLevel(lvl)
	if got := Inorder(r); !slices.Equal(got, []int{4, 2, 7, 5, 1, 3, 6}) {
		t.Errorf("got %v", got)
	}
	if got := tree.ToLevel(r); !slices.Equal(got, lvl) {
		t.Errorf("tree not restored: %v", got)
	}
}
