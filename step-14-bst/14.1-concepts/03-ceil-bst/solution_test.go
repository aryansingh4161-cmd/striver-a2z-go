package ceilbst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestCeil(t *testing.T) {
	r := bst.Build([]int{8, 4, 12, 2, 6, 10, 14})
	for _, tc := range [][2]int{{11, 12}, {6, 6}, {15, -1}, {1, 2}, {7, 8}} {
		if got := Ceil(r, tc[0]); got != tc[1] {
			t.Errorf("Ceil(%d) = %d", tc[0], got)
		}
	}
}
