package iterativepreorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestPreorder(t *testing.T) {
	if got := Preorder(tree.FromLevel([]int{1, 2, 3, 4, 5, tree.Null, 6})); !slices.Equal(got, []int{1, 2, 4, 5, 3, 6}) {
		t.Errorf("got %v", got)
	}
	if len(Preorder(nil)) != 0 {
		t.Error("nil")
	}
}
