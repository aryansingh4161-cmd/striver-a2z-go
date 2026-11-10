package floorbst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestFloor(t *testing.T) {
	r := bst.Build([]int{8, 4, 12, 2, 6, 10, 14})
	for _, tc := range [][2]int{{11, 10}, {6, 6}, {1, -1}, {15, 14}, {5, 4}} {
		if got := Floor(r, tc[0]); got != tc[1] {
			t.Errorf("Floor(%d) = %d", tc[0], got)
		}
	}
}
