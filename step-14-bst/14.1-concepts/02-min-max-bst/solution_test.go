package minmaxbst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestMinMax(t *testing.T) {
	r := bst.Build([]int{8, 3, 10, 1, 6, 14, 4, 7, 13})
	if Min(r) != 1 || Max(r) != 14 {
		t.Error("min/max")
	}
	one := bst.Build([]int{5})
	if Min(one) != 5 || Max(one) != 5 {
		t.Error("single")
	}
}
