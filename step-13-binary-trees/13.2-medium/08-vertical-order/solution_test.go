package verticalorder

import (
	"reflect"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestVerticalOrder(t *testing.T) {
	N := tree.Null
	if got := VerticalOrder(tree.FromLevel([]int{3, 9, 20, N, N, 15, 7})); !reflect.DeepEqual(got, [][]int{{9}, {3, 15}, {20}, {7}}) {
		t.Errorf("got %v", got)
	}
	// 5 and 6 share row 2, col 0 → sorted by value
	if got := VerticalOrder(tree.FromLevel([]int{1, 2, 3, 4, 6, 5, 7})); !reflect.DeepEqual(got, [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}) {
		t.Errorf("got %v", got)
	}
}
