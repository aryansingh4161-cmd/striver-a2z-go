package levelorder

import (
	"reflect"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestLevelOrder(t *testing.T) {
	r := tree.FromLevel([]int{3, 9, 20, tree.Null, tree.Null, 15, 7})
	if got := LevelOrder(r); !reflect.DeepEqual(got, [][]int{{3}, {9, 20}, {15, 7}}) {
		t.Errorf("got %v", got)
	}
	if got := LevelOrder(nil); len(got) != 0 {
		t.Errorf("nil: %v", got)
	}
}
