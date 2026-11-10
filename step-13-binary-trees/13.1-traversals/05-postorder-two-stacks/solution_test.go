package postordertwostacks

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestPostorder(t *testing.T) {
	if got := Postorder(tree.FromLevel([]int{1, 2, 3, 4, 5, tree.Null, 6})); !slices.Equal(got, []int{4, 5, 2, 6, 3, 1}) {
		t.Errorf("got %v", got)
	}
}
