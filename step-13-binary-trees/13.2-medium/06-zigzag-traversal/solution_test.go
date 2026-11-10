package zigzagtraversal

import (
	"reflect"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestZigzag(t *testing.T) {
	N := tree.Null
	if got := Zigzag(tree.FromLevel([]int{3, 9, 20, N, N, 15, 7})); !reflect.DeepEqual(got, [][]int{{3}, {20, 9}, {15, 7}}) {
		t.Errorf("got %v", got)
	}
	if got := Zigzag(tree.FromLevel([]int{1, 2, 3, 4, 5, 6, 7})); !reflect.DeepEqual(got, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}) {
		t.Errorf("got %v", got)
	}
}
