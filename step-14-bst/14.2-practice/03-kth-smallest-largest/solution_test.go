package kthsmallestlargest

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestKth(t *testing.T) {
	r := bst.Build([]int{5, 3, 6, 2, 4, 1})
	for k := 1; k <= 6; k++ {
		if KthSmallest(r, k) != k || KthLargest(r, k) != 7-k {
			t.Errorf("k=%d", k)
		}
	}
}
