package bottomview

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBottomView(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{20, 8, 22, 5, 3, N, 25, N, N, 10, 14})
	if got := BottomView(r); !slices.Equal(got, []int{5, 10, 3, 14, 25}) {
		t.Errorf("got %v", got)
	}
}
