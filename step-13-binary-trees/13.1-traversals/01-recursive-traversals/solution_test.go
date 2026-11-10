package recursivetraversals

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestTraversals(t *testing.T) {
	r := tree.FromLevel([]int{1, 2, 3, 4, 5})
	if got := Preorder(r); !slices.Equal(got, []int{1, 2, 4, 5, 3}) {
		t.Errorf("pre %v", got)
	}
	if got := Inorder(r); !slices.Equal(got, []int{4, 2, 5, 1, 3}) {
		t.Errorf("in %v", got)
	}
	if got := Postorder(r); !slices.Equal(got, []int{4, 5, 2, 3, 1}) {
		t.Errorf("post %v", got)
	}
	if len(Preorder(nil)) != 0 {
		t.Error("nil")
	}
}
