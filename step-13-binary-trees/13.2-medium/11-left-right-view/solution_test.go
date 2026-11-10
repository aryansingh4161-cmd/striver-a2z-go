package leftrightview

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestViews(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{1, 2, 3, N, 5, N, 4})
	if got := RightView(r); !slices.Equal(got, []int{1, 3, 4}) {
		t.Errorf("right %v", got)
	}
	if got := LeftView(r); !slices.Equal(got, []int{1, 2, 5}) {
		t.Errorf("left %v", got)
	}
	r2 := tree.FromLevel([]int{1, 2, 3, 4})
	if got := RightView(r2); !slices.Equal(got, []int{1, 3, 4}) {
		t.Errorf("right deep-left %v", got)
	}
}
