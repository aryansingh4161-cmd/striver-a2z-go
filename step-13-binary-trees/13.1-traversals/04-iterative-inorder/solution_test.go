package iterativeinorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestInorder(t *testing.T) {
	if got := Inorder(tree.FromLevel([]int{1, 2, 3, 4, 5, tree.Null, 6})); !slices.Equal(got, []int{4, 2, 5, 1, 3, 6}) {
		t.Errorf("got %v", got)
	}
	if len(Inorder(nil)) != 0 {
		t.Error("nil")
	}
}
