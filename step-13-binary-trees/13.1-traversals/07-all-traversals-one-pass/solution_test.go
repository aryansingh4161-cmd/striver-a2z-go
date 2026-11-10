package alltraversalsonepass

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestAll(t *testing.T) {
	pre, in, post := AllTraversals(tree.FromLevel([]int{1, 2, 3, 4, 5}))
	if !slices.Equal(pre, []int{1, 2, 4, 5, 3}) || !slices.Equal(in, []int{4, 2, 5, 1, 3}) || !slices.Equal(post, []int{4, 5, 2, 3, 1}) {
		t.Errorf("%v %v %v", pre, in, post)
	}
}
