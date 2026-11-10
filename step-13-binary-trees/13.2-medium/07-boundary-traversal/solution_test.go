package boundarytraversal

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBoundary(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{1, 2, 7, 3, N, N, 8, N, 4, 9, N, 5, 6, 10, 11})
	if got := Boundary(r); !slices.Equal(got, []int{1, 2, 3, 4, 5, 6, 10, 11, 9, 8, 7}) {
		t.Errorf("got %v", got)
	}
	if got := Boundary(tree.FromLevel([]int{1})); !slices.Equal(got, []int{1}) {
		t.Errorf("single: %v", got)
	}
	if got := Boundary(tree.FromLevel([]int{1, 2, 3, 4, 5, 6, 7})); !slices.Equal(got, []int{1, 2, 4, 5, 6, 7, 3}) {
		t.Errorf("full: %v", got)
	}
}
