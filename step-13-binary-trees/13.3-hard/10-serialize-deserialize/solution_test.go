package serializedeserialize

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestRoundTrip(t *testing.T) {
	N := tree.Null
	for _, lvl := range [][]int{{1, 2, 3, N, N, 4, 5}, {}, {1}, {-5, N, 7, 8}, {1, 2, 3, 4, 5, 6, 7}} {
		r := tree.FromLevel(lvl)
		if got := tree.ToLevel(Deserialize(Serialize(r))); !slices.Equal(got, lvl) {
			t.Errorf("level round trip %v → %v", lvl, got)
		}
		if got := tree.ToLevel(DeserializePre(SerializePre(r))); !slices.Equal(got, lvl) {
			t.Errorf("preorder round trip %v → %v", lvl, got)
		}
	}
	if s := Serialize(tree.FromLevel([]int{1, 2, 3, N, N, 4, 5})); s != "1,2,3,#,#,4,5,#,#,#,#" {
		t.Errorf("format %q", s)
	}
}
